package jira

import (
	"fmt"
	"net/http"
)

const (
	// AssigneeAutomatic represents the value of the "Assignee: Automatic" of JIRA
	AssigneeAutomatic = "-1"
)

// IssueService handles Issues for the JIRA instance / API.
//
// JIRA API docs: https://docs.atlassian.com/jira/REST/latest/#api/2/issue
type IssueService struct {
	client *Client
}

// Issue represents a JIRA issue.
type Issue struct {
	Expand string       `json:"expand,omitempty"`
	ID     string       `json:"id,omitempty"`
	Self   string       `json:"self,omitempty"`
	Key    string       `json:"key,omitempty"`
	Fields *IssueFields `json:"fields,omitempty"`
}

// IssueFields represents single fields of a JIRA issue.
// Every JIRA issue has several fields attached.
type IssueFields struct {
	// TODO Missing fields
	//	* "timespent": null,
	//	* "aggregatetimespent": null,
	//	* "workratio": -1,
	//	* "lastViewed": null,
	//	* "labels": [],
	//	* "timeestimate": null,
	//	* "aggregatetimeoriginalestimate": null,
	//	* "timeoriginalestimate": null,
	//	* "timetracking": {},
	//	* "attachment": [],
	//	* "aggregatetimeestimate": null,
	//	* "subtasks": [],
	//	* "environment": null,
	//	* "duedate": null,
	Type              IssueType     `json:"issuetype"`
	Project           Project       `json:"project,omitempty"`
	Resolution        *Resolution   `json:"resolution,omitempty"`
	Priority          *Priority     `json:"priority,omitempty"`
	Resolutiondate    string        `json:"resolutiondate,omitempty"`
	Created           string        `json:"created,omitempty"`
	Watches           *Watches      `json:"watches,omitempty"`
	Assignee          *Assignee     `json:"assignee,omitempty"`
	Updated           string        `json:"updated,omitempty"`
	Description       string        `json:"description,omitempty"`
	Summary           string        `json:"summary"`
	Creator           *Assignee     `json:"Creator,omitempty"`
	Reporter          *Assignee     `json:"reporter,omitempty"`
	Components        []*Component  `json:"components,omitempty"`
	Status            *Status       `json:"status,omitempty"`
	Progress          *Progress     `json:"progress,omitempty"`
	AggregateProgress *Progress     `json:"aggregateprogress,omitempty"`
	Worklog           []*Worklog    `json:"worklog.worklogs,omitempty"`
	IssueLinks        []*IssueLink  `json:"issuelinks,omitempty"`
	Comments          []*Comment    `json:"comment.comments,omitempty"`
	FixVersions       []*FixVersion `json:"fixVersions,omitempty"`
}

// IssueType represents a type of a JIRA issue.
// Typical types are "Request", "Bug", "Story", ...
type IssueType struct {
	Self        string `json:"self,omitempty"`
	ID          string `json:"id,omitempty"`
	Description string `json:"description,omitempty"`
	IconURL     string `json:"iconUrl,omitempty"`
	Name        string `json:"name,omitempty"`
	Subtask     bool   `json:"subtask,omitempty"`
}

// Project represents a JIRA Project.
type Project struct {
	Self       string            `json:"self,omitempty"`
	ID         string            `json:"id,omitempty"`
	Key        string            `json:"key,omitempty"`
	Name       string            `json:"name,omitempty"`
	AvatarURLs map[string]string `json:"avatarUrls,omitempty"`
}

// Resolution represents a resolution of a JIRA issue.
// Typical types are "Fixed", "Suspended", "Won't Fix", ...
type Resolution struct {
	Self        string `json:"self"`
	ID          string `json:"id"`
	Description string `json:"description"`
	Name        string `json:"name"`
}

// Priority represents a priority of a JIRA issue.
// Typical types are "Normal", "Moderate", "Urgent", ...
type Priority struct {
	Self    string `json:"self,omitempty"`
	IconURL string `json:"iconUrl,omitempty"`
	Name    string `json:"name,omitempty"`
	ID      string `json:"id,omitempty"`
}

// Watches represents a type of how many user are "observing" a JIRA issue to track the status / updates.
type Watches struct {
	Self       string `json:"self,omitempty"`
	WatchCount int    `json:"watchCount,omitempty"`
	IsWatching bool   `json:"isWatching,omitempty"`
}

// Assignee represents a user who is this JIRA issue assigned to.
type Assignee struct {
	Self         string            `json:"self,omitempty"`
	Name         string            `json:"name,omitempty"`
	EmailAddress string            `json:"emailAddress,omitempty"`
	AvatarURLs   map[string]string `json:"avatarUrls,omitempty"`
	DisplayName  string            `json:"displayName,omitempty"`
	Active       bool              `json:"active,omitempty"`
}

// Component represents a "component" of a JIRA issue.
// Components can be user defined in every JIRA instance.
type Component struct {
	Self string `json:"self,omitempty"`
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

// Status represents the current status of a JIRA issue.
// Typical status are "Open", "In Progress", "Closed", ...
// Status can be user defined in every JIRA instance.
type Status struct {
	Self           string         `json:"self"`
	Description    string         `json:"description"`
	IconURL        string         `json:"iconUrl"`
	Name           string         `json:"name"`
	ID             string         `json:"id"`
	StatusCategory StatusCategory `json:"statusCategory"`
}

// StatusCategory represents the category a status belongs to.
// Those categories can be user defined in every JIRA instance.
type StatusCategory struct {
	Self      string `json:"self"`
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Key       string `json:"key"`
	ColorName string `json:"colorName"`
}

// Progress represents the progress of a JIRA issue.
type Progress struct {
	Progress int `json:"progress"`
	Total    int `json:"total"`
}

// Worklog represents the work log of a JIRA issue.
// JIRA Wiki: https://confluence.atlassian.com/jira/logging-work-on-an-issue-185729605.html
type Worklog struct {
	// TODO Add Worklogs
}

// IssueLink represents a link between two issues in JIRA.
type IssueLink struct {
	ID           string        `json:"id"`
	Self         string        `json:"self"`
	Type         IssueLinkType `json:"type"`
	OutwardIssue Issue         `json:"outwardIssue"`
	InwardIssue  Issue         `json:"inwardIssue"`
	Comment      Comment       `json:"comment"`
}

// IssueLinkType represents a type of a link between to issues in JIRA.
// Typical issue link types are "Related to", "Duplicate", "Is blocked by", etc.
type IssueLinkType struct {
	ID      string `json:"id"`
	Self    string `json:"self"`
	Name    string `json:"name"`
	Inward  string `json:"inward"`
	Outward string `json:"outward"`
}

// Comment represents a comment by a person to an issue in JIRA.
type Comment struct {
	Self         string            `json:"self"`
	Name         string            `json:"name"`
	Author       Assignee          `json:"author"`
	Body         string            `json:"body"`
	UpdateAuthor Assignee          `json:"updateAuthor"`
	Updated      string            `json:"updated"`
	Created      string            `json:"created"`
	Visibility   CommentVisibility `json:"visibility"`
}

// FixVersion represents a software release in which an issue is fixed.
type FixVersion struct {
	Archived        *bool  `json:"archived,omitempty"`
	ID              string `json:"id,omitempty"`
	Name            string `json:"name,omitempty"`
	ProjectID       int    `json:"projectId,omitempty"`
	ReleaseDate     string `json:"releaseDate,omitempty"`
	Released        *bool  `json:"released,omitempty"`
	Self            string `json:"self,omitempty"`
	UserReleaseDate string `json:"userReleaseDate,omitempty"`
}

// CommentVisibility represents he visibility of a comment.
// E.g. Type could be "role" and Value "Administrators"
type CommentVisibility struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

// Get returns a full representation of the issue for the given issue key.
// JIRA will attempt to identify the issue by the issueIdOrKey path parameter.
// This can be an issue id, or an issue key.
// If the issue cannot be found via an exact match, JIRA will also look for the issue in a case-insensitive way, or by looking to see if the issue was moved.
//
// JIRA API docs: https://docs.atlassian.com/jira/REST/latest/#api/2/issue-getIssue
func (s *IssueService) Get(issueID string) (*Issue, *http.Response, error) {
	apiEndpoint := fmt.Sprintf("rest/api/2/issue/%s", issueID)
	req, err := s.client.NewRequest("GET", apiEndpoint, nil)
	if err != nil {
		return nil, nil, err
	}

	issue := new(Issue)
	resp, err := s.client.Do(req, issue)
	if err != nil {
		return nil, resp, err
	}

	return issue, resp, nil
}

// Create creates an issue or a sub-task from a JSON representation.
// Creating a sub-task is similar to creating a regular issue, with two important differences:
// The issueType field must correspond to a sub-task issue type and you must provide a parent field in the issue create request containing the id or key of the parent issue.
//
// JIRA API docs: https://docs.atlassian.com/jira/REST/latest/#api/2/issue-createIssues
func (s *IssueService) Create(issue *Issue) (*Issue, *http.Response, error) {
	apiEndpoint := "rest/api/2/issue/"
	req, err := s.client.NewRequest("POST", apiEndpoint, issue)
	if err != nil {
		return nil, nil, err
	}

	responseIssue := new(Issue)
	resp, err := s.client.Do(req, responseIssue)
	if err != nil {
		return nil, resp, err
	}

	return responseIssue, resp, nil
}

// AddComment adds a new comment to issueID.
//
// JIRA API docs: https://docs.atlassian.com/jira/REST/latest/#api/2/issue-addComment
func (s *IssueService) AddComment(issueID string, comment *Comment) (*Comment, *http.Response, error) {
	apiEndpoint := fmt.Sprintf("rest/api/2/issue/%s/comment", issueID)
	req, err := s.client.NewRequest("POST", apiEndpoint, comment)
	if err != nil {
		return nil, nil, err
	}

	responseComment := new(Comment)
	resp, err := s.client.Do(req, responseComment)
	if err != nil {
		return nil, resp, err
	}

	return responseComment, resp, nil
}

// AddLink adds a link between two issues.
//
// JIRA API docs: https://docs.atlassian.com/jira/REST/latest/#api/2/issueLink
func (s *IssueService) AddLink(issueLink *IssueLink) (*http.Response, error) {
	apiEndpoint := fmt.Sprintf("rest/api/2/issueLink")
	req, err := s.client.NewRequest("POST", apiEndpoint, issueLink)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req, nil)
	return resp, err
}
