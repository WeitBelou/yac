package yac_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/weitbelou/yac"
)

func BenchmarkYacRouterOne(b *testing.B) {
	cases := []string{
		"/users",
		"/users/messages",
		"/users/messages/info",
		"/users/stats",
		"/users/profile",
	}

	router := yac.Router{}
	for _, c := range cases {
		router.Handle(http.MethodPost, c, emptyHandlerFunc)
	}

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/", nil)

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, c := range cases {
			req.URL.Path = c

			router.ServeHTTP(w, req)
		}
	}
}

func BenchmarkHTTPMuxOne(b *testing.B) {
	cases := []string{
		"/users",
		"/users/messages",
		"/users/messages/info",
		"/users/stats",
		"/users/profile",
	}

	router := http.ServeMux{}
	for _, c := range cases {
		router.Handle(c, emptyHandlerFunc)
	}

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/", nil)

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, c := range cases {
			req.URL.Path = c
			router.ServeHTTP(w, req)
		}
	}
}

func TestYacRouterOne(t *testing.T) {
	cases := []string{
		"/users",
		"/users/messages",
		"/users/messages/info",
		"/users/stats",
		"/users/profile",
	}

	router := yac.Router{}
	for _, c := range cases {
		router.Handle(http.MethodPost, c, echoHandlerFunc)
	}

	for _, c := range cases {
		w := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodPost, c, nil)

		router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusOK, w.Result().StatusCode)
		assert.Equal(t, fmt.Sprintf("POST:%s", c), w.Body.String())
	}
}

func TestHTTPMuxOne(t *testing.T) {
	cases := []string{
		"/users",
		"/users/messages",
		"/users/messages/info",
		"/users/stats",
		"/users/profile",
	}

	router := http.ServeMux{}
	for _, c := range cases {
		router.Handle(c, echoHandlerFunc)
	}

	for _, c := range cases {
		w := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodPost, c, nil)

		router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusOK, w.Result().StatusCode)
		assert.Equal(t, fmt.Sprintf("POST:%s", c), w.Body.String())
	}
}
