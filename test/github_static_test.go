package yac

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var githubStatic = []route{
	// OAuth Authorizations
	{method: "GET", pattern: "/authorizations",
		path: "/authorizations"},
	{method: "POST", pattern: "/authorizations",
		path: "/authorizations"},

	// Activity
	{method: "GET", pattern: "/events",
		path: "/events"},
	{method: "GET", pattern: "/feeds",
		path: "/feeds"},
	{method: "GET", pattern: "/notifications",
		path: "/notifications"},
	{method: "PUT", pattern: "/notifications",
		path: "/notifications"},
	{method: "GET", pattern: "/user/starred",
		path: "/user/starred"},
	{method: "GET", pattern: "/user/subscriptions",
		path: "/user/subscriptions"},

	// Gists
	{method: "GET", pattern: "/gists",
		path: "/gists"},
	{method: "GET", pattern: "/gists/public",
		path: "/gists/public"},
	{method: "GET", pattern: "/gists/starred",
		path: "/gists/starred"},
	{method: "POST", pattern: "/gists",
		path: "/gists"},

	// Issues
	{method: "GET", pattern: "/issues",
		path: "/issues"},
	{method: "GET", pattern: "/user/issues",
		path: "/user/issues"},

	// Miscellaneous
	{method: "GET", pattern: "/emojis",
		path: "/emojis"},
	{method: "GET", pattern: "/gitignore/templates",
		path: "/gitignore/templates"},
	{method: "POST", pattern: "/markdown",
		path: "/markdown"},
	{method: "POST", pattern: "/markdown/raw",
		path: "/markdown/raw"},
	{method: "GET", pattern: "/meta",
		path: "/meta"},
	{method: "GET", pattern: "/rate_limit",
		path: "/rate_limit"},

	// Organizations
	{method: "GET", pattern: "/user/orgs",
		path: "/user/orgs"},
	{method: "GET", pattern: "/user/teams",
		path: "/user/teams"},

	// Repositories
	{method: "GET", pattern: "/user/repos",
		path: "/user/repos"},
	{method: "GET", pattern: "/repositories",
		path: "/repositories"},
	{method: "POST", pattern: "/user/repos",
		path: "/user/repos"},

	// Search
	{method: "GET", pattern: "/search/repositories",
		path: "/search/repositories"},
	{method: "GET", pattern: "/search/code",
		path: "/search/code"},
	{method: "GET", pattern: "/search/issues",
		path: "/search/issues"},
	{method: "GET", pattern: "/search/users",
		path: "/search/users"},

	// Users
	{method: "GET", pattern: "/user",
		path: "/user"},
	{method: "PATCH", pattern: "/user",
		path: "/user"},
	{method: "GET", pattern: "/users",
		path: "/users"},
	{method: "GET", pattern: "/user/emails",
		path: "/user/emails"},
	{method: "POST", pattern: "/user/emails",
		path: "/user/emails"},
	{method: "DELETE", pattern: "/user/emails",
		path: "/user/emails"},
	{method: "GET", pattern: "/user/followers",
		path: "/user/followers"},
	{method: "GET", pattern: "/user/following",
		path: "/user/following"},
	{method: "GET", pattern: "/user/keys",
		path: "/user/keys"},
	{method: "POST", pattern: "/user/keys",
		path: "/user/keys"},
}

// Tests that all static API routes resolves correctly
func TestGitHubStatic(t *testing.T) {
	router, err := createRouter(githubStatic, emptyHandler)
	require.Nil(t, err, "can not create router: %v", err)

	req, w := createRequestResponse()

	for _, route := range githubStatic {
		resetRequestResponse(req, w, route.method, route.path)
		router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusOK, w.Code, "can not resolve route %+v", route)
	}
}

// Bench for static API resolving
func BenchmarkGithubStatic(b *testing.B) {
	router, err := createRouter(githubStatic, emptyHandler)
	require.Nil(b, err, "can not create router: %v", err)

	req, w := createRequestResponse()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, route := range githubStatic {
			resetRequestResponse(req, w, route.method, route.path)
			router.ServeHTTP(w, req)
		}
	}
}
