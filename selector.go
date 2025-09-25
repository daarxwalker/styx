package styx

import "strings"

const (
	Self        = "&"
	Child       = ">"
	NextSibling = "+"
	Sibling     = "~"
)

func WithSelector(parts ...string) Node {
	key := strings.Join(parts, " ")
	return func(n *node) {
		n.key = key
		n.nodes = append(
			n.nodes, &node{
				nodeType: nodeTypeSelector,
				key:      key,
			},
		)
	}
}

func Class(value string) string {
	return "." + value
}

func Id(value string) string {
	return "#" + value
}

func Attribute(name, value string) string {
	return `[` + name + `="` + value + `"]`
}

func stripSelectorPrefix(sel string) string {
	if len(sel) > 0 && (sel[0] == '.' || sel[0] == '#') {
		return sel[1:]
	}
	return sel
}
