package jira

import (
	//"fmt"
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
	Values     []Value `json:"values"`
}

type Value struct {
	ID   int    `json:"id"`
	Self string `json:"self"`
	Name string `json:"name"`
	Type string `json:"type"`
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
	// Works only for 1.0 endpoints
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
	apiEndpoint := "/rest/agile/1.0/board"
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
