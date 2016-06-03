package jira

import (
	"fmt"
	"net/http"
)

// AuthenticationService handles authentication for the JIRA instance / API.
//
// JIRA API docs: https://docs.atlassian.com/jira/REST/latest/#authentication
type AuthenticationService struct {
	client *Client
}

// Session represents a Session JSON response by the JIRA API.
type Session struct {
	Self    string `json:"self,omitempty"`
	Name    string `json:"name,omitempty"`
	Session struct {
		Name  string `json:"name"`
		Value string `json:"value"`
	} `json:"session,omitempty"`
	LoginInfo struct {
		FailedLoginCount    int    `json:"failedLoginCount"`
		LoginCount          int    `json:"loginCount"`
		LastFailedLoginTime string `json:"lastFailedLoginTime"`
		PreviousLoginTime   string `json:"previousLoginTime"`
	} `json:"loginInfo"`
	Cookies []*http.Cookie
}

// AcquireSessionCookie creates a new session for a user in JIRA.
// Once a session has been successfully created it can be used to access any of JIRA's remote APIs and also the web UI by passing the appropriate HTTP Cookie header.
// The header will by automatically applied to every API request.
// Note that it is generally preferrable to use HTTP BASIC authentication with the REST API.
// However, this resource may be used to mimic the behaviour of JIRA's log-in page (e.g. to display log-in errors to a user).
//
// JIRA API docs: https://docs.atlassian.com/jira/REST/latest/#auth/1/session
func (s *AuthenticationService) AcquireSessionCookie(username, password string) (bool, error) {
	apiEndpoint := "rest/auth/1/session"
	body := struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}{
		username,
		password,
	}

	req, err := s.client.NewRequest("POST", apiEndpoint, body)
	if err != nil {
		return false, err
	}

	session := new(Session)
	resp, err := s.client.Do(req, session)
	session.Cookies = resp.Cookies()

	if err != nil {
		return false, fmt.Errorf("Auth at JIRA instance failed (HTTP(S) request). %s", err)
	}
	if resp != nil && resp.StatusCode != 200 {
		return false, fmt.Errorf("Auth at JIRA instance failed (HTTP(S) request). Status code: %d", resp.StatusCode)
	}

	s.client.session = session

	return true, nil
}

// Authenticated reports if the current Client has an authenticated session with JIRA
func (s *AuthenticationService) Authenticated() bool {
	if s != nil {
		return s.client.session != nil
	}
	return false
}

// TODO Missing API Call GET (Returns information about the currently authenticated user's session)
// See https://docs.atlassian.com/jira/REST/latest/#auth/1/session
// TODO Missing API Call DELETE (Logs the current user out of JIRA, destroying the existing session, if any.)
// See https://docs.atlassian.com/jira/REST/latest/#auth/1/session
