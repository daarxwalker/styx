package styx

func (b *Builder) Hidden(modifiers ...Modifier) *Builder {
	className := "hidden"
	b.createClass(className, modifiers...)
	b.createStyle(className, "display:none;", modifiers...)
	return b
}

func (b *Builder) Inline(modifiers ...Modifier) *Builder {
	className := "inline"
	b.createClass(className, modifiers...)
	b.createStyle(className, "display:inline;", modifiers...)
	return b
}

func (b *Builder) Block(modifiers ...Modifier) *Builder {
	className := "block"
	b.createClass(className, modifiers...)
	b.createStyle(className, "display:block;", modifiers...)
	return b
}

func (b *Builder) InlineBlock(modifiers ...Modifier) *Builder {
	className := "inline-block"
	b.createClass(className, modifiers...)
	b.createStyle(className, "display:inline-block;", modifiers...)
	return b
}

func (b *Builder) Flex(modifiers ...Modifier) *Builder {
	className := "flex"
	b.createClass(className, modifiers...)
	b.createStyle(className, "display:flex;", modifiers...)
	return b
}

func (b *Builder) InlineFlex(modifiers ...Modifier) *Builder {
	className := "inline-flex"
	b.createClass(className, modifiers...)
	b.createStyle(className, "display:inline-flex;", modifiers...)
	return b
}

func (b *Builder) Grid(modifiers ...Modifier) *Builder {
	className := "grid"
	b.createClass(className, modifiers...)
	b.createStyle(className, "display:grid;", modifiers...)
	return b
}

func (b *Builder) InlineGrid(modifiers ...Modifier) *Builder {
	className := "inline-grid"
	b.createClass(className, modifiers...)
	b.createStyle(className, "display:grid;", modifiers...)
	return b
}
