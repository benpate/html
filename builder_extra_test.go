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

	// ReadString returns the current buffer WITHOUT closing open tags,
	// and resets the buffer for further writing.
	b := New()
	b.HTML()
	b.Head().InnerHTML("<title>x</title>")

	header := b.ReadString()
	require.Contains(t, header, "<html>")
	require.Contains(t, header, "<title>x</title>")

	// After ReadString, the buffer is empty and we can keep building
	b.Body()
	rest := b.String()
	require.Contains(t, rest, "<body>")
	require.NotContains(t, rest, "<html>") // already read out
}
