package cloud

import (
	"context"
	"fmt"
	"net/http"
	"testing"
)

func TestInsightIconService_Get(t *testing.T) {
	setup()
	defer teardown()
	testMux.HandleFunc("/jsm/insight/workspace/g2778e1d-939d-581d-c8e2-9d5g59de456b/v1/icon/68", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testRequestURL(t, r, "/jsm/insight/workspace/g2778e1d-939d-581d-c8e2-9d5g59de456b/v1/icon/68")
		fmt.Fprint(w, `{"id":"68","name":"Mac OS","url16":"https://api.atlassian.com/jsm/insight/workspace/g2778e1d-939d-581d-c8e2-9d5g59de456b/v1/icon/68/icon.png?size=16","url48":"https://api.atlassian.com/jsm/insight/workspace/g2778e1d-939d-581d-c8e2-9d5g59de456b/v1/icon/68/icon.png?size=48"}`)
	})
	if objectSchemas, err := testClient.Insight.Icon.Get(context.Background(), "g2778e1d-939d-581d-c8e2-9d5g59de456b", "68"); err != nil {
		t.Errorf("Error given: %s", err)
	} else if objectSchemas == nil {
		t.Error("Expected objectSchemas. Object is nil")
	}
}

func TestInsightIconService_GetGlobal(t *testing.T) {
	setup()
	defer teardown()
	testMux.HandleFunc("/jsm/insight/workspace/g2778e1d-939d-581d-c8e2-9d5g59de456b/v1/icon/global", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testRequestURL(t, r, "/jsm/insight/workspace/g2778e1d-939d-581d-c8e2-9d5g59de456b/v1/icon/global")
		fmt.Fprint(w, `[{"id":"68","name":"Mac OS","url16":"https://api.atlassian.com/jsm/insight/workspace/g2778e1d-939d-581d-c8e2-9d5g59de456b/v1/icon/68/icon.png?size=16","url48":"https://api.atlassian.com/jsm/insight/workspace/g2778e1d-939d-581d-c8e2-9d5g59de456b/v1/icon/68/icon.png?size=48"},{"id":"69","name":"Marker","url16":"https://api.atlassian.com/jsm/insight/workspace/g2778e1d-939d-581d-c8e2-9d5g59de456b/v1/icon/69/icon.png?size=16","url48":"https://api.atlassian.com/jsm/insight/workspace/g2778e1d-939d-581d-c8e2-9d5g59de456b/v1/icon/69/icon.png?size=48"},{"id":"70","name":"Master","url16":"https://api.atlassian.com/jsm/insight/workspace/g2778e1d-939d-581d-c8e2-9d5g59de456b/v1/icon/70/icon.png?size=16","url48":"https://api.atlassian.com/jsm/insight/workspace/g2778e1d-939d-581d-c8e2-9d5g59de456b/v1/icon/70/icon.png?size=48"},{"id":"71","name":"Memory Slot","url16":"https://api.atlassian.com/jsm/insight/workspace/g2778e1d-939d-581d-c8e2-9d5g59de456b/v1/icon/71/icon.png?size=16","url48":"https://api.atlassian.com/jsm/insight/workspace/g2778e1d-939d-581d-c8e2-9d5g59de456b/v1/icon/71/icon.png?size=48"},{"id":"72","name":"Mic","url16":"https://api.atlassian.com/jsm/insight/workspace/g2778e1d-939d-581d-c8e2-9d5g59de456b/v1/icon/72/icon.png?size=16","url48":"https://api.atlassian.com/jsm/insight/workspace/f1668d0c-828c-470c-b7d1-8c4f48cd345a/v1/icon/72/icon.png?size=48"}]`)
	})
	if objectSchemas, err := testClient.Insight.Icon.GetGlobal(context.Background(), "g2778e1d-939d-581d-c8e2-9d5g59de456b"); err != nil {
		t.Errorf("Error given: %s", err)
	} else if objectSchemas == nil {
		t.Error("Expected objectSchemas. Object is nil")
	}
}
