package jira

import (
	"bytes"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/jarcoal/httpmock"
	"io/ioutil"
	"net/http"
	"reflect"
	"testing"
)

func TestAuthenticationService_AcquireSessionCookie_Failure(t *testing.T) {
	setup()
	defer teardown()
	testMux.HandleFunc("/rest/auth/1/session", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		testRequestURL(t, r, "/rest/auth/1/session")
		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			t.Errorf("Error in read body: %s", err)
		}
		if bytes.Index(b, []byte(`"username":"foo"`)) < 0 {
			t.Error("No username found")
		}
		if bytes.Index(b, []byte(`"password":"bar"`)) < 0 {
			t.Error("No password found")
		}

		// Emulate error
		w.WriteHeader(http.StatusInternalServerError)
	})

	res, err := testClient.Authentication.AcquireSessionCookie("foo", "bar")
	if err == nil {
		t.Errorf("Expected error, but no error given")
	}
	if res == true {
		t.Error("Expected error, but result was true")
	}

	if testClient.Authentication.Authenticated() != false {
		t.Error("Expected false, but result was true")
	}
}

func TestAuthenticationService_AcquireSessionCookie_Success(t *testing.T) {
	setup()
	defer teardown()
	testMux.HandleFunc("/rest/auth/1/session", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		testRequestURL(t, r, "/rest/auth/1/session")
		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			t.Errorf("Error in read body: %s", err)
		}
		if bytes.Index(b, []byte(`"username":"foo"`)) < 0 {
			t.Error("No username found")
		}
		if bytes.Index(b, []byte(`"password":"bar"`)) < 0 {
			t.Error("No password found")
		}

		fmt.Fprint(w, `{"session":{"name":"JSESSIONID","value":"12345678901234567890"},"loginInfo":{"failedLoginCount":10,"loginCount":127,"lastFailedLoginTime":"2016-03-16T04:22:35.386+0000","previousLoginTime":"2016-03-16T04:22:35.386+0000"}}`)
	})

	res, err := testClient.Authentication.AcquireSessionCookie("foo", "bar")
	if err != nil {
		t.Errorf("No error expected. Got %s", err)
	}
	if res == false {
		t.Error("Expected result was true. Got false")
	}

	if testClient.Authentication.Authenticated() != true {
		t.Error("Expected true, but result was false")
	}

	if testClient.Authentication.authType != authTypeSession {
		t.Errorf("Expected authType %d. Got %d", authTypeSession, testClient.Authentication.authType)
	}
}

func TestAuthenticationService_SetBasicAuth(t *testing.T) {
	setup()
	defer teardown()

	testClient.Authentication.SetBasicAuth("test-user", "test-password")

	if testClient.Authentication.username != "test-user" {
		t.Errorf("Expected username test-user. Got %s", testClient.Authentication.username)
	}

	if testClient.Authentication.password != "test-password" {
		t.Errorf("Expected password test-password. Got %s", testClient.Authentication.password)
	}

	if testClient.Authentication.authType != authTypeBasic {
		t.Errorf("Expected authType %d. Got %d", authTypeBasic, testClient.Authentication.authType)
	}
}

func TestAuthenticationService_SetOauth2(t *testing.T) {
	setup()
	defer teardown()

	testClient.Authentication.SetOauth2("token")

	if testClient.Authentication.accessToken != "token" {
		t.Errorf("Expected accessToken='token'. Got '%s'", testClient.Authentication.accessToken)
	}

	if testClient.Authentication.authType != authTypeOauth2 {
		t.Errorf("Expected authType %d. Got %d", authTypeOauth2, testClient.Authentication.authType)
	}
}

func TestAuthenticationService_Authenticated(t *testing.T) {
	// Skip setup() because we don't want a fully setup client
	testClient = new(Client)

	// Test before we've attempted to authenticate
	if testClient.Authentication.Authenticated() != false {
		t.Error("Expected false, but result was true")
	}
}

func TestAuthenticationService_Authenticated_WithBasicAuth(t *testing.T) {
	setup()
	defer teardown()

	testClient.Authentication.SetBasicAuth("test-user", "test-password")

	// Test before we've attempted to authenticate
	if testClient.Authentication.Authenticated() != true {
		t.Error("Expected true, but result was false")
	}
}

func TestAuthenticationService_Authenticated_WithBasicAuthButNoUsername(t *testing.T) {
	setup()
	defer teardown()

	testClient.Authentication.SetBasicAuth("", "test-password")

	// Test before we've attempted to authenticate
	if testClient.Authentication.Authenticated() != false {
		t.Error("Expected false, but result was true")
	}
}

func TestAuthenticationService_Authenticated_WithOAuth2(t *testing.T) {
	setup()
	defer teardown()

	testClient.Authentication.SetOauth2("token")

	// Test before we've attempted to authenticate
	if testClient.Authentication.Authenticated() != true {
		t.Error("Expected true, but result was false")
	}
}

func TestAithenticationService_GetUserInfo_AccessForbidden_Fail(t *testing.T) {
	setup()
	defer teardown()
	testMux.HandleFunc("/rest/auth/1/session", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			testMethod(t, r, "POST")
			testRequestURL(t, r, "/rest/auth/1/session")
			b, err := ioutil.ReadAll(r.Body)
			if err != nil {
				t.Errorf("Error in read body: %s", err)
			}
			if bytes.Index(b, []byte(`"username":"foo"`)) < 0 {
				t.Error("No username found")
			}
			if bytes.Index(b, []byte(`"password":"bar"`)) < 0 {
				t.Error("No password found")
			}

			fmt.Fprint(w, `{"session":{"name":"JSESSIONID","value":"12345678901234567890"},"loginInfo":{"failedLoginCount":10,"loginCount":127,"lastFailedLoginTime":"2016-03-16T04:22:35.386+0000","previousLoginTime":"2016-03-16T04:22:35.386+0000"}}`)
		}

		if r.Method == "GET" {
			testMethod(t, r, "GET")
			testRequestURL(t, r, "/rest/auth/1/session")

			w.WriteHeader(http.StatusForbidden)
		}
	})

	testClient.Authentication.AcquireSessionCookie("foo", "bar")

	_, err := testClient.Authentication.GetCurrentUser()
	if err == nil {
		t.Errorf("Non nil error expect, received nil")
	}
}

func TestAuthenticationService_GetUserInfo_NonOkStatusCode_Fail(t *testing.T) {
	setup()
	defer teardown()

	testMux.HandleFunc("/rest/auth/1/session", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			testMethod(t, r, "POST")
			testRequestURL(t, r, "/rest/auth/1/session")
			b, err := ioutil.ReadAll(r.Body)
			if err != nil {
				t.Errorf("Error in read body: %s", err)
			}
			if bytes.Index(b, []byte(`"username":"foo"`)) < 0 {
				t.Error("No username found")
			}
			if bytes.Index(b, []byte(`"password":"bar"`)) < 0 {
				t.Error("No password found")
			}

			fmt.Fprint(w, `{"session":{"name":"JSESSIONID","value":"12345678901234567890"},"loginInfo":{"failedLoginCount":10,"loginCount":127,"lastFailedLoginTime":"2016-03-16T04:22:35.386+0000","previousLoginTime":"2016-03-16T04:22:35.386+0000"}}`)
		}

		if r.Method == "GET" {
			testMethod(t, r, "GET")
			testRequestURL(t, r, "/rest/auth/1/session")
			//any status but 200
			w.WriteHeader(240)
		}
	})

	testClient.Authentication.AcquireSessionCookie("foo", "bar")

	_, err := testClient.Authentication.GetCurrentUser()
	if err == nil {
		t.Errorf("Non nil error expect, received nil")
	}
}

func TestAuthenticationService_GetUserInfo_FailWithoutLogin(t *testing.T) {
	// no setup() required here
	testClient = new(Client)

	_, err := testClient.Authentication.GetCurrentUser()
	if err == nil {
		t.Errorf("Expected error, but got %s", err)
	}
}

func TestAuthenticationService_GetUserInfo_Success(t *testing.T) {
	setup()
	defer teardown()

	testUserInfo := new(Session)
	testUserInfo.Name = "foo"
	testUserInfo.Self = "https://my.jira.com/rest/api/latest/user?username=foo"
	testUserInfo.LoginInfo.FailedLoginCount = 12
	testUserInfo.LoginInfo.LastFailedLoginTime = "2016-09-06T16:41:23.949+0200"
	testUserInfo.LoginInfo.LoginCount = 357
	testUserInfo.LoginInfo.PreviousLoginTime = "2016-09-07T11:36:23.476+0200"

	testMux.HandleFunc("/rest/auth/1/session", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			testMethod(t, r, "POST")
			testRequestURL(t, r, "/rest/auth/1/session")
			b, err := ioutil.ReadAll(r.Body)
			if err != nil {
				t.Errorf("Error in read body: %s", err)
			}
			if bytes.Index(b, []byte(`"username":"foo"`)) < 0 {
				t.Error("No username found")
			}
			if bytes.Index(b, []byte(`"password":"bar"`)) < 0 {
				t.Error("No password found")
			}

			fmt.Fprint(w, `{"session":{"name":"JSESSIONID","value":"12345678901234567890"},"loginInfo":{"failedLoginCount":10,"loginCount":127,"lastFailedLoginTime":"2016-03-16T04:22:35.386+0000","previousLoginTime":"2016-03-16T04:22:35.386+0000"}}`)
		}

		if r.Method == "GET" {
			testMethod(t, r, "GET")
			testRequestURL(t, r, "/rest/auth/1/session")
			fmt.Fprint(w, `{"self":"https://my.jira.com/rest/api/latest/user?username=foo","name":"foo","loginInfo":{"failedLoginCount":12,"loginCount":357,"lastFailedLoginTime":"2016-09-06T16:41:23.949+0200","previousLoginTime":"2016-09-07T11:36:23.476+0200"}}`)
		}
	})

	testClient.Authentication.AcquireSessionCookie("foo", "bar")

	userinfo, err := testClient.Authentication.GetCurrentUser()
	if err != nil {
		t.Errorf("Nil error expect, received %s", err)
	}
	equal := reflect.DeepEqual(*testUserInfo, *userinfo)

	if !equal {
		t.Error("The user information doesn't match")
	}
}

func TestAuthenticationService_Logout_Success(t *testing.T) {
	setup()
	defer teardown()

	testMux.HandleFunc("/rest/auth/1/session", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			testMethod(t, r, "POST")
			testRequestURL(t, r, "/rest/auth/1/session")
			b, err := ioutil.ReadAll(r.Body)
			if err != nil {
				t.Errorf("Error in read body: %s", err)
			}
			if bytes.Index(b, []byte(`"username":"foo"`)) < 0 {
				t.Error("No username found")
			}
			if bytes.Index(b, []byte(`"password":"bar"`)) < 0 {
				t.Error("No password found")
			}

			fmt.Fprint(w, `{"session":{"name":"JSESSIONID","value":"12345678901234567890"},"loginInfo":{"failedLoginCount":10,"loginCount":127,"lastFailedLoginTime":"2016-03-16T04:22:35.386+0000","previousLoginTime":"2016-03-16T04:22:35.386+0000"}}`)
		}

		if r.Method == "DELETE" {
			// return 204
			w.WriteHeader(http.StatusNoContent)
		}
	})

	testClient.Authentication.AcquireSessionCookie("foo", "bar")

	err := testClient.Authentication.Logout()
	if err != nil {
		t.Errorf("Expected nil error, got %s", err)
	}
}

func TestAuthenticationService_Logout_FailWithoutLogin(t *testing.T) {
	setup()
	defer teardown()

	testMux.HandleFunc("/rest/auth/1/session", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "DELETE" {
			// 401
			w.WriteHeader(http.StatusUnauthorized)
		}
	})
	err := testClient.Authentication.Logout()
	if err == nil {
		t.Error("Expected not nil, got nil")
	}
}

func TestAuthenticationService_GetOauth2AccessToken(t *testing.T) {
	setup()
	defer teardown()
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("POST", jiraOauth2URL, httpmock.NewStringResponder(200, `{"access_token":"token","token_type":"Bearer","expires_in":900}`))
	a, err := testClient.Authentication.GetOauth2AccessToken("clientId", "userKey", "READ", []byte("secret"))
	if err != nil {
		t.Errorf("Expected no error getting oauth2 access token, but got : %s", err)
	}
	if a != nil {
		if a.AccessToken != "token" {
			t.Errorf("Unexpected access token : %s", a.AccessToken)
		}
		if a.TokenType != "Bearer" {
			t.Errorf("Unexpected access token type : %s", a.TokenType)
		}
		if a.ExpiresIn != 900 {
			t.Errorf("Unexpected access token expiration : %d", a.ExpiresIn)
		}
	}
}

func TestAuthenticationService_createJWT(t *testing.T) {
	testURL := "https://example.com"
	secret := []byte("secret")
	jwtStr, err := createJWT("oauthId", testURL, testURL, "user", secret)
	if err != nil {
		t.Errorf("Unexpected error creating JWT: %s", err)
	}

	type JiraClaims struct {
		InstanceURL string `json:"tnt"`
		jwt.StandardClaims
	}

	token, err := jwt.ParseWithClaims(jwtStr, &JiraClaims{}, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if err != nil {
		t.Errorf("Unexpected error parsing JWT: %s", err)
	}

	if claims, ok := token.Claims.(*JiraClaims); !ok {
		t.Errorf("Unexpected error processing custom Jira Claims: %s", err)
	} else {
		if claims.InstanceURL != testURL {
			t.Errorf("Exepcted %s in claims 'tnt' but got: %s", testURL, testURL)
		}
		if claims.Audience != testURL {
			t.Errorf("Exepcted %s in claims 'aud' but got: %s", testURL, testURL)
		}
		if claims.Issuer != "urn:atlassian:connect:clientid:oauthId" {
			t.Errorf("Exepcted urn:atlassian:connect:clientid:oauthId in claims 'iss' but got: %s", claims.Issuer)
		}
		if claims.Subject != "urn:atlassian:connect:userkey:user" {
			t.Errorf("Exepcted urn:atlassian:connect:userkey:user in claims 'sub' but got: %s", claims.Subject)
		}
	}
}
