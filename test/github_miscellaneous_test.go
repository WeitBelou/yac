package yac

import "testing"

var githubMiscellaneous = []route{
	{method: "GET", pattern: "/gitignore/templates/{str:name}",
		path: "/gitignore/templates/test", params: `{"name":"test"}`},
	{method: "GET", pattern: "/emojis",
		path: "/emojis", params: `{}`},
	{method: "GET", pattern: "/gitignore/templates",
		path: "/gitignore/templates", params: `{}`},
	{method: "POST", pattern: "/markdown",
		path: "/markdown", params: `{}`},
	{method: "POST", pattern: "/markdown/raw",
		path: "/markdown/raw", params: `{}`},
	{method: "GET", pattern: "/meta",
		path: "/meta", params: `{}`},
	{method: "GET", pattern: "/rate_limit",
		path: "/rate_limit", params: `{}`},
}

// Tests that all 'miscellaneous' API routes resolves correctly
func TestGitHubResolveMiscellaneous(t *testing.T) {
	testResolve(t, githubMiscellaneous)
}

// Bench for 'miscellaneous' API resolving
func BenchmarkGithubResolveMiscellaneous(b *testing.B) {
	benchResolve(b, githubMiscellaneous)
}

// Test for 'miscellaneous' API params parsing
func TestGitHubParamsMiscellaneous(t *testing.T) {
	testParams(t, githubMiscellaneous)
}
