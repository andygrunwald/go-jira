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
	MaxResults int `json:"maxResults"`
	StartAt int `json:"startAt"`
	Total int `json:"total"`
	IsLast bool `json:"isLast"`
	Values []struct {
		ID int `json:"id"`
		Self string `json:"self"`
		Name string `json:"name"`
		Type string `json:"type"`
	} `json:"values"`
}

type BoardListSettings struct {
	startAt int
	maxResults int
	boardType string
	name string
	projectKeyOrId string
}


// Get all boards form jira
//
// JIRA API docs: https://docs.atlassian.com/jira/REST/latest/#api/2/project-getAllProjects
func (s *BoardService) GetList(bs *BoardListSettings) (*BoardsList, *http.Response, error) {


	apiEndpoint := "/rest/agile/1.0/board"
	req, err := s.client.NewRequest("GET", apiEndpoint, nil)

	if bs != nil {
		values := req.URL.Query()
		values.Add(name, value)
	}

	if err != nil {
		return nil, nil, err
	}

	boards := new(BoardsList)
	resp, err := s.client.Do(req, boards)
	if err != nil {
		return nil, resp, err
	}
	return boards, resp, nil
}
