package yac

import (
	"net/http"
	"testing"

	"net/http/httptest"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type route struct {
	method  string
	pattern string
}

var staticRoutes = []route{
	{http.MethodGet, "/users"},
	{http.MethodPost, "/users"},
	{http.MethodGet, "/about"},
}

func TestRouterResolveStatic(t *testing.T) {
	router := NewRouter()
	for _, r := range staticRoutes {
		err := router.Route(r.pattern, r.method, emptyHandlerFunc)
		require.Nil(t, err, "can not set route %+v: %v", r, err)
	}

	for _, r := range staticRoutes {
		req := httptest.NewRequest(r.method, r.pattern, nil)
		w := httptest.NewRecorder()

		router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusOK, w.Code, "can not resolve route %+v", r)
	}
}

func TestRouterResolveNotFound(t *testing.T) {
	router := NewRouter()
	for _, r := range staticRoutes {
		err := router.Route(r.pattern, r.method, emptyHandlerFunc)
		require.Nil(t, err, "can not set route %+v: %v", r, err)
	}

	req := httptest.NewRequest(http.MethodGet, "/notfound", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusNotFound, w.Code, "shouldn't be found")
}
