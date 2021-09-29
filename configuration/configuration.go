package configuration

import (
	"github.com/knadh/koanf"
	"github.com/skilld-labs/http-event-adapter/log"
)

const (
	Tag = "configuration"
)

type ProviderConfig struct {
	Logger log.Logger
	Source interface{}
}

type Provider interface {
	Get(string) interface{}
	GetBool(string) bool
	GetFloat(string) float64
	GetFloats(string) []float64
	GetString(string) string
	GetStrings(string) []string
	GetMapStringString(string) map[string]string
	GetMapStringStrings(string) map[string][]string
	GetMapStringBool(string) map[string]bool
	GetKeys() []string
	GetBytes(string) []byte
	GetSlices(string) []*koanf.Koanf
	Load(string, interface{})
	Set(string, interface{})
}
