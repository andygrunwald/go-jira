package main

import (
	"context"
	"fmt"
	"os"
	"strings"

	jira "github.com/andygrunwald/go-jira/v2/cloud"
	"github.com/trivago/tgo/tcontainer"
)

func main() {
	jiraURL := "https://issues.apache.org/jira/"
	username := "my.username"
	password := "my.secret.password"

	customFieldName := "customfield_10220"
	customFieldValue := "foo"

	tp := jira.BasicAuthTransport{
		Username: strings.TrimSpace(username),
		APIToken: strings.TrimSpace(password),
	}

	client, err := jira.NewClient(strings.TrimSpace(jiraURL), tp.Client())
	if err != nil {
		fmt.Printf("\nerror: %v\n", err)
		os.Exit(1)
	}

	unknowns := tcontainer.NewMarshalMap()
	unknowns[customFieldName] = customFieldValue

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
				Name: "Bug",
			},
			Project: jira.Project{
				Key: "PROJ1",
			},
			Summary:  "Just a demo issue",
			Unknowns: unknowns,
		},
	}

	issue, _, err := client.Issue.Create(context.Background(), &i)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s: %v\n", issue.Key, issue.Self)
}
