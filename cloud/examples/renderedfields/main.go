package main

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	jira "github.com/andygrunwald/go-jira/v2/cloud"
)

func main() {
	jiraURL := "https://issues.apache.org/jira/"
	username := "my.username"
	password := "my.secret.password"
	issueId := "MESOS-3325"

	var tp *http.Client

	if strings.TrimSpace(username) == "" {
		tp = nil
	} else {

		ba := jira.BasicAuthTransport{
			Username: strings.TrimSpace(username),
			APIToken: strings.TrimSpace(password),
		}
		tp = ba.Client()
	}

	client, err := jira.NewClient(strings.TrimSpace(jiraURL), tp)
	if err != nil {
		fmt.Printf("\nerror: %v\n", err)
		return
	}

	fmt.Printf("Targeting %s for issue %s\n", strings.TrimSpace(jiraURL), issueId)

	options := &jira.GetQueryOptions{Expand: "renderedFields"}
	u, _, err := client.Issue.Get(context.Background(), issueId, options)

	if err != nil {
		fmt.Printf("\n==> error: %v\n", err)
		return
	}

	fmt.Printf("RenderedFields: %+v\n", *u.RenderedFields)

	for _, c := range u.RenderedFields.Comments.Comments {
		fmt.Printf("  %+v\n", c)
	}
}
