package styx

import (
	"bytes"
	"fmt"
	"io"

	"github.com/dchest/uniuri"
)

type Styx struct {
	nodes []*node
}

func New() *Styx {
	return &Styx{
		nodes: make([]*node, 0),
	}
}

func (s *Styx) Global(nodes ...Node) {
	globalNode := &node{
		nodeType: nodeTypeGlobal,
	}
	for _, nd := range nodes {
		nd(globalNode)
	}
	s.nodes = append(s.nodes, globalNode.nodes...)
}

func (s *Styx) Rule(nodes ...Node) string {
	ruleNode := &node{
		nodeType: nodeTypeRule,
	}
	for _, nd := range nodes {
		nd(ruleNode)
	}
	var selector string
	for _, nd := range ruleNode.nodes {
		if nd.nodeType == nodeTypeSelector {
			selector = nd.key
			break
		}
	}
	if len(selector) == 0 {
		selector = "." + uniuri.New()
		ruleNode.nodes = append(
			ruleNode.nodes, &node{
				nodeType: nodeTypeSelector,
				key:      selector,
			},
		)
	}
	sortNodes(ruleNode)
	s.nodes = append(s.nodes, ruleNode)
	return stripSelectorPrefix(selector)
}

func (s *Styx) Write(w io.Writer) error {
	_, err := w.Write(s.Bytes())
	if err != nil {
		return fmt.Errorf("failed to write styles: %w", err)
	}
	return nil
}

func (s *Styx) Join(other *Styx) {
	s.nodes = append(s.nodes, other.nodes...)
}

func (s *Styx) Bytes() []byte {
	return s.build().Bytes()
}

func (s *Styx) String() string {
	return s.build().String()
}

func (s *Styx) build() *bytes.Buffer {
	b := new(bytes.Buffer)
	for _, n := range s.nodes {
		build(b, n.nodes)
	}
	return b
}
