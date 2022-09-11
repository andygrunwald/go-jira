package cloud

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"testing"
)

func TestFilterService_GetList(t *testing.T) {
	setup()
	defer teardown()
	testAPIEndpoint := "/rest/api/2/filter"
	raw, err := os.ReadFile("../testing/mock-data/all_filters.json")
	if err != nil {
		t.Error(err.Error())
	}
	testMux.HandleFunc(testAPIEndpoint, func(writer http.ResponseWriter, request *http.Request) {
		testMethod(t, request, http.MethodGet)
		testRequestURL(t, request, testAPIEndpoint)
		fmt.Fprint(writer, string(raw))
	})

	filters, _, err := testClient.Filter.GetList(context.Background())
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
	testAPIEndpoint := "/rest/api/2/filter/10000"
	raw, err := os.ReadFile("../testing/mock-data/filter.json")
	if err != nil {
		t.Error(err.Error())
	}
	testMux.HandleFunc(testAPIEndpoint, func(writer http.ResponseWriter, request *http.Request) {
		testMethod(t, request, http.MethodGet)
		testRequestURL(t, request, testAPIEndpoint)
		fmt.Fprint(writer, string(raw))
	})

	filter, _, err := testClient.Filter.Get(context.Background(), 10000)
	if filter == nil {
		t.Errorf("Expected Filter, got nil")
	}
	if err != nil {
		t.Errorf("Error given: %s", err)
	}

}

func TestFilterService_GetFavouriteList(t *testing.T) {
	setup()
	defer teardown()
	testAPIEndpoint := "/rest/api/2/filter/favourite"
	raw, err := os.ReadFile("../testing/mock-data/favourite_filters.json")
	if err != nil {
		t.Error(err.Error())
	}
	testMux.HandleFunc(testAPIEndpoint, func(writer http.ResponseWriter, request *http.Request) {
		testMethod(t, request, http.MethodGet)
		testRequestURL(t, request, testAPIEndpoint)
		fmt.Fprint(writer, string(raw))
	})

	filters, _, err := testClient.Filter.GetFavouriteList(context.Background())
	if filters == nil {
		t.Error("Expected Filters list. Filters list is nil")
	}
	if err != nil {
		t.Errorf("Error given: %s", err)
	}
}

func TestFilterService_GetMyFilters(t *testing.T) {
	setup()
	defer teardown()
	testAPIEndpoint := "/rest/api/3/filter/my"
	raw, err := os.ReadFile("../testing/mock-data/my_filters.json")
	if err != nil {
		t.Error(err.Error())
	}
	testMux.HandleFunc(testAPIEndpoint, func(writer http.ResponseWriter, request *http.Request) {
		testMethod(t, request, http.MethodGet)
		testRequestURL(t, request, testAPIEndpoint)
		fmt.Fprint(writer, string(raw))
	})

	opts := GetMyFiltersQueryOptions{}
	filters, _, err := testClient.Filter.GetMyFilters(context.Background(), &opts)
	if err != nil {
		t.Errorf("Error given: %s", err)
	}
	if filters == nil {
		t.Errorf("Expected Filters, got nil")
	}
}

func TestFilterService_Search(t *testing.T) {
	setup()
	defer teardown()
	testAPIEndpoint := "/rest/api/3/filter/search"
	raw, err := os.ReadFile("../testing/mock-data/search_filters.json")
	if err != nil {
		t.Error(err.Error())
	}
	testMux.HandleFunc(testAPIEndpoint, func(writer http.ResponseWriter, request *http.Request) {
		testMethod(t, request, http.MethodGet)
		testRequestURL(t, request, testAPIEndpoint)
		fmt.Fprint(writer, string(raw))
	})

	opt := FilterSearchOptions{}
	filters, _, err := testClient.Filter.Search(context.Background(), &opt)
	if err != nil {
		t.Errorf("Error given: %s", err)
	}
	if filters == nil {
		t.Errorf("Expected Filters, got nil")
	}
}
