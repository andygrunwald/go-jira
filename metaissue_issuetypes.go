package jira

import (
	"context"
	"errors"
	"fmt"
)

type GetIssueTypesResponse struct {
	Values []*MetaIssueType `json:"values,omitempty"`
	Total  int              `json:"total,omitempty"`
}

// GetIssueTypesForProject returns the issue types for a given project, and is designed to replace the
// /rest/api/2/issue/createmeta endpoint, which is called using GetCreateMetaWithOptionsWithContext. That endpoint
// is deprecated as of Jira Server 9.0.0.
// See: https://docs.atlassian.com/software/jira/docs/api/REST/9.0.0/#issue-getCreateIssueMetaProjectIssueTypes
func (s *IssueService) GetIssueTypesForProject(ctx context.Context, projectKey string) (*GetIssueTypesResponse, *Response, error) {
	if projectKey == "" {
		return nil, nil, errors.New("project key cannot be empty")
	}

	apiEndpoint := fmt.Sprintf("rest/api/2/issue/createmeta/%s/issuetypes", projectKey)

	req, err := s.client.NewRequestWithContext(ctx, "GET", apiEndpoint, nil)
	if err != nil {
		return nil, nil, err
	}

	meta := new(GetIssueTypesResponse)
	resp, err := s.client.Do(req, meta)

	if err != nil {
		return nil, resp, err
	}

	return meta, resp, nil
}

type FieldMetaSchema struct {
	Type   string
	Items  string
	Custom string
	System string
}

type FieldMetaAllowedValue struct {
	ID      string
	Name    string
	Value   string
	IconURL string
}

type FieldMeta struct {
	FieldID         string                   `json:"fieldId"`
	Name            string                   `json:"name"`
	Required        bool                     `json:"required"`
	AutoCompleteURL string                   `json:"autoCompleteUrl"`
	Schema          FieldMetaSchema          `json:"schema"`
	HasDefaultValue bool                     `json:"hasDefaultValue"`
	AllowedValues   []*FieldMetaAllowedValue `json:"allowedValues"`
	Operations      []string                 `json:"operations"`
}

type GetFieldsResponse struct {
	Values []*FieldMeta `json:"values,omitempty"`
	Total  int          `json:"total,omitempty"`
}

func (s *IssueService) GetFieldsForIssueType(ctx context.Context, projectKey string, issueTypeID string) (*GetFieldsResponse, *Response, error) {
	if projectKey == "" {
		return nil, nil, errors.New("project key cannot be empty")
	}
	if issueTypeID == "" {
		return nil, nil, errors.New("Issue Type ID cannot be empty")
	}

	apiEndpoint := fmt.Sprintf("rest/api/2/issue/createmeta/%s/issuetypes/%s", projectKey, issueTypeID)

	req, err := s.client.NewRequestWithContext(ctx, "GET", apiEndpoint, nil)
	if err != nil {
		return nil, nil, err
	}

	meta := new(GetFieldsResponse)
	resp, err := s.client.Do(req, meta)

	if err != nil {
		return nil, resp, err
	}

	return meta, resp, nil
}
