package onpremise

import (
	"context"
	"net/http"
)

// ResolutionService handles resolutions for the Jira instance / API.
//
// Jira API docs: https://developer.atlassian.com/cloud/jira/platform/rest/#api-Resolution
type ResolutionService service

// Resolution represents a resolution of a Jira issue.
// Typical types are "Fixed", "Suspended", "Won't Fix", ...
type Resolution struct {
	Self        string `json:"self" structs:"self"`
	ID          string `json:"id" structs:"id"`
	Description string `json:"description" structs:"description"`
	Name        string `json:"name" structs:"name"`
}

// GetList gets all resolutions from Jira
//
// Jira API docs: https://developer.atlassian.com/cloud/jira/platform/rest/#api-api-2-resolution-get
//
// TODO Double check this method if this works as expected, is using the latest API and the response is complete
// This double check effort is done for v2 - Remove this two lines if this is completed.
func (s *ResolutionService) GetList(ctx context.Context) ([]Resolution, *Response, error) {
	apiEndpoint := "rest/api/2/resolution"
	req, err := s.client.NewRequest(ctx, http.MethodGet, apiEndpoint, nil)
	if err != nil {
		return nil, nil, err
	}

	resolutionList := []Resolution{}
	resp, err := s.client.Do(req, &resolutionList)
	if err != nil {
		return nil, resp, NewJiraError(resp, err)
	}
	return resolutionList, resp, nil
}
