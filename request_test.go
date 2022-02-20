package jira

import (
	"encoding/json"
	"net/http"
	"reflect"
	"testing"
)

func TestRequestService_Create(t *testing.T) {
	setup()
	defer teardown()

	var (
		wantRequester = "qm:a713c8ea-1075-4e30-9d96-891a7d181739:5ad6d3581db05e2a66fa80b"
		gotRequester  string

		wantParticipants = []string{
			"qm:a713c8ea-1075-4e30-9d96-891a7d181739:5ad6d3581db05e2a66fa80b",
		}
		gotParticipants []string
	)

	testMux.HandleFunc("/rest/servicedeskapi/request", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		testRequestURL(t, r, "/rest/servicedeskapi/request")

		var payload struct {
			Requester    string   `json:"raiseOnBehalfOf,omitempty"`
			Participants []string `json:"requestParticipants,omitempty"`
		}

		if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
			t.Fatal(err)
		}

		gotRequester = payload.Requester
		gotParticipants = payload.Participants

		w.Write([]byte(`{
		  "_expands": [
			"participant",
			"status",
			"sla",
			"requestType",
			"serviceDesk",
			"attachment",
			"action",
			"comment"
		  ],
		  "issueId": "107001",
		  "issueKey": "HELPDESK-1",
		  "requestTypeId": "25",
		  "serviceDeskId": "10",
		  "createdDate": {
			"iso8601": "2015-10-08T14:42:00+0700",
			"jira": "2015-10-08T14:42:00.000+0700",
			"friendly": "Monday 14:42 PM",
			"epochMillis": 1444290120000
		  },
		  "reporter": {
			"accountId": "qm:a713c8ea-1075-4e30-9d96-891a7d181739:5ad6d3581db05e2a66fa80b",
			"name": "qm:a713c8ea-1075-4e30-9d96-891a7d181739:5ad6d3581db05e2a66fa80b",
			"key": "qm:a713c8ea-1075-4e30-9d96-891a7d181739:5ad6d3581db05e2a66fa80b",
			"emailAddress": "fred@example.com",
			"displayName": "Fred F. User",
			"active": true,
			"timeZone": "Australia/Sydney",
			"_links": {
			  "jiraRest": "https://your-domain.atlassian.net/rest/api/2/user?username=qm:a713c8ea-1075-4e30-9d96-891a7d181739:5ad6d3581db05e2a66fa80b",
			  "avatarUrls": {
				"48x48": "https://avatar-cdn.atlassian.com/9bc3b5bcb0db050c6d7660b28a5b86c9?s=48&d=https%3A%2F%2Fsecure.gravatar.com%2Favatar%2F9bc3b5bcb0db050c6d7660b28a5b86c9%3Fd%3Dmm%26s%3D48%26noRedirect%3Dtrue",
				"24x24": "https://avatar-cdn.atlassian.com/9bc3b5bcb0db050c6d7660b28a5b86c9?s=24&d=https%3A%2F%2Fsecure.gravatar.com%2Favatar%2F9bc3b5bcb0db050c6d7660b28a5b86c9%3Fd%3Dmm%26s%3D24%26noRedirect%3Dtrue",
				"16x16": "https://avatar-cdn.atlassian.com/9bc3b5bcb0db050c6d7660b28a5b86c9?s=16&d=https%3A%2F%2Fsecure.gravatar.com%2Favatar%2F9bc3b5bcb0db050c6d7660b28a5b86c9%3Fd%3Dmm%26s%3D16%26noRedirect%3Dtrue",
				"32x32": "https://avatar-cdn.atlassian.com/9bc3b5bcb0db050c6d7660b28a5b86c9?s=32&d=https%3A%2F%2Fsecure.gravatar.com%2Favatar%2F9bc3b5bcb0db050c6d7660b28a5b86c9%3Fd%3Dmm%26s%3D32%26noRedirect%3Dtrue"
			  },
			  "self": "https://your-domain.atlassian.net/rest/api/2/user?username=qm:a713c8ea-1075-4e30-9d96-891a7d181739:5ad6d3581db05e2a66fa80b"
			}
		  },
		  "requestFieldValues": [
			{
			  "fieldId": "summary",
			  "label": "What do you need?",
			  "value": "Request JSD help via REST"
			},
			{
			  "fieldId": "description",
			  "label": "Why do you need this?",
			  "value": "I need a new *mouse* for my Mac",
			  "renderedValue": {
				"html": "<p>I need a new <b>mouse</b> for my Mac</p>"
			  }
			}
		  ],
		  "currentStatus": {
			"status": "Waiting for Support",
			"statusCategory": "NEW",
			"statusDate": {
			  "iso8601": "2015-10-08T14:01:00+0700",
			  "jira": "2015-10-08T14:01:00.000+0700",
			  "friendly": "Today 14:01 PM",
			  "epochMillis": 1444287660000
			}
		  },
		  "_links": {
			"jiraRest": "https://your-domain.atlassian.net/rest/api/2/issue/107001",
			"web": "https://your-domain.atlassian.net/servicedesk/customer/portal/10/HELPDESK-1",
			"self": "https://your-domain.atlassian.net/rest/servicedeskapi/request/107001",
			"agent": "https://your-domain.atlassian.net/browse/HELPDESK-1"
		  }
		}`))
	})

	request := &Request{
		ServiceDeskID: "10",
		TypeID:        "25",
		FieldValues: []RequestFieldValue{
			{
				FieldID: "summary",
				Value:   "Request JSD help via REST",
			},
			{
				FieldID: "description",
				Value:   "I need a new *mouse* for my Mac",
			},
		},
	}

	_, _, err := testClient.Request.Create(wantRequester, wantParticipants, request)
	if err != nil {
		t.Fatal(err)
	}

	if wantRequester != gotRequester {
		t.Fatalf("want requester: %q, got %q", wantRequester, gotRequester)
	}

	if !reflect.DeepEqual(wantParticipants, gotParticipants) {
		t.Fatalf("want participants: %v, got %v", wantParticipants, gotParticipants)
	}
}

func TestRequestService_CreateComment(t *testing.T) {
	setup()
	defer teardown()

	testMux.HandleFunc("/rest/servicedeskapi/request/HELPDESK-1/comment", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		testRequestURL(t, r, "/rest/servicedeskapi/request/HELPDESK-1/comment")

		w.Write([]byte(`{
		  "_expands": [
			"attachment",
			"renderedBody"
		  ],
		  "id": "1000",
		  "body": "Hello there",
		  "public": true,
		  "author": {
			"accountId": "qm:a713c8ea-1075-4e30-9d96-891a7d181739:5ad6d3581db05e2a66fa80b",
			"name": "qm:a713c8ea-1075-4e30-9d96-891a7d181739:5ad6d3581db05e2a66fa80b",
			"key": "qm:a713c8ea-1075-4e30-9d96-891a7d181739:5ad6d3581db05e2a66fa80b",
			"emailAddress": "fred@example.com",
			"displayName": "Fred F. User",
			"active": true,
			"timeZone": "Australia/Sydney",
			"_links": {
			  "jiraRest": "https://your-domain.atlassian.net/rest/api/2/user?username=qm:a713c8ea-1075-4e30-9d96-891a7d181739:5ad6d3581db05e2a66fa80b",
			  "avatarUrls": {
				"48x48": "https://avatar-cdn.atlassian.com/9bc3b5bcb0db050c6d7660b28a5b86c9?s=48&d=https%3A%2F%2Fsecure.gravatar.com%2Favatar%2F9bc3b5bcb0db050c6d7660b28a5b86c9%3Fd%3Dmm%26s%3D48%26noRedirect%3Dtrue",
				"24x24": "https://avatar-cdn.atlassian.com/9bc3b5bcb0db050c6d7660b28a5b86c9?s=24&d=https%3A%2F%2Fsecure.gravatar.com%2Favatar%2F9bc3b5bcb0db050c6d7660b28a5b86c9%3Fd%3Dmm%26s%3D24%26noRedirect%3Dtrue",
				"16x16": "https://avatar-cdn.atlassian.com/9bc3b5bcb0db050c6d7660b28a5b86c9?s=16&d=https%3A%2F%2Fsecure.gravatar.com%2Favatar%2F9bc3b5bcb0db050c6d7660b28a5b86c9%3Fd%3Dmm%26s%3D16%26noRedirect%3Dtrue",
				"32x32": "https://avatar-cdn.atlassian.com/9bc3b5bcb0db050c6d7660b28a5b86c9?s=32&d=https%3A%2F%2Fsecure.gravatar.com%2Favatar%2F9bc3b5bcb0db050c6d7660b28a5b86c9%3Fd%3Dmm%26s%3D32%26noRedirect%3Dtrue"
			  },
			  "self": "https://your-domain.atlassian.net/rest/api/2/user?username=qm:a713c8ea-1075-4e30-9d96-891a7d181739:5ad6d3581db05e2a66fa80b"
			}
		  },
		  "created": {
			"iso8601": "2015-10-09T10:22:00+0700",
			"jira": "2015-10-09T10:22:00.000+0700",
			"friendly": "Today 10:22 AM",
			"epochMillis": 1444360920000
		  },
		  "_links": {
			"self": "https://your-domain.atlassian.net/rest/servicedeskapi/request/2000/comment/1000"
		  }
		}`))
	})

	comment := &RequestComment{
		Body:   "Hello there",
		Public: true,
	}

	_, _, err := testClient.Request.CreateComment("HELPDESK-1", comment)
	if err != nil {
		t.Fatal(err)
	}
}
