package yac

import (
	"fmt"
	"net/http"
)

type methods map[string]http.Handler
type routes map[string]methods

type Router struct {
	rs routes
}

func (r Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	ms, pathFound := r.rs[req.URL.Path]
	if !pathFound {
		defaultNotFoundHandler(w, req)
		return
	}

	h, methodAllowed := ms[req.Method]
	if !methodAllowed {
		defaultMethodNotAllowedHandler(w, req)
		return
	}

	h.ServeHTTP(w, req)
}

// Handle registers handler for given route.
func (r *Router) Handle(method string, path string, h http.Handler) error {
	errs := validateRoute(method, path, h)
	if len(errs) != 0 {
		return fmt.Errorf("invalid route: %v", errs)
	}

	if len(r.rs) == 0 {
		r.rs = make(routes, 1)
	}

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
