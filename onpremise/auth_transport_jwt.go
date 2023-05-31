package onpremise

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"net/http"
	"net/url"
	"sort"
	"strings"
	"time"

	jwt "github.com/golang-jwt/jwt/v4"
)

// JWTAuthTransport is an http.RoundTripper that authenticates all requests
// using Jira's JWT based authentication.
//
// NOTE: this form of auth should be used by add-ons installed from the Atlassian marketplace.
//
// Jira docs: https://developer.atlassian.com/cloud/jira/platform/understanding-jwt
// Examples in other languages:
//
//	https://bitbucket.org/atlassian/atlassian-jwt-ruby/src/d44a8e7a4649e4f23edaa784402655fda7c816ea/lib/atlassian/jwt.rb
//	https://bitbucket.org/atlassian/atlassian-jwt-py/src/master/atlassian_jwt/url_utils.py
type JWTAuthTransport struct {
	Secret []byte
	Issuer string

	// Transport is the underlying HTTP transport to use when making requests.
	// It will default to http.DefaultTransport if nil.
	Transport http.RoundTripper
}

func (t *JWTAuthTransport) Client() *http.Client {
	return &http.Client{Transport: t}
}

func (t *JWTAuthTransport) transport() http.RoundTripper {
	if t.Transport != nil {
		return t.Transport
	}
	return http.DefaultTransport
}

// RoundTrip adds the session object to the request.
func (t *JWTAuthTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req2 := cloneRequest(req) // per RoundTripper contract
	exp := time.Duration(59) * time.Second
	qsh := t.createQueryStringHash(req.Method, req2.URL)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss": t.Issuer,
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(exp).Unix(),
		"qsh": qsh,
	})

	jwtStr, err := token.SignedString(t.Secret)
	if err != nil {
		return nil, fmt.Errorf("jwtAuth: error signing JWT: %w", err)
	}

	req2.Header.Set("Authorization", fmt.Sprintf("JWT %s", jwtStr))
	return t.transport().RoundTrip(req2)
}

func (t *JWTAuthTransport) createQueryStringHash(httpMethod string, jiraURL *url.URL) string {
	canonicalRequest := t.canonicalizeRequest(httpMethod, jiraURL)
	h := sha256.Sum256([]byte(canonicalRequest))
	return hex.EncodeToString(h[:])
}

func (t *JWTAuthTransport) canonicalizeRequest(httpMethod string, jiraURL *url.URL) string {
	path := "/" + strings.Replace(strings.Trim(jiraURL.Path, "/"), "&", "%26", -1)

	var canonicalQueryString []string
	for k, v := range jiraURL.Query() {
		if k == "jwt" {
			continue
		}
		param := url.QueryEscape(k)
		value := url.QueryEscape(strings.Join(v, ""))
		canonicalQueryString = append(canonicalQueryString, strings.Replace(strings.Join([]string{param, value}, "="), "+", "%20", -1))
	}
	sort.Strings(canonicalQueryString)
	return fmt.Sprintf("%s&%s&%s", strings.ToUpper(httpMethod), path, strings.Join(canonicalQueryString, "&"))
}
