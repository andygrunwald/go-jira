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
