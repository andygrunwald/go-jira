package main

import (
	"context"
	"fmt"

	jira "github.com/andygrunwald/go-jira/v2/cloud"
)

func main() {
	tp := jira.BasicAuthTransport{
		Username: "<username>",
		APIToken: "<api-token>",
	}
	jiraClient, _ := jira.NewClient("https://go-jira-opensource.atlassian.net/", tp.Client())

	// Running JQL query
	jql := "type = Bug and Status NOT IN (Resolved)"
	fmt.Printf("Usecase: Running a JQL query '%s'\n", jql)
	options := &jira.SearchOptionsV2{
		Fields: []string{"*all"},
	}
	issues, resp, err := jiraClient.Issue.SearchV2JQL(context.Background(), jql, options)
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
