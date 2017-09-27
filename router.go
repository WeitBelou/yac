package yac

import (
	"fmt"
	"net/http"
)

type methods map[string]http.Handler
type routes map[string]methods

type router struct {
	rs routes
}

// Creates new router
func NewRouter() router {
	return router{rs: make(routes)}
}

func (r router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	ms, pathFound := r.rs[req.URL.Path]
	if !pathFound {
		handleNotFound(w, req)
		return
	}

	h, methodAllowed := ms[req.Method]
	if !methodAllowed {
		handleMethodNotAllowed(w, req)
		return
	}

	h.ServeHTTP(w, req)
}

// Add handler for given path and method.
// Returns error if there such route already exists.
func (r router) Handle(method string, path string, h http.Handler) error {
	ms, pathExists := r.rs[path]
	if !pathExists {
		r.rs[path] = make(methods, 1)
		r.rs[path][method] = h
		return nil
	}

	_, methodExists := ms[method]
	if methodExists {
		return fmt.Errorf("route with path '%s' and method '%s' already exists", path, method)
	}

	r.rs[path][method] = h
	return nil
}
