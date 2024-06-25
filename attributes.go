package html

import "strings"

// Aria adds a "aria-*" attribute to the Element
func (element *Element) Aria(name string, value string) *Element {
	return element.Attr("aria-"+name, value)
}

// Class adds a "class" attribute to the Element.
// Multiple class names can be passed, and are separated by spaces
func (element *Element) Class(values ...string) *Element {
	return element.Attr("class", strings.Join(values, " "))
}

// For adds a "for" attribute to the Element
func (element *Element) For(value string) *Element {
	return element.Attr("for", value)
}

// Data adds a "data-" attribute to the Element
func (element *Element) Data(name string, value string) *Element {
	return element.Attr("data-"+name, value)
}

// Href adds an "href" attribute to the Element
func (element *Element) Href(url string) *Element {
	return element.Attr("href", url)
}

// ID adds an "id" attribute to the Element
func (element *Element) ID(value string) *Element {
	return element.Attr("id", value)
}

// Label adds a "label" attribute to the Element
func (element *Element) Label(value string) *Element {
	return element.Attr("label", value)
}

// List adds a "list" attribute to the Element
func (element *Element) List(value string) *Element {
	return element.Attr("list", value)
}

// Media adds a "name" attribute to the Element
func (element *Element) Media(value string) *Element {
	return element.Attr("media", value)
}

// Name adds a "name" attribute to the Element
func (element *Element) Name(value string) *Element {
	return element.Attr("name", value)
}

// Rel adds a "rel" (relationship) attribute to the Element (valid values listed at https://developer.mozilla.org/en-US/docs/Web/HTML/Attributes/rel)
func (element *Element) Rel(value string) *Element {
	return element.Attr("rel", value)
}

// Role adds a "role" attribute to the Element (for WAI-ARIA)
func (element *Element) Role(value string) *Element {
	return element.Attr("role", value)
}

// Script adds a "data-script" attribute to the Element (for https://hyperscript.org)
func (element *Element) Script(value string) *Element {
	return element.Attr("data-script", value)
}

// Src adds a "src" attribute to the Element
func (element *Element) Src(value string) *Element {
	return element.Attr("src", value)
}

// Srcset adds a "srcset" attribute to the Element
func (element *Element) Srcset(value string) *Element {
	return element.Attr("srcset", value)
}

// Stype adds a "style" attribute to the Element
// Multiple style definitions can be passed, and are separated by semicolons.
func (element *Element) Style(values ...string) *Element {
	return element.Attr("style", strings.Join(values, "; "))
}

// TabIndex adds a "tabIndex" attribute to the Element
func (element *Element) TabIndex(value string) *Element {
	return element.Attr("tabIndex", value)
}

// Type adds a "type" attribute to the Element
func (element *Element) Type(value string) *Element {
	return element.Attr("type", value)
}

// Value adds a "value" attribute to the Element
func (element *Element) Value(value string) *Element {
	return element.Attr("value", value)
}
