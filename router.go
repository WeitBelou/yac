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
	r.handleRequest(w, req)
}

// Handles request: iterate over all routes before finds first matching route.
func (r *Router) handleRequest(w http.ResponseWriter, req *http.Request) {
	if !r.routes.HasPattern(req.URL.Path) {
		pathNotFound(w)
	}
}

// Default path not found response.
func pathNotFound(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "Not found!")
}
