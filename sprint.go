package jira

import (
	"fmt"
)

// SprintService handles sprints in JIRA Agile API.
// See https://docs.atlassian.com/jira-software/REST/cloud/
type SprintService struct {
	client *Client
}

// IssuesWrapper represents a wrapper struct for moving issues to sprint
type IssuesWrapper struct {
	Issues []string `json:"issues"`
}

// Wrapper struct for search result
type IssuesInSprintResult struct {
	Issues []Issue `json:"issues"`
}

// MoveIssuesToSprint moves issues to a sprint, for a given sprint Id.
// Issues can only be moved to open or active sprints.
// The maximum number of issues that can be moved in one operation is 50.
//
// JIRA API docs: https://docs.atlassian.com/jira-software/REST/cloud/#agile/1.0/sprint-moveIssuesToSprint
func (s *SprintService) MoveIssuesToSprint(sprintID int, issueIDs []string) (*Response, error) {
	apiEndpoint := fmt.Sprintf("rest/agile/1.0/sprint/%d/issue", sprintID)

	payload := IssuesWrapper{Issues: issueIDs}

	req, err := s.client.NewRequest("POST", apiEndpoint, payload)

	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req, nil)
	return resp, err
}

// For a given sprint Id, return all issues currently associated with that sprint
//
//  JIRA API Docs: https://docs.atlassian.com/jira-software/REST/cloud/#agile/1.0/board/{boardId}/sprint-getIssuesForSprint
func (s *SprintService) GetIssuesInSprint(sprintID int) ([]Issue, *Response, error) {
	apiEndpoint := fmt.Sprintf("rest/agile/1.0/sprint/%d/issue", sprintID)

	req, err := s.client.NewRequest("GET", apiEndpoint, nil)

	if err != nil {
		return nil, nil, err
	}

	result := new(IssuesInSprintResult)
	resp, err := s.client.Do(req, result)
	return result.Issues, resp, err
}
