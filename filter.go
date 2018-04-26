package jira

import (
	"github.com/pkg/errors"
)

// FilterService handles filters for the JIRA instance / API.
//
// JIRA API docs: https://developer.atlassian.com/cloud/jira/platform/rest/#api-api-2-filter-get
type FilterService struct {
	client *Client
}

// Filter is JIRA is basically a saved search
type Filter struct {
	Self            string `json:"self,omitempty"`
	ID              string `json:"id,omitempty"`
	Name            string `json:"name,omitempty"`
	Description     string `json:"description,omitempty"`
	Owner           *User  `json:"owner,omitempty"`
	Jql             string `json:"jql,omitempty"`
	ViewURL         string `json:"viewUrl,omitempty"`
	SearchURL       string `json:"searchUrl,omitempty"`
	Favourite       bool   `json:"favourite,omitempty"`
	FavouritedCount int    `json:"favouritedCount,omitempty"`
}

// GetList returns all filters for the current user
//
// JIRA API docs: https://developer.atlassian.com/cloud/jira/platform/rest/#api-api-2-filter-get
func (s *FilterService) GetList(opt *GetQueryOptions) (*[]Filter, *Response, error) {
	endpoint := "rest/api/2/filter"
	return s.callWithOptions(endpoint, opt)
}

// // Get returns a filter given an ID
// //
// // JIRA API docs: https://developer.atlassian.com/cloud/jira/platform/rest/#api-api-2-filter-id-get
// func (s *FilterService) Get(ID string, opt *GetQueryOptions) (*Filter, *Response, error) {
// 	endpoint := fmt.Sprintf("rest/api/2/filter/%s", ID)
// 	return s.callWithOptions(endpoint, opt)
// }

// callWithOptions creates a request with options, issues it and returns the response
func (s *FilterService) callWithOptions(endpoint string, opt *GetQueryOptions) (*[]Filter, *Response, error) {
	url, err := addOptions(endpoint, opt)
	if err != nil {
		return nil, nil, errors.Wrap(err, "could not add query options")
	}

	req, err := s.client.NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}

	filters := new([]Filter)
	resp, err := s.client.Do(req, filters)

	if err != nil {
		jerr := NewJiraError(resp, err)
		return nil, resp, jerr
	}

	return filters, resp, nil
}
