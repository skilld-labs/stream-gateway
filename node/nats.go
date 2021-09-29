package node

import (
	"../pin"
	"../transport"
)

type NatsNode struct {
	Name      string
	Transport transport.Transport
	Pins      []pin.Pin
}

func NewNatsNode(cfg NodeConfig) *NatsNode {
	return &NatsNode{Name: cfg.Name}
}

func (n *NatsNode) AddPin(p pin.Pin) {
	n.Pins = append(n.Pins, p)
}

func (n *NatsNode) AddTransport(t transport.Transport) {
	n.Transport = t
}
