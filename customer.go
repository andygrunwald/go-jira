package jira

import (
	"context"
	"encoding/json"
	"fmt"
)

// CustomerService handles ServiceDesk customers for the Jira instance / API.
type CustomerService struct {
	client *Client
}

// Customer represents a ServiceDesk customer.
type Customer struct {
	AccountID    string    `json:"accountId,omitempty" structs:"accountId,omitempty"`
	Name         string    `json:"name,omitempty" structs:"name,omitempty"`
	Key          string    `json:"key,omitempty" structs:"key,omitempty"`
	EmailAddress string    `json:"emailAddress,omitempty" structs:"emailAddress,omitempty"`
	DisplayName  string    `json:"displayName,omitempty" structs:"displayName,omitempty"`
	Active       *bool     `json:"active,omitempty" structs:"active,omitempty"`
	TimeZone     string    `json:"timeZone,omitempty" structs:"timeZone,omitempty"`
	Links        *SelfLink `json:"_links,omitempty" structs:"_links,omitempty"`
}

// CustomerListOptions is the query options for listing customers.
type CustomerListOptions struct {
	Query string `url:"query,omitempty"`
	Start int    `url:"start,omitempty"`
	Limit int    `url:"limit,omitempty"`
}

// CustomerList is a page of customers.
type CustomerList struct {
	Values  []Customer `json:"values,omitempty" structs:"values,omitempty"`
	Start   int        `json:"start,omitempty" structs:"start,omitempty"`
	Limit   int        `json:"limit,omitempty" structs:"limit,omitempty"`
	IsLast  bool       `json:"isLastPage,omitempty" structs:"isLastPage,omitempty"`
	Expands []string   `json:"_expands,omitempty" structs:"_expands,omitempty"`
}

// CreateWithContext creates a ServiceDesk customer.
//
// https://developer.atlassian.com/cloud/jira/service-desk/rest/api-group-customer/#api-rest-servicedeskapi-customer-post
func (c *CustomerService) CreateWithContext(ctx context.Context, email, displayName string) (*Customer, *Response, error) {
	apiEndpoint := "rest/servicedeskapi/customer"

	payload := struct {
		Email       string `json:"email"`
		DisplayName string `json:"displayName"`
	}{
		Email:       email,
		DisplayName: displayName,
	}

	req, err := c.client.NewRequestWithContext(ctx, "POST", apiEndpoint, payload)
	if err != nil {
		return nil, nil, err
	}
	resp, err := c.client.Do(req, nil)
	if err != nil {
		return nil, resp, NewJiraError(resp, err)
	}
	defer resp.Body.Close()

	responseCustomer := new(Customer)
	if err := json.NewDecoder(resp.Body).Decode(responseCustomer); err != nil {
		return nil, resp, fmt.Errorf("could not unmarshall the data into struct")
	}

	return responseCustomer, resp, nil
}

// Create wraps CreateWithContext using the background context.
func (c *CustomerService) Create(email, displayName string) (*Customer, *Response, error) {
	return c.CreateWithContext(context.Background(), email, displayName)
}
