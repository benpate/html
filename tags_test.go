package html

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDiv(t *testing.T) {

	b := New()
	div := b.Div()

	require.Equal(t, div.name, "div")
	require.True(t, div.container)
	require.False(t, div.closed)
	require.False(t, div.endBracket)

	div.EndBracket()
	require.True(t, div.endBracket)
	require.False(t, div.closed)

	div.Close()
	require.True(t, div.closed)

	require.Equal(t, "<div></div>", b.String())
}

func TestInput(t *testing.T) {

	b := New()
	div := b.Input("text", "FullName").Attr("value", "John Connor")

	require.Equal(t, div.name, "input")
	require.False(t, div.container)
	require.False(t, div.closed)
	require.False(t, div.endBracket)

	div.Close()
	require.True(t, div.endBracket)
	require.True(t, div.closed)

	require.Equal(t, `<input type="text" name="FullName" value="John Connor">`, b.String())
}

func TestOption(t *testing.T) {
	b := New()
	b.Option("John Connor", "1")

	require.Equal(t, `<option value="1">John Connor</option>`, b.String())
}
