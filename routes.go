package yac

import "fmt"

type Route map[string]Handler

type Routes map[string]Route

func NewRoutes() Routes {
	return make(map[string]Route)
}

func (r Routes) Add(pattern, method string, handler Handler) error {
	return nil
}

func (r Routes) Has(pattern, method string) bool {
	return false
}

type ErrRouteAlreadyExists struct {
	pattern string
	method  string
}

func (e ErrRouteAlreadyExists) Error() string {
	return fmt.Sprintf("route with pattern '%s' and method '%s' already exists", e.pattern, e.method)
}
