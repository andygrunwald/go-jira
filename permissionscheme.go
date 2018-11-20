package jira

import "fmt"

// PermissionSchemeService handles permissionschemes for the JIRA instance / API.
//
// JIRA API docs: https://docs.atlassian.com/jira/REST/latest/#api/3/permissionscheme
type PermissionSchemeService struct {
	client *Client
}
type PermissionSchemes struct {
	PermissionSchemes []PermissionScheme `json:"permissionSchemes" structs:"permissionSchemes"`
}

type Permission struct {
	ID     int    `json:"id" structs:"id"`
	Self   string `json:"expand" structs:"expand"`
	Holder Holder `json:"holder" structs:"holder"`
	Name   string `json:"permission" structs:"permission"`
}

type Holder struct {
	Type      string `json:"type" structs:"type"`
	Parameter string `json:"parameter" structs:"parameter"`
	Expand    string `json:"expand" structs:"expand"`
}

// GetList returns a list of all permission schemes
func (s *PermissionSchemeService) GetList() (*PermissionSchemes, *Response, error) {
	apiEndpoint := "/rest/api/3/permissionscheme"
	req, err := s.client.NewRequest("GET", apiEndpoint, nil)
	if err != nil {
		return nil, nil, err
	}

	pss := new(PermissionSchemes)
	resp, err := s.client.Do(req, &pss)
	if err != nil {
		jerr := NewJiraError(resp, err)
		return nil, resp, jerr
	}

	return pss, resp, nil
}

// Get returns a full representation of the permission scheme for the schemeID
func (s *PermissionSchemeService) Get(schemeID int) (*PermissionScheme, *Response, error) {
	apiEndpoint := fmt.Sprintf("/rest/api/3/permissionscheme/%d", schemeID)
	req, err := s.client.NewRequest("GET", apiEndpoint, nil)
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
