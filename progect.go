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
	Expand     string `json:"expand"`
	Self       string `json:"self"`
	ID         string `json:"id"`
	Key        string `json:"key"`
	Name       string `json:"name"`
	AvatarUrls struct {
		Four8X48  string `json:"48x48"`
		Two4X24   string `json:"24x24"`
		One6X16   string `json:"16x16"`
		Three2X32 string `json:"32x32"`
	} `json:"avatarUrls"`
	ProjectTypeKey  string `json:"projectTypeKey"`
	ProjectCategory struct {
		Self        string `json:"self"`
		ID          string `json:"id"`
		Name        string `json:"name"`
		Description string `json:"description"`
	} `json:"projectCategory,omitempty"`
}

// Full project description
// Can't name it project because it exist in issue
type FullProject struct {
	Expand      string `json:"expand"`
	Self        string `json:"self"`
	ID          string `json:"id"`
	Key         string `json:"key"`
	Description string `json:"description"`
	Lead        struct {
		Self       string `json:"self"`
		Name       string `json:"name"`
		AvatarUrls struct {
			Four8X48  string `json:"48x48"`
			Two4X24   string `json:"24x24"`
			One6X16   string `json:"16x16"`
			Three2X32 string `json:"32x32"`
		} `json:"avatarUrls"`
		DisplayName string `json:"displayName"`
		Active      bool   `json:"active"`
	} `json:"lead"`
	Components []struct {
		Self        string `json:"self"`
		ID          string `json:"id"`
		Name        string `json:"name"`
		Description string `json:"description"`
		Lead        struct {
			Self       string `json:"self"`
			Name       string `json:"name"`
			AvatarUrls struct {
				Four8X48  string `json:"48x48"`
				Two4X24   string `json:"24x24"`
				One6X16   string `json:"16x16"`
				Three2X32 string `json:"32x32"`
			} `json:"avatarUrls"`
			DisplayName string `json:"displayName"`
			Active      bool   `json:"active"`
		} `json:"lead"`
		AssigneeType string `json:"assigneeType"`
		Assignee     struct {
			Self       string `json:"self"`
			Name       string `json:"name"`
			AvatarUrls struct {
				Four8X48  string `json:"48x48"`
				Two4X24   string `json:"24x24"`
				One6X16   string `json:"16x16"`
				Three2X32 string `json:"32x32"`
			} `json:"avatarUrls"`
			DisplayName string `json:"displayName"`
			Active      bool   `json:"active"`
		} `json:"assignee"`
		RealAssigneeType string `json:"realAssigneeType"`
		RealAssignee     struct {
			Self       string `json:"self"`
			Name       string `json:"name"`
			AvatarUrls struct {
				Four8X48  string `json:"48x48"`
				Two4X24   string `json:"24x24"`
				One6X16   string `json:"16x16"`
				Three2X32 string `json:"32x32"`
			} `json:"avatarUrls"`
			DisplayName string `json:"displayName"`
			Active      bool   `json:"active"`
		} `json:"realAssignee"`
		IsAssigneeTypeValid bool   `json:"isAssigneeTypeValid"`
		Project             string `json:"project"`
		ProjectID           int    `json:"projectId"`
	} `json:"components"`
	IssueTypes []struct {
		Self        string `json:"self"`
		ID          string `json:"id"`
		Description string `json:"description"`
		IconURL     string `json:"iconUrl"`
		Name        string `json:"name"`
		Subtask     bool   `json:"subtask"`
		AvatarID    int    `json:"avatarId"`
	} `json:"issueTypes"`
	URL          string        `json:"url"`
	Email        string        `json:"email"`
	AssigneeType string        `json:"assigneeType"`
	Versions     []interface{} `json:"versions"`
	Name         string        `json:"name"`
	Roles        struct {
		Developers string `json:"Developers"`
	} `json:"roles"`
	AvatarUrls struct {
		Four8X48  string `json:"48x48"`
		Two4X24   string `json:"24x24"`
		One6X16   string `json:"16x16"`
		Three2X32 string `json:"32x32"`
	} `json:"avatarUrls"`
	ProjectCategory struct {
		Self        string `json:"self"`
		ID          string `json:"id"`
		Name        string `json:"name"`
		Description string `json:"description"`
	} `json:"projectCategory"`
}

// Get all projects form jira
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
