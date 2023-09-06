package onpremise

import (
	"context"
	"fmt"
	"net/http"
)

// ComponentService handles components for the Jira instance / API.//
// Jira API docs: https://docs.atlassian.com/software/jira/docs/api/REST/7.10.1/#api/2/component
type ComponentService service

// CreateComponentOptions are passed to the ComponentService.Create function to create a new Jira component
type ComponentOptions struct {
	Self         string `json:"self,omitempty" structs:"self,omitempty"`
	Name         string `json:"name,omitempty" structs:"name,omitempty"`
	Description  string `json:"description,omitempty" structs:"description,omitempty"`
	Lead         *User  `json:"lead,omitempty" structs:"lead,omitempty"`
	LeadUserName string `json:"leadUserName,omitempty" structs:"leadUserName,omitempty"`
	AssigneeType string `json:"assigneeType,omitempty" structs:"assigneeType,omitempty"`
	Assignee     *User  `json:"assignee,omitempty" structs:"assignee,omitempty"`
	Project      string `json:"project,omitempty" structs:"project,omitempty"`
	ProjectID    int    `json:"projectId,omitempty" structs:"projectId,omitempty"`
}

// Create creates a new Jira component based on the given options.
//
// TODO Double check this method if this works as expected, is using the latest API and the response is complete
// This double check effort is done for v2 - Remove this two lines if this is completed.
func (s *ComponentService) Create(ctx context.Context, options *ComponentOptions) (*ProjectComponent, *Response, error) {
	apiEndpoint := "rest/api/2/component"
	req, err := s.client.NewRequest(ctx, http.MethodPost, apiEndpoint, options)
	if err != nil {
		return nil, nil, err
	}

	component := new(ProjectComponent)
	resp, err := s.client.Do(req, component)

	if err != nil {
		return nil, resp, NewJiraError(resp, err)
	}

	return component, resp, nil
}

// GetListByProject gets all component from Jira project
//
// Jira API docs: https://developer.atlassian.com/cloud/jira/platform/rest/v2/api-group-project-components/#api-rest-api-2-project-projectidorkey-components-get
//
// TODO Double check this method if this works as expected, is using the latest API and the response is complete
// This double check effort is done for v2 - Remove this two lines if this is completed.
func (s *ComponentService) GetListByProject(ctx context.Context, projectID string) ([]ComponentOptions, *Response, error) {
	apiEndpoint := fmt.Sprintf("/rest/api/2/project/%s/components", projectID)
	req, err := s.client.NewRequest(ctx, http.MethodGet, apiEndpoint, nil)
	if err != nil {
		return nil, nil, err
	}

	componentList := []ComponentOptions{}
	resp, err := s.client.Do(req, &componentList)
	if err != nil {
		return nil, resp, NewJiraError(resp, err)
	}
	return componentList, resp, nil
}
