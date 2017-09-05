package yac

import (
	"log"
	"net/http"
	"time"
)

// Wrapper that takes handler and returns modified handler
type Wrapper func(h http.HandlerFunc) http.HandlerFunc

// Helper for wrappers slice
type Wrappers []Wrapper

// Wraps route with slice of wrappers
func (ws Wrappers) Wrap(h http.HandlerFunc) http.HandlerFunc {
	for _, w := range ws {
		h = w(h)
	}
	return h
}

// Logs requests
func Logger(h http.HandlerFunc) http.HandlerFunc {
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
