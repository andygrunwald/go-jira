package jira

import (
	"net/http"
	"net/url"
	"testing"
)

func TestErrorResponse_Empty(t *testing.T) {
	u, _ := url.Parse("https://issues.apache.org/jira/browse/MESOS-5040")
	r := &http.Response{
		Request: &http.Request{
			Method: "POST",
			URL:    u,
		},
		StatusCode: 200,
	}

	mockData := []struct {
		Response ErrorResponse
		Expected string
	}{
		{
			Response: ErrorResponse{},
			Expected: "[] map[]",
		},
		{
			Response: ErrorResponse{
				ErrorMessages: []string{"foo", "bar"},
			},
			Expected: "[foo bar] map[]",
		},
		{
			Response: ErrorResponse{
				Errors: map[string]string{"Foo": "Bar"},
			},
			Expected: "[] map[Foo:Bar]",
		},
		{
			Response: ErrorResponse{
				ErrorMessages: []string{"foo", "bar"},
				Errors:        map[string]string{"Foo": "Bar"},
			},
			Expected: "[foo bar] map[Foo:Bar]",
		},
		{
			Response: ErrorResponse{
				Response: r,
			},
			Expected: "POST https://issues.apache.org/jira/browse/MESOS-5040: 200 [] map[]",
		},
		{
			Response: ErrorResponse{
				Response:      r,
				ErrorMessages: []string{"foo", "bar"},
			},
			Expected: "POST https://issues.apache.org/jira/browse/MESOS-5040: 200 [foo bar] map[]",
		},
		{
			Response: ErrorResponse{
				Response: r,
				Errors:   map[string]string{"Foo": "Bar"},
			},
			Expected: "POST https://issues.apache.org/jira/browse/MESOS-5040: 200 [] map[Foo:Bar]",
		},
		{
			Response: ErrorResponse{
				Response:      r,
				ErrorMessages: []string{"foo", "bar"},
				Errors:        map[string]string{"Foo": "Bar"},
			},
			Expected: "POST https://issues.apache.org/jira/browse/MESOS-5040: 200 [foo bar] map[Foo:Bar]",
		},
	}

	for _, data := range mockData {
		got := data.Response.Error()
		if got != data.Expected {
			t.Errorf("Response is different as expected. Expected \"%s\". Got \"%s\"", data.Expected, got)
		}
	}
}
