package yac

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"

	"github.com/weitbelou/yac"
)

type routeTestCase struct {
	method  string
	pattern string

	path       string
	pathParams map[string]string
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
func createRouter(routes []routeTestCase, handler http.HandlerFunc) (http.Handler, error) {
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
