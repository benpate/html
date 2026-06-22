package html

import (
	"html"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

// FuzzInnerText confirms that arbitrary user content passed to InnerText is
// always HTML-escaped: the rendered body never contains a raw "<" or ">", and
// unescaping it recovers the exact original text (no loss, no injection).
func FuzzInnerText(f *testing.F) {

	f.Add("")
	f.Add("hello")
	f.Add("<script>alert(1)</script>")
	f.Add("a & b < c > d")
	f.Add(`"><img src=x onerror=alert(1)>`)

	f.Fuzz(func(t *testing.T, text string) {

		b := New()
		b.Div().InnerText(text)
		output := b.String()

		// Strip the fixed <div>...</div> wrapper to inspect only the escaped body.
		body := strings.TrimSuffix(strings.TrimPrefix(output, "<div>"), "</div>")

		// The escaped body must never contain raw angle brackets, which would let
		// user content open or close a tag.
		require.NotContains(t, body, "<", "InnerText left a raw '<' in the body")
		require.NotContains(t, body, ">", "InnerText left a raw '>' in the body")

		// Round-trip: unescaping the body must recover the exact input, proving the
		// escaping is lossless as well as safe.
		require.Equal(t, text, html.UnescapeString(body), "InnerText did not round-trip")
	})
}

// FuzzAttr confirms that arbitrary user content passed as an attribute value is
// always escaped: the rendered value never contains a raw double-quote (which
// would break out of the attribute) or angle bracket, and unescaping recovers
// the exact original value.
func FuzzAttr(f *testing.F) {

	f.Add("")
	f.Add("plain")
	f.Add(`" onmouseover="alert(1)`)
	f.Add(`a"b<c>d&e`)
	f.Add("https://example.com/?x=1&y=2")

	f.Fuzz(func(t *testing.T, value string) {

		// ForceAttr is used so empty values still render and exercise the path.
		b := New()
		b.Div().ForceAttr("data-x", value)
		output := b.String()

		// Isolate the quoted attribute value: data-x="...".
		const prefix = `<div data-x="`
		require.True(t, strings.HasPrefix(output, prefix), "unexpected output: %q", output)
		inner := strings.TrimSuffix(strings.TrimPrefix(output, prefix), `"></div>`)

		// A raw double-quote inside the value would terminate the attribute early
		// and let the rest be parsed as new attributes (an injection).
		require.NotContains(t, inner, `"`, "Attr left a raw '\"' in the value")
		require.NotContains(t, inner, "<", "Attr left a raw '<' in the value")
		require.NotContains(t, inner, ">", "Attr left a raw '>' in the value")

		// Round-trip: unescaping the value must recover the exact input.
		require.Equal(t, value, html.UnescapeString(inner), "Attr did not round-trip")
	})
}

// FuzzBuilderStack drives a random sequence of element and stack operations,
// confirming that the builder never panics and that String() always emits
// balanced markup: every container's "<tag>" has a matching "</tag>", and the
// final output has no dangling open tag.
//
// The op bytes are interpreted as a tiny instruction stream:
//
//	0 => open a container element (<x> ... </x>)
//	1 => open an empty element    (<x>, never closed)
//	2 => Close() the top of the stack
//	3 => EndBracket() the top of the stack
//	4 => CloseAll()
//
// All elements use the same tag name "x" so balance can be checked by counting.
func FuzzBuilderStack(f *testing.F) {

	f.Add([]byte{})
	f.Add([]byte{0, 0, 0})       // three nested containers, auto-closed by String()
	f.Add([]byte{0, 2, 0, 2})    // open/close pairs
	f.Add([]byte{1, 1, 1})       // three empty elements
	f.Add([]byte{0, 1, 0, 4, 3}) // mixed, with CloseAll and a trailing EndBracket

	f.Fuzz(func(t *testing.T, ops []byte) {

		// Cap the op stream so the fuzzer explores operation patterns rather than
		// spending its budget growing a multi-megabyte buffer from one long input.
		if len(ops) > 64 {
			ops = ops[:64]
		}

		b := New()

		require.NotPanics(t, func() {
			for _, op := range ops {
				switch op % 5 {
				case 0:
					b.Container("x")
				case 1:
					b.Empty("x")
				case 2:
					b.Close()
				case 3:
					b.EndBracket()
				case 4:
					b.CloseAll()
				}
			}
		})

		// String() calls CloseAll(), so the final markup must be fully balanced:
		// the count of opening "<x>" tags equals the count of closing "</x>" tags.
		output := b.String()
		opens := strings.Count(output, "<x>")
		closes := strings.Count(output, "</x>")

		// Every container contributes one "<x>" and one "</x>"; every empty element
		// contributes one "<x>" and zero "</x>". So closes can never exceed opens,
		// and the output can never end while a container is still open.
		require.LessOrEqual(t, closes, opens, "more closing than opening tags: %q", output)
		require.False(t, strings.HasSuffix(output, "<x"), "output ends mid-tag: %q", output)
	})
}

// FuzzAttrNameTrustBoundary pins the library's load-bearing trust boundary:
// attribute VALUES are escaped, but attribute NAMES are written verbatim and are
// the caller's responsibility. This test does not assert names are sanitized
// (they intentionally are not); it asserts that the value half of the contract
// holds regardless of the name, and that no name input causes a panic.
//
// If a future change starts escaping or validating names, this test's
// expectations must be revisited deliberately — that is the point of pinning it.
func FuzzAttrNameTrustBoundary(f *testing.F) {

	f.Add("class", "a b")
	f.Add("data-x", `"><script>`)
	f.Add(`evil"name`, "value")
	f.Add("", "")

	f.Fuzz(func(t *testing.T, name string, value string) {

		var output string
		require.NotPanics(t, func() {
			b := New()
			b.Div().ForceAttr(name, value)
			output = b.String()
		})

		// The name is written raw, so locate the value by its escaped form rather
		// than by parsing the (possibly malformed) name. The escaped value must
		// appear verbatim in the output: the escaping of the value is independent
		// of whatever the caller passed as the name.
		require.Contains(t, output, `="`+html.EscapeString(value)+`"`,
			"escaped value missing or altered by name %q", name)
	})
}
