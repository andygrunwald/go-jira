package main

import (
	"bufio"
	"fmt"
	"golang.org/x/term"
	"net/http"
	"os"
	"strings"
	"syscall"

	jira "github.com/andygrunwald/go-jira"
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
			Password: strings.TrimSpace(password),
		}
		tp = ba.Client()
	}

	client, err := jira.NewClient(tp, strings.TrimSpace(jiraURL))
	if err != nil {
		fmt.Printf("\nerror: %v\n", err)
		return
	}

	fmt.Printf("Targeting %s for issue %s\n", strings.TrimSpace(jiraURL), key)

	options := &jira.GetQueryOptions{Expand: "renderedFields"}
	u, _, err := client.Issue.Get(key, options)

	if err != nil {
		fmt.Printf("\n==> error: %v\n", err)
		return
	}

	fmt.Printf("RenderedFields: %+v\n", *u.RenderedFields)

	for _, c := range u.RenderedFields.Comments.Comments {
		fmt.Printf("  %+v\n", c)
	}
}
