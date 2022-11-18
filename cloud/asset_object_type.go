package cloud

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/mcl-de/go-jira/v2/cloud/model/apps/assets"
)

// AssetsObjectTypeService handles object type in Insight App for the Jira instance / API.
type AssetsObjectTypeService service

// Get resource to return an object type
// Reference: https://developer.atlassian.com/cloud/assets/rest/api-group-objecttype/#api-objecttype-id-get
func (a *AssetsObjectTypeService) Get(ctx context.Context, workspaceID, id string) (*assets.ObjectType, error) {
	apiEndPoint := fmt.Sprintf(`%s/jsm/assets/workspace/%s/v1/objecttype/%s`, assetURL, workspaceID, id)

	req, err := a.client.NewRequest(ctx, http.MethodGet, apiEndPoint, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/json")

	res, err := a.client.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	err = CheckResponse(req, res)
	if err != nil {
		return nil, err
	}

	object := new(assets.ObjectType)
	err = json.NewDecoder(res.Body).Decode(&object)

	return object, err
}

// Update updates an object type and returns it
// Reference: https://developer.atlassian.com/cloud/asset/rest/api-group-objecttype/#api-objecttype-id-put
func (a *AssetsObjectTypeService) Update(ctx context.Context, workspaceID, id string, body assets.PutObjectType) (*assets.ObjectType, error) {
	apiEndPoint := fmt.Sprintf(`%s/jsm/assets/workspace/%s/v1/objecttype/%s`, assetURL, workspaceID, id)

	req, err := a.client.NewRequest(ctx, http.MethodPut, apiEndPoint, body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/json")

	res, err := a.client.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	err = CheckResponse(req, res)
	if err != nil {
		return nil, err
	}

	object := new(assets.ObjectType)
	err = json.NewDecoder(res.Body).Decode(&object)

	return object, err
}

// Delete delete an object type
// Reference: https://developer.atlassian.com/cloud/asset/rest/api-group-objecttype/#api-objecttype-id-delete
func (a *AssetsObjectTypeService) Delete(ctx context.Context, workspaceID, id string) error {
	apiEndPoint := fmt.Sprintf(`%s/jsm/asset/workspace/%s/v1/objecttype/%s`, assetURL, workspaceID, id)

	req, err := a.client.NewRequest(ctx, http.MethodDelete, apiEndPoint, nil)
	if err != nil {
		return err
	}

	res, err := a.client.client.Do(req)
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
// Reference: https://developer.atlassian.com/cloud/assets/rest/api-group-objecttype/#api-objecttype-id-attributes-get
func (a *AssetsObjectTypeService) GetAttributes(ctx context.Context, workspaceID, id string, options *assets.ObjectTypeAttributeOptions) ([]assets.ObjectTypeAttribute, error) {
	apiEndPoint := fmt.Sprintf(`%s/jsm/assets/workspace/%s/v1/objecttype/%s/attributes`, assetURL, workspaceID, id)
	url, err := addOptions(apiEndPoint, options)
	if err != nil {
		return nil, err
	}

	req, err := a.client.NewRequest(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/json")

	res, err := a.client.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	err = CheckResponse(req, res)
	if err != nil {
		return nil, err
	}

	var attributes []assets.ObjectTypeAttribute
	err = json.NewDecoder(res.Body).Decode(&attributes)

	return attributes, err
}

// UpdatePosition change position for the given object type
// Reference: https://developer.atlassian.com/cloud/assets/rest/api-group-objecttype/#api-objecttype-id-position-post
func (a *AssetsObjectTypeService) UpdatePosition(ctx context.Context, workspaceID, id string, body assets.PostObjectTypePosition) (*assets.ObjectType, error) {
	apiEndPoint := fmt.Sprintf(`%s/jsm/assets/workspace/%s/v1/objecttype/%s/position`, assetURL, workspaceID, id)

	req, err := a.client.NewRequest(ctx, http.MethodPost, apiEndPoint, body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/json")

	res, err := a.client.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	err = CheckResponse(req, res)
	if err != nil {
		return nil, err
	}

	object := new(assets.ObjectType)
	err = json.NewDecoder(res.Body).Decode(&object)

	return object, err
}

// Create creates an object type
// Reference: https://developer.atlassian.com/cloud/assets/rest/api-group-objecttype/#api-objecttype-create-post
func (a *AssetsObjectTypeService) Create(ctx context.Context, workspaceID string, body assets.CreateObjectType) (*assets.ObjectType, error) {
	apiEndPoint := fmt.Sprintf(`%s/jsm/assets/workspace/%s/v1/objecttype/create`, assetURL, workspaceID)

	req, err := a.client.NewRequest(ctx, http.MethodPost, apiEndPoint, body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/json")

	res, err := a.client.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	err = CheckResponse(req, res)
	if err != nil {
		return nil, err
	}

	object := new(assets.ObjectType)
	err = json.NewDecoder(res.Body).Decode(&object)

	return object, err
}
