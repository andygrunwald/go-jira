package jira

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
)

func TestOrganizationService_GetAllOrganizationsWithContext(t *testing.T) {
	setup()
	defer teardown()
	testMux.HandleFunc("/rest/servicedeskapi/organization", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testRequestURL(t, r, "/rest/servicedeskapi/organization")

		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `{ "_expands": [], "size": 1, "start": 1, "limit": 1, "isLastPage": false, "_links": { "base": "https://your-domain.atlassian.net/rest/servicedeskapi", "context": "context", "next": "https://your-domain.atlassian.net/rest/servicedeskapi/organization?start=2&limit=1", "prev": "https://your-domain.atlassian.net/rest/servicedeskapi/organization?start=0&limit=1" }, "values": [ { "id": "1", "name": "Charlie Cakes Franchises", "_links": { "self": "https://your-domain.atlassian.net/rest/servicedeskapi/organization/1" } } ] }`)
	})

	result, _, err := testClient.Organization.GetAllOrganizations(0, 50, "")

	if result == nil {
		t.Error("Expected Organizations. Result is nil")
	} else if result.Size != 1 {
		t.Errorf("Expected size to be 1, but got %d", result.Size)
	}

	if err != nil {
		t.Errorf("Error given: %s", err)
	}
}

func TestOrganizationService_CreateOrganization(t *testing.T) {
	setup()
	defer teardown()
	testMux.HandleFunc("/rest/servicedeskapi/organization", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		testRequestURL(t, r, "/rest/servicedeskapi/organization")

		o := new(OrganizationCreationDTO)
		json.NewDecoder(r.Body).Decode(&o)
		w.WriteHeader(http.StatusCreated)
		fmt.Fprintf(w, `{ "id": "1", "name": "%s", "_links": { "self": "https://your-domain.atlassian.net/rest/servicedeskapi/organization/1" } }`, o.Name)
	})

	name := "MyOrg"
	o, _, err := testClient.Organization.CreateOrganization(name)

	if o == nil {
		t.Error("Expected Organization. Result is nil")
	} else if o.Name != name {
		t.Errorf("Expected name to be %s, but got %s", name, o.Name)
	}

	if err != nil {
		t.Errorf("Error given: %s", err)
	}
}

func TestOrganizationService_GetOrganization(t *testing.T) {
	setup()
	defer teardown()
	testMux.HandleFunc("/rest/servicedeskapi/organization/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testRequestURL(t, r, "/rest/servicedeskapi/organization/1")

		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, `{ "id": "1", "name": "name", "_links": { "self": "https://your-domain.atlassian.net/rest/servicedeskapi/organization/1" } }`)
	})

	id := 1
	o, _, err := testClient.Organization.GetOrganization(id)

	if err != nil {
		t.Errorf("Error given: %s", err)
	}

	if o == nil {
		t.Error("Expected Organization. Result is nil")
	} else if o.Name != "name" {
		t.Errorf("Expected name to be name, but got %s", o.Name)
	}
}

func TestOrganizationService_DeleteOrganization(t *testing.T) {
	setup()
	defer teardown()
	testMux.HandleFunc("/rest/servicedeskapi/organization/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
		testRequestURL(t, r, "/rest/servicedeskapi/organization/1")

		w.WriteHeader(http.StatusNoContent)
	})

	_, err := testClient.Organization.DeleteOrganization(1)

	if err != nil {
		t.Errorf("Error given: %s", err)
	}
}

func TestOrganizationService_GetPropertiesKeys(t *testing.T) {
	setup()
	defer teardown()
	testMux.HandleFunc("/rest/servicedeskapi/organization/1/property", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testRequestURL(t, r, "/rest/servicedeskapi/organization/1/property")

		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, `{
			"keys": [
			  {
				"self": "/rest/servicedeskapi/organization/1/property/propertyKey",
				"key": "organization.attributes"
			  }
			]
		  }`)
	})

	pk, _, err := testClient.Organization.GetPropertiesKeys(1)

	if err != nil {
		t.Errorf("Error given: %s", err)
	}

	if pk == nil {
		t.Error("Expected Keys. Result is nil")
	} else if pk.Keys[0].Key != "organization.attributes" {
		t.Errorf("Expected name to be organization.attributes, but got %s", pk.Keys[0].Key)
	}
}

func TestOrganizationService_GetProperty(t *testing.T) {
	setup()
	defer teardown()
	testMux.HandleFunc("/rest/servicedeskapi/organization/1/property/organization.attributes", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testRequestURL(t, r, "/rest/servicedeskapi/organization/1/property/organization.attributes")

		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, `{
			"key": "organization.attributes",
			"value": {
			  "phone": "0800-1233456789",
			  "mail": "charlie@example.com"
			}
		  }`)
	})

	key := "organization.attributes"
	ep, _, err := testClient.Organization.GetProperty(1, key)

	if err != nil {
		t.Errorf("Error given: %s", err)
	}

	if ep == nil {
		t.Error("Expected Entity. Result is nil")
	} else if ep.Key != key {
		t.Errorf("Expected name to be %s, but got %s", key, ep.Key)
	}
}

func TestOrganizationService_SetProperty(t *testing.T) {
	setup()
	defer teardown()
	testMux.HandleFunc("/rest/servicedeskapi/organization/1/property/organization.attributes", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PUT")
		testRequestURL(t, r, "/rest/servicedeskapi/organization/1/property/organization.attributes")

		w.WriteHeader(http.StatusOK)
	})

	key := "organization.attributes"
	_, err := testClient.Organization.SetProperty(1, key)

	if err != nil {
		t.Errorf("Error given: %s", err)
	}
}

func TestOrganizationService_DeleteProperty(t *testing.T) {
	setup()
	defer teardown()
	testMux.HandleFunc("/rest/servicedeskapi/organization/1/property/organization.attributes", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
		testRequestURL(t, r, "/rest/servicedeskapi/organization/1/property/organization.attributes")

		w.WriteHeader(http.StatusOK)
	})

	key := "organization.attributes"
	_, err := testClient.Organization.DeleteProperty(1, key)

	if err != nil {
		t.Errorf("Error given: %s", err)
	}
}

func TestOrganizationService_GetUsers(t *testing.T) {
	setup()
	defer teardown()
	testMux.HandleFunc("/rest/servicedeskapi/organization/1/user", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testRequestURL(t, r, "/rest/servicedeskapi/organization/1/user")

		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, `{
			"_expands": [],
			"size": 1,
			"start": 1,
			"limit": 1,
			"isLastPage": false,
			"_links": {
			  "base": "https://your-domain.atlassian.net/rest/servicedeskapi",
			  "context": "context",
			  "next": "https://your-domain.atlassian.net/rest/servicedeskapi/organization/1/user?start=2&limit=1",
			  "prev": "https://your-domain.atlassian.net/rest/servicedeskapi/organization/1/user?start=0&limit=1"
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
					"48x48": "https://avatar-cdn.atlassian.com/9bc3b5bcb0db050c6d7660b28a5b86c9?s=48",
					"24x24": "https://avatar-cdn.atlassian.com/9bc3b5bcb0db050c6d7660b28a5b86c9?s=24",
					"16x16": "https://avatar-cdn.atlassian.com/9bc3b5bcb0db050c6d7660b28a5b86c9?s=16",
					"32x32": "https://avatar-cdn.atlassian.com/9bc3b5bcb0db050c6d7660b28a5b86c9?s=32"
				  },
				  "self": "https://your-domain.atlassian.net/rest/api/2/user?username=qm:a713c8ea-1075-4e30-9d96-891a7d181739:5ad6d3581db05e2a66fa80b"
				}
			  },
			  {
				"accountId": "qm:a713c8ea-1075-4e30-9d96-891a7d181739:5ad6d3a01db05e2a66fa80bd",
				"name": "qm:a713c8ea-1075-4e30-9d96-891a7d181739:5ad6d3a01db05e2a66fa80bd",
				"key": "qm:a713c8ea-1075-4e30-9d96-891a7d181739:5ad6d3a01db05e2a66fa80bd",
				"emailAddress": "bob@example.com",
				"displayName": "Bob D. Builder",
				"active": true,
				"timeZone": "Australia/Sydney",
				"_links": {
				  "jiraRest": "https://your-domain.atlassian.net/rest/api/2/user?username=qm:a713c8ea-1075-4e30-9d96-891a7d181739:5ad6d3a01db05e2a66fa80bd",
				  "avatarUrls": {
					"48x48": "https://avatar-cdn.atlassian.com/9bc3b5bcb0db050c6d7660b28a5b86c9?s=48",
					"24x24": "https://avatar-cdn.atlassian.com/9bc3b5bcb0db050c6d7660b28a5b86c9?s=24",
					"16x16": "https://avatar-cdn.atlassian.com/9bc3b5bcb0db050c6d7660b28a5b86c9?s=16",
					"32x32": "https://avatar-cdn.atlassian.com/9bc3b5bcb0db050c6d7660b28a5b86c9?s=32"
				  },
				  "self": "https://your-domain.atlassian.net/rest/api/2/user?username=qm:a713c8ea-1075-4e30-9d96-891a7d181739:5ad6d3a01db05e2a66fa80bd"
				}
			  }
			]
		  }`)
	})

	users, _, err := testClient.Organization.GetUsers(1, 0, 50)

	if err != nil {
		t.Errorf("Error given: %s", err)
	}

	if users == nil {
		t.Error("Expected Organizations. Result is nil")
	} else if users.Size != 1 {
		t.Errorf("Expected size to be 1, but got %d", users.Size)
	}

	if err != nil {
		t.Errorf("Error given: %s", err)
	}
}

func TestOrganizationService_AddUsers(t *testing.T) {
	setup()
	defer teardown()
	testMux.HandleFunc("/rest/servicedeskapi/organization/1/user", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		testRequestURL(t, r, "/rest/servicedeskapi/organization/1/user")

		w.WriteHeader(http.StatusNoContent)
	})

	users := OrganizationUsersDTO{
		AccountIds: []string{
			"qm:a713c8ea-1075-4e30-9d96-891a7d181739:5ad6d3581db05e2a66fa80b",
			"qm:a713c8ea-1075-4e30-9d96-891a7d181739:5ad6d3a01db05e2a66fa80bd",
		},
	}
	_, err := testClient.Organization.AddUsers(1, users)

	if err != nil {
		t.Errorf("Error given: %s", err)
	}
}

func TestOrganizationService_RemoveUsers(t *testing.T) {
	setup()
	defer teardown()
	testMux.HandleFunc("/rest/servicedeskapi/organization/1/user", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
		testRequestURL(t, r, "/rest/servicedeskapi/organization/1/user")

		w.WriteHeader(http.StatusNoContent)
	})

	users := OrganizationUsersDTO{
		AccountIds: []string{
			"qm:a713c8ea-1075-4e30-9d96-891a7d181739:5ad6d3581db05e2a66fa80b",
			"qm:a713c8ea-1075-4e30-9d96-891a7d181739:5ad6d3a01db05e2a66fa80bd",
		},
	}
	_, err := testClient.Organization.RemoveUsers(1, users)

	if err != nil {
		t.Errorf("Error given: %s", err)
	}
}
