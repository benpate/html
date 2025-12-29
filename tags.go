package html

// A creates a new anchor element with the specified href
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/a
func (b *Builder) A(href string) *Element {
	return b.Container("a").Attr("href", href)
}

// Audio creates a new abbreviation element
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/audio
func (b *Builder) Audio() *Element {
	return b.Container("audio")
}

// B creates a new bold element with the specified class
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/b
func (b *Builder) B(class string) *Element {
	return b.Container("b").Class(class)
}

// Body creates a new body element
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/body
func (b *Builder) Body() *Element {
	return b.Container("body")
}

// BR creates a new line break element
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/br
func (b *Builder) BR() *Element {
	return b.Empty("br")
}

// Button creates a new button element
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/button
func (b *Builder) Button() *Element {
	return b.Container("button")
}

// Datalist creates a new datalist element with the specified ID
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/datalist
func (b *Builder) Datalist(id string) *Element {
	return b.Container("datalist").ID(id).EndBracket()
}

// Div creates a new div element
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/div
func (b *Builder) Div() *Element {
	return b.Container("div")
}

// Figure creates a new figure element
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/figure
func (b *Builder) Figure() *Element {
	return b.Container("figure")
}

// FigCaption creates a new figcaption element with the specified inner HTML
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/figcaption
func (b *Builder) FigCaption(innerHTML string) *Element {
	return b.Container("figcaption").InnerHTML(innerHTML).Close()
}

// Form creates a new form element
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/form
func (b *Builder) Form(method string, action string) *Element {
	return b.Container("form").Attr("method", method).Attr("action", action)
}

// H1 creates a new h1 element
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/h1
func (b *Builder) H1() *Element {
	return b.Container("h1")
}

// H2 creates a new h2 element
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/h2
func (b *Builder) H2() *Element {
	return b.Container("h2")
}

// H3 creates a new h3 element
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/h3
func (b *Builder) H3() *Element {
	return b.Container("h3")
}

// Head creates a new head element
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/head
func (b *Builder) Head() *Element {
	return b.Container("head")
}

// HTML creates a new html element
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/html
func (b *Builder) HTML() *Element {
	return b.Container("html")
}

// I creates a new italics element
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/i
func (b *Builder) I(classes ...string) *Element {
	return b.Container("i").Class(classes...)
}

// Img creates a new image element
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/img
func (b *Builder) Img(src string) *Element {
	return b.Container("img").Src(src)
}

// Input creates a new input element
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/input
func (b *Builder) Input(t string, name string) *Element {
	return b.Empty("input").Type(t).Name(name)
}

// Label creates a new label element
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/label
func (b *Builder) Label(forID string) *Element {
	return b.Container("label").For(forID)
}

// Link creates a new link element
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/link
func (b *Builder) Link(rel string, href string) *Element {
	return b.Empty("link").Attr("rel", rel).Attr("href", href).Close()
}

// OptGroup creates a new meta element
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/optgroup
func (b *Builder) OptGroup(label string) *Element {
	return b.Container("optgroup").Label(label).EndBracket()
}

// Option creates a new option element
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/option
func (b *Builder) Option(label string, value string) *Element {
	return b.Container("option").ForceAttr("value", value).InnerText(label).Close()
}

// OptionSelected creates a new option element with the "selected" attribute set to true
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/option
func (b *Builder) OptionSelected(label string, value string) *Element {
	return b.Container("option").ForceAttr("value", value).Attr("selected", "true").InnerText(label).Close()
}

// Picture creates a new picture element
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/picture
func (b *Builder) Picture() *Element {
	return b.Container("picture")
}

// Script creates a new script element
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/script
func (b *Builder) Script() *Element {
	return b.Container("script")
}

// Stylesheet creates a new link element for a stylesheet
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/link
func (b *Builder) Stylesheet(url string) *Element {
	return b.Empty("link").Rel("stylesheet").Href(url)
}

// Select creates a new select element
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/select
func (b *Builder) Select(name string) *Element {
	return b.Container("select").Name(name)
}

// Span creates a new span element
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/span
func (b *Builder) Span() *Element {
	return b.Container("span")
}

// Source creates a new source element
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/source
func (b *Builder) Source() *Element {
	return b.Container("source")
}

// Textarea creates a new textarea element
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/textarea
func (b *Builder) Textarea(name string) *Element {
	return b.Container("textarea").Name(name)
}

// Title creates a new title element
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/title
func (b *Builder) Title(value string) *Element {
	return b.Container("title").InnerText(value).Close()
}

// Table creates a new table element
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/table
func (b *Builder) Table() *Element {
	return b.Container("table")
}

// TBody creates a new table body element
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/tbody
func (b *Builder) TBody() *Element {
	return b.Container("tbody")
}

// TD creates a new table data element
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/td
func (b *Builder) TD() *Element {
	return b.Container("td")
}

// TH creates a new table header element
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/th
func (b *Builder) TH() *Element {
	return b.Container("th")
}

// TR creates a new table row element
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/tr
func (b *Builder) TR() *Element {
	return b.Container("tr")
}
