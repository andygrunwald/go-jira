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
	issues, resp, warningMessages, err := jiraClient.Issue.Search(jql, nil)
	if err != nil {
		panic(err)
	}
	checkWarnings(warningMessages)
	outputResponse(issues, resp)

	fmt.Println("")
	fmt.Println("")

	// Running an empty JQL query to get all tickets
	jql = ""
	fmt.Printf("Usecase: Running an empty JQL query to get all tickets\n")
	issues, resp, warningMessages, err = jiraClient.Issue.Search(jql, nil)
	if err != nil {
		panic(err)
	}
	checkWarnings(warningMessages)
	outputResponse(issues, resp)
}

func outputResponse(issues []jira.Issue, resp *jira.Response) {
	fmt.Printf("Response Code: %d\n", resp.StatusCode)
	fmt.Println("==================================")
	for _, i := range issues {
		fmt.Printf("%s (%s/%s): %+v\n", i.Key, i.Fields.Type.Name, i.Fields.Priority.Name, i.Fields.Summary)
	}
}

func checkWarnings(warningMessages []jira.WarningMsg) {
	if len(warningMessages) > 0 {
		fmt.Printf("Warning messages in response:\n")
		for i, warn := range warningMessages {
			fmt.Printf("Warning %d: %s\n", i, warn)
		}
	}
}
