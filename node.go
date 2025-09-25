package styx

type Node func(n *node)

type nodeType int

type node struct {
	nodeType nodeType
	key      string
	value    string
	nodes    []*node
}

const (
	nodeTypeGlobal = iota
	nodeTypeRule
	nodeTypeSelector
	nodeTypeProperty
)

func GetSelector(n Node) string {
	globalNode := &node{
		nodeType: nodeTypeGlobal,
	}
	n(globalNode)
	if len(globalNode.nodes) > 0 {
		return stripSelectorPrefix(globalNode.nodes[0].key)
	}
	return ""
}

func sortNodes(n *node) {
	if len(n.nodes) <= 1 {
		return
	}
	selectorIndex := -1
	for i, nd := range n.nodes {
		if nd.nodeType == nodeTypeSelector {
			selectorIndex = i
			break
		}
	}
	if selectorIndex > 0 {
		selectorNode := n.nodes[selectorIndex]
		copy(n.nodes[1:selectorIndex+1], n.nodes[0:selectorIndex])
		n.nodes[0] = selectorNode
	}
}
