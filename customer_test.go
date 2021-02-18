package jira

import (
	"fmt"
	"net/http"
	"testing"
)

func TestCustomerService_Create(t *testing.T) {
	setup()
	defer teardown()

	const (
		wantDisplayName  = "Fred F. User"
		wantEmailAddress = "fred@example.com"
	)

	testMux.HandleFunc("/rest/servicedeskapi/customer", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		testRequestURL(t, r, "/rest/servicedeskapi/customer")

		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, `{
		  "accountId": "qm:a713c8ea-1075-4e30-9d96-891a7d181739:5ad6d3581db05e2a66fa80b",
		  "name": "qm:a713c8ea-1075-4e30-9d96-891a7d181739:5ad6d3581db05e2a66fa80b",
		  "key": "qm:a713c8ea-1075-4e30-9d96-891a7d181739:5ad6d3581db05e2a66fa80b",
		  "emailAddress": "%s",
		  "displayName": "%s",
		  "active": true,
		  "timeZone": "Australia/Sydney",
		  "_links": {
			"jiraRest": "https://your-domain.atlassian.net/rest/api/2/user?username=qm:a713c8ea-1075-4e30-9d96-891a7d181739:5ad6d3581db05e2a66fa80b",
			"avatarUrls": {
			  "48x48": "https://avatar-cdn.atlassian.com/image",
			  "24x24": "https://avatar-cdn.atlassian.com/image",
			  "16x16": "https://avatar-cdn.atlassian.com/image",
			  "32x32": "https://avatar-cdn.atlassian.com/image"
			},
			"self": "https://your-domain.atlassian.net/rest/api/2/user?username=qm:a713c8ea-1075-4e30-9d96-891a7d181739:5ad6d3581db05e2a66fa80b"
		  }
		}`, wantEmailAddress, wantDisplayName)
	})

	gotCustomer, _, err := testClient.Customer.Create(wantEmailAddress, wantDisplayName)
	if err != nil {
		t.Fatal(err)
	}

	if want, got := wantDisplayName, gotCustomer.DisplayName; want != got {
		t.Fatalf("want display name: %q, got %q", want, got)
	}

	if want, got := wantEmailAddress, gotCustomer.EmailAddress; want != got {
		t.Fatalf("want email address: %q, got %q", want, got)
	}
}
