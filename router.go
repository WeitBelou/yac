package yac

import (
	"fmt"
	"net/http"
	"sync"
)

type methods map[string]http.Handler
type routes map[string]methods

type Router struct {
	mu sync.Mutex
	rs routes

	NotFound         http.Handler
	MethodNotAllowed http.Handler
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	h := r.getHandler(req.URL.Path, req.Method)

	h.ServeHTTP(w, req)
}

// getHandler returns handler or fallback
func (r *Router) getHandler(path, method string) http.Handler {
	ms, pathFound := r.rs[path]
	if !pathFound {
		if r.NotFound == nil {
			return http.HandlerFunc(notFound)
		}
		return r.NotFound
	}

	h, methodAllowed := ms[method]
	if !methodAllowed {
		if r.MethodNotAllowed == nil {
			return http.HandlerFunc(methodNotAllowed)
		}
		return r.MethodNotAllowed
	}
	return h
}

// Handle registers handler for given route.
func (r *Router) Handle(method string, path string, h http.Handler) error {
	r.mu.Lock()
	defer r.mu.Unlock()

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
