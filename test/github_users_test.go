package yac

import "testing"

var githubUsers = []route{
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
