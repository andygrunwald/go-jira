package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"syscall"

	jira "github.com/andygrunwald/go-jira"
	"github.com/trivago/tgo/tcontainer"
	"golang.org/x/crypto/ssh/terminal"
)

func main() {
	r := bufio.NewReader(os.Stdin)

	fmt.Print("Jira URL: ")
	jiraURL, _ := r.ReadString('\n')

	fmt.Print("Jira Username: ")
	username, _ := r.ReadString('\n')

	fmt.Print("Jira Password: ")
	bytePassword, _ := terminal.ReadPassword(int(syscall.Stdin))
	password := string(bytePassword)

	fmt.Print("Custom field name (i.e. customfield_10220): ")
	customFieldName, _ := r.ReadString('\n')

	fmt.Print("Custom field value: ")
	customFieldValue, _ := r.ReadString('\n')

	tp := jira.BasicAuthTransport{
		Username: strings.TrimSpace(username),
		Password: strings.TrimSpace(password),
	}

	client, err := jira.NewClient(tp.Client(), strings.TrimSpace(jiraURL))
	if err != nil {
		fmt.Printf("\nerror: %v\n", err)
		return
	}

	unknowns := tcontainer.NewMarshalMap()
	unknowns[customFieldName] = customFieldValue

	i := jira.Issue{
		Fields: &jira.IssueFields{
			Assignee: &jira.User{
				Name: "myuser",
			},
			Reporter: &jira.User{
				Name: "youruser",
			},
			Description: "Test Issue",
			Type: jira.IssueType{
				Name: "Bug",
			},
			Project: jira.Project{
				Key: "PROJ1",
			},
			Summary:  "Just a demo issue",
			Unknowns: unknowns,
		},
	}

	issue, resp, err := client.Issue.Create(&i)
	if err != nil {
		_, err := ioutil.ReadAll(resp.Body)
		panic(err)
	}

	fmt.Printf("%s: %v\n", issue.Key, issue.Self)
}
