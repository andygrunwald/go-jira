package jira

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestSprintGetList(t *testing.T) {
	setup()
	defer teardown()

	testAPIEndpoint := "/rest/agile/1.0/board/123/sprint"

	raw, err := ioutil.ReadFile("./mocks/sprints.json")
	if err != nil {
		t.Error(err.Error())
	}

	testMux.HandleFunc(testAPIEndpoint, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testRequestURL(t, r, testAPIEndpoint)
		fmt.Fprint(w, string(raw))
	})

	sprints, _, err := testClient.Sprint.GetList("123")

	if err != nil {
		t.Errorf("Got error: %v", err)
	}

	if sprints == nil {
		t.Error("Expected sprint list. Got nil.")
	}

	if len(sprints) != 4 {
		t.Errorf("Expected 4 transitions. Got %d", len(sprints))
	}
}

func TestMoveIssueToSprint(t *testing.T) {
	setup()
	defer teardown()

	testAPIEndpoint := "/rest/agile/1.0/sprint/123/issue"

	issuesToMove := []string{"KEY-1", "KEY-2"}

	testMux.HandleFunc(testAPIEndpoint, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		testRequestURL(t, r, testAPIEndpoint)

		decoder := json.NewDecoder(r.Body)
		var payload IssuesWrapper
		err := decoder.Decode(&payload)
		if err != nil {
			t.Error("Got error: %v", err)
		}

		if payload.Issues[0] != issuesToMove[0] {
			t.Errorf("Expected %s to be in payload, got %s instead", issuesToMove[0], payload.Issues[0])
		}
	})
	_, err := testClient.Sprint.AddIssuesToSprint(123, issuesToMove)

	if err != nil {
		t.Error("Got error: %v", err)
	}
}
