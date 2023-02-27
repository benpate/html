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

func TestInnerHTML(t *testing.T) {

	b := New()
	b.Div().Class("first").InnerHTML("Hello, World!")
	require.Equal(t, `<div class="first">Hello, World!</div>`, b.String())
}

func TestInnerHTML_WithHTML(t *testing.T) {

	b := New()
	b.Div().Class("first").InnerHTML("Hello<br>World!")
	require.Equal(t, `<div class="first">Hello<br>World!</div>`, b.String())
}

func TestInnerText(t *testing.T) {

	b := New()
	b.Div().Class("first").InnerText("Hello, World!")
	require.Equal(t, `<div class="first">Hello, World!</div>`, b.String())
}

func TestInnerText_WithHTML(t *testing.T) {

	b := New()
	b.Div().Class("first").InnerText("Hello<br>World!")
	require.Equal(t, `<div class="first">Hello&lt;br&gt;World!</div>`, b.String())
}
