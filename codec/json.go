package codec

type JSONCodec struct {
	Depth    int
	KeyValue bool
}

func (d *JSONCodec) Encode() ([]byte, error) {
	return []byte{}, nil
}

func (d *JSONCodec) Decode() (map[string]interface{}, error) {
	return nil, nil
}
