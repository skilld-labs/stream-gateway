package streamgateway

import "../node"

type StreamNodes []*node.Node

type StreamGateway interface {
	AddNode(node.Node) error
}

func (s *StreamGateway) AddNode(n node.Node) error {
	StreamNodes = append(StreamNodes, &n)
}
