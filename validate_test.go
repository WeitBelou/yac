package yac

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

var emptyHandler = http.HandlerFunc(func(_ http.ResponseWriter, _ *http.Request) {})

func TestValidateRoute(t *testing.T) {
	cases := []struct {
		method  string
		path    string
		handler http.Handler
		errs    []error
	}{
		{errs: []error{errEmptyMethod, errEmptyPath, errNilHandler}},
		{path: "/", errs: []error{errEmptyMethod, errNilHandler}},
		{method: "GET", errs: []error{errEmptyPath, errNilHandler}},
		{handler: emptyHandler, errs: []error{errEmptyMethod, errEmptyPath}},
		{path: "/", handler: emptyHandler, errs: []error{errEmptyMethod}},
		{method: "GET", handler: emptyHandler, errs: []error{errEmptyPath}},
		{method: "GET", path: "/", errs: []error{errNilHandler}},
	}

	for _, c := range cases {
		errs := validateRoute(c.method, c.path, c.handler)
		assert.Equal(t, c.errs, errs)
	}
}
