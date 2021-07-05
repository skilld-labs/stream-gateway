package node

import (
	"../codec"
	"../configuration"
	"../log"
)

type NodeConfig struct {
	Type   string `yaml:"type"`
	Listen int    `yaml:"listen"`
	URI    string `yaml:"uri"`
	//auth to handle still
}

type Node interface {
	AddPin(Route) error
}

type Route struct {
	Path  string
	Codec codec.Codec
}

func NewFromConfig(p configuration.Provider, l log.Logger, nodeID, nodeConfigPath string) (Node, error) {
	var n Node
	node := NodeConfig{}
	p.Load(nodeConfigPath, &node)

	switch node.Type {
	case "http":
		//todo
		//if Listen defined = input
		//if Uri    defined = output

		hn, err := NewHTTPNode(HTTPNodeConfig{Logger: l, Config: p, Port: node.Listen})
		if err != nil {
			l.Fatal(err.Error())
		}
		n = &hn
	case "ftp":
	case "nats":
	default:
	}

	return n, nil
}
