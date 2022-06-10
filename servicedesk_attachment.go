package jira

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"

	"golang.org/x/sync/errgroup"
)

type DateDTO struct {
	ISO8601     string `json:"iso8601,omitempty" structs:"iso8601,omitempty"`
	JIRA        string `json:"jira,omitempty" structs:"jira,omitempty"`
	Friendly    string `json:"friendly,omitempty" structs:"friendly,omitempty"`
	EpochMillis int64  `json:"epochMillis,omitempty" structs:"epochMillis,omitempty"`
}

const boundary = "rjMhuusQtNPAshccLgTNLBqzpuVSjpRm"

type TemporaryFile struct {
	Name string
	File io.Reader
}

type AttachedTemporaryFile struct {
	TemporaryAttachments []TemporaryAttachment `json:"temporaryAttachments,omitempty" structs:"temporaryAttachments,omitempty"`
}

type TemporaryAttachment struct {
	TemporaryAttachmentID string `json:"temporaryAttachmentId,omitempty" structs:"temporaryAttachmentId,omitempty"`
	FileName              string `json:"fileName,omitempty" structs:"fileName,omitempty"`
}

// AttachTemporaryFile attaches temporary file to a ServiceDesk.
//
// https://developer.atlassian.com/cloud/jira/service-desk/rest/api-group-servicedesk/#api-rest-servicedeskapi-servicedesk-servicedeskid-attachtemporaryfile-post
func (s *ServiceDeskService) AttachTemporaryFile(ctx context.Context, serviceDeskID int, files []TemporaryFile) (*AttachedTemporaryFile, *Response, error) {
	pipeRead, pipeWrite := io.Pipe()

	// Set up multipart body for reading
	multipartWriter := multipart.NewWriter(pipeWrite)
	err := multipartWriter.SetBoundary(boundary)
	if err != nil {
		return nil, nil, err
	}

	var resp *Response

	wg, gCtx := errgroup.WithContext(ctx)
	wg.Go(func() error {
		var err error
		defer func() {
			if err != nil {
				pipeWrite.CloseWithError(err)
			} else {
				pipeWrite.Close()
			}
		}()

		var formFile io.Writer
		for _, v := range files {
			formFile, err = multipartWriter.CreateFormFile("file", v.Name)
			if err != nil {
				return fmt.Errorf(`file "%s": %w`, v.Name, err)
			}

			_, err = io.Copy(formFile, v.File)
			if err != nil {
				return fmt.Errorf(`file "%s" copy: %w`, v.Name, err)
			}
		}

		return multipartWriter.Close()
	})
	wg.Go(func() error {
		apiEndpoint := fmt.Sprintf("rest/servicedeskapi/servicedesk/%d/attachTemporaryFile", serviceDeskID)
		req, err := s.client.NewRawRequestWithContext(gCtx, http.MethodPost, apiEndpoint, pipeRead)
		if err != nil {
			return err
		}

		req.Header.Set("Content-Type", multipartWriter.FormDataContentType())
		req.Header.Set("Accept", "application/json")
		req.Header.Set("X-Atlassian-Token", "no-check")

		resp, err = s.client.Do(req, nil)
		if err != nil {
			return NewJiraError(resp, err)
		}

		return nil
	})
	err = wg.Wait()
	if err != nil {
		return nil, resp, err
	}
	defer resp.Body.Close()

	commentList := new(AttachedTemporaryFile)
	if err := json.NewDecoder(resp.Body).Decode(commentList); err != nil {
		return nil, resp, err
	}

	return commentList, resp, nil
}
