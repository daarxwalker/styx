package styx

func (b *Builder) BgColor(color Color, modifiers ...Modifier) *Builder {
	className := "bg-" + string(color)
	b.createClass(className, modifiers...)
	b.createStyle(className, "background-color:var(--color-"+string(color)+");", modifiers...)
	return b
}
