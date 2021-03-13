package htmlbuilder

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

// Value adds a "value" attribute to the Element
func (element *Element) Value(value string) *Element {
	return element.Attr("value", value)
}
