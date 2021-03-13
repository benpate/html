package builder

func (b *Builder) BR() *Element {
	return b.Empty("br")
}

func (b *Builder) Div() *Element {
	return b.Container("div")
}

func (b *Builder) Input() *Element {
	return b.Empty("input")
}

func (b *Builder) Label(forID string) *Element {
	return b.Container("label").For(forID)
}

func (b *Builder) Span() *Element {
	return b.Container("span")
}
