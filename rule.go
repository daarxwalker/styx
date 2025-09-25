package styx

func WithRule(nodes ...Node) Node {
	return func(n *node) {
		nestedRuleNode := &node{
			nodeType: nodeTypeRule,
		}
		for _, nd := range nodes {
			nd(nestedRuleNode)
		}
		sortNodes(nestedRuleNode)
		n.nodes = append(n.nodes, nestedRuleNode)
	}
}
