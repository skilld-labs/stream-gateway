package codec

type JSONCodec struct {
	Depth    int  `yaml:"depth"`
	KeyValue bool `yaml:"keyValue"`
}

func (d *JSONCodec) Encode() ([]byte, error) {
	return []byte{}, nil
}

func (d *JSONCodec) Decode() (map[string]interface{}, error) {
	return nil, nil
}

func NewJSONCodec(c map[string]interface{}) *JSONCodec {
	return &JSONCodec{
		Depth:    c["depth"].(int),
		KeyValue: c["keyValue"].(bool),
	}
}
