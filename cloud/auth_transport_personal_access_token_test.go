package cloud

import (
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

	client, _ := NewClient(patTransport.Client(), testServer.URL)
	client.User.GetSelf()

}
