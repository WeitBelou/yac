package yac

import (
	"errors"
	"net/http"
)

var (
	errEmptyMethod = errors.New("empty method")
	errEmptyPath   = errors.New("empty path")
	errNilHandler  = errors.New("nil handler")
)

// validateRoute checks route and returns slice of errors if any
func validateRoute(method string, path string, h http.Handler) []error {
	var errs []error
	if method == "" {
		errs = append(errs, errEmptyMethod)
	}
	if path == "" {
		errs = append(errs, errEmptyPath)
	}
	if h == nil {
		errs = append(errs, errNilHandler)
	}

	return errs
}
