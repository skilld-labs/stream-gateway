package main

import (
	"flag"

	"./codec"
	"./configuration"
	"./gateway"
	"./log"
	"./node"
	"./pin"
	"./transport"
	"github.com/davecgh/go-spew/spew"
)

func main() {
	config := flag.String("config", "config.yaml", "The path for the config file")
	flag.Parse()

	l := log.NewJsonLogger(&log.LoggerConfiguration{Verbosity: log.Err})
	cfg, err := configuration.NewKoanfProvider(configuration.ProviderConfig{
		Logger: l,
		Source: *config,
	})
	if err != nil {
		l.Fatal(err.Error())
	}

	g := gateway.Gateway{}

	for nodeName, _ := range cfg.Get("nodes").(map[string]interface{}) {
		pp := []pin.Pin{}
		tr := transport.Transport{}

		n := node.Load(node.NodeConfig{
			Name:   nodeName,
			URI:    cfg.GetString("nodes." + nodeName + ".transport.uri"),
			Listen: cfg.GetString("nodes." + nodeName + ".transport.listen"),
			Type:   cfg.GetString("nodes." + nodeName + ".transport.typ"),
		})

		//this throws an error because it tries to load the codec config into a Codec but it's a map[string]interface{}
		cfg.Load("nodes."+nodeName+".pins", &pp)
		for _, p := range pp {
			for _, pin := range cfg.Get("nodes." + nodeName + ".pins").([]interface{}) {
				if p.ID == pin.(map[string]interface{})["id"] {
					p.Codec = codec.Load(
						pin.(map[string]interface{})["codec"].(map[string]interface{}),
						codec.CodecConfig{Logger: l},
					)
				}
			}
			n.AddPin(p)
		}
		cfg.Load("nodes."+nodeName+".transport", &tr)
		n.AddTransport(tr)

		g.AddNode(n)
	}
	spew.Dump(g)
}
