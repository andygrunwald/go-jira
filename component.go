package jira

import (
	"context"
	"fmt"
)

// ComponentService handles components for the Jira instance / API.//
// Jira API docs: https://docs.atlassian.com/software/jira/docs/api/REST/7.10.1/#api/2/component
type ComponentService struct {
	client *Client
}

// CreateComponentOptions are passed to the ComponentService.Create function to create a new Jira component
type CreateComponentOptions struct {
	Name         string `json:"name,omitempty" structs:"name,omitempty"`
	Description  string `json:"description,omitempty" structs:"description,omitempty"`
	Lead         *User  `json:"lead,omitempty" structs:"lead,omitempty"`
	LeadUserName string `json:"leadUserName,omitempty" structs:"leadUserName,omitempty"`
	AssigneeType string `json:"assigneeType,omitempty" structs:"assigneeType,omitempty"`
	Assignee     *User  `json:"assignee,omitempty" structs:"assignee,omitempty"`
	Project      string `json:"project,omitempty" structs:"project,omitempty"`
	ProjectID    int    `json:"projectId,omitempty" structs:"projectId,omitempty"`
}

// CreateWithContext creates a new Jira component based on the given options.
func (s *ComponentService) CreateWithContext(ctx context.Context, options *CreateComponentOptions) (*ProjectComponent, *Response, error) {
	apiEndpoint := "rest/api/2/component"
	req, err := s.client.NewRequestWithContext(ctx, "POST", apiEndpoint, options)
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

// Create wraps CreateWithContext using the background context.
func (s *ComponentService) Create(options *CreateComponentOptions) (*ProjectComponent, *Response, error) {
	return s.CreateWithContext(context.Background(), options)
}

// GetWithContext returns a full representation of the JIRA component for the given component ID.
// Jira API docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-project-components/
func (s *ComponentService) GetWithContext(ctx context.Context, componentID string) (*ProjectComponent, *Response, error) {
	apiEndpoint := fmt.Sprintf("rest/api/2/component/%s", componentID)
	req, err := s.client.NewRequestWithContext(ctx, "GET", apiEndpoint, nil)
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

// Get wraps GetWithContext using the background context.
func (s *ComponentService) Get(componentID string) (*ProjectComponent, *Response, error) {
	return s.GetWithContext(context.Background(), componentID)
}
