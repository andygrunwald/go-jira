package jira

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// ComponentService handles components for the JIRA instance / API.
//
// JIRA API docs: https://docs.atlassian.com/software/jira/docs/api/REST/7.10.1/#api/2/component
type ComponentService struct {
	client *Client
}

// CreateComponentOptions are passed to the ComponentService.Create function to create a new JIRA component
type CreateComponentOptions struct {
	Name         string
	Description  string
	LeadUserName string
	AssigneeType string
	Lead         *User
	Assignee     *User
	Project      string
	ProjectID    *int
}

// FullComponent is a JIRA component with all details filled in
type FullComponent struct {
	Name                string `json:"name,omitempty" structs:"name,omitempty"`
	Description         string `json:"description,omitempty" structs:"description,omitempty"`
	Lead                *User  `json:"lead,omitempty" structs:"lead,omitempty"`
	Assignee            *User  `json:"assignee,omitempty" structs:"assignee,omitempty"`
	LeadUserName        string `json:"leadUserName,omitempty" structs:"leadUserName,omitempty"`
	AssigneeType        string `json:"assigneeType,omitempty" structs:"assigneeType,omitempty"`
	RealAssigneeType    string `json:"realAssigneeType,omitempty" structs:"realAssigneeType,omitempty"`
	RealAssignee        *User  `json:"realAssignee,omitempty" structs:"realAssignee,omitempty"`
	IsAssigneeTypeValid bool   `json:"isAssigneeTypeValid,omitempty" structs:"isAssigneeTypeValid,omitempty"`
	Project             string `json:"project,omitempty" structs:"project,omitempty"`
	ProjectID           int    `json:"projectId,omitempty" structs:"projectId,omitempty"`
}

// Create creates a new JIRA component based on the given options.
func (s *ComponentService) Create(options *CreateComponentOptions) (*FullComponent, *Response, error) {
	apiEndpoint := "rest/api/2/component"
	req, err := s.client.NewRequest("POST", apiEndpoint, options)
	if err != nil {
		return nil, nil, err
	}
	resp, err := s.client.Do(req, nil)
	if err != nil {
		// incase of error return the resp for further inspection
		return nil, resp, err
	}

	component := new(FullComponent)
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, resp, fmt.Errorf("Could not read the returned data")
	}
	err = json.Unmarshal(data, component)
	if err != nil {
		return nil, resp, fmt.Errorf("Could not unmarshall the data into struct")
	}
	return component, resp, nil
}
