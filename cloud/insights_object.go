package cloud

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/andygrunwald/go-jira/v2/cloud/models/apps/insights"
)

// GetObject returns a single object by id
// Reference: https://developer.atlassian.com/cloud/insight/rest/api-group-object/#api-object-id-get
func (i *InsightsService) GetObject(ctx context.Context, workspaceID, id string) (*insights.Object, error) {
	apiEndPoint := fmt.Sprintf(`%s/jsm/insight/workspace/%s/v1/object/%s`, insightsURL, workspaceID, id)

	req, err := i.client.NewRequest(ctx, http.MethodGet, apiEndPoint, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/json")

	res, err := i.client.client.Do(req)
	if err != nil {
		return nil, err
	}

	switch res.StatusCode {
	case http.StatusUnauthorized:
		return nil, fmt.Errorf("%s: %w", req.URL.String(), ErrUnauthorized)
	case http.StatusNotFound:
		return nil, fmt.Errorf("%s: %w", req.URL.String(), ErrNotFound)
	case http.StatusInternalServerError:
		return nil, fmt.Errorf("%s: %w", req.URL.String(), ErrUnknown)
	}

	object := new(insights.Object)
	err = json.NewDecoder(res.Body).Decode(&object)

	return object, err
}

// UpdateObject updates an object and returns it
// Reference: https://developer.atlassian.com/cloud/insight/rest/api-group-object/#api-object-id-put
func (i *InsightsService) UpdateObject(ctx context.Context, workspaceID, id string, body insights.PutObject) (*insights.Object, error) {
	apiEndPoint := fmt.Sprintf(`%s/jsm/insight/workspace/%s/v1/object/%s`, insightsURL, workspaceID, id)

	req, err := i.client.NewRequest(ctx, http.MethodGet, apiEndPoint, body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/json")

	res, err := i.client.client.Do(req)
	if err != nil {
		return nil, err
	}

	switch res.StatusCode {
	case http.StatusBadRequest:
		responseBody, err := io.ReadAll(res.Body)
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("%s: %w\n%s", req.URL.String(), ErrValidation, responseBody)
	case http.StatusUnauthorized:
		return nil, fmt.Errorf("%s: %w", req.URL.String(), ErrUnauthorized)
	case http.StatusNotFound:
		return nil, fmt.Errorf("%s: %w", req.URL.String(), ErrNotFound)
	case http.StatusInternalServerError:
		return nil, fmt.Errorf("%s: %w", req.URL.String(), ErrUnknown)
	}

	object := new(insights.Object)
	err = json.NewDecoder(res.Body).Decode(&object)

	return object, err
}

// DeleteObject delete an object
// Reference: https://developer.atlassian.com/cloud/insight/rest/api-group-object/#api-object-id-delete
func (i *InsightsService) DeleteObject(ctx context.Context, workspaceID, id string) error {
	apiEndPoint := fmt.Sprintf(`%s/jsm/insight/workspace/%s/v1/object/%s`, insightsURL, workspaceID, id)

	req, err := i.client.NewRequest(ctx, http.MethodDelete, apiEndPoint, nil)
	if err != nil {
		return err
	}

	res, err := i.client.client.Do(req)
	if err != nil {
		return err
	}

	switch res.StatusCode {
	case http.StatusBadRequest:
		responseBody, err := io.ReadAll(res.Body)
		if err != nil {
			return err
		}

		return fmt.Errorf("%s: %w\n%s", req.URL.String(), ErrValidation, responseBody)
	case http.StatusUnauthorized:
		return fmt.Errorf("%s: %w", req.URL.String(), ErrUnauthorized)
	case http.StatusNotFound:
		return fmt.Errorf("%s: %w", req.URL.String(), ErrNotFound)
	case http.StatusInternalServerError:
		return fmt.Errorf("%s: %w", req.URL.String(), ErrUnknown)
	}

	return err
}

// GetObjectAttributes list all attributes for the given object
// Reference: https://developer.atlassian.com/cloud/insight/rest/api-group-object/#api-object-id-attributes-get
func (i *InsightsService) GetObjectAttributes(ctx context.Context, workspaceID, id string) ([]insights.ObjectAttribute, error) {
	apiEndPoint := fmt.Sprintf(`%s/jsm/insight/workspace/%s/v1/object/%s/attributes`, insightsURL, workspaceID, id)

	req, err := i.client.NewRequest(ctx, http.MethodGet, apiEndPoint, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/json")

	res, err := i.client.client.Do(req)
	if err != nil {
		return nil, err
	}

	switch res.StatusCode {
	case http.StatusUnauthorized:
		return nil, fmt.Errorf("%s: %w", req.URL.String(), ErrUnauthorized)
	case http.StatusNotFound:
		return nil, fmt.Errorf("%s: %w", req.URL.String(), ErrNotFound)
	case http.StatusInternalServerError:
		return nil, fmt.Errorf("%s: %w", req.URL.String(), ErrUnknown)
	}

	var attributes []insights.ObjectAttribute
	err = json.NewDecoder(res.Body).Decode(&attributes)

	return attributes, err
}

// GetObjectReferenceInfo find all references for an object
// Reference: https://developer.atlassian.com/cloud/insight/rest/api-group-object/#api-object-id-referenceinfo-get
func (i *InsightsService) GetObjectReferenceInfo(ctx context.Context, workspaceID, id string) ([]insights.ObjectReferenceTypeInfo, error) {
	apiEndPoint := fmt.Sprintf(`%s/jsm/insight/workspace/%s/v1/object/%s/referenceinfo`, insightsURL, workspaceID, id)

	req, err := i.client.NewRequest(ctx, http.MethodGet, apiEndPoint, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/json")

	res, err := i.client.client.Do(req)
	if err != nil {
		return nil, err
	}

	switch res.StatusCode {
	case http.StatusUnauthorized:
		return nil, fmt.Errorf("%s: %w", req.URL.String(), ErrUnauthorized)
	case http.StatusNotFound:
		return nil, fmt.Errorf("%s: %w", req.URL.String(), ErrNotFound)
	case http.StatusInternalServerError:
		return nil, fmt.Errorf("%s: %w", req.URL.String(), ErrUnknown)
	}

	var referenceInfos []insights.ObjectReferenceTypeInfo
	err = json.NewDecoder(res.Body).Decode(&referenceInfos)

	return referenceInfos, err
}

// CreateObject create a new object
// Reference: https://developer.atlassian.com/cloud/insight/rest/api-group-object/#api-object-create-post
func (i *InsightsService) CreateObject(ctx context.Context, workspaceID string, body insights.PutObject) (*insights.Object, error) {
	apiEndPoint := fmt.Sprintf(`%s/jsm/insight/workspace/%s/v1/object/create`, insightsURL, workspaceID)

	req, err := i.client.NewRequest(ctx, http.MethodPost, apiEndPoint, body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/json")

	res, err := i.client.client.Do(req)
	if err != nil {
		return nil, err
	}

	switch res.StatusCode {
	case http.StatusBadRequest:
		responseBody, err := io.ReadAll(res.Body)
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("%s: %w\n%s", req.URL.String(), ErrValidation, responseBody)
	case http.StatusUnauthorized:
		return nil, fmt.Errorf("%s: %w", req.URL.String(), ErrUnauthorized)
	case http.StatusInternalServerError:
		return nil, fmt.Errorf("%s: %w", req.URL.String(), ErrUnknown)
	}

	object := new(insights.Object)
	err = json.NewDecoder(res.Body).Decode(&object)

	return object, err
}
