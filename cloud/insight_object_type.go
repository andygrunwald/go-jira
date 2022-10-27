package cloud

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/mcl-de/go-jira/v2/cloud/model/apps/insight"
)

// InsightObjectTypeService handles object type in Insight App for the Jira instance / API.
type InsightObjectTypeService service

// Get resource to return an object type
// Reference: https://developer.atlassian.com/cloud/insight/rest/api-group-objecttype/#api-objecttype-id-get
func (i *InsightObjectTypeService) Get(ctx context.Context, workspaceID, id string) (*insight.ObjectType, error) {
	apiEndPoint := fmt.Sprintf(`%s/jsm/insight/workspace/%s/v1/objecttype/%s`, insightURL, workspaceID, id)

	req, err := i.client.NewRequest(ctx, http.MethodGet, apiEndPoint, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/json")

	res, err := i.client.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	err = CheckResponse(req, res)
	if err != nil {
		return nil, err
	}

	object := new(insight.ObjectType)
	err = json.NewDecoder(res.Body).Decode(&object)

	return object, err
}

// Update updates an object type and returns it
// Reference: https://developer.atlassian.com/cloud/insight/rest/api-group-objecttype/#api-objecttype-id-put
func (i *InsightObjectTypeService) Update(ctx context.Context, workspaceID, id string, body insight.PutObjectType) (*insight.ObjectType, error) {
	apiEndPoint := fmt.Sprintf(`%s/jsm/insight/workspace/%s/v1/objecttype/%s`, insightURL, workspaceID, id)

	req, err := i.client.NewRequest(ctx, http.MethodPut, apiEndPoint, body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/json")

	res, err := i.client.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	err = CheckResponse(req, res)
	if err != nil {
		return nil, err
	}

	object := new(insight.ObjectType)
	err = json.NewDecoder(res.Body).Decode(&object)

	return object, err
}

// Delete delete an object type
// Reference: https://developer.atlassian.com/cloud/insight/rest/api-group-objecttype/#api-objecttype-id-delete
func (i *InsightObjectTypeService) Delete(ctx context.Context, workspaceID, id string) error {
	apiEndPoint := fmt.Sprintf(`%s/jsm/insight/workspace/%s/v1/objecttype/%s`, insightURL, workspaceID, id)

	req, err := i.client.NewRequest(ctx, http.MethodDelete, apiEndPoint, nil)
	if err != nil {
		return err
	}

	res, err := i.client.client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	err = CheckResponse(req, res)
	if err != nil {
		return err
	}

	return nil
}

// GetAttributes list all attributes for the given object type
// Reference: https://developer.atlassian.com/cloud/insight/rest/api-group-objecttype/#api-objecttype-id-attributes-get
func (i *InsightObjectTypeService) GetAttributes(ctx context.Context, workspaceID, id string, options *insight.ObjectTypeAttributeOptions) ([]insight.ObjectTypeAttribute, error) {
	apiEndPoint := fmt.Sprintf(`%s/jsm/insight/workspace/%s/v1/objecttype/%s/attributes`, insightURL, workspaceID, id)
	url, err := addOptions(apiEndPoint, options)
	if err != nil {
		return nil, err
	}

	req, err := i.client.NewRequest(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/json")

	res, err := i.client.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	err = CheckResponse(req, res)
	if err != nil {
		return nil, err
	}

	var attributes []insight.ObjectTypeAttribute
	err = json.NewDecoder(res.Body).Decode(&attributes)

	return attributes, err
}

// UpdatePosition change position for the given object type
// Reference: https://developer.atlassian.com/cloud/insight/rest/api-group-objecttype/#api-objecttype-id-position-post
func (i *InsightObjectTypeService) UpdatePosition(ctx context.Context, workspaceID, id string, body insight.PostObjectTypePosition) (*insight.ObjectType, error) {
	apiEndPoint := fmt.Sprintf(`%s/jsm/insight/workspace/%s/v1/objecttype/%s/position`, insightURL, workspaceID, id)

	req, err := i.client.NewRequest(ctx, http.MethodPost, apiEndPoint, body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/json")

	res, err := i.client.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	err = CheckResponse(req, res)
	if err != nil {
		return nil, err
	}

	object := new(insight.ObjectType)
	err = json.NewDecoder(res.Body).Decode(&object)

	return object, err
}

// Create creates an object type
// Reference: https://developer.atlassian.com/cloud/insight/rest/api-group-objecttype/#api-objecttype-create-post
func (i *InsightObjectTypeService) Create(ctx context.Context, workspaceID string, body insight.CreateObjectType) (*insight.ObjectType, error) {
	apiEndPoint := fmt.Sprintf(`%s/jsm/insight/workspace/%s/v1/objecttype/create`, insightURL, workspaceID)

	req, err := i.client.NewRequest(ctx, http.MethodPost, apiEndPoint, body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/json")

	res, err := i.client.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	err = CheckResponse(req, res)
	if err != nil {
		return nil, err
	}

	object := new(insight.ObjectType)
	err = json.NewDecoder(res.Body).Decode(&object)

	return object, err
}