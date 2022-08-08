package main

import (
	"bufio"
	"fmt"
	"golang.org/x/term"
	"log"
	"os"
	"strconv"
	"syscall"
	"time"

	jira "github.com/andygrunwald/go-jira"
)

// This example implement the behaviour of Jira's "Complete sprint" button.
// It creates a new sprint and close the sprint currently active.
// Then all issues under the closed sprint are moved to the newly created sprint.
func main() {
	sc := bufio.NewScanner(os.Stdin)

	fmt.Print("Jira URL: ")
	sc.Scan()
	jiraURL := sc.Text()

	fmt.Print("Jira Username: ")
	sc.Scan()
	username := sc.Text()

	fmt.Print("Jira Password: ")
	bytePassword, _ := term.ReadPassword(int(syscall.Stdin))
	password := string(bytePassword)

	fmt.Print("\nJira Board ID : ")
	sc.Scan()
	boardIDstr := sc.Text()
	boardID, err := strconv.Atoi(boardIDstr)
	if err != nil {
		log.Fatal(err)
	}

	tp := jira.BasicAuthTransport{
		Username: username,
		Password: password,
	}

	client, err := jira.NewClient(tp.Client(), jiraURL)
	if err != nil {
		log.Println(err)
		return
	}

	// create a new sprint
	start := time.Now()
	end := start.AddDate(0, 0, 7)
	s := jira.Sprint{
		Name:          "New Sprint",
		StartDate:     &start,
		EndDate:       &end,
		OriginBoardID: boardID,
	}
	futureSprint, _, err := client.Sprint.Create(&s)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("New sprint created ID:", futureSprint.ID)

	// get the sprint currently active
	sprints, _, err := client.Board.GetAllSprintsWithOptions(boardID, &jira.GetAllSprintsOptions{
		State: "active",
	})
	if err != nil {
		log.Fatal(err)
	}
	if len(sprints.Values) != 1 {
		log.Fatal("Retrieved active sprint list has invalid length")
	}
	activeSprint := sprints.Values[0]
	log.Println("Active sprint retrieved ID:", activeSprint.ID)

	// get all non-subtask issues of the active sprint
	issues, _, err := client.Sprint.GetIssuesForSprintWithOptions(activeSprint.ID, &jira.GetIssuesForSprintOptions{
		Jql: "type NOT IN ('Sub-task')",
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Println("All non-subtask issues of the active sprint retrieved length:", len(issues))
	for _, i := range issues {
		log.Printf("\t(%s) %s", i.ID, i.Fields.Summary)
	}

	// complete the active sprint
	completeParam := make(map[string]interface{})
	completeParam["state"] = "closed"
	resp, err := client.Sprint.UpdateSprint(activeSprint.ID, completeParam)
	if err != nil {
		log.Fatal(err)
	}
	resp.Body.Close()
	log.Println("Active sprint completed")

	// move the issues previously under the active sprint to the newly created sprint
	var issueIDs []string
	for _, i := range issues {
		issueIDs = append(issueIDs, i.ID)
	}
	resp, err = client.Sprint.MoveIssuesToSprint(futureSprint.ID, issueIDs)
	if err != nil {
		log.Fatal(err)
	}
	resp.Body.Close()
	log.Println("The issues from the active sprint have been moved to the new sprint")

	// start the newly created sprint
	startParam := make(map[string]interface{})
	startParam["state"] = "active"
	resp, err = client.Sprint.UpdateSprint(futureSprint.ID, startParam)
	if err != nil {
		log.Fatal(err)
	}
	resp.Body.Close()
	log.Println("The new sprint has been started")
}
