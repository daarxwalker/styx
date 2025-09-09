package styx

func (b *Builder) W(value string, modifiers ...Modifier) *Builder {
	className := "w-" + value
	b.createClass(className, modifiers...)
	b.createStyle(className, "width:"+value+";", modifiers...)
	return b
}

func (b *Builder) MinW(value string, modifiers ...Modifier) *Builder {
	className := "min-w-" + value
	b.createClass(className, modifiers...)
	b.createStyle(className, "min-width:"+value+";", modifiers...)
	return b
}

func (b *Builder) MaxW(value string, modifiers ...Modifier) *Builder {
	className := "max-w-" + value
	b.createClass(className, modifiers...)
	b.createStyle(className, "max-width:"+value+";", modifiers...)
	return b
}

func (b *Builder) H(value string, modifiers ...Modifier) *Builder {
	className := "h-" + value
	b.createClass(className, modifiers...)
	b.createStyle(className, "height:"+value+";", modifiers...)
	return b
}

func (b *Builder) MinH(value string, modifiers ...Modifier) *Builder {
	className := "min-h-" + value
	b.createClass(className, modifiers...)
	b.createStyle(className, "min-height:"+value+";", modifiers...)
	return b
}

func (b *Builder) MaxH(value string, modifiers ...Modifier) *Builder {
	className := "max-h-" + value
	b.createClass(className, modifiers...)
	b.createStyle(className, "max-height:"+value+";", modifiers...)
	return b
}
