package yac

import (
	"fmt"
	"net/http"

	"regexp"
	"strings"

	"github.com/armon/go-radix"
)

type Route struct {
	Methods map[string]http.HandlerFunc
}

func NewRoute(method string, h http.HandlerFunc) *Route {
	return &Route{Methods: map[string]http.HandlerFunc{method: h}}
}

// Add method's handler returns error if method already exists.
func (r *Route) AddMethod(method string, h http.HandlerFunc) error {
	if _, ok := r.Methods[method]; ok {
		return fmt.Errorf("'%s' already has handler", method)
	}

	r.Methods[method] = h
	return nil
}

// Helper for routes tree
type Routes struct {
	tree *radix.Tree
}

func NewRoutes() *Routes {
	return &Routes{tree: radix.New()}
}

// Inserts route and returns
func (rs *Routes) Insert(pattern, method string, h http.HandlerFunc) error {
	return nil
}

// Looks for best matching route
func (rs *Routes) Get(path string) (*Route, error) {
	// If there is static route with such path
	if route, ok := rs.tree.Get(path); ok {
		return UnsafeCastToRoutePointer(route), nil
	}

	return nil, nil
}

// Extracts param patterns from beginning of string and returns it + remaining part
func ExtractParamPattern(path string) (param string, remaining string, err error) {
	parts := strings.SplitN(path, "/", 2)

	// Check if params matches pattern
	if !paramPattern.MatchString(parts[0]) {
		return "", path, fmt.Errorf("first part '%s' doesn't match param pattern '%s'",
			parts[0], paramPattern)
	}

	if len(parts) == 1 {
		return parts[0], "", nil
	}

	return parts[0], parts[1], nil
}

var paramPattern = regexp.MustCompile(`{(int|str|hex|oid)([[:word:]]+)}`)
