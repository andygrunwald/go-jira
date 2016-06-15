package jira

import (
	"fmt"
)

// ProjectService handles projects for the JIRA instance / API.
//
// JIRA API docs: https://docs.atlassian.com/jira/REST/latest/#api/2/project
type ProjectService struct {
	client *Client
}

// ProjectList represent a list of Projects
type ProjectList []struct {
	Expand          string          `json:"expand"`
	Self            string          `json:"self"`
	ID              string          `json:"id"`
	Key             string          `json:"key"`
	Name            string          `json:"name"`
	AvatarUrls      AvatarUrls      `json:"avatarUrls"`
	ProjectTypeKey  string          `json:"projectTypeKey"`
	ProjectCategory ProjectCategory `json:"projectCategory,omitempty"`
}

// ProjectCategory represents a single project category
type ProjectCategory struct {
	Self        string `json:"self"`
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// Project represents a JIRA Project.
type Project struct {
	Expand       string             `json:"expand,omitempty"`
	Self         string             `json:"self,omitempty"`
	ID           string             `json:"id,omitempty"`
	Key          string             `json:"key,omitempty"`
	Description  string             `json:"description,omitempty"`
	Lead         User               `json:"lead,omitempty"`
	Components   []ProjectComponent `json:"components,omitempty"`
	IssueTypes   []IssueType        `json:"issueTypes,omitempty"`
	URL          string             `json:"url,omitempty"`
	Email        string             `json:"email,omitempty"`
	AssigneeType string             `json:"assigneeType,omitempty"`
	Versions     []interface{}      `json:"versions,omitempty"`
	Name         string             `json:"name,omitempty"`
	Roles        struct {
		Developers string `json:"Developers,omitempty"`
	} `json:"roles,omitempty"`
	AvatarUrls      AvatarUrls      `json:"avatarUrls,omitempty"`
	ProjectCategory ProjectCategory `json:"projectCategory,omitempty"`
}

// ProjectComponent represents a single component of a project
type ProjectComponent struct {
	Self                string `json:"self"`
	ID                  string `json:"id"`
	Name                string `json:"name"`
	Description         string `json:"description"`
	Lead                User   `json:"lead"`
	AssigneeType        string `json:"assigneeType"`
	Assignee            User   `json:"assignee"`
	RealAssigneeType    string `json:"realAssigneeType"`
	RealAssignee        User   `json:"realAssignee"`
	IsAssigneeTypeValid bool   `json:"isAssigneeTypeValid"`
	Project             string `json:"project"`
	ProjectID           int    `json:"projectId"`
}

// GetList gets all projects form JIRA
//
// JIRA API docs: https://docs.atlassian.com/jira/REST/latest/#api/2/project-getAllProjects
func (s *ProjectService) GetList() (*ProjectList, *Response, error) {
	apiEndpoint := "rest/api/2/project"
	req, err := s.client.NewRequest("GET", apiEndpoint, nil)
	if err != nil {
		return nil, nil, err
	}

	projectList := new(ProjectList)
	resp, err := s.client.Do(req, projectList)
	if err != nil {
		return nil, resp, err
	}
	return projectList, resp, nil
}

// Get returns a full representation of the project for the given issue key.
// JIRA will attempt to identify the project by the projectIdOrKey path parameter.
// This can be an project id, or an project key.
//
// JIRA API docs: https://docs.atlassian.com/jira/REST/latest/#api/2/project-getProject
func (s *ProjectService) Get(projectID string) (*Project, *Response, error) {
	apiEndpoint := fmt.Sprintf("/rest/api/2/project/%s", projectID)
	req, err := s.client.NewRequest("GET", apiEndpoint, nil)
	if err != nil {
		return nil, nil, err
	}

	project := new(Project)
	resp, err := s.client.Do(req, project)
	if err != nil {
		return nil, resp, err
	}
	return project, resp, nil
}
