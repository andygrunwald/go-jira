package cloud

import (
	"context"
	"fmt"
	"net/http"
)

// RoleService handles roles for the Jira instance / API.
//
// Jira API docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/#api-group-Role
type RoleService service

// Role represents a Jira product role
type Role struct {
	Self        string   `json:"self" structs:"self"`
	Name        string   `json:"name" structs:"name"`
	ID          int      `json:"id" structs:"id"`
	Description string   `json:"description" structs:"description"`
	Actors      []*Actor `json:"actors" structs:"actors"`
}

// Actor represents a Jira actor
type Actor struct {
	ID          int        `json:"id" structs:"id"`
	DisplayName string     `json:"displayName" structs:"displayName"`
	Type        string     `json:"type" structs:"type"`
	Name        string     `json:"name" structs:"name"`
	AvatarURL   string     `json:"avatarUrl" structs:"avatarUrl"`
	ActorUser   *ActorUser `json:"actorUser" structs:"actoruser"`
}

// ActorUser contains the account id of the actor/user
type ActorUser struct {
	AccountID string `json:"accountId" structs:"accountId"`
}

// GetList returns a list of all available project roles
//
// Jira API docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/#api-api-3-role-get
//
// TODO Double check this method if this works as expected, is using the latest API and the response is complete
// This double check effort is done for v2 - Remove this two lines if this is completed.
func (s *RoleService) GetList(ctx context.Context) (*[]Role, *Response, error) {
	apiEndpoint := "rest/api/3/role"
	req, err := s.client.NewRequest(ctx, http.MethodGet, apiEndpoint, nil)
	if err != nil {
		return nil, nil, err
	}
	roles := new([]Role)
	resp, err := s.client.Do(req, roles)
	if err != nil {
		jerr := NewJiraError(resp, err)
		return nil, resp, jerr
	}
	return roles, resp, err
}

// Get retreives a single Role from Jira
//
// Jira API docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/#api-api-3-role-id-get
//
// TODO Double check this method if this works as expected, is using the latest API and the response is complete
// This double check effort is done for v2 - Remove this two lines if this is completed.
func (s *RoleService) Get(ctx context.Context, roleID int) (*Role, *Response, error) {
	apiEndpoint := fmt.Sprintf("rest/api/3/role/%d", roleID)
	req, err := s.client.NewRequest(ctx, http.MethodGet, apiEndpoint, nil)
	if err != nil {
		return nil, nil, err
	}
	role := new(Role)
	resp, err := s.client.Do(req, role)
	if err != nil {
		jerr := NewJiraError(resp, err)
		return nil, resp, jerr
	}
	if role.Self == "" {
		return nil, resp, fmt.Errorf("no role with ID %d found", roleID)
	}

	return role, resp, err
}
