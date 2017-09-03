package yac

import "testing"

var githubActivity = []route{
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
	{method: "GET", pattern: "/repos/{str:owner}/{str:repo}/events",
		path: "/repos/test/test/events", params: `{"owner": "test", "repo": "test"}`},
	{method: "GET", pattern: "/networks/{str:owner}/{str:repo}/events",
		path: "/networks/test/test/events", params: `{"owner": "test", "repo": "test"}`},
	{method: "GET", pattern: "/orgs/{str:org}/events",
		path: "/orgs/test/events", params: `{"org": "test"}`},
	{method: "GET", pattern: "/users/{str:user}/received_events",
		path: "/users/test/received_events", params: `{"user": "test"}`},
	{method: "GET", pattern: "/users/{str:user}/received_events/public",
		path: "/users/test/received_events/public", params: `{"user": "test"}`},
	{method: "GET", pattern: "/users/{str:user}/events",
		path: "/users/test/events", params: `{"user": "test"}`},
	{method: "GET", pattern: "/users/{str:user}/events/public",
		path: "/users/test/events/public", params: `{"user": "test"}`},
	{method: "GET", pattern: "/users/{str:user}/events/orgs/{str:org}",
		path: "/users/test/events/orgs/test", params: `{"user": "test", "org": "test"}`},
	{method: "GET", pattern: "/repos/{str:owner}/{str:repo}/notifications",
		path: "/repos/test/test/notifications", params: `{"owner": "test", "repo": "test"}`},
	{method: "PUT", pattern: "/repos/{str:owner}/{str:repo}/notifications",
		path: "/repos/test/test/notifications", params: `{"owner": "test", "repo": "test"}`},
	{method: "GET", pattern: "/notifications/threads/{int:id}",
		path: "/notifications/threads/123", params: `{"id": "123"}`},
	{method: "PATCH", pattern: "/notifications/threads/{int:id}",
		path: "/notifications/threads/123", params: `{"id": "123"}`},
	{method: "GET", pattern: "/notifications/threads/{int:id}/subscription",
		path: "/notifications/threads/123/subscription", params: `{"id": "123"}`},
	{method: "PUT", pattern: "/notifications/threads/{int:id}/subscription",
		path: "/notifications/threads/123/subscription", params: `{"id": "123"}`},
	{method: "DELETE", pattern: "/notifications/threads/{int:id}/subscription",
		path: "/notifications/threads/123/subscription", params: `{"id": "123"}`},
	{method: "GET", pattern: "/repos/{str:owner}/{str:repo}/stargazers",
		path: "/repos/test/test/stargazers", params: `{"owner": "test", "repo": "test"}`},
	{method: "GET", pattern: "/users/{str:user}/starred",
		path: "/users/test/starred", params: `{"user": "test"}`},
	{method: "GET", pattern: "/user/starred/{str:owner}/{str:repo}",
		path: "/user/starred/test/test", params: `{"owner": "test", "repo": "test"}`},
	{method: "PUT", pattern: "/user/starred/{str:owner}/{str:repo}",
		path: "/user/starred/test/test", params: `{"owner": "test", "repo": "test"}`},
	{method: "DELETE", pattern: "/user/starred/{str:owner}/{str:repo}",
		path: "/user/starred/test/test", params: `{"owner": "test", "repo": "test"}`},
	{method: "GET", pattern: "/repos/{str:owner}/{str:repo}/subscribers",
		path: "/repos/test/test/subscribers", params: `{"owner": "test", "repo": "test"}`},
	{method: "GET", pattern: "/users/{str:user}/subscriptions",
		path: "/users/test/subscriptions", params: `{"user": "test"}`},
	{method: "GET", pattern: "/repos/{str:owner}/{str:repo}/subscription",
		path: "/repos/test/test/subscription", params: `{"owner": "test", "repo": "test"}`},
	{method: "PUT", pattern: "/repos/{str:owner}/{str:repo}/subscription",
		path: "/repos/test/test/subscription", params: `{"owner": "test", "repo": "test"}`},
	{method: "DELETE", pattern: "/repos/{str:owner}/{str:repo}/subscription",
		path: "/repos/test/test/subscription", params: `{"owner": "test", "repo": "test"}`},
	{method: "GET", pattern: "/user/subscriptions/{str:owner}/{str:repo}",
		path: "/user/subscriptions/test/test", params: `{"owner": "test", "repo": "test"}`},
	{method: "PUT", pattern: "/user/subscriptions/{str:owner}/{str:repo}",
		path: "/user/subscriptions/test/test", params: `{"owner": "test", "repo": "test"}`},
	{method: "DELETE", pattern: "/user/subscriptions/{str:owner}/{str:repo}",
		path: "/user/subscriptions/test/test", params: `{"owner": "test", "repo": "test"}`},
}

// Tests that all 'activity' API routes resolves correctly
func TestGitHubResolveActivity(t *testing.T) {
	testResolve(t, githubActivity)
}

// Bench for 'activity' API resolving
func BenchmarkGithubResolveActivity(b *testing.B) {
	benchResolve(b, githubActivity)
}

// Test for 'activity' API params parsing
func TestGitHubParamsActivity(t *testing.T) {
	testParams(t, githubActivity)
}
