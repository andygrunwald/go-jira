package cloud

import (
	"fmt"
	"net/http"
)

// APITokenAuthTransport is an http.RoundTripper that authenticates all requests
// using a Personal API Token.
//
// Jira docs: https://support.atlassian.com/atlassian-account/docs/manage-api-tokens-for-your-atlassian-account/
// Create a token: https://id.atlassian.com/manage-profile/security/api-tokens
type APITokenAuthTransport struct {
	// Token is the API key.
	Token string

	// Transport is the underlying HTTP transport to use when making requests.
	// It will default to http.DefaultTransport if nil.
	Transport http.RoundTripper
}

// RoundTrip implements the RoundTripper interface. We just add the
// API token header and return the RoundTripper for this transport type.
func (t *APITokenAuthTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req2 := cloneRequest(req) // per RoundTripper contract

	req2.Header.Set("Authorization", fmt.Sprintf("Bearer %s", t.Token))
	return t.transport().RoundTrip(req2)
}

// Client returns an *http.Client that makes requests that are authenticated
// using the API token. This is a nice little bit of sugar
// so we can just get the client instead of creating the client in the calling code.
// If it's necessary to send more information on client init, the calling code can
// always skip this and set the transport itself.
func (t *APITokenAuthTransport) Client() *http.Client {
	return &http.Client{Transport: t}
}

func (t *APITokenAuthTransport) transport() http.RoundTripper {
	if t.Transport != nil {
		return t.Transport
	}
	return http.DefaultTransport
}
