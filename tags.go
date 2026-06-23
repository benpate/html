package html

// A creates a new anchor element with the specified href
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/a
func (builder *Builder) A(href string) *Element {
	return builder.Container("a").Attr("href", href)
}

// Audio creates a new audio element
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/audio
func (builder *Builder) Audio() *Element {
	return builder.Container("audio")
}

// B creates a new bold element with the specified class
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/b
func (builder *Builder) B(class string) *Element {
	return builder.Container("b").Class(class)
}

// Body creates a new body element
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/body
func (builder *Builder) Body() *Element {
	return builder.Container("body")
}

// BR creates a new line break element
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/br
func (builder *Builder) BR() *Element {
	return builder.Empty("br")
}

// Button creates a new button element
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/button
func (builder *Builder) Button() *Element {
	return builder.Container("button")
}

// Datalist creates a new datalist element with the specified ID
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/datalist
func (builder *Builder) Datalist(id string) *Element {
	return builder.Container("datalist").ID(id).EndBracket()
}

// Div creates a new div element
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/div
func (builder *Builder) Div() *Element {
	return builder.Container("div")
}

// Figure creates a new figure element
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/figure
func (builder *Builder) Figure() *Element {
	return builder.Container("figure")
}

// FigCaption creates a new figcaption element with the specified inner HTML
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/figcaption
func (builder *Builder) FigCaption(innerHTML string) *Element {
	return builder.Container("figcaption").InnerHTML(innerHTML)
}

// Form creates a new form element
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/form
func (builder *Builder) Form(method string, action string) *Element {
	return builder.Container("form").Attr("method", method).Attr("action", action)
}

// H1 creates a new h1 element
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/h1
func (builder *Builder) H1() *Element {
	return builder.Container("h1")
}

// H2 creates a new h2 element
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/h2
func (builder *Builder) H2() *Element {
	return builder.Container("h2")
}

// H3 creates a new h3 element
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/h3
func (builder *Builder) H3() *Element {
	return builder.Container("h3")
}

// Head creates a new head element
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/head
func (builder *Builder) Head() *Element {
	return builder.Container("head")
}

// HTML creates a new html element
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/html
func (builder *Builder) HTML() *Element {
	return builder.Container("html")
}

// I creates a new italics element
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/i
func (builder *Builder) I(classes ...string) *Element {
	return builder.Container("i").Class(classes...)
}

// Img creates a new image element
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/img
func (builder *Builder) Img(src string) *Element {
	return builder.Empty("img").Src(src)
}

// Input creates a new input element
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/input
func (builder *Builder) Input(t string, name string) *Element {
	return builder.Empty("input").Type(t).Name(name)
}

// Label creates a new label element
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/label
func (builder *Builder) Label(forID string) *Element {
	return builder.Container("label").For(forID)
}

// Link creates a new link element
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/link
func (builder *Builder) Link(rel string, href string) *Element {
	return builder.Empty("link").Attr("rel", rel).Attr("href", href).Close()
}

// OptGroup creates a new optgroup element
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/optgroup
func (builder *Builder) OptGroup(label string) *Element {
	return builder.Container("optgroup").Label(label).EndBracket()
}

// Option creates a new option element
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/option
func (builder *Builder) Option(label string, value string) *Element {
	return builder.Container("option").ForceAttr("value", value).InnerText(label)
}

// OptionSelected creates a new option element with the "selected" attribute set to true
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/option
func (builder *Builder) OptionSelected(label string, value string) *Element {
	return builder.Container("option").ForceAttr("value", value).Attr("selected", "true").InnerText(label)
}

// Picture creates a new picture element
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/picture
func (builder *Builder) Picture() *Element {
	return builder.Container("picture")
}

// Script creates a new script element
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/script
func (builder *Builder) Script() *Element {
	return builder.Container("script")
}

// Stylesheet creates a new link element for a stylesheet
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/link
func (builder *Builder) Stylesheet(url string) *Element {
	return builder.Empty("link").Rel("stylesheet").Href(url)
}

// Select creates a new select element
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/select
func (builder *Builder) Select(name string) *Element {
	return builder.Container("select").Name(name)
}

// Span creates a new span element
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/span
func (builder *Builder) Span() *Element {
	return builder.Container("span")
}

// Source creates a new source element
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/source
func (builder *Builder) Source() *Element {
	return builder.Empty("source")
}

// Textarea creates a new textarea element
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/textarea
func (builder *Builder) Textarea(name string) *Element {
	return builder.Container("textarea").Name(name)
}

// Title creates a new title element
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/title
func (builder *Builder) Title(value string) *Element {
	return builder.Container("title").InnerText(value)
}

// Table creates a new table element
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/table
func (builder *Builder) Table() *Element {
	return builder.Container("table")
}

// TBody creates a new table body element
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/tbody
func (builder *Builder) TBody() *Element {
	return builder.Container("tbody")
}

// TD creates a new table data element
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/td
func (builder *Builder) TD() *Element {
	return builder.Container("td")
}

// TH creates a new table header element
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/th
func (builder *Builder) TH() *Element {
	return builder.Container("th")
}

// TR creates a new table row element
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/tr
func (builder *Builder) TR() *Element {
	return builder.Container("tr")
}
