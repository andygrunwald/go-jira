package cloud

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/mcl-de/go-jira/v2/cloud/model/apps/assets"
)

// AssetsIconService handles Icon in Assets App for the Jira instance / API.
type AssetsIconService service

// Get returns single icon by id
// Reference: https://developer.atlassian.com/cloud/assets/rest/api-group-iql/#api-iql-objects-get
func (a *AssetsIconService) Get(ctx context.Context, workspaceID, id string) (*assets.Icon, error) {
	apiEndPoint := fmt.Sprintf(`%s/jsm/assets/workspace/%s/v1/icon/%s`, assetURL, workspaceID, id)

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

	icon := new(assets.Icon)
	err = json.NewDecoder(res.Body).Decode(&icon)

	return icon, err
}

// GetGlobal returns all global icons i.e. icons not associated with a particular object schema
// Reference: https://developer.atlassian.com/cloud/assets/rest/api-group-iql/#api-iql-objects-get
func (a *AssetsIconService) GetGlobal(ctx context.Context, workspaceID string) ([]assets.Icon, error) {
	apiEndPoint := fmt.Sprintf(`%s/jsm/assets/workspace/%s/v1/icon/global`, assetURL, workspaceID)

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

	icons := []assets.Icon{}
	err = json.NewDecoder(res.Body).Decode(&icons)

	return icons, err
}
