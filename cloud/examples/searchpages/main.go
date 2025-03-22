package main

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	jira "github.com/andygrunwald/go-jira/v2/cloud"
)

func main() {
	jiraURL := "https://issues.apache.org/jira/"
	username := "my.username"
	password := "my.secret.password"
	jiraProjectKey := "TES"

	tp := jira.BasicAuthTransport{
		Username: strings.TrimSpace(username),
		APIToken: strings.TrimSpace(password),
	}

	client, err := jira.NewClient(strings.TrimSpace(jiraURL), tp.Client())
	if err != nil {
		log.Fatal(err)
	}

	var issues []jira.Issue

	// appendFunc will append jira issues to []jira.Issue
	appendFunc := func(i jira.Issue) (err error) {
		issues = append(issues, i)
		return err
	}

	// SearchPages will page through results and pass each issue to appendFunc
	// In this example, we'll search for all the issues in the target project
	err = client.Issue.SearchPages(context.Background(), fmt.Sprintf(`project=%s`, strings.TrimSpace(jiraProjectKey)), nil, appendFunc)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%d issues found.\n", len(issues))

	for _, i := range issues {
		t := time.Time(i.Fields.Created) // convert go-jira.Time to time.Time for manipulation
		date := t.Format("2006-01-02")
		clock := t.Format("15:04")
		fmt.Printf("Creation Date: %s\nCreation Time: %s\nIssue Key: %s\nIssue Summary: %s\n\n", date, clock, i.Key, i.Fields.Summary)
	}

}
