package html

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// TestTags exercises every tag-building convenience method and confirms the
// exact HTML it produces.
func TestTags(t *testing.T) {

	test := func(expected string, build func(b *Builder)) {
		t.Helper()
		b := New()
		build(b)
		require.Equal(t, expected, b.String())
	}

	test(`<a href="http://x.com"></a>`, func(b *Builder) { b.A("http://x.com") })
	test(`<audio></audio>`, func(b *Builder) { b.Audio() })
	test(`<b class="bold"></b>`, func(b *Builder) { b.B("bold") })
	test(`<body></body>`, func(b *Builder) { b.Body() })
	test(`<br>`, func(b *Builder) { b.BR() })
	test(`<button></button>`, func(b *Builder) { b.Button() })
	test(`<datalist id="dl"></datalist>`, func(b *Builder) { b.Datalist("dl") })
	test(`<div></div>`, func(b *Builder) { b.Div() })
	test(`<figure></figure>`, func(b *Builder) { b.Figure() })
	test(`<figcaption>caption</figcaption>`, func(b *Builder) { b.FigCaption("caption") })
	test(`<form method="post" action="/submit"></form>`, func(b *Builder) { b.Form("post", "/submit") })
	test(`<h1></h1>`, func(b *Builder) { b.H1() })
	test(`<h2></h2>`, func(b *Builder) { b.H2() })
	test(`<h3></h3>`, func(b *Builder) { b.H3() })
	test(`<head></head>`, func(b *Builder) { b.Head() })
	test(`<html></html>`, func(b *Builder) { b.HTML() })
	test(`<i class="fa fa-star"></i>`, func(b *Builder) { b.I("fa", "fa-star") })
	test(`<img src="pic.jpg"></img>`, func(b *Builder) { b.Img("pic.jpg") })
	test(`<input type="text" name="field">`, func(b *Builder) { b.Input("text", "field") })
	test(`<label for="field"></label>`, func(b *Builder) { b.Label("field") })
	test(`<link rel="stylesheet" href="style.css">`, func(b *Builder) { b.Link("stylesheet", "style.css") })
	test(`<optgroup label="Group"></optgroup>`, func(b *Builder) { b.OptGroup("Group") })
	test(`<option value="v">Label</option>`, func(b *Builder) { b.Option("Label", "v") })
	test(`<option value="v" selected="true">Label</option>`, func(b *Builder) { b.OptionSelected("Label", "v") })
	test(`<picture></picture>`, func(b *Builder) { b.Picture() })
	test(`<script></script>`, func(b *Builder) { b.Script() })
	test(`<link rel="stylesheet" href="style.css">`, func(b *Builder) { b.Stylesheet("style.css") })
	test(`<select name="choice"></select>`, func(b *Builder) { b.Select("choice") })
	test(`<span></span>`, func(b *Builder) { b.Span() })
	test(`<source></source>`, func(b *Builder) { b.Source() })
	test(`<textarea name="bio"></textarea>`, func(b *Builder) { b.Textarea("bio") })
	test(`<title>My Page</title>`, func(b *Builder) { b.Title("My Page") })
	test(`<table></table>`, func(b *Builder) { b.Table() })
	test(`<tbody></tbody>`, func(b *Builder) { b.TBody() })
	test(`<td></td>`, func(b *Builder) { b.TD() })
	test(`<th></th>`, func(b *Builder) { b.TH() })
	test(`<tr></tr>`, func(b *Builder) { b.TR() })
}

func TestTags_Nested(t *testing.T) {

	// A representative nested structure exercises the auto-close stack behavior.
	b := New()
	b.Table()
	b.TR()
	b.TD().InnerText("cell")

	require.Equal(t, `<table><tr><td>cell</td></tr></table>`, b.String())
}

func TestTitle_EscapesContent(t *testing.T) {
	b := New()
	b.Title("A & B")
	require.Equal(t, `<title>A &amp; B</title>`, b.String())
}
