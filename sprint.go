package jira

import (
	"context"
	"fmt"

	"github.com/google/go-querystring/query"
)

// SprintService handles sprints in Jira Agile API.
// See https://docs.atlassian.com/jira-software/REST/cloud/
type SprintService struct {
	client *Client
}

// IssuesWrapper represents a wrapper struct for moving issues to sprint
type IssuesWrapper struct {
	Issues []string `json:"issues"`
}

// IssuesInSprintResult represents a wrapper struct for search result
type IssuesInSprintResult struct {
	Issues []Issue `json:"issues"`
}

type GetIssuesForSprintOptions struct {
	// Jql filters results to issues that match the given JQL.
	Jql string `url:"jql,omitempty"`

	SearchOptions
}

// MoveIssuesToSprintWithContext moves issues to a sprint, for a given sprint Id.
// Issues can only be moved to open or active sprints.
// The maximum number of issues that can be moved in one operation is 50.
//
// Jira API docs: https://docs.atlassian.com/jira-software/REST/cloud/#agile/1.0/sprint-moveIssuesToSprint
// Caller must close resp.Body
func (s *SprintService) MoveIssuesToSprintWithContext(ctx context.Context, sprintID int, issueIDs []string) (*Response, error) {
	apiEndpoint := fmt.Sprintf("rest/agile/1.0/sprint/%d/issue", sprintID)

	payload := IssuesWrapper{Issues: issueIDs}

	req, err := s.client.NewRequestWithContext(ctx, "POST", apiEndpoint, payload)

	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req, nil)
	if err != nil {
		err = NewJiraError(resp, err)
	}
	return resp, err
}

// MoveIssuesToSprint wraps MoveIssuesToSprintWithContext using the background context.
// Caller must close resp.Body
func (s *SprintService) MoveIssuesToSprint(sprintID int, issueIDs []string) (*Response, error) {
	return s.MoveIssuesToSprintWithContext(context.Background(), sprintID, issueIDs)
}

// GetIssuesForSprintWithContext returns all issues in a sprint, for a given sprint Id.
// This only includes issues that the user has permission to view.
// By default, the returned issues are ordered by rank.
//
// Jira API Docs: https://docs.atlassian.com/jira-software/REST/cloud/#agile/1.0/sprint-getIssuesForSprint
func (s *SprintService) GetIssuesForSprintWithContext(ctx context.Context, sprintID int) ([]Issue, *Response, error) {
	apiEndpoint := fmt.Sprintf("rest/agile/1.0/sprint/%d/issue", sprintID)

	req, err := s.client.NewRequestWithContext(ctx, "GET", apiEndpoint, nil)

	if err != nil {
		return nil, nil, err
	}

	result := new(IssuesInSprintResult)
	resp, err := s.client.Do(req, result)
	if err != nil {
		err = NewJiraError(resp, err)
	}

	return result.Issues, resp, err
}

// GetIssuesForSprint wraps GetIssuesForSprintWithContext using the background context.
func (s *SprintService) GetIssuesForSprint(sprintID int) ([]Issue, *Response, error) {
	return s.GetIssuesForSprintWithContext(context.Background(), sprintID)
}

// GetIssuesForSprintWithOptionsWithContext returns all issues in a sprint, for a given sprint Id.
// The issues are filtered out with the given options.
// This only includes issues that the user has permission to view.
// By default, the returned issues are ordered by rank.
//
// Jira API Docs: https://docs.atlassian.com/jira-software/REST/cloud/#agile/1.0/sprint-getIssuesForSprint
func (s *SprintService) GetIssuesForSprintWithOptionsWithContext(ctx context.Context, sprintID int, options *GetIssuesForSprintOptions) ([]Issue, *Response, error) {
	apiEndpoint := fmt.Sprintf("rest/agile/1.0/sprint/%d/issue", sprintID)
	url, err := addOptions(apiEndpoint, options)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequestWithContext(ctx, "GET", url, nil)

	if err != nil {
		return nil, nil, err
	}

	result := new(IssuesInSprintResult)
	resp, err := s.client.Do(req, result)
	if err != nil {
		err = NewJiraError(resp, err)
	}

	return result.Issues, resp, err
}

// GetIssuesForSprintWithOptions wraps GetIssuesForSprintWithOptionsWithContext using the background context.
func (s *SprintService) GetIssuesForSprintWithOptions(sprintID int, options *GetIssuesForSprintOptions) ([]Issue, *Response, error) {
	return s.GetIssuesForSprintWithOptionsWithContext(context.Background(), sprintID, options)
}

// GetIssueWithContext returns a full representation of the issue for the given issue key.
// Jira will attempt to identify the issue by the issueIdOrKey path parameter.
// This can be an issue id, or an issue key.
// If the issue cannot be found via an exact match, Jira will also look for the issue in a case-insensitive way, or by looking to see if the issue was moved.
//
// # The given options will be appended to the query string
//
// Jira API docs: https://docs.atlassian.com/jira-software/REST/7.3.1/#agile/1.0/issue-getIssue
//
// TODO: create agile service for holding all agile apis' implementation
func (s *SprintService) GetIssueWithContext(ctx context.Context, issueID string, options *GetQueryOptions) (*Issue, *Response, error) {
	apiEndpoint := fmt.Sprintf("rest/agile/1.0/issue/%s", issueID)

	req, err := s.client.NewRequestWithContext(ctx, "GET", apiEndpoint, nil)

	if err != nil {
		return nil, nil, err
	}

	if options != nil {
		q, err := query.Values(options)
		if err != nil {
			return nil, nil, err
		}
		req.URL.RawQuery = q.Encode()
	}

	issue := new(Issue)
	resp, err := s.client.Do(req, issue)

	if err != nil {
		jerr := NewJiraError(resp, err)
		return nil, resp, jerr
	}

	return issue, resp, nil
}

// GetIssue wraps GetIssueWithContext using the background context.
func (s *SprintService) GetIssue(issueID string, options *GetQueryOptions) (*Issue, *Response, error) {
	return s.GetIssueWithContext(context.Background(), issueID, options)
}

// CreateWithContext creates a sprint from a JSON representation.
//
// Jira API docs: https://docs.atlassian.com/jira-software/REST/7.3.1/#agile/1.0/sprint-createSprint
func (s *SprintService) CreateWithContext(ctx context.Context, sprint *Sprint) (*Sprint, *Response, error) {
	apiEndpoint := "rest/agile/1.0/sprint"
	req, err := s.client.NewRequestWithContext(ctx, "POST", apiEndpoint, sprint)
	if err != nil {
		return nil, nil, err
	}

	created := new(Sprint)
	resp, err := s.client.Do(req, created)
	if err != nil {
		jerr := NewJiraError(resp, err)
		return nil, resp, jerr
	}

	return created, resp, nil
}

// Create wraps CreateWithContext using the background context.
func (s *SprintService) Create(sprint *Sprint) (*Sprint, *Response, error) {
	return s.CreateWithContext(context.Background(), sprint)
}

// UpdateSprintWithContext partially updates a sprint from a JSON representation. The sprint is found by ID.
//
// Jira API docs: https://docs.atlassian.com/jira-software/REST/7.3.1/#agile/1.0/sprint-partiallyUpdateSprint
// Caller must close resp.Body
func (s *SprintService) UpdateSprintWithContext(ctx context.Context, sprintID int, data map[string]interface{}) (*Response, error) {
	apiEndpoint := fmt.Sprintf("rest/agile/1.0/sprint/%v", sprintID)

	req, err := s.client.NewRequestWithContext(ctx, "POST", apiEndpoint, data)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req, nil)
	if err != nil {
		jerr := NewJiraError(resp, err)
		return resp, jerr
	}

	return resp, nil
}

// UpdateSprint wraps UpdateSprintWithContext using the background context.
// Caller must close resp.Body
func (s *SprintService) UpdateSprint(sprintID int, data map[string]interface{}) (*Response, error) {
	return s.UpdateSprintWithContext(context.Background(), sprintID, data)
}
