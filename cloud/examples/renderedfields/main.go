package main

import (
	"bufio"
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"
	"syscall"

	"golang.org/x/term"

	jira "github.com/daveoy/go-jira/v2/cloud"
)

func main() {
	r := bufio.NewReader(os.Stdin)

	fmt.Print("Jira URL: ")
	jiraURL, _ := r.ReadString('\n')

	fmt.Print("Jira Issue key: ")
	key, _ := r.ReadString('\n')
	key = strings.TrimSpace(key)

	fmt.Print("Jira Username: ")
	username, _ := r.ReadString('\n')

	fmt.Print("Jira Password: ")
	bytePassword, _ := term.ReadPassword(int(syscall.Stdin))
	password := string(bytePassword)

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

	fmt.Printf("Targeting %s for issue %s\n", strings.TrimSpace(jiraURL), key)

	options := &jira.GetQueryOptions{Expand: "renderedFields"}
	u, _, err := client.Issue.Get(context.Background(), key, options)

	if err != nil {
		fmt.Printf("\n==> error: %v\n", err)
		return
	}

	fmt.Printf("RenderedFields: %+v\n", *u.RenderedFields)

	for _, c := range u.RenderedFields.Comments.Comments {
		fmt.Printf("  %+v\n", c)
	}
}
