package jira

import (
	"time"
	"strings"
)

const WLTimeFormat = "2006-01-02T15:04:05Z"
const WLDateFormat = "2006-01-02"

// SprintService handles sprints in JIRA Agile API.
// See https://docs.atlassian.com/jira-software/REST/cloud/
type WorklogService struct {
	client *Client
}

type TSWorkLogOptions struct {
	Username string  `url:"username"`
	DateFrom *WLDate `url:"dateFrom"`
	DateTo   *WLDate `url:"dateTo"`
}

// WorkLog time custom format
type WLTime struct {
	time.Time
}

// WorkLog date custom format
type WLDate struct {
	time.Time
}

// TimeSheet WorkLog represents one entry of a WorkLog
type TSWorkLog struct {
	Self             string   `json:"self,omitempty" structs:"self,omitempty"`
	Author           *User    `json:"author,omitempty" structs:"author,omitempty"`
	UpdateAuthor     *User    `json:"updateAuthor,omitempty" structs:"updateAuthor,omitempty"`
	Comment          string   `json:"comment,omitempty" structs:"comment,omitempty"`
	Created          *WLTime  `json:"dateCreated,omitempty" structs:"dateCreated,omitempty"`
	Updated          *WLTime  `json:"dateUpdated,omitempty" structs:"dateUpdated,omitempty"`
	Started          *WLTime  `json:"dateStarted,omitempty" structs:"dateStarted,omitempty"`
	TimeSpent        string   `json:"timeSpent,omitempty" structs:"timeSpent,omitempty"`
	TimeSpentSeconds int      `json:"timeSpentSeconds,omitempty" structs:"timeSpentSeconds,omitempty"`
	ID               int64    `json:"id,omitempty" structs:"id,omitempty"`
	Issue            *TSIssue `json:"issue,omitempty" structs:"issue,omitempty"`
}

// TimeSheet Issue object represents a JIRA issue.
type TSIssue struct {
	Expand    string       `json:"expand,omitempty" structs:"expand,omitempty"`
	ID        int64        `json:"id,omitempty" structs:"id,omitempty"`
	Self      string       `json:"self,omitempty" structs:"self,omitempty"`
	Key       string       `json:"key,omitempty" structs:"key,omitempty"`
	Fields    *IssueFields `json:"fields,omitempty" structs:"fields,omitempty"`
	Changelog *Changelog   `json:"changelog,omitempty" structs:"changelog,omitempty"`
}

// UnmarshalJSON will transform the TimeSheet WorkLog time into a time.Time
// during the transformation of the JSON response
func (t *WLTime) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), "\"")
	// Ignore null, like in the main JSON package.
	if s == "null" {
		t.Time = time.Time{}
		return
	}

	t.Time, err = time.Parse(WLTimeFormat, s)
	return
}

// UnmarshalJSON will transform the TimeSheet WorkLog date into a time.Time
// during the transformation of the JSON response
func (d *WLDate) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), "\"")
	// Ignore null, like in the main JSON package.
	if s == "null" {
		d.Time = time.Time{}
		return
	}

	d.Time, err = time.Parse(WLDateFormat, s)
	return
}

// GetWorkLogs returns worklogs for a user on date range
func (w *WorklogService) GetWorkLogs(options *TSWorkLogOptions) (*[]TSWorkLog, *Response, error) {
	apiEndpoint := "/rest/tempo-timesheets/3/worklogs"

	url, err := addOptions(apiEndpoint, options)
	req, err := w.client.NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}

	result := new([]TSWorkLog)
	resp, err := w.client.Do(req, result)
	if err != nil {
		err = NewJiraError(resp, err)
	}

	return result, resp, err
}
