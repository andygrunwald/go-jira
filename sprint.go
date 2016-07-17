package jira

import (
	"fmt"
	"time"
)

// SprintService handles sprints in JIRA Agile API.
type SprintService struct {
	client *Client
}

// Wrapper struct for search result
type sprintsResult struct {
	Sprints []Sprint `json:"values"`
}

// Sprint represents a sprint on JIRA agile board
type Sprint struct {
	ID            int        `json:"id"`
	Name          string     `json:"name"`
	CompleteDate  *time.Time `json:"completeDate"`
	EndDate       *time.Time `json:"endDate"`
	StartDate     *time.Time `json:"startDate"`
	OriginBoardID int        `json:"originBoardId"`
	Self          string     `json:"self"`
	State         string     `json:"state"`
}

// Wrapper struct for moving issues to sprint
type IssuesWrapper struct {
	Issues []string `json:"issues"`
}

// GetList gets sprints for given board
//
// JIRA API docs: https://docs.atlassian.com/jira-software/REST/cloud/#agile/1.0/board/{boardId}/sprint
func (s *SprintService) GetList(boardID string) ([]Sprint, *Response, error) {
	apiEndpoint := fmt.Sprintf("rest/agile/1.0/board/%s/sprint", boardID)
	req, err := s.client.NewRequest("GET", apiEndpoint, nil)
	if err != nil {
		return nil, nil, err
	}

	result := new(sprintsResult)
	resp, err := s.client.Do(req, result)
	return result.Sprints, resp, err
}

func (s *SprintService) AddIssuesToSprint(sprintID int, issueIDs []string) (*Response, error) {
	apiEndpoint := fmt.Sprintf("rest/agile/1.0/sprint/%d/issue", sprintID)

	payload := IssuesWrapper{Issues: issueIDs}

	req, err := s.client.NewRequest("POST", apiEndpoint, payload)

	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req, nil)
	return resp, err
}
