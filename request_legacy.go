// +build !go1.13

package jira

import (
	"context"
	"io"
	"net/http"
)

func newRequestWithContext(ctx context.Context, method, url string, body io.Reader) (*http.Request, error) {
	r, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	return r.WithContext(ctx), nil
}
