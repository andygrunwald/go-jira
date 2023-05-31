package cloud

import (
	"context"
	"net/http"
)

// StatusService handles staties for the Jira instance / API.
//
// Jira API docs: https://developer.atlassian.com/cloud/jira/platform/rest/v2/#api-group-Workflow-statuses
type StatusService service

// Status represents the current status of a Jira issue.
// Typical status are "Open", "In Progress", "Closed", ...
// Status can be user defined in every Jira instance.
type Status struct {
	Self           string         `json:"self" structs:"self"`
	Description    string         `json:"description" structs:"description"`
	IconURL        string         `json:"iconUrl" structs:"iconUrl"`
	Name           string         `json:"name" structs:"name"`
	ID             string         `json:"id" structs:"id"`
	StatusCategory StatusCategory `json:"statusCategory" structs:"statusCategory"`
}

// GetAllStatuses returns a list of all statuses associated with workflows.
//
// Jira API docs: https://developer.atlassian.com/cloud/jira/platform/rest/v2/#api-rest-api-2-status-get
//
// TODO Double check this method if this works as expected, is using the latest API and the response is complete
// This double check effort is done for v2 - Remove this two lines if this is completed.
func (s *StatusService) GetAllStatuses(ctx context.Context) ([]Status, *Response, error) {
	apiEndpoint := "rest/api/2/status"
	req, err := s.client.NewRequest(ctx, http.MethodGet, apiEndpoint, nil)

	if err != nil {
		return nil, nil, err
	}

	statusList := []Status{}
	resp, err := s.client.Do(req, &statusList)
	if err != nil {
		return nil, resp, NewJiraError(resp, err)
	}

	return statusList, resp, nil
}
