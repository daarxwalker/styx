package styx

import "strings"

func WithAtRule(nodes ...Node) Node {
	return func(n *node) {
		nestedRuleNode := &node{
			nodeType: nodeTypeRule,
		}
		for _, nd := range nodes {
			nd(nestedRuleNode)
		}
		var hasSelector bool
		for _, nd := range nestedRuleNode.nodes {
			if nd.nodeType == nodeTypeSelector {
				hasSelector = true
				if !strings.HasPrefix(nd.key, "@") {
					nd.key = "@" + nd.key
				}
			}
		}
		if !hasSelector {
			return
		}
		n.nodes = append(n.nodes, nestedRuleNode)
	}
}

func WithCharsetRule(charset string) Node {
	return WithAtRule(
		WithSelector("charset \"" + charset + "\""),
	)
}

func WithColorProfileRule(name string, nodes ...Node) Node {
	return WithAtRule(
		WithSelector("color-profile "+name),
		WithFragment(nodes...),
	)
}

func WithContainerRule(query string, nodes ...Node) Node {
	return WithAtRule(
		WithSelector("container "+query),
		WithFragment(nodes...),
	)
}

func WithCounterStyleRule(name string, nodes ...Node) Node {
	return WithAtRule(
		WithSelector("counter-style "+name),
		WithFragment(nodes...),
	)
}

func WithFontFaceRule(nodes ...Node) Node {
	return WithAtRule(
		WithSelector("font-face"),
		WithFragment(nodes...),
	)
}

func WithFontFeatureValuesRule(name string, nodes ...Node) Node {
	return WithAtRule(
		WithSelector("font-feature-values "+name),
		WithFragment(nodes...),
	)
}

func WithFontPaletteValuesRule(name string, nodes ...Node) Node {
	return WithAtRule(
		WithSelector("font-palette-values "+name),
		WithFragment(nodes...),
	)
}

func WithImportRule(url string) Node {
	return WithAtRule(
		WithSelector("import \"" + url + "\""),
	)
}

func WithKeyframesRule(name string, nodes ...Node) Node {
	return WithAtRule(
		WithSelector("keyframes "+name),
		WithFragment(nodes...),
	)
}

func WithLayerRule(name string, nodes ...Node) Node {
	return WithAtRule(
		WithSelector("layer "+name),
		WithFragment(nodes...),
	)
}

func WithMediaRule(query string, nodes ...Node) Node {
	return WithAtRule(
		WithSelector("media "+query),
		WithFragment(nodes...),
	)
}

func WithNamespaceRule(ns string) Node {
	return WithAtRule(
		WithSelector("namespace " + ns),
	)
}

func WithPageRule(name string, nodes ...Node) Node {
	return WithAtRule(
		WithSelector("page "+name),
		WithFragment(nodes...),
	)
}

func WithPositionTryRule(name string, nodes ...Node) Node {
	return WithAtRule(
		WithSelector("position-try "+name),
		WithFragment(nodes...),
	)
}

func WithPropertyRule(name string, nodes ...Node) Node {
	return WithAtRule(
		WithSelector("property "+name),
		WithFragment(nodes...),
	)
}

func WithScopeRule(name string, nodes ...Node) Node {
	return WithAtRule(
		WithSelector("scope "+name),
		WithFragment(nodes...),
	)
}

func WithStartingStyleRule(nodes ...Node) Node {
	return WithAtRule(
		WithSelector("starting-style"),
		WithFragment(nodes...),
	)
}

func WithSupportsRule(condition string, nodes ...Node) Node {
	return WithAtRule(
		WithSelector("supports "+condition),
		WithFragment(nodes...),
	)
}

func WithViewTransitionRule(name string, nodes ...Node) Node {
	return WithAtRule(
		WithSelector("view-transition "+name),
		WithFragment(nodes...),
	)
}
