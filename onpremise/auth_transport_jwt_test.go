package onpremise

import (
	"context"
	"net/http"
	"strings"
	"testing"
)

func TestJWTAuthTransport_HeaderContainsJWT(t *testing.T) {
	setup()
	defer teardown()

	sharedSecret := []byte("ssshh,it's a secret")
	issuer := "add-on.key"

	jwtTransport := &JWTAuthTransport{
		Secret: sharedSecret,
		Issuer: issuer,
	}

	testMux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// look for the presence of the JWT in the header
		val := r.Header.Get("Authorization")
		if !strings.Contains(val, "JWT ") {
			t.Errorf("request does not contain JWT in the Auth header")
		}
	})

	jwtClient, _ := NewClient(testServer.URL, jwtTransport.Client())
	jwtClient.Issue.Get(context.Background(), "TEST-1", nil)
}
