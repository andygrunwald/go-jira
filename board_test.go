package jira

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestBoardsGetAll(t *testing.T) {
	setup()
	defer teardown()
	testAPIEdpoint := "/rest/agile/1.0/board"

	raw, err := ioutil.ReadFile("./mocks/all_boards.json")
	if err != nil {
		t.Error(err.Error())
	}
	testMux.HandleFunc(testAPIEdpoint, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testRequestURL(t, r, testAPIEdpoint)
		fmt.Fprint(w, string(raw))
	})

	projects, _, err := testClient.Board.GetList(nil)
	if projects == nil {
		t.Error("Expected boards list. Boards list is nil")
	}
	if err != nil {
		t.Errorf("Error given: %s", err)
	}
}

// Test with params
func TestBoardsGetFiltered(t *testing.T) {
	setup()
	defer teardown()
	testAPIEdpoint := "/rest/agile/1.0/board"

	raw, err := ioutil.ReadFile("./mocks/all_boards_filtered.json")
	if err != nil {
		t.Error(err.Error())
	}
	testMux.HandleFunc(testAPIEdpoint, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testRequestURL(t, r, testAPIEdpoint)
		fmt.Fprint(w, string(raw))
	})

	boardsListOptions := BoardListOptions{
		BoardType:      "scrum",
		Name:           "Test",
		ProjectKeyOrId: "TE",
		StartAt:        1,
		MaxResults:     10,
	}

	projects, _, err := testClient.Board.GetList(&boardsListOptions)
	if projects == nil {
		t.Error("Expected boards list. Boards list is nil")
	}
	if err != nil {
		t.Errorf("Error given: %s", err)
	}
}
