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
