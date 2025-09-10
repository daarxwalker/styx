package styx

import "fmt"

// Flex direction

func (b *Builder) FlexRow(modifiers ...Modifier) *Builder {
	className := "flex-row"
	b.createClass(className, modifiers...)
	b.createStyle(className, "flex-direction:row;", modifiers...)
	return b
}

func (b *Builder) FlexRowReverse(modifiers ...Modifier) *Builder {
	className := "flex-row-reverse"
	b.createClass(className, modifiers...)
	b.createStyle(className, "flex-direction:row-reverse;", modifiers...)
	return b
}

func (b *Builder) FlexCol(modifiers ...Modifier) *Builder {
	className := "flex-col"
	b.createClass(className, modifiers...)
	b.createStyle(className, "flex-direction:column;", modifiers...)
	return b
}

func (b *Builder) FlexColReverse(modifiers ...Modifier) *Builder {
	className := "flex-col-reverse"
	b.createClass(className, modifiers...)
	b.createStyle(className, "flex-direction:column-reverse;", modifiers...)
	return b
}

// Flex wrap

func (b *Builder) FlexNoWrap(modifiers ...Modifier) *Builder {
	className := "flex-nowrap"
	b.createClass(className, modifiers...)
	b.createStyle(className, "flex-wrap:nowrap;", modifiers...)
	return b
}

func (b *Builder) FlexWrap(modifiers ...Modifier) *Builder {
	className := "flex-wrap"
	b.createClass(className, modifiers...)
	b.createStyle(className, "flex-wrap:wrap;", modifiers...)
	return b
}

func (b *Builder) FlexWrapReverse(modifiers ...Modifier) *Builder {
	className := "flex-wrap-reverse"
	b.createClass(className, modifiers...)
	b.createStyle(className, "flex-wrap:wrap-reverse;", modifiers...)
	return b
}

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

func (b *Builder) JustifyBetween(modifiers ...Modifier) *Builder {
	className := "justify-between"
	b.createClass(className, modifiers...)
	b.createStyle(className, "justify-content:space-between;", modifiers...)
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

// Grid

func (b *Builder) GridCols(value any, modifiers ...Modifier) *Builder {
	var colsValue string
	className := "grid-cols-"
	switch v := value.(type) {
	case string:
		className += v
		colsValue = v
	case int:
		strVal := fmt.Sprintf("%d", v)
		className += strVal
		colsValue = "repeat(" + strVal + ",minmax(0,1fr))"
	default:
		return b
	}
	b.createClass(className, modifiers...)
	b.createStyle(className, "grid-template-columns:"+colsValue+";", modifiers...)
	return b
}

func (b *Builder) GridRows(value any, modifiers ...Modifier) *Builder {
	var rowsValue string
	className := "grid-rows-"
	switch v := value.(type) {
	case string:
		className += v
		rowsValue = v
	case int:
		strVal := fmt.Sprintf("%d", v)
		className += strVal
		rowsValue = "repeat(" + strVal + ",minmax(0,1fr))"
	default:
		return b
	}
	b.createClass(className, modifiers...)
	b.createStyle(className, "grid-template-rows:"+rowsValue+";", modifiers...)
	return b
}
