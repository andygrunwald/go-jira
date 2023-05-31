package cloud

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"testing"
)

func TestStatusService_GetAllStatuses(t *testing.T) {
	setup()
	defer teardown()
	testapiEndpoint := "/rest/api/2/status"

	raw, err := os.ReadFile("../testing/mock-data/all_statuses.json")
	if err != nil {
		t.Error(err.Error())
	}

	testMux.HandleFunc(testapiEndpoint, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testRequestURL(t, r, testapiEndpoint)
		fmt.Fprint(w, string(raw))
	})

	statusList, _, err := testClient.Status.GetAllStatuses(context.Background())

	if statusList == nil {
		t.Error("Expected statusList. statusList is nill")
	}

	if err != nil {
		t.Errorf("Error given: %s", err)
	}
}
