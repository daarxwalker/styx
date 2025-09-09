package styx

func (b *Builder) Shadow(size Modifier, modifiers ...Modifier) *Builder {
	className := "shadow-" + string(size)
	b.createClass(className, modifiers...)
	b.createStyle(className, "box-shadow:var(--shadow-"+string(size)+");", modifiers...)
	return b
}

func (b *Builder) Opacity(value string, modifiers ...Modifier) *Builder {
	className := "opacity-" + value
	b.createClass(className, modifiers...)
	b.createStyle(className, "opacity:"+value+";", modifiers...)
	return b
}
