package yac

import (
	"log"
	"net/http"
	"regexp"
	"time"
)

// Wrapper that takes route and returns modified route
type Wrapper func(route Route) Route

// Helper for wrappers slice
type Wrappers []Wrapper

// Wraps route with slice of wrappers
func (ws Wrappers) Wrap(r Route) Route {
	for _, w := range ws {
		r = w(r)
	}
	return r
}

// Logs requests
func Logger(r Route) Route {
	r.Handler = func(w http.ResponseWriter, req *http.Request) {
		start := time.Now()
		defer log.Printf(
			"%s\t%s\t%s",
			req.Method,
			req.RequestURI,
			time.Since(start),
		)

		r.Handler(w, req)
	}
	return r
}

// Compiles pattern of route
func patternCompiler(route Route) Route {
	re, err := regexp.Compile(convertSimplePatternToRegexp(route.Pattern))
	if err != nil {
		log.Panicf("can not compile pattern: %v", err)
	}

	return Route{
		Method:  route.Method,
		Pattern: route.Pattern,
		Handler: route.Handler,
		matcher: re,
	}
}

// Inserts params into request's context
func paramsInserter(route Route) Route {
	newRoute := route
	newRoute.Handler = func(w http.ResponseWriter, req *http.Request) {
		params, err := newParams(req, route.matcher, req.URL.Path)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		route.Handler(w, putParamsToRequest(req, params))
	}
	return newRoute
}
