package main

import (
	"bufio"
	"context"
	"fmt"
	"github.com/andygrunwald/go-jira"
	"golang.org/x/term"
	"log"
	"os"
	"strings"
	"syscall"
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

	client, err := jira.NewClient(tp.Client(), strings.TrimSpace(jiraURL))
	if err != nil {
		fmt.Printf("\nerror: %v\n", err)
		return
	}

	newRecord := &jira.AuditRecord{
		Summary:  "User created",
		Category: "USER_MANAGEMENT",
		ObjectItem: &jira.AuditRecordObjectItem{
			ID:       "usr",
			Name:     "user",
			TypeName: "USER",
		},
		ChangedValues: []*jira.AuditRecordChangedValue{
			{
				FieldName:   "email",
				ChangedTo:   "newuser@atlassian.com",
				ChangedFrom: "user@atlassian.com",
			},
		},

		AssociatedItems: []*jira.AuditRecordAssociatedItem{
			{
				ID:         "jira-software-users",
				Name:       "jira-software-users",
				TypeName:   "GROUP",
				ParentID:   "1",
				ParentName: "Jira Internal Directory",
			},
		},
	}

	response, err := client.Audit.Add(context.Background(), newRecord)
	if err != nil {
		if response != nil {
			log.Println("Response HTTP Response", response.StatusCode)
		}
		log.Fatal(err)
	}

	log.Println(response.StatusCode)
}
