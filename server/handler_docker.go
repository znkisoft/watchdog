package server

import (
	"net/http"
)

type Transport struct {
	HTTPTransport http.Transport
}

func NewTransport() (*Transport, error) {
	return &Transport{
		HTTPTransport: http.Transport{},
	}, nil
}

func (t *Transport) ContainerProxy(w http.ResponseWriter, r *http.Request) (any, error) {

	return nil, nil
}
