package yac

import "fmt"

type Method string
type Pattern string

type Routes map[Pattern]map[Method]Handler

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

	if rs[Pattern(pattern)] == nil {
		rs[Pattern(pattern)] = make(map[Method]Handler)
	}

	rs[Pattern(pattern)][Method(method)] = h
	return nil
}

// Checks if route in routes
func (rs Routes) Has(pattern, method string) bool {
	_, ok := rs[Pattern(pattern)][Method(method)]
	return ok
}

type ErrRouteAlreadyExists struct {
	pattern string
	method  string
}

func (e ErrRouteAlreadyExists) Error() string {
	return fmt.Sprintf("route with pattern '%s' and method '%s' already exists", e.pattern, e.method)
}
