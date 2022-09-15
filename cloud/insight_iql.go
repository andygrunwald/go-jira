package cloud

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/mcl-de/go-jira/v2/cloud/model/apps/insight"
)

// InsightIQLService handles Insight Query Language (IQL) in Insight App for the Jira instance / API.
type InsightIQLService service

// GetObjects returns found objects based on Insight Query Language (IQL)
// Reference: https://developer.atlassian.com/cloud/insight/rest/api-group-iql/#api-iql-objects-get
func (i *InsightIQLService) GetObjects(ctx context.Context, workspaceID string, options insight.IQLGetObjectsOptions) (*insight.ObjectListResult, error) {
	apiEndPoint := fmt.Sprintf(`%s/jsm/insight/workspace/%s/v1/iql/objects`, insightURL, workspaceID)

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

	result := new(insight.ObjectListResult)
	err = json.NewDecoder(res.Body).Decode(&result)

	return result, err
}
