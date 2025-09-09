package styx

type FontFace string

func (b *Builder) Font(value string, modifiers ...Modifier) *Builder {
	className := "font-" + value
	b.createClass(className, modifiers...)
	b.createStyle(className, "font:"+value+";", modifiers...)
	return b
}

func (b *Builder) FontFamily(fontFamilyName string, modifiers ...Modifier) *Builder {
	className := "font-" + fontFamilyName
	b.createClass(className, modifiers...)
	b.createStyle(className, "font-family:"+fontFamilyName+",sans-serif;", modifiers...)
	return b
}

func (b *Builder) FontWeight(value string, modifiers ...Modifier) *Builder {
	className := "font-" + value
	b.createClass(className, modifiers...)
	b.createStyle(className, "font-weight:"+value+";", modifiers...)
	return b
}

func (b *Builder) FontSize(value string, modifiers ...Modifier) *Builder {
	className := "font-" + value
	b.createClass(className, modifiers...)
	b.createStyle(className, "font-size:var(--color-"+value+");", modifiers...)
	return b
}

func (b *Builder) Antialiased(modifiers ...Modifier) *Builder {
	className := "antialiased"
	b.createClass(className, modifiers...)
	b.createStyle(className, "-webkit-font-smoothing:antialiased;-moz-osx-font-smoothing:grayscale;", modifiers...)
	return b
}

func (b *Builder) SubpixelAntialiased(modifiers ...Modifier) *Builder {
	className := "subpixel-antialiased"
	b.createClass(className, modifiers...)
	b.createStyle(className, "-webkit-font-smoothing:auto;-moz-osx-font-smoothing:auto;", modifiers...)
	return b
}

func (b *Builder) TextColor(color Color, modifiers ...Modifier) *Builder {
	className := "text-" + string(color)
	b.createClass(className, modifiers...)
	b.createStyle(className, "color:var(--color-"+string(color)+");", modifiers...)
	return b
}

func (b *Builder) Underline(modifiers ...Modifier) *Builder {
	className := "underline"
	b.createClass(className, modifiers...)
	b.createStyle(className, "text-decoration:underline;", modifiers...)
	return b
}

func (b *Builder) NoUnderline(modifiers ...Modifier) *Builder {
	className := "no-underline"
	b.createClass(className, modifiers...)
	b.createStyle(className, "text-decoration:no-underline;", modifiers...)
	return b
}

func (b *Builder) TextAlign(value string, modifiers ...Modifier) *Builder {
	className := "text-" + value
	b.createClass(className, modifiers...)
	b.createStyle(className, "text-align:"+value+";", modifiers...)
	return b
}

func (b *Builder) TextLeft(value string, modifiers ...Modifier) *Builder {
	className := "text-" + value
	b.createClass(className, modifiers...)
	b.createStyle(className, "text-align:"+value+";", modifiers...)
	return b
}

func (b *Builder) LineHeight(value string, modifiers ...Modifier) *Builder {
	className := "line-height-" + value
	b.createClass(className, modifiers...)
	b.createStyle(className, "line-height:"+value+";", modifiers...)
	return b
}

func (b *Builder) ListStyle(value string, modifiers ...Modifier) *Builder {
	className := "list-style-" + value
	b.createClass(className, modifiers...)
	b.createStyle(className, "list-style:"+value+";", modifiers...)
	return b
}
