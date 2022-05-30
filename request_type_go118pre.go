//go:build !go1.18
// +build !go1.18

package jira

// RequestTypeService handles Requesttypes for the Jira instance / API.
//
// Jira API docs: https://developer.atlassian.com/cloud/jira/service-desk/rest/api-group-requesttype/
type RequestTypeService struct {
	client *Client
}
