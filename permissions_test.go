package jira

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestPermissionsService_GetBulk(t *testing.T) {
	setup()
	defer teardown()
	testAPIEndpoint := "/rest/api/2/permissions/check"

	req, err := ioutil.ReadFile("./mocks/check_permissions_request.json")
	if err != nil {
		t.Error(err.Error())
	}
	var reqBody map[string]interface{}
	if err := json.Unmarshal(req, &reqBody); err != nil {
		t.Error(err.Error())
	}

	resp, err := ioutil.ReadFile("./mocks/check_permissions.json")
	if err != nil {
		t.Error(err.Error())
	}
	testMux.HandleFunc(testAPIEndpoint, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		testRequestURL(t, r, testAPIEndpoint)
		testRequestBody(t, r, reqBody)
		fmt.Fprint(w, string(resp))
	})

	bulkPerms, _, err := testClient.Permissions.GetBulkPermissions(
		"",
		[]string{"ADMINISTER"},
		[]string{"EDIT_ISSUES"},
		[]int{10001},
		[]int{10010, 10011, 10012, 10013, 10014})
	if bulkPerms == nil {
		t.Error("Expected bulkPerms list. BulkPerms list is nil")
	}
	if err != nil {
		t.Errorf("Error given: %s", err)
	}
}
