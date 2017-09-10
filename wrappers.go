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

// Returns new wrappers with "wrappers" added
func (ws Wrappers) Add(wrappers ...Wrapper) Wrappers {
	newWs := append(ws, wrappers...)
	return newWs
}

// Wraps route with slice of wrappers
func (ws Wrappers) Wrap(h Handler) Handler {
	for _, w := range ws {
		h = w(h)
	}
	return h
}

// Logs requests
func Logger(logger log.Logger, h Handler) Handler {
	return HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		start := time.Now()
		defer logger.Printf(
			"%s\t%s\t%s",
			req.Method,
			req.RequestURI,
			time.Since(start),
		)

		h.ServeHTTP(w, req)
	})
}
