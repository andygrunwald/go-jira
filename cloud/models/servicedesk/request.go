package servicedesk

import (
	"github.com/andygrunwald/go-jira/v2/cloud/models"
)

type CreateRequest struct {
	ServiceDeskID string         `json:"serviceDeskId,omitempty" structs:"serviceDeskId,omitempty"`
	TypeID        string         `json:"requestTypeId,omitempty" structs:"requestTypeId,omitempty"`
	FieldValues   map[string]any `json:"requestFieldValues,omitempty" structs:"requestFieldValues,omitempty"`
	Participants  []string       `json:"requestParticipants,omitempty" structs:"requestParticipants,omitempty"`
	Requester     string         `json:"raiseOnBehalfOf,omitempty" structs:"raiseOnBehalfOf,omitempty"`
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
	Links         *models.SelfLink    `json:"_links,omitempty" structs:"_links,omitempty"`
	Expands       []string            `json:"_expands,omitempty" structs:"_expands,omitempty"`
}

// RequestFieldValue is a request field.
type RequestFieldValue struct {
	FieldID string `json:"fieldId,omitempty" structs:"fieldId,omitempty"`
	Label   string `json:"label,omitempty" structs:"label,omitempty"`
	Value   string `json:"value,omitempty" structs:"value,omitempty"`
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
	ID      string           `json:"id,omitempty" structs:"id,omitempty"`
	Body    string           `json:"body,omitempty" structs:"body,omitempty"`
	Public  bool             `json:"public" structs:"public"`
	Author  *Customer        `json:"author,omitempty" structs:"author,omitempty"`
	Created *RequestDate     `json:"created,omitempty" structs:"created,omitempty"`
	Links   *models.SelfLink `json:"_links,omitempty" structs:"_links,omitempty"`
	Expands []string         `json:"_expands,omitempty" structs:"_expands,omitempty"`
}

// RequestCommentListOptions is the query options for listing comments for a ServiceDesk request.
type RequestCommentListOptions struct {
	Public   *bool    `url:"public,omitempty" query:"public"`
	Internal *bool    `url:"internal,omitempty" query:"internal"`
	Expand   []string `url:"expand,omitempty" query:"expand"`
	Start    int      `url:"start,omitempty" query:"start"`
	Limit    int      `url:"limit,omitempty" query:"limit"`
}

type CreateRequestComment struct {
	Body   string `json:"body" structs:"body"`
	Public bool   `json:"public" structs:"public"`
}
