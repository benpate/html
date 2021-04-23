package html

// Class adds a "class" attribute to the Element
func (element *Element) Class(value string) *Element {
	return element.Attr("class", value)
}

// For adds a "for" attribute to the Element
func (element *Element) For(value string) *Element {
	return element.Attr("for", value)
}

// ID adds an "id" attribute to the Element
func (element *Element) ID(value string) *Element {
	return element.Attr("id", value)
}

// Name adds a "name" attribute to the Element
func (element *Element) Name(value string) *Element {
	return element.Attr("name", value)
}

// Script adds a "data-script" attribute to the Element (hyperscript)
func (element *Element) Script(value string) *Element {
	return element.Attr("data-script", value)
}

// Type adds a "type" attribute to the Element
func (element *Element) Type(value string) *Element {
	return element.Attr("type", value)
}

// Value adds a "value" attribute to the Element
func (element *Element) Value(value string) *Element {
	return element.Attr("value", value)
}
