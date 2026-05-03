package html

// HxGet adds a "hx-get" attribute to the Element
func (element *Element) HxGet(value string) *Element {
	return element.Attr("hx-get", value)
}

// HxPost adds a "hx-post" attribute to the Element
func (element *Element) HxPost(value string) *Element {
	return element.Attr("hx-post", value)
}

// HxTarget adds a "hx-target" attribute to the Element
func (element *Element) HxTarget(value string) *Element {
	return element.Attr("hx-target", value)
}

// HxTrigger adds a "hx-trigger" attribute to the Element
func (element *Element) HxTrigger(value string) *Element {
	return element.Attr("hx-trigger", value)
}

// HxSwap adds a "hx-swap" attribute to the Element
func (element *Element) HxSwap(value string) *Element {
	return element.Attr("hx-swap", value)
}
