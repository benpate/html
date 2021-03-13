package builder

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
