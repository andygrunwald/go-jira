package main

import (
	"fmt"

	jira "github.com/andygrunwald/go-jira"
)

func main() {
	jiraClient, _ := jira.NewClient(nil, "https://issues.apache.org/jira/")

	// Running JQL query

	jql := "project = Mesos and type = Bug and Status NOT IN (Resolved)"
	fmt.Printf("Usecase: Running a JQL query '%s'\n", jql)
	issues, resp, err := jiraClient.Issue.Search(jql, nil)
	if err != nil {
		panic(err)
	}
	outputResponse(issues, resp)

	fmt.Println("")
	fmt.Println("")

	// Running an empty JQL query to get all tickets
	jql = ""
	fmt.Printf("Usecase: Running an empty JQL query to get all tickets\n")
	issues, resp, err = jiraClient.Issue.Search(jql, nil)
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
