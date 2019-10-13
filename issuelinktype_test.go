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

func TestIssueLinkTypeService_Get(t *testing.T) {
	setup()
	defer teardown()
	testMux.HandleFunc("/rest/api/2/issueLinkType/123", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testRequestURL(t, r, "/rest/api/2/issueLinkType/123")

		fmt.Fprint(w, `{"id": "123","name": "Blocked","inward": "Blocked","outward": "Blocked",
		"self": "https://www.example.com/jira/rest/api/2/issueLinkType/123"}`)
	})

	if linkType, _, err := testClient.IssueLinkType.Get("123"); err != nil {
		t.Errorf("Error given: %s", err)
	} else if linkType == nil {
		t.Error("Expected linkType. LinkType is nil")
	}
}
