package yac

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestExtractPathParams(t *testing.T) {
	pattern := regexp.MustCompile(`/users/(?P<id>\d+)`)
	pathParams := extractPathParams(pattern, "/users/12")

	expectedPathParams := map[string]interface{}{"id": "12"}

	assert.Equal(t, expectedPathParams, pathParams)
}

func TestPostBody(t *testing.T) {
	bodyBuf := bytes.NewBufferString(`{"some": ["json-like", "body"]}`)

	const path = "/"
	req := httptest.NewRequest(http.MethodPost, path, bytes.NewBuffer(bodyBuf.Bytes()))

	p, err := newParams(req, regexp.MustCompile(path), path)
	if err != nil {
		t.Fail()
	}

	assert.JSONEq(t, bodyBuf.String(), string(p.Body))
}

func TestNilBody(t *testing.T) {
	const path = "/"
	req := httptest.NewRequest(http.MethodPost, path, nil)

	p, err := newParams(req, regexp.MustCompile(path), path)
	require.Nil(t, err, "can not extract params: %v", err)

	assert.Equal(t, p.Body, make([]byte, 0))
}
