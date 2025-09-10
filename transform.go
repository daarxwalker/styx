package styx

func (b *Builder) Rotate(value string, modifiers ...Modifier) *Builder {
	className := "rotate-"
	if detectNegativeValue(value) {
		className = negateClass(className, value)
	}
	b.createClass(className, modifiers...)
	b.createStyle(className, "rotate:"+value+";", modifiers...)
	return b
}
