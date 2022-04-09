package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"syscall"

	jira "github.com/andygrunwald/go-jira"
	"golang.org/x/term"
)

func main() {
	r := bufio.NewReader(os.Stdin)

	fmt.Print("Jira URL: ")
	jiraURL, _ := r.ReadString('\n')

	fmt.Print("Jira Username: ")
	username, _ := r.ReadString('\n')

	fmt.Print("Jira Password: ")
	bytePassword, _ := term.ReadPassword(int(syscall.Stdin))
	password := string(bytePassword)

	fmt.Print("Jira Issue ID: ")
	issueId, _ := r.ReadString('\n')
	issueId = strings.TrimSpace(issueId)

	fmt.Print("Label: ")
	label, _ := r.ReadString('\n')
	label = strings.TrimSpace(label)

	tp := jira.BasicAuthTransport{
		Username: strings.TrimSpace(username),
		Password: strings.TrimSpace(password),
	}

	client, err := jira.NewClient(tp.Client(), strings.TrimSpace(jiraURL))
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

	resp, err := client.Issue.UpdateIssue(issueId, c)

	if err != nil {
		fmt.Println(err)
	}
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))

	issue, _, _ := client.Issue.Get(issueId, nil)

	fmt.Printf("Issue: %s:%s\n", issue.Key, issue.Fields.Summary)
	fmt.Printf("\tLabels: %+v\n", issue.Fields.Labels)
}
