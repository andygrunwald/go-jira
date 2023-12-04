package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"
	"syscall"

	"golang.org/x/term"

	jira "github.com/daveoy/go-jira/v2/onpremise"
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
