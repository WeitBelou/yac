package yac

import (
	"fmt"
	"net/http"
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
	path := req.URL.Path
	if !r.routes.HasPattern(path) {
		pathNotFound(w)
		return
	}

	handler, err := r.routes.Get(path, req.Method)
	if err != nil {
		methodNotAllowed(w)
		return
	}

	handler.ServeHTTP(w, req)
}

// Default path not found response.
func pathNotFound(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("Not found!"))
}

// Default method not found response
func methodNotAllowed(w http.ResponseWriter) {
	w.WriteHeader(http.StatusMethodNotAllowed)
	w.Write([]byte("Method not allowed!"))
}
