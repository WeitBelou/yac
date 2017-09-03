package yac

import "testing"

var githubOauth = []route{
	{method: "GET", pattern: "/authorizations",
		path: "/authorizations", params: `{}`},
	{method: "POST", pattern: "/authorizations",
		path: "/authorizations", params: `{}`},
	{method: "GET", pattern: "/authorizations/{int:id}",
		path: "/authorizations/123", params: `{"id": "123"}`},
	{method: "PATCH", pattern: "/authorizations/{int:id}",
		path: "/authorizations/123", params: `{"id": "123"}`},
	{method: "DELETE", pattern: "/authorizations/{int:id}",
		path: "/authorizations/123", params: `{"id": "123"}`},
	{method: "PUT", pattern: "/authorizations/clients/{int:client_id}",
		path: "/authorizations/clients/123", params: `{"client_id": "123"}`},
	{method: "DELETE", pattern: "/applications/{int:client_id}/tokens",
		path: "/applications/123/tokens", params: `{"client_id": "123"}`},
	{method: "GET", pattern: "/applications/{int:client_id}/tokens/{str:access_token}",
		path: "/applications/123/tokens/test", params: `{"client_id": "123", "access_token": "test"}`},
	{method: "DELETE", pattern: "/applications/{int:client_id}/tokens/{str:access_token}",
		path: "/applications/123/tokens/test", params: `{"client_id": "123", "access_token": "test"}`},
}

// Tests that all 'oauth' API routes resolves correctly
func TestGitHubResolveOauth(t *testing.T) {
	testResolve(t, githubOauth)
}

// Bench for 'oauth' API resolving
func BenchmarkGithubResolveOauth(b *testing.B) {
	benchResolve(b, githubOauth)
}

// Test for 'oauth' params parsing
func TestGitHubParamsOauth(t *testing.T) {
	testParams(t, githubOauth)
}
