package onpremise

import (
	"context"
	"errors"
	"fmt"
	"net/http"
)

// StatusCategoryService handles status categories for the Jira instance / API.
//
// Use it to obtain a list of all status categories and the details of a category.
// Status categories provided a mechanism for categorizing statuses.
//
// Jira API docs: https://developer.atlassian.com/cloud/jira/platform/rest/v2/api-group-workflow-status-categories/#api-group-workflow-status-categories
type StatusCategoryService service

// StatusCategory represents the category a status belongs to.
// Those categories can be user defined in every Jira instance.
type StatusCategory struct {
	Self      string `json:"self" structs:"self"`
	ID        int    `json:"id" structs:"id"`
	Name      string `json:"name" structs:"name"`
	Key       string `json:"key" structs:"key"`
	ColorName string `json:"colorName" structs:"colorName"`
}

// These constants are the keys of the default Jira status categories
const (
	StatusCategoryComplete   = "done"
	StatusCategoryInProgress = "indeterminate"
	StatusCategoryToDo       = "new"
	StatusCategoryUndefined  = "undefined"
)

// GetList returns a list of all status categories.
//
// Jira API docs: https://docs.atlassian.com/software/jira/docs/api/REST/7.6.1/#api/2/statuscategory-getStatusCategories
func (s *StatusCategoryService) GetList(ctx context.Context) ([]StatusCategory, *Response, error) {
	apiEndpoint := "/rest/api/2/statuscategory"
	req, err := s.client.NewRequest(ctx, http.MethodGet, apiEndpoint, nil)
	if err != nil {
		return nil, nil, err
	}

	var statusCategories []StatusCategory
	resp, err := s.client.Do(req, &statusCategories)
	if err != nil {
		return nil, resp, NewJiraError(resp, err)
	}

	return statusCategories, resp, nil
}

// Get returns a full representation of the StatusCategory having the given id or key.
//
// statusCategoryID represents the ID or key of the status category.
//
// Jira API docs: https://docs.atlassian.com/software/jira/docs/api/REST/7.6.1/#api/2/statuscategory-getStatusCategory
func (s *StatusCategoryService) Get(ctx context.Context, statusCategoryID string) (*StatusCategory, *Response, error) {
	if statusCategoryID == "" {
		return nil, nil, errors.New("no status category id set")
	}

	apiEndpoint := fmt.Sprintf("/rest/api/2/statuscategory/%v", statusCategoryID)
	req, err := s.client.NewRequest(ctx, http.MethodGet, apiEndpoint, nil)
	if err != nil {
		return nil, nil, err
	}

	statusCategory := new(StatusCategory)
	resp, err := s.client.Do(req, statusCategory)
	if err != nil {
		return nil, resp, NewJiraError(resp, err)
	}

	return statusCategory, resp, nil
}
