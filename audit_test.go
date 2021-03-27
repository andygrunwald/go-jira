package jira

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
	"time"
)

func TestAuditService_Search(t *testing.T) {

	dateMocked, err := time.Parse(DateFormatJira, "2021-03-26T23:34:46.397-0700")
	if err != nil {
		t.Fatal(err)
	}

	testCases := []struct {
		name               string
		opts               *AuditSearchOptionsScheme
		offset, limit      int
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		params             map[string]string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
		errorMjs           string
	}{
		{
			name: "SearchAuditRecordsWhenTheParametersAreCorrect",
			opts: &AuditSearchOptionsScheme{
				Filter:     "created",
				From:       dateMocked, // Last 2 hours
				To:         time.Time{},
				ProjectIDs: []string{"10001", "10000"},
				UserIDs:    []string{"cjt9", "mz4w"},
			},
			offset:         0,
			limit:          50,
			mockFile:       "./mocks/get-audit-records.json",
			wantHTTPMethod: http.MethodGet,
			endpoint:       "/rest/api/2/auditing/record",
			params: map[string]string{
				"filter":     "created",
				"from":       "2021-03-26T23:34:46.397-0700",
				"limit":      "50",
				"offset":     "0",
				"projectIds": "10001,10000",
				"userIds":    "cjt9,mz4w",
			},
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name: "SearchAuditRecordsWhenTheContextIsNil",
			opts: &AuditSearchOptionsScheme{
				Filter:     "created",
				From:       dateMocked, // Last 2 hours
				To:         time.Time{},
				ProjectIDs: []string{"10001", "10000"},
				UserIDs:    []string{"cjt9", "mz4w"},
			},
			offset:         0,
			limit:          50,
			mockFile:       "./mocks/get-audit-records.json",
			wantHTTPMethod: http.MethodGet,
			endpoint:       "/rest/api/2/auditing/record",
			params: map[string]string{
				"filter":     "created",
				"from":       "2021-03-26T23:34:46.397-0700",
				"limit":      "50",
				"offset":     "0",
				"projectIds": "10001,10000",
				"userIds":    "cjt9,mz4w",
			},
			context:            nil,
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
			errorMjs:           "error!, please provide a valid ctx value",
		},

		{
			name: "SearchAuditRecordsWhenTheLimitIsNotSet",
			opts: &AuditSearchOptionsScheme{
				Filter:     "created",
				From:       dateMocked, // Last 2 hours
				To:         time.Time{},
				ProjectIDs: []string{"10001", "10000"},
				UserIDs:    []string{"cjt9", "mz4w"},
			},
			offset:         0,
			mockFile:       "./mocks/get-audit-records.json",
			wantHTTPMethod: http.MethodGet,
			endpoint:       "/rest/api/2/auditing/record",
			params: map[string]string{
				"filter":     "created",
				"from":       "2021-03-26T23:34:46.397-0700",
				"limit":      "0",
				"offset":     "0",
				"projectIds": "10001,10000",
				"userIds":    "cjt9,mz4w",
			},
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
			errorMjs:           "error!, the response struct won't return any records",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			//Init the HTTP mock server
			setup()
			defer teardown()

			expectedResponseAsBytes, err := ioutil.ReadFile(testCase.mockFile)
			if err != nil {
				t.Fatal(err)
			}

			//Create the custom HTTP handle
			testMux.HandleFunc(testCase.endpoint, func(w http.ResponseWriter, r *http.Request) {
				testMethod(t, r, testCase.wantHTTPMethod)
				testRequestURL(t, r, testCase.endpoint)
				testRequestParams(t, r, testCase.params)

				_, err := fmt.Fprint(w, string(expectedResponseAsBytes))
				if err != nil {
					w.WriteHeader(500)
				}

			})

			records, response, err := testClient.Audit.Search(testCase.context, testCase.opts, testCase.offset, testCase.limit)

			if testCase.wantErr {

				if !ErrorContains(err, testCase.errorMjs) {
					t.Fatalf("Error Wanted: (%v), Error Returned: (%v)", testCase.errorMjs, err)
				}

				t.Logf("Error Wanted: (%v), Error Returned: (%v)", testCase.errorMjs, err)
				return
			}

			if err != nil {
				if response != nil {
					t.Fatal("Response HTTP Response", response.StatusCode)
				}

				t.Fatal(err)
			}

			for index, record := range records.Records {
				t.Logf("----------------------------------")
				t.Logf("Record #%v", index+1)
				t.Logf("Record ID: %v", record.ID)
				t.Logf("Record Created by: %v", record.Created)
				t.Logf("Record Category: %v", record.Category)
				t.Logf("Record Source IP: %v", record.RemoteAddress)
				t.Logf("Record Event: %v", record.EventSource)
				t.Logf("----------------------------------")
			}

		})
	}

}

func ErrorContains(out error, want string) bool {
	if out == nil {
		return want == ""
	}
	if want == "" {
		return false
	}
	return strings.Contains(out.Error(), want)
}
