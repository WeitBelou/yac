package yac

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRoutes_HasOnEmpty(t *testing.T) {
	rs := NewRoutes()
	assert.False(t, rs.Has("/users", "GET"))
}

func TestRoutes_Add(t *testing.T) {
	pattern := "/users"
	method := "GET"

	rs := NewRoutes()
	rs.Add(pattern, method, emptyHandler)

	assert.True(t, rs.Has(pattern, method))
}
