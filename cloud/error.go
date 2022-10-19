package cloud

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
)

// JiraRestError error message from Jira REST API
// See https://docs.atlassian.com/jira/REST/cloud/#error-responses
type JiraRestError struct {
	ErrorMessages []string          `json:"errorMessages"`
	Errors        map[string]string `json:"errors"`
}

var ErrValidation = &ResponseError{statusCode: http.StatusBadRequest}
var ErrUnauthorized = &ResponseError{statusCode: http.StatusUnauthorized}
var ErrNotFound = &ResponseError{statusCode: http.StatusNotFound}
var ErrUnknown = &ResponseError{statusCode: http.StatusInternalServerError}
var ErrNoBody = errors.New("no body in response error")

type ResponseError struct {
	status     string
	statusCode int
	url        *url.URL
	body       []byte
}

// Error is a short string representing the error
func (r *ResponseError) Error() string {
	return fmt.Sprintf("%s: %d %s", r.url.String(), r.statusCode, r.status)
}

// Body returns the response body if present
func (r *ResponseError) Body() []byte {
	return r.body
}

func (r *ResponseError) JiraRESTError() (*JiraRestError, error) {
	if r.body == nil {
		return nil, ErrNoBody
	}

	var jiraErr JiraRestError
	err := json.Unmarshal(r.body, &jiraErr)
	return &jiraErr, err
}

// Is checks if the error status code is equal
func (r *ResponseError) Is(tgt error) bool {
	target, ok := tgt.(*ResponseError)
	if !ok {
		return false
	}

	return r.statusCode == target.statusCode
}
