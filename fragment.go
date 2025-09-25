package styx

func WithFragment(nodes ...Node) Node {
	return func(n *node) {
		for _, nd := range nodes {
			nd(n)
		}
	}
}
