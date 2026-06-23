package html

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestElement_Attr_SkipsEmpty(t *testing.T) {
	// Attr skips empty values; ForceAttr writes them
	b := New()
	b.Span().Attr("data-x", "")
	require.Equal(t, `<span></span>`, b.String())
}

func TestElement_ForceAttr_DoesntSkipEmpty(t *testing.T) {
	// Attr skips empty values; ForceAttr writes them
	b := New()
	b.Span().ForceAttr("data-x", "")
	require.Equal(t, `<span data-x=""></span>`, b.String())
}

func TestElement_ForceAttr_AfterEndBracket(t *testing.T) {

	// Once the end bracket is written, no further attributes can be added
	b := New()
	e := b.Span()
	e.EndBracket()
	e.ForceAttr("late", "value") // should be a no-op

	require.Equal(t, `<span></span>`, b.String())
}

func TestElement_InnerHTML_OnClosedElement(t *testing.T) {

	// Adding content to an already-closed element is a no-op
	b := New()
	e := b.Span()
	e.Close()
	e.InnerHTML("ignored")

	require.Equal(t, `<span></span>`, b.String())
}

func TestElement_Close_Idempotent(t *testing.T) {

	// Closing an already-closed element is a no-op: it must not write a second
	// closing tag, regardless of how many times it is called.
	b := New()
	e := b.Span()
	e.Close()
	e.Close()
	e.Close()

	require.Equal(t, `<span></span>`, b.String())
}

func TestElement_InnerHTML_Empty(t *testing.T) {

	// InnerHTML with empty content still closes the element
	b := New()
	b.Div().InnerHTML("")
	require.Equal(t, `<div></div>`, b.String())
}

func TestElement_EndBracket_Idempotent(t *testing.T) {

	// Calling EndBracket twice writes only one ">"
	b := New()
	e := b.Div()
	e.EndBracket()
	e.EndBracket()
	require.Equal(t, `<div></div>`, b.String())
}

func TestElement_NonContainer_NoClosingTag(t *testing.T) {

	// An "empty" (non-container) element never gets a closing tag
	b := New()
	b.Empty("br")
	require.Equal(t, `<br>`, b.String())
}

func TestElement_NonContainer_EndBracketPopsStack(t *testing.T) {

	// Regression: EndBracket on a non-container element must pop it off the stack.
	// Previously it marked the element closed but left it on the stack, so the
	// following String()/CloseAll() spun forever on an element it could not pop.
	b := New()
	b.Empty("br").EndBracket()

	require.NotPanics(t, func() {
		require.Equal(t, `<br>`, b.String())
	})
}

func TestElement_InnerText_Escapes(t *testing.T) {

	b := New()
	b.Div().InnerText(`<script>alert("x")</script>`)
	require.Equal(t, `<div>&lt;script&gt;alert(&#34;x&#34;)&lt;/script&gt;</div>`, b.String())
}
