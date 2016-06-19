package jira

import "fmt"

// BoardService handles Agile Boards for the JIRA instance / API.
//
// JIRA API docs: https://docs.atlassian.com/jira-software/REST/server/
type BoardService struct {
	client *Client
}

// BoardsList reflects a list of agile boards
type BoardsList struct {
	MaxResults int     `json:"maxResults"`
	StartAt    int     `json:"startAt"`
	Total      int     `json:"total"`
	IsLast     bool    `json:"isLast"`
	Values     []Board `json:"values"`
}

// Board represents a JIRA agile board
type Board struct {
	ID       int    `json:"id,omitempty"`
	Self     string `json:"self,omitempty"`
	Name     string `json:"name,omitempty"`
	Type     string `json:"type,omitempty"`
	FilterID int    `json:"filterId,omitempty"`
}

// BoardListOptions specifies the optional parameters to the BoardService.GetList
type BoardListOptions struct {
	// BoardType filters results to boards of the specified type.
	// Valid values: scrum, kanban.
	BoardType string `url:"boardType,omitempty"`
	// Name filters results to boards that match or partially match the specified name.
	Name string `url:"name,omitempty"`
	// ProjectKeyOrID filters results to boards that are relevant to a project.
	// Relevance meaning that the JQL filter defined in board contains a reference to a project.
	ProjectKeyOrID string `url:"projectKeyOrId,omitempty"`

	SearchOptions
}

// GetList will return all boards from JIRA
//
// JIRA API docs: https://docs.atlassian.com/jira-software/REST/cloud/#agile/1.0/board-getAllBoards
func (s *BoardService) GetList(opt *BoardListOptions) (*BoardsList, *Response, error) {
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

// Get will return the board for the given boardID.
// This board will only be returned if the user has permission to view it.
//
// JIRA API docs: https://docs.atlassian.com/jira-software/REST/cloud/#agile/1.0/board-getBoard
func (s *BoardService) Get(boardID int) (*Board, *Response, error) {
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

// Create creates a new board. Board name, type and filter Id is required.
// name - Must be less than 255 characters.
// type - Valid values: scrum, kanban
// filterId - Id of a filter that the user has permissions to view.
// Note, if the user does not have the 'Create shared objects' permission and tries to create a shared board, a private
// board will be created instead (remember that board sharing depends on the filter sharing).
//
// JIRA API docs: https://docs.atlassian.com/jira-software/REST/cloud/#agile/1.0/board-createBoard
func (s *BoardService) Create(board *Board) (*Board, *Response, error) {
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

// Delete will delete an agile board.
//
// https://docs.atlassian.com/jira-software/REST/cloud/#agile/1.0/board-deleteBoard
func (s *BoardService) Delete(boardID int) (*Board, *Response, error) {
	apiEndpoint := fmt.Sprintf("rest/agile/1.0/board/%v", boardID)
	req, err := s.client.NewRequest("DELETE", apiEndpoint, nil)
	if err != nil {
		return nil, nil, err
	}

	resp, err := s.client.Do(req, nil)
	return nil, resp, err
}
