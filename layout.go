package styx

// Position

func (b *Builder) Absolute(modifiers ...Modifier) *Builder {
	className := "absolute"
	b.createClass(className, modifiers...)
	b.createStyle(className, "position:absolute;", modifiers...)
	return b
}

func (b *Builder) Relative(modifiers ...Modifier) *Builder {
	className := "relative"
	b.createClass(className, modifiers...)
	b.createStyle(className, "position:relative;", modifiers...)
	return b
}

func (b *Builder) Fixed(modifiers ...Modifier) *Builder {
	className := "fixed"
	b.createClass(className, modifiers...)
	b.createStyle(className, "position:fixed;", modifiers...)
	return b
}

func (b *Builder) Sticky(modifiers ...Modifier) *Builder {
	className := "sticky"
	b.createClass(className, modifiers...)
	b.createStyle(className, "position:sticky;", modifiers...)
	return b
}

func (b *Builder) Static(modifiers ...Modifier) *Builder {
	className := "static"
	b.createClass(className, modifiers...)
	b.createStyle(className, "position:static;", modifiers...)
	return b
}

func (b *Builder) Top(value string, modifiers ...Modifier) *Builder {
	className := "top-" + value
	b.createClass(className, modifiers...)
	b.createStyle(className, "top:"+value+";", modifiers...)
	return b
}

func (b *Builder) Right(value string, modifiers ...Modifier) *Builder {
	className := "right-" + value
	b.createClass(className, modifiers...)
	b.createStyle(className, "top:"+value+";", modifiers...)
	return b
}

func (b *Builder) Bottom(value string, modifiers ...Modifier) *Builder {
	className := "bottom-" + value
	b.createClass(className, modifiers...)
	b.createStyle(className, "top:"+value+";", modifiers...)
	return b
}

func (b *Builder) Left(value string, modifiers ...Modifier) *Builder {
	className := "left-" + value
	b.createClass(className, modifiers...)
	b.createStyle(className, "top:"+value+";", modifiers...)
	return b
}

func (b *Builder) Inset(value string, modifiers ...Modifier) *Builder {
	className := "inset-" + value
	b.createClass(className, modifiers...)
	b.createStyle(className, "top:"+value+";right:"+value+";bottom:"+value+";left:"+value+";", modifiers...)
	return b
}

func (b *Builder) InsetX(value string, modifiers ...Modifier) *Builder {
	className := "inset-" + value
	b.createClass(className, modifiers...)
	b.createStyle(className, "right:"+value+";left:"+value+";", modifiers...)
	return b
}

func (b *Builder) InsetY(value string, modifiers ...Modifier) *Builder {
	className := "inset-" + value
	b.createClass(className, modifiers...)
	b.createStyle(className, "top:"+value+";bottom:"+value+";", modifiers...)
	return b
}

// Box sizing

func (b *Builder) BoxBorder(modifiers ...Modifier) *Builder {
	className := "box-border"
	b.createClass(className, modifiers...)
	b.createStyle(className, "box-sizing:border-box;", modifiers...)
	return b
}

func (b *Builder) BoxContent(modifiers ...Modifier) *Builder {
	className := "box-content"
	b.createClass(className, modifiers...)
	b.createStyle(className, "box-sizing:context-box;", modifiers...)
	return b
}

// Overflow

func (b *Builder) Overflow(value string, modifiers ...Modifier) *Builder {
	className := "overflow-" + value
	b.createClass(className, modifiers...)
	b.createStyle(className, "overflow:"+value+";", modifiers...)
	return b
}

func (b *Builder) OverflowX(value string, modifiers ...Modifier) *Builder {
	className := "overflow-x-" + value
	b.createClass(className, modifiers...)
	b.createStyle(className, "overflow-x:"+value+";", modifiers...)
	return b
}

func (b *Builder) OverflowY(value string, modifiers ...Modifier) *Builder {
	className := "overflow-y-" + value
	b.createClass(className, modifiers...)
	b.createStyle(className, "overflow-y:"+value+";", modifiers...)
	return b
}

// Visibility

func (b *Builder) Visible(modifiers ...Modifier) *Builder {
	className := "visible"
	b.createClass(className, modifiers...)
	b.createStyle(className, "visibility:visible;", modifiers...)
	return b
}

func (b *Builder) Invisible(modifiers ...Modifier) *Builder {
	className := "invisible"
	b.createClass(className, modifiers...)
	b.createStyle(className, "visibility:invisible;", modifiers...)
	return b
}

func (b *Builder) Collapse(modifiers ...Modifier) *Builder {
	className := "collapse"
	b.createClass(className, modifiers...)
	b.createStyle(className, "visibility:collapse;", modifiers...)
	return b
}

// Z-index

func (b *Builder) Z(value string, modifiers ...Modifier) *Builder {
	className := "z-" + value
	b.createClass(className, modifiers...)
	b.createStyle(className, "z-index:"+value+";", modifiers...)
	return b
}

// Other

func (b *Builder) Aspect(value string, modifiers ...Modifier) *Builder {
	className := "aspect-" + value
	b.createClass(className, modifiers...)
	b.createStyle(className, "aspect-ratio:"+value+";", modifiers...)
	return b
}
