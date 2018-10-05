package main

import (
	"fmt"

	jira "github.com/andygrunwald/go-jira"
)

func main() {
	jiraClient, _ := jira.NewClient(nil, "https://jira.atlassian.com/")
	req, _ := jiraClient.NewRequest("GET", "/rest/api/2/project", nil)

	projects := new([]jira.Project)
	_, err := jiraClient.Do(req, projects)
	if err != nil {
		panic(err)
	}

	for _, project := range *projects {
		fmt.Printf("%s: %s\n", project.Key, project.Name)
	}
}
