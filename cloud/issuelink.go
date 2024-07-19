package cloud

import (
	"context"
	"fmt"
	"net/http"
)

// IssueLinkService handles issue links for the Jira instance / API.
//
// Jira API docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-issue-links/#api-group-issue-links
type IssueLinkService service

type CreateLink struct {
	Type         IssueLinkType `json:"type"`
	InwardIssue  LinkedIssue   `json:"inwardIssue"`
	OutwardIssue LinkedIssue   `json:"outwardIssue"`
	Comment      *Commentv3    `json:"comment,omitempty"`
}

type LinkedIssue struct {
	ID  string `json:"id"`
	Key string `json:"key"`
}

// TODO rename it once the current Comment struct is named correctly
type Commentv3 struct {
	Body       string           `json:"body"`
	Visibility Visibility       `json:"visibility"`
	Properties []EntityProperty `json:"properties"`
	// Additional Properties
}

type Visibility struct {
	Type       string `json:"type"`
	Value      string `json:"value"`
	Identifier string `json:"identifier,omitempty"`
	// Additional Properties
}

// IssueLink represents a link between two issues in Jira.
// TODO rename it once the current IssueLink struct is named correctly
type IssueLinkv3 struct {
	ID           string        `json:"id"`
	Self         string        `json:"self"`
	Type         IssueLinkType `json:"type"`
	OutwardIssue LinkedIssue   `json:"outwardIssue"`
	InwardIssue  LinkedIssue   `json:"inwardIssue"`
}

// Create creates an issue link.
//
// Jira API docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-issue-links/#api-rest-api-3-issuelink-post
func (s *IssueLinkService) Create(ctx context.Context, body CreateLink) (*Response, error) {
	apiEndpoint := "rest/api/3/issueLink"
	req, err := s.client.NewRequest(ctx, http.MethodPost, apiEndpoint, body)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req, nil)
	if err != nil {
		return resp, NewJiraError(resp, err)
	}
	defer resp.Body.Close()

	return resp, nil
}

// Get get an issue link.
//
// Jira API docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-issue-links/#api-rest-api-3-issuelink-linkid-get
func (s *IssueLinkService) Get(ctx context.Context, linkID string) (*IssueLinkv3, *Response, error) {
	apiEndpoint := fmt.Sprintf("rest/api/3/issueLink/%s", linkID)
	req, err := s.client.NewRequest(ctx, http.MethodGet, apiEndpoint, nil)
	if err != nil {
		return nil, nil, err
	}

	responseCustomer := new(IssueLinkv3)
	resp, err := s.client.Do(req, &responseCustomer)
	if err != nil {
		return nil, resp, NewJiraError(resp, err)
	}

	return responseCustomer, resp, nil
}

// Delete deletes an issue link.
//
// Jira API docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-issue-links/#api-rest-api-3-issuelink-linkid-delete
func (s *IssueLinkService) Delete(ctx context.Context, linkID string) (*Response, error) {
	apiEndpoint := fmt.Sprintf("rest/api/3/issueLink/%s", linkID)
	req, err := s.client.NewRequest(ctx, http.MethodDelete, apiEndpoint, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req, nil)
	if err != nil {
		return resp, NewJiraError(resp, err)
	}

	return resp, nil
}
