package yac

import (
	"reflect"
	"regexp"
	"testing"

	"encoding/json"

	"github.com/stretchr/testify/assert"
)

func TestSimplifiedPattern(t *testing.T) {
	pattern := regexp.MustCompile(convertSimplePatternToRegexp("/users/{hex:id}"))
	pathParams := extractPathParams(pattern, "/users/234feabc1357346781234524")

	expectedPathParams := map[string]interface{}{"id": "234feabc1357346781234524"}

	assert.Equal(t, expectedPathParams, pathParams)
}

func TestSimplifiedPatternEmpty(t *testing.T) {
	pattern := regexp.MustCompile(convertSimplePatternToRegexp("/users/{id}"))
	pathParams := extractPathParams(pattern, "/users/234feabc1357346781234524")

	expectedPathParams := map[string]interface{}{"id": "234feabc1357346781234524"}

	if !reflect.DeepEqual(pathParams, expectedPathParams) {
		t.Fatalf("expected: %v but got: %v", expectedPathParams, pathParams)
	}
}

func TestRelativePath(t *testing.T) {
	const basePath = "/api/v1"
	const absolutePath = "/api/v1/users/1234feabc1357346781234524"

	relPath, err := relativePath(basePath, absolutePath)
	if err != nil {
		t.Fatal(err)
	}

	if relPath != "/users/1234feabc1357346781234524" {
		t.Fail()
	}
}

func TestPatternConversion(t *testing.T) {
	pattern := "/repos/{str:owner}/{str:repo}/issues/comments"
	re := convertSimplePatternToRegexp(pattern)
	assert.Equal(t, `^/repos/(?P<owner>[[:alnum:]]+)/(?P<repo>[[:alnum:]]+)/issues/comments/?$`, re)
}

func TestExtractParams(t *testing.T) {
	path := "/repos/owner/repo/issues/comments"
	pattern := regexp.MustCompile(`^/repos/(?P<owner>[[:alnum:]]+)/(?P<repo>[[:alnum:]]+)/issues/comments/?$`)
	params := extractPathParams(pattern, path)

	js, err := json.Marshal(params)
	assert.Nil(t, err, "can not marshal params: %v", err)

	assert.JSONEq(t, `{"owner":"owner","repo":"repo"}`, string(js))
}
