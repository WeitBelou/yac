package yac

import "testing"

var githubRepositories = []route{
	{method: "GET", pattern: "/user/repos",
		path: "/user/repos"},
	{method: "GET", pattern: "/repositories",
		path: "/repositories"},
	{method: "POST", pattern: "/user/repos",
		path: "/user/repos"},
	{method: "GET", pattern: "/users/{str:user}/repos",
		path: "/users/test/repos", params: `{"user":"test"}`},
	{method: "GET", pattern: "/orgs/{str:org}/repos",
		path: "/orgs/test/repos", params: `{"org":"test"}`},
	{method: "POST", pattern: "/orgs/{str:org}/repos",
		path: "/orgs/test/repos", params: `{"org":"test"}`},
	{method: "GET", pattern: "/repos/{str:owner}/{str:repo}",
		path: "/repos/test/test", params: `{"owner":"test","repo":"test"}`},
	{method: "PATCH", pattern: "/repos/{str:owner}/{str:repo}",
		path: "/repos/test/test", params: `{"owner":"test","repo":"test"}`},
	{method: "GET", pattern: "/repos/{str:owner}/{str:repo}/contributors",
		path: "/repos/test/test/contributors", params: `{"owner":"test","repo":"test"}`},
	{method: "GET", pattern: "/repos/{str:owner}/{str:repo}/languages",
		path: "/repos/test/test/languages", params: `{"owner":"test","repo":"test"}`},
	{method: "GET", pattern: "/repos/{str:owner}/{str:repo}/teams",
		path: "/repos/test/test/teams", params: `{"owner":"test","repo":"test"}`},
	{method: "GET", pattern: "/repos/{str:owner}/{str:repo}/tags",
		path: "/repos/test/test/tags", params: `{"owner":"test","repo":"test"}`},
	{method: "GET", pattern: "/repos/{str:owner}/{str:repo}/branches",
		path: "/repos/test/test/branches", params: `{"owner":"test","repo":"test"}`},
	{method: "GET", pattern: "/repos/{str:owner}/{str:repo}/branches/{str:branch}",
		path: "/repos/test/test/branches/test", params: `{"owner":"test","repo":"test","branch":"test"}`},
	{method: "DELETE", pattern: "/repos/{str:owner}/{str:repo}",
		path: "/repos/test/test", params: `{"owner":"test","repo":"test"}`},
	{method: "GET", pattern: "/repos/{str:owner}/{str:repo}/collaborators",
		path: "/repos/test/test/collaborators", params: `{"owner":"test","repo":"test"}`},
	{method: "GET", pattern: "/repos/{str:owner}/{str:repo}/collaborators/{str:user}",
		path: "/repos/test/test/collaborators/test", params: `{"owner":"test","repo":"test","user":"test"}`},
	{method: "PUT", pattern: "/repos/{str:owner}/{str:repo}/collaborators/{str:user}",
		path: "/repos/test/test/collaborators/test", params: `{"owner":"test","repo":"test","user":"test"}`},
	{method: "DELETE", pattern: "/repos/{str:owner}/{str:repo}/collaborators/{str:user}",
		path: "/repos/test/test/collaborators/test", params: `{"owner":"test","repo":"test","user":"test"}`},
	{method: "GET", pattern: "/repos/{str:owner}/{str:repo}/comments",
		path: "/repos/test/test/comments", params: `{"owner":"test","repo":"test"}`},
	{method: "GET", pattern: "/repos/{str:owner}/{str:repo}/commits/{hex:sha}/comments",
		path:   "/repos/test/test/commits/dead12beef/comments",
		params: `{"owner":"test","repo":"test","sha":"dead12beef"}`},
	{method: "POST", pattern: "/repos/{str:owner}/{str:repo}/commits/{hex:sha}/comments",
		path:   "/repos/test/test/commits/dead12beef/comments",
		params: `{"owner":"test","repo":"test","sha":"dead12beef"}`},
	{method: "GET", pattern: "/repos/{str:owner}/{str:repo}/comments/{int:id}",
		path: "/repos/test/test/comments/123", params: `{"owner":"test","repo":"test","id":"123"}`},
	{method: "PATCH", pattern: "/repos/{str:owner}/{str:repo}/comments/{int:id}",
		path: "/repos/test/test/comments/123", params: `{"owner":"test","repo":"test","id":"123"}`},
	{method: "DELETE", pattern: "/repos/{str:owner}/{str:repo}/comments/{int:id}",
		path: "/repos/test/test/comments/123", params: `{"owner":"test","repo":"test","id":"123"}`},
	{method: "GET", pattern: "/repos/{str:owner}/{str:repo}/commits",
		path: "/repos/test/test/commits", params: `{"owner":"test","repo":"test"}`},
	{method: "GET", pattern: "/repos/{str:owner}/{str:repo}/commits/{hex:sha}",
		path:   "/repos/test/test/commits/dead12beef",
		params: `{"owner":"test","repo":"test","sha":"dead12beef"}`},
	{method: "GET", pattern: "/repos/{str:owner}/{str:repo}/readme",
		path: "/repos/test/test/readme", params: `{"owner":"test","repo":"test"}`},
	{method: "GET", pattern: "/repos/{str:owner}/{str:repo}/contents/{str:path}",
		path: "/repos/test/test/contents/pasdasd", params: `{"owner":"test","repo":"test","path":"pasdasd"}`},
	{method: "PUT", pattern: "/repos/{str:owner}/{str:repo}/contents/{str:path}",
		path: "/repos/test/test/contents/pathdas", params: `{"owner":"test","repo":"test","path":"pathdas"}`},
	{method: "DELETE", pattern: "/repos/{str:owner}/{str:repo}/contents/{str:path}",
		path: "/repos/test/test/contents/pathsda", params: `{"owner":"test","repo":"test","path":"pathsda"}`},
	{method: "GET", pattern: "/repos/{str:owner}/{str:repo}/{str:archive_format}/{str:ref}",
		path:   "/repos/test/test/test/test",
		params: `{"owner":"test","repo":"test","archive_format":"test","ref":"test"}`},
	{method: "GET", pattern: "/repos/{str:owner}/{str:repo}/keys",
		path: "/repos/test/test/keys", params: `{"owner":"test","repo":"test"}`},
	{method: "GET", pattern: "/repos/{str:owner}/{str:repo}/keys/{int:id}",
		path: "/repos/test/test/keys/123", params: `{"owner":"test","repo":"test","id":"123"}`},
	{method: "POST", pattern: "/repos/{str:owner}/{str:repo}/keys",
		path: "/repos/test/test/keys", params: `{"owner":"test","repo":"test"}`},
	{method: "PATCH", pattern: "/repos/{str:owner}/{str:repo}/keys/{int:id}",
		path: "/repos/test/test/keys/123", params: `{"owner":"test","repo":"test","id":"123"}`},
	{method: "DELETE", pattern: "/repos/{str:owner}/{str:repo}/keys/{int:id}",
		path: "/repos/test/test/keys/123", params: `{"owner":"test","repo":"test","id":"123"}`},
	{method: "GET", pattern: "/repos/{str:owner}/{str:repo}/downloads",
		path: "/repos/test/test/downloads", params: `{"owner":"test","repo":"test"}`},
	{method: "GET", pattern: "/repos/{str:owner}/{str:repo}/downloads/{int:id}",
		path: "/repos/test/test/downloads/123", params: `{"owner":"test","repo":"test","id":"123"}`},
	{method: "DELETE", pattern: "/repos/{str:owner}/{str:repo}/downloads/{int:id}",
		path: "/repos/test/test/downloads/123", params: `{"owner":"test","repo":"test","id":"123"}`},
	{method: "GET", pattern: "/repos/{str:owner}/{str:repo}/forks",
		path: "/repos/test/test/forks", params: `{"owner":"test","repo":"test"}`},
	{method: "POST", pattern: "/repos/{str:owner}/{str:repo}/forks",
		path: "/repos/test/test/forks", params: `{"owner":"test","repo":"test"}`},
	{method: "GET", pattern: "/repos/{str:owner}/{str:repo}/hooks",
		path: "/repos/test/test/hooks", params: `{"owner":"test","repo":"test"}`},
	{method: "GET", pattern: "/repos/{str:owner}/{str:repo}/hooks/{int:id}",
		path: "/repos/test/test/hooks/123", params: `{"owner":"test","repo":"test","id":"123"}`},
	{method: "POST", pattern: "/repos/{str:owner}/{str:repo}/hooks",
		path: "/repos/test/test/hooks", params: `{"owner":"test","repo":"test"}`},
	{method: "PATCH", pattern: "/repos/{str:owner}/{str:repo}/hooks/{int:id}",
		path: "/repos/test/test/hooks/123", params: `{"owner":"test","repo":"test","id":"123"}`},
	{method: "POST", pattern: "/repos/{str:owner}/{str:repo}/hooks/{int:id}/tests",
		path: "/repos/test/test/hooks/123/tests", params: `{"owner":"test","repo":"test","id":"123"}`},
	{method: "DELETE", pattern: "/repos/{str:owner}/{str:repo}/hooks/{int:id}",
		path: "/repos/test/test/hooks/123", params: `{"owner":"test","repo":"test","id":"123"}`},
	{method: "POST", pattern: "/repos/{str:owner}/{str:repo}/merges",
		path: "/repos/test/test/merges", params: `{"owner":"test","repo":"test"}`},
	{method: "GET", pattern: "/repos/{str:owner}/{str:repo}/releases",
		path: "/repos/test/test/releases", params: `{"owner":"test","repo":"test"}`},
	{method: "GET", pattern: "/repos/{str:owner}/{str:repo}/releases/{int:id}",
		path: "/repos/test/test/releases/123", params: `{"owner":"test","repo":"test","id":"123"}`},
	{method: "POST", pattern: "/repos/{str:owner}/{str:repo}/releases",
		path: "/repos/test/test/releases", params: `{"owner":"test","repo":"test"}`},
	{method: "PATCH", pattern: "/repos/{str:owner}/{str:repo}/releases/{int:id}",
		path: "/repos/test/test/releases/123", params: `{"owner":"test","repo":"test","id":"123"}`},
	{method: "DELETE", pattern: "/repos/{str:owner}/{str:repo}/releases/{int:id}",
		path: "/repos/test/test/releases/123", params: `{"owner":"test","repo":"test","id":"123"}`},
	{method: "GET", pattern: "/repos/{str:owner}/{str:repo}/releases/{int:id}/assets",
		path: "/repos/test/test/releases/123/assets", params: `{"owner":"test","repo":"test","id":"123"}`},
	{method: "GET", pattern: "/repos/{str:owner}/{str:repo}/stats/contributors",
		path: "/repos/test/test/stats/contributors", params: `{"owner":"test","repo":"test"}`},
	{method: "GET", pattern: "/repos/{str:owner}/{str:repo}/stats/commit_activity",
		path: "/repos/test/test/stats/commit_activity", params: `{"owner":"test","repo":"test"}`},
	{method: "GET", pattern: "/repos/{str:owner}/{str:repo}/stats/code_frequency",
		path: "/repos/test/test/stats/code_frequency", params: `{"owner":"test","repo":"test"}`},
	{method: "GET", pattern: "/repos/{str:owner}/{str:repo}/stats/participation",
		path: "/repos/test/test/stats/participation", params: `{"owner":"test","repo":"test"}`},
	{method: "GET", pattern: "/repos/{str:owner}/{str:repo}/stats/punch_card",
		path: "/repos/test/test/stats/punch_card", params: `{"owner":"test","repo":"test"}`},
	{method: "GET", pattern: "/repos/{str:owner}/{str:repo}/statuses/{str:ref}",
		path: "/repos/test/test/statuses/test", params: `{"owner":"test","repo":"test","ref":"test"}`},
	{method: "POST", pattern: "/repos/{str:owner}/{str:repo}/statuses/{str:ref}",
		path: "/repos/test/test/statuses/test", params: `{"owner":"test","repo":"test","ref":"test"}`},
}

// Tests that all 'repositories' API routes resolves correctly
func TestGitHubResolveRepositories(t *testing.T) {
	testResolve(t, githubRepositories)
}

// Bench for 'repositories' API resolving
func BenchmarkGithubResolveRepositories(b *testing.B) {
	benchResolve(b, githubRepositories)
}

// Test for 'repositories' API params parsing
func TestGitHubParamsRepositories(t *testing.T) {
	testParams(t, githubRepositories)
}
