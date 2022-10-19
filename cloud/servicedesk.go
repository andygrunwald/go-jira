package cloud

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/google/go-querystring/query"
)

// ServiceDeskService handles ServiceDesk for the Jira instance / API.
type ServiceDeskService service

// ServiceDeskOrganizationDTO is a DTO for ServiceDesk organizations
type ServiceDeskOrganizationDTO struct {
	OrganizationID int `json:"organizationId,omitempty" structs:"organizationId,omitempty"`
}

// GetOrganizations returns a list of
// all organizations associated with a service desk.
//
// https://developer.atlassian.com/cloud/jira/service-desk/rest/api-group-organization/#api-rest-servicedeskapi-servicedesk-servicedeskid-organization-get
//
// TODO Double check this method if this works as expected, is using the latest API and the response is complete
// This double check effort is done for v2 - Remove this two lines if this is completed.
func (s *ServiceDeskService) GetOrganizations(ctx context.Context, serviceDeskID interface{}, start int, limit int, accountID string) (*PagedDTO, *Response, error) {
	apiEndPoint := fmt.Sprintf("rest/servicedeskapi/servicedesk/%v/organization?start=%d&limit=%d", serviceDeskID, start, limit)
	if accountID != "" {
		apiEndPoint += fmt.Sprintf("&accountId=%s", accountID)
	}

	req, err := s.client.NewRequest(ctx, http.MethodGet, apiEndPoint, nil)
	req.Header.Set("Accept", "application/json")

	if err != nil {
		return nil, nil, err
	}

	orgs := new(PagedDTO)
	resp, err := s.client.Do(req, &orgs)
	if err != nil {
		jerr := NewJiraError(resp, err)
		return nil, resp, jerr
	}

	return orgs, resp, nil
}

// AddOrganization adds an organization to
// a service desk. If the organization ID is already
// associated with the service desk, no change is made
// and the resource returns a 204 success code.
//
// https://developer.atlassian.com/cloud/jira/service-desk/rest/api-group-organization/#api-rest-servicedeskapi-servicedesk-servicedeskid-organization-post
// Caller must close resp.Body
//
// TODO Double check this method if this works as expected, is using the latest API and the response is complete
// This double check effort is done for v2 - Remove this two lines if this is completed.
func (s *ServiceDeskService) AddOrganization(ctx context.Context, serviceDeskID interface{}, organizationID int) (*Response, error) {
	apiEndPoint := fmt.Sprintf("rest/servicedeskapi/servicedesk/%v/organization", serviceDeskID)

	organization := ServiceDeskOrganizationDTO{
		OrganizationID: organizationID,
	}

	req, err := s.client.NewRequest(ctx, http.MethodPost, apiEndPoint, organization)

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

// RemoveOrganization removes an organization
// from a service desk. If the organization ID does not
// match an organization associated with the service desk,
// no change is made and the resource returns a 204 success code.
//
// https://developer.atlassian.com/cloud/jira/service-desk/rest/api-group-organization/#api-rest-servicedeskapi-servicedesk-servicedeskid-organization-delete
// Caller must close resp.Body
//
// TODO Double check this method if this works as expected, is using the latest API and the response is complete
// This double check effort is done for v2 - Remove this two lines if this is completed.
func (s *ServiceDeskService) RemoveOrganization(ctx context.Context, serviceDeskID interface{}, organizationID int) (*Response, error) {
	apiEndPoint := fmt.Sprintf("rest/servicedeskapi/servicedesk/%v/organization", serviceDeskID)

	organization := ServiceDeskOrganizationDTO{
		OrganizationID: organizationID,
	}

	req, err := s.client.NewRequest(ctx, http.MethodDelete, apiEndPoint, organization)

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

// AddCustomers adds customers to the given service desk.
//
// https://developer.atlassian.com/cloud/jira/service-desk/rest/api-group-servicedesk/#api-rest-servicedeskapi-servicedesk-servicedeskid-customer-post
//
// TODO Double check this method if this works as expected, is using the latest API and the response is complete
// This double check effort is done for v2 - Remove this two lines if this is completed.
func (s *ServiceDeskService) AddCustomers(ctx context.Context, serviceDeskID interface{}, acountIDs ...string) (*Response, error) {
	apiEndpoint := fmt.Sprintf("rest/servicedeskapi/servicedesk/%v/customer", serviceDeskID)

	payload := struct {
		AccountIDs []string `json:"accountIds"`
	}{
		AccountIDs: acountIDs,
	}
	req, err := s.client.NewRequest(ctx, http.MethodPost, apiEndpoint, payload)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req, nil)
	if err != nil {
		return resp, NewJiraError(resp, err)
	}

	defer resp.Body.Close()
	_, _ = io.Copy(io.Discard, resp.Body)

	return resp, nil
}

// RemoveCustomers removes customers to the given service desk.
//
// https://developer.atlassian.com/cloud/jira/service-desk/rest/api-group-servicedesk/#api-rest-servicedeskapi-servicedesk-servicedeskid-customer-delete
//
// TODO Double check this method if this works as expected, is using the latest API and the response is complete
// This double check effort is done for v2 - Remove this two lines if this is completed.
func (s *ServiceDeskService) RemoveCustomers(ctx context.Context, serviceDeskID interface{}, acountIDs ...string) (*Response, error) {
	apiEndpoint := fmt.Sprintf("rest/servicedeskapi/servicedesk/%v/customer", serviceDeskID)

	payload := struct {
		AccountIDs []string `json:"accountIDs"`
	}{
		AccountIDs: acountIDs,
	}
	req, err := s.client.NewRequest(ctx, http.MethodDelete, apiEndpoint, payload)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req, nil)
	if err != nil {
		return resp, NewJiraError(resp, err)
	}

	defer resp.Body.Close()
	_, _ = io.Copy(io.Discard, resp.Body)

	return resp, nil
}

// ListCustomers lists customers for a ServiceDesk.
//
// https://developer.atlassian.com/cloud/jira/service-desk/rest/api-group-servicedesk/#api-rest-servicedeskapi-servicedesk-servicedeskid-customer-get
//
// TODO Double check this method if this works as expected, is using the latest API and the response is complete
// This double check effort is done for v2 - Remove this two lines if this is completed.
func (s *ServiceDeskService) ListCustomers(ctx context.Context, serviceDeskID interface{}, options *CustomerListOptions) (*CustomerList, *Response, error) {
	apiEndpoint := fmt.Sprintf("rest/servicedeskapi/servicedesk/%v/customer", serviceDeskID)
	req, err := s.client.NewRequest(ctx, http.MethodGet, apiEndpoint, nil)
	if err != nil {
		return nil, nil, err
	}

	// this is an experiemntal endpoint
	req.Header.Set("X-ExperimentalApi", "opt-in")

	if options != nil {
		q, err := query.Values(options)
		if err != nil {
			return nil, nil, err
		}
		req.URL.RawQuery = q.Encode()
	}

	resp, err := s.client.Do(req, nil)
	if err != nil {
		return nil, resp, NewJiraError(resp, err)
	}
	defer resp.Body.Close()

	customerList := new(CustomerList)
	if err := json.NewDecoder(resp.Body).Decode(customerList); err != nil {
		return nil, resp, fmt.Errorf("could not unmarshall the data into struct")
	}

	return customerList, resp, nil
}
