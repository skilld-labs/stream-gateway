package node

import (
	"../pin"
	"../transport"
)

type Node interface {
	AddPin(pin.Pin)
	AddTransport(transport.Transport)
}

type NodeConfig struct {
	Name   string
	URI    string
	Listen string
	Type   string
}

func Load(cfg NodeConfig) Node {
	var node Node
	switch cfg.Type {
	case "http":
		node = NewHTTPNode(cfg)
	case "nats":
		node = NewNatsNode(cfg)
	default:
	}
	return node
}
