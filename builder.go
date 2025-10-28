package html

import (
	"io"
	"strings"
)

// Builder collects tags and attributes into a strings.Builder efficiently.
type Builder struct {
	*strings.Builder
	last *Element
}

// New generates a fully initialized Builder
func New() *Builder {
	var builder strings.Builder
	return &Builder{
		Builder: &builder,
	}
}

// Element adds a new HTML element to the builder
func (builder *Builder) Element(name string, container bool) *Element {

	result := &Element{
		builder:   builder,
		name:      name,
		container: container,
	}

	// If there's already another element on the stack then
	// we probably have some cleanup to do for it first.
	if builder.last != nil {

		// If the last element is NOT a container, then it should
		// be closed and popped from the stack.
		if !builder.last.container {
			builder.Close()
		} else {
			// If it IS a container, then it should at least
			// get an EndBracket
			builder.last.EndBracket()
		}

		// Point the result at the rest of the stack
		result.parent = builder.last
	}

	// Save this current element on the stack.
	builder.last = result

	return result.Start()
}

// Empty creates a new "empty" or non-container element that WILL NOT have an end tag
func (builder *Builder) Empty(name string) *Element {
	return builder.Element(name, false)
}

// Container creates a new "container" element that WILL have an end tag.
func (builder *Builder) Container(name string) *Element {
	return builder.Element(name, true)
}

// Space adds a single space character to the buffer
func (builder *Builder) Space() *Builder {
	builder.Grow(1)
	builder.WriteString(" ")
	return builder
}

// SubTree generates a new Builder that shares this Builder's string buffer.
// This is useful when sending a Builder to another function, so that the
// other function can maintain it's own stack of elements -- and potentially
// call .CloseAll() -- without affecting this current builder.
func (builder *Builder) SubTree() *Builder {

	// If we're beginning a sub-tree, then guarantee that the most recent tag has at
	// least been ended properly.
	builder.EndBracket()

	return &Builder{
		Builder: builder.Builder,
	}
}

// EndBracket adds an end bracket to the last tag on the stack
func (builder *Builder) EndBracket() *Builder {

	if builder.last == nil {
		return builder
	}

	// Write the closing HTML to the buffer
	builder.last.EndBracket()

	return builder
}

// Close completes the last tag on the stack, then pops it off of the stack
func (builder *Builder) Close() *Builder {

	if builder.last == nil {
		return builder
	}

	// Write the closing HTML to the buffer
	builder.last.Close()

	return builder
}

// CloseAll calls .Cload() until the stack is empty.
func (builder *Builder) CloseAll() *Builder {

	for builder.last != nil {
		builder.Close()
	}

	return builder
}

// ReadString reads the existing contents of the buffer but DOES NOT
// close any existing tags.  This is useful for reading the header
// of an HTML document before calling a subroutine that will fill in
// its body.
func (builder *Builder) ReadString() string {

	// if there was any extra space remaining in the old buffer,
	// then we'll use that in the new buffer, too
	extraSpace := builder.Cap() - builder.Len()

	result := builder.String()
	builder.Reset()

	if extraSpace > 0 {
		builder.Grow(extraSpace)
	}

	// Return the contents of the original buffer so far.
	return result
}

// String returns the assembled HTML as a string.  It overrides
// the default behavior of the strings.Builder by also calling
// CloseAll() on all unclosed tags in the stack before generating HTML.
func (builder *Builder) String() string {
	builder.CloseAll()
	return builder.Builder.String()
}

// Bytes returns the assembled HTML as a slice of bytes.
func (builder *Builder) Bytes() []byte {
	builder.CloseAll()
	return []byte(builder.Builder.String())
}

// Reader returns the string as an io.Reader.
func (builder *Builder) Reader() io.Reader {
	return strings.NewReader(builder.String())
}
