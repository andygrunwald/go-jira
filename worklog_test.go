package jira

import (
	"testing"
	"net/http"
	"fmt"
	"io/ioutil"
	"time"
)

func TestWorklogService_GetWorkLogs(t *testing.T) {
	setup()
	defer teardown()

	testAPIEndpoint := "/rest/tempo-timesheets/3/worklogs"

	raw, err := ioutil.ReadFile("./mocks/tempo_worklogs.json")
	if err != nil {
		t.Error(err.Error())
	}

	testMux.HandleFunc(testAPIEndpoint, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testRequestURL(t, r, testAPIEndpoint)
		fmt.Fprint(w, string(raw))
	})

	dateFrom, _ := time.Parse(WLDateFormat, "2018-04-01")
	dateTo, _ := time.Parse(WLDateFormat, "2018-04-10")
	options := TSWorkLogOptions{
		Username: "smgladkovskiy",
		DateFrom: &WLDate{dateFrom},
		DateTo:   &WLDate{dateTo},
	}
	worklogs, _, err := testClient.Worklog.GetWorkLogs(&options)
	if worklogs == nil {
		t.Error("Expected worklogs. Worklogs is nil")
	}
	if err != nil {
		t.Errorf("Error given: %s", err)
	}
}
