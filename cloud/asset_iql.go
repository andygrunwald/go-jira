package cloud

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/mcl-de/go-jira/v2/cloud/model/apps/assets"
)

// InsightIQLService handles Insight Query Language (IQL) in Insight App for the Jira instance / API.
type AssetsIQLService service

// GetObjects returns found objects based on Insight Query Language (IQL)
// Reference: https://developer.atlassian.com/cloud/assets/rest/api-group-aql/#api-iql-objects-get
func (a *AssetsIQLService) GetObjects(ctx context.Context, workspaceID string, options assets.IQLGetObjectsOptions) (*assets.ObjectListResult, error) {
	apiEndPoint := fmt.Sprintf(`%s/jsm/assets/workspace/%s/v1/iql/objects`, assetURL, workspaceID)
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

	result := new(assets.ObjectListResult)
	err = json.NewDecoder(res.Body).Decode(&result)

	return result, err
}
