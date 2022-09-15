package cloud

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/mcl-de/go-jira/v2/cloud/model/apps/insight"
)

// InsightIconService handles Icon in Insight App for the Jira instance / API.
type InsightIconService service

// Get returns single icon by id
// Reference: https://developer.atlassian.com/cloud/insight/rest/api-group-iql/#api-iql-objects-get
func (i *InsightIconService) Get(ctx context.Context, workspaceID, id string) (*insight.Icon, error) {
	apiEndPoint := fmt.Sprintf(`%s/jsm/insight/workspace/%s/v1/icon/%s`, insightURL, workspaceID, id)

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
	case http.StatusNotFound:
		return nil, fmt.Errorf("%s: %w", req.URL.String(), ErrNotFound)
	case http.StatusInternalServerError:
		return nil, fmt.Errorf("%s: %w", req.URL.String(), ErrUnknown)
	}

	icon := new(insight.Icon)
	err = json.NewDecoder(res.Body).Decode(&icon)

	return icon, err
}

// GetGlobal returns all global icons i.e. icons not associated with a particular object schema
// Reference: https://developer.atlassian.com/cloud/insight/rest/api-group-iql/#api-iql-objects-get
func (i *InsightIconService) GetGlobal(ctx context.Context, workspaceID string) ([]insight.Icon, error) {
	apiEndPoint := fmt.Sprintf(`%s/jsm/insight/workspace/%s/v1/icon/global`, insightURL, workspaceID)

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

	icons := []insight.Icon{}
	err = json.NewDecoder(res.Body).Decode(&icons)

	return icons, err
}
