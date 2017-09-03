package yac

import "testing"

var githubSearch = []route{
	{method: "GET", pattern: "/search/repositories",
		path: "/search/repositories"},
	{method: "GET", pattern: "/search/code",
		path: "/search/code"},
	{method: "GET", pattern: "/search/issues",
		path: "/search/issues"},
	{method: "GET", pattern: "/search/users",
		path: "/search/users"},
}

// Tests that all 'search' API routes resolves correctly
func TestGitHubResolveSearch(t *testing.T) {
	testResolve(t, githubSearch)
}

// Bench for 'search' API resolving
func BenchmarkGithubResolveSearch(b *testing.B) {
	benchResolve(b, githubSearch)
}

// Test for 'search' API params parsing
func TestGitHubParamsSearch(t *testing.T) {
	testParams(t, githubSearch)
}
