package jira

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

// A Client manages communication with the JIRA API.
type Client struct {
	// HTTP client used to communicate with the API.
	client *http.Client

	// Base URL for API requests.
	baseURL *url.URL

	// Session storage if the user authentificate with a Session cookie
	session *Session

	// Services used for talking to different parts of the JIRA API.
	Authentication *AuthenticationService
	Issue          *IssueService
}

// NewClient returns a new JIRA API client.
// If a nil httpClient is provided, http.DefaultClient will be used.
// To use API methods which require authentication you can follow the preferred solution and
// provide an http.Client that will perform the authentication for you with OAuth and HTTP Basic (such as that provided by the golang.org/x/oauth2 library).
// As an alternative you can use Session Cookie based authentication provided by this package as well.
// See https://docs.atlassian.com/jira/REST/latest/#authentication
// baseURL is the HTTP endpoint of your JIRA instance and should always be specified with a trailing slash.
func NewClient(httpClient *http.Client, baseURL string) (*Client, error) {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	parsedBaseURL, err := url.Parse(baseURL)
	if err != nil {
		return nil, err
	}

	c := &Client{
		client:  httpClient,
		baseURL: parsedBaseURL,
	}
	c.Authentication = &AuthenticationService{client: c}
	c.Issue = &IssueService{client: c}

	return c, nil
}

// NewRequest creates an API request.
// A relative URL can be provided in urlStr, in which case it is resolved relative to the baseURL of the Client.
// Relative URLs should always be specified without a preceding slash.
// If specified, the value pointed to by body is JSON encoded and included as the request body.
func (c *Client) NewRequest(method, urlStr string, body interface{}) (*http.Request, error) {
	rel, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	u := c.baseURL.ResolveReference(rel)

	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	// Set session cookie if there is one
	if c.session != nil {
		req.Header.Set("Cookie", fmt.Sprintf("%s=%s", c.session.Session.Name, c.session.Session.Value))
	}

	return req, nil
}

// NewMultiPartRequest creates an API request including a multi-part file
// A relative URL can be provided in urlStr, in which case it is resolved relative to the baseURL of the Client.
// Relative URLs should always be specified without a preceding slash.
// If specified, the value pointed to by buf is a multipart form
func (c *Client) NewMultiPartRequest(method, urlStr string, buf *bytes.Buffer) (*http.Request, error) {
	rel, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	u := c.baseURL.ResolveReference(rel)

	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	// Set required headers
	req.Header.Set("X-Atlassian-Token", "nocheck")

	// Set session cookie if there is one
	if c.session != nil {
		req.Header.Set("Cookie", fmt.Sprintf("%s=%s", c.session.Session.Name, c.session.Session.Value))
	}

	return req, nil
}

// Do sends an API request and returns the API response.
// The API response is JSON decoded and stored in the value pointed to by v, or returned as an error if an API error has occurred.
func (c *Client) Do(req *http.Request, v interface{}) (*http.Response, error) {
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	err = CheckResponse(resp)
	if err != nil {
		// Even though there was an error, we still return the response
		// in case the caller wants to inspect it further
		return resp, err
	}

	if v != nil {
		// Open a NewDecoder and defer closing the reader only if there is a provided interface to decode to
		defer resp.Body.Close()
		err = json.NewDecoder(resp.Body).Decode(v)
	}

	return resp, err
}

// Do sends an API request and returns the API response.
// The API response is JSON decoded and stored in the value pointed to by v, or returned as an error if an API error has occurred.
// The caller is expected to consume the Body, and needs to call Body.Close when the response has been handled
func (c *Client) DoNoClose(req *http.Request, v interface{}) (*http.Response, error) {
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	err = CheckResponse(resp)
	if err != nil {
		// Even though there was an error, we still return the response
		// in case the caller wants to inspect it further
		return resp, err
	}

	if v != nil {
		err = json.NewDecoder(resp.Body).Decode(v)
	}

	return resp, err
}

//// Authenticated reports if the current Client has an authenticated session with JIRA
//func (c *Client) Authenticated() bool {
//	if c != nil {
//		return c.session != nil
//	} else {
//		return false
//	}
//}

// CheckResponse checks the API response for errors, and returns them if present.
// A response is considered an error if it has a status code outside the 200 range.
// API error responses are expected to have either no response body, or a JSON response body that maps to ErrorResponse.
// Any other response body will be silently ignored.
func CheckResponse(r *http.Response) error {
	if c := r.StatusCode; 200 <= c && c <= 299 {
		return nil
	}

	errorResponse := &ErrorResponse{Response: r}
	data, err := ioutil.ReadAll(r.Body)
	if err == nil && data != nil {
		json.Unmarshal(data, errorResponse)
	}
	return errorResponse
}
