package onpremise

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

// Test that the cookie in the transport is the cookie returned in the header
func TestCookieAuthTransport_SessionObject_Exists(t *testing.T) {
	setup()
	defer teardown()

	testCookie := &http.Cookie{Name: "test", Value: "test"}

	testMux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		cookies := r.Cookies()

		if len(cookies) < 1 {
			t.Errorf("No cookies set")
		}

		if cookies[0].Name != testCookie.Name {
			t.Errorf("Cookie names don't match, expected %v, got %v", testCookie.Name, cookies[0].Name)
		}

		if cookies[0].Value != testCookie.Value {
			t.Errorf("Cookie values don't match, expected %v, got %v", testCookie.Value, cookies[0].Value)
		}
	})

	tp := &CookieAuthTransport{
		Username:      "username",
		Password:      "password",
		AuthURL:       "https://some.jira.com/rest/auth/1/session",
		SessionObject: []*http.Cookie{testCookie},
	}

	basicAuthClient, _ := NewClient(testServer.URL, tp.Client())
	req, _ := basicAuthClient.NewRequest(context.Background(), http.MethodGet, ".", nil)
	basicAuthClient.Do(req, nil)
}

// Test that an empty cookie in the transport is not returned in the header
func TestCookieAuthTransport_SessionObject_ExistsWithEmptyCookie(t *testing.T) {
	setup()
	defer teardown()

	emptyCookie := &http.Cookie{Name: "empty_cookie", Value: ""}
	testCookie := &http.Cookie{Name: "test", Value: "test"}

	testMux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		cookies := r.Cookies()

		if len(cookies) > 1 {
			t.Errorf("The empty cookie should not have been added")
		}

		if cookies[0].Name != testCookie.Name {
			t.Errorf("Cookie names don't match, expected %v, got %v", testCookie.Name, cookies[0].Name)
		}

		if cookies[0].Value != testCookie.Value {
			t.Errorf("Cookie values don't match, expected %v, got %v", testCookie.Value, cookies[0].Value)
		}
	})

	tp := &CookieAuthTransport{
		Username:      "username",
		Password:      "password",
		AuthURL:       "https://some.jira.com/rest/auth/1/session",
		SessionObject: []*http.Cookie{emptyCookie, testCookie},
	}

	basicAuthClient, _ := NewClient(testServer.URL, tp.Client())
	req, _ := basicAuthClient.NewRequest(context.Background(), http.MethodGet, ".", nil)
	basicAuthClient.Do(req, nil)
}

// Test that if no cookie is in the transport, it checks for a cookie
func TestCookieAuthTransport_SessionObject_DoesNotExist(t *testing.T) {
	setup()
	defer teardown()

	testCookie := &http.Cookie{Name: "does_not_exist", Value: "does_not_exist"}

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		http.SetCookie(w, testCookie)
		w.Write([]byte(`OK`))
	}))
	defer ts.Close()

	testMux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		cookies := r.Cookies()

		if len(cookies) < 1 {
			t.Errorf("No cookies set")
		}

		if cookies[0].Name != testCookie.Name {
			t.Errorf("Cookie names don't match, expected %v, got %v", testCookie.Name, cookies[0].Name)
		}

		if cookies[0].Value != testCookie.Value {
			t.Errorf("Cookie values don't match, expected %v, got %v", testCookie.Value, cookies[0].Value)
		}
	})

	tp := &CookieAuthTransport{
		Username: "username",
		Password: "password",
		AuthURL:  ts.URL,
	}

	basicAuthClient, _ := NewClient(testServer.URL, tp.Client())
	req, _ := basicAuthClient.NewRequest(context.Background(), http.MethodGet, ".", nil)
	basicAuthClient.Do(req, nil)
}
