package cloud

import (
	"context"
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
	Comment      Commentv3     `json:"comment"`
}

type LinkedIssue struct {
	ID  string `json:"id"`
	Key string `json:"key"`
}

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

// CreateLink creates an issue link.
//
// Jira API docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-issue-links/#api-rest-api-3-issuelink-post
func (s *IssueLinkService) CreateLink(ctx context.Context, body CreateLink) (*Response, error) {
	apiEndpoint := "rest/api/3/issueLink"
	req, err := s.client.NewRequest(ctx, http.MethodPost, apiEndpoint, body)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req, nil)
	if err != nil {
		return resp, err
	}
	defer resp.Body.Close()

	return resp, nil
}
