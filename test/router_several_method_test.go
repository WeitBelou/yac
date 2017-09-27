package yac_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/weitbelou/yac"
)

var cases = []struct {
	method string
	path   string
}{
	{http.MethodGet, "/users"},
	{http.MethodPost, "/users"},
	{http.MethodDelete, "/users"},
	{http.MethodPatch, "/users"},
	{http.MethodGet, "/users/messages"},
	{http.MethodPost, "/users/messages"},
}

func BenchmarkYacRouterSeveral(b *testing.B) {
	router := yac.Router{}
	for _, r := range cases {
		router.Handle(r.method, r.path, emptyHandlerFunc)
	}

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/", nil)

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, c := range cases {
			req.Method = c.method
			req.URL.Path = c.path

			router.ServeHTTP(w, req)
		}
	}
}

func getMuxHandlerForMethods(base http.Handler, methods []string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		for _, m := range methods {
			if req.Method == m {
				base.ServeHTTP(w, req)
				return
			}
		}
		http.NotFound(w, req)
	})
}

func BenchmarkHTTPMuxSeveral(b *testing.B) {
	cases := []struct {
		path    string
		methods []string
	}{
		{"/users", []string{http.MethodGet, http.MethodPost}},
		{"/users/messages", []string{http.MethodGet, http.MethodPost, http.MethodDelete, http.MethodPatch}},
	}

	router := http.ServeMux{}
	for _, c := range cases {
		router.Handle(c.path, getMuxHandlerForMethods(emptyHandlerFunc, c.methods))
	}

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/", nil)

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, c := range cases {
			req.URL.Path = c.path
			for _, m := range c.methods {
				req.Method = m
				router.ServeHTTP(w, req)
			}
		}
	}
}

func TestHTTPMuxSeveral(t *testing.T) {
	cases := []struct {
		path    string
		methods []string
	}{
		{"/users", []string{http.MethodGet, http.MethodPost}},
		{"/users/messages", []string{http.MethodGet, http.MethodPost, http.MethodDelete, http.MethodPatch}},
	}

	router := http.ServeMux{}
	for _, c := range cases {
		router.Handle(c.path, getMuxHandlerForMethods(echoHandlerFunc, c.methods))
	}

	for _, c := range cases {
		for _, m := range c.methods {
			w := httptest.NewRecorder()
			req := httptest.NewRequest(m, c.path, nil)

			router.ServeHTTP(w, req)
			assert.Equal(t, http.StatusOK, w.Result().StatusCode)
			assert.Equal(t, fmt.Sprintf("%s:%s", m, c.path), w.Body.String())
		}
	}
}

func TestYacRouterSeveral(t *testing.T) {
	router := yac.Router{}
	for _, r := range cases {
		router.Handle(r.method, r.path, echoHandlerFunc)
	}

	for _, c := range cases {
		w := httptest.NewRecorder()
		req := httptest.NewRequest(c.method, c.path, nil)

		router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusOK, w.Result().StatusCode)
		assert.Equal(t, fmt.Sprintf("%s:%s", c.method, c.path), w.Body.String())
	}
}
