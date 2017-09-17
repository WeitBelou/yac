package params

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExtractParamName(t *testing.T) {
	var cases = []struct {
		param string
		name  string
	}{
		{"{id}", "id"},
		{"{user_id}", "user_id"},
	}

	for _, c := range cases {
		name, err := extractParamName(c.param)
		assert.Nil(t, err, "can not extract params: %v", err)
		assert.Equal(t, c.name, name)
	}
}

func TestExtractParamsFromPath(t *testing.T) {
	var cases = []struct {
		pattern string
		path    string
		params  Params
	}{
		{"/users", "/users", Params{}},
		{"/users/{id}", "/users/12", Params{"id": "12"}},
		{"/users/{user_id}/posts/{post_id}", "/users/12/posts/123",
			Params{"user_id": "12", "post_id": "123"}},
	}

	for _, c := range cases {
		params, err := extractParams(c.pattern, c.path)
		assert.Nil(t, err, "can not extract params from path: %v", err)

		assert.Equal(t, c.params, params)
	}
}
