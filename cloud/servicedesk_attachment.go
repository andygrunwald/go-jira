package cloud

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"

	"github.com/andygrunwald/go-jira/v2/cloud/models/servicedesk"
	"golang.org/x/sync/errgroup"
)

const boundary = "rjMhuusQtNPAshccLgTNLBqzpuVSjpRm"

// AttachTemporaryFile attaches temporary file to a ServiceDesk.
//
// https://developer.atlassian.com/cloud/jira/service-desk/rest/api-group-servicedesk/#api-rest-servicedeskapi-servicedesk-servicedeskid-attachtemporaryfile-post
func (s *ServiceDeskService) AttachTemporaryFile(ctx context.Context, serviceDeskID int, files []servicedesk.TemporaryFile) (*servicedesk.AttachedTemporaryFile, *Response, error) {
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
		req, err := s.client.NewRawRequest(gCtx, http.MethodPost, apiEndpoint, pipeRead)
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

	commentList := new(servicedesk.AttachedTemporaryFile)
	if err := json.NewDecoder(resp.Body).Decode(commentList); err != nil {
		return nil, resp, err
	}

	return commentList, resp, nil
}
