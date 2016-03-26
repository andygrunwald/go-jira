package jira

import (
	"fmt"
	"net/http"
)

// ErrorResponse reports one or more errors caused by an API request.
type ErrorResponse struct {
	Response      *http.Response    // HTTP response that caused this error
	ErrorMessages []string          `json:"errorMessages,omitempty"`
	Errors        map[string]string `json:"errors,omitempty"`
}

func (r *ErrorResponse) Error() string {
	if r.Response == nil {
		return fmt.Sprintf("%v %+v", r.ErrorMessages, r.Errors)
	}
	return fmt.Sprintf("%v %v: %d %v %+v",
		r.Response.Request.Method, r.Response.Request.URL,
		r.Response.StatusCode, r.ErrorMessages, r.Errors)
}
