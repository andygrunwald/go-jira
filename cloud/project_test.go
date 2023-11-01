package cloud

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"testing"
)

func TestProjectService_GetAll(t *testing.T) {
	setup()
	defer teardown()
	testapiEndpoint := "/rest/api/2/project"

	raw, err := os.ReadFile("../testing/mock-data/all_projects.json")
	if err != nil {
		t.Error(err.Error())
	}
	testMux.HandleFunc(testapiEndpoint, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testRequestURL(t, r, "/rest/api/2/project?expand=issueTypes")
		fmt.Fprint(w, string(raw))
	})

	projects, _, err := testClient.Project.GetAll(context.Background(), &GetQueryOptions{Expand: "issueTypes"})
	if projects == nil {
		t.Error("Expected project list. Project list is nil")
	}
	if err != nil {
		t.Errorf("Error given: %s", err)
	}
}

func TestProjectService_Get(t *testing.T) {
	setup()
	defer teardown()
	testapiEndpoint := "/rest/api/2/project/12310505"

	raw, err := os.ReadFile("../testing/mock-data/project.json")
	if err != nil {
		t.Error(err.Error())
	}
	testMux.HandleFunc(testapiEndpoint, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testRequestURL(t, r, testapiEndpoint)
		fmt.Fprint(w, string(raw))
	})

	projects, _, err := testClient.Project.Get(context.Background(), "12310505")
	if err != nil {
		t.Errorf("Error given: %s", err)
	}
	if projects == nil {
		t.Error("Expected project list. Project list is nil")
		return
	}
	if len(projects.Roles) != 9 {
		t.Errorf("Expected 9 roles but got %d", len(projects.Roles))
	}
}

func TestProjectService_Get_NoProject(t *testing.T) {
	setup()
	defer teardown()
	testapiEndpoint := "/rest/api/2/project/99999999"

	testMux.HandleFunc(testapiEndpoint, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testRequestURL(t, r, testapiEndpoint)
		fmt.Fprint(w, nil)
	})

	projects, resp, err := testClient.Project.Get(context.Background(), "99999999")
	if projects != nil {
		t.Errorf("Expected nil. Got %+v", projects)
	}

	if resp.Status == "404" {
		t.Errorf("Expected status 404. Got %s", resp.Status)
	}
	if err == nil {
		t.Errorf("Error given: %s", err)
	}
}

func TestProjectService_GetPermissionScheme_Failure(t *testing.T) {
	setup()
	defer teardown()
	testapiEndpoint := "/rest/api/2/project/99999999/permissionscheme"

	testMux.HandleFunc(testapiEndpoint, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testRequestURL(t, r, testapiEndpoint)
		fmt.Fprint(w, nil)
	})

	permissionScheme, resp, err := testClient.Project.GetPermissionScheme(context.Background(), "99999999")
	if permissionScheme != nil {
		t.Errorf("Expected nil. Got %+v", permissionScheme)
	}

	if resp.Status == "404" {
		t.Errorf("Expected status 404. Got %s", resp.Status)
	}
	if err == nil {
		t.Errorf("Error given: %s", err)
	}
}

func TestProjectService_GetPermissionScheme_Success(t *testing.T) {
	setup()
	defer teardown()
	testapiEndpoint := "/rest/api/2/project/99999999/permissionscheme"

	testMux.HandleFunc(testapiEndpoint, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testRequestURL(t, r, testapiEndpoint)
		fmt.Fprint(w, `{
			"expand": "permissions,user,group,projectRole,field,all",
			"id": 10201,
			"self": "https://www.example.com/rest/api/2/permissionscheme/10201",
			"name": "Project for specific-users",
			"description": "Projects that can only see for people belonging to specific-users group"
		}`)
	})

	permissionScheme, resp, err := testClient.Project.GetPermissionScheme(context.Background(), "99999999")
	if permissionScheme.ID != 10201 {
		t.Errorf("Expected Permission Scheme ID. Got %+v", permissionScheme)
	}

	if resp.Status == "404" {
		t.Errorf("Expected status 404. Got %s", resp.Status)
	}
	if err != nil {
		t.Errorf("Error given: %s", err)
	}
}

func TestProjectService_Find_Success(t *testing.T) {
	setup()
	defer teardown()
	testapiEndpoint := "/rest/api/2/project/search"

	testMux.HandleFunc(testapiEndpoint, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testRequestURL(t, r, testapiEndpoint+"?startAt=0&maxResults=2")
		fmt.Fprint(w, `{
			"self": "https://your-domain.atlassian.net/rest/api/2/project/search?startAt=0&maxResults=2",
			"nextPage": "https://your-domain.atlassian.net/rest/api/2/project/search?startAt=2&maxResults=2",
			"maxResults": 2,
			"startAt": 0,
			"total": 7,
			"isLast": false,
			"values": [
				{
					"self": "https://your-domain.atlassian.net/rest/api/2/project/EX",
					"id": "10000",
					"key": "EX",
					"name": "Example",
					"avatarUrls": {
						"48x48": "https://your-domain.atlassian.net/secure/projectavatar?size=large&pid=10000",
						"24x24": "https://your-domain.atlassian.net/secure/projectavatar?size=small&pid=10000",
						"16x16": "https://your-domain.atlassian.net/secure/projectavatar?size=xsmall&pid=10000",
						"32x32": "https://your-domain.atlassian.net/secure/projectavatar?size=medium&pid=10000"
					},
					"projectCategory": {
						"self": "https://your-domain.atlassian.net/rest/api/2/projectCategory/10000",
						"id": "10000",
						"name": "FIRST",
						"description": "First Project Category"
					},
					"simplified": false,
					"style": "classic",
					"insight": {
						"totalIssueCount": 100,
						"lastIssueUpdateTime": "2023-10-27T00:46:39.889+0000"
					}
				},
				{
					"self": "https://your-domain.atlassian.net/rest/api/2/project/ABC",
					"id": "10001",
					"key": "ABC",
					"name": "Alphabetical",
					"avatarUrls": {
						"48x48": "https://your-domain.atlassian.net/secure/projectavatar?size=large&pid=10001",
						"24x24": "https://your-domain.atlassian.net/secure/projectavatar?size=small&pid=10001",
						"16x16": "https://your-domain.atlassian.net/secure/projectavatar?size=xsmall&pid=10001",
						"32x32": "https://your-domain.atlassian.net/secure/projectavatar?size=medium&pid=10001"
					},
					"projectCategory": {
						"self": "https://your-domain.atlassian.net/rest/api/2/projectCategory/10000",
						"id": "10000",
						"name": "FIRST",
						"description": "First Project Category"
					},
					"simplified": false,
					"style": "classic",
					"insight": {
						"totalIssueCount": 100,
						"lastIssueUpdateTime": "2023-10-27T00:46:39.889+0000"
					}
				}
			]
		}`)
	})

	if projects, _, err := testClient.Project.Find(context.Background(), WithStartAt(0), WithMaxResults(2)); err != nil {
		t.Errorf("Error given: %s", err)
	} else if len(projects) != 2 {
		t.Errorf("Expected 2 projects. Projects is %d", len(projects))
	} else if projects[0].ID != "10000" {
		t.Errorf("Expected 10000. Projects[0].ID is %s", projects[0].ID)
	}
}
