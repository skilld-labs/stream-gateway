package main

import (
	"flag"

	"./configuration"
	"./log"
	"./node"
)

func main() {
	//configuration, logger etc

	config := flag.String("config", "config.yaml", "The path for the config file")
	flag.Parse()

	l := log.NewJsonLogger(&log.LoggerConfiguration{})
	cfg, err := configuration.NewKoanfProvider(configuration.ProviderConfig{
		Logger: l,
		Source: *config,
	})
	if err != nil {
		l.Fatal(err.Error())
	}

	testNode, err := node.NewFromConfig(cfg, l, "test", "inputs.test.transport")
	if err != nil {
		l.Fatal(err.Error())
	}

	err = testNode.AddPin(node.Route{})
	if err != nil {
		l.Fatal(err.Error())
	}

}
