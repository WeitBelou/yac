package yac

import "testing"

var githubPullRequests = []route{
	{method: "GET", pattern: "/repos/{str:owner}/{str:repo}/pulls",
		path: "/repos/test/test/pulls", params: `{"owner":"test","repo":"test"}`},
	{method: "GET", pattern: "/repos/{str:owner}/{str:repo}/pulls/{str:user}",
		path: "/repos/test/test/pulls/test", params: `{"owner":"test","repo":"test","user":"test"}`},
	{method: "POST", pattern: "/repos/{str:owner}/{str:repo}/pulls",
		path: "/repos/test/test/pulls", params: `{"owner":"test","repo":"test"}`},
	{method: "PATCH", pattern: "/repos/{str:owner}/{str:repo}/pulls/{str:user}",
		path: "/repos/test/test/pulls/test", params: `{"owner":"test","repo":"test","user":"test"}`},
	{method: "GET", pattern: "/repos/{str:owner}/{str:repo}/pulls/{str:user}/commits",
		path: "/repos/test/test/pulls/test/commits", params: `{"owner":"test","repo":"test","user":"test"}`},
	{method: "GET", pattern: "/repos/{str:owner}/{str:repo}/pulls/{str:user}/files",
		path: "/repos/test/test/pulls/test/files", params: `{"owner":"test","repo":"test","user":"test"}`},
	{method: "GET", pattern: "/repos/{str:owner}/{str:repo}/pulls/{str:user}/merge",
		path: "/repos/test/test/pulls/test/merge", params: `{"owner":"test","repo":"test","user":"test"}`},
	{method: "PUT", pattern: "/repos/{str:owner}/{str:repo}/pulls/{str:user}/merge",
		path: "/repos/test/test/pulls/test/merge", params: `{"owner":"test","repo":"test","user":"test"}`},
	{method: "GET", pattern: "/repos/{str:owner}/{str:repo}/pulls/{str:user}/comments",
		path: "/repos/test/test/pulls/test/comments", params: `{"owner":"test","repo":"test","user":"test"}`},
	{method: "GET", pattern: "/repos/{str:owner}/{str:repo}/pulls/comments",
		path: "/repos/test/test/pulls/comments", params: `{"owner":"test","repo":"test"}`},
	{method: "GET", pattern: "/repos/{str:owner}/{str:repo}/pulls/comments/{str:user}",
		path: "/repos/test/test/pulls/comments/test", params: `{"owner":"test","repo":"test","user":"test"}`},
	{method: "PUT", pattern: "/repos/{str:owner}/{str:repo}/pulls/{str:user}/comments",
		path: "/repos/test/test/pulls/test/comments", params: `{"owner":"test","repo":"test","user":"test"}`},
	{method: "PATCH", pattern: "/repos/{str:owner}/{str:repo}/pulls/comments/{str:user}",
		path: "/repos/test/test/pulls/comments/test", params: `{"owner":"test","repo":"test","user":"test"}`},
	{method: "DELETE", pattern: "/repos/{str:owner}/{str:repo}/pulls/comments/{str:user}",
		path: "/repos/test/test/pulls/comments/test", params: `{"owner":"test","repo":"test","user":"test"}`},
}

// Tests that all 'pull requests' API routes resolves correctly
func TestGitHubResolvePullRequests(t *testing.T) {
	testResolve(t, githubPullRequests)
}

// Bench for 'pull requests' API resolving
func BenchmarkGithubResolvePullRequests(b *testing.B) {
	benchResolve(b, githubPullRequests)
}

// Test for 'pull requests' API params parsing
func TestGitHubParamsPullRequests(t *testing.T) {
	testParams(t, githubPullRequests)
}
