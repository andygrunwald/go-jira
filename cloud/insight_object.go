package cloud

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/mcl-de/go-jira/v2/cloud/model/apps/insight"
)

// InsightObjectService handles object in Insight App for the Jira instance / API.
type InsightObjectService service

// Get returns a single object by id
// Reference: https://developer.atlassian.com/cloud/insight/rest/api-group-object/#api-object-id-get
func (i *InsightObjectService) Get(ctx context.Context, workspaceID, id string) (*insight.Object, error) {
	apiEndPoint := fmt.Sprintf(`%s/jsm/insight/workspace/%s/v1/object/%s`, insightURL, workspaceID, id)

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

	object := new(insight.Object)
	err = json.NewDecoder(res.Body).Decode(&object)

	return object, err
}

// Update updates an object and returns it
// Reference: https://developer.atlassian.com/cloud/insight/rest/api-group-object/#api-object-id-put
func (i *InsightObjectService) Update(ctx context.Context, workspaceID, id string, body insight.CreateOrUpdateObject) (*insight.Object, error) {
	apiEndPoint := fmt.Sprintf(`%s/jsm/insight/workspace/%s/v1/object/%s`, insightURL, workspaceID, id)

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

	object := new(insight.Object)
	err = json.NewDecoder(res.Body).Decode(&object)

	return object, err
}

// Delete delete an object
// Reference: https://developer.atlassian.com/cloud/insight/rest/api-group-object/#api-object-id-delete
func (i *InsightObjectService) Delete(ctx context.Context, workspaceID, id string) error {
	apiEndPoint := fmt.Sprintf(`%s/jsm/insight/workspace/%s/v1/object/%s`, insightURL, workspaceID, id)

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

// GetAttributes list all attributes for the given object
// Reference: https://developer.atlassian.com/cloud/insight/rest/api-group-object/#api-object-id-attributes-get
func (i *InsightObjectService) GetAttributes(ctx context.Context, workspaceID, id string) ([]insight.ObjectAttribute, error) {
	apiEndPoint := fmt.Sprintf(`%s/jsm/insight/workspace/%s/v1/object/%s/attributes`, insightURL, workspaceID, id)

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

	var attributes []insight.ObjectAttribute
	err = json.NewDecoder(res.Body).Decode(&attributes)

	return attributes, err
}

// GetHistory list history for the given object
// Reference: https://developer.atlassian.com/cloud/insight/rest/api-group-object/#api-object-id-history-get
func (i *InsightObjectService) GetHistory(ctx context.Context, workspaceID, id string, options *insight.ObjectHistoryOptions) ([]insight.ObjectHistory, error) {
	apiEndPoint := fmt.Sprintf(`%s/jsm/insight/workspace/%s/v1/object/%s/history`, insightURL, workspaceID, id)
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

	var history []insight.ObjectHistory
	err = json.NewDecoder(res.Body).Decode(&history)

	return history, err
}

// GetReferenceInfo find all references for an object
// Reference: https://developer.atlassian.com/cloud/insight/rest/api-group-object/#api-object-id-referenceinfo-get
func (i *InsightObjectService) GetReferenceInfo(ctx context.Context, workspaceID, id string) ([]insight.ObjectReferenceTypeInfo, error) {
	apiEndPoint := fmt.Sprintf(`%s/jsm/insight/workspace/%s/v1/object/%s/referenceinfo`, insightURL, workspaceID, id)

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

	var referenceInfos []insight.ObjectReferenceTypeInfo
	err = json.NewDecoder(res.Body).Decode(&referenceInfos)

	return referenceInfos, err
}

// Create creates a new object
// Reference: https://developer.atlassian.com/cloud/insight/rest/api-group-object/#api-object-create-post
func (i *InsightObjectService) Create(ctx context.Context, workspaceID string, body insight.CreateOrUpdateObject) (*insight.Object, error) {
	apiEndPoint := fmt.Sprintf(`%s/jsm/insight/workspace/%s/v1/object/create`, insightURL, workspaceID)

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

	object := new(insight.Object)
	err = json.NewDecoder(res.Body).Decode(&object)

	return object, err
}

// NavlistIQL returns a list of objects based on an IQL. Note that the preferred endpoint is IQL from the InsightIQLService
// Reference: https://developer.atlassian.com/cloud/insight/rest/api-group-object/#api-object-navlist-iql-post
func (i *InsightObjectService) NavlistIQL(ctx context.Context, workspaceID string, body insight.ObjectNavlistIQL) (*insight.ObjectListResult, error) {
	apiEndPoint := fmt.Sprintf(`%s/jsm/insight/workspace/%s/v1/object/navlist/iql`, insightURL, workspaceID)

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

	result := new(insight.ObjectListResult)
	err = json.NewDecoder(res.Body).Decode(&result)

	return result, err
}