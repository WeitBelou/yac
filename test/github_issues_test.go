package yac

import (
	"testing"
)

var githubIssues = []route{
	{method: "GET", pattern: "/issues",
		path: "/issues"},
	{method: "GET", pattern: "/user/issues",
		path: "/user/issues"},
	{method: "GET", pattern: "/orgs/{str:org}/issues",
		path: "/orgs/test/issues", params: `{"org":"test"}`},
	{method: "GET", pattern: "/repos/{str:owner}/{str:repo}/issues",
		path: "/repos/test/test/issues", params: `{"owner": "test", "repo": "test"}`},
	{method: "GET", pattern: "/repos/{str:owner}/{str:repo}/issues/{str:user}",
		path: "/repos/test/test/issues/test", params: `{"owner": "test", "repo": "test","user":"test"}`},
	{method: "POST", pattern: "/repos/{str:owner}/{str:repo}/issues",
		path: "/repos/test/test/issues", params: `{"owner": "test", "repo": "test"}`},
	{method: "PATCH", pattern: "/repos/{str:owner}/{str:repo}/issues/{str:user}",
		path: "/repos/test/test/issues/test", params: `{"owner": "test", "repo": "test", "user": "test"}`},
	{method: "GET", pattern: "/repos/{str:owner}/{str:repo}/assignees",
		path: "/repos/test/test/assignees", params: `{"owner": "test", "repo": "test"}`},
	{method: "GET", pattern: "/repos/{str:owner}/{str:repo}/assignees/{str:assignee}",
		path: "/repos/test/test/assignees/test", params: `{"owner": "test", "repo": "test", "assignee": "test"}`},
	{method: "GET", pattern: "/repos/{str:owner}/{str:repo}/issues/{str:user}/comments",
		path: "/repos/test/test/issues/test/comments", params: `{"owner": "test","repo": "test","user":"test"}`},
	{method: "GET", pattern: "/repos/{str:owner}/{str:repo}/issues/comments",
		path: "/repos/test/test/issues/comments", params: `{"owner": "test", "repo": "test"}`},
	{method: "GET", pattern: "/repos/{str:owner}/{str:repo}/issues/comments/{int:id}",
		path: "/repos/test/test/issues/comments/123", params: `{"owner":"test","repo":"test","id":"123"}`},
	{method: "POST", pattern: "/repos/{str:owner}/{str:repo}/issues/{str:user}/comments",
		path: "/repos/test/test/issues/test/comments", params: `{"owner":"test","repo":"test","user":"test"}`},
	{method: "PATCH", pattern: "/repos/{str:owner}/{str:repo}/issues/comments/{int:id}",
		path: "/repos/test/test/issues/comments/123", params: `{"owner":"test","repo":"test","id":"123"}`},
	{method: "DELETE", pattern: "/repos/{str:owner}/{str:repo}/issues/comments/{int:id}",
		path: "/repos/test/test/issues/comments/123", params: `{"owner":"test","repo":"test","id":"123"}`},
	{method: "GET", pattern: "/repos/{str:owner}/{str:repo}/issues/{str:user}/events",
		path: "/repos/test/test/issues/test/events", params: `{"owner":"test","repo":"test","user":"test"}`},
	{method: "GET", pattern: "/repos/{str:owner}/{str:repo}/issues/events",
		path: "/repos/test/test/issues/events", params: `{"owner":"test","repo":"test"}`},
	{method: "GET", pattern: "/repos/{str:owner}/{str:repo}/issues/events/{int:id}",
		path: "/repos/test/test/issues/events/123", params: `{"owner":"test","repo":"test","id":"123"}`},
	{method: "GET", pattern: "/repos/{str:owner}/{str:repo}/labels",
		path: "/repos/test/test/labels", params: `{"owner":"test","repo":"test"}`},
	{method: "GET", pattern: "/repos/{str:owner}/{str:repo}/labels/{str:name}",
		path: "/repos/test/test/labels/test", params: `{"owner":"test","repo":"test","name":"test"}`},
	{method: "POST", pattern: "/repos/{str:owner}/{str:repo}/labels",
		path: "/repos/test/test/labels", params: `{"owner":"test","repo":"test"}`},
	{method: "PATCH", pattern: "/repos/{str:owner}/{str:repo}/labels/{str:name}",
		path: "/repos/test/test/labels/test", params: `{"owner":"test","repo":"test","name":"test"}`},
	{method: "DELETE", pattern: "/repos/{str:owner}/{str:repo}/labels/{str:name}",
		path: "/repos/test/test/labels/test", params: `{"owner":"test","repo":"test","name":"test"}`},
	{method: "GET", pattern: "/repos/{str:owner}/{str:repo}/issues/{str:user}/labels",
		path: "/repos/test/test/issues/test/labels", params: `{"owner":"test","repo":"test","user":"test"}`},
	{method: "POST", pattern: "/repos/{str:owner}/{str:repo}/issues/{str:user}/labels",
		path: "/repos/test/test/issues/test/labels", params: `{"owner":"test","repo":"test","user":"test"}`},
	{method: "DELETE", pattern: "/repos/{str:owner}/{str:repo}/issues/{str:user}/labels/{str:name}",
		path: "/repos/test/test/issues/test/labels/test",
		params: `{"owner":"test","repo":"test","name":"test","user":"test"}`},
	{method: "PUT", pattern: "/repos/{str:owner}/{str:repo}/issues/{str:user}/labels",
		path: "/repos/test/test/issues/test/labels", params: `{"owner":"test","repo":"test","user":"test"}`},
	{method: "DELETE", pattern: "/repos/{str:owner}/{str:repo}/issues/{str:user}/labels",
		path: "/repos/test/test/issues/test/labels", params: `{"owner":"test","repo":"test","user":"test"}`},
	{method: "GET", pattern: "/repos/{str:owner}/{str:repo}/milestones/{str:user}/labels",
		path: "/repos/test/test/milestones/test/labels", params: `{"owner":"test","repo":"test","user":"test"}`},
	{method: "GET", pattern: "/repos/{str:owner}/{str:repo}/milestones",
		path: "/repos/test/test/milestones", params: `{"owner":"test","repo":"test"}`},
	{method: "GET", pattern: "/repos/{str:owner}/{str:repo}/milestones/{str:user}",
		path: "/repos/test/test/milestones/test", params: `{"owner":"test","repo":"test","user":"test"}`},
	{method: "POST", pattern: "/repos/{str:owner}/{str:repo}/milestones",
		path: "/repos/test/test/milestones", params: `{"owner":"test","repo":"test"}`},
	{method: "PATCH", pattern: "/repos/{str:owner}/{str:repo}/milestones/{str:user}",
		path: "/repos/test/test/milestones/test", params: `{"owner":"test","repo":"test","user":"test"}`},
	{method: "DELETE", pattern: "/repos/{str:owner}/{str:repo}/milestones/{str:user}",
		path: "/repos/test/test/milestones/test", params: `{"owner":"test","repo":"test","user":"test"}`},
}

// Tests that all 'issues' API routes resolves correctly
func TestGitHubResolveIssues(t *testing.T) {
	testResolve(t, githubIssues)
}

// Bench for 'issues' API resolving
func BenchmarkGithubResolveIssues(b *testing.B) {
	benchResolve(b, githubIssues)
}

// Test for 'issues' API params parsing
func TestGitHubParamsIssues(t *testing.T) {
	testParams(t, githubIssues)
}
