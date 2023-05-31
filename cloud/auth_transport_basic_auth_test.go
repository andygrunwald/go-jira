package cloud

import (
	"context"
	"net/http"
	"testing"
)

func TestBasicAuthTransport(t *testing.T) {
	setup()
	defer teardown()

	username, apiToken := "username", "api_token"

	testMux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		u, p, ok := r.BasicAuth()
		if !ok {
			t.Errorf("request does not contain basic auth credentials")
		}
		if u != username {
			t.Errorf("request contained basic auth username %q, want %q", u, username)
		}
		if p != apiToken {
			t.Errorf("request contained basic auth password %q, want %q", p, apiToken)
		}
	})

	tp := &BasicAuthTransport{
		Username: username,
		APIToken: apiToken,
	}

	basicAuthClient, _ := NewClient(testServer.URL, tp.Client())
	req, _ := basicAuthClient.NewRequest(context.Background(), http.MethodGet, ".", nil)
	basicAuthClient.Do(req, nil)
}

func TestBasicAuthTransport_transport(t *testing.T) {
	// default transport
	tp := &BasicAuthTransport{}
	if tp.transport() != http.DefaultTransport {
		t.Errorf("Expected http.DefaultTransport to be used.")
	}

	// custom transport
	tp = &BasicAuthTransport{
		Transport: &http.Transport{},
	}
	if tp.transport() == http.DefaultTransport {
		t.Errorf("Expected custom transport to be used.")
	}
}
