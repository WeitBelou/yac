package yac

import (
	"net/http"
	"testing"

	"net/http/httptest"

	"github.com/weitbelou/yac"
)

// http://developer.github.com/v3/
var githubAPI = []yac.Route{
	// OAuth Authorizations
	{Method: "GET", Pattern: "/authorizations"},
	{Method: "GET", Pattern: "/authorizations/{int:id}"},
	{Method: "POST", Pattern: "/authorizations"},

	{Method: "PUT", Pattern: "/authorizations/clients/{int:client_id}"},
	{Method: "PATCH", Pattern: "/authorizations/{int:id}"},
	{Method: "DELETE", Pattern: "/authorizations/{int:id}"},
	{Method: "GET", Pattern: "/applications/{int:client_id}/tokens/{str:access_token}"},
	{Method: "DELETE", Pattern: "/applications/{int:client_id}/tokens"},
	{Method: "DELETE", Pattern: "/applications/{int:client_id}/tokens/{str:access_token}"},

	// Activity
	{Method: "GET", Pattern: "/events"},
	{Method: "GET", Pattern: "/repos/{str:owner}/{str:repo}/events"},
	{Method: "GET", Pattern: "/networks/{str:owner}/{str:repo}/events"},
	{Method: "GET", Pattern: "/orgs/{str:org}/events"},
	{Method: "GET", Pattern: "/users/{str:user}/received_events"},
	{Method: "GET", Pattern: "/users/{str:user}/received_events/public"},
	{Method: "GET", Pattern: "/users/{str:user}/events"},
	{Method: "GET", Pattern: "/users/{str:user}/events/public"},
	{Method: "GET", Pattern: "/users/{str:user}/events/orgs/{str:org}"},
	{Method: "GET", Pattern: "/feeds"},
	{Method: "GET", Pattern: "/notifications"},
	{Method: "GET", Pattern: "/repos/{str:owner}/{str:repo}/notifications"},
	{Method: "PUT", Pattern: "/notifications"},
	{Method: "PUT", Pattern: "/repos/{str:owner}/{str:repo}/notifications"},
	{Method: "GET", Pattern: "/notifications/threads/{int:id}"},
	{Method: "PATCH", Pattern: "/notifications/threads/{int:id}"},
	{Method: "GET", Pattern: "/notifications/threads/{int:id}/subscription"},
	{Method: "PUT", Pattern: "/notifications/threads/{int:id}/subscription"},
	{Method: "DELETE", Pattern: "/notifications/threads/{int:id}/subscription"},
	{Method: "GET", Pattern: "/repos/{str:owner}/{str:repo}/stargazers"},
	{Method: "GET", Pattern: "/users/{str:user}/starred"},
	{Method: "GET", Pattern: "/user/starred"},
	{Method: "GET", Pattern: "/user/starred/{str:owner}/{str:repo}"},
	{Method: "PUT", Pattern: "/user/starred/{str:owner}/{str:repo}"},
	{Method: "DELETE", Pattern: "/user/starred/{str:owner}/{str:repo}"},
	{Method: "GET", Pattern: "/repos/{str:owner}/{str:repo}/subscribers"},
	{Method: "GET", Pattern: "/users/{str:user}/subscriptions"},
	{Method: "GET", Pattern: "/user/subscriptions"},
	{Method: "GET", Pattern: "/repos/{str:owner}/{str:repo}/subscription"},
	{Method: "PUT", Pattern: "/repos/{str:owner}/{str:repo}/subscription"},
	{Method: "DELETE", Pattern: "/repos/{str:owner}/{str:repo}/subscription"},
	{Method: "GET", Pattern: "/user/subscriptions/{str:owner}/{str:repo}"},
	{Method: "PUT", Pattern: "/user/subscriptions/{str:owner}/{str:repo}"},
	{Method: "DELETE", Pattern: "/user/subscriptions/{str:owner}/{str:repo}"},

	// Gists
	{Method: "GET", Pattern: "/users/str:user/gists"},
	{Method: "GET", Pattern: "/gists"},
	{Method: "GET", Pattern: "/gists/public"},
	{Method: "GET", Pattern: "/gists/starred"},
	{Method: "GET", Pattern: "/gists/{int:id}"},
	{Method: "POST", Pattern: "/gists"},
	{Method: "PATCH", Pattern: "/gists/{int:id}"},
	{Method: "PUT", Pattern: "/gists/{int:id}/star"},
	{Method: "DELETE", Pattern: "/gists/{int:id}/star"},
	{Method: "GET", Pattern: "/gists/{int:id}/star"},
	{Method: "POST", Pattern: "/gists/{int:id}/forks"},
	{Method: "DELETE", Pattern: "/gists/{int:id}"},

	// Git Data
	{Method: "GET", Pattern: "/repos/{str:owner}/{str:repo}/git/blobs/{hex:sha}"},
	{Method: "POST", Pattern: "/repos/{str:owner}/{str:owner}/git/blobs"},
	{Method: "GET", Pattern: "/repos/{str:owner}/{str:owner}/git/commits/{hex:sha}"},
	{Method: "POST", Pattern: "/repos/{str:owner}/{str:owner}/git/commits"},
	{Method: "GET", Pattern: "/repos/{str:owner}/{str:owner}/git/refs/{str:ref}"},
	{Method: "GET", Pattern: "/repos/{str:owner}/{str:owner}/git/refs"},
	{Method: "POST", Pattern: "/repos/{str:owner}/{str:owner}/git/refs"},
	{Method: "PATCH", Pattern: "/repos/{str:owner}/{str:owner}/git/refs/{str:ref}"},
	{Method: "DELETE", Pattern: "/repos/{str:owner}/{str:owner}/git/refs/{str:ref}"},
	{Method: "GET", Pattern: "/repos/{str:owner}/{str:owner}/git/tags/{hex:sha}"},
	{Method: "POST", Pattern: "/repos/{str:owner}/{str:owner}/git/tags"},
	{Method: "GET", Pattern: "/repos/{str:owner}/{str:owner}/git/trees/{hex:sha}"},
	{Method: "POST", Pattern: "/repos/{str:owner}/{str:owner}/git/trees"},

	// Issues
	{Method: "GET", Pattern: "/issues"},
	{Method: "GET", Pattern: "/user/issues"},
	{Method: "GET", Pattern: "/orgs/{str:org}/issues"},
	{Method: "GET", Pattern: "/repos/{str:owner}/{str:owner}/issues"},
	{Method: "GET", Pattern: "/repos/{str:owner}/{str:owner}/issues/{int:user}"},
	{Method: "POST", Pattern: "/repos/{str:owner}/{str:owner}/issues"},
	{Method: "PATCH", Pattern: "/repos/{str:owner}/{str:owner}/issues/{int:user}"},
	{Method: "GET", Pattern: "/repos/{str:owner}/{str:owner}/assignees"},
	{Method: "GET", Pattern: "/repos/{str:owner}/{str:owner}/assignees/{str:assignee}"},
	{Method: "GET", Pattern: "/repos/{str:owner}/{str:owner}/issues/{int:user}/comments"},
	{Method: "GET", Pattern: "/repos/{str:owner}/{str:owner}/issues/comments"},
	{Method: "GET", Pattern: "/repos/{str:owner}/{str:owner}/issues/comments/{int:id}"},
	{Method: "POST", Pattern: "/repos/{str:owner}/{str:owner}/issues/{int:user}/comments"},
	{Method: "PATCH", Pattern: "/repos/{str:owner}/{str:owner}/issues/comments/{int:id}"},
	{Method: "DELETE", Pattern: "/repos/{str:owner}/{str:owner}/issues/comments/{int:id}"},
	{Method: "GET", Pattern: "/repos/{str:owner}/{str:owner}/issues/{int:user}/events"},
	{Method: "GET", Pattern: "/repos/{str:owner}/{str:owner}/issues/events"},
	{Method: "GET", Pattern: "/repos/{str:owner}/{str:owner}/issues/events/{int:id}"},
	{Method: "GET", Pattern: "/repos/{str:owner}/{str:owner}/labels"},
	{Method: "GET", Pattern: "/repos/{str:owner}/{str:owner}/labels/{str:name}"},
	{Method: "POST", Pattern: "/repos/{str:owner}/{str:owner}/labels"},
	{Method: "PATCH", Pattern: "/repos/{str:owner}/{str:owner}/labels/{str:name}"},
	{Method: "DELETE", Pattern: "/repos/{str:owner}/{str:owner}/labels/{str:name}"},
	{Method: "GET", Pattern: "/repos/{str:owner}/{str:owner}/issues/{int:user}/labels"},
	{Method: "POST", Pattern: "/repos/{str:owner}/{str:owner}/issues/{int:user}/labels"},
	{Method: "DELETE", Pattern: "/repos/{str:owner}/{str:owner}/issues/{int:user}/labels/{str:name}"},
	{Method: "PUT", Pattern: "/repos/{str:owner}/{str:owner}/issues/{int:user}/labels"},
	{Method: "DELETE", Pattern: "/repos/{str:owner}/{str:owner}/issues/{int:user}/labels"},
	{Method: "GET", Pattern: "/repos/{str:owner}/{str:owner}/milestones/{int:user}/labels"},
	{Method: "GET", Pattern: "/repos/{str:owner}/{str:owner}/milestones"},
	{Method: "GET", Pattern: "/repos/{str:owner}/{str:owner}/milestones/{int:user}"},
	{Method: "POST", Pattern: "/repos/{str:owner}/{str:owner}/milestones"},
	{Method: "PATCH", Pattern: "/repos/{str:owner}/{str:owner}/milestones/{int:user}"},
	{Method: "DELETE", Pattern: "/repos/{str:owner}/{str:owner}/milestones/{int:user}"},

	// Miscellaneous
	{Method: "GET", Pattern: "/emojis"},
	{Method: "GET", Pattern: "/gitignore/templates"},
	{Method: "GET", Pattern: "/gitignore/templates/{str:name}"},
	{Method: "POST", Pattern: "/markdown"},
	{Method: "POST", Pattern: "/markdown/raw"},
	{Method: "GET", Pattern: "/meta"},
	{Method: "GET", Pattern: "/rate_limit"},

	// Organizations
	{Method: "GET", Pattern: "/users/{str:user}/orgs"},
	{Method: "GET", Pattern: "/user/orgs"},
	{Method: "GET", Pattern: "/orgs/{str:org}"},
	{Method: "PATCH", Pattern: "/orgs/{str:org}"},
	{Method: "GET", Pattern: "/orgs/{str:org}/members"},
	{Method: "GET", Pattern: "/orgs/{str:org}/members/{str:user}"},
	{Method: "DELETE", Pattern: "/orgs/{str:org}/members/{str:user}"},
	{Method: "GET", Pattern: "/orgs/{str:org}/public_members"},
	{Method: "GET", Pattern: "/orgs/{str:org}/public_members/{str:user}"},
	{Method: "PUT", Pattern: "/orgs/{str:org}/public_members/{str:user}"},
	{Method: "DELETE", Pattern: "/orgs/{str:org}/public_members/{str:user}"},
	{Method: "GET", Pattern: "/orgs/{str:org}/teams"},
	{Method: "GET", Pattern: "/teams/{int:id}"},
	{Method: "POST", Pattern: "/orgs/{str:org}/teams"},
	{Method: "PATCH", Pattern: "/teams/{int:id}"},
	{Method: "DELETE", Pattern: "/teams/{int:id}"},
	{Method: "GET", Pattern: "/teams/{int:id}/members"},
	{Method: "GET", Pattern: "/teams/{int:id}/members/{str:user}"},
	{Method: "PUT", Pattern: "/teams/{int:id}/members/{str:user}"},
	{Method: "DELETE", Pattern: "/teams/{int:id}/members/{str:user}"},
	{Method: "GET", Pattern: "/teams/{int:id}/repos"},
	{Method: "GET", Pattern: "/teams/{int:id}/repos/{str:owner}/{str:repo}"},
	{Method: "PUT", Pattern: "/teams/{int:id}/repos/{str:owner}/{str:repo}"},
	{Method: "DELETE", Pattern: "/teams/{int:id}/repos/{str:owner}/{str:repo}"},
	{Method: "GET", Pattern: "/user/teams"},

	// Pull Requests
	{Method: "GET", Pattern: "/repos/{str:owner}/{str:owner}/pulls"},
	{Method: "GET", Pattern: "/repos/{str:owner}/{str:owner}/pulls/{int:user}"},
	{Method: "POST", Pattern: "/repos/{str:owner}/{str:owner}/pulls"},
	{Method: "PATCH", Pattern: "/repos/{str:owner}/{str:owner}/pulls/{int:user}"},
	{Method: "GET", Pattern: "/repos/{str:owner}/{str:owner}/pulls/{int:user}/commits"},
	{Method: "GET", Pattern: "/repos/{str:owner}/{str:owner}/pulls/{int:user}/files"},
	{Method: "GET", Pattern: "/repos/{str:owner}/{str:owner}/pulls/{int:user}/merge"},
	{Method: "PUT", Pattern: "/repos/{str:owner}/{str:owner}/pulls/{int:user}/merge"},
	{Method: "GET", Pattern: "/repos/{str:owner}/{str:owner}/pulls/{int:user}/comments"},
	{Method: "GET", Pattern: "/repos/{str:owner}/{str:owner}/pulls/comments"},
	{Method: "GET", Pattern: "/repos/{str:owner}/{str:owner}/pulls/comments/{int:user}"},
	{Method: "PUT", Pattern: "/repos/{str:owner}/{str:owner}/pulls/{int:user}/comments"},
	{Method: "PATCH", Pattern: "/repos/{str:owner}/{str:owner}/pulls/comments/{int:user}"},
	{Method: "DELETE", Pattern: "/repos/{str:owner}/{str:owner}/pulls/comments/{int:user}"},

	// Repositories
	{Method: "GET", Pattern: "/user/repos"},
	{Method: "GET", Pattern: "/users/{str:user}/repos"},
	{Method: "GET", Pattern: "/orgs/{str:org}/repos"},
	{Method: "GET", Pattern: "/repositories"},
	{Method: "POST", Pattern: "/user/repos"},
	{Method: "POST", Pattern: "/orgs/{str:org}/repos"},
	{Method: "GET", Pattern: "/repos/{str:owner}/{str:repo}"},
	{Method: "PATCH", Pattern: "/repos/{str:owner}/{str:repo}"},
	{Method: "GET", Pattern: "/repos/{str:owner}/{str:owner}/contributors"},
	{Method: "GET", Pattern: "/repos/{str:owner}/{str:owner}/languages"},
	{Method: "GET", Pattern: "/repos/{str:owner}/{str:owner}/teams"},
	{Method: "GET", Pattern: "/repos/{str:owner}/{str:owner}/tags"},
	{Method: "GET", Pattern: "/repos/{str:owner}/{str:owner}/branches"},
	{Method: "GET", Pattern: "/repos/{str:owner}/{str:owner}/branches/{str:branch}"},
	{Method: "DELETE", Pattern: "/repos/{str:owner}/{str:repo}"},
	{Method: "GET", Pattern: "/repos/{str:owner}/{str:owner}/collaborators"},
	{Method: "GET", Pattern: "/repos/{str:owner}/{str:owner}/collaborators/{str:user}"},
	{Method: "PUT", Pattern: "/repos/{str:owner}/{str:owner}/collaborators/{str:user}"},
	{Method: "DELETE", Pattern: "/repos/{str:owner}/{str:owner}/collaborators/{str:user}"},
	{Method: "GET", Pattern: "/repos/{str:owner}/{str:owner}/comments"},
	{Method: "GET", Pattern: "/repos/{str:owner}/{str:owner}/commits/{hex:sha}/comments"},
	{Method: "POST", Pattern: "/repos/{str:owner}/{str:owner}/commits/{hex:sha}/comments"},
	{Method: "GET", Pattern: "/repos/{str:owner}/{str:owner}/comments/{int:id}"},
	{Method: "PATCH", Pattern: "/repos/{str:owner}/{str:owner}/comments/{int:id}"},
	{Method: "DELETE", Pattern: "/repos/{str:owner}/{str:owner}/comments/{int:id}"},
	{Method: "GET", Pattern: "/repos/{str:owner}/{str:owner}/commits"},
	{Method: "GET", Pattern: "/repos/{str:owner}/{str:owner}/commits/{hex:sha}"},
	{Method: "GET", Pattern: "/repos/{str:owner}/{str:owner}/readme"},
	{Method: "GET", Pattern: "/repos/{str:owner}/{str:owner}/contents/*path"},
	{Method: "PUT", Pattern: "/repos/{str:owner}/{str:owner}/contents/*path"},
	{Method: "DELETE", Pattern: "/repos/{str:owner}/{str:owner}/contents/*path"},
	{Method: "GET", Pattern: "/repos/{str:owner}/{str:owner}/{str:archive_format}/{str:ref}"},
	{Method: "GET", Pattern: "/repos/{str:owner}/{str:owner}/keys"},
	{Method: "GET", Pattern: "/repos/{str:owner}/{str:owner}/keys/{int:id}"},
	{Method: "POST", Pattern: "/repos/{str:owner}/{str:owner}/keys"},
	{Method: "PATCH", Pattern: "/repos/{str:owner}/{str:owner}/keys/{int:id}"},
	{Method: "DELETE", Pattern: "/repos/{str:owner}/{str:owner}/keys/{int:id}"},
	{Method: "GET", Pattern: "/repos/{str:owner}/{str:owner}/downloads"},
	{Method: "GET", Pattern: "/repos/{str:owner}/{str:owner}/downloads/{int:id}"},
	{Method: "DELETE", Pattern: "/repos/{str:owner}/{str:owner}/downloads/{int:id}"},
	{Method: "GET", Pattern: "/repos/{str:owner}/{str:owner}/forks"},
	{Method: "POST", Pattern: "/repos/{str:owner}/{str:owner}/forks"},
	{Method: "GET", Pattern: "/repos/{str:owner}/{str:owner}/hooks"},
	{Method: "GET", Pattern: "/repos/{str:owner}/{str:owner}/hooks/{int:id}"},
	{Method: "POST", Pattern: "/repos/{str:owner}/{str:owner}/hooks"},
	{Method: "PATCH", Pattern: "/repos/{str:owner}/{str:owner}/hooks/{int:id}"},
	{Method: "POST", Pattern: "/repos/{str:owner}/{str:owner}/hooks/{int:id}/tests"},
	{Method: "DELETE", Pattern: "/repos/{str:owner}/{str:owner}/hooks/{int:id}"},
	{Method: "POST", Pattern: "/repos/{str:owner}/{str:owner}/merges"},
	{Method: "GET", Pattern: "/repos/{str:owner}/{str:owner}/releases"},
	{Method: "GET", Pattern: "/repos/{str:owner}/{str:owner}/releases/{int:id}"},
	{Method: "POST", Pattern: "/repos/{str:owner}/{str:owner}/releases"},
	{Method: "PATCH", Pattern: "/repos/{str:owner}/{str:owner}/releases/{int:id}"},
	{Method: "DELETE", Pattern: "/repos/{str:owner}/{str:owner}/releases/{int:id}"},
	{Method: "GET", Pattern: "/repos/{str:owner}/{str:owner}/releases/{int:id}/assets"},
	{Method: "GET", Pattern: "/repos/{str:owner}/{str:owner}/stats/contributors"},
	{Method: "GET", Pattern: "/repos/{str:owner}/{str:owner}/stats/commit_activity"},
	{Method: "GET", Pattern: "/repos/{str:owner}/{str:owner}/stats/code_frequency"},
	{Method: "GET", Pattern: "/repos/{str:owner}/{str:owner}/stats/participation"},
	{Method: "GET", Pattern: "/repos/{str:owner}/{str:owner}/stats/punch_card"},
	{Method: "GET", Pattern: "/repos/{str:owner}/{str:owner}/statuses/{str:ref}"},
	{Method: "POST", Pattern: "/repos/{str:owner}/{str:owner}/statuses/{str:ref}"},

	// Search
	{Method: "GET", Pattern: "/search/repositories"},
	{Method: "GET", Pattern: "/search/code"},
	{Method: "GET", Pattern: "/search/issues"},
	{Method: "GET", Pattern: "/search/users"},
	{Method: "GET", Pattern: "/legacy/issues/search/{str:owner}/{str:repo}sitory/{str:state}/{str:keyword}"},
	{Method: "GET", Pattern: "/legacy/repos/search/{str:keyword}"},
	{Method: "GET", Pattern: "/legacy/user/search/{str:keyword}"},
	{Method: "GET", Pattern: "/legacy/user/email/{str:email}"},

	// Users
	{Method: "GET", Pattern: "/users/{str:user}"},
	{Method: "GET", Pattern: "/user"},
	{Method: "PATCH", Pattern: "/user"},
	{Method: "GET", Pattern: "/users"},
	{Method: "GET", Pattern: "/user/emails"},
	{Method: "POST", Pattern: "/user/emails"},
	{Method: "DELETE", Pattern: "/user/emails"},
	{Method: "GET", Pattern: "/users/{str:user}/followers"},
	{Method: "GET", Pattern: "/user/followers"},
	{Method: "GET", Pattern: "/users/{str:user}/following"},
	{Method: "GET", Pattern: "/user/following"},
	{Method: "GET", Pattern: "/user/following/{str:user}"},
	{Method: "GET", Pattern: "/users/{str:user}/following/{str:target_user}"},
	{Method: "PUT", Pattern: "/user/following/{str:user}"},
	{Method: "DELETE", Pattern: "/user/following/{str:user}"},
	{Method: "GET", Pattern: "/users/{str:user}/keys"},
	{Method: "GET", Pattern: "/user/keys"},
	{Method: "GET", Pattern: "/user/keys/{int:id}"},
	{Method: "POST", Pattern: "/user/keys"},
	{Method: "PATCH", Pattern: "/user/keys/{int:id}"},
	{Method: "DELETE", Pattern: "/user/keys/{int:id}"},
}

// Empty handler to return 200 for resolved routes.
func emptyHandler(_ http.ResponseWriter, _ *http.Request) {}

// Check that all routes resolved correctly
func TestGitHubResolve(t *testing.T) {
	router, err := yac.NewRouter("")
	if err != nil {
		t.Fatalf("can not create router: %v", err)
	}

	for _, route := range githubAPI {
		if err := router.Route(route.Pattern, route.Method, emptyHandler); err != nil {
			t.Fatalf("can not init route '%+v': %v", route, err)
		}
	}

	for _, route := range githubAPI {
		req := httptest.NewRequest(string(route.Method), string(route.Pattern), nil)
		w := httptest.NewRecorder()

		router.ServeHTTP(w, req)

		if w.Code != http.StatusOK {
			t.Errorf("can not resolve route '%+v'", route)
		}

	}
}
