package transport

import (
	"net/http"
)

type HTTPTransport struct {
	Listen string
	URI    string
}

func (t *HTTPTransport) ListenAndServe() error {
	//?
	return http.ListenAndServe(":"+t.Listen, nil)
}
