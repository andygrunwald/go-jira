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

	boardsListOptions := &BoardListOptions{
		BoardType:      "scrum",
		Name:           "Test",
		ProjectKeyOrId: "TE",
		StartAt:        1,
		MaxResults:     10,
	}

	projects, _, err := testClient.Board.GetList(boardsListOptions)
	if projects == nil {
		t.Error("Expected boards list. Boards list is nil")
	}
	if err != nil {
		t.Errorf("Error given: %s", err)
	}
}

func TestBoardGet(t *testing.T) {
	setup()
	defer teardown()
	testAPIEdpoint := "/rest/agile/1.0/board/1"

	testMux.HandleFunc(testAPIEdpoint, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testRequestURL(t, r, testAPIEdpoint)
		fmt.Fprint(w, `{"id":4,"self":"https://test.jira.org/rest/agile/1.0/board/1","name":"Test Weekly","type":"scrum"}`)
	})

	board, _, err := testClient.Board.Get("1")
	if board == nil {
		t.Error("Expected board list. Board list is nil")
	}
	if err != nil {
		t.Errorf("Error given: %s", err)
	}
}

func TestBoardGet_NoBoard(t *testing.T) {
	setup()
	defer teardown()
	testAPIEndpoint := "/rest/api/2/board/99999999"

	testMux.HandleFunc(testAPIEndpoint, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testRequestURL(t, r, testAPIEndpoint)
		fmt.Fprint(w, nil)
	})

	board, resp, err := testClient.Board.Get("99999999")
	if board != nil {
		t.Errorf("Expected nil. Got %s", err)
	}

	if resp.Status == "404" {
		t.Errorf("Expected status 404. Got %s", resp.Status)
	}
	if err == nil {
		t.Errorf("Error given: %s", err)
	}
}

func TestBoardCreate(t *testing.T) {
	setup()
	defer teardown()
	testMux.HandleFunc("/rest/agile/1.0/board", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		testRequestURL(t, r, "/rest/agile/1.0/board")

		w.WriteHeader(http.StatusCreated)
		fmt.Fprint(w, `{"id":17,"self":"https://test.jira.org/rest/agile/1.0/board/17","name":"Test","type":"kanban"}`)
	})

	b := &Board{
		Name:     "Test",
		Type:     "kanban",
		FilterId: 17,
	}
	issue, _, err := testClient.Board.Create(b)
	if issue == nil {
		t.Error("Expected board. Board is nil")
	}
	if err != nil {
		t.Errorf("Error given: %s", err)
	}
}
