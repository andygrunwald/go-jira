package onpremise

import (
	"fmt"
	"net/http"
)

// BearerAuthTransport is a http.RoundTripper that authenticates all requests
// using Jira's bearer (oauth 2.0 (3lo)) based authentication.
type BearerAuthTransport struct {
	Token string

	// Transport is the underlying HTTP transport to use when making requests.
	// It will default to http.DefaultTransport if nil.
	Transport http.RoundTripper
}

// RoundTrip implements the RoundTripper interface.  We just add the
// bearer token and return the RoundTripper for this transport type.
func (t *BearerAuthTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req2 := cloneRequest(req) // per RoundTripper contract

	req2.Header.Set("Authorization", fmt.Sprintf("Bearer %s", t.Token))
	return t.transport().RoundTrip(req2)
}

// Client returns an *http.Client that makes requests that are authenticated
// using HTTP Basic Authentication.  This is a nice little bit of sugar
// so we can just get the client instead of creating the client in the calling code.
// If it's necessary to send more information on client init, the calling code can
// always skip this and set the transport itself.
func (t *BearerAuthTransport) Client() *http.Client {
	return &http.Client{Transport: t}
}

func (t *BearerAuthTransport) transport() http.RoundTripper {
	if t.Transport != nil {
		return t.Transport
	}
	return http.DefaultTransport
}
