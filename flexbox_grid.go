package styx

// Gap

func (b *Builder) Gap(value string, modifiers ...Modifier) *Builder {
	className := "gap-" + value
	b.createClass(className, modifiers...)
	b.createStyle(className, "gap:"+value+";", modifiers...)
	return b
}

func (b *Builder) GapX(value string, modifiers ...Modifier) *Builder {
	className := "gap-x-" + value
	b.createClass(className, modifiers...)
	b.createStyle(className, "gap-x:"+value+";", modifiers...)
	return b
}

func (b *Builder) GapY(value string, modifiers ...Modifier) *Builder {
	className := "gap-y-" + value
	b.createClass(className, modifiers...)
	b.createStyle(className, "gap-y:"+value+";", modifiers...)
	return b
}

// Place items

func (b *Builder) PlaceItems(value string, modifiers ...Modifier) *Builder {
	className := "place-items-" + value
	b.createClass(className, modifiers...)
	b.createStyle(className, "place-items:"+value+";", modifiers...)
	return b
}

func (b *Builder) PlaceItemsCenter(modifiers ...Modifier) *Builder {
	b.PlaceItems("center", modifiers...)
	return b
}

// Justify content

func (b *Builder) Justify(value string, modifiers ...Modifier) *Builder {
	className := "justify-" + value
	b.createClass(className, modifiers...)
	b.createStyle(className, "justify-content:"+value+";", modifiers...)
	return b
}

func (b *Builder) JustifyStart(modifiers ...Modifier) *Builder {
	b.Justify("start", modifiers...)
	return b
}

func (b *Builder) JustifyEnd(modifiers ...Modifier) *Builder {
	b.Justify("end", modifiers...)
	return b
}

func (b *Builder) JustifyCenter(modifiers ...Modifier) *Builder {
	b.Justify("center", modifiers...)
	return b
}

// Align items

func (b *Builder) Items(value string, modifiers ...Modifier) *Builder {
	className := "items-" + value
	b.createClass(className, modifiers...)
	b.createStyle(className, "align-items:"+value+";", modifiers...)
	return b
}

func (b *Builder) ItemsStart(modifiers ...Modifier) *Builder {
	b.Items("start", modifiers...)
	return b
}

func (b *Builder) ItemsEnd(modifiers ...Modifier) *Builder {
	b.Items("end", modifiers...)
	return b
}

func (b *Builder) ItemsCenter(modifiers ...Modifier) *Builder {
	b.Items("center", modifiers...)
	return b
}
