package cloud

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/andygrunwald/go-jira/v2/cloud/models"
	"github.com/andygrunwald/go-jira/v2/cloud/models/servicedesk"
	"github.com/google/go-querystring/query"
)

// RequestService handles ServiceDesk customer requests for the Jira instance / API.
type RequestService service

// Create creates a new request.
//
// https://developer.atlassian.com/cloud/jira/service-desk/rest/api-group-request/#api-rest-servicedeskapi-request-post
func (r *RequestService) Create(ctx context.Context, request *servicedesk.CreateRequest) (*servicedesk.Request, *Response, error) {
	apiEndpoint := "rest/servicedeskapi/request"

	req, err := r.client.NewRequest(ctx, http.MethodPost, apiEndpoint, request)
	if err != nil {
		return nil, nil, err
	}

	responseRequest := new(servicedesk.Request)
	resp, err := r.client.Do(req, responseRequest)
	if err != nil {
		return nil, resp, NewJiraError(resp, err)
	}

	return responseRequest, resp, nil
}

// CreateComment creates a comment on a request.
//
// https://developer.atlassian.com/cloud/jira/service-desk/rest/api-group-request/#api-rest-servicedeskapi-request-issueidorkey-comment-post
func (r *RequestService) CreateComment(ctx context.Context, issueIDOrKey string, comment *servicedesk.RequestComment) (*servicedesk.RequestComment, *Response, error) {
	apiEndpoint := fmt.Sprintf("rest/servicedeskapi/request/%v/comment", issueIDOrKey)

	req, err := r.client.NewRequest(ctx, http.MethodPost, apiEndpoint, comment)
	if err != nil {
		return nil, nil, err
	}

	responseComment := new(servicedesk.RequestComment)
	resp, err := r.client.Do(req, responseComment)
	if err != nil {
		return nil, resp, NewJiraError(resp, err)
	}

	return responseComment, resp, nil
}

func (r *RequestService) CreateAttachment(ctx context.Context, idOrKey string, request servicedesk.CreateRequestAttachment) (*servicedesk.AttachmentCreateResultDTO, *Response, error) {
	apiEndpoint := fmt.Sprintf("rest/servicedeskapi/request/%s/attachment", idOrKey)
	req, err := r.client.NewRequest(ctx, http.MethodPost, apiEndpoint, request)
	if err != nil {
		return nil, nil, err
	}

	req.Header.Set("Accept", "application/json")

	resp, err := r.client.Do(req, nil)
	if err != nil {
		return nil, resp, NewJiraError(resp, err)
	}
	defer resp.Body.Close()

	attachment := new(servicedesk.AttachmentCreateResultDTO)
	if err = json.NewDecoder(resp.Body).Decode(attachment); err != nil {
		return nil, resp, err
	}

	return attachment, resp, nil
}

// GetRequestParticipants returns all request participants
//
// https://developer.atlassian.com/cloud/jira/service-desk/rest/api-group-request/#api-rest-servicedeskapi-request-issueidorkey-participant-get
func (r *RequestService) GetRequestParticipants(ctx context.Context, idOrKey string) (*models.PagedDTOT[servicedesk.UserDTO], *Response, error) {
	apiEndpoint := fmt.Sprintf("rest/servicedeskapi/request/%s/participant", idOrKey)

	req, err := r.client.NewRequest(ctx, http.MethodGet, apiEndpoint, nil)
	if err != nil {
		return nil, nil, err
	}

	responseRequest := new(models.PagedDTOT[servicedesk.UserDTO])
	resp, err := r.client.Do(req, responseRequest)
	if err != nil {
		return nil, resp, err
	}

	return responseRequest, resp, nil
}

type ChangeRequestParticipants struct {
	AccountIDs []string `json:"accountIds"`
}

// AddRequestParticipants adds a request participants to a request
//
// https://developer.atlassian.com/cloud/jira/service-desk/rest/api-group-request/#api-rest-servicedeskapi-request-issueidorkey-participant-post
func (r *RequestService) AddRequestParticipants(ctx context.Context, idOrKey string, participants ChangeRequestParticipants) (*models.PagedDTOT[servicedesk.UserDTO], *Response, error) {
	apiEndpoint := fmt.Sprintf("rest/servicedeskapi/request/%s/participant", idOrKey)

	req, err := r.client.NewRequest(ctx, http.MethodPost, apiEndpoint, participants)
	if err != nil {
		return nil, nil, err
	}

	responseRequest := new(models.PagedDTOT[servicedesk.UserDTO])
	resp, err := r.client.Do(req, responseRequest)
	if err != nil {
		return nil, resp, err
	}

	return responseRequest, resp, nil
}

// RemoveRequestParticipants removes a request participants from a request
//
// https://developer.atlassian.com/cloud/jira/service-desk/rest/api-group-request/#api-rest-servicedeskapi-request-issueidorkey-participant-delete
func (r *RequestService) RemoveRequestParticipants(ctx context.Context, idOrKey string, participants ChangeRequestParticipants) (*models.PagedDTOT[servicedesk.UserDTO], *Response, error) {
	apiEndpoint := fmt.Sprintf("rest/servicedeskapi/request/%s/participant", idOrKey)

	req, err := r.client.NewRequest(ctx, http.MethodDelete, apiEndpoint, participants)
	if err != nil {
		return nil, nil, err
	}

	responseRequest := new(models.PagedDTOT[servicedesk.UserDTO])
	resp, err := r.client.Do(req, responseRequest)
	if err != nil {
		return nil, resp, err
	}

	return responseRequest, resp, nil
}

// CreateRequestComments create a comment for a ServiceDesk request.
//
// https://developer.atlassian.com/cloud/jira/service-desk/rest/api-group-request/#api-rest-servicedeskapi-request-issueidorkey-comment-post
func (s *ServiceDeskService) CreateRequestComments(ctx context.Context, idOrKey string, request servicedesk.CreateRequestComment) (*servicedesk.CommentDTO, *Response, error) {
	apiEndpoint := fmt.Sprintf("rest/servicedeskapi/request/%s/comment", idOrKey)

	req, err := s.client.NewRequest(ctx, http.MethodPost, apiEndpoint, request)
	if err != nil {
		return nil, nil, err
	}

	req.Header.Set("Accept", "application/json")

	resp, err := s.client.Do(req, nil)
	if err != nil {
		return nil, resp, NewJiraError(resp, err)
	}
	defer resp.Body.Close()

	comment := new(servicedesk.CommentDTO)
	if err = json.NewDecoder(resp.Body).Decode(comment); err != nil {
		return nil, resp, err
	}

	return comment, resp, nil
}

// ListRequestComments lists comments for a ServiceDesk request.
//
// https://developer.atlassian.com/cloud/jira/service-desk/rest/api-group-request/#api-rest-servicedeskapi-request-issueidorkey-comment-get
func (s *ServiceDeskService) ListRequestComments(ctx context.Context, idOrKey string, options *servicedesk.RequestCommentListOptions) (*models.PagedDTOT[servicedesk.CommentDTO], *Response, error) {
	apiEndpoint := fmt.Sprintf("rest/servicedeskapi/request/%s/comment", idOrKey)

	req, err := s.client.NewRequest(ctx, http.MethodGet, apiEndpoint, nil)
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

	commentList := new(models.PagedDTOT[servicedesk.CommentDTO])
	if err := json.NewDecoder(resp.Body).Decode(commentList); err != nil {
		return nil, resp, err
	}

	return commentList, resp, nil
}
