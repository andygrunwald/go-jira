package jira

import (
	"encoding/json"
	"net/http"
	"testing"
)

func TestSprintService_MoveIssuesToSprint(t *testing.T) {
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
			t.Errorf("Got error: %v", err)
		}

		if payload.Issues[0] != issuesToMove[0] {
			t.Errorf("Expected %s to be in payload, got %s instead", issuesToMove[0], payload.Issues[0])
		}
	})
	_, err := testClient.Sprint.MoveIssuesToSprint(123, issuesToMove)

	if err != nil {
		t.Errorf("Got error: %v", err)
	}
}
