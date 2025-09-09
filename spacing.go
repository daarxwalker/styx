package styx

func (b *Builder) P(value string, modifiers ...Modifier) *Builder {
	className := "p-" + value
	b.createClass(className, modifiers...)
	b.createStyle(className, "padding:"+value+";", modifiers...)
	return b
}

func (b *Builder) PX(value string, modifiers ...Modifier) *Builder {
	className := "px-" + value
	b.createClass(className, modifiers...)
	b.createStyle(className, "padding-right:"+value+";padding-left:"+value+";", modifiers...)
	return b
}

func (b *Builder) PY(value string, modifiers ...Modifier) *Builder {
	className := "py-" + value
	b.createClass(className, modifiers...)
	b.createStyle(className, "padding-top:"+value+";padding-bottom:"+value+";", modifiers...)
	return b
}

func (b *Builder) PT(value string, modifiers ...Modifier) *Builder {
	className := "pt-" + value
	b.createClass(className, modifiers...)
	b.createStyle(className, "padding-top:"+value+";", modifiers...)
	return b
}

func (b *Builder) PR(value string, modifiers ...Modifier) *Builder {
	className := "pr-" + value
	b.createClass(className, modifiers...)
	b.createStyle(className, "padding-right:"+value+";", modifiers...)
	return b
}

func (b *Builder) PB(value string, modifiers ...Modifier) *Builder {
	className := "pb-" + value
	b.createClass(className, modifiers...)
	b.createStyle(className, "padding-bottom:"+value+";", modifiers...)
	return b
}

func (b *Builder) PL(value string, modifiers ...Modifier) *Builder {
	className := "pl-" + value
	b.createClass(className, modifiers...)
	b.createStyle("pl-"+value, "padding-left:"+value+";", modifiers...)
	return b
}

func (b *Builder) M(value string, modifiers ...Modifier) *Builder {
	className := "m-" + value
	b.createClass(className, modifiers...)
	b.createStyle(className, "margin:"+value+";", modifiers...)
	return b
}

func (b *Builder) MX(value string, modifiers ...Modifier) *Builder {
	className := "mx-" + value
	b.createClass(className, modifiers...)
	b.createStyle(className, "margin-right:"+value+";margin-left:"+value+";", modifiers...)
	return b
}

func (b *Builder) MY(value string, modifiers ...Modifier) *Builder {
	className := "my-" + value
	b.createClass(className, modifiers...)
	b.createStyle(className, "margin-top:"+value+";margin-bottom:"+value+";", modifiers...)
	return b
}

func (b *Builder) MT(value string, modifiers ...Modifier) *Builder {
	className := "mt-" + value
	b.createClass(className, modifiers...)
	b.createStyle(className, "margin-top:"+value+";", modifiers...)
	return b
}

func (b *Builder) MR(value string, modifiers ...Modifier) *Builder {
	className := "mr-" + value
	b.createClass(className, modifiers...)
	b.createStyle(className, "margin-right:"+value+";", modifiers...)
	return b
}

func (b *Builder) MB(value string, modifiers ...Modifier) *Builder {
	className := "mb-" + value
	b.createClass(className, modifiers...)
	b.createStyle(className, "margin-bottom:"+value+";", modifiers...)
	return b
}

func (b *Builder) ML(value string, modifiers ...Modifier) *Builder {
	className := "ml-" + value
	b.createClass(className, modifiers...)
	b.createStyle(className, "margin-left:"+value+";", modifiers...)
	return b
}
