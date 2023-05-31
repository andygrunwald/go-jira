package onpremise

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// VersionService handles Versions for the Jira instance / API.
//
// Jira API docs: https://docs.atlassian.com/jira/REST/latest/#api/2/version
type VersionService service

// Version represents a single release version of a project
type Version struct {
	Self            string `json:"self,omitempty" structs:"self,omitempty"`
	ID              string `json:"id,omitempty" structs:"id,omitempty"`
	Name            string `json:"name,omitempty" structs:"name,omitempty"`
	Description     string `json:"description,omitempty" structs:"description,omitempty"`
	Archived        *bool  `json:"archived,omitempty" structs:"archived,omitempty"`
	Released        *bool  `json:"released,omitempty" structs:"released,omitempty"`
	ReleaseDate     string `json:"releaseDate,omitempty" structs:"releaseDate,omitempty"`
	UserReleaseDate string `json:"userReleaseDate,omitempty" structs:"userReleaseDate,omitempty"`
	ProjectID       int    `json:"projectId,omitempty" structs:"projectId,omitempty"` // Unlike other IDs, this is returned as a number
	StartDate       string `json:"startDate,omitempty" structs:"startDate,omitempty"`
}

// Get gets version info from Jira
//
// Jira API docs: https://developer.atlassian.com/cloud/jira/platform/rest/#api-api-2-version-id-get
//
// TODO Double check this method if this works as expected, is using the latest API and the response is complete
// This double check effort is done for v2 - Remove this two lines if this is completed.
func (s *VersionService) Get(ctx context.Context, versionID int) (*Version, *Response, error) {
	apiEndpoint := fmt.Sprintf("/rest/api/2/version/%v", versionID)
	req, err := s.client.NewRequest(ctx, http.MethodGet, apiEndpoint, nil)
	if err != nil {
		return nil, nil, err
	}

	version := new(Version)
	resp, err := s.client.Do(req, version)
	if err != nil {
		return nil, resp, NewJiraError(resp, err)
	}
	return version, resp, nil
}

// Create creates a version in Jira.
//
// Jira API docs: https://developer.atlassian.com/cloud/jira/platform/rest/#api-api-2-version-post
//
// TODO Double check this method if this works as expected, is using the latest API and the response is complete
// This double check effort is done for v2 - Remove this two lines if this is completed.
func (s *VersionService) Create(ctx context.Context, version *Version) (*Version, *Response, error) {
	apiEndpoint := "/rest/api/2/version"
	req, err := s.client.NewRequest(ctx, http.MethodPost, apiEndpoint, version)
	if err != nil {
		return nil, nil, err
	}

	resp, err := s.client.Do(req, nil)
	if err != nil {
		return nil, resp, err
	}
	defer resp.Body.Close()

	responseVersion := new(Version)
	err = json.NewDecoder(resp.Body).Decode(&responseVersion)
	if err != nil {
		return nil, resp, err
	}

	return responseVersion, resp, nil
}

// Update updates a version from a JSON representation.
//
// Jira API docs: https://developer.atlassian.com/cloud/jira/platform/rest/#api-api-2-version-id-put
// Caller must close resp.Body
//
// TODO Double check this method if this works as expected, is using the latest API and the response is complete
// This double check effort is done for v2 - Remove this two lines if this is completed.
func (s *VersionService) Update(ctx context.Context, version *Version) (*Version, *Response, error) {
	apiEndpoint := fmt.Sprintf("rest/api/2/version/%v", version.ID)
	req, err := s.client.NewRequest(ctx, http.MethodPut, apiEndpoint, version)
	if err != nil {
		return nil, nil, err
	}
	resp, err := s.client.Do(req, nil)
	if err != nil {
		jerr := NewJiraError(resp, err)
		return nil, resp, jerr
	}

	// This is just to follow the rest of the API's convention of returning a version.
	// Returning the same pointer here is pointless, so we return a copy instead.
	ret := *version
	return &ret, resp, nil
}
