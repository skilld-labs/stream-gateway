package pin

import (
	"../codec"
)

type Pin struct {
	Codec codec.Codec `koanf:"codec"`
	ID    string      `koanf:"id"`
}
