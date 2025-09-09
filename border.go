package styx

// Radius

func (b *Builder) Rounded(value string, modifiers ...Modifier) *Builder {
	className := "rounded-" + value
	b.createClass(className, modifiers...)
	b.createStyle(className, "border-radius:"+value+";", modifiers...)
	return b
}

// Border

func (b *Builder) Border(value string, modifiers ...Modifier) *Builder {
	className := "border-" + value
	b.createClass(className, modifiers...)
	b.createStyle(className, "border-width:"+value+";", modifiers...)
	return b
}

func (b *Builder) BorderX(value string, modifiers ...Modifier) *Builder {
	className := "border-x-" + value
	b.createClass(className, modifiers...)
	b.createStyle(className, "border-inline-width:"+value+";", modifiers...)
	return b
}

func (b *Builder) BorderY(value string, modifiers ...Modifier) *Builder {
	className := "border-y-" + value
	b.createClass(className, modifiers...)
	b.createStyle(className, "border-block-width:"+value+";", modifiers...)
	return b
}

func (b *Builder) BorderTop(value string, modifiers ...Modifier) *Builder {
	className := "border-t-" + value
	b.createClass(className, modifiers...)
	b.createStyle(className, "border-top-width:"+value+";", modifiers...)
	return b
}

func (b *Builder) BorderRight(value string, modifiers ...Modifier) *Builder {
	className := "border-r-" + value
	b.createClass(className, modifiers...)
	b.createStyle(className, "border-right-width:"+value+";", modifiers...)
	return b
}

func (b *Builder) BorderBottom(value string, modifiers ...Modifier) *Builder {
	className := "border-b-" + value
	b.createClass(className, modifiers...)
	b.createStyle(className, "border-bottom-width:"+value+";", modifiers...)
	return b
}

func (b *Builder) BorderLeft(value string, modifiers ...Modifier) *Builder {
	className := "border-l-" + value
	b.createClass(className, modifiers...)
	b.createStyle(className, "border-left-width:"+value+";", modifiers...)
	return b
}

func (b *Builder) BorderColor(color Color, modifiers ...Modifier) *Builder {
	className := "border-" + string(color)
	b.createClass(className, modifiers...)
	b.createStyle(className, "border-color:var(--color-"+string(color)+");", modifiers...)
	return b
}

func (b *Builder) BorderStyle(value string, modifiers ...Modifier) *Builder {
	className := "border-" + value
	b.createClass(className, modifiers...)
	b.createStyle(className, "border-style:"+value+";", modifiers...)
	return b
}

func (b *Builder) BorderCollapse(modifiers ...Modifier) *Builder {
	className := "border-collapse"
	b.createClass(className, modifiers...)
	b.createStyle(className, "border-collapse:collapse;", modifiers...)
	return b
}

func (b *Builder) BorderSeparate(modifiers ...Modifier) *Builder {
	className := "border-separate"
	b.createClass(className, modifiers...)
	b.createStyle(className, "border-collapse:separate;", modifiers...)
	return b
}

func (b *Builder) BorderSpacing(value string, modifiers ...Modifier) *Builder {
	className := "border-spacing-" + value
	b.createClass(className, modifiers...)
	b.createStyle(className, "border-spacing:"+value+";", modifiers...)
	return b
}

// Outline

func (b *Builder) Outline(value string, modifiers ...Modifier) *Builder {
	className := "outline-" + value
	b.createClass(className, modifiers...)
	b.createStyle(className, "outline-width:"+value+";", modifiers...)
	return b
}

func (b *Builder) OutlineColor(color Color, modifiers ...Modifier) *Builder {
	className := "outline-" + string(color)
	b.createClass(className, modifiers...)
	b.createStyle(className, "outline-color:var(--color-"+string(color)+");", modifiers...)
	return b
}

func (b *Builder) OutlineStyle(value string, modifiers ...Modifier) *Builder {
	className := "outline-" + value
	b.createClass(className, modifiers...)
	b.createStyle(className, "outline-style:"+value+";", modifiers...)
	return b
}
