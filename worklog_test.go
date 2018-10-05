package jira

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"testing"
	"time"
)

func TestWorklogService_GetWorkLogs(t *testing.T) {
	setup()
	defer teardown()

	testAPIEndpoint := "/rest/tempo-timesheets/3/worklogs"

	raw, err := ioutil.ReadFile("./mocks/tempo_timesheets_worklogs.json")
	if err != nil {
		t.Error(err.Error())
	}

	testMux.HandleFunc(testAPIEndpoint, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testRequestURL(t, r, testAPIEndpoint)
		fmt.Fprint(w, string(raw))
	})

	params := url.Values{}
	dateFrom, _ := time.Parse(TTWorklogDateFormat, "2018-04-01")
	dateTo, _ := time.Parse(TTWorklogDateFormat, "2018-04-10")
	params.Add("dateFrom", dateFrom.Format(TTWorklogDateFormat))
	params.Add("dateTo", dateTo.Format(TTWorklogDateFormat))
	params.Add("username", "smgladkovskiy")

	worklogs, _, err := testClient.Worklog.GetWorkLogs(&params)
	if worklogs == nil {
		t.Error("Expected worklogs. Worklogs is nil")
	}
	if err != nil {
		t.Errorf("Error given: %s", err)
	}
}
