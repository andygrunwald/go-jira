package cloud

import (
	"context"
	"net/http"

	"github.com/mcl-de/go-jira/v2/cloud/models/servicedesk"
)

// CustomerService handles ServiceDesk customers for the Jira instance / API.
type CustomerService service

// Create creates a ServiceDesk customer.
//
// https://developer.atlassian.com/cloud/jira/service-desk/rest/api-group-customer/#api-rest-servicedeskapi-customer-post
func (c *CustomerService) Create(ctx context.Context, email, displayName string) (*servicedesk.Customer, *Response, error) {
	const apiEndpoint = "rest/servicedeskapi/customer"

	payload := struct {
		Email       string `json:"email"`
		DisplayName string `json:"displayName"`
	}{
		Email:       email,
		DisplayName: displayName,
	}

	req, err := c.client.NewRequest(ctx, http.MethodPost, apiEndpoint, payload)
	if err != nil {
		return nil, nil, err
	}

	responseCustomer := new(servicedesk.Customer)
	resp, err := c.client.Do(req, responseCustomer)
	if err != nil {
		return nil, resp, NewJiraError(resp, err)
	}

	return responseCustomer, resp, nil
}
