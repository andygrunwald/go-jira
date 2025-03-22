package main

import (
	"context"
	"fmt"
	"io"
	"strings"

	jira "github.com/andygrunwald/go-jira/v2/cloud"
)

func main() {
	jiraURL := "https://issues.apache.org/jira/"
	username := "my.username"
	password := "my.secret.password"
	issueId := "MESOS-3325"
	label := "example-label"

	tp := jira.BasicAuthTransport{
		Username: strings.TrimSpace(username),
		APIToken: strings.TrimSpace(password),
	}

	client, err := jira.NewClient(strings.TrimSpace(jiraURL), tp.Client())
	if err != nil {
		fmt.Printf("\nerror: %v\n", err)
		return
	}

	type Labels struct {
		Add string `json:"add" structs:"add"`
	}

	type Update struct {
		Labels []Labels `json:"labels" structs:"labels"`
	}

	c := map[string]interface{}{
		"update": Update{
			Labels: []Labels{
				{
					Add: label,
				},
			},
		},
	}

	resp, err := client.Issue.UpdateIssue(context.Background(), issueId, c)

	if err != nil {
		fmt.Println(err)
	}
	body, _ := io.ReadAll(resp.Body)
	fmt.Println(string(body))

	issue, _, _ := client.Issue.Get(context.Background(), issueId, nil)

	fmt.Printf("Issue: %s:%s\n", issue.Key, issue.Fields.Summary)
	fmt.Printf("\tLabels: %+v\n", issue.Fields.Labels)
}
