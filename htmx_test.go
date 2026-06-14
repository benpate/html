package html

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHtmxAttributes(t *testing.T) {

	test := func(expected string, apply func(e *Element)) {
		t.Helper()
		b := New()
		apply(b.Div())
		require.Equal(t, expected, b.String())
	}

	test(`<div hx-get="/load"></div>`, func(e *Element) { e.HxGet("/load") })
	test(`<div hx-post="/save"></div>`, func(e *Element) { e.HxPost("/save") })
	test(`<div hx-target="#result"></div>`, func(e *Element) { e.HxTarget("#result") })
	test(`<div hx-trigger="click"></div>`, func(e *Element) { e.HxTrigger("click") })
	test(`<div hx-swap="outerHTML"></div>`, func(e *Element) { e.HxSwap("outerHTML") })
}

func TestHtmx_Combined(t *testing.T) {

	b := New()
	b.Button().HxGet("/data").HxTarget("#out").HxSwap("innerHTML").InnerText("Load")

	require.Equal(t, `<button hx-get="/data" hx-target="#out" hx-swap="innerHTML">Load</button>`, b.String())
}
