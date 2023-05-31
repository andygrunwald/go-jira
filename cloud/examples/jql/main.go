package main

import (
	"context"
	"fmt"

	jira "github.com/andygrunwald/go-jira/v2/cloud"
)

func main() {
	jiraClient, _ := jira.NewClient("https://issues.apache.org/jira/", nil)

	// Running JQL query

	jql := "project = Mesos and type = Bug and Status NOT IN (Resolved)"
	fmt.Printf("Usecase: Running a JQL query '%s'\n", jql)
	issues, resp, err := jiraClient.Issue.Search(context.Background(), jql, nil)
	if err != nil {
		panic(err)
	}
	outputResponse(issues, resp)

	fmt.Println("")
	fmt.Println("")

	// Running an empty JQL query to get all tickets
	jql = ""
	fmt.Printf("Usecase: Running an empty JQL query to get all tickets\n")
	issues, resp, err = jiraClient.Issue.Search(context.Background(), jql, nil)
	if err != nil {
		panic(err)
	}
	outputResponse(issues, resp)
}

func outputResponse(issues []jira.Issue, resp *jira.Response) {
	fmt.Printf("Call to %s\n", resp.Request.URL)
	fmt.Printf("Response Code: %d\n", resp.StatusCode)
	fmt.Println("==================================")
	for _, i := range issues {
		fmt.Printf("%s (%s/%s): %+v\n", i.Key, i.Fields.Type.Name, i.Fields.Priority.Name, i.Fields.Summary)
	}
}
