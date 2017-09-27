package yac_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/weitbelou/yac"
)

var emptyHandlerFunc = http.HandlerFunc(func(http.ResponseWriter, *http.Request) {})

func TestRouterRouteExists(t *testing.T) {
	router := yac.Router{}

	err := router.Handle(http.MethodGet, "/", emptyHandlerFunc)
	require.Nil(t, err, "can not add route to empty router: %v", err)

	err = router.Handle(http.MethodGet, "/", emptyHandlerFunc)
	assert.NotNil(t, err, "must return error if we trying to add duplicated route")
	assert.Equal(t, "route with path '/' and method 'GET' already exists", err.Error(), "unexpected type")
}

func TestRouterInvalidRoute(t *testing.T) {
	router := yac.Router{}

	err := router.Handle("", "/", emptyHandlerFunc)
	require.NotNil(t, err, "can not add route to empty router: %v", err)
	assert.Equal(t, "invalid route: [empty method]", err.Error())
}

var echoHandlerFunc = http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "%s:%s", req.Method, req.URL.Path)
})

func TestRouterStaticPath(t *testing.T) {
	cases := []struct {
		path   string
		method string
	}{
		{"/", http.MethodGet},
		{"/", http.MethodPost},
		{"/users", http.MethodGet},
		{"/users", http.MethodPost},
	}

	router := yac.Router{}
	for _, c := range cases {
		err := router.Handle(c.method, c.path, echoHandlerFunc)
		require.Nil(t, err, "can not add route: %v", err)
	}

	for _, c := range cases {
		w := httptest.NewRecorder()
		req := httptest.NewRequest(c.method, c.path, nil)

		router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusOK, w.Code, "failed to respond")
		assert.Equal(t, fmt.Sprintf("%s:%s", req.Method, req.URL.Path), w.Body.String())
	}
}

func TestRouterNotFound(t *testing.T) {
	router := yac.Router{}

	err := router.Handle(http.MethodGet, "/", emptyHandlerFunc)
	require.Nil(t, err, "can not add route: %v", err)

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/users", nil)

	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusNotFound, w.Code)
	assert.JSONEq(t, `{"status": 404, "error": "'/users' not found"}`, w.Body.String())
}

func TestRouterMethodNotAllowed(t *testing.T) {
	router := yac.Router{}

	err := router.Handle(http.MethodGet, "/", emptyHandlerFunc)
	require.Nil(t, err, "can not add route: %v", err)

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/", nil)

	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusMethodNotAllowed, w.Code)
	assert.JSONEq(t, `{"status": 405, "error": "method 'POST' not allowed on path '/'"}`, w.Body.String())
}

func TestRouterConcurrentUse(t *testing.T) {
	router := yac.Router{}

	cases := []struct {
		path   string
		method string
	}{
		{"/", http.MethodGet},
		{"/", http.MethodPost},
		{"/users", http.MethodGet},
		{"/users", http.MethodPost},
	}

	for _, c := range cases {
		err := router.Handle(c.method, c.path, echoHandlerFunc)
		require.Nil(t, err, "can not init router")
	}

	n := 1000
	wg := sync.WaitGroup{}
	wg.Add(n)

	for i := 0; i < n; i++ {
		go func() {
			defer wg.Done()

			for _, c := range cases {
				w := httptest.NewRecorder()
				req := httptest.NewRequest(c.method, c.path, nil)

				router.ServeHTTP(w, req)
				assert.Equal(t, http.StatusOK, w.Code, "failed to respond")
				assert.Equal(t, fmt.Sprintf("%s:%s", req.Method, req.URL.Path), w.Body.String())
			}
		}()
	}

	wg.Wait()
}

func TestRouterConcurrentInit(t *testing.T) {
	path := "/"
	method := http.MethodGet
	handler := emptyHandlerFunc

	router := yac.Router{}
	router.Handle(method, path, handler)

	n := 1000

	wg := sync.WaitGroup{}
	wg.Add(n)

	for i := 0; i < n; i++ {
		go func() {
			defer wg.Done()
			err := router.Handle(method, path, handler)
			assert.NotNil(t, err, "can init router twice")
		}()

	}
	wg.Wait()
}
