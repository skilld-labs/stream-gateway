package transport

type Transport struct {
	Typ    string `yaml:"typ"`
	Listen int    `yaml:"listen"`
	Uri    string `yaml:uri`
}
