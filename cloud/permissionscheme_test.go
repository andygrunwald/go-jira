package cloud

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"testing"
)

func TestPermissionSchemeService_GetList(t *testing.T) {
	setup()
	defer teardown()
	testAPIEndpoint := "/rest/api/3/permissionscheme"

	raw, err := os.ReadFile("../testing/mock-data/all_permissionschemes.json")
	if err != nil {
		t.Error(err.Error())
	}
	testMux.HandleFunc(testAPIEndpoint, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testRequestURL(t, r, testAPIEndpoint)
		fmt.Fprint(w, string(raw))
	})

	permissionScheme, _, err := testClient.PermissionScheme.GetList(context.Background())
	if err != nil {
		t.Errorf("Error given: %v", err)
	}
	if permissionScheme == nil {
		t.Error("Expected permissionScheme list. PermissionScheme list is nil")
		return
	}
	if len(permissionScheme.PermissionSchemes) != 2 {
		t.Errorf("Expected %d permissionSchemes but got %d", 2, len(permissionScheme.PermissionSchemes))
	}
}

func TestPermissionSchemeService_GetList_NoList(t *testing.T) {
	setup()
	defer teardown()
	testAPIEndpoint := "/rest/api/3/permissionscheme"

	raw, err := os.ReadFile("../testing/mock-data/no_permissionschemes.json")
	if err != nil {
		t.Error(err.Error())
	}
	testMux.HandleFunc(testAPIEndpoint, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testRequestURL(t, r, testAPIEndpoint)
		fmt.Fprint(w, string(raw))
	})

	permissionScheme, _, err := testClient.PermissionScheme.GetList(context.Background())
	if permissionScheme != nil {
		t.Errorf("Expected permissionScheme list has %d entries but should be nil", len(permissionScheme.PermissionSchemes))
	}
	if err == nil {
		t.Errorf("No error given")
	}
}

func TestPermissionSchemeService_Get(t *testing.T) {
	setup()
	defer teardown()
	testapiEndpoint := "/rest/api/3/permissionscheme/10100"
	raw, err := os.ReadFile("../testing/mock-data/permissionscheme.json")
	if err != nil {
		t.Error(err.Error())
	}
	testMux.HandleFunc(testapiEndpoint, func(writer http.ResponseWriter, request *http.Request) {
		testMethod(t, request, http.MethodGet)
		testRequestURL(t, request, testapiEndpoint)
		fmt.Fprint(writer, string(raw))
	})

	permissionScheme, _, err := testClient.PermissionScheme.Get(context.Background(), 10100)
	if permissionScheme == nil {
		t.Errorf("Expected permissionscheme, got nil")
	}
	if err != nil {
		t.Errorf("Error given: %s", err)
	}
}

func TestPermissionSchemeService_Get_NoScheme(t *testing.T) {
	setup()
	defer teardown()
	testapiEndpoint := "/rest/api/3/permissionscheme/99999"
	raw, err := os.ReadFile("../testing/mock-data/no_permissionscheme.json")
	if err != nil {
		t.Error(err.Error())
	}
	testMux.HandleFunc(testapiEndpoint, func(writer http.ResponseWriter, request *http.Request) {
		testMethod(t, request, http.MethodGet)
		testRequestURL(t, request, testapiEndpoint)
		fmt.Fprint(writer, string(raw))
	})

	permissionScheme, _, err := testClient.PermissionScheme.Get(context.Background(), 99999)
	if permissionScheme != nil {
		t.Errorf("Expected nil, got permissionschme %v", permissionScheme)
	}
	if err == nil {
		t.Errorf("No error given")
	}
}
