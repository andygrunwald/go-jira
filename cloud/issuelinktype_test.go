package cloud

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"testing"
)

func TestIssueLinkTypeService_GetList(t *testing.T) {
	setup()
	defer teardown()
	testAPIEndpoint := "/rest/api/2/issueLinkType"

	raw, err := os.ReadFile("../testing/mock-data/all_issuelinktypes.json")
	if err != nil {
		t.Error(err.Error())
	}
	testMux.HandleFunc(testAPIEndpoint, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testRequestURL(t, r, testAPIEndpoint)
		fmt.Fprint(w, string(raw))
	})

	linkTypes, _, err := testClient.IssueLinkType.GetList(context.Background())
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
		testMethod(t, r, http.MethodGet)
		testRequestURL(t, r, "/rest/api/2/issueLinkType/123")

		fmt.Fprint(w, `{"id": "123","name": "Blocked","inward": "Blocked","outward": "Blocked",
		"self": "https://www.example.com/jira/rest/api/2/issueLinkType/123"}`)
	})

	if linkType, _, err := testClient.IssueLinkType.Get(context.Background(), "123"); err != nil {
		t.Errorf("Error given: %s", err)
	} else if linkType == nil {
		t.Error("Expected linkType. LinkType is nil")
	}
}

func TestIssueLinkTypeService_Create(t *testing.T) {
	setup()
	defer teardown()
	testMux.HandleFunc("/rest/api/2/issueLinkType", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		testRequestURL(t, r, "/rest/api/2/issueLinkType")

		w.WriteHeader(http.StatusCreated)
		fmt.Fprint(w, `{"id":"10021","name":"Problem/Incident","inward":"is caused by",
		"outward":"causes","self":"https://www.example.com/jira/rest/api/2/issueLinkType/10021"}`)
	})

	lt := &IssueLinkType{
		Name:    "Problem/Incident",
		Inward:  "is caused by",
		Outward: "causes",
	}

	if linkType, _, err := testClient.IssueLinkType.Create(context.Background(), lt); err != nil {
		t.Errorf("Error given: %s", err)
	} else if linkType == nil {
		t.Error("Expected linkType. LinkType is nil")
	}
}

func TestIssueLinkTypeService_Update(t *testing.T) {
	setup()
	defer teardown()
	testMux.HandleFunc("/rest/api/2/issueLinkType/100", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPut)
		testRequestURL(t, r, "/rest/api/2/issueLinkType/100")

		w.WriteHeader(http.StatusNoContent)
	})

	lt := &IssueLinkType{
		ID:      "100",
		Name:    "Problem/Incident",
		Inward:  "is caused by",
		Outward: "causes",
	}

	if linkType, _, err := testClient.IssueLinkType.Update(context.Background(), lt); err != nil {
		t.Errorf("Error given: %s", err)
	} else if linkType == nil {
		t.Error("Expected linkType. LinkType is nil")
	}
}

func TestIssueLinkTypeService_Delete(t *testing.T) {
	setup()
	defer teardown()
	testMux.HandleFunc("/rest/api/2/issueLinkType/100", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)
		testRequestURL(t, r, "/rest/api/2/issueLinkType/100")

		w.WriteHeader(http.StatusNoContent)
	})

	resp, err := testClient.IssueLinkType.Delete(context.Background(), "100")
	if resp.StatusCode != http.StatusNoContent {
		t.Error("Expected issue not deleted.")
	}
	if err != nil {
		t.Errorf("Error given: %s", err)
	}
}
