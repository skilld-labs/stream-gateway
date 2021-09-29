package node

import (
	"../pin"
	"../transport"
)

type HTTPNode struct {
	Name      string
	Transport transport.Transport
	Pins      []pin.Pin
}

func NewHTTPNode(cfg NodeConfig) *HTTPNode {
	//here handle Transport. Client if URI is defined, Server if Listen is defined
	return &HTTPNode{Name: cfg.Name}
}

func (n *HTTPNode) AddPin(p pin.Pin) {
	n.Pins = append(n.Pins, p)
}

func (n *HTTPNode) AddTransport(t transport.Transport) {
	n.Transport = t
}
