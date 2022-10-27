package onpremise

import (
	"context"
	"fmt"
	"net/http"

	"github.com/google/go-querystring/query"
)

// ProjectService handles projects for the Jira instance / API.
//
// Jira API docs: https://docs.atlassian.com/jira/REST/latest/#api/2/project
type ProjectService service

// ProjectList represent a list of Projects
type ProjectList []struct {
	Expand          string          `json:"expand" structs:"expand"`
	Self            string          `json:"self" structs:"self"`
	ID              string          `json:"id" structs:"id"`
	Key             string          `json:"key" structs:"key"`
	Name            string          `json:"name" structs:"name"`
	AvatarUrls      AvatarUrls      `json:"avatarUrls" structs:"avatarUrls"`
	ProjectTypeKey  string          `json:"projectTypeKey" structs:"projectTypeKey"`
	ProjectCategory ProjectCategory `json:"projectCategory,omitempty" structs:"projectsCategory,omitempty"`
	IssueTypes      []IssueType     `json:"issueTypes,omitempty" structs:"issueTypes,omitempty"`
}

// ProjectCategory represents a single project category
type ProjectCategory struct {
	Self        string `json:"self" structs:"self,omitempty"`
	ID          string `json:"id" structs:"id,omitempty"`
	Name        string `json:"name" structs:"name,omitempty"`
	Description string `json:"description" structs:"description,omitempty"`
}

// Project represents a Jira Project.
type Project struct {
	Expand          string             `json:"expand,omitempty" structs:"expand,omitempty"`
	Self            string             `json:"self,omitempty" structs:"self,omitempty"`
	ID              string             `json:"id,omitempty" structs:"id,omitempty"`
	Key             string             `json:"key,omitempty" structs:"key,omitempty"`
	Description     string             `json:"description,omitempty" structs:"description,omitempty"`
	Lead            User               `json:"lead,omitempty" structs:"lead,omitempty"`
	Components      []ProjectComponent `json:"components,omitempty" structs:"components,omitempty"`
	IssueTypes      []IssueType        `json:"issueTypes,omitempty" structs:"issueTypes,omitempty"`
	URL             string             `json:"url,omitempty" structs:"url,omitempty"`
	Email           string             `json:"email,omitempty" structs:"email,omitempty"`
	AssigneeType    string             `json:"assigneeType,omitempty" structs:"assigneeType,omitempty"`
	Versions        []Version          `json:"versions,omitempty" structs:"versions,omitempty"`
	Name            string             `json:"name,omitempty" structs:"name,omitempty"`
	Roles           map[string]string  `json:"roles,omitempty" structs:"roles,omitempty"`
	AvatarUrls      AvatarUrls         `json:"avatarUrls,omitempty" structs:"avatarUrls,omitempty"`
	ProjectCategory ProjectCategory    `json:"projectCategory,omitempty" structs:"projectCategory,omitempty"`
}

// ProjectComponent represents a single component of a project
type ProjectComponent struct {
	Self                string `json:"self" structs:"self,omitempty"`
	ID                  string `json:"id" structs:"id,omitempty"`
	Name                string `json:"name" structs:"name,omitempty"`
	Description         string `json:"description" structs:"description,omitempty"`
	Lead                User   `json:"lead,omitempty" structs:"lead,omitempty"`
	AssigneeType        string `json:"assigneeType" structs:"assigneeType,omitempty"`
	Assignee            User   `json:"assignee" structs:"assignee,omitempty"`
	RealAssigneeType    string `json:"realAssigneeType" structs:"realAssigneeType,omitempty"`
	RealAssignee        User   `json:"realAssignee" structs:"realAssignee,omitempty"`
	IsAssigneeTypeValid bool   `json:"isAssigneeTypeValid" structs:"isAssigneeTypeValid,omitempty"`
	Project             string `json:"project" structs:"project,omitempty"`
	ProjectID           int    `json:"projectId" structs:"projectId,omitempty"`
}

// PermissionScheme represents the permission scheme for the project
type PermissionScheme struct {
	Expand      string       `json:"expand" structs:"expand,omitempty"`
	Self        string       `json:"self" structs:"self,omitempty"`
	ID          int          `json:"id" structs:"id,omitempty"`
	Name        string       `json:"name" structs:"name,omitempty"`
	Description string       `json:"description" structs:"description,omitempty"`
	Permissions []Permission `json:"permissions" structs:"permissions,omitempty"`
}

// GetAll returns all projects form Jira with optional query params, like &GetQueryOptions{Expand: "issueTypes"} to get
// a list of all projects and their supported issuetypes.
//
// Jira API docs: https://developer.atlassian.com/cloud/jira/platform/rest/v2/api-group-projects/#api-rest-api-2-project-get
//
// TODO Double check this method if this works as expected, is using the latest API and the response is complete
// This double check effort is done for v2 - Remove this two lines if this is completed.
func (s *ProjectService) GetAll(ctx context.Context, options *GetQueryOptions) (*ProjectList, *Response, error) {
	apiEndpoint := "rest/api/2/project"
	req, err := s.client.NewRequest(ctx, http.MethodGet, apiEndpoint, nil)
	if err != nil {
		return nil, nil, err
	}

	if options != nil {
		q, err := query.Values(options)
		if err != nil {
			return nil, nil, err
		}
		req.URL.RawQuery = q.Encode()
	}

	projectList := new(ProjectList)
	resp, err := s.client.Do(req, projectList)
	if err != nil {
		jerr := NewJiraError(resp, err)
		return nil, resp, jerr
	}

	return projectList, resp, nil
}

// Get returns a full representation of the project for the given issue key.
// Jira will attempt to identify the project by the projectIdOrKey path parameter.
// This can be an project id, or an project key.
//
// Jira API docs: https://docs.atlassian.com/jira/REST/latest/#api/2/project-getProject
//
// TODO Double check this method if this works as expected, is using the latest API and the response is complete
// This double check effort is done for v2 - Remove this two lines if this is completed.
func (s *ProjectService) Get(ctx context.Context, projectID string) (*Project, *Response, error) {
	apiEndpoint := fmt.Sprintf("rest/api/2/project/%s", projectID)
	req, err := s.client.NewRequest(ctx, http.MethodGet, apiEndpoint, nil)
	if err != nil {
		return nil, nil, err
	}

	project := new(Project)
	resp, err := s.client.Do(req, project)
	if err != nil {
		jerr := NewJiraError(resp, err)
		return nil, resp, jerr
	}

	return project, resp, nil
}

// GetPermissionScheme returns a full representation of the permission scheme for the project
// Jira will attempt to identify the project by the projectIdOrKey path parameter.
// This can be an project id, or an project key.
//
// Jira API docs: https://docs.atlassian.com/jira/REST/latest/#api/2/project-getProject
//
// TODO Double check this method if this works as expected, is using the latest API and the response is complete
// This double check effort is done for v2 - Remove this two lines if this is completed.
func (s *ProjectService) GetPermissionScheme(ctx context.Context, projectID string) (*PermissionScheme, *Response, error) {
	apiEndpoint := fmt.Sprintf("/rest/api/2/project/%s/permissionscheme", projectID)
	req, err := s.client.NewRequest(ctx, http.MethodGet, apiEndpoint, nil)
	if err != nil {
		return nil, nil, err
	}

	ps := new(PermissionScheme)
	resp, err := s.client.Do(req, ps)
	if err != nil {
		jerr := NewJiraError(resp, err)
		return nil, resp, jerr
	}

	return ps, resp, nil
}
