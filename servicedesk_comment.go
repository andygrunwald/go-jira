//go:build go1.18
// +build go1.18

package jira

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/google/go-querystring/query"
)

// RequestCommentListOptions is the query options for listing comments for a ServiceDesk request.
type RequestCommentListOptions struct {
	Public   *bool    `url:"public,omitempty"`
	Internal *bool    `url:"internal,omitempty"`
	Expand   []string `url:"expand,omitempty"`
	Start    int      `url:"start,omitempty"`
	Limit    int      `url:"limit,omitempty"`
}

// CommentDTO represents a comment to an request in ServiceDesk.
type CommentDTO struct {
	ID           string           `json:"id,omitempty" structs:"id,omitempty"`
	Body         string           `json:"body,omitempty" structs:"body,omitempty"`
	RenderedBody RenderedValueDTO `json:"renderedBody,omitempty" structs:"renderedBody,omitempty"`
	Author       User             `json:"author,omitempty" structs:"author,omitempty"`
	Created      string           `json:"created,omitempty" structs:"created,omitempty"`
	Attachments  Attachment       `json:"attachments,omitempty" structs:"attachments,omitempty"`
	Public       string           `json:"public,omitempty" structs:"public,omitempty"`
	Expands      []string         `json:"_expands,omitempty" structs:"_expands,omitempty"`
	Links        *SelfLink        `json:"_links,omitempty" structs:"_links,omitempty"`
}

type RenderedValueDTO struct {
	HTML string `json:"html,omitempty" structs:"html,omitempty"`
}

// ListRequestCommentsWithContext lists comments for a ServiceDesk request.
//
// https://developer.atlassian.com/cloud/jira/service-desk/rest/api-group-request/#api-rest-servicedeskapi-request-issueidorkey-comment-get
func (s *ServiceDeskService) ListRequestCommentsWithContext(ctx context.Context, idOrKey string, options *RequestCommentListOptions) (*PagedDTOT[CommentDTO], *Response, error) {
	apiEndpoint := fmt.Sprintf("rest/servicedeskapi/request/%s/comment", idOrKey)
	req, err := s.client.NewRequestWithContext(ctx, "GET", apiEndpoint, nil)
	if err != nil {
		return nil, nil, err
	}

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

	commentList := new(PagedDTOT[CommentDTO])
	if err := json.NewDecoder(resp.Body).Decode(commentList); err != nil {
		return nil, resp, fmt.Errorf("could not unmarshall the data into struct")
	}

	return commentList, resp, nil
}
