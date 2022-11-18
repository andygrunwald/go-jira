package cloud

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/mcl-de/go-jira/v2/cloud/model/apps/assets"
)

// AssetObjectSchemaService handles object schema in Insight App for the Jira instance / API.
type AssetsObjectSchemaService service

// List resource to find object schemas
// Reference: https://developer.atlassian.com/cloud/assets/rest/api-group-objectschema/#api-objectschema-list-get
func (a *AssetsObjectSchemaService) List(ctx context.Context, workspaceID string) (*assets.GenericList[assets.ObjectSchema], error) {
	apiEndPoint := fmt.Sprintf(`%s/jsm/assets/workspace/%s/v1/objectschema/list`, assetURL, workspaceID)

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

	list := new(assets.GenericList[assets.ObjectSchema])
	err = json.NewDecoder(res.Body).Decode(&list)

	return list, err
}

// Create creates an object schema
// Reference: https://developer.atlassian.com/cloud/assets/rest/api-group-objectschema/#api-objectschema-create-post
func (a *AssetsObjectSchemaService) Create(ctx context.Context, workspaceID string, body assets.CreateOrUpdateObjectSchema) (*assets.ObjectSchema, error) {
	apiEndPoint := fmt.Sprintf(`%s/jsm/assets/workspace/%s/v1/objectschema/create`, assetURL, workspaceID)

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

	schema := new(assets.ObjectSchema)
	err = json.NewDecoder(res.Body).Decode(&schema)

	return schema, err
}

// Get returns a schema by id
// Reference: https://developer.atlassian.com/cloud/assets/rest/api-group-objectschema/#api-objectschema-id-get
func (a *AssetsObjectSchemaService) Get(ctx context.Context, workspaceID, id string) (*assets.ObjectSchema, error) {
	apiEndPoint := fmt.Sprintf(`%s/jsm/assets/workspace/%s/v1/objectschema/%s`, assetURL, workspaceID, id)

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

	schema := new(assets.ObjectSchema)
	err = json.NewDecoder(res.Body).Decode(&schema)

	return schema, err
}

// Update updates an object schema
// Reference: https://developer.atlassian.com/cloud/assets/rest/api-group-objectschema/#api-objectschema-id-put
func (a *AssetsObjectSchemaService) Update(ctx context.Context, workspaceID, id string, body assets.CreateOrUpdateObjectSchema) (*assets.ObjectSchema, error) {
	apiEndPoint := fmt.Sprintf(`%s/jsm/assets/workspace/%s/v1/objectschema/%s`, assetURL, workspaceID, id)

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

	schema := new(assets.ObjectSchema)
	err = json.NewDecoder(res.Body).Decode(&schema)

	return schema, err
}

// Delete deletes an object schema
// Reference: https://developer.atlassian.com/cloud/asset/rest/api-group-objectschema/#api-objectschema-id-delete
func (a *AssetsObjectSchemaService) Delete(ctx context.Context, workspaceID, id string) (*assets.ObjectSchema, error) {
	apiEndPoint := fmt.Sprintf(`%s/jsm/assets/workspace/%s/v1/objectschema/%s`, assetURL, workspaceID, id)

	req, err := a.client.NewRequest(ctx, http.MethodDelete, apiEndPoint, nil)
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

	schema := new(assets.ObjectSchema)
	err = json.NewDecoder(res.Body).Decode(&schema)

	return schema, err
}

// GetAttributes find all object type attributes for this object schema
// Reference: https://developer.atlassian.com/cloud/assets/rest/api-group-objectschema/#api-objectschema-id-attributes-get
func (a *AssetsObjectSchemaService) GetAttributes(ctx context.Context, workspaceID, id string, options *assets.ObjectSchemaAttributeOptions) ([]assets.ObjectTypeAttribute, error) {
	apiEndPoint := fmt.Sprintf(`%s/jsm/assets/workspace/%s/v1/objectschema/%s/attributes`, assetURL, workspaceID, id)
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

// GetObjectTypes find all object types for this object schema
// Reference: https://developer.atlassian.com/cloud/insight/rest/api-group-objectschema/#api-objectschema-id-attributes-get
func (a *AssetsObjectSchemaService) GetObjectTypes(ctx context.Context, workspaceID, id string, options *assets.ObjectSchemaObjectTypeOptions) ([]assets.ObjectType, error) {
	apiEndPoint := fmt.Sprintf(`%s/jsm/assets/workspace/%s/v1/objectschema/%s/objecttypes/flat`, assetURL, workspaceID, id)
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

	var objectTypes []assets.ObjectType
	err = json.NewDecoder(res.Body).Decode(&objectTypes)

	return objectTypes, err
}
