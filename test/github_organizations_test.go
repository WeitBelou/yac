package yac

import "testing"

var githubOrganizations = []route{
	{method: "GET", pattern: "/user/orgs",
		path: "/user/orgs"},
	{method: "GET", pattern: "/user/teams",
		path: "/user/teams"},
	{method: "GET", pattern: "/users/{str:user}/orgs",
		path: "/users/test/orgs", params: `{"user":"test"}`},
	{method: "GET", pattern: "/orgs/{str:org}",
		path: "/orgs/test", params: `{"org":"test"}`},
	{method: "PATCH", pattern: "/orgs/{str:org}",
		path: "/orgs/test", params: `{"org":"test"}`},
	{method: "GET", pattern: "/orgs/{str:org}/members",
		path: "/orgs/test/members", params: `{"org":"test"}`},
	{method: "GET", pattern: "/orgs/{str:org}/members/{str:user}",
		path: "/orgs/test/members/test", params: `{"org":"test","user":"test"}`},
	{method: "DELETE", pattern: "/orgs/{str:org}/members/{str:user}",
		path: "/orgs/test/members/test", params: `{"org":"test","user":"test"}`},
	{method: "GET", pattern: "/orgs/{str:org}/public_members",
		path: "/orgs/test/public_members", params: `{"org":"test"}`},
	{method: "GET", pattern: "/orgs/{str:org}/public_members/{str:user}",
		path: "/orgs/test/public_members/test", params: `{"org":"test","user":"test"}`},
	{method: "PUT", pattern: "/orgs/{str:org}/public_members/{str:user}",
		path: "/orgs/test/public_members/test", params: `{"org":"test","user":"test"}`},
	{method: "DELETE", pattern: "/orgs/{str:org}/public_members/{str:user}",
		path: "/orgs/test/public_members/test", params: `{"org":"test","user":"test"}`},
	{method: "GET", pattern: "/orgs/{str:org}/teams",
		path: "/orgs/test/teams", params: `{"org":"test"}`},
	{method: "GET", pattern: "/teams/{int:id}",
		path: "/teams/123", params: `{"id":"123"}`},
	{method: "POST", pattern: "/orgs/{str:org}/teams",
		path: "/orgs/test/teams", params: `{"org":"test"}`},
	{method: "PATCH", pattern: "/teams/{int:id}",
		path: "/teams/123", params: `{"id":"123"}`},
	{method: "DELETE", pattern: "/teams/{int:id}",
		path: "/teams/123", params: `{"id":"123"}`},
	{method: "GET", pattern: "/teams/{int:id}/members",
		path: "/teams/123/members", params: `{"id":"123"}`},
	{method: "GET", pattern: "/teams/{int:id}/members/{str:user}",
		path: "/teams/123/members/test", params: `{"id":"123","user":"test"}`},
	{method: "PUT", pattern: "/teams/{int:id}/members/{str:user}",
		path: "/teams/123/members/test", params: `{"id":"123","user":"test"}`},
	{method: "DELETE", pattern: "/teams/{int:id}/members/{str:user}",
		path: "/teams/123/members/test", params: `{"id":"123","user":"test"}`},
	{method: "GET", pattern: "/teams/{int:id}/repos",
		path: "/teams/123/repos", params: `{"id":"123"}`},
	{method: "GET", pattern: "/teams/{int:id}/repos/{str:owner}/{str:repo}",
		path: "/teams/123/repos/test/test", params: `{"id":"123","owner":"test","repo":"test"}`},
	{method: "PUT", pattern: "/teams/{int:id}/repos/{str:owner}/{str:repo}",
		path: "/teams/123/repos/test/test", params: `{"id":"123","owner":"test","repo":"test"}`},
	{method: "DELETE", pattern: "/teams/{int:id}/repos/{str:owner}/{str:repo}",
		path: "/teams/123/repos/test/test", params: `{"id":"123","owner":"test","repo":"test"}`},
}

// Tests that all 'organizations' API routes resolves correctly
func TestGitHubResolveOrganizations(t *testing.T) {
	testResolve(t, githubOrganizations)
}

// Bench for 'organizations' API resolving
func BenchmarkGithubResolveOrganizations(b *testing.B) {
	benchResolve(b, githubOrganizations)
}

// Test for 'organizations' API params parsing
func TestGitHubParamsOrganizations(t *testing.T) {
	testParams(t, githubOrganizations)
}
