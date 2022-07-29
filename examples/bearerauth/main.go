package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"syscall"

	"golang.org/x/term"

	jira "github.com/andygrunwald/go-jira"
)

func main() {
	r := bufio.NewReader(os.Stdin)

	fmt.Print("Jira URL: ")
	jiraURL, _ := r.ReadString('\n')

	fmt.Print("Jira Personal Access Token: ")
	byteToken, _ := term.ReadPassword(int(syscall.Stdin))
	token := string(byteToken)

	tp := jira.BearerAuthTransport{
		Token: string(token),
	}

	client, err := jira.NewClient(tp.Client(), strings.TrimSpace(jiraURL))
	if err != nil {
		fmt.Printf("\nerror: %v\n", err)
		return
	}

	u, _, err := client.User.Get("admin")

	if err != nil {
		fmt.Printf("\nerror: %v\n", err)
		return
	}

	fmt.Printf("\nEmail: %v\nSuccess!\n", u.EmailAddress)

}
