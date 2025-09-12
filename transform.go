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

func (b *Builder) ScaleX(value string, modifiers ...Modifier) *Builder {
	className := "scale-x-"
	if detectNegativeValue(value) {
		className = negateClass(className, value)
	}
	b.createClass(className, modifiers...)
	b.createStyle(className, "--scale-x:"+value+";scale:var(--scale-x) var(--scale-y);", modifiers...)
	return b
}

func (b *Builder) ScaleY(value string, modifiers ...Modifier) *Builder {
	className := "scale-y-"
	if detectNegativeValue(value) {
		className = negateClass(className, value)
	}
	b.createClass(className, modifiers...)
	b.createStyle(className, "--scale-y:"+value+";scale:var(--scale-x) var(--scale-y);", modifiers...)
	return b
}

func (b *Builder) Origin(value string, modifiers ...Modifier) *Builder {
	className := "origin-" + value
	b.createClass(className, modifiers...)
	b.createStyle(className, "transform-origin:"+value+";", modifiers...)
	return b
}

func (b *Builder) Translate(value string, modifiers ...Modifier) *Builder {
	className := "translate-"
	if detectNegativeValue(value) {
		className = negateClass(className, value)
	}
	b.createClass(className, modifiers...)
	b.createStyle(
		className, "--translate-x:"+value+";--translate-y:"+value+";translate:var(--translate-x) var(--translate-y);",
		modifiers...,
	)
	return b
}

func (b *Builder) TranslateX(value string, modifiers ...Modifier) *Builder {
	className := "translate-x-"
	if detectNegativeValue(value) {
		className = negateClass(className, value)
	}
	b.createClass(className, modifiers...)
	b.createStyle(
		className, "--translate-x:"+value+";translate:var(--translate-x) var(--translate-y);",
		modifiers...,
	)
	return b
}

func (b *Builder) TranslateY(value string, modifiers ...Modifier) *Builder {
	className := "translate-y-"
	if detectNegativeValue(value) {
		className = negateClass(className, value)
	}
	b.createClass(className, modifiers...)
	b.createStyle(
		className, "--translate-y:"+value+";translate:var(--translate-x) var(--translate-y);",
		modifiers...,
	)
	return b
}
