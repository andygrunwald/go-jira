package cloud

import (
	"context"
	"fmt"
	"net/http"
	"testing"
)

func TestIssueLinkService_Create(t *testing.T) {
	setup()
	defer teardown()
	testMux.HandleFunc("/rest/api/3/issueLink", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		testRequestURL(t, r, "/rest/api/3/issueLink")
	})

	if _, err := testClient.IssueLink.Create(context.Background(), CreateLink{
		Type: IssueLinkType{ID: "1"},
	}); err != nil {
		t.Errorf("Error given: %s", err)
	}
}

func TestIssueLinkService_Get(t *testing.T) {
	setup()
	defer teardown()
	testMux.HandleFunc("/rest/api/3/issueLink/123", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testRequestURL(t, r, "/rest/api/3/issueLink/123")

		fmt.Fprint(w, `{"id": "10001", "inwardIssue": {"fields": {"issuetype": {"avatarId": 10002, "description": "A problem with the software.", "entityId": "9d7dd6f7-e8b6-4247-954b-7b2c9b2a5ba2", "hierarchyLevel": 0, "iconUrl": "https://your-domain.atlassian.net/secure/viewavatar?size=xsmall&avatarId=10316&avatarType=issuetype\",", "id": "1", "name": "Bug", "scope": {"project": {"id": "10000", "key": "KEY", "name": "Next Gen Project"}, "type": "PROJECT"}, "self": "https://your-domain.atlassian.net/rest/api/3/issueType/1", "subtask": false}, "priority": {"description": "Very little impact.", "iconUrl": "https://your-domain.atlassian.net/images/icons/priorities/trivial.png", "id": "2", "name": "Trivial", "self": "https://your-domain.atlassian.net/rest/api/3/priority/5", "statusColor": "#cfcfcf"}, "status": {"description": "The issue is closed.", "iconUrl": "https://your-domain.atlassian.net/images/icons/closed.gif", "id": "5", "name": "Closed", "self": "https://your-domain.atlassian.net/rest/api/3/status/5", "statusCategory": {"colorName": "green", "id": 9, "key": "completed", "self": "https://your-domain.atlassian.net/rest/api/3/statuscategory/9"}}}, "id": "10004", "key": "PR-3", "self": "https://your-domain.atlassian.net/rest/api/3/issue/PR-3"}, "outwardIssue": {"fields": {"issuetype": {"avatarId": 1, "description": "A task that needs to be done.", "hierarchyLevel": 0, "iconUrl": "https://your-domain.atlassian.net/secure/viewavatar?size=xsmall&avatarId=10299&avatarType=issuetype\",", "id": "3", "name": "Task", "self": "https://your-domain.atlassian.net/rest/api/3/issueType/3", "subtask": false}, "priority": {"description": "Major loss of function.", "iconUrl": "https://your-domain.atlassian.net/images/icons/priorities/major.png", "id": "1", "name": "Major", "self": "https://your-domain.atlassian.net/rest/api/3/priority/3", "statusColor": "#009900"}, "status": {"description": "The issue is currently being worked on.", "iconUrl": "https://your-domain.atlassian.net/images/icons/progress.gif", "id": "10000", "name": "In Progress", "self": "https://your-domain.atlassian.net/rest/api/3/status/10000", "statusCategory": {"colorName": "yellow", "id": 1, "key": "in-flight", "name": "In Progress", "self": "https://your-domain.atlassian.net/rest/api/3/statuscategory/1"}}}, "id": "10004L", "key": "PR-2", "self": "https://your-domain.atlassian.net/rest/api/3/issue/PR-2"}, "type": {"id": "1000", "inward": "Duplicated by", "name": "Duplicate", "outward": "Duplicates", "self": "https://your-domain.atlassian.net/rest/api/3/issueLinkType/1000"}}`)
	})

	if linkType, _, err := testClient.IssueLink.Get(context.Background(), "123"); err != nil {
		t.Errorf("Error given: %s", err)
	} else if linkType == nil {
		t.Error("Expected linkType. LinkType is nil")
	}
}

func TestIssueLinkService_Delete(t *testing.T) {
	setup()
	defer teardown()
	testMux.HandleFunc("/rest/api/3/issueLink/123", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)
		testRequestURL(t, r, "/rest/api/3/issueLink/123")
	})

	if _, err := testClient.IssueLink.Delete(context.Background(), "123"); err != nil {
		t.Errorf("Error given: %s", err)
	}
}
