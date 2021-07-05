package codec

type Codec interface {
	Decode() (map[string]interface{}, error)
	Encode() ([]byte, error)
}
