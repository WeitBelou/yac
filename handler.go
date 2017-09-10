package yac

import "net/http"

type Handler interface {
	ServeHTTP(w http.ResponseWriter, req *http.Request)
}

type HandlerFunc func(w http.ResponseWriter, req *http.Request)

func (h HandlerFunc) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	h(w, req)
}
