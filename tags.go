package html

func (b *Builder) A(href string) *Element {
	return b.Container("a").Attr("href", href)
}

func (b *Builder) Audio() *Element {
	return b.Container("audio")
}

func (b *Builder) B(class string) *Element {
	return b.Container("b").Class(class)
}

func (b *Builder) Body() *Element {
	return b.Container("body")
}

func (b *Builder) BR() *Element {
	return b.Empty("br")
}

func (b *Builder) Button() *Element {
	return b.Container("button")
}

func (b *Builder) Datalist(id string) *Element {
	return b.Container("datalist").ID(id).EndBracket()
}

func (b *Builder) Div() *Element {
	return b.Container("div")
}

func (b *Builder) Form(method string, action string) *Element {
	return b.Container("form").Attr("method", method).Attr("action", action)
}

func (b *Builder) H1() *Element {
	return b.Container("h1")
}

func (b *Builder) H2() *Element {
	return b.Container("h2")
}
func (b *Builder) H3() *Element {
	return b.Container("h3")
}

func (b *Builder) Head() *Element {
	return b.Container("head")
}

func (b *Builder) HTML() *Element {
	return b.Container("html")
}

func (b *Builder) I(classes ...string) *Element {
	return b.Container("i").Class(classes...)
}

func (b *Builder) Img(src string) *Element {
	return b.Container("img").Src(src)
}

func (b *Builder) Input(t string, name string) *Element {
	return b.Empty("input").Type(t).Name(name)
}

func (b *Builder) Label(forID string) *Element {
	return b.Container("label").For(forID)
}

func (b *Builder) Link(rel string, href string) *Element {
	return b.Empty("link").Attr("rel", rel).Attr("href", href).Close()
}

func (b *Builder) OptGroup(label string) *Element {
	return b.Container("optgroup").Label(label).EndBracket()
}

func (b *Builder) Option(label string, value string) *Element {
	return b.Container("option").ForceAttr("value", value).InnerText(label).Close()
}

func (b *Builder) OptionSelected(label string, value string) *Element {
	return b.Container("option").ForceAttr("value", value).Attr("selected", "true").InnerText(label).Close()
}

func (b *Builder) Script() *Element {
	return b.Container("script")
}

func (b *Builder) Stylesheet(url string) *Element {
	return b.Empty("link").Rel("stylesheet").Href(url)
}

func (b *Builder) Select(name string) *Element {
	return b.Container("select").Name(name)
}

func (b *Builder) Span() *Element {
	return b.Container("span")
}

func (b *Builder) Source() *Element {
	return b.Container("source")
}

func (b *Builder) Textarea(name string) *Element {
	return b.Container("textarea").Name(name)
}

func (b *Builder) Title(value string) *Element {
	return b.Container("title").InnerText(value).Close()
}

func (b *Builder) Table() *Element {
	return b.Container("table")
}

func (b *Builder) TBody() *Element {
	return b.Container("tbody")
}

func (b *Builder) TD() *Element {
	return b.Container("td")
}

func (b *Builder) TH() *Element {
	return b.Container("th")
}

func (b *Builder) TR() *Element {
	return b.Container("tr")
}
