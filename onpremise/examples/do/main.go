package main

import (
	"context"
	"fmt"
	"net/http"

	jira "github.com/andygrunwald/go-jira/v2/onpremise"
)

func main() {
	jiraClient, _ := jira.NewClient("https://jira.atlassian.com/", nil)
	req, _ := jiraClient.NewRequest(context.Background(), http.MethodGet, "/rest/api/2/project", nil)

	projects := new([]jira.Project)
	res, err := jiraClient.Do(req, projects)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	for _, project := range *projects {
		fmt.Printf("%s: %s\n", project.Key, project.Name)
	}
}
