package styx

func (b *Builder) Cursor(value string, modifiers ...Modifier) *Builder {
	className := "cursor-" + value
	b.createClass(className, modifiers...)
	b.createStyle(className, "cursor:"+value+";", modifiers...)
	return b
}

func (b *Builder) PointerEvents(value string, modifiers ...Modifier) *Builder {
	className := "pointer-events-" + value
	b.createClass(className, modifiers...)
	b.createStyle(className, "pointer-events:"+value+";", modifiers...)
	return b
}

func (b *Builder) Select(value string, modifiers ...Modifier) *Builder {
	className := "select-" + value
	b.createClass(className, modifiers...)
	b.createStyle(className, "user-select:"+value+";", modifiers...)
	return b
}

func (b *Builder) WillChange(value string, modifiers ...Modifier) *Builder {
	className := "will-change-" + value
	b.createClass(className, modifiers...)
	b.createStyle(className, "will-change:"+value+";", modifiers...)
	return b
}

func (b *Builder) Appearance(value string, modifiers ...Modifier) *Builder {
	className := "appearance-" + value
	b.createClass(className, modifiers...)
	b.createStyle(className, "appearance:"+value+";", modifiers...)
	return b
}
