package main

import (
	"context"
	"fmt"
	"strings"

	jira "github.com/andygrunwald/go-jira/v2/onpremise"
)

func main() {
	jiraURL := "https://issues.apache.org/jira/"
	username := "my.username"
	password := "my.secret.password"

	tp := jira.BasicAuthTransport{
		Username: strings.TrimSpace(username),
		Password: strings.TrimSpace(password),
	}

	client, err := jira.NewClient(strings.TrimSpace(jiraURL), tp.Client())
	if err != nil {
		fmt.Printf("\nerror: %v\n", err)
		return
	}

	u, _, err := client.User.Get(context.Background(), "admin")

	if err != nil {
		fmt.Printf("\nerror: %v\n", err)
		return
	}

	fmt.Printf("\nEmail: %v\nSuccess!\n", u.EmailAddress)

}
