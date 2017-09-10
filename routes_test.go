package yac

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

var emptyHandlerFunc = HandlerFunc(func(_ http.ResponseWriter, _ *http.Request) {})

func TestRoutes_Has_OnEmpty(t *testing.T) {
	rs := NewRoutes()
	assert.False(t, rs.Has("/users", "GET"), "can find route in empty routes")
}

func TestRoutes_Add_InEmpty(t *testing.T) {
	pattern := "/users"
	method := "GET"

	rs := NewRoutes()
	rs.Add(pattern, method, emptyHandlerFunc)

	assert.True(t, rs.Has(pattern, method))
}

func TestRoutes_Add_Duplicated(t *testing.T) {
	pattern := "/users"
	method := "GET"

	rs := NewRoutes()

	err := rs.Add(pattern, method, emptyHandlerFunc)
	assert.Nil(t, err, "can not insert route in empty routes: %v", err)

	err = rs.Add(pattern, method, emptyHandlerFunc)
	assert.NotNil(t, err, "can insert duplicated item")
}
