package styx

func (b *Builder) Container(modifiers ...Modifier) *Builder {
	className := "container"
	b.createClass(className, modifiers...)
	b.createStyle(className, "max-width:1200px;margin:0 auto;", modifiers...)
	return b
}
