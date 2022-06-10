//go:build go1.18
// +build go1.18

package jira

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/go-querystring/query"
)

// RequestCommentListOptions is the query options for listing comments for a ServiceDesk request.
type RequestCommentListOptions struct {
	Public   *bool    `url:"public,omitempty" query:"public"`
	Internal *bool    `url:"internal,omitempty" query:"internal"`
	Expand   []string `url:"expand,omitempty" query:"expand"`
	Start    int      `url:"start,omitempty" query:"start"`
	Limit    int      `url:"limit,omitempty" query:"limit"`
}

// CommentDTO represents a comment to an request in ServiceDesk.
type CommentDTO struct {
	ID           string
	Body         string            `json:"body,omitempty" structs:"body,omitempty"`
	RenderedBody *RenderedValueDTO `json:"renderedBody,omitempty" structs:"renderedBody,omitempty"`
	Author       UserDTO           `json:"author,omitempty" structs:"author,omitempty"`
	Created      DateDTO           `json:"created,omitempty" structs:"created,omitempty"`
	Attachments  *Attachment       `json:"attachments,omitempty" structs:"attachments,omitempty"`
	Public       bool              `json:"public,omitempty" structs:"public,omitempty"`
	Expands      []string          `json:"_expands,omitempty" structs:"_expands,omitempty"`
	Links        *SelfLink         `json:"_links,omitempty" structs:"_links,omitempty"`
}

type UserDTO struct {
	AccountID    string       `json:"accountId,omitempty" structs:"accountId,omitempty"`
	Name         string       `json:"name,omitempty" structs:"name,omitempty"`
	Key          string       `json:"key,omitempty" structs:"key,omitempty"`
	EmailAddress string       `json:"emailAddress,omitempty" structs:"emailAddress,omitempty"`
	DisplayName  string       `json:"displayName,omitempty" structs:"displayName,omitempty"`
	Active       bool         `json:"active" structs:"active"`
	TimeZone     string       `json:"timeZone,omitempty" structs:"timeZone,omitempty"`
	Links        *UserLinkDTO `json:"_links,omitempty" structs:"_links,omitempty"`
}

type UserLinkDTO struct {
	Self       string            `json:"self,omitempty" structs:"self,omitempty"`
	JIRARest   string            `json:"jiraRest,omitempty" structs:"jiraRest,omitempty"`
	AvatarUrls map[string]string `json:"avatarUrls,omitempty" structs:"avatarUrls,omitempty"`
}

type RenderedValueDTO struct {
	HTML string `json:"html,omitempty" structs:"html,omitempty"`
}

// ListRequestCommentsWithContext lists comments for a ServiceDesk request.
//
// https://developer.atlassian.com/cloud/jira/service-desk/rest/api-group-request/#api-rest-servicedeskapi-request-issueidorkey-comment-get
func (s *ServiceDeskService) ListRequestCommentsWithContext(ctx context.Context, idOrKey string, options *RequestCommentListOptions) (*PagedDTOT[CommentDTO], *Response, error) {
	apiEndpoint := fmt.Sprintf("rest/servicedeskapi/request/%s/comment", idOrKey)
	req, err := s.client.NewRequestWithContext(ctx, http.MethodGet, apiEndpoint, nil)
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
		return nil, resp, err
	}

	return commentList, resp, nil
}

type CreateRequestComment struct {
	Body   string `json:"body" structs:"body"`
	Public bool   `json:"public" structs:"public"`
}

// CreateRequestCommentsWithContext create a comment for a ServiceDesk request.
//
// https://developer.atlassian.com/cloud/jira/service-desk/rest/api-group-request/#api-rest-servicedeskapi-request-issueidorkey-comment-post
func (s *ServiceDeskService) CreateRequestCommentsWithContext(ctx context.Context, idOrKey string, request CreateRequestComment) (*CommentDTO, *Response, error) {
	apiEndpoint := fmt.Sprintf("rest/servicedeskapi/request/%s/comment", idOrKey)
	req, err := s.client.NewRequestWithContext(ctx, http.MethodPost, apiEndpoint, request)
	if err != nil {
		return nil, nil, err
	}

	req.Header.Set("Accept", "application/json")

	resp, err := s.client.Do(req, nil)
	if err != nil {
		return nil, resp, NewJiraError(resp, err)
	}
	defer resp.Body.Close()

	comment := new(CommentDTO)
	if err = json.NewDecoder(resp.Body).Decode(comment); err != nil {
		return nil, resp, err
	}

	return comment, resp, nil
}

type AttachmentDTO struct {
	Filename string            `json:"filename,omitempty" structs:"filename,omitempty"`
	Author   UserDTO           `json:"author,omitempty" structs:"author,omitempty"`
	Created  DateDTO           `json:"created,omitempty" structs:"created,omitempty"`
	Size     int64             `json:"size,omitempty" structs:"size,omitempty"`
	MIMEType string            `json:"mimeType,omitempty" structs:"mimeType,omitempty"`
	Links    AttachmentLinkDTO `json:"_links,omitempty" structs:"_links,omitempty"`
}

type AttachmentLinkDTO struct {
	Self      string `json:"self,omitempty" structs:"self,omitempty"`
	JIRARest  string `json:"jiraRest,omitempty" structs:"jiraRest,omitempty"`
	Content   string `json:"content,omitempty" structs:"content,omitempty"`
	Thumbnail string `json:"thumbnail,omitempty" structs:"thumbnail,omitempty"`
}

type AttachmentCreateResultDTO struct {
	Comment     CommentDTO               `json:"comment,omitempty" structs:"comment,omitempty"`
	Attachments PagedDTOT[AttachmentDTO] `json:"attachments,omitempty" structs:"attachments,omitempty"`
}

type CreateRequestAttachment struct {
	TemporaryAttachmentIDs []string              `json:"temporaryAttachmentIds,omitempty" structs:"temporaryAttachmentIds,omitempty"`
	AdditionalComment      *AdditionalCommentDTO `json:"additionalComment,omitempty" structs:"additionalComment,omitempty"`
	Public                 bool                  `json:"public" structs:"public"`
}

type AdditionalCommentDTO struct {
	Body string `json:"body,omitempty" structs:"body,omitempty"`
}

func (s *ServiceDeskService) CreateAttachment(ctx context.Context, idOrKey string, request CreateRequestAttachment) (*AttachmentCreateResultDTO, *Response, error) {
	apiEndpoint := fmt.Sprintf("rest/servicedeskapi/request/%s/attachment", idOrKey)
	req, err := s.client.NewRequestWithContext(ctx, http.MethodPost, apiEndpoint, request)
	if err != nil {
		return nil, nil, err
	}

	req.Header.Set("Accept", "application/json")

	resp, err := s.client.Do(req, nil)
	if err != nil {
		return nil, resp, NewJiraError(resp, err)
	}
	defer resp.Body.Close()

	attachment := new(AttachmentCreateResultDTO)
	if err = json.NewDecoder(resp.Body).Decode(attachment); err != nil {
		return nil, resp, err
	}

	return attachment, resp, nil
}
