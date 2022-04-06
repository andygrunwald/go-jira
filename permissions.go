package jira

import (
	"context"
)

// PermissionsService handles projects for the Jira instance / API.
//
// Jira API docs: https://developer.atlassian.com/cloud/jira/platform/rest/v2/api-group-permissions/
type PermissionsService struct {
	client *Client
}

type BulkProjectPermissions struct {
	Permission   string `json:"permission"`
	ProjectIDs   []int  `json:"projects"`
	IssueTypeIDs []int  `json:"issues"`
}

type BulkPermissions struct {
	ProjectPermissions []*BulkProjectPermissions `json:"projectPermissions"`
	GlobalPermissions  []string                  `json:"globalPermissions"`
}
type requestProjectPermission struct {
	Permissions  []string `json:"permissions"`
	ProjectIDs   []int    `json:"projects"`
	IssueTypeIDs []int    `json:"issues"`
}

type requestBulkPermissions struct {
	AccountID          string                     `json:"accountId,omitempty"`
	GlobalPermissions  []string                   `json:"globalPermissions,omitempty"`
	ProjectPermissions []requestProjectPermission `json:"projectPermissions,omitempty"`
}

// GetBulkPermissionsWithContext takes a list of projects with issue type IDs and
// returns permissions given user (or currently-authed user if that is not set)
// is granted on those projects and issue types. globalPermissions can also be
// passed, which will check non-project-specific permissions too.
//
// https://developer.atlassian.com/cloud/jira/platform/rest/v2/api-group-permissions/#api-rest-api-2-permissions-check-post
func (s *PermissionsService) GetBulkPermissionsWithContext(ctx context.Context, accountID string, globalPermissions, permissions []string, projectIDs []int, issueTypeIDs []int) (*BulkPermissions, *Response, error) {
	reqBody := requestBulkPermissions{
		AccountID:         accountID,
		GlobalPermissions: globalPermissions,
		ProjectPermissions: []requestProjectPermission{{
			Permissions:  permissions,
			ProjectIDs:   projectIDs,
			IssueTypeIDs: issueTypeIDs,
		}},
	}

	req, err := s.client.NewRequestWithContext(ctx, "POST", "/rest/api/2/permissions/check", reqBody)
	if err != nil {
		return nil, nil, err
	}

	bs := new(BulkPermissions)
	resp, err := s.client.Do(req, bs)
	if err != nil {
		return nil, resp, NewJiraError(resp, err)
	}

	return bs, resp, nil
}

func (s *PermissionsService) GetBulkPermissions(accountID string, globalPermissions, permissions []string, projectIDs, issueTypeIDs []int) (*BulkPermissions, *Response, error) {
	return s.GetBulkPermissionsWithContext(context.Background(), accountID, globalPermissions, permissions, projectIDs, issueTypeIDs)
}
