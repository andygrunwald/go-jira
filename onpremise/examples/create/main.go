package main

import (
	"context"
	"fmt"
	"strings"

	jira "github.com/andygrunwald/go-jira/v2/onpremise"
)

func main() {
	jiraURL := "https://issues.apache.org/jira/"
	username := "my.username"
	password := "my.secret.password"

	tp := jira.BasicAuthTransport{
		Username: strings.TrimSpace(username),
		Password: strings.TrimSpace(password),
	}

	client, err := jira.NewClient(strings.TrimSpace(jiraURL), tp.Client())
	if err != nil {
		fmt.Printf("\nerror: %v\n", err)
		return
	}

	i := jira.Issue{
		Fields: &jira.IssueFields{
			Assignee: &jira.User{
				AccountID: "my-user-account-id",
			},
			Reporter: &jira.User{
				AccountID: "your-user-account-id",
			},
			Description: "Test Issue",
			Type: jira.IssueType{
				Name: "Bug",
			},
			Project: jira.Project{
				Key: "PROJ1",
			},
			Summary: "Just a demo issue",
		},
	}

	issue, _, err := client.Issue.Create(context.Background(), &i)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s: %+v\n", issue.Key, issue.Self)
}
