package jira

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"strings"
	"testing"
)

func TestIssueGet_Success(t *testing.T) {
	setup()
	defer teardown()
	testMux.HandleFunc("/rest/api/2/issue/10002", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testRequestURL(t, r, "/rest/api/2/issue/10002")

		fmt.Fprint(w, `{"expand":"renderedFields,names,schema,transitions,operations,editmeta,changelog,versionedRepresentations","id":"10002","self":"http://www.example.com/jira/rest/api/2/issue/10002","key":"EX-1","fields":{"watcher":{"self":"http://www.example.com/jira/rest/api/2/issue/EX-1/watchers","isWatching":false,"watchCount":1,"watchers":[{"self":"http://www.example.com/jira/rest/api/2/user?username=fred","name":"fred","displayName":"Fred F. User","active":false}]},"attachment":[{"self":"http://www.example.com/jira/rest/api/2.0/attachments/10000","filename":"picture.jpg","author":{"self":"http://www.example.com/jira/rest/api/2/user?username=fred","name":"fred","avatarUrls":{"48x48":"http://www.example.com/jira/secure/useravatar?size=large&ownerId=fred","24x24":"http://www.example.com/jira/secure/useravatar?size=small&ownerId=fred","16x16":"http://www.example.com/jira/secure/useravatar?size=xsmall&ownerId=fred","32x32":"http://www.example.com/jira/secure/useravatar?size=medium&ownerId=fred"},"displayName":"Fred F. User","active":false},"created":"2016-03-16T04:22:37.461+0000","size":23123,"mimeType":"image/jpeg","content":"http://www.example.com/jira/attachments/10000","thumbnail":"http://www.example.com/jira/secure/thumbnail/10000"}],"sub-tasks":[{"id":"10000","type":{"id":"10000","name":"","inward":"Parent","outward":"Sub-task"},"outwardIssue":{"id":"10003","key":"EX-2","self":"http://www.example.com/jira/rest/api/2/issue/EX-2","fields":{"status":{"iconUrl":"http://www.example.com/jira//images/icons/statuses/open.png","name":"Open"}}}}],"description":"example bug report","project":{"self":"http://www.example.com/jira/rest/api/2/project/EX","id":"10000","key":"EX","name":"Example","avatarUrls":{"48x48":"http://www.example.com/jira/secure/projectavatar?size=large&pid=10000","24x24":"http://www.example.com/jira/secure/projectavatar?size=small&pid=10000","16x16":"http://www.example.com/jira/secure/projectavatar?size=xsmall&pid=10000","32x32":"http://www.example.com/jira/secure/projectavatar?size=medium&pid=10000"},"projectCategory":{"self":"http://www.example.com/jira/rest/api/2/projectCategory/10000","id":"10000","name":"FIRST","description":"First Project Category"}},"comment":{"comments":[{"self":"http://www.example.com/jira/rest/api/2/issue/10010/comment/10000","id":"10000","author":{"self":"http://www.example.com/jira/rest/api/2/user?username=fred","name":"fred","displayName":"Fred F. User","active":false},"body":"Lorem ipsum dolor sit amet, consectetur adipiscing elit. Pellentesque eget venenatis elit. Duis eu justo eget augue iaculis fermentum. Sed semper quam laoreet nisi egestas at posuere augue semper.","updateAuthor":{"self":"http://www.example.com/jira/rest/api/2/user?username=fred","name":"fred","displayName":"Fred F. User","active":false},"created":"2016-03-16T04:22:37.356+0000","updated":"2016-03-16T04:22:37.356+0000","visibility":{"type":"role","value":"Administrators"}}]},"issuelinks":[{"id":"10001","type":{"id":"10000","name":"Dependent","inward":"depends on","outward":"is depended by"},"outwardIssue":{"id":"10004L","key":"PRJ-2","self":"http://www.example.com/jira/rest/api/2/issue/PRJ-2","fields":{"status":{"iconUrl":"http://www.example.com/jira//images/icons/statuses/open.png","name":"Open"}}}},{"id":"10002","type":{"id":"10000","name":"Dependent","inward":"depends on","outward":"is depended by"},"inwardIssue":{"id":"10004","key":"PRJ-3","self":"http://www.example.com/jira/rest/api/2/issue/PRJ-3","fields":{"status":{"iconUrl":"http://www.example.com/jira//images/icons/statuses/open.png","name":"Open"}}}}],"worklog":{"worklogs":[{"self":"http://www.example.com/jira/rest/api/2/issue/10010/worklog/10000","author":{"self":"http://www.example.com/jira/rest/api/2/user?username=fred","name":"fred","displayName":"Fred F. User","active":false},"updateAuthor":{"self":"http://www.example.com/jira/rest/api/2/user?username=fred","name":"fred","displayName":"Fred F. User","active":false},"comment":"I did some work here.","updated":"2016-03-16T04:22:37.471+0000","visibility":{"type":"group","value":"jira-developers"},"started":"2016-03-16T04:22:37.471+0000","timeSpent":"3h 20m","timeSpentSeconds":12000,"id":"100028","issueId":"10002"}]},"updated":"2016-04-06T02:36:53.594-0700","timetracking":{"originalEstimate":"10m","remainingEstimate":"3m","timeSpent":"6m","originalEstimateSeconds":600,"remainingEstimateSeconds":200,"timeSpentSeconds":400}},"names":{"watcher":"watcher","attachment":"attachment","sub-tasks":"sub-tasks","description":"description","project":"project","comment":"comment","issuelinks":"issuelinks","worklog":"worklog","updated":"updated","timetracking":"timetracking"},"schema":{}}`)
	})

	issue, _, err := testClient.Issue.Get("10002")
	if issue == nil {
		t.Error("Expected issue. Issue is nil")
	}
	if err != nil {
		t.Errorf("Error given: %s", err)
	}
}

func TestIssueCreate(t *testing.T) {
	setup()
	defer teardown()
	testMux.HandleFunc("/rest/api/2/issue/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		testRequestURL(t, r, "/rest/api/2/issue/")

		w.WriteHeader(http.StatusCreated)
		fmt.Fprint(w, `{"expand":"renderedFields,names,schema,transitions,operations,editmeta,changelog,versionedRepresentations","id":"10002","self":"http://www.example.com/jira/rest/api/2/issue/10002","key":"EX-1","fields":{"watcher":{"self":"http://www.example.com/jira/rest/api/2/issue/EX-1/watchers","isWatching":false,"watchCount":1,"watchers":[{"self":"http://www.example.com/jira/rest/api/2/user?username=fred","name":"fred","displayName":"Fred F. User","active":false}]},"attachment":[{"self":"http://www.example.com/jira/rest/api/2.0/attachments/10000","filename":"picture.jpg","author":{"self":"http://www.example.com/jira/rest/api/2/user?username=fred","name":"fred","avatarUrls":{"48x48":"http://www.example.com/jira/secure/useravatar?size=large&ownerId=fred","24x24":"http://www.example.com/jira/secure/useravatar?size=small&ownerId=fred","16x16":"http://www.example.com/jira/secure/useravatar?size=xsmall&ownerId=fred","32x32":"http://www.example.com/jira/secure/useravatar?size=medium&ownerId=fred"},"displayName":"Fred F. User","active":false},"created":"2016-03-16T04:22:37.461+0000","size":23123,"mimeType":"image/jpeg","content":"http://www.example.com/jira/attachments/10000","thumbnail":"http://www.example.com/jira/secure/thumbnail/10000"}],"sub-tasks":[{"id":"10000","type":{"id":"10000","name":"","inward":"Parent","outward":"Sub-task"},"outwardIssue":{"id":"10003","key":"EX-2","self":"http://www.example.com/jira/rest/api/2/issue/EX-2","fields":{"status":{"iconUrl":"http://www.example.com/jira//images/icons/statuses/open.png","name":"Open"}}}}],"description":"example bug report","project":{"self":"http://www.example.com/jira/rest/api/2/project/EX","id":"10000","key":"EX","name":"Example","avatarUrls":{"48x48":"http://www.example.com/jira/secure/projectavatar?size=large&pid=10000","24x24":"http://www.example.com/jira/secure/projectavatar?size=small&pid=10000","16x16":"http://www.example.com/jira/secure/projectavatar?size=xsmall&pid=10000","32x32":"http://www.example.com/jira/secure/projectavatar?size=medium&pid=10000"},"projectCategory":{"self":"http://www.example.com/jira/rest/api/2/projectCategory/10000","id":"10000","name":"FIRST","description":"First Project Category"}},"comment":{"comments":[{"self":"http://www.example.com/jira/rest/api/2/issue/10010/comment/10000","id":"10000","author":{"self":"http://www.example.com/jira/rest/api/2/user?username=fred","name":"fred","displayName":"Fred F. User","active":false},"body":"Lorem ipsum dolor sit amet, consectetur adipiscing elit. Pellentesque eget venenatis elit. Duis eu justo eget augue iaculis fermentum. Sed semper quam laoreet nisi egestas at posuere augue semper.","updateAuthor":{"self":"http://www.example.com/jira/rest/api/2/user?username=fred","name":"fred","displayName":"Fred F. User","active":false},"created":"2016-03-16T04:22:37.356+0000","updated":"2016-03-16T04:22:37.356+0000","visibility":{"type":"role","value":"Administrators"}}]},"issuelinks":[{"id":"10001","type":{"id":"10000","name":"Dependent","inward":"depends on","outward":"is depended by"},"outwardIssue":{"id":"10004L","key":"PRJ-2","self":"http://www.example.com/jira/rest/api/2/issue/PRJ-2","fields":{"status":{"iconUrl":"http://www.example.com/jira//images/icons/statuses/open.png","name":"Open"}}}},{"id":"10002","type":{"id":"10000","name":"Dependent","inward":"depends on","outward":"is depended by"},"inwardIssue":{"id":"10004","key":"PRJ-3","self":"http://www.example.com/jira/rest/api/2/issue/PRJ-3","fields":{"status":{"iconUrl":"http://www.example.com/jira//images/icons/statuses/open.png","name":"Open"}}}}],"worklog":{"worklogs":[{"self":"http://www.example.com/jira/rest/api/2/issue/10010/worklog/10000","author":{"self":"http://www.example.com/jira/rest/api/2/user?username=fred","name":"fred","displayName":"Fred F. User","active":false},"updateAuthor":{"self":"http://www.example.com/jira/rest/api/2/user?username=fred","name":"fred","displayName":"Fred F. User","active":false},"comment":"I did some work here.","updated":"2016-03-16T04:22:37.471+0000","visibility":{"type":"group","value":"jira-developers"},"started":"2016-03-16T04:22:37.471+0000","timeSpent":"3h 20m","timeSpentSeconds":12000,"id":"100028","issueId":"10002"}]},"updated":"2016-04-06T02:36:53.594-0700","timetracking":{"originalEstimate":"10m","remainingEstimate":"3m","timeSpent":"6m","originalEstimateSeconds":600,"remainingEstimateSeconds":200,"timeSpentSeconds":400}},"names":{"watcher":"watcher","attachment":"attachment","sub-tasks":"sub-tasks","description":"description","project":"project","comment":"comment","issuelinks":"issuelinks","worklog":"worklog","updated":"updated","timetracking":"timetracking"},"schema":{}}`)
	})

	i := &Issue{
		Fields: &IssueFields{
			Description: "example bug report",
		},
	}
	issue, _, err := testClient.Issue.Create(i)
	if issue == nil {
		t.Error("Expected issue. Issue is nil")
	}
	if err != nil {
		t.Errorf("Error given: %s", err)
	}
}

func TestIssueAddComment(t *testing.T) {
	setup()
	defer teardown()
	testMux.HandleFunc("/rest/api/2/issue/10000/comment", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		testRequestURL(t, r, "/rest/api/2/issue/10000/comment")

		w.WriteHeader(http.StatusCreated)
		fmt.Fprint(w, `{"self":"http://www.example.com/jira/rest/api/2/issue/10010/comment/10000","id":"10000","author":{"self":"http://www.example.com/jira/rest/api/2/user?username=fred","name":"fred","displayName":"Fred F. User","active":false},"body":"Lorem ipsum dolor sit amet, consectetur adipiscing elit. Pellentesque eget venenatis elit. Duis eu justo eget augue iaculis fermentum. Sed semper quam laoreet nisi egestas at posuere augue semper.","updateAuthor":{"self":"http://www.example.com/jira/rest/api/2/user?username=fred","name":"fred","displayName":"Fred F. User","active":false},"created":"2016-03-16T04:22:37.356+0000","updated":"2016-03-16T04:22:37.356+0000","visibility":{"type":"role","value":"Administrators"}}`)
	})

	c := &Comment{
		Body: "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Pellentesque eget venenatis elit. Duis eu justo eget augue iaculis fermentum. Sed semper quam laoreet nisi egestas at posuere augue semper.",
		Visibility: CommentVisibility{
			Type:  "role",
			Value: "Administrators",
		},
	}
	comment, _, err := testClient.Issue.AddComment("10000", c)
	if comment == nil {
		t.Error("Expected Comment. Comment is nil")
	}
	if err != nil {
		t.Errorf("Error given: %s", err)
	}
}

func TestIssueAddLink(t *testing.T) {
	setup()
	defer teardown()
	testMux.HandleFunc("/rest/api/2/issueLink", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		testRequestURL(t, r, "/rest/api/2/issueLink")

		w.WriteHeader(http.StatusOK)
	})

	il := &IssueLink{
		Type: IssueLinkType{
			Name: "Duplicate",
		},
		InwardIssue: Issue{
			Key: "HSP-1",
		},
		OutwardIssue: Issue{
			Key: "MKY-1",
		},
		Comment: Comment{
			Body: "Linked related issue!",
			Visibility: CommentVisibility{
				Type:  "group",
				Value: "jira-software-users",
			},
		},
	}
	resp, err := testClient.Issue.AddLink(il)
	if resp == nil {
		t.Error("Expected response. Response is nil")
	}
	if resp.StatusCode != 200 {
		t.Errorf("Expected Status code 200. Given %d", resp.StatusCode)
	}
	if err != nil {
		t.Errorf("Error given: %s", err)
	}
}

func TestIssueFields(t *testing.T) {
	setup()
	defer teardown()
	testMux.HandleFunc("/rest/api/2/issue/10002", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testRequestURL(t, r, "/rest/api/2/issue/10002")

		fmt.Fprint(w, `{"expand":"renderedFields,names,schema,transitions,operations,editmeta,changelog,versionedRepresentations","id":"10002","self":"http://www.example.com/jira/rest/api/2/issue/10002","key":"EX-1","fields":{"labels":["test"],"watcher":{"self":"http://www.example.com/jira/rest/api/2/issue/EX-1/watchers","isWatching":false,"watchCount":1,"watchers":[{"self":"http://www.example.com/jira/rest/api/2/user?username=fred","name":"fred","displayName":"Fred F. User","active":false}]},"attachment":[{"self":"http://www.example.com/jira/rest/api/2.0/attachments/10000","filename":"picture.jpg","author":{"self":"http://www.example.com/jira/rest/api/2/user?username=fred","name":"fred","avatarUrls":{"48x48":"http://www.example.com/jira/secure/useravatar?size=large&ownerId=fred","24x24":"http://www.example.com/jira/secure/useravatar?size=small&ownerId=fred","16x16":"http://www.example.com/jira/secure/useravatar?size=xsmall&ownerId=fred","32x32":"http://www.example.com/jira/secure/useravatar?size=medium&ownerId=fred"},"displayName":"Fred F. User","active":false},"created":"2016-03-16T04:22:37.461+0000","size":23123,"mimeType":"image/jpeg","content":"http://www.example.com/jira/attachments/10000","thumbnail":"http://www.example.com/jira/secure/thumbnail/10000"}],"sub-tasks":[{"id":"10000","type":{"id":"10000","name":"","inward":"Parent","outward":"Sub-task"},"outwardIssue":{"id":"10003","key":"EX-2","self":"http://www.example.com/jira/rest/api/2/issue/EX-2","fields":{"status":{"iconUrl":"http://www.example.com/jira//images/icons/statuses/open.png","name":"Open"}}}}],"description":"example bug report","project":{"self":"http://www.example.com/jira/rest/api/2/project/EX","id":"10000","key":"EX","name":"Example","avatarUrls":{"48x48":"http://www.example.com/jira/secure/projectavatar?size=large&pid=10000","24x24":"http://www.example.com/jira/secure/projectavatar?size=small&pid=10000","16x16":"http://www.example.com/jira/secure/projectavatar?size=xsmall&pid=10000","32x32":"http://www.example.com/jira/secure/projectavatar?size=medium&pid=10000"},"projectCategory":{"self":"http://www.example.com/jira/rest/api/2/projectCategory/10000","id":"10000","name":"FIRST","description":"First Project Category"}},"comment":{"comments":[{"self":"http://www.example.com/jira/rest/api/2/issue/10010/comment/10000","id":"10000","author":{"self":"http://www.example.com/jira/rest/api/2/user?username=fred","name":"fred","displayName":"Fred F. User","active":false},"body":"Lorem ipsum dolor sit amet, consectetur adipiscing elit. Pellentesque eget venenatis elit. Duis eu justo eget augue iaculis fermentum. Sed semper quam laoreet nisi egestas at posuere augue semper.","updateAuthor":{"self":"http://www.example.com/jira/rest/api/2/user?username=fred","name":"fred","displayName":"Fred F. User","active":false},"created":"2016-03-16T04:22:37.356+0000","updated":"2016-03-16T04:22:37.356+0000","visibility":{"type":"role","value":"Administrators"}}]},"issuelinks":[{"id":"10001","type":{"id":"10000","name":"Dependent","inward":"depends on","outward":"is depended by"},"outwardIssue":{"id":"10004L","key":"PRJ-2","self":"http://www.example.com/jira/rest/api/2/issue/PRJ-2","fields":{"status":{"iconUrl":"http://www.example.com/jira//images/icons/statuses/open.png","name":"Open"}}}},{"id":"10002","type":{"id":"10000","name":"Dependent","inward":"depends on","outward":"is depended by"},"inwardIssue":{"id":"10004","key":"PRJ-3","self":"http://www.example.com/jira/rest/api/2/issue/PRJ-3","fields":{"status":{"iconUrl":"http://www.example.com/jira//images/icons/statuses/open.png","name":"Open"}}}}],"worklog":{"worklogs":[{"self":"http://www.example.com/jira/rest/api/2/issue/10010/worklog/10000","author":{"self":"http://www.example.com/jira/rest/api/2/user?username=fred","name":"fred","displayName":"Fred F. User","active":false},"updateAuthor":{"self":"http://www.example.com/jira/rest/api/2/user?username=fred","name":"fred","displayName":"Fred F. User","active":false},"comment":"I did some work here.","updated":"2016-03-16T04:22:37.471+0000","visibility":{"type":"group","value":"jira-developers"},"started":"2016-03-16T04:22:37.471+0000","timeSpent":"3h 20m","timeSpentSeconds":12000,"id":"100028","issueId":"10002"}]},"updated":"2016-04-06T02:36:53.594-0700","timetracking":{"originalEstimate":"10m","remainingEstimate":"3m","timeSpent":"6m","originalEstimateSeconds":600,"remainingEstimateSeconds":200,"timeSpentSeconds":400}},"names":{"watcher":"watcher","attachment":"attachment","sub-tasks":"sub-tasks","description":"description","project":"project","comment":"comment","issuelinks":"issuelinks","worklog":"worklog","updated":"updated","timetracking":"timetracking"},"schema":{}}`)
	})

	issue, _, err := testClient.Issue.Get("10002")
	if issue == nil {
		t.Error("Expected issue. Issue is nil")
	}
	if !reflect.DeepEqual(issue.Fields.Labels, []string{"test"}) {
		t.Error("Expected labels for the returned issue")
	}

	if err != nil {
		t.Errorf("Error given: %s", err)
	}
}

func TestIssueDownloadAttachment(t *testing.T) {
	var testAttachment = "Here is an attachment"

	setup()
	defer teardown()
	testMux.HandleFunc("/secure/attachment/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testRequestURL(t, r, "/secure/attachment/10000/")

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(testAttachment))
	})

	resp, err := testClient.Issue.DownloadAttachment("10000")
	if resp == nil {
		t.Error("Expected response. Response is nil")
	}
	defer resp.Body.Close()

	attachment, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Error("Expected attachment text", err)
	}
	if string(attachment) != testAttachment {
		t.Errorf("Expecting an attachment: %s", string(attachment))
	}

	if resp.StatusCode != 200 {
		t.Errorf("Expected Status code 200. Given %d", resp.StatusCode)
	}
	if err != nil {
		t.Errorf("Error given: %s", err)
	}
}

func TestIssueDownloadAttachment_BadStatus(t *testing.T) {

	setup()
	defer teardown()
	testMux.HandleFunc("/secure/attachment/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testRequestURL(t, r, "/secure/attachment/10000/")

		w.WriteHeader(http.StatusForbidden)
	})

	resp, err := testClient.Issue.DownloadAttachment("10000")
	if resp == nil {
		t.Error("Expected response. Response is nil")
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusForbidden {
		t.Errorf("Expected Status code %d. Given %d", http.StatusForbidden, resp.StatusCode)
	}
	if err == nil {
		t.Errorf("Error expected")
	}
}

func TestIssuePostAttachment(t *testing.T) {
	var testAttachment = "Here is an attachment"

	setup()
	defer teardown()
	testMux.HandleFunc("/rest/api/2/issue/10000/attachments", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		testRequestURL(t, r, "/rest/api/2/issue/10000/attachments")
		status := http.StatusOK

		file, _, err := r.FormFile("file")
		if err != nil {
			status = http.StatusNotAcceptable
		}
		if file == nil {
			status = http.StatusNoContent
		} else {

			// Read the file into memory
			data, err := ioutil.ReadAll(file)
			if err != nil {
				status = http.StatusInternalServerError
			}
			if string(data) != testAttachment {
				status = http.StatusNotAcceptable
			}

			w.WriteHeader(status)
			fmt.Fprint(w, `[{"self":"http://jira/jira/rest/api/2/attachment/228924","id":"228924","filename":"example.jpg","author":{"self":"http://jira/jira/rest/api/2/user?username=test","name":"test","emailAddress":"test@test.com","avatarUrls":{"16x16":"http://jira/jira/secure/useravatar?size=small&avatarId=10082","48x48":"http://jira/jira/secure/useravatar?avatarId=10082"},"displayName":"Tester","active":true},"created":"2016-05-24T00:25:17.000-0700","size":32280,"mimeType":"image/jpeg","content":"http://jira/jira/secure/attachment/228924/example.jpg","thumbnail":"http://jira/jira/secure/thumbnail/228924/_thumb_228924.png"}]`)
			file.Close()
		}
	})

	reader := strings.NewReader(testAttachment)

	issue, resp, err := testClient.Issue.PostAttachment("10000", reader, "attachment")

	if issue == nil {
		t.Error("Expected response. Response is nil")
	}

	if resp.StatusCode != 200 {
		t.Errorf("Expected Status code 200. Given %d", resp.StatusCode)
	}

	if err != nil {
		t.Errorf("Error given: %s", err)
	}
}

func TestIssuePostAttachment_NoResponse(t *testing.T) {
	var testAttachment = "Here is an attachment"

	setup()
	defer teardown()
	testMux.HandleFunc("/rest/api/2/issue/10000/attachments", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		testRequestURL(t, r, "/rest/api/2/issue/10000/attachments")
		w.WriteHeader(http.StatusOK)
	})
	reader := strings.NewReader(testAttachment)

	_, _, err := testClient.Issue.PostAttachment("10000", reader, "attachment")

	if err == nil {
		t.Errorf("Error expected: %s", err)
	}
}

func TestIssuePostAttachment_NoFilename(t *testing.T) {
	var testAttachment = "Here is an attachment"

	setup()
	defer teardown()
	testMux.HandleFunc("/rest/api/2/issue/10000/attachments", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		testRequestURL(t, r, "/rest/api/2/issue/10000/attachments")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `[{"self":"http://jira/jira/rest/api/2/attachment/228924","id":"228924","filename":"example.jpg","author":{"self":"http://jira/jira/rest/api/2/user?username=test","name":"test","emailAddress":"test@test.com","avatarUrls":{"16x16":"http://jira/jira/secure/useravatar?size=small&avatarId=10082","48x48":"http://jira/jira/secure/useravatar?avatarId=10082"},"displayName":"Tester","active":true},"created":"2016-05-24T00:25:17.000-0700","size":32280,"mimeType":"image/jpeg","content":"http://jira/jira/secure/attachment/228924/example.jpg","thumbnail":"http://jira/jira/secure/thumbnail/228924/_thumb_228924.png"}]`)
	})
	reader := strings.NewReader(testAttachment)

	_, _, err := testClient.Issue.PostAttachment("10000", reader, "")

	if err != nil {
		t.Errorf("Error expected: %s", err)
	}
}

func TestIssuePostAttachment_NoAttachment(t *testing.T) {

	setup()
	defer teardown()
	testMux.HandleFunc("/rest/api/2/issue/10000/attachments", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		testRequestURL(t, r, "/rest/api/2/issue/10000/attachments")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `[{"self":"http://jira/jira/rest/api/2/attachment/228924","id":"228924","filename":"example.jpg","author":{"self":"http://jira/jira/rest/api/2/user?username=test","name":"test","emailAddress":"test@test.com","avatarUrls":{"16x16":"http://jira/jira/secure/useravatar?size=small&avatarId=10082","48x48":"http://jira/jira/secure/useravatar?avatarId=10082"},"displayName":"Tester","active":true},"created":"2016-05-24T00:25:17.000-0700","size":32280,"mimeType":"image/jpeg","content":"http://jira/jira/secure/attachment/228924/example.jpg","thumbnail":"http://jira/jira/secure/thumbnail/228924/_thumb_228924.png"}]`)
	})

	_, _, err := testClient.Issue.PostAttachment("10000", nil, "attachment")

	if err != nil {
		t.Errorf("Error given: %s", err)
	}
}
