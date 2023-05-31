package onpremise

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"reflect"
	"testing"
)

func TestAuthenticationService_AcquireSessionCookie_Failure(t *testing.T) {
	setup()
	defer teardown()
	testMux.HandleFunc("/rest/auth/1/session", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		testRequestURL(t, r, "/rest/auth/1/session")
		b, err := io.ReadAll(r.Body)
		if err != nil {
			t.Errorf("Error in read body: %s", err)
		}
		if !bytes.Contains(b, []byte(`"username":"foo"`)) {
			t.Error("No username found")
		}
		if !bytes.Contains(b, []byte(`"password":"bar"`)) {
			t.Error("No password found")
		}

		// Emulate error
		w.WriteHeader(http.StatusInternalServerError)
	})

	res, err := testClient.Authentication.AcquireSessionCookie(context.Background(), "foo", "bar")
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
		testMethod(t, r, http.MethodPost)
		testRequestURL(t, r, "/rest/auth/1/session")
		b, err := io.ReadAll(r.Body)
		if err != nil {
			t.Errorf("Error in read body: %s", err)
		}
		if !bytes.Contains(b, []byte(`"username":"foo"`)) {
			t.Error("No username found")
		}
		if !bytes.Contains(b, []byte(`"password":"bar"`)) {
			t.Error("No password found")
		}

		fmt.Fprint(w, `{"session":{"name":"JSESSIONID","value":"12345678901234567890"},"loginInfo":{"failedLoginCount":10,"loginCount":127,"lastFailedLoginTime":"2016-03-16T04:22:35.386+0000","previousLoginTime":"2016-03-16T04:22:35.386+0000"}}`)
	})

	res, err := testClient.Authentication.AcquireSessionCookie(context.Background(), "foo", "bar")
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

func TestAuthenticationService_GetUserInfo_AccessForbidden_Fail(t *testing.T) {
	setup()
	defer teardown()
	testMux.HandleFunc("/rest/auth/1/session", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			testMethod(t, r, http.MethodPost)
			testRequestURL(t, r, "/rest/auth/1/session")
			b, err := io.ReadAll(r.Body)
			if err != nil {
				t.Errorf("Error in read body: %s", err)
			}
			if !bytes.Contains(b, []byte(`"username":"foo"`)) {
				t.Error("No username found")
			}
			if !bytes.Contains(b, []byte(`"password":"bar"`)) {
				t.Error("No password found")
			}

			fmt.Fprint(w, `{"session":{"name":"JSESSIONID","value":"12345678901234567890"},"loginInfo":{"failedLoginCount":10,"loginCount":127,"lastFailedLoginTime":"2016-03-16T04:22:35.386+0000","previousLoginTime":"2016-03-16T04:22:35.386+0000"}}`)
		}

		if r.Method == http.MethodGet {
			testMethod(t, r, http.MethodGet)
			testRequestURL(t, r, "/rest/auth/1/session")

			w.WriteHeader(http.StatusForbidden)
		}
	})

	testClient.Authentication.AcquireSessionCookie(context.Background(), "foo", "bar")

	_, err := testClient.Authentication.GetCurrentUser(context.Background())
	if err == nil {
		t.Errorf("Non nil error expect, received nil")
	}
}

func TestAuthenticationService_GetUserInfo_NonOkStatusCode_Fail(t *testing.T) {
	setup()
	defer teardown()

	testMux.HandleFunc("/rest/auth/1/session", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			testMethod(t, r, http.MethodPost)
			testRequestURL(t, r, "/rest/auth/1/session")
			b, err := io.ReadAll(r.Body)
			if err != nil {
				t.Errorf("Error in read body: %s", err)
			}
			if !bytes.Contains(b, []byte(`"username":"foo"`)) {
				t.Error("No username found")
			}
			if !bytes.Contains(b, []byte(`"password":"bar"`)) {
				t.Error("No password found")
			}

			fmt.Fprint(w, `{"session":{"name":"JSESSIONID","value":"12345678901234567890"},"loginInfo":{"failedLoginCount":10,"loginCount":127,"lastFailedLoginTime":"2016-03-16T04:22:35.386+0000","previousLoginTime":"2016-03-16T04:22:35.386+0000"}}`)
		}

		if r.Method == http.MethodGet {
			testMethod(t, r, http.MethodGet)
			testRequestURL(t, r, "/rest/auth/1/session")
			//any status but 200
			w.WriteHeader(240)
		}
	})

	testClient.Authentication.AcquireSessionCookie(context.Background(), "foo", "bar")

	_, err := testClient.Authentication.GetCurrentUser(context.Background())
	if err == nil {
		t.Errorf("Non nil error expect, received nil")
	}
}

func TestAuthenticationService_GetUserInfo_FailWithoutLogin(t *testing.T) {
	// no setup() required here
	testClient = new(Client)

	_, err := testClient.Authentication.GetCurrentUser(context.Background())
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
		if r.Method == http.MethodPost {
			testMethod(t, r, http.MethodPost)
			testRequestURL(t, r, "/rest/auth/1/session")
			b, err := io.ReadAll(r.Body)
			if err != nil {
				t.Errorf("Error in read body: %s", err)
			}
			if !bytes.Contains(b, []byte(`"username":"foo"`)) {
				t.Error("No username found")
			}
			if !bytes.Contains(b, []byte(`"password":"bar"`)) {
				t.Error("No password found")
			}

			fmt.Fprint(w, `{"session":{"name":"JSESSIONID","value":"12345678901234567890"},"loginInfo":{"failedLoginCount":10,"loginCount":127,"lastFailedLoginTime":"2016-03-16T04:22:35.386+0000","previousLoginTime":"2016-03-16T04:22:35.386+0000"}}`)
		}

		if r.Method == http.MethodGet {
			testMethod(t, r, http.MethodGet)
			testRequestURL(t, r, "/rest/auth/1/session")
			fmt.Fprint(w, `{"self":"https://my.jira.com/rest/api/latest/user?username=foo","name":"foo","loginInfo":{"failedLoginCount":12,"loginCount":357,"lastFailedLoginTime":"2016-09-06T16:41:23.949+0200","previousLoginTime":"2016-09-07T11:36:23.476+0200"}}`)
		}
	})

	testClient.Authentication.AcquireSessionCookie(context.Background(), "foo", "bar")

	userinfo, err := testClient.Authentication.GetCurrentUser(context.Background())
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
		if r.Method == http.MethodPost {
			testMethod(t, r, http.MethodPost)
			testRequestURL(t, r, "/rest/auth/1/session")
			b, err := io.ReadAll(r.Body)
			if err != nil {
				t.Errorf("Error in read body: %s", err)
			}
			if !bytes.Contains(b, []byte(`"username":"foo"`)) {
				t.Error("No username found")
			}
			if !bytes.Contains(b, []byte(`"password":"bar"`)) {
				t.Error("No password found")
			}

			fmt.Fprint(w, `{"session":{"name":"JSESSIONID","value":"12345678901234567890"},"loginInfo":{"failedLoginCount":10,"loginCount":127,"lastFailedLoginTime":"2016-03-16T04:22:35.386+0000","previousLoginTime":"2016-03-16T04:22:35.386+0000"}}`)
		}

		if r.Method == http.MethodDelete {
			// return 204
			w.WriteHeader(http.StatusNoContent)
		}
	})

	testClient.Authentication.AcquireSessionCookie(context.Background(), "foo", "bar")

	err := testClient.Authentication.Logout(context.Background())
	if err != nil {
		t.Errorf("Expected nil error, got %s", err)
	}
}

func TestAuthenticationService_Logout_FailWithoutLogin(t *testing.T) {
	setup()
	defer teardown()

	testMux.HandleFunc("/rest/auth/1/session", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodDelete {
			// 401
			w.WriteHeader(http.StatusUnauthorized)
		}
	})
	err := testClient.Authentication.Logout(context.Background())
	if err == nil {
		t.Error("Expected not nil, got nil")
	}
}
