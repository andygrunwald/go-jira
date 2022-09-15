package cloud

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/mcl-de/go-jira/v2/cloud/model/apps/insight"
)

func TestInsightObjectTypeService_Get(t *testing.T) {
	setup()
	defer teardown()
	testMux.HandleFunc("/jsm/insight/workspace/g2778e1d-939d-581d-c8e2-9d5g59de456b/v1/objecttype/23", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testRequestURL(t, r, "/jsm/insight/workspace/g2778e1d-939d-581d-c8e2-9d5g59de456b/v1/objecttype/23")
		fmt.Fprint(w, `{"workspaceId":"g2778e1d-939d-581d-c8e2-9d5g59de456b","globalId":"g2778e1d-939d-581d-c8e2-9d5g59de456b:23","id":"23","name":"Office","description":"Lorem ipsum dolor sit amet, consectetur adipiscing elit. Proin nec ex.","icon":{"id":"13","name":"Building","url16":"https://api.atlassian.com/jsm/insight/workspace/g2778e1d-939d-581d-c8e2-9d5g59de456b/v1/objecttype/23/icon.png?size=16","url48":"https://api.atlassian.com/jsm/insight/workspace/g2778e1d-939d-581d-c8e2-9d5g59de456b/v1/objecttype/23/icon.png?size=48"},"position":2,"created":"2021-02-16T19:36:51.951Z","updated":"2021-04-16T15:17:03.384Z","objectCount":4,"objectSchemaId":"6","inherited":false,"abstractObjectType":false,"parentObjectTypeInherited":false}`)
	})
	if objectType, err := testClient.Insight.ObjectType.Get(context.Background(), "g2778e1d-939d-581d-c8e2-9d5g59de456b", "23"); err != nil {
		t.Errorf("Error given: %s", err)
	} else if objectType == nil {
		t.Error("Expected objectType. Object is nil")
	}
}

func TestInsightObjectTypeService_Update(t *testing.T) {
	setup()
	defer teardown()
	testMux.HandleFunc("/jsm/insight/workspace/g2778e1d-939d-581d-c8e2-9d5g59de456b/v1/objecttype/23", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPut)
		testRequestURL(t, r, "/jsm/insight/workspace/g2778e1d-939d-581d-c8e2-9d5g59de456b/v1/objecttype/23")

		fmt.Fprint(w, `{"workspaceId":"g2778e1d-939d-581d-c8e2-9d5g59de456b","globalId":"g2778e1d-939d-581d-c8e2-9d5g59de456b:23","id":"23","name":"Office","description":"Lorem ipsum dolor sit amet, consectetur adipiscing elit. Proin nec ex.","icon":{"id":"13","name":"Building","url16":"https://api.atlassian.com/jsm/insight/workspace/g2778e1d-939d-581d-c8e2-9d5g59de456b/v1/objecttype/23/icon.png?size=16","url48":"https://api.atlassian.com/jsm/insight/workspace/g2778e1d-939d-581d-c8e2-9d5g59de456b/v1/objecttype/23/icon.png?size=48"},"position":2,"created":"2021-02-16T19:36:51.951Z","updated":"2021-04-16T15:17:03.384Z","objectCount":4,"objectSchemaId":"6","inherited":false,"abstractObjectType":false,"parentObjectTypeInherited":false}`)
	})

	if object, err := testClient.Insight.ObjectType.Update(context.Background(), "g2778e1d-939d-581d-c8e2-9d5g59de456b", "23", insight.PutObjectType{}); err != nil {
		t.Errorf("Error given: %s", err)
	} else if object == nil {
		t.Error("Expected object. Object is nil")
	}
}

func TestInsightObjectTypeService_Delete(t *testing.T) {
	setup()
	defer teardown()
	testMux.HandleFunc("/jsm/insight/workspace/g2778e1d-939d-581d-c8e2-9d5g59de456b/v1/objecttype/23", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)
		testRequestURL(t, r, "/jsm/insight/workspace/g2778e1d-939d-581d-c8e2-9d5g59de456b/v1/objecttype/23")
	})

	if err := testClient.Insight.ObjectType.Delete(context.Background(), "g2778e1d-939d-581d-c8e2-9d5g59de456b", "23"); err != nil {
		t.Errorf("Error given: %s", err)
	}
}

func TestInsightObjectTypeService_GetAttributes(t *testing.T) {
	setup()
	defer teardown()
	testMux.HandleFunc("/jsm/insight/workspace/g2778e1d-939d-581d-c8e2-9d5g59de456b/v1/objecttype/23/attributes", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testRequestURL(t, r, "/jsm/insight/workspace/g2778e1d-939d-581d-c8e2-9d5g59de456b/v1/objecttype/23/attributes?onlyValueEditable=true&orderByName=true&orderByRequired=true")
		fmt.Fprint(w, `[{"workspaceId":"g2778e1d-939d-581d-c8e2-9d5g59de456b","globalId":"g2778e1d-939d-581d-c8e2-9d5g59de456b:637","id":"637","objectTypeAttribute":{"workspaceId":"g2778e1d-939d-581d-c8e2-9d5g59de456b","globalId":"g2778e1d-939d-581d-c8e2-9d5g59de456b:134","id":"134","name":"Key","label":false,"type":0,"defaultType":{"id":0,"name":"Text"},"editable":false,"system":true,"sortable":true,"summable":false,"indexed":true,"minimumCardinality":1,"maximumCardinality":1,"removable":false,"hidden":false,"includeChildObjectTypes":false,"uniqueAttribute":false,"options":"","position":0},"objectTypeAttributeId":"134","objectAttributeValues":[{"value":"ITSM-88","displayValue":"ITSM-88","searchValue":"ITSM-88","referencedType":false}],"objectId":"88"},{"workspaceId":"g2778e1d-939d-581d-c8e2-9d5g59de456b","globalId":"g2778e1d-939d-581d-c8e2-9d5g59de456b:640","id":"640","objectTypeAttribute":{"workspaceId":"g2778e1d-939d-581d-c8e2-9d5g59de456b","globalId":"g2778e1d-939d-581d-c8e2-9d5g59de456b:135","id":"135","name":"Name","label":true,"type":0,"description":"The name of the object","defaultType":{"id":0,"name":"Text"},"editable":true,"system":false,"sortable":true,"summable":false,"indexed":true,"minimumCardinality":1,"maximumCardinality":1,"suffix":"","removable":false,"hidden":false,"includeChildObjectTypes":false,"uniqueAttribute":true,"regexValidation":"","iql":"","options":"","position":1},"objectTypeAttributeId":"135","objectAttributeValues":[{"value":"SYD-1","displayValue":"SYD-1","searchValue":"SYD-1","referencedType":false}],"objectId":"88"},{"workspaceId":"g2778e1d-939d-581d-c8e2-9d5g59de456b","globalId":"g2778e1d-939d-581d-c8e2-9d5g59de456b:638","id":"638","objectTypeAttribute":{"workspaceId":"g2778e1d-939d-581d-c8e2-9d5g59de456b","globalId":"g2778e1d-939d-581d-c8e2-9d5g59de456b:136","id":"136","name":"Created","label":false,"type":0,"defaultType":{"id":6,"name":"DateTime"},"editable":false,"system":true,"sortable":true,"summable":false,"indexed":true,"minimumCardinality":1,"maximumCardinality":1,"removable":false,"hidden":false,"includeChildObjectTypes":false,"uniqueAttribute":false,"options":"","position":2},"objectTypeAttributeId":"136","objectAttributeValues":[{"value":"2021-02-16T20:04:41.527Z","displayValue":"16/Feb/21 8:04 PM","searchValue":"2021-02-16T20:04:41.527Z","referencedType":false}],"objectId":"88"},{"workspaceId":"f1668d0c-828c-470c-b7d1-8c4f48cd345a","globalId":"f1668d0c-828c-470c-b7d1-8c4f48cd345a:639","id":"639","objectTypeAttribute":{"workspaceId":"f1668d0c-828c-470c-b7d1-8c4f48cd345a","globalId":"f1668d0c-828c-470c-b7d1-8c4f48cd345a:137","id":"137","name":"Updated","label":false,"type":0,"defaultType":{"id":6,"name":"DateTime"},"editable":false,"system":true,"sortable":true,"summable":false,"indexed":true,"minimumCardinality":1,"maximumCardinality":1,"removable":false,"hidden":false,"includeChildObjectTypes":false,"uniqueAttribute":false,"options":"","position":3},"objectTypeAttributeId":"137","objectAttributeValues":[{"value":"2021-04-20T14:55:02.816Z","displayValue":"20/Apr/21 2:55 PM","searchValue":"2021-04-20T14:55:02.816Z","referencedType":false}],"objectId":"88"},{"workspaceId":"f1668d0c-828c-470c-b7d1-8c4f48cd345a","globalId":"f1668d0c-828c-470c-b7d1-8c4f48cd345a:641","id":"641","objectTypeAttribute":{"workspaceId":"f1668d0c-828c-470c-b7d1-8c4f48cd345a","globalId":"f1668d0c-828c-470c-b7d1-8c4f48cd345a:144","id":"144","name":"City","label":false,"type":1,"referenceType":{"workspaceId":"f1668d0c-828c-470c-b7d1-8c4f48cd345a","globalId":"f1668d0c-828c-470c-b7d1-8c4f48cd345a:4","id":"4","name":"Reference","description":"Reference","color":"49a6ed","url16":"https://api.atlassian.com/jsm/insight/workspace/f1668d0c-828c-470c-b7d1-8c4f48cd345a/v1/config/referencetype/4/image.png?size=16","removable":false},"referenceObjectTypeId":"24","referenceObjectType":{"workspaceId":"f1668d0c-828c-470c-b7d1-8c4f48cd345a","globalId":"f1668d0c-828c-470c-b7d1-8c4f48cd345a:24","id":"24","name":"City","type":0,"icon":{"id":"28","name":"Cottage","url16":"https://api.atlassian.com/jsm/insight/workspace/f1668d0c-828c-470c-b7d1-8c4f48cd345a/v1/objecttype/24/icon.png?size=16","url48":"https://api.atlassian.com/jsm/insight/workspace/f1668d0c-828c-470c-b7d1-8c4f48cd345a/v1/objecttype/24/icon.png?size=48"},"position":3,"created":"2021-02-16T19:58:45.698Z","updated":"2021-04-16T15:17:03.393Z","objectCount":0,"objectSchemaId":"6","inherited":false,"abstractObjectType":false,"parentObjectTypeInherited":false},"editable":true,"system":false,"sortable":true,"summable":false,"indexed":true,"minimumCardinality":0,"maximumCardinality":1,"removable":true,"hidden":false,"includeChildObjectTypes":false,"uniqueAttribute":false,"options":"","position":4},"objectTypeAttributeId":"144","objectAttributeValues":[{"referencedObject":{"workspaceId":"f1668d0c-828c-470c-b7d1-8c4f48cd345a","globalId":"f1668d0c-828c-470c-b7d1-8c4f48cd345a:87","id":"87","label":"Sydney","objectKey":"ITSM-87","avatar":{"workspaceId":"f1668d0c-828c-470c-b7d1-8c4f48cd345a","url16":"https://api.atlassian.com/jsm/insight/workspace/f1668d0c-828c-470c-b7d1-8c4f48cd345a/v1/objecttype/24/icon.png?size=16","url48":"https://api.atlassian.com/jsm/insight/workspace/f1668d0c-828c-470c-b7d1-8c4f48cd345a/v1/objecttype/24/icon.png?size=48","url72":"https://api.atlassian.com/jsm/insight/workspace/f1668d0c-828c-470c-b7d1-8c4f48cd345a/v1/objecttype/24/icon.png?size=72","url144":"https://api.atlassian.com/jsm/insight/workspace/f1668d0c-828c-470c-b7d1-8c4f48cd345a/v1/objecttype/24/icon.png?size=144","url288":"https://api.atlassian.com/jsm/insight/workspace/f1668d0c-828c-470c-b7d1-8c4f48cd345a/v1/objecttype/24/icon.png?size=288","objectId":"87","mediaClientConfig":{"clientId":"1a2s3d4f-dc47-44b0-9t0r-1h2h3yd68e9q","mediaBaseUrl":"https://api.media.atlassian.com","mediaJwtToken":"eyJhbGciOiJIUzI1NiJ9.eyJpc3MiOiIxYTJzM2Q0Zi1kYzQ3LTQ0YjAtOXQwci0xaDJoM3lkNjhlOXEiLCJhY2Nlc3MiOnsidXJuOmZpbGVzdG9yZTpmaWxlOjg0MTIzZXJ0LTEyM2MtNGIxMi0xMmM1LTBiODZkYzgxMjNmZiI6WyJyZWFkIl19LCJleHAiOjE2MjYxNTY1NjcsIm5iZiI6MTYyNjE1NTkwN30.YjicbagPLbzapp3eEZbCQ7Z9V8Uc0WeBledyTw-Qu0s","fileId":"84123ert-123c-4b12-12c5-0b86dc8123ff"}},"objectType":{"workspaceId":"f1668d0c-828c-470c-b7d1-8c4f48cd345a","globalId":"f1668d0c-828c-470c-b7d1-8c4f48cd345a:24","id":"24","name":"City","type":0,"icon":{"id":"28","name":"Cottage","url16":"https://api.atlassian.com/jsm/insight/workspace/f1668d0c-828c-470c-b7d1-8c4f48cd345a/v1/objecttype/24/icon.png?size=16","url48":"https://api.atlassian.com/jsm/insight/workspace/f1668d0c-828c-470c-b7d1-8c4f48cd345a/v1/objecttype/24/icon.png?size=48"},"position":3,"created":"2021-02-16T19:58:45.698Z","updated":"2021-04-16T15:17:03.393Z","objectCount":0,"objectSchemaId":"6","inherited":false,"abstractObjectType":false,"parentObjectTypeInherited":false},"created":"2021-02-16T20:04:26.445Z","updated":"2021-02-16T20:04:26.445Z","hasAvatar":false,"timestamp":1613505866445,"_links":{"self":"https://api.atlassian.com/jsm/insight/workspace/f1668d0c-828c-470c-b7d1-8c4f48cd345a/v1/object/87"},"name":"Sydney"},"displayValue":"Sydney","searchValue":"ITSM-87","referencedType":true}],"objectId":"88"},{"workspaceId":"f1668d0c-828c-470c-b7d1-8c4f48cd345a","globalId":"f1668d0c-828c-470c-b7d1-8c4f48cd345a:44662","id":"44662","objectTypeAttribute":{"workspaceId":"f1668d0c-828c-470c-b7d1-8c4f48cd345a","globalId":"f1668d0c-828c-470c-b7d1-8c4f48cd345a:265","id":"265","name":"Placeholder","label":false,"type":0,"defaultType":{"id":0,"name":"Text"},"editable":true,"system":false,"sortable":true,"summable":false,"indexed":true,"minimumCardinality":0,"maximumCardinality":1,"removable":true,"hidden":false,"includeChildObjectTypes":false,"uniqueAttribute":false,"options":"","position":5},"objectTypeAttributeId":"265","objectAttributeValues":[{"value":"A placeholder value","displayValue":"A placeholder value","searchValue":"A placeholder value","referencedType":false}],"objectId":"88"}]`)
	})
	if attributes, err := testClient.Insight.ObjectType.GetAttributes(context.Background(), "g2778e1d-939d-581d-c8e2-9d5g59de456b", "23", &insight.ObjectTypeAttributeOptions{
		OnlyValueEditable: true,
		OrderByName:       true,
		OrderByRequired:   true,
	}); err != nil {
		t.Errorf("Error given: %s", err)
	} else if attributes == nil {
		t.Error("Expected attributes. Object is nil")
	}
}

func TestInsightObjectTypeService_UpdatePosition(t *testing.T) {
	setup()
	defer teardown()
	testMux.HandleFunc("/jsm/insight/workspace/g2778e1d-939d-581d-c8e2-9d5g59de456b/v1/objecttype/23/position", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		testRequestURL(t, r, "/jsm/insight/workspace/g2778e1d-939d-581d-c8e2-9d5g59de456b/v1/objecttype/23/position")
		fmt.Fprint(w, `{"workspaceId":"g2778e1d-939d-581d-c8e2-9d5g59de456b","globalId":"g2778e1d-939d-581d-c8e2-9d5g59de456b:23","id":"23","name":"Office","description":"Lorem ipsum dolor sit amet, consectetur adipiscing elit. Proin nec ex.","icon":{"id":"13","name":"Building","url16":"https://api.atlassian.com/jsm/insight/workspace/g2778e1d-939d-581d-c8e2-9d5g59de456b/v1/objecttype/23/icon.png?size=16","url48":"https://api.atlassian.com/jsm/insight/workspace/g2778e1d-939d-581d-c8e2-9d5g59de456b/v1/objecttype/23/icon.png?size=48"},"position":2,"created":"2021-02-16T19:36:51.951Z","updated":"2021-04-16T15:17:03.384Z","objectCount":4,"objectSchemaId":"6","inherited":false,"abstractObjectType":false,"parentObjectTypeInherited":false}`)
	})
	if objectType, err := testClient.Insight.ObjectType.UpdatePosition(context.Background(), "g2778e1d-939d-581d-c8e2-9d5g59de456b", "23", insight.PostObjectTypePosition{}); err != nil {
		t.Errorf("Error given: %s", err)
	} else if objectType == nil {
		t.Error("Expected objectType. Object is nil")
	}
}

func TestInsightObjectTypeService_Create(t *testing.T) {
	setup()
	defer teardown()
	testMux.HandleFunc("/jsm/insight/workspace/g2778e1d-939d-581d-c8e2-9d5g59de456b/v1/objecttype/create", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		testRequestURL(t, r, "/jsm/insight/workspace/g2778e1d-939d-581d-c8e2-9d5g59de456b/v1/objecttype/create")
		fmt.Fprint(w, `{"workspaceId":"g2778e1d-939d-581d-c8e2-9d5g59de456b","globalId":"g2778e1d-939d-581d-c8e2-9d5g59de456b:23","id":"23","name":"Office","description":"Lorem ipsum dolor sit amet, consectetur adipiscing elit. Proin nec ex.","icon":{"id":"13","name":"Building","url16":"https://api.atlassian.com/jsm/insight/workspace/g2778e1d-939d-581d-c8e2-9d5g59de456b/v1/objecttype/23/icon.png?size=16","url48":"https://api.atlassian.com/jsm/insight/workspace/g2778e1d-939d-581d-c8e2-9d5g59de456b/v1/objecttype/23/icon.png?size=48"},"position":2,"created":"2021-02-16T19:36:51.951Z","updated":"2021-04-16T15:17:03.384Z","objectCount":4,"objectSchemaId":"6","inherited":false,"abstractObjectType":false,"parentObjectTypeInherited":false}`)
	})
	if objectType, err := testClient.Insight.ObjectType.Create(context.Background(), "g2778e1d-939d-581d-c8e2-9d5g59de456b", insight.CreateObjectType{}); err != nil {
		t.Errorf("Error given: %s", err)
	} else if objectType == nil {
		t.Error("Expected objectType. Object is nil")
	}
}
