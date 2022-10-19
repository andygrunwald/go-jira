package cloud

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"testing"
)

func TestPriorityService_GetList(t *testing.T) {
	setup()
	defer teardown()
	testapiEndpoint := "/rest/api/2/priority"

	raw, err := os.ReadFile("../testing/mock-data/all_priorities.json")
	if err != nil {
		t.Error(err.Error())
	}
	testMux.HandleFunc(testapiEndpoint, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testRequestURL(t, r, testapiEndpoint)
		fmt.Fprint(w, string(raw))
	})

	priorities, _, err := testClient.Priority.GetList(context.Background())
	if priorities == nil {
		t.Error("Expected priority list. Priority list is nil")
	}
	if err != nil {
		t.Errorf("Error given: %s", err)
	}
}
