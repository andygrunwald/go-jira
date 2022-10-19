package onpremise

import (
	"context"
	"fmt"
	"net/http"
)

// OrganizationService handles Organizations for the Jira instance / API.
//
// Jira API docs: https://developer.atlassian.com/cloud/jira/service-desk/rest/api-group-organization/
type OrganizationService service

// OrganizationCreationDTO is DTO for creat organization API
type OrganizationCreationDTO struct {
	Name string `json:"name,omitempty" structs:"name,omitempty"`
}

// SelfLink Stores REST API URL to the organization.
type SelfLink struct {
	Self string `json:"self,omitempty" structs:"self,omitempty"`
}

// Organization contains Organization data
type Organization struct {
	ID    string    `json:"id,omitempty" structs:"id,omitempty"`
	Name  string    `json:"name,omitempty" structs:"name,omitempty"`
	Links *SelfLink `json:"_links,omitempty" structs:"_links,omitempty"`
}

// OrganizationUsersDTO contains organization user ids
type OrganizationUsersDTO struct {
	AccountIds []string `json:"accountIds,omitempty" structs:"accountIds,omitempty"`
}

// PagedDTO is response of a paged list
type PagedDTO struct {
	Size       int           `json:"size,omitempty" structs:"size,omitempty"`
	Start      int           `json:"start,omitempty" structs:"start,omitempty"`
	Limit      int           `limit:"size,omitempty" structs:"limit,omitempty"`
	IsLastPage bool          `json:"isLastPage,omitempty" structs:"isLastPage,omitempty"`
	Values     []interface{} `values:"isLastPage,omitempty" structs:"values,omitempty"`
	Expands    []string      `json:"_expands,omitempty" structs:"_expands,omitempty"`
}

// PropertyKey contains Property key details.
type PropertyKey struct {
	Self string `json:"self,omitempty" structs:"self,omitempty"`
	Key  string `json:"key,omitempty" structs:"key,omitempty"`
}

// PropertyKeys contains an array of PropertyKey
type PropertyKeys struct {
	Keys []PropertyKey `json:"keys,omitempty" structs:"keys,omitempty"`
}

// GetAllOrganizations returns a list of organizations in
// the Jira Service Management instance.
// Use this method when you want to present a list
// of organizations or want to locate an organization
// by name.
//
// Jira API docs: https://developer.atlassian.com/cloud/jira/service-desk/rest/api-group-organization/#api-group-organization
//
// TODO Double check this method if this works as expected, is using the latest API and the response is complete
// This double check effort is done for v2 - Remove this two lines if this is completed.
func (s *OrganizationService) GetAllOrganizations(ctx context.Context, start int, limit int, accountID string) (*PagedDTO, *Response, error) {
	apiEndPoint := fmt.Sprintf("rest/servicedeskapi/organization?start=%d&limit=%d", start, limit)
	if accountID != "" {
		apiEndPoint += fmt.Sprintf("&accountId=%s", accountID)
	}

	req, err := s.client.NewRequest(ctx, http.MethodGet, apiEndPoint, nil)
	req.Header.Set("Accept", "application/json")

	if err != nil {
		return nil, nil, err
	}

	v := new(PagedDTO)
	resp, err := s.client.Do(req, v)
	if err != nil {
		jerr := NewJiraError(resp, err)
		return nil, resp, jerr
	}

	return v, resp, nil
}

// CreateOrganization creates an organization by
// passing the name of the organization.
//
// Jira API docs: https://developer.atlassian.com/cloud/jira/service-desk/rest/api-group-organization/#api-rest-servicedeskapi-organization-post
//
// TODO Double check this method if this works as expected, is using the latest API and the response is complete
// This double check effort is done for v2 - Remove this two lines if this is completed.
func (s *OrganizationService) CreateOrganization(ctx context.Context, name string) (*Organization, *Response, error) {
	apiEndPoint := "rest/servicedeskapi/organization"

	organization := OrganizationCreationDTO{
		Name: name,
	}

	req, err := s.client.NewRequest(ctx, http.MethodPost, apiEndPoint, organization)
	req.Header.Set("Accept", "application/json")

	if err != nil {
		return nil, nil, err
	}

	o := new(Organization)
	resp, err := s.client.Do(req, &o)
	if err != nil {
		jerr := NewJiraError(resp, err)
		return nil, resp, jerr
	}

	return o, resp, nil
}

// GetOrganization returns details of an
// organization. Use this method to get organization
// details whenever your application component is
// passed an organization ID but needs to display
// other organization details.
//
// Jira API docs: https://developer.atlassian.com/cloud/jira/service-desk/rest/api-group-organization/#api-rest-servicedeskapi-organization-organizationid-get
//
// TODO Double check this method if this works as expected, is using the latest API and the response is complete
// This double check effort is done for v2 - Remove this two lines if this is completed.
func (s *OrganizationService) GetOrganization(ctx context.Context, organizationID int) (*Organization, *Response, error) {
	apiEndPoint := fmt.Sprintf("rest/servicedeskapi/organization/%d", organizationID)

	req, err := s.client.NewRequest(ctx, http.MethodGet, apiEndPoint, nil)
	req.Header.Set("Accept", "application/json")

	if err != nil {
		return nil, nil, err
	}

	o := new(Organization)
	resp, err := s.client.Do(req, &o)
	if err != nil {
		jerr := NewJiraError(resp, err)
		return nil, resp, jerr
	}

	return o, resp, nil
}

// DeleteOrganization deletes an organization. Note that
// the organization is deleted regardless
// of other associations it may have.
// For example, associations with service desks.
//
// Jira API docs: https://developer.atlassian.com/cloud/jira/service-desk/rest/api-group-organization/#api-rest-servicedeskapi-organization-organizationid-delete
// Caller must close resp.Body
//
// TODO Double check this method if this works as expected, is using the latest API and the response is complete
// This double check effort is done for v2 - Remove this two lines if this is completed.
func (s *OrganizationService) DeleteOrganization(ctx context.Context, organizationID int) (*Response, error) {
	apiEndPoint := fmt.Sprintf("rest/servicedeskapi/organization/%d", organizationID)

	req, err := s.client.NewRequest(ctx, http.MethodDelete, apiEndPoint, nil)

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

// GetPropertiesKeys returns the keys of
// all properties for an organization. Use this resource
// when you need to find out what additional properties
// items have been added to an organization.
//
// https://developer.atlassian.com/cloud/jira/service-desk/rest/api-group-organization/#api-rest-servicedeskapi-organization-organizationid-property-get
//
// TODO Double check this method if this works as expected, is using the latest API and the response is complete
// This double check effort is done for v2 - Remove this two lines if this is completed.
func (s *OrganizationService) GetPropertiesKeys(ctx context.Context, organizationID int) (*PropertyKeys, *Response, error) {
	apiEndPoint := fmt.Sprintf("rest/servicedeskapi/organization/%d/property", organizationID)

	req, err := s.client.NewRequest(ctx, http.MethodGet, apiEndPoint, nil)
	req.Header.Set("Accept", "application/json")

	if err != nil {
		return nil, nil, err
	}

	pk := new(PropertyKeys)
	resp, err := s.client.Do(req, &pk)
	if err != nil {
		jerr := NewJiraError(resp, err)
		return nil, resp, jerr
	}

	return pk, resp, nil
}

// GetProperty returns the value of a property
// from an organization. Use this method to obtain the JSON
// content for an organization's property.
//
// https://developer.atlassian.com/cloud/jira/service-desk/rest/api-group-organization/#api-rest-servicedeskapi-organization-organizationid-property-propertykey-get
//
// TODO Double check this method if this works as expected, is using the latest API and the response is complete
// This double check effort is done for v2 - Remove this two lines if this is completed.
func (s *OrganizationService) GetProperty(ctx context.Context, organizationID int, propertyKey string) (*EntityProperty, *Response, error) {
	apiEndPoint := fmt.Sprintf("rest/servicedeskapi/organization/%d/property/%s", organizationID, propertyKey)

	req, err := s.client.NewRequest(ctx, http.MethodGet, apiEndPoint, nil)
	req.Header.Set("Accept", "application/json")

	if err != nil {
		return nil, nil, err
	}

	ep := new(EntityProperty)
	resp, err := s.client.Do(req, &ep)
	if err != nil {
		jerr := NewJiraError(resp, err)
		return nil, resp, jerr
	}

	return ep, resp, nil
}

// SetProperty sets the value of a
// property for an organization. Use this
// resource to store custom data against an organization.
//
// https://developer.atlassian.com/cloud/jira/service-desk/rest/api-group-organization/#api-rest-servicedeskapi-organization-organizationid-property-propertykey-put
// Caller must close resp.Body
//
// TODO Double check this method if this works as expected, is using the latest API and the response is complete
// This double check effort is done for v2 - Remove this two lines if this is completed.
func (s *OrganizationService) SetProperty(ctx context.Context, organizationID int, propertyKey string) (*Response, error) {
	apiEndPoint := fmt.Sprintf("rest/servicedeskapi/organization/%d/property/%s", organizationID, propertyKey)

	req, err := s.client.NewRequest(ctx, http.MethodPut, apiEndPoint, nil)
	req.Header.Set("Accept", "application/json")

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

// DeleteProperty removes a property from an organization.
//
// https://developer.atlassian.com/cloud/jira/service-desk/rest/api-group-organization/#api-rest-servicedeskapi-organization-organizationid-property-propertykey-delete
// Caller must close resp.Body
//
// TODO Double check this method if this works as expected, is using the latest API and the response is complete
// This double check effort is done for v2 - Remove this two lines if this is completed.
func (s *OrganizationService) DeleteProperty(ctx context.Context, organizationID int, propertyKey string) (*Response, error) {
	apiEndPoint := fmt.Sprintf("rest/servicedeskapi/organization/%d/property/%s", organizationID, propertyKey)

	req, err := s.client.NewRequest(ctx, http.MethodDelete, apiEndPoint, nil)
	req.Header.Set("Accept", "application/json")

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

// GetUsers returns all the users
// associated with an organization. Use this
// method where you want to provide a list of
// users for an organization or determine if
// a user is associated with an organization.
//
// https://developer.atlassian.com/cloud/jira/service-desk/rest/api-group-organization/#api-rest-servicedeskapi-organization-organizationid-user-get
//
// TODO Double check this method if this works as expected, is using the latest API and the response is complete
// This double check effort is done for v2 - Remove this two lines if this is completed.
func (s *OrganizationService) GetUsers(ctx context.Context, organizationID int, start int, limit int) (*PagedDTO, *Response, error) {
	apiEndPoint := fmt.Sprintf("rest/servicedeskapi/organization/%d/user?start=%d&limit=%d", organizationID, start, limit)

	req, err := s.client.NewRequest(ctx, http.MethodGet, apiEndPoint, nil)
	req.Header.Set("Accept", "application/json")

	if err != nil {
		return nil, nil, err
	}

	users := new(PagedDTO)
	resp, err := s.client.Do(req, &users)
	if err != nil {
		jerr := NewJiraError(resp, err)
		return nil, resp, jerr
	}

	return users, resp, nil
}

// AddUsers adds users to an organization.
//
// https://developer.atlassian.com/cloud/jira/service-desk/rest/api-group-organization/#api-rest-servicedeskapi-organization-organizationid-user-post
// Caller must close resp.Body
//
// TODO Double check this method if this works as expected, is using the latest API and the response is complete
// This double check effort is done for v2 - Remove this two lines if this is completed.
func (s *OrganizationService) AddUsers(ctx context.Context, organizationID int, users OrganizationUsersDTO) (*Response, error) {
	apiEndPoint := fmt.Sprintf("rest/servicedeskapi/organization/%d/user", organizationID)

	req, err := s.client.NewRequest(ctx, http.MethodPost, apiEndPoint, users)

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

// RemoveUsers removes users from an organization.
//
// https://developer.atlassian.com/cloud/jira/service-desk/rest/api-group-organization/#api-rest-servicedeskapi-organization-organizationid-user-delete
// Caller must close resp.Body
//
// TODO Double check this method if this works as expected, is using the latest API and the response is complete
// This double check effort is done for v2 - Remove this two lines if this is completed.
func (s *OrganizationService) RemoveUsers(ctx context.Context, organizationID int, users OrganizationUsersDTO) (*Response, error) {
	apiEndPoint := fmt.Sprintf("rest/servicedeskapi/organization/%d/user", organizationID)

	req, err := s.client.NewRequest(ctx, http.MethodDelete, apiEndPoint, nil)
	req.Header.Set("Accept", "application/json")

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
