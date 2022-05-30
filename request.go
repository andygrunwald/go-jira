package jira

import (
	"context"
	"fmt"
	"net/http"
)

// RequestService handles ServiceDesk customer requests for the Jira instance / API.
type RequestService struct {
	client *Client
}

// Request represents a ServiceDesk customer request.
type Request struct {
	IssueID       string              `json:"issueId,omitempty" structs:"issueId,omitempty"`
	IssueKey      string              `json:"issueKey,omitempty" structs:"issueKey,omitempty"`
	TypeID        string              `json:"requestTypeId,omitempty" structs:"requestTypeId,omitempty"`
	ServiceDeskID string              `json:"serviceDeskId,omitempty" structs:"serviceDeskId,omitempty"`
	Reporter      *Customer           `json:"reporter,omitempty" structs:"reporter,omitempty"`
	FieldValues   []RequestFieldValue `json:"requestFieldValues,omitempty" structs:"requestFieldValues,omitempty"`
	Status        *RequestStatus      `json:"currentStatus,omitempty" structs:"currentStatus,omitempty"`
	Links         *SelfLink           `json:"_links,omitempty" structs:"_links,omitempty"`
	Expands       []string            `json:"_expands,omitempty" structs:"_expands,omitempty"`
}

// RequestFieldValue is a request field.
type RequestFieldValue struct {
	FieldID string `json:"fieldId,omitempty" structs:"fieldId,omitempty"`
	Label   string `json:"label,omitempty" structs:"label,omitempty"`
	Value   any    `json:"value,omitempty" structs:"value,omitempty"`
}

// RequestDate is the date format used in requests.
type RequestDate struct {
	ISO8601  string `json:"iso8601,omitempty" structs:"iso8601,omitempty"`
	Jira     string `json:"jira,omitempty" structs:"jira,omitempty"`
	Friendly string `json:"friendly,omitempty" structs:"friendly,omitempty"`
	Epoch    int64  `json:"epoch,omitempty" structs:"epoch,omitempty"`
}

// RequestStatus is the status for a request.
type RequestStatus struct {
	Status   string
	Category string
	Date     RequestDate
}

// RequestComment is a comment for a request.
type RequestComment struct {
	ID      string       `json:"id,omitempty" structs:"id,omitempty"`
	Body    string       `json:"body,omitempty" structs:"body,omitempty"`
	Public  bool         `json:"public" structs:"public"`
	Author  *Customer    `json:"author,omitempty" structs:"author,omitempty"`
	Created *RequestDate `json:"created,omitempty" structs:"created,omitempty"`
	Links   *SelfLink    `json:"_links,omitempty" structs:"_links,omitempty"`
	Expands []string     `json:"_expands,omitempty" structs:"_expands,omitempty"`
}

// CreateRequest represents a ServiceDesk customer request.
type CreateRequest struct {
	ServiceDeskID string         `json:"serviceDeskId,omitempty" structs:"serviceDeskId,omitempty"`
	TypeID        string         `json:"requestTypeId,omitempty" structs:"requestTypeId,omitempty"`
	FieldValues   map[string]any `json:"requestFieldValues,omitempty" structs:"requestFieldValues,omitempty"`
	Participants  []string       `json:"requestParticipants,omitempty" structs:"requestParticipants,omitempty"`
	Requester     string         `json:"raiseOnBehalfOf,omitempty" structs:"raiseOnBehalfOf,omitempty"`
}

// CreateWithContext creates a new request.
//
// https://developer.atlassian.com/cloud/jira/service-desk/rest/api-group-request/#api-rest-servicedeskapi-request-post
func (r *RequestService) CreateWithContext(ctx context.Context, request *CreateRequest) (*Request, *Response, error) {
	apiEndpoint := "rest/servicedeskapi/request"

	req, err := r.client.NewRequestWithContext(ctx, http.MethodPost, apiEndpoint, request)
	if err != nil {
		return nil, nil, err
	}

	responseRequest := new(Request)
	resp, err := r.client.Do(req, responseRequest)
	if err != nil {
		return nil, resp, NewJiraError(resp, err)
	}

	return responseRequest, resp, nil
}

// Create wraps CreateWithContext using the background context.
func (r *RequestService) Create(request *CreateRequest) (*Request, *Response, error) {
	return r.CreateWithContext(context.Background(), request)
}

// CreateCommentWithContext creates a comment on a request.
//
// https://developer.atlassian.com/cloud/jira/service-desk/rest/api-group-request/#api-rest-servicedeskapi-request-issueidorkey-comment-post
func (r *RequestService) CreateCommentWithContext(ctx context.Context, issueIDOrKey string, comment *RequestComment) (*RequestComment, *Response, error) {
	apiEndpoint := fmt.Sprintf("rest/servicedeskapi/request/%v/comment", issueIDOrKey)

	req, err := r.client.NewRequestWithContext(ctx, "POST", apiEndpoint, comment)
	if err != nil {
		return nil, nil, err
	}

	responseComment := new(RequestComment)
	resp, err := r.client.Do(req, responseComment)
	if err != nil {
		return nil, resp, NewJiraError(resp, err)
	}

	return responseComment, resp, nil
}

// CreateComment wraps CreateCommentWithContext using the background context.
func (r *RequestService) CreateComment(issueIDOrKey string, comment *RequestComment) (*RequestComment, *Response, error) {
	return r.CreateCommentWithContext(context.Background(), issueIDOrKey, comment)
}
