package html

import (
	"html"
)

// Element represents a element that is being written into the provided strings.Builder
type Element struct {
	builder    *Builder
	parent     *Element
	name       string
	container  bool
	endBracket bool
	closed     bool
}

// Start writes the initial tag name and opening bracket for this tag.
func (element *Element) Start() *Element {

	var growSize int

	if element.container {
		growSize = 2*len(element.name) + 5 // "<" + name + ">" + "</" + name + ">"
	} else {
		growSize = len(element.name) + 2 // "<" + name + ">"
	}

	// Grow the buffer and write the new tag name
	element.builder.Grow(growSize)
	element.builder.WriteRune('<')
	element.builder.WriteString(element.name)

	return element
}

// Attr writes the attribute into the string builder.  It converts
// the value (second parameter) into a string, and then uses html.EscapeString
// to escape the attribute value.  Empty values are not written to the builder.
//
// The NAME is written into the tag verbatim, without escaping, so it must never
// come from user input: a name containing a quote closes the attribute, and one
// containing ">" closes the tag, letting the rest be parsed as markup.
func (element *Element) Attr(name string, value string) *Element {

	// RULE: An empty value writes no attribute at all.
	if value == "" {
		return element
	}

	// Otherwise, write the attribute to the builder
	return element.ForceAttr(name, value)
}

// ForceAttr writes the attribute into the string builder.  It converts
// the value (second parameter) into a string, and then uses html.EscapeString
// to escape the attribute value.  Empty values ARE written to the builder.
//
// As with Attr, the NAME is written verbatim and must never come from user
// input. Values are escaped; names are the caller's responsibility.
// #nosec G104 -- strings.Builder's Write methods always return a nil error
func (element *Element) ForceAttr(name string, value string) *Element {

	// RULE: Once the end bracket is written there is nowhere left to put an
	// attribute, so the call is silently ignored.
	if element.endBracket {
		return element
	}

	// escape the value
	value = html.EscapeString(value)

	// length of: space + name + quote + escaped value + quote
	element.builder.Grow(len(name) + len(value) + 4)

	// write values to the builder
	element.builder.WriteRune(' ')
	element.builder.WriteString(name)
	element.builder.WriteString(`="`)
	element.builder.WriteString(value)
	element.builder.WriteRune('"')

	return element
}

// EndBracket writes the final ">" of the beginning element to the strings.Builder
// It uses an internal variable to prevent duplicate calls
func (element *Element) EndBracket() *Element {

	// RULE: The end bracket is written exactly once.
	if element.endBracket {
		return element
	}

	element.endBracket = true
	element.builder.WriteRune('>')

	// A non-container element has no closing tag, so its end bracket closes it
	// permanently. Pop it off the stack now; otherwise CloseAll would spin on an
	// element that is closed but never removed.
	if !element.container {
		element.closed = true
		element.builder.last = element.parent
	}

	return element
}

// InnerHTML does three things:
// 1) closes the beginning element (if needed)
// 2) appends innerHTML (if provided)
// 3) writes an ending element to the builder (ie. </element> )
// InnerHTML DOES NOT escape its content, so it is not safe to use
// with user-generated content.
func (element *Element) InnerHTML(innerHTML string) *Element {

	// RULE: A closed element cannot take any more content.
	if element.closed {
		return element
	}

	// Only need to write additional content if innerHTML is not empty
	if innerHTML != "" {

		if !element.endBracket {
			element.builder.WriteRune('>')
			element.endBracket = true
		}

		// Write innerHTML (if present)
		element.builder.Grow(len(innerHTML))
		element.builder.WriteString(innerHTML)
	}

	// Write the remaining end element
	return element.Close()
}

// InnerText works similarly to the InnerHTML method.  In addition,
// it HTML escapes its content so that user-generated-content can
// be inserted into the output safely.
func (element *Element) InnerText(text string) *Element {
	return element.InnerHTML(html.EscapeString(text))
}

// Close writes the necessary closing tag for this element and marks it closed
func (element *Element) Close() *Element {

	// RULE: Closing twice writes a second end tag, so do it once.
	if element.closed {
		return element
	}

	// Mark this element as closed.
	element.EndBracket()

	// If this is a CONTAINER element, then add the ending tag too.
	if element.container {
		// correct buffer size is set when we create the tag...
		element.builder.WriteString("</")
		element.builder.WriteString(element.name)
		element.builder.WriteRune('>')
	}

	// Mark the element closed
	element.closed = true

	// Pop this element off the builder's stack.
	element.builder.last = element.parent

	// And so it is finished
	return element
}
