package html

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// TestAttributes exercises every attribute helper against a <span> element.
func TestAttributes(t *testing.T) {

	test := func(expected string, apply func(e *Element)) {
		t.Helper()
		b := New()
		apply(b.Span())
		require.Equal(t, expected, b.String())
	}

	test(`<span aria-label="Hello"></span>`, func(e *Element) { e.Aria("label", "Hello") })
	test(`<span class="a b c"></span>`, func(e *Element) { e.Class("a", "b", "c") })
	test(`<span for="field"></span>`, func(e *Element) { e.For("field") })
	test(`<span data-id="5"></span>`, func(e *Element) { e.Data("id", "5") })
	test(`<span href="http://x.com"></span>`, func(e *Element) { e.Href("http://x.com") })
	test(`<span id="main"></span>`, func(e *Element) { e.ID("main") })
	test(`<span label="My Label"></span>`, func(e *Element) { e.Label("My Label") })
	test(`<span list="datalist1"></span>`, func(e *Element) { e.List("datalist1") })
	test(`<span media="screen"></span>`, func(e *Element) { e.Media("screen") })
	test(`<span name="field"></span>`, func(e *Element) { e.Name("field") })
	test(`<span rel="next"></span>`, func(e *Element) { e.Rel("next") })
	test(`<span role="button"></span>`, func(e *Element) { e.Role("button") })
	test(`<span data-script="on click add .foo"></span>`, func(e *Element) { e.Script("on click", "add .foo") })
	test(`<span src="pic.jpg"></span>`, func(e *Element) { e.Src("pic.jpg") })
	test(`<span srcset="pic-2x.jpg 2x"></span>`, func(e *Element) { e.SrcSet("pic-2x.jpg 2x") })
	test(`<span style="color:red; font-weight:bold"></span>`, func(e *Element) { e.Style("color:red", "font-weight:bold") })
	test(`<span tabIndex="0"></span>`, func(e *Element) { e.TabIndex("0") })
	test(`<span type="text"></span>`, func(e *Element) { e.Type("text") })
}

func TestAttribute_Value_ForcesEmpty(t *testing.T) {

	// Value() always writes, even when empty
	b := New()
	b.Span().Value("")
	require.Equal(t, `<span value=""></span>`, b.String())

	b2 := New()
	b2.Span().Value("hello")
	require.Equal(t, `<span value="hello"></span>`, b2.String())
}

func TestAttribute_EscapesValue(t *testing.T) {

	// Attribute values are HTML-escaped
	b := New()
	b.Span().ID(`a"&<b`)
	require.Equal(t, `<span id="a&#34;&amp;&lt;b"></span>`, b.String())
}

func TestAttribute_MultipleAttributes(t *testing.T) {

	b := New()
	b.Div().ID("box").Class("a", "b").Data("x", "1")
	require.Equal(t, `<div id="box" class="a b" data-x="1"></div>`, b.String())
}
