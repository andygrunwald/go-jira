package cloud

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
)

// GroupService handles Groups for the Jira instance / API.
//
// Jira API docs: https://docs.atlassian.com/jira/REST/server/#api/2/group
type GroupService service

// groupMembersResult is only a small wrapper around the Group* methods
// to be able to parse the results
type groupMembersResult struct {
	StartAt    int           `json:"startAt"`
	MaxResults int           `json:"maxResults"`
	Total      int           `json:"total"`
	Members    []GroupMember `json:"values"`
}

// Response body of https://developer.atlassian.com/cloud/jira/platform/rest/v2/api-group-groups/#api-rest-api-2-group-member-get
type getGroupMembersResult struct {
	IsLast     bool          `json:"isLast"`
	MaxResults int           `json:"maxResults"`
	NextPage   string        `json:"nextPage"`
	Total      int           `json:"total"`
	StartAt    int           `json:"startAt"`
	Values     []GroupMember `json:"values"`
}

// Group represents a Jira group
type Group struct {
	ID     string       `json:"groupId,omitempty" structs:"groupId,omitempty"`
	Name   string       `json:"name,omitempty" structs:"name,omitempty"`
	Self   string       `json:"self,omitempty" structs:"self,omitempty"`
	Users  GroupMembers `json:"users,omitempty" structs:"users,omitempty"`
	Expand string       `json:"expand,omitempty" structs:"expand,omitempty"`
}

// GroupMembers represent members in a Jira group
type GroupMembers struct {
	Size       int           `json:"size,omitempty" structs:"size,omitempty"`
	Items      []GroupMember `json:"items,omitempty" structs:"items,omitempty"`
	MaxResults int           `json:"max-results,omitempty" structs:"max-results.omitempty"`
	StartIndex int           `json:"start-index,omitempty" structs:"start-index,omitempty"`
	EndIndex   int           `json:"end-index,omitempty" structs:"end-index,omitempty"`
}

// GroupMember reflects a single member of a group
type GroupMember struct {
	Self         string `json:"self,omitempty"`
	Name         string `json:"name,omitempty"`
	Key          string `json:"key,omitempty"`
	AccountID    string `json:"accountId,omitempty"`
	EmailAddress string `json:"emailAddress,omitempty"`
	DisplayName  string `json:"displayName,omitempty"`
	Active       bool   `json:"active,omitempty"`
	TimeZone     string `json:"timeZone,omitempty"`
	AccountType  string `json:"accountType,omitempty"`
}

// GroupSearchOptions specifies the optional parameters for the Get Group methods
type GroupSearchOptions struct {
	StartAt              int
	MaxResults           int
	IncludeInactiveUsers bool
}

type Groups struct {
	Groups []Group `json:"groups,omitempty"`
	Header string  `json:"header,omitempty"`
	Total  int     `json:"total,omitempty"`
}

// Get returns a paginated list of members of the specified group and its subgroups.
// Users in the page are ordered by user names.
// User of this resource is required to have sysadmin or admin permissions.
//
// Jira API docs: https://docs.atlassian.com/jira/REST/server/#api/2/group-getUsersFromGroup
//
// WARNING: This API only returns the first page of group members
//
// TODO Double check this method if this works as expected, is using the latest API and the response is complete
// This double check effort is done for v2 - Remove this two lines if this is completed.
// Deprecated: Use GetGroupMembers instead
func (s *GroupService) Get(ctx context.Context, name string, options *GroupSearchOptions) ([]GroupMember, *Response, error) {
	var apiEndpoint string
	if options == nil {
		apiEndpoint = fmt.Sprintf("/rest/api/2/group/member?groupname=%s", url.QueryEscape(name))
	} else {
		// TODO use addOptions
		apiEndpoint = fmt.Sprintf(
			"/rest/api/2/group/member?groupname=%s&startAt=%d&maxResults=%d&includeInactiveUsers=%t",
			url.QueryEscape(name),
			options.StartAt,
			options.MaxResults,
			options.IncludeInactiveUsers,
		)
	}
	req, err := s.client.NewRequest(ctx, http.MethodGet, apiEndpoint, nil)
	if err != nil {
		return nil, nil, err
	}

	group := new(groupMembersResult)
	resp, err := s.client.Do(req, group)
	if err != nil {
		return nil, resp, err
	}
	return group.Members, resp, nil
}

// Add adds a user to a group.
//
// The account ID of the user, which uniquely identifies the user across all Atlassian products.
// For example, 5b10ac8d82e05b22cc7d4ef5.
//
// Jira API docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-groups/#api-rest-api-3-group-user-post
func (s *GroupService) AddUserByGroupName(ctx context.Context, groupName string, accountID string) (*Group, *Response, error) {
	apiEndpoint := fmt.Sprintf("/rest/api/3/group/user?groupname=%s", groupName)
	var user struct {
		AccountID string `json:"accountId"`
	}
	user.AccountID = accountID
	req, err := s.client.NewRequest(ctx, http.MethodPost, apiEndpoint, &user)
	if err != nil {
		return nil, nil, err
	}

	responseGroup := new(Group)
	resp, err := s.client.Do(req, responseGroup)
	if err != nil {
		jerr := NewJiraError(resp, err)
		return nil, resp, jerr
	}

	return responseGroup, resp, nil
}

// Remove removes a user from a group.
//
// The account ID of the user, which uniquely identifies the user across all Atlassian products.
// For example, 5b10ac8d82e05b22cc7d4ef5.
//
// Jira API docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-groups/#api-rest-api-3-group-user-delete
// Caller must close resp.Body
func (s *GroupService) RemoveUserByGroupName(ctx context.Context, groupName string, accountID string) (*Response, error) {
	apiEndpoint := fmt.Sprintf("/rest/api/3/group/user?groupname=%s&accountId=%s", groupName, accountID)
	req, err := s.client.NewRequest(ctx, http.MethodDelete, apiEndpoint, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req, nil)
	if err != nil {
		jerr := NewJiraError(resp, err)
		return resp, jerr
	}

	return resp, nil
}

// Sets case insensitive search
func WithCaseInsensitive() searchF {
	return func(s search) search {
		s = append(s, searchParam{name: "caseInsensitive", value: "true"})
		return s
	}
}

// Sets query string for filtering group names.
func WithGroupNameContains(contains string) searchF {
	return func(s search) search {
		s = append(s, searchParam{name: "query", value: contains})
		return s
	}
}

// Sets excluded group names.
func WithExcludedGroupNames(excluded []string) searchF {
	return func(s search) search {
		for _, name := range excluded {
			s = append(s, searchParam{name: "exclude", value: name})
		}

		return s
	}
}

// Sets excluded group ids.
func WithExcludedGroupsIds(excluded []string) searchF {
	return func(s search) search {
		for _, id := range excluded {
			s = append(s, searchParam{name: "excludeId", value: id})
		}

		return s
	}
}

// Search for the groups
// It can search by groupId, accountId or userName
// Apart from returning groups it also returns total number of groups
//
// Jira API docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-groups/#api-rest-api-3-groups-picker-get
func (s *GroupService) Find(ctx context.Context, tweaks ...searchF) ([]Group, *Response, error) {
	search := []searchParam{}
	for _, f := range tweaks {
		search = f(search)
	}

	apiEndpoint := "/rest/api/3/groups/picker"

	queryString := ""
	for _, param := range search {
		queryString += fmt.Sprintf("%s=%s&", param.name, param.value)
	}

	if queryString != "" {
		apiEndpoint += "?" + queryString
	}

	req, err := s.client.NewRequest(ctx, http.MethodGet, apiEndpoint, nil)
	if err != nil {
		return nil, nil, err
	}

	groups := Groups{}
	resp, err := s.client.Do(req, &groups)
	if err != nil {
		return nil, resp, NewJiraError(resp, err)
	}

	return groups.Groups, resp, nil
}

func WithInactiveUsers() searchF {
	return func(s search) search {
		s = append(s, searchParam{name: "includeInactiveUsers", value: "true"})
		return s
	}
}

// Search for the group members
// It can filter out inactive users
// Apart from returning group members it also returns total number of group members
//
// Jira API docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-groups/#api-rest-api-3-group-member-get
func (s *GroupService) GetGroupMembers(ctx context.Context, groupId, groupName string, tweaks ...searchF) ([]GroupMember, *Response, error) {
	search := []searchParam{}
	for _, f := range tweaks {
		search = f(search)
	}

	apiEndpoint := fmt.Sprintf("/rest/api/3/group/member?groupid=%s&groupname=%s", groupId, groupName)

	queryString := ""
	for _, param := range search {
		queryString += fmt.Sprintf("%s=%s&", param.name, param.value)
	}

	if queryString != "" {
		apiEndpoint += "&" + queryString
	}

	req, err := s.client.NewRequest(ctx, http.MethodGet, apiEndpoint, nil)
	if err != nil {
		return nil, nil, err
	}

	group := new(getGroupMembersResult)
	resp, err := s.client.Do(req, group)
	if err != nil {
		return nil, resp, NewJiraError(resp, err)
	}

	return group.Values, resp, nil
}
