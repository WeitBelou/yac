package yac

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strings"
)

type Route struct {
	Method  string
	Pattern string
	Handler http.HandlerFunc

	matcher *regexp.Regexp
}

// Checks if router the same: method and pattern equal
func (r Route) Same(o Route) bool {
	return (r.Method == o.Method) && (r.Pattern == o.Pattern)
}

// Returns true if this route's matcher 'simpler' than other: matcher has less named subexpression
// If 'o' doesn't contain matcher (i.e. empty Route) than current considered to be simple
func (r Route) Simpler(o Route) bool {
	return o.matcher == nil || len(r.matcher.SubexpNames()) < len(o.matcher.SubexpNames())
}

// Helper for slice of routes
type Routes []Route

// Returns new router
func NewRouter() *Router {
	return &Router{
		routes:   make(Routes, 0),
		wrappers: Wrappers{patternCompiler, paramsInserter},
	}
}

type Router struct {
	routes Routes

	wrappers Wrappers
}

// Add wrappers to router
// Wrappers will be applied to handler function of every route.
func (r *Router) AddWrappers(wrappers ...Wrapper) {
	r.wrappers = append(r.wrappers, wrappers...)
}

// Add new route to routes.
func (r *Router) Route(route Route) error {
	for _, rt := range r.routes {
		if rt.Same(route) {
			return fmt.Errorf("route already exists")
		}
	}

	r.routes = append(r.routes, r.wrappers.Wrap(route))
	return nil
}

// Adds Get handler
func (r *Router) Get(pattern string, handler http.HandlerFunc) error {
	return r.Route(Route{Method: http.MethodGet, Pattern: pattern, Handler: handler})
}

// Adds Post handler
func (r *Router) Post(pattern string, handler http.HandlerFunc) error {
	return r.Route(Route{Method: http.MethodPost, Pattern: pattern, Handler: handler})
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
	var pathFound bool

	var matched Route

	path := req.URL.Path
	for _, route := range r.routes {
		if route.matcher.MatchString(path) {
			pathFound = true
			if route.Method == req.Method {
				if route.Simpler(matched) {
					matched = route
				}
			}
		}
	}

	if matched.Handler != nil {
		matched.Handler(w, req)
		return
	}

	if pathFound {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "method %s not allowed at path %s", req.Method, req.URL)
		return
	}

	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintf(w, "path %s not found", req.URL)
}

// Pretty prints routes
func (r *Router) PrintRoutes() {
	log.Println(strings.Repeat("-", 10))

	for _, route := range r.routes {
		log.Printf("'%s': '%s'", route.Method, route.Pattern)
	}

	log.Println(strings.Repeat("-", 10))
}
