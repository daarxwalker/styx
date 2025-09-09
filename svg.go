package styx

func (b *Builder) Fill(value string, modifiers ...Modifier) *Builder {
	className := "fill-" + value
	b.createClass(className, modifiers...)
	b.createStyle(className, "fill:"+value+";", modifiers...)
	return b
}

func (b *Builder) Stroke(value string, modifiers ...Modifier) *Builder {
	className := "stroke-" + value
	b.createClass(className, modifiers...)
	b.createStyle(className, "stroke"+value+";", modifiers...)
	return b
}
