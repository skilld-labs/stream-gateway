package node

import (
	"../configuration"
	"../log"
	"../transport"
	"github.com/davecgh/go-spew/spew"
)

type HTTPNode struct {
	Name      string
	Transport transport.Transport
}

type HTTPNodeConfig struct {
	Logger log.Logger
	Config configuration.Provider
	Port   int
}

func NewHTTPNode(cfg HTTPNodeConfig) (HTTPNode, error) {
	return HTTPNode{}, nil
}

func (n *HTTPNode) AddPin(r Route) error {
	spew.Dump("ntm")
	return nil
}
