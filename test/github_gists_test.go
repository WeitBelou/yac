package yac

import "testing"

var githubGists = []route{
	{method: "GET", pattern: "/gists",
		path: "/gists", params: `{}`},
	{method: "GET", pattern: "/gists/public",
		path: "/gists/public", params: `{}`},
	{method: "GET", pattern: "/gists/starred",
		path: "/gists/starred", params: `{}`},
	{method: "POST", pattern: "/gists",
		path: "/gists", params: `{}`},
	{method: "GET", pattern: "/users/{str:user}/gists",
		path: "/users/test/gists", params: `{"user": "test"}`},
	{method: "GET", pattern: "/gists/{int:id}",
		path: "/gists/123", params: `{"id": "123"}`},
	{method: "PATCH", pattern: "/gists/{int:id}",
		path: "/gists/123", params: `{"id": "123"}`},
	{method: "PUT", pattern: "/gists/{int:id}/star",
		path: "/gists/123/star", params: `{"id": "123"}`},
	{method: "DELETE", pattern: "/gists/{int:id}/star",
		path: "/gists/123/star", params: `{"id": "123"}`},
	{method: "GET", pattern: "/gists/{int:id}/star",
		path: "/gists/123/star", params: `{"id": "123"}`},
	{method: "POST", pattern: "/gists/{int:id}/forks",
		path: "/gists/123/forks", params: `{"id": "123"}`},
	{method: "DELETE", pattern: "/gists/{int:id}",
		path: "/gists/123", params: `{"id": "123"}`},
}

// Tests that all 'gists' API routes resolves correctly
func TestGitHubResolveGists(t *testing.T) {
	testResolve(t, githubGists)
}

// Bench for 'gists' API resolving
func BenchmarkGithubResolveGists(b *testing.B) {
	benchResolve(b, githubGists)
}

// Test for 'gists' API params parsing
func TestGitHubParamsGists(t *testing.T) {
	testParams(t, githubGists)
}
