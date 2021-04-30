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

	tp := jira.BasicAuthTransport{
		Username: strings.TrimSpace(username),
		Password: strings.TrimSpace(password),
	}

	client, err := jira.NewClient(tp.Client(), strings.TrimSpace(jiraURL))
	if err != nil {
		fmt.Printf("\nerror: %v\n", err)
		return
	}

	// the the records associated with the creation on the last 2 hours
	options := &jira.AuditSearchOptionsScheme{
		Filter:     "created",
		From:       time.Now().Add(time.Duration(-2) * time.Hour), // Last 2 hours
		To:         time.Time{},
		ProjectIDs: nil,
		UserIDs:    nil,
	}

	records, response, err := client.Audit.Search(context.Background(), options, 0, 50)
	if err != nil {
		if response != nil {
			log.Println("Response HTTP Response", response.StatusCode)
		}
		log.Fatal(err)
	}

	for index, record := range records.Records {

		log.Println("----------------------------------")
		log.Printf("Record #%v", index+1)
		log.Printf("Record ID: %v", record.ID)
		log.Printf("Record Created by: %v", record.Created)
		log.Printf("Record Category: %v", record.Category)
		log.Printf("Record Source IP: %v", record.RemoteAddress)
		log.Printf("Record Event: %v", record.EventSource)
		log.Println("----------------------------------")
	}

}
