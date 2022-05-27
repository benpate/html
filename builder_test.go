package html

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAutoEndBracket(t *testing.T) {

	b := New()
	b.BR()
	b.BR()

	require.NotNil(t, b.last)
	assert.Nil(t, b.last.parent)
}

func TestRead(t *testing.T) {

	b := New()

	b.H1().EndBracket()
	require.Equal(t, "<h1>", b.ReadString())

	b.WriteString("hello world")
	require.Equal(t, "hello world</h1>", b.String())
}
