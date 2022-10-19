package main

import (
	"context"
	"fmt"

	jira "github.com/andygrunwald/go-jira/v2/cloud"
)

func main() {
	jiraURL := "https://go-jira-opensource.atlassian.net/"

	// Jira docs: https://support.atlassian.com/atlassian-account/docs/manage-api-tokens-for-your-atlassian-account/
	// Create a new API token: https://id.atlassian.com/manage-profile/security/api-tokens
	tp := jira.BasicAuthTransport{
		Username: "<username>",
		APIToken: "<api-token>",
	}
	client, err := jira.NewClient(jiraURL, tp.Client())
	if err != nil {
		panic(err)
	}

	c := &jira.ComponentCreateOptions{
		Name:         "Dummy component",
		AssigneeType: jira.AssigneeTypeUnassigned,
		Project:      "BUG",
	}
	component, _, err := client.Component.Create(context.Background(), c)
	if err != nil {
		panic(err)
	}

	fmt.Printf("component: %+v\n", component)
	fmt.Println("Success!")
}
