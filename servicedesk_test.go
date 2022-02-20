package jira

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"sort"
	"strconv"
	"testing"
)

func TestServiceDeskService_GetOrganizations(t *testing.T) {
	setup()
	defer teardown()
	testMux.HandleFunc("/rest/servicedeskapi/servicedesk/10001/organization", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testRequestURL(t, r, "/rest/servicedeskapi/servicedesk/10001/organization")

		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `{
			"_expands": [],
			"size": 3,
			"start": 3,
			"limit": 3,
			"isLastPage": false,
			"_links": {
			  "base": "https://your-domain.atlassian.net/rest/servicedeskapi",
			  "context": "context",
			  "next": "https://your-domain.atlassian.net/rest/servicedeskapi/servicedesk/10001/organization?start=6&limit=3",
			  "prev": "https://your-domain.atlassian.net/rest/servicedeskapi/servicedesk/10001/organization?start=0&limit=3"
			},
			"values": [
			  {
				"id": "1",
				"name": "Charlie Cakes Franchises",
				"_links": {
				  "self": "https://your-domain.atlassian.net/rest/servicedeskapi/organization/1"
				}
			  },
			  {
				"id": "2",
				"name": "Atlas Coffee Co",
				"_links": {
				  "self": "https://your-domain.atlassian.net/rest/servicedeskapi/organization/2"
				}
			  },
			  {
				"id": "3",
				"name": "The Adjustment Bureau",
				"_links": {
				  "self": "https://your-domain.atlassian.net/rest/servicedeskapi/organization/3"
				}
			  }
			]
		  }`)
	})

	orgs, _, err := testClient.ServiceDesk.GetOrganizations(10001, 3, 3, "")

	if orgs == nil {
		t.Error("Expected Organizations. Result is nil")
	} else if orgs.Size != 3 {
		t.Errorf("Expected size to be 3, but got %d", orgs.Size)
	}

	if err != nil {
		t.Errorf("Error given: %s", err)
	}
}

func TestServiceDeskService_AddOrganizations(t *testing.T) {
	setup()
	defer teardown()
	testMux.HandleFunc("/rest/servicedeskapi/servicedesk/10001/organization", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		testRequestURL(t, r, "/rest/servicedeskapi/servicedesk/10001/organization")

		w.WriteHeader(http.StatusNoContent)
	})

	_, err := testClient.ServiceDesk.AddOrganization(10001, 1)

	if err != nil {
		t.Errorf("Error given: %s", err)
	}
}

func TestServiceDeskService_RemoveOrganizations(t *testing.T) {
	setup()
	defer teardown()
	testMux.HandleFunc("/rest/servicedeskapi/servicedesk/10001/organization", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
		testRequestURL(t, r, "/rest/servicedeskapi/servicedesk/10001/organization")

		w.WriteHeader(http.StatusNoContent)
	})

	_, err := testClient.ServiceDesk.RemoveOrganization(10001, 1)

	if err != nil {
		t.Errorf("Error given: %s", err)
	}
}

func TestServiceDeskServiceStringServiceDeskID_GetOrganizations(t *testing.T) {
	setup()
	defer teardown()
	testMux.HandleFunc("/rest/servicedeskapi/servicedesk/TEST/organization", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testRequestURL(t, r, "/rest/servicedeskapi/servicedesk/TEST/organization")

		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `{
			"_expands": [],
			"size": 3,
			"start": 3,
			"limit": 3,
			"isLastPage": false,
			"_links": {
			  "base": "https://your-domain.atlassian.net/rest/servicedeskapi",
			  "context": "context",
			  "next": "https://your-domain.atlassian.net/rest/servicedeskapi/servicedesk/TEST/organization?start=6&limit=3",
			  "prev": "https://your-domain.atlassian.net/rest/servicedeskapi/servicedesk/TEST/organization?start=0&limit=3"
			},
			"values": [
			  {
				"id": "1",
				"name": "Charlie Cakes Franchises",
				"_links": {
				  "self": "https://your-domain.atlassian.net/rest/servicedeskapi/organization/1"
				}
			  },
			  {
				"id": "2",
				"name": "Atlas Coffee Co",
				"_links": {
				  "self": "https://your-domain.atlassian.net/rest/servicedeskapi/organization/2"
				}
			  },
			  {
				"id": "3",
				"name": "The Adjustment Bureau",
				"_links": {
				  "self": "https://your-domain.atlassian.net/rest/servicedeskapi/organization/3"
				}
			  }
			]
		  }`)
	})

	orgs, _, err := testClient.ServiceDesk.GetOrganizations("TEST", 3, 3, "")

	if orgs == nil {
		t.Error("Expected Organizations. Result is nil")
	} else if orgs.Size != 3 {
		t.Errorf("Expected size to be 3, but got %d", orgs.Size)
	}

	if err != nil {
		t.Errorf("Error given: %s", err)
	}
}

func TestServiceDeskServiceStringServiceDeskID_AddOrganizations(t *testing.T) {
	setup()
	defer teardown()
	testMux.HandleFunc("/rest/servicedeskapi/servicedesk/TEST/organization", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		testRequestURL(t, r, "/rest/servicedeskapi/servicedesk/TEST/organization")

		w.WriteHeader(http.StatusNoContent)
	})

	_, err := testClient.ServiceDesk.AddOrganization("TEST", 1)

	if err != nil {
		t.Errorf("Error given: %s", err)
	}
}

func TestServiceDeskServiceStringServiceDeskID_RemoveOrganizations(t *testing.T) {
	setup()
	defer teardown()
	testMux.HandleFunc("/rest/servicedeskapi/servicedesk/TEST/organization", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
		testRequestURL(t, r, "/rest/servicedeskapi/servicedesk/TEST/organization")

		w.WriteHeader(http.StatusNoContent)
	})

	_, err := testClient.ServiceDesk.RemoveOrganization("TEST", 1)

	if err != nil {
		t.Errorf("Error given: %s", err)
	}
}

func TestServiceDeskService_AddCustomers(t *testing.T) {
	tests := []struct {
		name          string
		serviceDeskID interface{}
	}{
		{
			name:          "string service desk id",
			serviceDeskID: "10000",
		},
		{
			name:          "int service desk id",
			serviceDeskID: 10000,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			setup()
			defer teardown()

			var (
				wantAccountIDs = []string{
					"qm:a713c8ea-1075-4e30-9d96-891a7d181739:5ad6d3a01db05e2a66fa80bd",
					"qm:a713c8ea-1075-4e30-9d96-891a7d181739:5ad6d3581db05e2a66fa80b",
				}
				gotAccountIDs []string
			)

			testMux.HandleFunc(fmt.Sprintf("/rest/servicedeskapi/servicedesk/%v/customer", test.serviceDeskID), func(w http.ResponseWriter, r *http.Request) {
				testMethod(t, r, "POST")
				testRequestURL(t, r, fmt.Sprintf("/rest/servicedeskapi/servicedesk/%v/customer", test.serviceDeskID))

				var payload struct {
					AccountIDs []string `json:"accountIds"`
				}

				if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
					t.Fatal(err)
				}

				gotAccountIDs = append(gotAccountIDs, payload.AccountIDs...)

				w.WriteHeader(http.StatusNoContent)
			})

			_, err := testClient.ServiceDesk.AddCustomers(test.serviceDeskID, wantAccountIDs...)

			if err != nil {
				t.Errorf("Error given: %s", err)
			}

			if want, got := len(wantAccountIDs), len(gotAccountIDs); want != got {
				t.Fatalf("want account id length: %d, got %d", want, got)
			}

			sort.Strings(wantAccountIDs)
			sort.Strings(gotAccountIDs)

			if !reflect.DeepEqual(wantAccountIDs, gotAccountIDs) {
				t.Fatalf("want account ids: %v, got %v", wantAccountIDs, gotAccountIDs)
			}
		})
	}
}

func TestServiceDeskService_RemoveCustomers(t *testing.T) {
	tests := []struct {
		name          string
		serviceDeskID interface{}
	}{
		{
			name:          "string service desk id",
			serviceDeskID: "10000",
		},
		{
			name:          "int service desk id",
			serviceDeskID: 10000,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			setup()
			defer teardown()

			var (
				wantAccountIDs = []string{
					"qm:a713c8ea-1075-4e30-9d96-891a7d181739:5ad6d3a01db05e2a66fa80bd",
					"qm:a713c8ea-1075-4e30-9d96-891a7d181739:5ad6d3581db05e2a66fa80b",
				}
				gotAccountIDs []string
			)

			testMux.HandleFunc(fmt.Sprintf("/rest/servicedeskapi/servicedesk/%v/customer", test.serviceDeskID), func(w http.ResponseWriter, r *http.Request) {
				testMethod(t, r, "DELETE")
				testRequestURL(t, r, fmt.Sprintf("/rest/servicedeskapi/servicedesk/%v/customer", test.serviceDeskID))

				var payload struct {
					AccountIDs []string `json:"accountIds"`
				}

				if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
					t.Fatal(err)
				}

				gotAccountIDs = append(gotAccountIDs, payload.AccountIDs...)

				w.WriteHeader(http.StatusNoContent)
			})

			_, err := testClient.ServiceDesk.RemoveCustomers(test.serviceDeskID, wantAccountIDs...)

			if err != nil {
				t.Errorf("Error given: %s", err)
			}

			if want, got := len(wantAccountIDs), len(gotAccountIDs); want != got {
				t.Fatalf("want account id length: %d, got %d", want, got)
			}

			sort.Strings(wantAccountIDs)
			sort.Strings(gotAccountIDs)

			if !reflect.DeepEqual(wantAccountIDs, gotAccountIDs) {
				t.Fatalf("want account ids: %v, got %v", wantAccountIDs, gotAccountIDs)
			}
		})
	}
}

func TestServiceDeskService_ListCustomers(t *testing.T) {
	tests := []struct {
		name          string
		serviceDeskID interface{}
	}{
		{
			name:          "string service desk id",
			serviceDeskID: "10000",
		},
		{
			name:          "int service desk id",
			serviceDeskID: 10000,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			setup()
			defer teardown()

			var (
				email       = "fred@example.com"
				wantOptions = &CustomerListOptions{
					Query: email,
					Start: 1,
					Limit: 10,
				}

				gotOptions = new(CustomerListOptions)
			)

			testMux.HandleFunc(fmt.Sprintf("/rest/servicedeskapi/servicedesk/%v/customer", test.serviceDeskID), func(w http.ResponseWriter, r *http.Request) {
				testMethod(t, r, "GET")
				testRequestURL(t, r, fmt.Sprintf("/rest/servicedeskapi/servicedesk/%v/customer", test.serviceDeskID))

				qs := r.URL.Query()
				gotOptions.Query = qs.Get("query")
				if start := qs.Get("start"); start != "" {
					gotOptions.Start, _ = strconv.Atoi(start)
				}
				if limit := qs.Get("limit"); limit != "" {
					gotOptions.Limit, _ = strconv.Atoi(limit)
				}

				w.Write([]byte(`{
				  "_expands": [],
				  "size": 1,
				  "start": 1,
				  "limit": 1,
				  "isLastPage": false,
				  "_links": {
					"base": "https://your-domain.atlassian.net/rest/servicedeskapi",
					"context": "context",
					"next": "https://your-domain.atlassian.net/rest/servicedeskapi/servicedesk/1/customer?start=2&limit=1",
					"prev": "https://your-domain.atlassian.net/rest/servicedeskapi/servicedesk/1/customer?start=0&limit=1"
				  },
				  "values": [
					{
					  "accountId": "qm:a713c8ea-1075-4e30-9d96-891a7d181739:5ad6d3581db05e2a66fa80b",
					  "name": "qm:a713c8ea-1075-4e30-9d96-891a7d181739:5ad6d3581db05e2a66fa80b",
					  "key": "qm:a713c8ea-1075-4e30-9d96-891a7d181739:5ad6d3581db05e2a66fa80b",
					  "emailAddress": "fred@example.com",
					  "displayName": "Fred F. User",
					  "active": true,
					  "timeZone": "Australia/Sydney",
					  "_links": {
						"jiraRest": "https://your-domain.atlassian.net/rest/api/2/user?username=qm:a713c8ea-1075-4e30-9d96-891a7d181739:5ad6d3581db05e2a66fa80b",
						"avatarUrls": {
						  "48x48": "https://avatar-cdn.atlassian.com/9bc3b5bcb0db050c6d7660b28a5b86c9?s=48&d=https%3A%2F%2Fsecure.gravatar.com%2Favatar%2F9bc3b5bcb0db050c6d7660b28a5b86c9%3Fd%3Dmm%26s%3D48%26noRedirect%3Dtrue",
						  "24x24": "https://avatar-cdn.atlassian.com/9bc3b5bcb0db050c6d7660b28a5b86c9?s=24&d=https%3A%2F%2Fsecure.gravatar.com%2Favatar%2F9bc3b5bcb0db050c6d7660b28a5b86c9%3Fd%3Dmm%26s%3D24%26noRedirect%3Dtrue",
						  "16x16": "https://avatar-cdn.atlassian.com/9bc3b5bcb0db050c6d7660b28a5b86c9?s=16&d=https%3A%2F%2Fsecure.gravatar.com%2Favatar%2F9bc3b5bcb0db050c6d7660b28a5b86c9%3Fd%3Dmm%26s%3D16%26noRedirect%3Dtrue",
						  "32x32": "https://avatar-cdn.atlassian.com/9bc3b5bcb0db050c6d7660b28a5b86c9?s=32&d=https%3A%2F%2Fsecure.gravatar.com%2Favatar%2F9bc3b5bcb0db050c6d7660b28a5b86c9%3Fd%3Dmm%26s%3D32%26noRedirect%3Dtrue"
						},
						"self": "https://your-domain.atlassian.net/rest/api/2/user?username=qm:a713c8ea-1075-4e30-9d96-891a7d181739:5ad6d3581db05e2a66fa80b"
					  }
					}
				  ]
				}`))
			})

			customerList, _, err := testClient.ServiceDesk.ListCustomers(test.serviceDeskID, wantOptions)
			if err != nil {
				t.Fatal(err)
			}

			if !reflect.DeepEqual(wantOptions, gotOptions) {
				t.Fatalf("want options: %#v, got %#v", wantOptions, gotOptions)
			}

			if want, got := 1, len(customerList.Values); want != got {
				t.Fatalf("want customer count: %d, got %d", want, got)
			}

			if want, got := email, customerList.Values[0].EmailAddress; want != got {
				t.Fatalf("want customer email: %q, got %q", want, got)
			}
		})
	}
}
