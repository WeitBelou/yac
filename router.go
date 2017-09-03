package yac

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
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
func NewRouter(root string) (*Router, error) {
	r := &Router{routes: make(Routes, 0)}

	newRoot, err := url.Parse(root)
	if err != nil {
		return nil, fmt.Errorf("invalid path format %s: %v", root, err)
	}

	r.root = newRoot

	return r, nil
}

type Router struct {
	root *url.URL

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
	relPath, err := relativePath(r.root.Path, req.URL.Path)
	if err != nil {
		http.NotFound(w, req)
	}

	r.handleRequest(w, req, relPath)
}

// Handles request: iterate over all routes before finds first matching route.
func (r *Router) handleRequest(w http.ResponseWriter, req *http.Request, path string) {
	var pathFound bool

	var matched *Route

	for _, route := range r.routes {
		// TODO: Here something very strange: occasionally route matches wrong path.
		if route.matcher.MatchString(path) {
			pathFound = true
			if route.Method == req.Method {
				if matched == nil || len(route.matcher.SubexpNames()) <= len(matched.matcher.SubexpNames()) {
					matched = &route
				}
			}
		}
	}

	if matched != nil {
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
