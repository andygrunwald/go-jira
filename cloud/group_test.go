package cloud

import (
	"context"
	"fmt"
	"net/http"
	"testing"
)

func TestGroupService_GetPage(t *testing.T) {
	setup()
	defer teardown()
	testMux.HandleFunc("/rest/api/2/group/member", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testRequestURL(t, r, "/rest/api/2/group/member?groupname=default")
		startAt := r.URL.Query().Get("startAt")
		if startAt == "0" {
			fmt.Fprint(w, `{"self":"http://www.example.com/jira/rest/api/2/group/member?includeInactiveUsers=false&maxResults=2&groupname=default&startAt=0","nextPage":"`+testServer.URL+`/rest/api/2/group/member?groupname=default&includeInactiveUsers=false&maxResults=2&startAt=2","maxResults":2,"startAt":0,"total":4,"isLast":false,"values":[{"self":"http://www.example.com/jira/rest/api/2/user?username=michael","name":"michael","key":"michael","emailAddress":"michael@example.com","displayName":"MichaelScofield","active":true,"timeZone":"Australia/Sydney"},{"self":"http://www.example.com/jira/rest/api/2/user?username=alex","name":"alex","key":"alex","emailAddress":"alex@example.com","displayName":"AlexanderMahone","active":true,"timeZone":"Australia/Sydney"}]}`)
		} else if startAt == "2" {
			fmt.Fprint(w, `{"self":"http://www.example.com/jira/rest/api/2/group/member?includeInactiveUsers=false&maxResults=2&groupname=default&startAt=2","maxResults":2,"startAt":2,"total":4,"isLast":true,"values":[{"self":"http://www.example.com/jira/rest/api/2/user?username=michael","name":"michael","key":"michael","emailAddress":"michael@example.com","displayName":"MichaelScofield","active":true,"timeZone":"Australia/Sydney"},{"self":"http://www.example.com/jira/rest/api/2/user?username=alex","name":"alex","key":"alex","emailAddress":"alex@example.com","displayName":"AlexanderMahone","active":true,"timeZone":"Australia/Sydney"}]}`)
		} else {
			t.Errorf("startAt %s", startAt)
		}
	})
	if page, resp, err := testClient.Group.Get(context.Background(), "default", &GroupSearchOptions{
		StartAt:              0,
		MaxResults:           2,
		IncludeInactiveUsers: false,
	}); err != nil {
		t.Errorf("Error given: %s %s", err, testServer.URL)
	} else if page == nil || len(page) != 2 {
		t.Error("Expected members. Group.Members is not 2 or is nil")
	} else {
		if resp.StartAt != 0 {
			t.Errorf("Expect Result StartAt to be 0, but is %d", resp.StartAt)
		}
		if resp.MaxResults != 2 {
			t.Errorf("Expect Result MaxResults to be 2, but is %d", resp.MaxResults)
		}
		if resp.Total != 4 {
			t.Errorf("Expect Result Total to be 4, but is %d", resp.Total)
		}
		if page, resp, err := testClient.Group.Get(context.Background(), "default", &GroupSearchOptions{
			StartAt:              2,
			MaxResults:           2,
			IncludeInactiveUsers: false,
		}); err != nil {
			t.Errorf("Error give: %s %s", err, testServer.URL)
		} else if page == nil || len(page) != 2 {
			t.Error("Expected members. Group.Members is not 2 or is nil")
		} else {
			if resp.StartAt != 2 {
				t.Errorf("Expect Result StartAt to be 2, but is %d", resp.StartAt)
			}
			if resp.MaxResults != 2 {
				t.Errorf("Expect Result MaxResults to be 2, but is %d", resp.MaxResults)
			}
			if resp.Total != 4 {
				t.Errorf("Expect Result Total to be 4, but is %d", resp.Total)
			}
		}
	}
}

func TestGroupService_Add(t *testing.T) {
	setup()
	defer teardown()
	testMux.HandleFunc("/rest/api/3/group/user", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		testRequestURL(t, r, "/rest/api/3/group/user?groupname=default")

		w.WriteHeader(http.StatusCreated)
		fmt.Fprint(w, `{"name":"default","self":"http://www.example.com/jira/rest/api/2/group?groupname=default","users":{"size":1,"items":[],"max-results":50,"start-index":0,"end-index":0},"expand":"users"}`)
	})

	if group, _, err := testClient.Group.AddUserByGroupName(context.Background(), "default", "5b10ac8d82e05b22cc7d4ef5"); err != nil {
		t.Errorf("Error given: %s", err)
	} else if group == nil {
		t.Error("Expected group. Group is nil")
	}
}

func TestGroupService_Remove(t *testing.T) {
	setup()
	defer teardown()
	testMux.HandleFunc("/rest/api/3/group/user", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)
		testRequestURL(t, r, "/rest/api/3/group/user?groupname=default")

		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `{"name":"default","self":"http://www.example.com/jira/rest/api/2/group?groupname=default","users":{"size":1,"items":[],"max-results":50,"start-index":0,"end-index":0},"expand":"users"}`)
	})

	if _, err := testClient.Group.RemoveUserByGroupName(context.Background(), "default", "5b10ac8d82e05b22cc7d4ef5"); err != nil {
		t.Errorf("Error given: %s", err)
	}
}

func TestGroupService_Find_Success(t *testing.T) {
	setup()
	defer teardown()
	testMux.HandleFunc("/rest/api/3/groups/picker", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testRequestURL(t, r, "/rest/api/3/groups/picker")

		fmt.Fprint(w, `{"header": "Showing 2 of 2 matching groups",
										"total": 2,
		                "groups": [{
											"name": "jdog-developers",
											"html": "<b>j</b>dog-developers",
											"groupId": "276f955c-63d7-42c8-9520-92d01dca0625"
										},
										{
											"name": "juvenal-bot",
											"html": "<b>j</b>uvenal-bot",
											"groupId": "6e87dc72-4f1f-421f-9382-2fee8b652487"
										}]}`)
	})

	if group, _, err := testClient.Group.Find(context.Background()); err != nil {
		t.Errorf("Error given: %s", err)
	} else if group == nil {
		t.Error("Expected group. Group is nil")
	} else if len(group) != 2 {
		t.Errorf("Expected 2 groups. Group is %d", len(group))
	}
}

func TestGroupService_Find_SuccessParams(t *testing.T) {
	setup()
	defer teardown()
	testMux.HandleFunc("/rest/api/3/groups/picker", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testRequestURL(t, r, "/rest/api/3/groups/picker?maxResults=2&caseInsensitive=true&excludeId=1&excludeId=2&exclude=test&query=test&accountId=123")

		fmt.Fprint(w, `{"header": "Showing 2 of 2 matching groups",
										"total": 2,
		                "groups": [{
											"name": "jdog-developers",
											"html": "<b>j</b>dog-developers",
											"groupId": "276f955c-63d7-42c8-9520-92d01dca0625"
										},
										{
											"name": "juvenal-bot",
											"html": "<b>j</b>uvenal-bot",
											"groupId": "6e87dc72-4f1f-421f-9382-2fee8b652487"
										}]}`)
	})

	if group, _, err := testClient.Group.Find(
		context.Background(),
		WithMaxResults(2),
		WithCaseInsensitive(),
		WithExcludedGroupsIds([]string{"1", "2"}),
		WithExcludedGroupNames([]string{"test"}),
		WithGroupNameContains("test"),
		WithAccountId("123"),
	); err != nil {
		t.Errorf("Error given: %s", err)
	} else if group == nil {
		t.Error("Expected group. Group is nil")
	} else if len(group) != 2 {
		t.Errorf("Expected 2 groups. Group is %d", len(group))
	}
}

func TestGroupService_GetGroupMembers_Success(t *testing.T) {
	setup()
	defer teardown()
	testMux.HandleFunc("/rest/api/3/group/member", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testRequestURL(t, r, "/rest/api/3/group/member?groupId=1&startAt=0&maxResults=2&includeInactiveUsers=true")

		fmt.Fprint(w, `{
			"self": "https://your-domain.atlassian.net/rest/api/3/group/member?groupname=jira-administrators&includeInactiveUsers=false&startAt=2&maxResults=2",
			"nextPage": "https://your-domain.atlassian.net/rest/api/3/group/member?groupname=jira-administrators&includeInactiveUsers=false&startAt=4&maxResults=2",
			"maxResults": 2,
			"startAt": 3,
			"total": 5,
			"isLast": false,
			"values": [
				{
					"self": "https://your-domain.atlassian.net/rest/api/3/user?accountId=5b10a2844c20165700ede21g",
					"name": "",
					"key": "",
					"accountId": "5b10a2844c20165700ede21g",
					"emailAddress": "mia@example.com",
					"avatarUrls": {},
					"displayName": "Mia",
					"active": true,
					"timeZone": "Australia/Sydney",
					"accountType": "atlassian"
				},
				{
					"self": "https://your-domain.atlassian.net/rest/api/3/user?accountId=5b10a0effa615349cb016cd8",
					"name": "",
					"key": "",
					"accountId": "5b10a0effa615349cb016cd8",
					"emailAddress": "will@example.com",
					"avatarUrls": {},
					"displayName": "Will",
					"active": false,
					"timeZone": "Australia/Sydney",
					"accountType": "atlassian"
				}
			]
		}`)
	})

	if members, _, err := testClient.Group.GetGroupMembers(context.Background(), "1", WithStartAt(0), WithMaxResults(2), WithInactiveUsers()); err != nil {
		t.Errorf("Error given: %s", err)
	} else if len(members) != 2 {
		t.Errorf("Expected 2 members. Members is %d", len(members))
	} else if members[0].AccountID != "5b10a2844c20165700ede21g" {
		t.Errorf("Expected 5b10a2844c20165700ede21g. Members[0].AccountId is %s", members[0].AccountID)
	}
}
