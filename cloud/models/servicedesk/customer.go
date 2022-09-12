package servicedesk

import "github.com/mcl-de/go-jira/v2/cloud/models"

// Customer represents a ServiceDesk customer.
type Customer struct {
	AccountID    string           `json:"accountId,omitempty" structs:"accountId,omitempty"`
	Name         string           `json:"name,omitempty" structs:"name,omitempty"`
	Key          string           `json:"key,omitempty" structs:"key,omitempty"`
	EmailAddress string           `json:"emailAddress,omitempty" structs:"emailAddress,omitempty"`
	DisplayName  string           `json:"displayName,omitempty" structs:"displayName,omitempty"`
	Active       *bool            `json:"active,omitempty" structs:"active,omitempty"`
	TimeZone     string           `json:"timeZone,omitempty" structs:"timeZone,omitempty"`
	Links        *models.SelfLink `json:"_links,omitempty" structs:"_links,omitempty"`
}

// CustomerListOptions is the query options for listing customers.
type CustomerListOptions struct {
	Query string `url:"query,omitempty"`
	Start int    `url:"start,omitempty"`
	Limit int    `url:"limit,omitempty"`
}
