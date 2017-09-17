package yac

import (
	"context"
	"fmt"
	"net/http"

	"github.com/weitbelou/yac/params"
)

type Router struct {
	routes Routes

	wrappers Wrappers
}

// Returns new router
func NewRouter() Router {
	return Router{
		routes:   NewRoutes(),
		wrappers: Wrappers{},
	}
}

// Add wrappers to router
// Wrappers will be applied to handler function of every route.
func (r *Router) AddWrappers(wrappers ...Wrapper) {
	r.wrappers = append(r.wrappers, wrappers...)
}

// Add new route to routes.
func (r *Router) Route(pattern, method string, h Handler) error {
	return r.routes.Add(pattern, method, h)
}

// Listen on given port
func (r *Router) ListenAndServe(port string) error {
	return http.ListenAndServe(fmt.Sprintf(":%s", port), r)
}

// Implements http.Handler interface
func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	handler, p, err := r.routes.Get(req.URL.Path, req.Method)

	switch err.(type) {
	case ErrPathNotFound:
		pathNotFound(w)
	case ErrMethodNotAllowed:
		methodNotAllowed(w)
	default:
		handler.ServeHTTP(w, req.WithContext(context.WithValue(req.Context(), params.ContextKey, p)))
	}
}

// Default response for "path not found".
func pathNotFound(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("Not found!"))
}

// Default response for "method not allowed"
func methodNotAllowed(w http.ResponseWriter) {
	w.WriteHeader(http.StatusMethodNotAllowed)
	w.Write([]byte("Method not allowed!"))
}
