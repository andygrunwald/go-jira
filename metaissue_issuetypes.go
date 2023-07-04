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
