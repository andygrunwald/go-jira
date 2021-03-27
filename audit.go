package jira

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

// AuditService interacts with the Jira audit Records to Add/Search entries
//
// Jira API docs: https://docs.atlassian.com/software/jira/docs/api/REST/8.5.13/#api/2/auditing
type AuditService struct{ client *Client }

type AuditSearchOptionsScheme struct {
	Filter     string
	From       time.Time
	To         time.Time
	ProjectIDs []string
	UserIDs    []string
}

const DateFormatJira = "2006-01-02T15:04:05.999-0700"

// Search methods find the auditing record filtering by date, project associated, entity and user who triggered the action
// Docs: https://docs.atlassian.com/software/jira/docs/api/REST/8.5.13/#api/2/auditing-getRecords
func (a *AuditService) Search(ctx context.Context, opts *AuditSearchOptionsScheme, offset, limit int) (result *AuditRecordScheme, response *Response, err error) {

	if ctx == nil {
		return nil, nil, fmt.Errorf("error!, please provide a valid ctx value")
	}

	if limit == 0 {
		return nil, nil, fmt.Errorf("error!, the response struct won't return any records")
	}

	params := url.Values{}
	params.Add("offset", strconv.Itoa(offset))
	params.Add("limit", strconv.Itoa(limit))

	if opts != nil {

		if opts.Filter != "" {
			params.Add("filter", opts.Filter)
		}

		if !opts.From.IsZero() {
			params.Add("from", opts.From.Format(DateFormatJira))
		}

		if !opts.To.IsZero() {
			params.Add("to", opts.To.Format(DateFormatJira))
		}

		var projects string
		for index, value := range opts.ProjectIDs {

			if index == 0 {
				projects = value
				continue
			}

			projects += "," + value
		}

		if len(projects) != 0 {
			params.Add("projectIds", projects)
		}

		var users string
		for index, value := range opts.UserIDs {

			if index == 0 {
				users = value
				continue
			}

			users += "," + value
		}

		if len(users) != 0 {
			params.Add("userIds", users)
		}

	}

	var endpoint = fmt.Sprintf("rest/api/2/auditing/record?%v", params.Encode())

	request, err := a.client.NewRequestWithContext(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	result = new(AuditRecordScheme)
	response, err = a.client.Do(request, result)
	if err != nil {
		return nil, nil, NewJiraError(response, err)
	}

	return
}

type AuditRecordScheme struct {
	Offset  int `json:"offset"`
	Limit   int `json:"limit"`
	Total   int `json:"total"`
	Records []struct {
		ID            int    `json:"id"`
		Summary       string `json:"summary"`
		RemoteAddress string `json:"remoteAddress"`
		AuthorKey     string `json:"authorKey"`
		Created       string `json:"created"`
		Category      string `json:"category"`
		EventSource   string `json:"eventSource"`
		ObjectItem    struct {
			ID       string `json:"id"`
			Name     string `json:"name"`
			TypeName string `json:"typeName"`
		} `json:"objectItem"`
		ChangedValues []struct {
			FieldName string `json:"fieldName"`
			ChangedTo string `json:"changedTo"`
		} `json:"changedValues,omitempty"`
		AssociatedItems []struct {
			ID       string `json:"id"`
			Name     string `json:"name"`
			TypeName string `json:"typeName"`
		} `json:"associatedItems"`
	} `json:"records"`
}
