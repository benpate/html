package html

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestContainer(t *testing.T) {

	b := New()

	b.Div().Class("first")
	b.Div().Class("second").InnerHTML("Hey-oh")

	require.Equal(t, `<div class="first"><div class="second">Hey-oh</div></div>`, b.String())
}
