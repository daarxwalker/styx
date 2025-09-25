package styx

import (
	"bytes"
)

func build(b *bytes.Buffer, nodes []*node) {
	count := len(nodes)
	for i, n := range nodes {
		switch n.nodeType {
		case nodeTypeRule:
			build(b, n.nodes)
		case nodeTypeSelector:
			if i > 0 {
				b.WriteString("}")
			}
			buildSelector(b, n)
		case nodeTypeProperty:
			buildProperty(b, n)
		}
		if i == count-1 {
			b.WriteString("}")
		}
	}
}

func buildSelector(b *bytes.Buffer, n *node) {
	b.WriteString(n.key)
	b.WriteString("{")
}

func buildProperty(b *bytes.Buffer, n *node) {
	b.WriteString(n.key)
	b.WriteString(":")
	b.WriteString(n.value)
	b.WriteString(";")
}
