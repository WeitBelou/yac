package yac

import (
	"net/http"
	"testing"
)

var githubStatic = []routeTestCase{
	// OAuth Authorizations
	{method: "GET", pattern: "/authorizations"},
	{method: "POST", pattern: "/authorizations"},

	// Activity
	{method: "GET", pattern: "/events"},
	{method: "GET", pattern: "/feeds"},
	{method: "GET", pattern: "/notifications"},
	{method: "PUT", pattern: "/notifications"},
	{method: "GET", pattern: "/user/starred"},
	{method: "GET", pattern: "/user/subscriptions"},

	// Gists
	{method: "GET", pattern: "/gists"},
	{method: "GET", pattern: "/gists/public"},
	{method: "GET", pattern: "/gists/starred"},
	{method: "POST", pattern: "/gists"},

	// Issues
	{method: "GET", pattern: "/issues"},
	{method: "GET", pattern: "/user/issues"},

	// Miscellaneous
	{method: "GET", pattern: "/emojis"},
	{method: "GET", pattern: "/gitignore/templates"},
	{method: "POST", pattern: "/markdown"},
	{method: "POST", pattern: "/markdown/raw"},
	{method: "GET", pattern: "/meta"},
	{method: "GET", pattern: "/rate_limit"},

	// Organizations
	{method: "GET", pattern: "/user/orgs"},
	{method: "GET", pattern: "/user/teams"},

	// Repositories
	{method: "GET", pattern: "/user/repos"},
	{method: "GET", pattern: "/repositories"},
	{method: "POST", pattern: "/user/repos"},

	// Search
	{method: "GET", pattern: "/search/repositories"},
	{method: "GET", pattern: "/search/code"},
	{method: "GET", pattern: "/search/issues"},
	{method: "GET", pattern: "/search/users"},

	// Users
	{method: "GET", pattern: "/user"},
	{method: "PATCH", pattern: "/user"},
	{method: "GET", pattern: "/users"},
	{method: "GET", pattern: "/user/emails"},
	{method: "POST", pattern: "/user/emails"},
	{method: "DELETE", pattern: "/user/emails"},
	{method: "GET", pattern: "/user/followers"},
	{method: "GET", pattern: "/user/following"},
	{method: "GET", pattern: "/user/keys"},
	{method: "POST", pattern: "/user/keys"},
}

var githubDynamic = []routeTestCase{
	// OAuth Authorizations
	{method: "GET", pattern: "/authorizations/{int:id}"},
	{method: "PUT", pattern: "/authorizations/clients/{int:client_id}"},
	{method: "PATCH", pattern: "/authorizations/{int:id}"},
	{method: "DELETE", pattern: "/authorizations/{int:id}"},
	{method: "GET", pattern: "/applications/{int:client_id}/tokens/{str:access_token}"},
	{method: "DELETE", pattern: "/applications/{int:client_id}/tokens"},
	{method: "DELETE", pattern: "/applications/{int:client_id}/tokens/{str:access_token}"},

	// Activity
	{method: "GET", pattern: "/repos/{str:owner}/{str:repo}/events"},
	{method: "GET", pattern: "/networks/{str:owner}/{str:repo}/events"},
	{method: "GET", pattern: "/orgs/{str:org}/events"},
	{method: "GET", pattern: "/users/{str:user}/received_events"},
	{method: "GET", pattern: "/users/{str:user}/received_events/public"},
	{method: "GET", pattern: "/users/{str:user}/events"},
	{method: "GET", pattern: "/users/{str:user}/events/public"},
	{method: "GET", pattern: "/users/{str:user}/events/orgs/{str:org}"},
	{method: "GET", pattern: "/repos/{str:owner}/{str:repo}/notifications"},
	{method: "PUT", pattern: "/repos/{str:owner}/{str:repo}/notifications"},
	{method: "GET", pattern: "/notifications/threads/{int:id}"},
	{method: "PATCH", pattern: "/notifications/threads/{int:id}"},
	{method: "GET", pattern: "/notifications/threads/{int:id}/subscription"},
	{method: "PUT", pattern: "/notifications/threads/{int:id}/subscription"},
	{method: "DELETE", pattern: "/notifications/threads/{int:id}/subscription"},
	{method: "GET", pattern: "/repos/{str:owner}/{str:repo}/stargazers"},
	{method: "GET", pattern: "/users/{str:user}/starred"},
	{method: "GET", pattern: "/user/starred/{str:owner}/{str:repo}"},
	{method: "PUT", pattern: "/user/starred/{str:owner}/{str:repo}"},
	{method: "DELETE", pattern: "/user/starred/{str:owner}/{str:repo}"},
	{method: "GET", pattern: "/repos/{str:owner}/{str:repo}/subscribers"},
	{method: "GET", pattern: "/users/{str:user}/subscriptions"},
	{method: "GET", pattern: "/repos/{str:owner}/{str:repo}/subscription"},
	{method: "PUT", pattern: "/repos/{str:owner}/{str:repo}/subscription"},
	{method: "DELETE", pattern: "/repos/{str:owner}/{str:repo}/subscription"},
	{method: "GET", pattern: "/user/subscriptions/{str:owner}/{str:repo}"},
	{method: "PUT", pattern: "/user/subscriptions/{str:owner}/{str:repo}"},
	{method: "DELETE", pattern: "/user/subscriptions/{str:owner}/{str:repo}"},

	// Gists
	{method: "GET", pattern: "/users/{str:user}/gists"},
	{method: "GET", pattern: "/gists/{int:id}"},
	{method: "PATCH", pattern: "/gists/{int:id}"},
	{method: "PUT", pattern: "/gists/{int:id}/star"},
	{method: "DELETE", pattern: "/gists/{int:id}/star"},
	{method: "GET", pattern: "/gists/{int:id}/star"},
	{method: "POST", pattern: "/gists/{int:id}/forks"},
	{method: "DELETE", pattern: "/gists/{int:id}"},

	// Git Data
	{method: "GET", pattern: "/repos/{str:owner}/{str:repo}/git/blobs/{hex:sha}"},
	{method: "POST", pattern: "/repos/{str:owner}/{str:owner}/git/blobs"},
	{method: "GET", pattern: "/repos/{str:owner}/{str:owner}/git/commits/{hex:sha}"},
	{method: "POST", pattern: "/repos/{str:owner}/{str:owner}/git/commits"},
	{method: "GET", pattern: "/repos/{str:owner}/{str:owner}/git/refs/{str:ref}"},
	{method: "GET", pattern: "/repos/{str:owner}/{str:owner}/git/refs"},
	{method: "POST", pattern: "/repos/{str:owner}/{str:owner}/git/refs"},
	{method: "PATCH", pattern: "/repos/{str:owner}/{str:owner}/git/refs/{str:ref}"},
	{method: "DELETE", pattern: "/repos/{str:owner}/{str:owner}/git/refs/{str:ref}"},
	{method: "GET", pattern: "/repos/{str:owner}/{str:owner}/git/tags/{hex:sha}"},
	{method: "POST", pattern: "/repos/{str:owner}/{str:owner}/git/tags"},
	{method: "GET", pattern: "/repos/{str:owner}/{str:owner}/git/trees/{hex:sha}"},
	{method: "POST", pattern: "/repos/{str:owner}/{str:owner}/git/trees"},

	// Issues
	{method: "GET", pattern: "/orgs/{str:org}/issues"},
	{method: "GET", pattern: "/repos/{str:owner}/{str:owner}/issues"},
	{method: "GET", pattern: "/repos/{str:owner}/{str:owner}/issues/{int:user}"},
	{method: "POST", pattern: "/repos/{str:owner}/{str:owner}/issues"},
	{method: "PATCH", pattern: "/repos/{str:owner}/{str:owner}/issues/{int:user}"},
	{method: "GET", pattern: "/repos/{str:owner}/{str:owner}/assignees"},
	{method: "GET", pattern: "/repos/{str:owner}/{str:owner}/assignees/{str:assignee}"},
	{method: "GET", pattern: "/repos/{str:owner}/{str:owner}/issues/{int:user}/comments"},
	{method: "GET", pattern: "/repos/{str:owner}/{str:owner}/issues/comments"},
	{method: "GET", pattern: "/repos/{str:owner}/{str:owner}/issues/comments/{int:id}"},
	{method: "POST", pattern: "/repos/{str:owner}/{str:owner}/issues/{int:user}/comments"},
	{method: "PATCH", pattern: "/repos/{str:owner}/{str:owner}/issues/comments/{int:id}"},
	{method: "DELETE", pattern: "/repos/{str:owner}/{str:owner}/issues/comments/{int:id}"},
	{method: "GET", pattern: "/repos/{str:owner}/{str:owner}/issues/{int:user}/events"},
	{method: "GET", pattern: "/repos/{str:owner}/{str:owner}/issues/events"},
	{method: "GET", pattern: "/repos/{str:owner}/{str:owner}/issues/events/{int:id}"},
	{method: "GET", pattern: "/repos/{str:owner}/{str:owner}/labels"},
	{method: "GET", pattern: "/repos/{str:owner}/{str:owner}/labels/{str:name}"},
	{method: "POST", pattern: "/repos/{str:owner}/{str:owner}/labels"},
	{method: "PATCH", pattern: "/repos/{str:owner}/{str:owner}/labels/{str:name}"},
	{method: "DELETE", pattern: "/repos/{str:owner}/{str:owner}/labels/{str:name}"},
	{method: "GET", pattern: "/repos/{str:owner}/{str:owner}/issues/{int:user}/labels"},
	{method: "POST", pattern: "/repos/{str:owner}/{str:owner}/issues/{int:user}/labels"},
	{method: "DELETE", pattern: "/repos/{str:owner}/{str:owner}/issues/{int:user}/labels/{str:name}"},
	{method: "PUT", pattern: "/repos/{str:owner}/{str:owner}/issues/{int:user}/labels"},
	{method: "DELETE", pattern: "/repos/{str:owner}/{str:owner}/issues/{int:user}/labels"},
	{method: "GET", pattern: "/repos/{str:owner}/{str:owner}/milestones/{int:user}/labels"},
	{method: "GET", pattern: "/repos/{str:owner}/{str:owner}/milestones"},
	{method: "GET", pattern: "/repos/{str:owner}/{str:owner}/milestones/{int:user}"},
	{method: "POST", pattern: "/repos/{str:owner}/{str:owner}/milestones"},
	{method: "PATCH", pattern: "/repos/{str:owner}/{str:owner}/milestones/{int:user}"},
	{method: "DELETE", pattern: "/repos/{str:owner}/{str:owner}/milestones/{int:user}"},

	// Miscellaneous
	{method: "GET", pattern: "/gitignore/templates/{str:name}"},

	// Organizations
	{method: "GET", pattern: "/users/{str:user}/orgs"},
	{method: "GET", pattern: "/orgs/{str:org}"},
	{method: "PATCH", pattern: "/orgs/{str:org}"},
	{method: "GET", pattern: "/orgs/{str:org}/members"},
	{method: "GET", pattern: "/orgs/{str:org}/members/{str:user}"},
	{method: "DELETE", pattern: "/orgs/{str:org}/members/{str:user}"},
	{method: "GET", pattern: "/orgs/{str:org}/public_members"},
	{method: "GET", pattern: "/orgs/{str:org}/public_members/{str:user}"},
	{method: "PUT", pattern: "/orgs/{str:org}/public_members/{str:user}"},
	{method: "DELETE", pattern: "/orgs/{str:org}/public_members/{str:user}"},
	{method: "GET", pattern: "/orgs/{str:org}/teams"},
	{method: "GET", pattern: "/teams/{int:id}"},
	{method: "POST", pattern: "/orgs/{str:org}/teams"},
	{method: "PATCH", pattern: "/teams/{int:id}"},
	{method: "DELETE", pattern: "/teams/{int:id}"},
	{method: "GET", pattern: "/teams/{int:id}/members"},
	{method: "GET", pattern: "/teams/{int:id}/members/{str:user}"},
	{method: "PUT", pattern: "/teams/{int:id}/members/{str:user}"},
	{method: "DELETE", pattern: "/teams/{int:id}/members/{str:user}"},
	{method: "GET", pattern: "/teams/{int:id}/repos"},
	{method: "GET", pattern: "/teams/{int:id}/repos/{str:owner}/{str:repo}"},
	{method: "PUT", pattern: "/teams/{int:id}/repos/{str:owner}/{str:repo}"},
	{method: "DELETE", pattern: "/teams/{int:id}/repos/{str:owner}/{str:repo}"},

	// Pull Requests
	{method: "GET", pattern: "/repos/{str:owner}/{str:owner}/pulls"},
	{method: "GET", pattern: "/repos/{str:owner}/{str:owner}/pulls/{int:user}"},
	{method: "POST", pattern: "/repos/{str:owner}/{str:owner}/pulls"},
	{method: "PATCH", pattern: "/repos/{str:owner}/{str:owner}/pulls/{int:user}"},
	{method: "GET", pattern: "/repos/{str:owner}/{str:owner}/pulls/{int:user}/commits"},
	{method: "GET", pattern: "/repos/{str:owner}/{str:owner}/pulls/{int:user}/files"},
	{method: "GET", pattern: "/repos/{str:owner}/{str:owner}/pulls/{int:user}/merge"},
	{method: "PUT", pattern: "/repos/{str:owner}/{str:owner}/pulls/{int:user}/merge"},
	{method: "GET", pattern: "/repos/{str:owner}/{str:owner}/pulls/{int:user}/comments"},
	{method: "GET", pattern: "/repos/{str:owner}/{str:owner}/pulls/comments"},
	{method: "GET", pattern: "/repos/{str:owner}/{str:owner}/pulls/comments/{int:user}"},
	{method: "PUT", pattern: "/repos/{str:owner}/{str:owner}/pulls/{int:user}/comments"},
	{method: "PATCH", pattern: "/repos/{str:owner}/{str:owner}/pulls/comments/{int:user}"},
	{method: "DELETE", pattern: "/repos/{str:owner}/{str:owner}/pulls/comments/{int:user}"},

	// Repositories
	{method: "GET", pattern: "/users/{str:user}/repos"},
	{method: "GET", pattern: "/orgs/{str:org}/repos"},
	{method: "POST", pattern: "/orgs/{str:org}/repos"},
	{method: "GET", pattern: "/repos/{str:owner}/{str:repo}"},
	{method: "PATCH", pattern: "/repos/{str:owner}/{str:repo}"},
	{method: "GET", pattern: "/repos/{str:owner}/{str:owner}/contributors"},
	{method: "GET", pattern: "/repos/{str:owner}/{str:owner}/languages"},
	{method: "GET", pattern: "/repos/{str:owner}/{str:owner}/teams"},
	{method: "GET", pattern: "/repos/{str:owner}/{str:owner}/tags"},
	{method: "GET", pattern: "/repos/{str:owner}/{str:owner}/branches"},
	{method: "GET", pattern: "/repos/{str:owner}/{str:owner}/branches/{str:branch}"},
	{method: "DELETE", pattern: "/repos/{str:owner}/{str:repo}"},
	{method: "GET", pattern: "/repos/{str:owner}/{str:owner}/collaborators"},
	{method: "GET", pattern: "/repos/{str:owner}/{str:owner}/collaborators/{str:user}"},
	{method: "PUT", pattern: "/repos/{str:owner}/{str:owner}/collaborators/{str:user}"},
	{method: "DELETE", pattern: "/repos/{str:owner}/{str:owner}/collaborators/{str:user}"},
	{method: "GET", pattern: "/repos/{str:owner}/{str:owner}/comments"},
	{method: "GET", pattern: "/repos/{str:owner}/{str:owner}/commits/{hex:sha}/comments"},
	{method: "POST", pattern: "/repos/{str:owner}/{str:owner}/commits/{hex:sha}/comments"},
	{method: "GET", pattern: "/repos/{str:owner}/{str:owner}/comments/{int:id}"},
	{method: "PATCH", pattern: "/repos/{str:owner}/{str:owner}/comments/{int:id}"},
	{method: "DELETE", pattern: "/repos/{str:owner}/{str:owner}/comments/{int:id}"},
	{method: "GET", pattern: "/repos/{str:owner}/{str:owner}/commits"},
	{method: "GET", pattern: "/repos/{str:owner}/{str:owner}/commits/{hex:sha}"},
	{method: "GET", pattern: "/repos/{str:owner}/{str:owner}/readme"},
	{method: "GET", pattern: "/repos/{str:owner}/{str:owner}/contents/*path"},
	{method: "PUT", pattern: "/repos/{str:owner}/{str:owner}/contents/*path"},
	{method: "DELETE", pattern: "/repos/{str:owner}/{str:owner}/contents/*path"},
	{method: "GET", pattern: "/repos/{str:owner}/{str:owner}/{str:archive_format}/{str:ref}"},
	{method: "GET", pattern: "/repos/{str:owner}/{str:owner}/keys"},
	{method: "GET", pattern: "/repos/{str:owner}/{str:owner}/keys/{int:id}"},
	{method: "POST", pattern: "/repos/{str:owner}/{str:owner}/keys"},
	{method: "PATCH", pattern: "/repos/{str:owner}/{str:owner}/keys/{int:id}"},
	{method: "DELETE", pattern: "/repos/{str:owner}/{str:owner}/keys/{int:id}"},
	{method: "GET", pattern: "/repos/{str:owner}/{str:owner}/downloads"},
	{method: "GET", pattern: "/repos/{str:owner}/{str:owner}/downloads/{int:id}"},
	{method: "DELETE", pattern: "/repos/{str:owner}/{str:owner}/downloads/{int:id}"},
	{method: "GET", pattern: "/repos/{str:owner}/{str:owner}/forks"},
	{method: "POST", pattern: "/repos/{str:owner}/{str:owner}/forks"},
	{method: "GET", pattern: "/repos/{str:owner}/{str:owner}/hooks"},
	{method: "GET", pattern: "/repos/{str:owner}/{str:owner}/hooks/{int:id}"},
	{method: "POST", pattern: "/repos/{str:owner}/{str:owner}/hooks"},
	{method: "PATCH", pattern: "/repos/{str:owner}/{str:owner}/hooks/{int:id}"},
	{method: "POST", pattern: "/repos/{str:owner}/{str:owner}/hooks/{int:id}/tests"},
	{method: "DELETE", pattern: "/repos/{str:owner}/{str:owner}/hooks/{int:id}"},
	{method: "POST", pattern: "/repos/{str:owner}/{str:owner}/merges"},
	{method: "GET", pattern: "/repos/{str:owner}/{str:owner}/releases"},
	{method: "GET", pattern: "/repos/{str:owner}/{str:owner}/releases/{int:id}"},
	{method: "POST", pattern: "/repos/{str:owner}/{str:owner}/releases"},
	{method: "PATCH", pattern: "/repos/{str:owner}/{str:owner}/releases/{int:id}"},
	{method: "DELETE", pattern: "/repos/{str:owner}/{str:owner}/releases/{int:id}"},
	{method: "GET", pattern: "/repos/{str:owner}/{str:owner}/releases/{int:id}/assets"},
	{method: "GET", pattern: "/repos/{str:owner}/{str:owner}/stats/contributors"},
	{method: "GET", pattern: "/repos/{str:owner}/{str:owner}/stats/commit_activity"},
	{method: "GET", pattern: "/repos/{str:owner}/{str:owner}/stats/code_frequency"},
	{method: "GET", pattern: "/repos/{str:owner}/{str:owner}/stats/participation"},
	{method: "GET", pattern: "/repos/{str:owner}/{str:owner}/stats/punch_card"},
	{method: "GET", pattern: "/repos/{str:owner}/{str:owner}/statuses/{str:ref}"},
	{method: "POST", pattern: "/repos/{str:owner}/{str:owner}/statuses/{str:ref}"},

	// Search
	{method: "GET", pattern: "/legacy/issues/search/{str:owner}/{str:repo}sitory/{str:state}/{str:keyword}"},
	{method: "GET", pattern: "/legacy/repos/search/{str:keyword}"},
	{method: "GET", pattern: "/legacy/user/search/{str:keyword}"},
	{method: "GET", pattern: "/legacy/user/email/{str:email}"},

	// Users
	{method: "GET", pattern: "/users/{str:user}"},
	{method: "GET", pattern: "/users/{str:user}/followers"},
	{method: "GET", pattern: "/users/{str:user}/following"},
	{method: "GET", pattern: "/user/following/{str:user}"},
	{method: "GET", pattern: "/users/{str:user}/following/{str:target_user}"},
	{method: "PUT", pattern: "/user/following/{str:user}"},
	{method: "DELETE", pattern: "/user/following/{str:user}"},
	{method: "GET", pattern: "/users/{str:user}/keys"},
	{method: "GET", pattern: "/user/keys/{int:id}"},
	{method: "PATCH", pattern: "/user/keys/{int:id}"},
	{method: "DELETE", pattern: "/user/keys/{int:id}"},
}

// http://developer.github.com/v3/
var githubAPI = make([]routeTestCase, 0, len(githubStatic)+len(githubDynamic))

// Initialize full GitHub API
func init() {
	githubAPI = append(githubAPI, githubStatic...)
	githubAPI = append(githubAPI, githubDynamic...)
}

// Empty handler to return 200 for resolved routes.
func emptyHandler(_ http.ResponseWriter, _ *http.Request) {}

// Tests that all static API routes resolves correctly
func TestGitHubStatic(t *testing.T) {
	router, err := createRouter(githubStatic, emptyHandler)
	if err != nil {
		t.Fatalf("can not create router: %v", err)
	}

	req, w := createRequestResponse()

	for _, route := range githubStatic {
		resetRequestResponse(req, w, route.method, route.pattern)
		router.ServeHTTP(w, req)

		if w.Code != http.StatusOK {
			t.Errorf("can not resolve route '%+v'", route)
		}
	}
}

// Bench for static API resolving
func BenchmarkGithubStatic(b *testing.B) {
	router, err := createRouter(githubStatic, emptyHandler)
	if err != nil {
		b.Fatalf("can not create router: %v", err)
	}

	req, w := createRequestResponse()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, route := range githubStatic {
			resetRequestResponse(req, w, route.method, route.pattern)
			router.ServeHTTP(w, req)
		}
	}
}
