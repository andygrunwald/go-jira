package jira

import (
	"fmt"
	"net/http"
	"testing"
)

func TestServiceDeskService_GetOrganizations(t *testing.T) {
	setup()
	defer teardown()
	testMux.HandleFunc("/rest/servicedeskapi/servicedesk/10001/organization", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testRequestURL(t, r, "/rest/servicedeskapi/servicedesk/10001/organization")

		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `{
			"_expands": [],
			"size": 3,
			"start": 3,
			"limit": 3,
			"isLastPage": false,
			"_links": {
			  "base": "https://your-domain.atlassian.net/rest/servicedeskapi",
			  "context": "context",
			  "next": "https://your-domain.atlassian.net/rest/servicedeskapi/servicedesk/10001/organization?start=6&limit=3",
			  "prev": "https://your-domain.atlassian.net/rest/servicedeskapi/servicedesk/10001/organization?start=0&limit=3"
			},
			"values": [
			  {
				"id": "1",
				"name": "Charlie Cakes Franchises",
				"_links": {
				  "self": "https://your-domain.atlassian.net/rest/servicedeskapi/organization/1"
				}
			  },
			  {
				"id": "2",
				"name": "Atlas Coffee Co",
				"_links": {
				  "self": "https://your-domain.atlassian.net/rest/servicedeskapi/organization/2"
				}
			  },
			  {
				"id": "3",
				"name": "The Adjustment Bureau",
				"_links": {
				  "self": "https://your-domain.atlassian.net/rest/servicedeskapi/organization/3"
				}
			  }
			]
		  }`)
	})

	orgs, _, err := testClient.ServiceDesk.GetOrganizations(10001, 3, 3, "")

	if orgs == nil {
		t.Error("Expected Organizations. Result is nil")
	} else if orgs.Size != 3 {
		t.Errorf("Expected size to be 3, but got %d", orgs.Size)
	}

	if err != nil {
		t.Errorf("Error given: %s", err)
	}
}

func TestServiceDeskService_AddOrganizations(t *testing.T) {
	setup()
	defer teardown()
	testMux.HandleFunc("/rest/servicedeskapi/servicedesk/10001/organization", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		testRequestURL(t, r, "/rest/servicedeskapi/servicedesk/10001/organization")

		w.WriteHeader(http.StatusNoContent)
	})

	_, err := testClient.ServiceDesk.AddOrganization(10001, 1)

	if err != nil {
		t.Errorf("Error given: %s", err)
	}
}

func TestServiceDeskService_RemoveOrganizations(t *testing.T) {
	setup()
	defer teardown()
	testMux.HandleFunc("/rest/servicedeskapi/servicedesk/10001/organization", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
		testRequestURL(t, r, "/rest/servicedeskapi/servicedesk/10001/organization")

		w.WriteHeader(http.StatusNoContent)
	})

	_, err := testClient.ServiceDesk.RemoveOrganization(10001, 1)

	if err != nil {
		t.Errorf("Error given: %s", err)
	}
}
