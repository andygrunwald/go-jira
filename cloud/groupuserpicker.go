package cloud

import (
	"context"
	"net/http"

	"github.com/mcl-de/go-jira/v2/cloud/models"
)

// GroupUserPickerService  handles Groups for the Jira instance / API.
//
// Jira API docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-group-and-user-picker/
type GroupUserPickerService service

// FindGroupAndUsers
//
// Jira API docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-group-and-user-picker/#api-group-group-and-user-picker
func (s *GroupUserPickerService) FindGroupAndUsers(ctx context.Context, opts models.FindGroupAndUsersQueryOptions) (*models.FoundUsersAndGroups, *Response, error) {
	u, err := addOptions("/rest/api/3/groupuserpicker", opts)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	found := new(models.FoundUsersAndGroups)
	resp, err := s.client.Do(req, found)
	if err != nil {
		return nil, resp, err
	}

	return found, resp, nil
}
