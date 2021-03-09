package main

import (
	"bufio"
	"fmt"
	jira "github.com/andygrunwald/go-jira"
	"golang.org/x/term"
	"log"
	"os"
	"strings"
	"syscall"
	"time"
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

	fmt.Print("\nJira Project Key: ") // e.g. TES or WOW
	jiraPK, _ := r.ReadString('\n')

	tp := jira.BasicAuthTransport{
		Username: strings.TrimSpace(username),
		Password: strings.TrimSpace(password),
	}

	client, err := jira.NewClient(tp.Client(), strings.TrimSpace(jiraURL))
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
	err = client.Issue.SearchPages(fmt.Sprintf(`project=%s`, strings.TrimSpace(jiraPK)), nil, appendFunc)
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
