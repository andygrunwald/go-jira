package cloud

import (
	"context"
	"fmt"
	"net/http"

	"github.com/andygrunwald/go-jira/v2/cloud/models"
	"github.com/andygrunwald/go-jira/v2/cloud/models/servicedesk"
	"github.com/google/go-querystring/query"
)

// GetRequestTypes returns all
// request types of a given service desk id.
//
// Jira API docs: https://developer.atlassian.com/cloud/jira/service-desk/rest/api-group-servicedesk/#api-rest-servicedeskapi-servicedesk-servicedeskid-requesttype-get
func (s *ServiceDeskService) GetRequestTypes(ctx context.Context, serviceDeskID int, options *servicedesk.RequestTypeOptions) (*models.PagedDTOT[servicedesk.RequestType], *Response, error) {
	apiEndPoint := fmt.Sprintf("rest/servicedeskapi/servicedesk/%d/requesttype", serviceDeskID)

	req, err := s.client.NewRequest(ctx, http.MethodGet, apiEndPoint, nil)
	if err != nil {
		return nil, nil, err
	}

	req.Header.Set("Accept", "application/json")

	if options != nil {
		q, err := query.Values(options)
		if err != nil {
			return nil, nil, err
		}

		req.URL.RawQuery = q.Encode()
	}

	o := new(models.PagedDTOT[servicedesk.RequestType])
	resp, err := s.client.Do(req, &o)
	if err != nil {
		return nil, resp, NewJiraError(resp, err)
	}

	return o, resp, nil
}

// GetRequestTypeFields returns fields for
// the provided request type id.
//
// Jira API docs: https://developer.atlassian.com/cloud/jira/service-desk/rest/api-group-servicedesk/#api-rest-servicedeskapi-servicedesk-servicedeskid-requesttype-requesttypeid-field-get
func (s *ServiceDeskService) GetRequestTypeFields(ctx context.Context, serviceDeskID, requestTypeID int) (*servicedesk.CustomerRequestCreateMeta, *Response, error) {
	apiEndPoint := fmt.Sprintf("rest/servicedeskapi/servicedesk/%d/requesttype/%d/field", serviceDeskID, requestTypeID)

	req, err := s.client.NewRequest(ctx, http.MethodGet, apiEndPoint, nil)
	if err != nil {
		return nil, nil, err
	}

	req.Header.Set("Accept", "application/json")

	o := new(servicedesk.CustomerRequestCreateMeta)
	resp, err := s.client.Do(req, &o)
	if err != nil {
		return nil, resp, NewJiraError(resp, err)
	}

	return o, resp, nil
}

// GetRequestTypeGroups returns groups for
// a given service desk id.
//
// Jira API docs: https://developer.atlassian.com/cloud/jira/service-desk/rest/api-group-servicedesk/#api-rest-servicedeskapi-servicedesk-servicedeskid-requesttypegroup-get
func (s *ServiceDeskService) GetRequestTypeGroups(ctx context.Context, serviceDeskID int) (*models.PagedDTOT[servicedesk.RequestTypeGroup], *Response, error) {
	apiEndPoint := fmt.Sprintf("rest/servicedeskapi/servicedesk/%d/requesttypegroup", serviceDeskID)

	req, err := s.client.NewRequest(ctx, http.MethodGet, apiEndPoint, nil)
	if err != nil {
		return nil, nil, err
	}

	req.Header.Set("Accept", "application/json")

	o := new(models.PagedDTOT[servicedesk.RequestTypeGroup])
	resp, err := s.client.Do(req, &o)
	if err != nil {
		return nil, resp, NewJiraError(resp, err)
	}

	return o, resp, nil
}
