package jira

import (
	"fmt"
	"net/http"
	"testing"
)

func TestFilterService_GetList(t *testing.T) {
	setup()
	defer teardown()
	testMux.HandleFunc("/rest/api/2/filter", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testRequestURL(t, r, "/rest/api/2/filter")

		fmt.Fprint(w, `{"expand":"renderedFields,names,schema,transitions,operations,editmeta,changelog,versionedRepresentations","id":"10002","self":"http://www.example.com/jira/rest/api/2/issue/10002","key":"EX-1","fields":{"watcher":{"self":"http://www.example.com/jira/rest/api/2/issue/EX-1/watchers","isWatching":false,"watchCount":1,"watchers":[{"self":"http://www.example.com/jira/rest/api/2/user?username=fred","name":"fred","displayName":"Fred F. User","active":false}]},"attachment":[{"self":"http://www.example.com/jira/rest/api/2.0/attachments/10000","filename":"picture.jpg","author":{"self":"http://www.example.com/jira/rest/api/2/user?username=fred","name":"fred","avatarUrls":{"48x48":"http://www.example.com/jira/secure/useravatar?size=large&ownerId=fred","24x24":"http://www.example.com/jira/secure/useravatar?size=small&ownerId=fred","16x16":"http://www.example.com/jira/secure/useravatar?size=xsmall&ownerId=fred","32x32":"http://www.example.com/jira/secure/useravatar?size=medium&ownerId=fred"},"displayName":"Fred F. User","active":false},"created":"2016-03-16T04:22:37.461+0000","size":23123,"mimeType":"image/jpeg","content":"http://www.example.com/jira/attachments/10000","thumbnail":"http://www.example.com/jira/secure/thumbnail/10000"}],"sub-tasks":[{"id":"10000","type":{"id":"10000","name":"","inward":"Parent","outward":"Sub-task"},"outwardIssue":{"id":"10003","key":"EX-2","self":"http://www.example.com/jira/rest/api/2/issue/EX-2","fields":{"status":{"iconUrl":"http://www.example.com/jira//images/icons/statuses/open.png","name":"Open"}}}}],"description":"example bug report","project":{"self":"http://www.example.com/jira/rest/api/2/project/EX","id":"10000","key":"EX","name":"Example","avatarUrls":{"48x48":"http://www.example.com/jira/secure/projectavatar?size=large&pid=10000","24x24":"http://www.example.com/jira/secure/projectavatar?size=small&pid=10000","16x16":"http://www.example.com/jira/secure/projectavatar?size=xsmall&pid=10000","32x32":"http://www.example.com/jira/secure/projectavatar?size=medium&pid=10000"},"projectCategory":{"self":"http://www.example.com/jira/rest/api/2/projectCategory/10000","id":"10000","name":"FIRST","description":"First Project Category"}},"comment":{"comments":[{"self":"http://www.example.com/jira/rest/api/2/issue/10010/comment/10000","id":"10000","author":{"self":"http://www.example.com/jira/rest/api/2/user?username=fred","name":"fred","displayName":"Fred F. User","active":false},"body":"Lorem ipsum dolor sit amet, consectetur adipiscing elit. Pellentesque eget venenatis elit. Duis eu justo eget augue iaculis fermentum. Sed semper quam laoreet nisi egestas at posuere augue semper.","updateAuthor":{"self":"http://www.example.com/jira/rest/api/2/user?username=fred","name":"fred","displayName":"Fred F. User","active":false},"created":"2016-03-16T04:22:37.356+0000","updated":"2016-03-16T04:22:37.356+0000","visibility":{"type":"role","value":"Administrators"}}]},"issuelinks":[{"id":"10001","type":{"id":"10000","name":"Dependent","inward":"depends on","outward":"is depended by"},"outwardIssue":{"id":"10004L","key":"PRJ-2","self":"http://www.example.com/jira/rest/api/2/issue/PRJ-2","fields":{"status":{"iconUrl":"http://www.example.com/jira//images/icons/statuses/open.png","name":"Open"}}}},{"id":"10002","type":{"id":"10000","name":"Dependent","inward":"depends on","outward":"is depended by"},"inwardIssue":{"id":"10004","key":"PRJ-3","self":"http://www.example.com/jira/rest/api/2/issue/PRJ-3","fields":{"status":{"iconUrl":"http://www.example.com/jira//images/icons/statuses/open.png","name":"Open"}}}}],"worklog":{"worklogs":[{"self":"http://www.example.com/jira/rest/api/2/issue/10010/worklog/10000","author":{"self":"http://www.example.com/jira/rest/api/2/user?username=fred","name":"fred","displayName":"Fred F. User","active":false},"updateAuthor":{"self":"http://www.example.com/jira/rest/api/2/user?username=fred","name":"fred","displayName":"Fred F. User","active":false},"comment":"I did some work here.","updated":"2016-03-16T04:22:37.471+0000","visibility":{"type":"group","value":"jira-developers"},"started":"2016-03-16T04:22:37.471+0000","timeSpent":"3h 20m","timeSpentSeconds":12000,"id":"100028","issueId":"10002"}]},"updated":"2016-04-06T02:36:53.594-0700","duedate":"2018-01-19","timetracking":{"originalEstimate":"10m","remainingEstimate":"3m","timeSpent":"6m","originalEstimateSeconds":600,"remainingEstimateSeconds":200,"timeSpentSeconds":400}},"names":{"watcher":"watcher","attachment":"attachment","sub-tasks":"sub-tasks","description":"description","project":"project","comment":"comment","issuelinks":"issuelinks","worklog":"worklog","updated":"updated","timetracking":"timetracking"},"schema":{}}`)
	})

	issue, _, err := testClient.Filter.GetList(nil)
	if issue == nil {
		t.Error("Expected issue. Issue is nil")
	}
	if err != nil {
		t.Errorf("Error given: %s", err)
	}
}
