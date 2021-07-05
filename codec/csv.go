package codec

type CSVCodec struct {
	Header    bool
	Delimiter string
}

func (c *CSVCodec) Encode() ([]byte, error) {
	return []byte{}, nil
}

func (c *CSVCodec) Decode() (map[string]interface{}, error) {
	return nil, nil
}
