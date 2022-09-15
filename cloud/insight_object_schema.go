package cloud

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/andygrunwald/go-jira/v2/cloud/model/apps/insight"
)

// InsightObjectSchemaService handles object schema in Insight App for the Jira instance / API.
type InsightObjectSchemaService service

// List resource to find object schemas
// Reference: https://developer.atlassian.com/cloud/insight/rest/api-group-objectschema/#api-objectschema-list-get
func (i *InsightObjectSchemaService) List(ctx context.Context, workspaceID string) (*insight.GenericList[insight.ObjectSchema], error) {
	apiEndPoint := fmt.Sprintf(`%s/jsm/insight/workspace/%s/v1/objectschema/list`, insightURL, workspaceID)

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

	switch res.StatusCode {
	case http.StatusUnauthorized:
		return nil, fmt.Errorf("%s: %w", req.URL.String(), ErrUnauthorized)
	case http.StatusInternalServerError:
		return nil, fmt.Errorf("%s: %w", req.URL.String(), ErrUnknown)
	}

	list := new(insight.GenericList[insight.ObjectSchema])
	err = json.NewDecoder(res.Body).Decode(&list)

	return list, err
}

// GetAttributes find all object type attributes for this object schema
// Reference: https://developer.atlassian.com/cloud/insight/rest/api-group-objectschema/#api-objectschema-id-attributes-get
func (i *InsightObjectSchemaService) GetAttributes(ctx context.Context, workspaceID, id string, options *insight.ObjectSchemaAttributeOptions) ([]insight.ObjectTypeAttribute, error) {
	apiEndPoint := fmt.Sprintf(`%s/jsm/insight/workspace/%s/v1/objectschema/%s/attributes`, insightURL, workspaceID, id)
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

	switch res.StatusCode {
	case http.StatusUnauthorized:
		return nil, fmt.Errorf("%s: %w", req.URL.String(), ErrUnauthorized)
	case http.StatusNotFound:
		return nil, fmt.Errorf("%s: %w", req.URL.String(), ErrNotFound)
	case http.StatusInternalServerError:
		return nil, fmt.Errorf("%s: %w", req.URL.String(), ErrUnknown)
	}

	var attributes []insight.ObjectTypeAttribute
	err = json.NewDecoder(res.Body).Decode(&attributes)

	return attributes, err
}
