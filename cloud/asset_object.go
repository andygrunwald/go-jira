package cloud

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/mcl-de/go-jira/v2/cloud/model/apps/assets"
)

// AssetsObjectService handles object in Assets App for the Jira instance / API.
type AssetsObjectService service

// Get returns a single object by id
// Reference: https://developer.atlassian.com/cloud/assets/rest/api-group-object/#api-object-id-get
func (a *AssetsObjectService) Get(ctx context.Context, workspaceID, id string) (*assets.Object, error) {
	apiEndPoint := fmt.Sprintf(`%s/jsm/assets/workspace/%s/v1/object/%s`, assetURL, workspaceID, id)

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

	object := new(assets.Object)
	err = json.NewDecoder(res.Body).Decode(&object)

	return object, err
}

// Update updates an object and returns it
// Reference: https://developer.atlassian.com/cloud/assets/rest/api-group-object/#api-object-id-put
func (a *AssetsObjectService) Update(ctx context.Context, workspaceID, id string, body assets.CreateOrUpdateObject) (*assets.Object, error) {
	apiEndPoint := fmt.Sprintf(`%s/jsm/assets/workspace/%s/v1/object/%s`, assetURL, workspaceID, id)

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

	object := new(assets.Object)
	err = json.NewDecoder(res.Body).Decode(&object)

	return object, err
}

// Delete delete an object
// Reference: https://developer.atlassian.com/cloud/assets/rest/api-group-object/#api-object-id-delete
func (a *AssetsObjectService) Delete(ctx context.Context, workspaceID, id string) error {
	apiEndPoint := fmt.Sprintf(`%s/jsm/assets/workspace/%s/v1/object/%s`, assetURL, workspaceID, id)

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

// GetAttributes list all attributes for the given object
// Reference: https://developer.atlassian.com/cloud/assets/rest/api-group-object/#api-object-id-attributes-get
func (a *AssetsObjectService) GetAttributes(ctx context.Context, workspaceID, id string) ([]assets.ObjectAttribute, error) {
	apiEndPoint := fmt.Sprintf(`%s/jsm/assets/workspace/%s/v1/object/%s/attributes`, assetURL, workspaceID, id)

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

	var attributes []assets.ObjectAttribute
	err = json.NewDecoder(res.Body).Decode(&attributes)

	return attributes, err
}

// GetHistory list history for the given object
// Reference: https://developer.atlassian.com/cloud/insight/rest/api-group-object/#api-object-id-history-get
func (a *AssetsObjectService) GetHistory(ctx context.Context, workspaceID, id string, options *assets.ObjectHistoryOptions) ([]assets.ObjectHistory, error) {
	apiEndPoint := fmt.Sprintf(`%s/jsm/assets/workspace/%s/v1/object/%s/history`, assetURL, workspaceID, id)
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

	var history []assets.ObjectHistory
	err = json.NewDecoder(res.Body).Decode(&history)

	return history, err
}

// GetReferenceInfo find all references for an object
// Reference: https://developer.atlassian.com/cloud/assets/rest/api-group-object/#api-object-id-referenceinfo-get
func (a *AssetsObjectService) GetReferenceInfo(ctx context.Context, workspaceID, id string) ([]assets.ObjectReferenceTypeInfo, error) {
	apiEndPoint := fmt.Sprintf(`%s/jsm/assets/workspace/%s/v1/object/%s/referenceinfo`, assetURL, workspaceID, id)

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

	var referenceInfos []assets.ObjectReferenceTypeInfo
	err = json.NewDecoder(res.Body).Decode(&referenceInfos)

	return referenceInfos, err
}

// Create creates a new object
// Reference: https://developer.atlassian.com/cloud/assets/rest/api-group-object/#api-object-create-post
func (a *AssetsObjectService) Create(ctx context.Context, workspaceID string, body assets.CreateOrUpdateObject) (*assets.Object, error) {
	apiEndPoint := fmt.Sprintf(`%s/jsm/assets/workspace/%s/v1/object/create`, assetURL, workspaceID)

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

	object := new(assets.Object)
	err = json.NewDecoder(res.Body).Decode(&object)

	return object, err
}

// NavlistIQL returns a list of objects based on an IQL. Note that the preferred endpoint is IQL from the InsightIQLService
// Reference: https://developer.atlassian.com/cloud/assets/rest/api-group-object/#api-object-navlist-iql-post
func (a *AssetsObjectService) NavlistIQL(ctx context.Context, workspaceID string, body assets.ObjectNavlistIQL) (*assets.ObjectListResult, error) {
	apiEndPoint := fmt.Sprintf(`%s/jsm/assets/workspace/%s/v1/object/navlist/iql`, assetURL, workspaceID)

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

	result := new(assets.ObjectListResult)
	err = json.NewDecoder(res.Body).Decode(&result)

	return result, err
}
