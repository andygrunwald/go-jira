package main

import (
	"fmt"

	jira "github.com/andygrunwald/go-jira/cloud"
)

func main() {
	jiraClient, _ := jira.NewClient("https://issues.apache.org/jira/", nil)
	issue, _, _ := jiraClient.Issue.Get("MESOS-3325", nil)

	fmt.Printf("%s: %+v\n", issue.Key, issue.Fields.Summary)
	fmt.Printf("Type: %s\n", issue.Fields.Type.Name)
	fmt.Printf("Priority: %s\n", issue.Fields.Priority.Name)
}
