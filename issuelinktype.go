package jira

import "fmt"

// IssueLinkTypeService handles issue link types for the JIRA instance / API.
//
// JIRA API docs: https://developer.atlassian.com/cloud/jira/platform/rest/v2/#api-group-Issue-link-types
type IssueLinkTypeService struct {
	client *Client
}

// GetList gets all of the issue link types from JIRA.
//
// JIRA API docs: https://developer.atlassian.com/cloud/jira/platform/rest/v2/#api-rest-api-2-issueLinkType-get
func (s *IssueLinkTypeService) GetList() ([]IssueLinkType, *Response, error) {
	apiEndpoint := "rest/api/2/issueLinkType"
	req, err := s.client.NewRequest("GET", apiEndpoint, nil)
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

// Get gets info of a specific issue link type from JIRA.
//
// JIRA API docs: https://developer.atlassian.com/cloud/jira/platform/rest/v2/#api-rest-api-2-issueLinkType-issueLinkTypeId-get
func (s *IssueLinkTypeService) Get(ID string) (*IssueLinkType, *Response, error) {
	apiEndPoint := fmt.Sprintf("rest/api/2/issueLinkType/%s", ID)
	req, err := s.client.NewRequest("GET", apiEndPoint, nil)
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
