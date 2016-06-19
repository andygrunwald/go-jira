package jira

import (
	"fmt"
	"net/http"
)

type BoardService struct {
	client *Client
}

//Type for boards list
type BoardsList struct {
	MaxResults int     `json:"maxResults"`
	StartAt    int     `json:"startAt"`
	Total      int     `json:"total"`
	IsLast     bool    `json:"isLast"`
	Values     []Board `json:"values"`
}

// Board represents a JIRA board
type Board struct {
	ID       int    `json:"id",omitempty"`
	Self     string `json:"self",omitempty"`
	Name     string `json:"name",omitempty"`
	Type     string `json:"type",omitempty"`
	FilterId int    `omitempty`
}

// BoardListOptions specifies the optional parameters to the BoardService.GetList
type BoardListOptions struct {
	// Filters results to boards of the specified type.
	// Valid values: scrum, kanban.
	BoardType string `url:"boardType,omitempty"`
	// Filters results to boards that match or partially match the specified name.
	Name string `url:"name,omitempty"`
	// Filters results to boards that are relevant to a project.
	// Relevance meaning that the jql filter defined in board contains a reference to a project.
	ProjectKeyOrId string `url:"projectKeyOrId,omitempty"`
	// ListOptions specifies the optional parameters to various List methods that
	// support pagination.
	// Pagination is used for the JIRA REST APIs to conserve server resources and limit
	// response size for resources that return potentially large collection of items.
	// A request to a pages API will result in a values array wrapped in a JSON object with some paging metadata
	// Default Pagination options
	// The starting index of the returned projects. Base index: 0.
	StartAt int `url:"startAt,omitempty"`
	// The maximum number of projects to return per page. Default: 50.
	MaxResults int `url:"maxResults,omitempty"`
}

// Get all boards form jira
//
// JIRA API docs: https://docs.atlassian.com/jira/REST/latest/#api/2/project-getAllProjects
func (s *BoardService) GetList(opt *BoardListOptions) (*BoardsList, *http.Response, error) {
	apiEndpoint := "rest/agile/1.0/board"
	url, err := addOptions(apiEndpoint, opt)
	req, err := s.client.NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}

	boards := new(BoardsList)
	resp, err := s.client.Do(req, boards)
	if err != nil {
		return nil, resp, err
	}

	return boards, resp, err
}

// Returns the board for the given board Id. This board will only be returned if the user has permission to view it.
func (s *BoardService) Get(boardID int) (*Board, *http.Response, error) {
	apiEndpoint := fmt.Sprintf("rest/agile/1.0/board/%v", boardID)
	req, err := s.client.NewRequest("GET", apiEndpoint, nil)
	if err != nil {
		return nil, nil, err
	}

	board := new(Board)
	resp, err := s.client.Do(req, board)
	if err != nil {
		return nil, resp, err
	}
	return board, resp, nil
}

// Creates a new board. Board name, type and filter Id is required.
// name - Must be less than 255 characters.
// type - Valid values: scrum, kanban
// filterId - Id of a filter that the user has permissions to view.
// Note, if the user does not have the 'Create shared objects' permission and tries to create a shared board, a private
// board will be created instead (remember that board sharing depends on the filter sharing).
//
// JIRA API docs: https://docs.atlassian.com/jira-software/REST/cloud/#agile/1.0/board-createBoard
func (s *BoardService) Create(board *Board) (*Board, *http.Response, error) {

	apiEndpoint := "rest/agile/1.0/board"
	req, err := s.client.NewRequest("POST", apiEndpoint, board)
	if err != nil {
		return nil, nil, err
	}

	responseBoard := new(Board)
	resp, err := s.client.Do(req, responseBoard)
	if err != nil {
		return nil, resp, err
	}
	return responseBoard, resp, nil
}

// Deletes the board.
//
// https://docs.atlassian.com/jira-software/REST/cloud/#agile/1.0/board-deleteBoard
func (s *BoardService) Delete(boardID int) (*Board, *http.Response, error) {
	apiEndpoint := fmt.Sprintf("rest/agile/1.0/board/%v", boardID)
	req, err := s.client.NewRequest("DELETE", apiEndpoint, nil)
	if err != nil {
		return nil, nil, err
	}

	resp, err := s.client.Do(req, nil)
	return nil, resp, err
}
