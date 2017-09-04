package yac

import (
	"net/http"
	"regexp"
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
