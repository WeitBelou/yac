package yac

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRoutesHasOnEmpty(t *testing.T) {
	rs := NewRoutes()
	assert.False(t, rs.Has("/users", "GET"), "can find route in empty routes")
}

func TestRoutesAdd(t *testing.T) {
	pattern := "/users"
	method := "GET"

	rs := NewRoutes()
	rs.Add(pattern, method, emptyHandler)

	assert.True(t, rs.Has(pattern, method))
}

func TestRoutesAddDuplicated(t *testing.T) {
	pattern := "/users"
	method := "GET"

	rs := NewRoutes()

	err := rs.Add(pattern, method, emptyHandler)
	assert.Nil(t, err, "can not insert route in empty routes: %v", err)

	err = rs.Add(pattern, method, emptyHandler)
	assert.NotNil(t, err, "can insert duplicated item")
}
