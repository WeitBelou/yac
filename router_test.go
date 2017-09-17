package yac

import (
	"net/http"
	"testing"

	"net/http/httptest"

	"encoding/json"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/weitbelou/yac/params"
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

func TestRouterResponse(t *testing.T) {
	router := createRouter(t, staticRoutes)

	req := httptest.NewRequest(http.MethodGet, "/users", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, "users", w.Body.String(), "wrong response for 'GET' '/users'")
}

func TestRouterResolvesDynamic(t *testing.T) {
	var cases = []struct {
		pattern string
		path    string
	}{
		{"/users/{id}", "/users/123"},
		{"/users/{user_id}/posts/{post_id}", "/users/1231/posts/1234"},
		{"/users/{user_id}/comments/{post_id}", "/users/1231/comments/1234"},
	}

	router := NewRouter()
	for _, c := range cases {
		err := router.Route(c.pattern, http.MethodGet, patternEchoHandler{pattern: c.pattern})
		require.Nil(t, err, "can not set route 'GET' '%s': %v", c.pattern, err)
	}

	for _, c := range cases {
		req := httptest.NewRequest(http.MethodGet, c.path, nil)
		w := httptest.NewRecorder()

		router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusOK, w.Code, "can not resolve '%s'", c.path)
		assert.Equal(t, c.pattern, w.Body.String(), "path '%s' routes to wrong pattern", c.path)
	}
}

type patternEchoHandler struct {
	pattern string
}

func (h patternEchoHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte(h.pattern))
}

func TestRouterParams(t *testing.T) {
	var cases = []struct {
		pattern string
		path    string
		params  string
	}{
		{"/users/{id}", "/users/123", `{"id": "123"}`},
		{"/users/{user_id}/posts/{post_id}", "/users/1231/posts/1234",
			`{"user_id": "1231", "post_id": "1234"}`},
		{"/users/{user_id}/comments/{comment_id}", "/users/1231/comments/1234",
			`{"user_id": "1231", "comment_id": "1234"}`},
	}

	router := NewRouter()
	for _, c := range cases {
		err := router.Route(c.pattern, http.MethodGet, paramsEchoHandlerFunc)
		require.Nil(t, err, "can not set route 'GET' '%s': %v", c.pattern, err)
	}

	for _, c := range cases {
		req := httptest.NewRequest(http.MethodGet, c.path, nil)
		w := httptest.NewRecorder()

		router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusOK, w.Code, "can not resolve '%s'", c.path)
		assert.JSONEq(t, c.params, w.Body.String(), "for pattern '%s' and path'%s'", c.pattern, c.path)
	}
}

// Writes params as json to response.
var paramsEchoHandlerFunc = HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
	p := req.Context().Value(params.ContextKey).(params.Params)
	j, _ := json.Marshal(p)

	w.Write(j)
})
