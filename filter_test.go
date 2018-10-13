package jira

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestFilterService_GetList(t *testing.T) {
	setup()
	defer teardown()
	testAPIEdpoint := "/rest/api/2/filter"
	raw, err := ioutil.ReadFile("./mocks/all_filters.json")
	if err != nil {
		t.Error(err.Error())
	}
	testMux.HandleFunc(testAPIEdpoint, func(writer http.ResponseWriter, request *http.Request) {
		testMethod(t, request, "GET")
		testRequestURL(t, request, testAPIEdpoint)
		fmt.Fprint(writer, string(raw))
	})

	filters, _, err := testClient.Filter.GetList()
	if filters == nil {
		t.Error("Expected Filters list. Filters list is nil")
	}
	if err != nil {
		t.Errorf("Error given: %s", err)
	}
}

func TestFilterService_Get(t *testing.T) {
	setup()
	defer teardown()
	testAPIEdpoint := "/rest/api/2/filter/10000"
	raw, err := ioutil.ReadFile("./mocks/filter.json")
	if err != nil {
		t.Error(err.Error())
	}
	testMux.HandleFunc(testAPIEdpoint, func(writer http.ResponseWriter, request *http.Request) {
		testMethod(t, request, "GET")
		testRequestURL(t, request, testAPIEdpoint)
		fmt.Fprintf(writer, string(raw))
	})

	filter, _, err := testClient.Filter.Get(10000)
	if filter == nil {
		t.Errorf("Expected Filter, got nil")
	}
	if err != nil {
		t.Errorf("Error given: %s", err)
	}


}
