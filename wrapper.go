package yac

import (
	"log"
	"net/http"
	"time"
)

// Wrapper that takes handler and returns modified handler
type Wrapper func(h Handler) Handler

// Helper for wrappers slice
type Wrappers []Wrapper

// Wraps route with slice of wrappers
func (ws Wrappers) Wrap(h Handler) Handler {
	for _, w := range ws {
		h = w(h)
	}
	return h
}

// Logs requests
func Logger(h Handler) Handler {
	return func(w http.ResponseWriter, req *http.Request) {
		start := time.Now()
		defer log.Printf(
			"%s\t%s\t%s",
			req.Method,
			req.RequestURI,
			time.Since(start),
		)

		h(w, req)
	}
}
