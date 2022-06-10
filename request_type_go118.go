//go:build go1.18
// +build go1.18

package jira

import (
	"context"
	"fmt"
	"github.com/google/go-querystring/query"
	"net/http"
)

// RequestTypeService handles Requesttypes for the Jira instance / API.
//
// Jira API docs: https://developer.atlassian.com/cloud/jira/service-desk/rest/api-group-requesttype/
type RequestTypeService struct {
	client *Client
}

// PagedDTOT is response of a paged list with generic support for values
type PagedDTOT[T any] struct {
	Size       int       `json:"size" structs:"size"`
	Start      int       `json:"start" structs:"start"`
	Limit      int       `limit:"limit,omitempty" structs:"limit,omitempty"`
	IsLastPage bool      `json:"isLastPage,omitempty" structs:"isLastPage,omitempty"`
	Values     []T       `json:"values,omitempty" structs:"values,omitempty"`
	Expands    []string  `json:"_expands,omitempty" structs:"_expands,omitempty"`
	Links      *SelfLink `json:"_links,omitempty" structs:"_links,omitempty"`
}

// RequestType contains RequestType data
type RequestType struct {
	ID            string           `json:"id,omitempty" structs:"id,omitempty"`
	Name          string           `json:"name,omitempty" structs:"name,omitempty"`
	Description   string           `json:"description,omitempty" structs:"description,omitempty"`
	HelpText      string           `json:"helpText,omitempty" structs:"helpText,omitempty"`
	IssueTypeID   string           `json:"issueTypeId,omitempty" structs:"issueTypeId,omitempty"`
	ServiceDeskID string           `json:"serviceDeskId,omitempty" structs:"serviceDeskId,omitempty"`
	GroupIDs      []string         `json:"groupIds,omitempty" structs:"groupIds,omitempty"`
	Icon          *RequestTypeIcon `json:"icon,omitempty" structs:"icon,omitempty"`
	Practice      string           `json:"practice,omitempty" structs:"practice,omitempty"`
	Expands       []string         `json:"_expands,omitempty" structs:"_expands,omitempty"`
	Links         *SelfLink        `json:"_links,omitempty" structs:"_links,omitempty"`
}

// RequestTypeIcon contains RequestType icon data
type RequestTypeIcon struct {
	ID    string               `json:"id,omitempty" structs:"id,omitempty"`
	Links *RequestTypeIconLink `json:"_links,omitempty" structs:"_links,omitempty"`
}

// RequestTypeIconLink contains RequestTypeIcon link data
type RequestTypeIconLink struct {
	IconURLs map[string]string `json:"iconUrls,omitempty" structs:"iconUrls,omitempty"`
}

// RequestTypeOptions is the query options for listing request types.
type RequestTypeOptions struct {
	SearchQuery string `url:"searchQuery,omitempty" query:"searchQuery"`
	GroupID     int    `url:"groupId,omitempty" query:"groupId"`
}

// GetWithContext returns all
// request types of a given service desk id.
//
// Jira API docs: https://developer.atlassian.com/cloud/jira/service-desk/rest/api-group-servicedesk/#api-rest-servicedeskapi-servicedesk-servicedeskid-requesttype-get
func (s *RequestTypeService) GetWithContext(ctx context.Context, serviceDeskID int, options *RequestTypeOptions) (*PagedDTOT[RequestType], *Response, error) {
	apiEndPoint := fmt.Sprintf("rest/servicedeskapi/servicedesk/%d/requesttype", serviceDeskID)

	req, err := s.client.NewRequestWithContext(ctx, http.MethodGet, apiEndPoint, nil)
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

	o := new(PagedDTOT[RequestType])
	resp, err := s.client.Do(req, &o)
	if err != nil {
		return nil, resp, NewJiraError(resp, err)
	}

	return o, resp, nil
}

// Get wraps GetWithContext using the background context.
func (s *RequestTypeService) Get(serviceDeskID int, options *RequestTypeOptions) (*PagedDTOT[RequestType], *Response, error) {
	return s.GetWithContext(context.Background(), serviceDeskID, options)
}

type CustomerRequestCreateMeta struct {
	RequestTypeFields         []RequestTypeField `json:"requestTypeFields,omitempty" structs:"requestTypeFields,omitempty"`
	CanRaiseOnBehalfOf        bool               `json:"canRaiseOnBehalfOf,omitempty" structs:"canRaiseOnBehalfOf,omitempty"`
	CanAddRequestParticipants bool               `json:"canAddRequestParticipants,omitempty" structs:"canAddRequestParticipants,omitempty"`
}

type RequestTypeField struct {
	FieldID       string                  `json:"fieldId,omitempty" structs:"fieldId,omitempty"`
	Name          string                  `json:"name,omitempty" structs:"name,omitempty"`
	Description   string                  `json:"description,omitempty" structs:"description,omitempty"`
	Required      bool                    `json:"required" structs:"required"`
	DefaultValues []RequestTypeFieldValue `json:"defaultValues,omitempty" structs:"defaultValues,omitempty"`
	ValidValues   []RequestTypeFieldValue `json:"validValues,omitempty" structs:"validValues,omitempty"`
	PresetValues  []string                `json:"presetValues,omitempty" structs:"presetValues,omitempty"`
	JiraSchema    *JiraSchema             `json:"jiraSchema,omitempty" structs:"jiraSchema,omitempty"`
	Visible       bool                    `json:"visible" structs:"visible"`
}

type RequestTypeFieldValue struct {
	Value    string                  `json:"value,omitempty" structs:"value,omitempty"`
	Label    string                  `json:"label,omitempty" structs:"label,omitempty"`
	Children []RequestTypeFieldValue `json:"children,omitempty" structs:"children,omitempty"`
}

type JiraSchema struct {
	Type          string            `json:"type,omitempty" structs:"type,omitempty"`
	Items         string            `json:"items,omitempty" structs:"items,omitempty"`
	System        string            `json:"system,omitempty" structs:"system,omitempty"`
	Custom        string            `json:"custom,omitempty" structs:"custom,omitempty"`
	CustomID      int               `json:"customId,omitempty" structs:"customId,omitempty"`
	Configuration map[string]string `json:"configuration,omitempty" structs:"configuration,omitempty"`
}

// GetFieldsWithContext returns fields for
// the provided request type id.
//
// Jira API docs: https://developer.atlassian.com/cloud/jira/service-desk/rest/api-group-servicedesk/#api-rest-servicedeskapi-servicedesk-servicedeskid-requesttype-requesttypeid-field-get
func (s *RequestTypeService) GetFieldsWithContext(ctx context.Context, serviceDeskID, requestTypeID int) (*CustomerRequestCreateMeta, *Response, error) {
	apiEndPoint := fmt.Sprintf("rest/servicedeskapi/servicedesk/%d/requesttype/%d/field", serviceDeskID, requestTypeID)

	req, err := s.client.NewRequestWithContext(ctx, http.MethodGet, apiEndPoint, nil)
	req.Header.Set("Accept", "application/json")

	if err != nil {
		return nil, nil, err
	}

	o := new(CustomerRequestCreateMeta)
	resp, err := s.client.Do(req, &o)
	if err != nil {
		return nil, resp, NewJiraError(resp, err)
	}

	return o, resp, nil
}

// GetFields wraps GetFieldsWithContext using the background context.
func (s *RequestTypeService) GetFields(serviceDeskID, requestTypeID int) (*CustomerRequestCreateMeta, *Response, error) {
	return s.GetFieldsWithContext(context.Background(), serviceDeskID, requestTypeID)
}

type RequestTypeGroup struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// GetGroupsWithContext returns groups for
// a given service desk id.
//
// Jira API docs: https://developer.atlassian.com/cloud/jira/service-desk/rest/api-group-servicedesk/#api-rest-servicedeskapi-servicedesk-servicedeskid-requesttype-requesttypeid-field-get
func (s *RequestTypeService) GetGroupsWithContext(ctx context.Context, serviceDeskID int) (*PagedDTOT[RequestTypeGroup], *Response, error) {
	apiEndPoint := fmt.Sprintf("rest/servicedeskapi/servicedesk/%d/requesttypegroup", serviceDeskID)

	req, err := s.client.NewRequestWithContext(ctx, http.MethodGet, apiEndPoint, nil)
	req.Header.Set("Accept", "application/json")

	if err != nil {
		return nil, nil, err
	}

	o := new(PagedDTOT[RequestTypeGroup])
	resp, err := s.client.Do(req, &o)
	if err != nil {
		return nil, resp, NewJiraError(resp, err)
	}

	return o, resp, nil
}

// GetGroups wraps GetGroupsWithContext using the background context.
func (s *RequestTypeService) GetGroups(serviceDeskID int) (*PagedDTOT[RequestTypeGroup], *Response, error) {
	return s.GetGroupsWithContext(context.Background(), serviceDeskID)
}