package codec

type CSVCodec struct {
	Header    bool   `yaml:"header"`
	Delimiter string `yaml:"delimiter"`
}

func (c *CSVCodec) Encode() ([]byte, error) {
	return []byte{}, nil
}

func (c *CSVCodec) Decode() (map[string]interface{}, error) {
	return nil, nil
}

func NewCSVCodec(c map[string]interface{}) *CSVCodec {
	return &CSVCodec{
		Header:    c["header"].(bool),
		Delimiter: c["delimiter"].(string),
	}
}
