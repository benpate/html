// Package html assembles HTML documents into a strings.Builder, without the
// indirection of a template.
//
// A Builder holds the output buffer and a stack of open elements. Each call
// returns the value you need next, so a document reads as a chain:
//
//	b := html.New()
//	b.Div().Class("wrapper")
//	b.Form().Attr("method", "post").Attr("action", "/save")
//	b.Input("text", "fullname").Value("John Connor").Close()
//	output := b.String()
//
// Elements come in two kinds. A container (Div, Span, Form) gets a closing tag;
// an empty one (BR, Img, Input) does not. Opening a new element automatically
// finishes whatever came before it, and String closes everything still on the
// stack, so a well-formed document falls out without matching every tag by hand.
// Close, CloseAll, and EndBracket drive the stack directly when the automatic
// behavior is not what you want.
//
// # What is escaped, and what is not
//
// This is the contract to keep in mind, because the package cannot tell trusted
// content from untrusted content on your behalf:
//
//   - Attribute VALUES are escaped. Attr, ForceAttr, and every helper built on
//     them (Class, Href, Data, and the rest) run the value through
//     html.EscapeString, so a value may safely come from user input.
//   - InnerText is escaped, and is how user-supplied text belongs in a document.
//   - InnerHTML is NOT escaped. It writes markup verbatim, by design.
//   - Attribute NAMES and tag names are NOT escaped. They are written into the
//     tag as given, so they must come from your own code, never from user input.
package html
