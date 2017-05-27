package main

import (
	"fmt"

	"github.com/andygrunwald/go-jira"
)

func main() {
	jiraClient, err := jira.NewClient(nil, "https://your.jira-instance.com/")
	if err != nil {
		panic(err)
	}

	res, err := jiraClient.Authentication.AcquireSessionCookie("username", "password")
	if err != nil || res == false {
		fmt.Printf("Result: %v\n", res)
		panic(err)
	}

	i := jira.Issue{
		Fields: &jira.IssueFields{
			Assignee: &jira.User{
				Name: "myuser",
			},
			Reporter: &jira.User{
				Name: "youruser",
			},
			Description: "Test Issue",
			Type: jira.IssueType{
				ID: "60",
			},
			Project: jira.Project{
				Name: "PROJ1",
			},
			Summary: "Just a demo issue",
		},
	}
	issue, _, err := jiraClient.Issue.Create(&i)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s: %+v\n", issue.Key, issue.Fields.Summary)
}
