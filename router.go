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

// Structure that stores supported methods for each route.
type Routes []Route

// Returns new router with root path == rootPath
func NewRouter() *Router {
	return &Router{routes: make(Routes, 0)}
}

type Router struct {
	routes Routes

	wrappers []WrapperFunc
}

// Add wrappers to router
func (r *Router) AddWrappers(wrappers ...WrapperFunc) {
	r.wrappers = append(r.wrappers, wrappers...)
}

// Add generic route to routes.
func (r *Router) Route(pattern string, method string, handler http.HandlerFunc) error {
	re, err := regexp.Compile(convertSimplePatternToRegexp(pattern))
	if err != nil {
		return fmt.Errorf("can not compile pattern: %v", err)
	}

	for _, route := range r.routes {
		if route.Method == method && route.Pattern == pattern {
			return fmt.Errorf("route already exists")
		}
	}

	r.routes = append(r.routes, Route{
		Pattern: pattern, Method: method,
		Handler: Wrap(handler, r.wrappers...), matcher: re,
	})

	return nil
}

// Adds Get handler
func (r *Router) Get(pattern string, handler http.HandlerFunc) error {
	return r.Route(pattern, http.MethodGet, handler)
}

// Adds Post handler
func (r *Router) Post(pattern string, handler http.HandlerFunc) error {
	return r.Route(pattern, http.MethodPost, handler)
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
				if matched.matcher == nil || len(route.matcher.SubexpNames()) <= len(matched.matcher.SubexpNames()) {
					matched = route
				}
			}
		}
	}

	if matched.Handler != nil {
		params, err := newParams(req, matched.matcher, path)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		matched.Handler(w, putParamsToRequest(req, params))
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
