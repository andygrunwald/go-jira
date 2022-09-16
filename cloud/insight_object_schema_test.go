package cloud

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/mcl-de/go-jira/v2/cloud/model/apps/insight"
)

func TestInsightObjectSchemaService_List(t *testing.T) {
	setup()
	defer teardown()
	testMux.HandleFunc("/jsm/insight/workspace/g2778e1d-939d-581d-c8e2-9d5g59de456b/v1/objectschema/list", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testRequestURL(t, r, "/jsm/insight/workspace/g2778e1d-939d-581d-c8e2-9d5g59de456b/v1/objectschema/list")
		fmt.Fprint(w, `{"startAt":0,"maxResults":25,"total":5,"values":[{"workspaceId":"g2778e1d-939d-581d-c8e2-9d5g59de456b","globalId":"g2778e1d-939d-581d-c8e2-9d5g59de456b:13","id":"13","name":"Discovery import","objectSchemaKey":"NS","status":"Ok","description":"","created":"2021-02-22T02:31:31.748Z","updated":"2021-03-26T12:12:46.132Z","objectCount":231,"objectTypeCount":23,"idAsInt":13,"canManage":true},{"workspaceId":"g2778e1d-939d-581d-c8e2-9d5g59de456b","globalId":"g2778e1d-939d-581d-c8e2-9d5g59de456b:6","id":"6","name":"ITSM","objectSchemaKey":"ITSM","status":"Ok","created":"2021-02-16T18:04:31.284Z","updated":"2021-02-16T18:04:31.288Z","objectCount":95,"objectTypeCount":34,"idAsInt":6,"canManage":true},{"workspaceId":"f1668d0c-828c-470c-b7d1-8c4f48cd345a","globalId":"f1668d0c-828c-470c-b7d1-8c4f48cd345a:1","id":"1","name":"Human Resources","objectSchemaKey":"HR","status":"Ok","created":"2021-02-15T22:05:30.709Z","updated":"2021-03-18T13:49:57.909Z","objectCount":1023,"objectTypeCount":14,"idAsInt":1,"canManage":true},{"workspaceId":"f1668d0c-828c-470c-b7d1-8c4f48cd345a","globalId":"f1668d0c-828c-470c-b7d1-8c4f48cd345a:27","id":"27","name":"Services","objectSchemaKey":"SVC","status":"Ok","description":"Contains the 'Service' object type and services your site uses across projects.","created":"2021-03-19T04:52:40.418Z","updated":"2021-03-19T04:52:40.428Z","objectCount":37,"objectTypeCount":1,"idAsInt":27,"canManage":false},{"workspaceId":"f1668d0c-828c-470c-b7d1-8c4f48cd345a","globalId":"f1668d0c-828c-470c-b7d1-8c4f48cd345a:30","id":"30","name":"Word life","objectSchemaKey":"WL","status":"Ok","created":"2021-03-28T23:19:49.290Z","updated":"2021-03-28T23:19:49.299Z","objectCount":0,"objectTypeCount":0,"idAsInt":30,"canManage":true}],"isLast":true}`)
	})
	if objectSchemas, err := testClient.Insight.ObjectSchema.List(context.Background(), "g2778e1d-939d-581d-c8e2-9d5g59de456b"); err != nil {
		t.Errorf("Error given: %s", err)
	} else if objectSchemas == nil {
		t.Error("Expected objectSchemas. Object is nil")
	}
}

func TestInsightObjectSchemaService_Create(t *testing.T) {
	setup()
	defer teardown()
	testMux.HandleFunc("/jsm/insight/workspace/g2778e1d-939d-581d-c8e2-9d5g59de456b/v1/objectschema/create", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		testRequestURL(t, r, "/jsm/insight/workspace/g2778e1d-939d-581d-c8e2-9d5g59de456b/v1/objectschema/create")
		fmt.Fprint(w, `{"workspaceId":"g2778e1d-939d-581d-c8e2-9d5g59de456b","globalId":"g2778e1d-939d-581d-c8e2-9d5g59de456b:39","id":"39","name":"Computers","objectSchemaKey":"COMP","status":"Ok","description":"The IT department schema","created":"2021-04-20T16:21:18.908Z","updated":"2021-04-20T16:21:18.912Z","objectCount":0,"objectTypeCount":0,"idAsInt":39}`)
	})
	if objectSchema, err := testClient.Insight.ObjectSchema.Create(context.Background(), "g2778e1d-939d-581d-c8e2-9d5g59de456b", insight.CreateOrUpdateObjectSchema{}); err != nil {
		t.Errorf("Error given: %s", err)
	} else if objectSchema == nil {
		t.Error("Expected objectSchema. Object is nil")
	}
}

func TestInsightObjectSchemaService_Get(t *testing.T) {
	setup()
	defer teardown()
	testMux.HandleFunc("/jsm/insight/workspace/g2778e1d-939d-581d-c8e2-9d5g59de456b/v1/objectschema/39", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testRequestURL(t, r, "/jsm/insight/workspace/g2778e1d-939d-581d-c8e2-9d5g59de456b/v1/objectschema/39")
		fmt.Fprint(w, `{"workspaceId":"g2778e1d-939d-581d-c8e2-9d5g59de456b","globalId":"g2778e1d-939d-581d-c8e2-9d5g59de456b:39","id":"39","name":"Computers","objectSchemaKey":"COMP","status":"Ok","description":"The IT department schema","created":"2021-04-20T16:21:18.908Z","updated":"2021-04-20T16:21:18.912Z","objectCount":0,"objectTypeCount":0,"idAsInt":39}`)
	})
	if objectSchema, err := testClient.Insight.ObjectSchema.Get(context.Background(), "g2778e1d-939d-581d-c8e2-9d5g59de456b", "39"); err != nil {
		t.Errorf("Error given: %s", err)
	} else if objectSchema == nil {
		t.Error("Expected objectSchema. Object is nil")
	}
}

func TestInsightObjectSchemaService_Update(t *testing.T) {
	setup()
	defer teardown()
	testMux.HandleFunc("/jsm/insight/workspace/g2778e1d-939d-581d-c8e2-9d5g59de456b/v1/objectschema/39", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPut)
		testRequestURL(t, r, "/jsm/insight/workspace/g2778e1d-939d-581d-c8e2-9d5g59de456b/v1/objectschema/39")
		fmt.Fprint(w, `{"workspaceId":"g2778e1d-939d-581d-c8e2-9d5g59de456b","globalId":"g2778e1d-939d-581d-c8e2-9d5g59de456b:39","id":"39","name":"Computers","objectSchemaKey":"COMP","status":"Ok","description":"The IT department schema","created":"2021-04-20T16:21:18.908Z","updated":"2021-04-20T16:21:18.912Z","objectCount":0,"objectTypeCount":0,"idAsInt":39}`)
	})
	if objectSchema, err := testClient.Insight.ObjectSchema.Update(context.Background(), "g2778e1d-939d-581d-c8e2-9d5g59de456b", "39", insight.CreateOrUpdateObjectSchema{}); err != nil {
		t.Errorf("Error given: %s", err)
	} else if objectSchema == nil {
		t.Error("Expected objectSchema. Object is nil")
	}
}

func TestInsightObjectSchemaService_Delete(t *testing.T) {
	setup()
	defer teardown()
	testMux.HandleFunc("/jsm/insight/workspace/g2778e1d-939d-581d-c8e2-9d5g59de456b/v1/objectschema/39", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)
		testRequestURL(t, r, "/jsm/insight/workspace/g2778e1d-939d-581d-c8e2-9d5g59de456b/v1/objectschema/39")
		fmt.Fprint(w, `{"workspaceId":"g2778e1d-939d-581d-c8e2-9d5g59de456b","globalId":"g2778e1d-939d-581d-c8e2-9d5g59de456b:39","id":"39","name":"Computers","objectSchemaKey":"COMP","status":"Ok","description":"The IT department schema","created":"2021-04-20T16:21:18.908Z","updated":"2021-04-20T16:21:18.912Z","objectCount":0,"objectTypeCount":0,"idAsInt":39}`)
	})
	if objectSchema, err := testClient.Insight.ObjectSchema.Delete(context.Background(), "g2778e1d-939d-581d-c8e2-9d5g59de456b", "39"); err != nil {
		t.Errorf("Error given: %s", err)
	} else if objectSchema == nil {
		t.Error("Expected objectSchema. Object is nil")
	}
}

func TestInsightObjectSchemaService_GetAttributes(t *testing.T) {
	setup()
	defer teardown()
	testMux.HandleFunc("/jsm/insight/workspace/g2778e1d-939d-581d-c8e2-9d5g59de456b/v1/objectschema/13/attributes", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testRequestURL(t, r, "/jsm/insight/workspace/g2778e1d-939d-581d-c8e2-9d5g59de456b/v1/objectschema/13/attributes?extended=true&onlyValueEditable=true")
		fmt.Fprint(w, `[{"workspaceId":"f1668d0c-828c-470c-b7d1-8c4f48cd345a","globalId":"f1668d0c-828c-470c-b7d1-8c4f48cd345a:868","id":"868","name":"Tier","label":false,"defaultType":{"id":10,"name":"Select"},"editable":true,"system":false,"sortable":true,"summable":false,"indexed":true,"minimumCardinality":1,"maximumCardinality":1,"removable":true,"hidden":false,"includeChildObjectTypes":false,"uniqueAttribute":false,"options":"Tier 1,Tier 2,Tier 3,Tier 4","position":5},{"workspaceId":"f1668d0c-828c-470c-b7d1-8c4f48cd345a","globalId":"f1668d0c-828c-470c-b7d1-8c4f48cd345a:867","id":"867","name":"Description","label":false,"defaultType":{"id":0,"name":"Text"},"editable":true,"system":false,"sortable":true,"summable":false,"indexed":true,"minimumCardinality":0,"maximumCardinality":1,"removable":true,"hidden":false,"includeChildObjectTypes":false,"uniqueAttribute":false,"options":"","position":4},{"workspaceId":"f1668d0c-828c-470c-b7d1-8c4f48cd345a","globalId":"f1668d0c-828c-470c-b7d1-8c4f48cd345a:869","id":"869","name":"Service relationships","label":false,"referenceType":{"workspaceId":"f1668d0c-828c-470c-b7d1-8c4f48cd345a","globalId":"f1668d0c-828c-470c-b7d1-8c4f48cd345a:36","id":"36","name":"Depends on","color":"42526E","url16":"https://api.atlassian.com/jsm/insight/workspace/f1668d0c-828c-470c-b7d1-8c4f48cd345a/v1/config/referencetype/36/image.png?size=16","removable":false,"objectSchemaId":"27"},"referenceObjectTypeId":"122","referenceObjectType":{"workspaceId":"f1668d0c-828c-470c-b7d1-8c4f48cd345a","globalId":"f1668d0c-828c-470c-b7d1-8c4f48cd345a:122","id":"122","name":"Service","type":0,"description":"This object type contains the services your site uses across projects.","icon":{"id":"164","name":"Service","url16":"https://api.atlassian.com/jsm/insight/workspace/f1668d0c-828c-470c-b7d1-8c4f48cd345a/v1/objecttype/122/icon.png?size=16","url48":"https://api.atlassian.com/jsm/insight/workspace/f1668d0c-828c-470c-b7d1-8c4f48cd345a/v1/objecttype/122/icon.png?size=48"},"position":0,"created":"2021-03-19T04:52:40.472Z","updated":"2021-03-19T04:52:40.472Z","objectCount":0,"objectSchemaId":"27","inherited":false,"abstractObjectType":false,"parentObjectTypeInherited":false},"editable":true,"system":false,"sortable":true,"summable":false,"indexed":true,"minimumCardinality":0,"maximumCardinality":-1,"removable":true,"hidden":false,"includeChildObjectTypes":false,"uniqueAttribute":false,"options":"","position":6},{"workspaceId":"f1668d0c-828c-470c-b7d1-8c4f48cd345a","globalId":"f1668d0c-828c-470c-b7d1-8c4f48cd345a:866","id":"866","name":"Updated","label":false,"defaultType":{"id":6,"name":"DateTime"},"editable":false,"system":true,"sortable":true,"summable":false,"indexed":true,"minimumCardinality":0,"maximumCardinality":1,"removable":true,"hidden":false,"includeChildObjectTypes":false,"uniqueAttribute":false,"options":"","position":3},{"workspaceId":"f1668d0c-828c-470c-b7d1-8c4f48cd345a","globalId":"f1668d0c-828c-470c-b7d1-8c4f48cd345a:865","id":"865","name":"Created","label":false,"defaultType":{"id":6,"name":"DateTime"},"editable":false,"system":true,"sortable":true,"summable":false,"indexed":true,"minimumCardinality":1,"maximumCardinality":1,"removable":true,"hidden":false,"includeChildObjectTypes":false,"uniqueAttribute":false,"options":"","position":2},{"workspaceId":"f1668d0c-828c-470c-b7d1-8c4f48cd345a","globalId":"f1668d0c-828c-470c-b7d1-8c4f48cd345a:863","id":"863","name":"Key","label":false,"defaultType":{"id":0,"name":"Text"},"editable":false,"system":true,"sortable":true,"summable":false,"indexed":true,"minimumCardinality":1,"maximumCardinality":1,"removable":true,"hidden":false,"includeChildObjectTypes":false,"uniqueAttribute":false,"options":"","position":0},{"workspaceId":"f1668d0c-828c-470c-b7d1-8c4f48cd345a","globalId":"f1668d0c-828c-470c-b7d1-8c4f48cd345a:870","id":"870","name":"Service ID","label":false,"defaultType":{"id":0,"name":"Text"},"editable":true,"system":false,"sortable":true,"summable":false,"indexed":true,"minimumCardinality":1,"maximumCardinality":1,"removable":true,"hidden":false,"includeChildObjectTypes":false,"uniqueAttribute":true,"options":"","position":7},{"workspaceId":"f1668d0c-828c-470c-b7d1-8c4f48cd345a","globalId":"f1668d0c-828c-470c-b7d1-8c4f48cd345a:864","id":"864","name":"Name","label":true,"defaultType":{"id":0,"name":"Text"},"editable":true,"system":false,"sortable":true,"summable":false,"indexed":true,"minimumCardinality":1,"maximumCardinality":1,"removable":false,"hidden":false,"includeChildObjectTypes":false,"uniqueAttribute":false,"options":"","position":1},{"workspaceId":"f1668d0c-828c-470c-b7d1-8c4f48cd345a","globalId":"f1668d0c-828c-470c-b7d1-8c4f48cd345a:871","id":"871","name":"Revision","label":false,"defaultType":{"id":0,"name":"Text"},"editable":true,"system":false,"sortable":true,"summable":false,"indexed":true,"minimumCardinality":1,"maximumCardinality":1,"removable":true,"hidden":true,"includeChildObjectTypes":false,"uniqueAttribute":false,"options":"","position":8}]`)
	})
	if attributes, err := testClient.Insight.ObjectSchema.GetAttributes(context.Background(), "g2778e1d-939d-581d-c8e2-9d5g59de456b", "13", &insight.ObjectSchemaAttributeOptions{
		OnlyValueEditable: true,
		Extended:          true,
	}); err != nil {
		t.Errorf("Error given: %s", err)
	} else if attributes == nil {
		t.Error("Expected attributes. Object is nil")
	}
}

func TestInsightObjectSchemaService_GetObjectTypes(t *testing.T) {
	setup()
	defer teardown()
	testMux.HandleFunc("/jsm/insight/workspace/g2778e1d-939d-581d-c8e2-9d5g59de456b/v1/objectschema/13/objecttypes/flat", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testRequestURL(t, r, "/jsm/insight/workspace/g2778e1d-939d-581d-c8e2-9d5g59de456b/v1/objectschema/13/objecttypes/flat?IncludeObjectCounts=true")
		fmt.Fprint(w, `[{"workspaceId":"g2778e1d-939d-581d-c8e2-9d5g59de456b","globalId":"g2778e1d-939d-581d-c8e2-9d5g59de456b:19","id":"19","name":"Employee","icon":{"id":"131","name":"Users","url16":"https://api.atlassian.com/jsm/insight/workspace/g2778e1d-939d-581d-c8e2-9d5g59de456b/v1/objecttype/19/icon.png?size=16","url48":"https://api.atlassian.com/jsm/insight/workspace/g2778e1d-939d-581d-c8e2-9d5g59de456b/v1/objecttype/19/icon.png?size=48"},"position":0,"created":"2021-02-16T18:32:38.173Z","updated":"2021-02-16T19:37:07.179Z","objectCount":0,"objectSchemaId":"6","inherited":false,"abstractObjectType":false,"parentObjectTypeInherited":false},{"workspaceId":"g2778e1d-939d-581d-c8e2-9d5g59de456b","globalId":"g2778e1d-939d-581d-c8e2-9d5g59de456b:23","id":"23","name":"Office","description":"Lorem ipsum dolor sit amet, consectetur adipiscing elit. Proin nec ex.","icon":{"id":"13","name":"Building","url16":"https://api.atlassian.com/jsm/insight/workspace/g2778e1d-939d-581d-c8e2-9d5g59de456b/v1/objecttype/23/icon.png?size=16","url48":"https://api.atlassian.com/jsm/insight/workspace/g2778e1d-939d-581d-c8e2-9d5g59de456b/v1/objecttype/23/icon.png?size=48"},"position":2,"created":"2021-02-16T19:36:51.951Z","updated":"2021-04-16T15:17:03.384Z","objectCount":0,"objectSchemaId":"6","inherited":false,"abstractObjectType":false,"parentObjectTypeInherited":false},{"workspaceId":"g2778e1d-939d-581d-c8e2-9d5g59de456b","globalId":"g2778e1d-939d-581d-c8e2-9d5g59de456b:24","id":"24","name":"City","icon":{"id":"28","name":"Cottage","url16":"https://api.atlassian.com/jsm/insight/workspace/g2778e1d-939d-581d-c8e2-9d5g59de456b/v1/objecttype/24/icon.png?size=16","url48":"https://api.atlassian.com/jsm/insight/workspace/f1668d0c-828c-470c-b7d1-8c4f48cd345a/v1/objecttype/24/icon.png?size=48"},"position":3,"created":"2021-02-16T19:58:45.698Z","updated":"2021-04-16T15:17:03.393Z","objectCount":0,"objectSchemaId":"6","inherited":false,"abstractObjectType":false,"parentObjectTypeInherited":false}]`)
	})
	if objectTypes, err := testClient.Insight.ObjectSchema.GetObjectTypes(context.Background(), "g2778e1d-939d-581d-c8e2-9d5g59de456b", "13", &insight.ObjectSchemaObjectTypeOptions{
		IncludeObjectCounts: true,
	}); err != nil {
		t.Errorf("Error given: %s", err)
	} else if objectTypes == nil {
		t.Error("Expected objectTypes. Object is nil")
	}
}
