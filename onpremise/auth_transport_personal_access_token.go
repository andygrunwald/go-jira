package onpremise

import "net/http"

// PATAuthTransport is an http.RoundTripper that authenticates all requests
// using the Personal Access Token specified.
// See here for more info: https://confluence.atlassian.com/enterprise/using-personal-access-tokens-1026032365.html
type PATAuthTransport struct {
	// Token is the key that was provided by Jira when creating the Personal Access Token.
	Token string

	// Transport is the underlying HTTP transport to use when making requests.
	// It will default to http.DefaultTransport if nil.
	Transport http.RoundTripper
}

// RoundTrip implements the RoundTripper interface.  We just add the
// basic auth and return the RoundTripper for this transport type.
func (t *PATAuthTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req2 := cloneRequest(req) // per RoundTripper contract
	req2.Header.Set("Authorization", "Bearer "+t.Token)
	return t.transport().RoundTrip(req2)
}

// Client returns an *http.Client that makes requests that are authenticated
// using HTTP Basic Authentication.  This is a nice little bit of sugar
// so we can just get the client instead of creating the client in the calling code.
// If it's necessary to send more information on client init, the calling code can
// always skip this and set the transport itself.
func (t *PATAuthTransport) Client() *http.Client {
	return &http.Client{Transport: t}
}

func (t *PATAuthTransport) transport() http.RoundTripper {
	if t.Transport != nil {
		return t.Transport
	}
	return http.DefaultTransport
}
