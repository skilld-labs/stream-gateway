package codec

import (
	"../log"
)

type Codec interface {
	Encode() ([]byte, error)
	Decode() (map[string]interface{}, error)
}

type CodecConfig struct {
	Logger log.Logger
}

func Load(cc map[string]interface{}, cfg CodecConfig) Codec {
	var codec Codec
	//var err error
	for typ, c := range cc {
		switch typ {
		case "json":
			codec = NewJSONCodec(c.(map[string]interface{}))
		case "csv":
			codec = NewCSVCodec(c.(map[string]interface{}))
		default:
			cfg.Logger.Warn("Codec '%s' unknown", typ)
		}
	}
	return codec
}
