package cloud

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"testing"
)

func TestRoleService_GetList_NoList(t *testing.T) {
	setup()
	defer teardown()
	testAPIEndpoint := "/rest/api/3/role"

	raw, err := os.ReadFile("../testing/mock-data/no_roles.json")
	if err != nil {
		t.Error(err.Error())
	}

	testMux.HandleFunc(testAPIEndpoint, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testRequestURL(t, r, testAPIEndpoint)
		fmt.Fprint(w, string(raw))
	})

	roles, _, err := testClient.Role.GetList(context.Background())
	if roles != nil {
		t.Errorf("Expected role list has %d entries but should be nil", len(*roles))
	}
	if err == nil {
		t.Errorf("No error given")
	}
}

func TestRoleService_GetList(t *testing.T) {
	setup()
	defer teardown()
	testAPIEndpoint := "/rest/api/3/role"

	raw, err := os.ReadFile("../testing/mock-data/all_roles.json")
	if err != nil {
		t.Error(err.Error())
	}
	testMux.HandleFunc(testAPIEndpoint, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testRequestURL(t, r, testAPIEndpoint)
		fmt.Fprint(w, string(raw))
	})

	roles, _, err := testClient.Role.GetList(context.Background())
	if err != nil {
		t.Errorf("Error given: %v", err)
	}
	if roles == nil {
		t.Error("Expected role list. Role list is nil")
		return
	}
	if len(*roles) != 2 {
		t.Errorf("Expected %d roles but got %d", 2, len(*roles))
	}
}

func TestRoleService_Get_NoRole(t *testing.T) {
	setup()
	defer teardown()
	testapiEndpoint := "/rest/api/3/role/99999"
	raw, err := os.ReadFile("../testing/mock-data/no_role.json")
	if err != nil {
		t.Error(err.Error())
	}
	testMux.HandleFunc(testapiEndpoint, func(writer http.ResponseWriter, request *http.Request) {
		testMethod(t, request, http.MethodGet)
		testRequestURL(t, request, testapiEndpoint)
		fmt.Fprint(writer, string(raw))
	})

	role, _, err := testClient.Role.Get(context.Background(), 99999)
	if role != nil {
		t.Errorf("Expected nil, got role %v", role)
	}
	if err == nil {
		t.Errorf("No error given")
	}
}

func TestRoleService_Get(t *testing.T) {
	setup()
	defer teardown()
	testapiEndpoint := "/rest/api/3/role/10002"
	raw, err := os.ReadFile("../testing/mock-data/role.json")
	if err != nil {
		t.Error(err.Error())
	}
	testMux.HandleFunc(testapiEndpoint, func(writer http.ResponseWriter, request *http.Request) {
		testMethod(t, request, http.MethodGet)
		testRequestURL(t, request, testapiEndpoint)
		fmt.Fprint(writer, string(raw))
	})

	role, _, err := testClient.Role.Get(context.Background(), 10002)
	if role == nil {
		t.Errorf("Expected Role, got nil")
	}
	if err != nil {
		t.Errorf("Error given: %s", err)
	}
}
