package cloud

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// IssueLinkTypeService handles issue link types for the Jira instance / API.
//
// Jira API docs: https://developer.atlassian.com/cloud/jira/platform/rest/v2/#api-group-Issue-link-types
type IssueLinkTypeService service

// GetList gets all of the issue link types from Jira.
//
// Jira API docs: https://developer.atlassian.com/cloud/jira/platform/rest/v2/#api-rest-api-2-issueLinkType-get
func (s *IssueLinkTypeService) GetList(ctx context.Context) ([]IssueLinkType, *Response, error) {
	apiEndpoint := "rest/api/2/issueLinkType"
	req, err := s.client.NewRequest(ctx, http.MethodGet, apiEndpoint, nil)
	if err != nil {
		return nil, nil, err
	}

	linkTypeList := []IssueLinkType{}
	resp, err := s.client.Do(req, &linkTypeList)
	if err != nil {
		return nil, resp, NewJiraError(resp, err)
	}
	return linkTypeList, resp, nil
}

// Get gets info of a specific issue link type from Jira.
//
// Jira API docs: https://developer.atlassian.com/cloud/jira/platform/rest/v2/#api-rest-api-2-issueLinkType-issueLinkTypeId-get
func (s *IssueLinkTypeService) Get(ctx context.Context, ID string) (*IssueLinkType, *Response, error) {
	apiEndPoint := fmt.Sprintf("rest/api/2/issueLinkType/%s", ID)
	req, err := s.client.NewRequest(ctx, http.MethodGet, apiEndPoint, nil)
	if err != nil {
		return nil, nil, err
	}

	linkType := new(IssueLinkType)
	resp, err := s.client.Do(req, linkType)
	if err != nil {
		return nil, resp, NewJiraError(resp, err)
	}
	return linkType, resp, nil
}

// Create creates an issue link type in Jira.
//
// Jira API docs: https://developer.atlassian.com/cloud/jira/platform/rest/v2/#api-rest-api-2-issueLinkType-post
func (s *IssueLinkTypeService) Create(ctx context.Context, linkType *IssueLinkType) (*IssueLinkType, *Response, error) {
	apiEndpoint := "/rest/api/2/issueLinkType"
	req, err := s.client.NewRequest(ctx, http.MethodPost, apiEndpoint, linkType)
	if err != nil {
		return nil, nil, err
	}

	resp, err := s.client.Do(req, nil)
	if err != nil {
		return nil, resp, err
	}

	responseLinkType := new(IssueLinkType)
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		e := fmt.Errorf("could not read the returned data")
		return nil, resp, NewJiraError(resp, e)
	}
	err = json.Unmarshal(data, responseLinkType)
	if err != nil {
		e := fmt.Errorf("could no unmarshal the data into struct")
		return nil, resp, NewJiraError(resp, e)
	}
	return linkType, resp, nil
}

// Update updates an issue link type.  The issue is found by key.
//
// Jira API docs: https://developer.atlassian.com/cloud/jira/platform/rest/v2/#api-rest-api-2-issueLinkType-issueLinkTypeId-put
// Caller must close resp.Body
func (s *IssueLinkTypeService) Update(ctx context.Context, linkType *IssueLinkType) (*IssueLinkType, *Response, error) {
	apiEndpoint := fmt.Sprintf("rest/api/2/issueLinkType/%s", linkType.ID)
	req, err := s.client.NewRequest(ctx, http.MethodPut, apiEndpoint, linkType)
	if err != nil {
		return nil, nil, err
	}
	resp, err := s.client.Do(req, nil)
	if err != nil {
		return nil, resp, NewJiraError(resp, err)
	}
	ret := *linkType
	return &ret, resp, nil
}

// Delete deletes an issue link type based on provided ID.
//
// Jira API docs: https://developer.atlassian.com/cloud/jira/platform/rest/v2/#api-rest-api-2-issueLinkType-issueLinkTypeId-delete
// Caller must close resp.Body
func (s *IssueLinkTypeService) Delete(ctx context.Context, ID string) (*Response, error) {
	apiEndpoint := fmt.Sprintf("rest/api/2/issueLinkType/%s", ID)
	req, err := s.client.NewRequest(ctx, http.MethodDelete, apiEndpoint, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req, nil)
	return resp, err
}
