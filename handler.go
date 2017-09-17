package yac

import "net/http"

type Handler interface {
	ServeHTTP(w http.ResponseWriter, req *http.Request)
}

// Adapter to use functions as handlers.
type HandlerFunc func(w http.ResponseWriter, req *http.Request)

func (h HandlerFunc) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	h(w, req)
}
