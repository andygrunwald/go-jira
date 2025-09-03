package main

import (
	"context"
	"fmt"

	jira "github.com/andygrunwald/go-jira/v2/cloud"
)

// GetAllIssues will implement pagination of api and get all the issues.
// Jira API has limitation as to maxResults it can return at one time.
// You may have usecase where you need to get all the issues according to jql
// This is where this example comes in.
func GetAllIssues(client *jira.Client, searchString string) ([]jira.Issue, error) {
	var issues []jira.Issue
	opt := &jira.SearchOptions{
		MaxResults: 1000, // Max results can go up to 1000
	}

	for {
		chunk, resp, err := client.Issue.Search(context.Background(), searchString, opt)
		if err != nil {
			return nil, err
		}

		issues = append(issues, chunk...)

		if resp.IsLast {
			return issues, nil
		}

		// Set the next page token for the next iteration
		opt.NextPageToken = resp.NextPageToken
	}
}

func main() {
	jiraClient, err := jira.NewClient("https://issues.apache.org/jira/", nil)
	if err != nil {
		panic(err)
	}

	jql := "project = Mesos and type = Bug and Status NOT IN (Resolved)"
	fmt.Printf("Usecase: Running a JQL query '%s'\n", jql)

	issues, err := GetAllIssues(jiraClient, jql)
	if err != nil {
		panic(err)
	}
	fmt.Println(issues)

}
