package yac

import (
	"log"
	"net/http"
	"time"
)

// Interface for handler wrappers
type Wrapper interface {
	// Takes Handler and returns new Handler
	Wrap(h Handler) Handler
}

// Adapter to use functions as wrappers
type WrapperFunc func(h Handler) Handler

// Calls w(h)
func (w WrapperFunc) Wrap(h Handler) Handler {
	return w(h)
}

// Helper for wrappers slice
type Wrappers []Wrapper

// Returns new wrappers with "wrappers" added
func (ws Wrappers) Add(wrappers ...Wrapper) Wrappers {
	return append(ws, wrappers...)
}

// Wraps route with slice of wrappers
func (ws Wrappers) Wrap(h Handler) Handler {
	for _, w := range ws {
		h = w.Wrap(h)
	}
	return h
}

// Requests logger
type LoggerWrapper struct {
	log.Logger
}

// Logs requests to associated logger
func (l LoggerWrapper) Wrap(handler Handler) Handler {
	return HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		start := time.Now()
		defer l.Printf(
			"%s\t%s\t%s",
			req.Method,
			req.RequestURI,
			time.Since(start),
		)

		handler.ServeHTTP(w, req)
	})
}
