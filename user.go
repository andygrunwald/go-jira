package jira

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/google/go-querystring/query"
)

// UserService handles users for the JIRA instance / API.
//
// JIRA API docs: https://docs.atlassian.com/jira/REST/cloud/#api/2/user
type UserService struct {
	client *Client
}

// User represents a JIRA user.
type User struct {
	Self            string     `json:"self,omitempty" structs:"self,omitempty"`
	Name            string     `json:"name,omitempty" structs:"name,omitempty"`
	Password        string     `json:"-"`
	Key             string     `json:"key,omitempty" structs:"key,omitempty"`
	EmailAddress    string     `json:"emailAddress,omitempty" structs:"emailAddress,omitempty"`
	AvatarUrls      AvatarUrls `json:"avatarUrls,omitempty" structs:"avatarUrls,omitempty"`
	DisplayName     string     `json:"displayName,omitempty" structs:"displayName,omitempty"`
	Active          bool       `json:"active,omitempty" structs:"active,omitempty"`
	TimeZone        string     `json:"timeZone,omitempty" structs:"timeZone,omitempty"`
	ApplicationKeys []string   `json:"applicationKeys,omitempty" structs:"applicationKeys,omitempty"`
}

// UserGroup represents the group list
type UserGroup struct {
	Self string `json:"self,omitempty" structs:"self,omitempty"`
	Name string `json:"name,omitempty" structs:"name,omitempty"`
}

type userSearchParam struct {
	name  string
	value string
}

type userSearch []userSearchParam

type userSearchF func(userSearch) userSearch

// Get gets user info from JIRA
//
// JIRA API docs: https://docs.atlassian.com/jira/REST/cloud/#api/2/user-getUser
func (s *UserService) Get(username string) (*User, *Response, error) {
	apiEndpoint := fmt.Sprintf("/rest/api/2/user?username=%s", username)
	req, err := s.client.NewRequest("GET", apiEndpoint, nil)
	if err != nil {
		return nil, nil, err
	}

	user := new(User)
	resp, err := s.client.Do(req, user)
	if err != nil {
		return nil, resp, NewJiraError(resp, err)
	}
	return user, resp, nil
}

// Create creates an user in JIRA.
//
// JIRA API docs: https://docs.atlassian.com/jira/REST/cloud/#api/2/user-createUser
func (s *UserService) Create(user *User) (*User, *Response, error) {
	apiEndpoint := "/rest/api/2/user"
	req, err := s.client.NewRequest("POST", apiEndpoint, user)
	if err != nil {
		return nil, nil, err
	}

	resp, err := s.client.Do(req, nil)
	if err != nil {
		return nil, resp, err
	}

	responseUser := new(User)
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		e := fmt.Errorf("Could not read the returned data")
		return nil, resp, NewJiraError(resp, e)
	}
	err = json.Unmarshal(data, responseUser)
	if err != nil {
		e := fmt.Errorf("Could not unmarshall the data into struct")
		return nil, resp, NewJiraError(resp, e)
	}
	return responseUser, resp, nil
}

// Delete deletes an user from JIRA.
// Returns http.StatusNoContent on success.
//
// JIRA API docs: https://developer.atlassian.com/cloud/jira/platform/rest/#api-api-2-user-delete
func (s *UserService) Delete(username string) (*Response, error) {
	apiEndpoint := fmt.Sprintf("/rest/api/2/user?username=%s", username)
	req, err := s.client.NewRequest("DELETE", apiEndpoint, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req, nil)
	if err != nil {
		return resp, NewJiraError(resp, err)
	}
	return resp, nil
}

// GetGroups returns the groups which the user belongs to
//
// JIRA API docs: https://docs.atlassian.com/jira/REST/cloud/#api/2/user-getUserGroups
func (s *UserService) GetGroups(username string) (*[]UserGroup, *Response, error) {
	apiEndpoint := fmt.Sprintf("/rest/api/2/user/groups?username=%s", username)
	req, err := s.client.NewRequest("GET", apiEndpoint, nil)
	if err != nil {
		return nil, nil, err
	}

	userGroups := new([]UserGroup)
	resp, err := s.client.Do(req, userGroups)
	if err != nil {
		return nil, resp, NewJiraError(resp, err)
	}
	return userGroups, resp, nil
}

// Get information about the current logged-in user
//
// JIRA API docs: https://developer.atlassian.com/cloud/jira/platform/rest/#api-api-2-myself-get
func (s *UserService) GetSelf() (*User, *Response, error) {
	const apiEndpoint = "rest/api/2/myself"
	req, err := s.client.NewRequest("GET", apiEndpoint, nil)
	if err != nil {
		return nil, nil, err
	}
	var user User
	resp, err := s.client.Do(req, &user)
	if err != nil {
		return nil, resp, NewJiraError(resp, err)
	}
	return &user, resp, nil
}

// WithMaxResults sets the max results to return
func WithMaxResults(maxResults int) userSearchF {
	return func(s userSearch) userSearch {
		s = append(s, userSearchParam{name: "maxResults", value: fmt.Sprintf("%d", maxResults)})
		return s
	}
}

// WithStartAt set the start pager
func WithStartAt(startAt int) userSearchF {
	return func(s userSearch) userSearch {
		s = append(s, userSearchParam{name: "startAt", value: fmt.Sprintf("%d", startAt)})
		return s
	}
}

// WithActive sets the active users lookup
func WithActive(active bool) userSearchF {
	return func(s userSearch) userSearch {
		s = append(s, userSearchParam{name: "includeActive", value: fmt.Sprintf("%t", active)})
		return s
	}
}

// WithInactive sets the inactive users lookup
func WithInactive(inactive bool) userSearchF {
	return func(s userSearch) userSearch {
		s = append(s, userSearchParam{name: "includeInactive", value: fmt.Sprintf("%t", inactive)})
		return s
	}
}

// Find searches for user info from JIRA:
// It can find users by email, username or name
//
// JIRA API docs: https://docs.atlassian.com/jira/REST/cloud/#api/2/user-findUsers
func (s *UserService) Find(property string, tweaks ...userSearchF) ([]User, *Response, error) {
	search := []userSearchParam{
		{
			name:  "username",
			value: property,
		},
	}
	for _, f := range tweaks {
		search = f(search)
	}

	var queryString = ""
	for _, param := range search {
		queryString += param.name + "=" + param.value + "&"
	}

	apiEndpoint := fmt.Sprintf("/rest/api/2/user/search?" + queryString[:len(queryString)-1])
	req, err := s.client.NewRequest("GET", apiEndpoint, nil)
	if err != nil {
		return nil, nil, err
	}

	users := []User{}
	resp, err := s.client.Do(req, &users)
	if err != nil {
		return nil, resp, NewJiraError(resp, err)
	}
	return users, resp, nil
}

// UserFindAssignableToProjectsOptions specifies the optional parameters to search users that
// support pagination.
// Pagination is used for the JIRA REST APIs to conserve server resources and limit
// response size for resources that return potentially large collection of users.
// A request to a pages API will result in a values array wrapped in a JSON object with some paging metadata
// Default Pagination options
type UserFindAssignableToProjectsOptions struct {
	// StartAt: The starting index of the returned projects. Base index: 0.
	StartAt int `url:"startAt,omitempty"`
	// MaxResults: The maximum number of projects to return per page. Default: 50.
	MaxResults int `url:"maxResults,omitempty"`
	// Query: A search input that is matched against appropriate user attributes to find relevant users.
	Query string `url:"query,omitempty"`
	// Username: A query string used to search username, display name, and email address. The string is matched to the starting letters of any word in the searched fields. For example, ar matches to the username archie but not mark. This field is deprecated in favour of query parameter.
	Username string `url:"username,omitempty"`
	// ProjectKeys: A comma-separated list of project keys (case sensitive).
	ProjectKeys string `url:"projectKeys"`
}

func (s *UserService) FindAssignableToProjects(options *UserFindAssignableToProjectsOptions) ([]User, *Response, error) {
	q, err := query.Values(options)
	if err != nil {
		return nil, nil, err
	}
	apiEndpoint := "/rest/api/2/user/assignable/multiProjectSearch?" + q.Encode()
	req, err := s.client.NewRequest("GET", apiEndpoint, nil)
	if err != nil {
		return nil, nil, err
	}

	users := []User{}
	resp, err := s.client.Do(req, &users)
	if err != nil {
		return nil, resp, NewJiraError(resp, err)
	}
	return users, resp, nil
}
