package jira

import (
	"fmt"
	"net/http"
)

type ProjectService struct {
	client *Client
}

// Project list type
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

type ProjectCategory struct {
	Self        string `json:"self"`
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// Full project description
// Can't name it project because it exist in issue
type FullProject struct {
	Expand       string             `json:"expand"`
	Self         string             `json:"self"`
	ID           string             `json:"id"`
	Key          string             `json:"key"`
	Description  string             `json:"description"`
	Lead         User               `json:"lead"`
	Components   []ProjectComponent `json:"components"`
	IssueTypes   []IssueType        `json:"issueTypes"`
	URL          string             `json:"url"`
	Email        string             `json:"email"`
	AssigneeType string             `json:"assigneeType"`
	Versions     []interface{}      `json:"versions"`
	Name         string             `json:"name"`
	Roles        struct {
		Developers string `json:"Developers"`
	} `json:"roles"`
	AvatarUrls      AvatarUrls      `json:"avatarUrls"`
	ProjectCategory ProjectCategory `json:"projectCategory"`
}

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
func (s *ProjectService) GetList() (*ProjectList, *http.Response, error) {
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
func (s *ProjectService) Get(projectID string) (*FullProject, *http.Response, error) {
	apiEndpoint := fmt.Sprintf("/rest/api/2/project/%s", projectID)
	req, err := s.client.NewRequest("GET", apiEndpoint, nil)
	if err != nil {
		return nil, nil, err
	}

	project := new(FullProject)
	resp, err := s.client.Do(req, project)
	if err != nil {
		return nil, resp, err
	}
	return project, resp, nil
}
