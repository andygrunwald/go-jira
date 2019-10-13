package jira

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestIssueLinkTypeService_GetList(t *testing.T) {
	setup()
	defer teardown()
	testAPIEndpoint := "/rest/api/2/issueLinkType"

	raw, err := ioutil.ReadFile("./mocks/all_issuelinktypes.json")
	if err != nil {
		t.Error(err.Error())
	}
	testMux.HandleFunc(testAPIEndpoint, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testRequestURL(t, r, testAPIEndpoint)
		fmt.Fprint(w, string(raw))
	})

	linkTypes, _, err := testClient.IssueLinkType.GetList()
	if linkTypes == nil {
		t.Error("Expected issueLinkType list. LinkTypes is nil")
	}
	if err != nil {
		t.Errorf("Error give: %s", err)
	}
}
