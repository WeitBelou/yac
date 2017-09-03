package yac

import "testing"

var githubUsers = []route{
	{method: "GET", pattern: "/user",
		path: "/user", params: `{}`},
	{method: "PATCH", pattern: "/user",
		path: "/user", params: `{}`},
	{method: "GET", pattern: "/users",
		path: "/users", params: `{}`},
	{method: "GET", pattern: "/user/emails",
		path: "/user/emails", params: `{}`},
	{method: "POST", pattern: "/user/emails",
		path: "/user/emails", params: `{}`},
	{method: "DELETE", pattern: "/user/emails",
		path: "/user/emails", params: `{}`},
	{method: "GET", pattern: "/user/followers",
		path: "/user/followers", params: `{}`},
	{method: "GET", pattern: "/user/following",
		path: "/user/following", params: `{}`},
	{method: "GET", pattern: "/user/keys",
		path: "/user/keys", params: `{}`},
	{method: "POST", pattern: "/user/keys",
		path: "/user/keys", params: `{}`},
}

// Tests that all 'users' API routes resolves correctly
func TestGitHubResolveUsers(t *testing.T) {
	testResolve(t, githubUsers)
}

// Bench for 'users' API resolving
func BenchmarkGithubResolveUsers(b *testing.B) {
	benchResolve(b, githubUsers)
}

// Test for 'users' API params parsing
func TestGitHubParamsUsers(t *testing.T) {
	testParams(t, githubUsers)
}
