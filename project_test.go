package jira

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestProjectGetAll(t *testing.T) {
	setup()
	defer teardown()
	testApiEdpoint := "/rest/api/2/project"

	raw, err := ioutil.ReadFile("./json_mocks/all_projects.json")
	if err != nil {
		t.Error(err.Error())
	}
	testMux.HandleFunc(testApiEdpoint, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testRequestURL(t, r, testApiEdpoint)
		fmt.Fprint(w, string(raw))
	})

	projects, _, err := testClient.Project.GetList()
	if projects == nil {
		t.Error("Expected project list. Project list is nil")
	}
	if err != nil {
		t.Errorf("Error given: %s", err)
	}
}

func TestProjectGet(t *testing.T) {
	setup()
	defer teardown()
	testApiEdpoint := "/rest/api/2/project/12310505"

	raw, err := ioutil.ReadFile("./json_mocks/project.json")
	if err != nil {
		t.Error(err.Error())
	}
	testMux.HandleFunc(testApiEdpoint, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testRequestURL(t, r, testApiEdpoint)
		fmt.Fprint(w, string(raw))
	})

	projects, _, err := testClient.Project.Get("12310505")
	if projects == nil {
		t.Error("Expected project list. Project list is nil")
	}
	if err != nil {
		t.Errorf("Error given: %s", err)
	}
}
