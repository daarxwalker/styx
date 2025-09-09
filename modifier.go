package styx

type Modifier string

const (
	Hover      Modifier = "hover"
	Focus      Modifier = "focus"
	Active     Modifier = "active"
	FirstChild Modifier = "first-child"
	LastChild  Modifier = "last-child"
	Checked    Modifier = "checked"
	Disabled   Modifier = "disabled"
)

func Not(selector string) Modifier {
	return Modifier("not(" + selector + ")")
}

func NthChild(value string) Modifier {
	return Modifier("nth-child(" + value + ")")
}

const (
	Before      Modifier = "before"
	After       Modifier = "after"
	Placeholder Modifier = "placeholder"
	Selection   Modifier = "selection"
)

const (
	Sm Modifier = "sm"
	Md Modifier = "md"
	Lg Modifier = "lg"
	Xl Modifier = "xl"
)

var (
	Pseudoselectors = map[Modifier]struct{}{
		Hover:      {},
		Focus:      {},
		Active:     {},
		FirstChild: {},
		LastChild:  {},
		Checked:    {},
		Disabled:   {},
	}
	Pseudoelements = map[Modifier]struct{}{
		Before:      {},
		After:       {},
		Placeholder: {},
		Selection:   {},
	}
)
