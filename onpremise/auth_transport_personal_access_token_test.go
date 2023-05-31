package onpremise

import (
	"context"
	"net/http"
	"testing"
)

func TestPATAuthTransport_HeaderContainsAuth(t *testing.T) {
	setup()
	defer teardown()

	token := "shhh, it's a token"

	patTransport := &PATAuthTransport{
		Token: token,
	}

	testMux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		val := r.Header.Get("Authorization")
		expected := "Bearer " + token
		if val != expected {
			t.Errorf("request does not contain bearer token in the Authorization header.")
		}
	})

	client, _ := NewClient(testServer.URL, patTransport.Client())
	client.User.GetSelf(context.Background())

}
