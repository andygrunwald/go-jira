package main

import (
	"fmt"

	jira "github.com/andygrunwald/go-jira"
)

//GetAllIssues takes a jira client and returns all issues for given JQL
func GetAllIssues(client *jira.Client, searchString string) ([]jira.Issue, error) {
	last := 0
	var issues []jira.Issue = nil
	for {
		opt := &jira.SearchOptions{
			MaxResults: 100,
			StartAt:    last,
		}

		chunk, resp, err := client.Issue.Search(searchString, opt)
		if err != nil {
			return nil, err
		}

		total := resp.Total
		if issues == nil {
			issues = make([]jira.Issue, 0, total)
		}
		issues = append(issues, chunk...)
		last = resp.StartAt + len(chunk)
		if last >= total {
			break
		}
	}
	return issues, nil
}

func main() {
	jiraClient, _ := jira.NewClient(nil, "https://issues.apache.org/jira/")

	// Jira API has limitation as to maxResults it can return at one time.
	// You may have usecase where you need to get all the issues according to jql
	// This is where this example comes in.
	jql := "project = Mesos and type = Bug and Status NOT IN (Resolved)"
	fmt.Printf("Usecase: Running a JQL query '%s'\n", jql)

	issues, err := GetAllIssues(jiraClient, jql)
	if err != nil {
		panic(err)
	}
	fmt.Println(issues)

}
