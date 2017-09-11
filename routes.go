package yac

import "fmt"

// Method name to handler
type Methods map[string]Handler

// Pattern to methods map
type Routes map[string]Methods

func NewRoutes() Routes {
	return make(Routes)
}

// Adds route to routes
// Returns error if there is such route.
//
// See: Has(pattern, method)
func (rs Routes) Add(pattern, method string, h Handler) error {
	if rs.Has(pattern, method) {
		return ErrRouteAlreadyExists{pattern: pattern, method: method}
	}

	if !rs.HasPattern(pattern) {
		rs[pattern] = make(Methods)
	}

	rs[pattern][method] = h
	return nil
}

// Returns handler for 'pattern' 'method'
// It's safe to ignore error if rs.Has() returns true
func (rs Routes) Get(pattern, method string) (Handler, error) {
	if !rs.Has(pattern, method) {
		return nil, ErrRouteNotFound{pattern: pattern, method: method}
	}

	return rs[pattern][method], nil
}

// Checks if route in routes
func (rs Routes) Has(pattern, method string) bool {
	_, ok := rs[pattern][method]
	return ok
}

// Checks if pattern in routes
func (rs Routes) HasPattern(pattern string) bool {
	_, ok := rs[pattern]
	return ok
}

type ErrRouteAlreadyExists struct {
	pattern string
	method  string
}

func (e ErrRouteAlreadyExists) Error() string {
	return fmt.Sprintf("route with pattern '%s' and method '%s' already exists", e.pattern, e.method)
}

type ErrRouteNotFound struct {
	pattern string
	method  string
}

func (e ErrRouteNotFound) Error() string {
	return fmt.Sprintf("route with pattern '%s' and method '%s' not found", e.pattern, e.method)
}
