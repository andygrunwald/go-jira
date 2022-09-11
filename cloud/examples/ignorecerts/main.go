package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"net/http"

	jira "github.com/andygrunwald/go-jira/v2/cloud"
)

func main() {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	jiraClient, _ := jira.NewClient("https://issues.apache.org/jira/", client)
	issue, _, _ := jiraClient.Issue.Get(context.Background(), "MESOS-3325", nil)

	fmt.Printf("%s: %+v\n", issue.Key, issue.Fields.Summary)
	fmt.Printf("Type: %s\n", issue.Fields.Type.Name)
	fmt.Printf("Priority: %s\n", issue.Fields.Priority.Name)

}
