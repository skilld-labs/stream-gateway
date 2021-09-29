package gateway

import (
	"../node"
)

type Gateway struct {
	Nodes []node.Node
}

func (g *Gateway) AddNode(n node.Node) {
	g.Nodes = append(g.Nodes, n)
}
