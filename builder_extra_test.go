package html

import (
	"io"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBuilder_Space(t *testing.T) {
	b := New()
	b.WriteString("a")
	b.Space()
	b.WriteString("b")
	require.Equal(t, "a b", b.String())
}

func TestBuilder_Bytes(t *testing.T) {
	b := New()
	b.Div().InnerText("hi")
	require.Equal(t, []byte(`<div>hi</div>`), b.Bytes())
}

func TestBuilder_Reader(t *testing.T) {
	b := New()
	b.Span().InnerText("content")

	data, err := io.ReadAll(b.Reader())
	require.NoError(t, err)
	require.Equal(t, `<span>content</span>`, string(data))
}

func TestBuilder_SubTree(t *testing.T) {

	b := New()
	b.Div() // open a container

	// A subtree shares the buffer but has its own element stack
	sub := b.SubTree()
	sub.Span().InnerText("inner")
	sub.CloseAll()

	// The parent builder can still close its own div
	require.Equal(t, `<div><span>inner</span></div>`, b.String())
}

func TestBuilder_EndBracket_NilStack(t *testing.T) {
	// EndBracket on an empty stack is a safe no-op
	b := New()
	require.NotPanics(t, func() { b.EndBracket() })
	require.Equal(t, "", b.String())
}

func TestBuilder_Close_NilStack(t *testing.T) {
	// Close on an empty stack is a safe no-op
	b := New()
	require.NotPanics(t, func() { b.Close() })
	require.Equal(t, "", b.String())
}

func TestBuilder_CloseAll(t *testing.T) {

	// CloseAll closes every open container on the stack
	b := New()
	b.HTML()
	b.Body()
	b.Div()
	b.CloseAll()

	require.Equal(t, `<html><body><div></div></body></html>`, b.String())
}

func TestBuilder_ReadString(t *testing.T) {

	// ReadString returns the current buffer WITHOUT closing open tags, leaving the
	// element stack intact so a later subroutine can fill in the body.
	b := New()
	b.HTML()
	b.Head().InnerHTML("<title>x</title>")

	// The header is read out exactly as written, with <html> still OPEN: it must
	// NOT contain a closing </html>, because ReadString does not close tags.
	header := b.ReadString()
	require.Equal(t, "<html><head><title>x</title></head>", header)
	require.NotContains(t, header, "</html>")

	// Because <html> is still open on the stack, the body nests inside it, and the
	// final String() closes <html> last -- producing well-formed, nested markup.
	b.Body()
	require.Equal(t, "<body></body></html>", b.String())
}

func TestBuilder_ReadString_MidTag(t *testing.T) {

	// Edge case: a freshly opened element has written "<div" but not yet its ">"
	// (the bracket is deferred until an attribute, child, or close). ReadString
	// reads the raw buffer as-is, so it returns the half-written tag.
	b := New()
	b.Div() // writes "<div", no ">" yet

	require.Equal(t, "<div", b.ReadString())

	// The element is still open, so EndBracket finishes the tag in the next chunk.
	b.EndBracket()
	require.Equal(t, "></div>", b.String())
}

func TestBuilder_ReadString_PreservesStackForNesting(t *testing.T) {

	// Regression: ReadString previously called the overridden String(), which ran
	// CloseAll() and emptied the stack -- so subsequent content was appended after
	// the closed tags instead of nested inside them.
	b := New()
	b.Div().EndBracket() // open a container and finish its tag, but leave it open

	read := b.ReadString()
	require.Equal(t, "<div>", read)

	// The div is still open, so the span nests inside it.
	b.Span().InnerText("inner")
	require.Equal(t, `<span>inner</span></div>`, b.String())
}
