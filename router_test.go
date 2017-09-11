package yac

import (
	"net/http"
	"testing"

	"net/http/httptest"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var usersHandler = HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
	w.Write([]byte("users"))
})

var staticRoutes = Routes{
	"/users": Methods{
		http.MethodGet:  usersHandler,
		http.MethodPost: emptyHandlerFunc,
	},
	"/about": Methods{
		http.MethodGet: emptyHandlerFunc,
	},
}

// Helper to create router from routes.
func createRouter(t *testing.T, routes Routes) Router {
	router := NewRouter()
	for pattern, methods := range routes {
		for method, handler := range methods {
			err := router.Route(pattern, method, handler)
			require.Nil(t, err, "can not set route '%s' '%s': %v", method, pattern, err)
		}
	}
	return router
}

func TestRouterResolveStatic(t *testing.T) {
	router := createRouter(t, staticRoutes)

	for pattern, methods := range staticRoutes {
		for method := range methods {
			req := httptest.NewRequest(method, pattern, nil)
			w := httptest.NewRecorder()

			router.ServeHTTP(w, req)
			assert.Equal(t, http.StatusOK, w.Code, "can not resolve route '%s' '%s'", method, pattern)
		}
	}
}

func TestRouterResolveNotFound(t *testing.T) {
	router := createRouter(t, staticRoutes)

	req := httptest.NewRequest(http.MethodGet, "/notfound", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusNotFound, w.Code, "shouldn't be found")
}

func TestRouterResolveMethodNotAllowed(t *testing.T) {
	router := createRouter(t, staticRoutes)

	req := httptest.NewRequest(http.MethodPatch, "/users", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusMethodNotAllowed, w.Code, "shouldn't be found")
}
