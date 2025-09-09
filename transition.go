package styx

func (b *Builder) Transition(modifiers ...Modifier) *Builder {
	className := "transition"
	style := `transition-property:color,background-color,border-color,outline-color,text-decoration-color,fill,stroke,opacity,box-shadow,transform,translate,scale,rotate,filter,-webkit-backdrop-filter,backdrop-filter,display,content-visibility,overlay,pointer-events;
transition-timing-function:cubic-bezier(0.4,0,0.2,1);
transition-duration:150ms;`
	b.createClass(className, modifiers...)
	b.createStyle(className, style, modifiers...)
	return b
}
