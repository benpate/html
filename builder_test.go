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

func TestPartialRead(t *testing.T) {

	b := New()

	// EndBracket opens the <h1> but does not close it, and ReadString does not
	// close open tags -- so the partial read returns just "<h1>".
	b.H1().EndBracket()
	require.Equal(t, "<h1>", b.ReadString())
}

func TestRead(t *testing.T) {

	b := New()

	// The <h1> stays open across ReadString, so content written afterward lands
	// inside it and the final String() closes it.
	b.H1().EndBracket()
	require.Equal(t, "<h1>", b.ReadString())

	b.WriteString("hello world")
	require.Equal(t, "hello world</h1>", b.String())
}
