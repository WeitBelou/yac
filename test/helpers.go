package yac

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"

	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/weitbelou/yac"
)

type route struct {
	method  string
	pattern string
	path    string
	params  string // As json
}

func (r route) String() string {
	return fmt.Sprintf("\t%s\n\t%s\n\t%s", r.method, r.pattern, r.path)
}

// Creates empty request and response writer
func createRequestResponse() (*http.Request, *httptest.ResponseRecorder) {
	return httptest.NewRequest("", "/", nil), httptest.NewRecorder()
}

// Resets request and response to new method and path
func resetRequestResponse(req *http.Request, w *httptest.ResponseRecorder, method, path string) {
	req.Method = method
	req.URL.Path = path

	w.Body = new(bytes.Buffer)
	w.HeaderMap = make(http.Header)
}

// Initialize router with list of routes
func createRouter(routes []route, handler http.HandlerFunc) (http.Handler, error) {
	router, err := yac.NewRouter("")
	if err != nil {
		return nil, fmt.Errorf("can not create router: %v", err)
	}

	for _, route := range routes {
		if err := router.Route(route.pattern, route.method, handler); err != nil {
			return nil, fmt.Errorf("can not init route '%+v': %v", route, err)
		}
	}

	return router, nil
}

// Empty handler to return 200 for resolved routes.
func emptyHandler(_ http.ResponseWriter, _ *http.Request) {}

// Params handler
func paramsHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params, err := yac.Params(req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	js, err := json.Marshal(params.PathParams)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Write(js)
}

// Helper for 'resolve' tests
func testResolve(t *testing.T, routes []route) {
	router, err := createRouter(routes, emptyHandler)
	require.Nil(t, err, "can not create router: %v", err)

	req, w := createRequestResponse()

	for _, route := range routes {
		resetRequestResponse(req, w, route.method, route.path)
		router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusOK, w.Code, "can not resolve route %s", route)
	}
}

// Helper for 'resolve' benchmarks
func benchResolve(b *testing.B, routes []route) {
	router, err := createRouter(routes, emptyHandler)
	require.Nil(b, err, "can not create router: %v", err)

	req, w := createRequestResponse()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, route := range routes {
			resetRequestResponse(req, w, route.method, route.path)
			router.ServeHTTP(w, req)
		}
	}
}

// Helper for 'params' tests
func testParams(t *testing.T, routes []route) {
	router, err := createRouter(routes, paramsHandler)
	require.Nil(t, err, "can not create router: %v", err)

	req, w := createRequestResponse()
	for _, route := range routes {
		resetRequestResponse(req, w, route.method, route.path)
		router.ServeHTTP(w, req)

		assert.JSONEq(t, route.params, w.Body.String(),
			"invalid params for \n%s", route)
	}
}
