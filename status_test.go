package jira

import (
	"fmt"
	"net/http"
	"os"
	"testing"
)

func TestStatusService_GetAllStatuses(t *testing.T) {
	setup()
	defer teardown()
	testAPIEdpoint := "/rest/api/2/status"

	raw, err := os.ReadFile("./mocks/all_statuses.json")
	if err != nil {
		t.Error(err.Error())
	}

	testMux.HandleFunc(testAPIEdpoint, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testRequestURL(t, r, testAPIEdpoint)
		fmt.Fprint(w, string(raw))
	})

	statusList, _, err := testClient.Status.GetAllStatuses()

	if statusList == nil {
		t.Error("Expected statusList. statusList is nill")
	}

	if err != nil {
		t.Errorf("Error given: %s", err)
	}
}
