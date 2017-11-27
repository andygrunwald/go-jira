package jira

import (
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

const (
	// HTTP Basic Authentication
	authTypeBasic = 1
	// HTTP Session Authentication
	authTypeSession = 2
	// Oauth Authentication
	authTypeOauth = 3
)

// AuthenticationService handles authentication for the JIRA instance / API.
//
// JIRA API docs: https://docs.atlassian.com/jira/REST/latest/#authentication
type AuthenticationService struct {
	client *Client

	// Authentication type
	authType int

	// Basic auth username
	username string

	// Basic auth password
	password string

	// Oauth access token
	accessToken string
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

// JSON response from the Jira Auth server: https://auth.atlassian.io/oauth2/token
type JiraAccessToken struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
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

	if resp != nil {
		session.Cookies = resp.Cookies()
	}

	if err != nil {
		return false, fmt.Errorf("Auth at JIRA instance failed (HTTP(S) request). %s", err)
	}
	if resp != nil && resp.StatusCode != 200 {
		return false, fmt.Errorf("Auth at JIRA instance failed (HTTP(S) request). Status code: %d", resp.StatusCode)
	}

	s.client.session = session
	s.authType = authTypeSession

	return true, nil
}

// Sets the access token on the Authentication Service (helpful in event access tokens are cached by the client code).
// Use GetOauth2AccessToken to get a new access token and set authType to Oauth
func (s *AuthenticationService) SetOauth(accessToken string) {
	s.authType = authTypeOauth
	s.accessToken = accessToken
}

// SetBasicAuth sets username and password for the basic auth against the JIRA instance.
func (s *AuthenticationService) SetBasicAuth(username, password string) {
	s.username = username
	s.password = password
	s.authType = authTypeBasic
}

// Authenticated reports if the current Client has authentication details for JIRA
func (s *AuthenticationService) Authenticated() bool {
	if s != nil {
		if s.authType == authTypeSession {
			return s.client.session != nil
		} else if s.authType == authTypeBasic {
			return s.username != ""
		} else if s.authType == authTypeOauth {
			return s.accessToken != ""
		}

	}
	return false
}

// Logout logs out the current user that has been authenticated and the session in the client is destroyed.
//
// JIRA API docs: https://docs.atlassian.com/jira/REST/latest/#auth/1/session
func (s *AuthenticationService) Logout() error {
	if s.authType != authTypeSession || s.client.session == nil {
		return fmt.Errorf("No user is authenticated yet.")
	}

	apiEndpoint := "rest/auth/1/session"
	req, err := s.client.NewRequest("DELETE", apiEndpoint, nil)
	if err != nil {
		return fmt.Errorf("Creating the request to log the user out failed : %s", err)
	}

	resp, err := s.client.Do(req, nil)
	if err != nil {
		return fmt.Errorf("Error sending the logout request: %s", err)
	}
	if resp.StatusCode != 204 {
		return fmt.Errorf("The logout was unsuccessful with status %d", resp.StatusCode)
	}

	// If logout successful, delete session
	s.client.session = nil

	return nil

}

// GetCurrentUser gets the details of the current user.
//
// JIRA API docs: https://docs.atlassian.com/jira/REST/latest/#auth/1/session
func (s *AuthenticationService) GetCurrentUser() (*Session, error) {
	if s == nil {
		return nil, fmt.Errorf("AUthenticaiton Service is not instantiated")
	}
	if s.authType != authTypeSession || s.client.session == nil {
		return nil, fmt.Errorf("No user is authenticated yet")
	}

	apiEndpoint := "rest/auth/1/session"
	req, err := s.client.NewRequest("GET", apiEndpoint, nil)
	if err != nil {
		return nil, fmt.Errorf("Could not create request for getting user info : %s", err)
	}

	resp, err := s.client.Do(req, nil)
	if err != nil {
		return nil, fmt.Errorf("Error sending request to get user info : %s", err)
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("Getting user info failed with status : %d", resp.StatusCode)
	}

	defer resp.Body.Close()
	ret := new(Session)
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Couldn't read body from the response : %s", err)
	}

	err = json.Unmarshal(data, &ret)

	if err != nil {
		return nil, fmt.Errorf("Could not unmarshall received user info : %s", err)
	}

	return ret, nil
}

// Attempts to retrieve an Oauth2 access token from the Jira Auth Server. Returns the details of the access token
// in a JiraAccessToken struct so that it can be cached by the client application if desired (oauth2 tokens expire
// after 15 min by default); else, this can be called multiple times to retrieve a new access token.
//
// oauthClientId, userKey, and sharedSecret are returned by Jira during the installation of the Add-on
// scope is the scope to place on the AccessToken returned (e.g., "READ", "READ WRITE", "READ ACT_AS_USER")
// See: https://developer.atlassian.com/cloud/jira/software/oauth-2-jwt-bearer-token-authorization-grant-type
func (s *AuthenticationService) GetOauth2AccessToken(oauthClientId, userKey, scope string, sharedSecret []byte) (*JiraAccessToken, error) {
	// create JSON Web Token
	jiraAuthURL := "https://auth.atlassian.io"
	jwtStr, err := createJWT(oauthClientId, s.client.baseURL.String(), jiraAuthURL, userKey, sharedSecret)
	if err != nil {
		return nil, fmt.Errorf("Error creating JSON Web Token : %s", err)
	}

	// Now make a request to the Jira auth server for an access token
	oauth2Route := fmt.Sprintf("%s/oauth2/token", jiraAuthURL)
	grantType := "urn:ietf:params:oauth:grant-type:jwt-bearer"
	form := url.Values{"grant_type": {grantType}, "assertion": {jwtStr}, "scope": {scope}}
	formBody := strings.NewReader(form.Encode())
	req, err := http.NewRequest("POST", oauth2Route, formBody)
	if err != nil {
		return nil, fmt.Errorf("Error creating Jira Auth Server request : %s", err)
	}
	req.Header.Set("Content-Length", strconv.FormatInt(int64(len(form.Encode())), 10))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=utf-8")
	req.Header.Set("alg", "HS256")
	req.Header.Set("typ", "JWT")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Error getting Oauth2 access token from %s : %s", oauth2Route, err)
	}

	// Finally, decode the JSON from the response
	defer resp.Body.Close()
	token := new(JiraAccessToken)
	if err := json.NewDecoder(resp.Body).Decode(token); err != nil {
		return nil, fmt.Errorf("Error decoding JSON from response %+v : %s", resp, err)
	}

	// set oauth type before leaving
	s.SetOauth(token.AccessToken)

	return token, nil
}

// Creates a JSON Web Token that will be used as an assertion to get an Oauth2 access token
// See: https://developer.atlassian.com/cloud/jira/software/understanding-jwt
func createJWT(oauthClientId, instanceURL, jiraAuthURL, userKey string, sharedSecret []byte) (string, error) {
	jwt := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss": fmt.Sprintf("urn:atlassian:connect:clientid:%s", oauthClientId),
		"tnt": instanceURL,
		"sub": fmt.Sprintf("urn:atlassian:connect:userkey:%s", userKey),
		"aud": jiraAuthURL,
		"iat": time.Now().Unix(),
		"exp": time.Now().Unix() + 10, // expires in 10 seconds
	})

	// sign the jwt using the signing method and the shared secret
	if signingStr, err := jwt.SigningString(); err != nil {
		return "", err
	} else {
		if _, err = jwt.Method.Sign(signingStr, sharedSecret); err != nil {
			return "", err
		}
	}

	// now return the signed JWT string to be included in the form params
	return jwt.SignedString(sharedSecret)
}
