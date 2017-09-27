package yac_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

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

func getMuxHandlerForMethods(methods ...string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		for _, m := range methods {
			if req.Method == m {
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
		router.Handle(c.path, getMuxHandlerForMethods(c.methods...))
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
