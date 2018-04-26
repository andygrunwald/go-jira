package main

import (
	"bufio"
	"fmt"
	"net/http/httputil"
	"os"
	"strings"
	"syscall"

	jira "github.com/andygrunwald/go-jira"
	"golang.org/x/crypto/ssh/terminal"
)

func main() {
	client, err := getAuthenticatedClient()
	if err != nil {
		panic(err)
	}

	_, resp, err := client.Filter.GetList(nil)
	if err != nil {
		panic(err)
	}
	raw, err := httputil.DumpResponse(resp.Response, true)
	if err != nil {
		panic(err)
	}
	fmt.Printf("RESPONSE: %s\n", string(raw))

}

func getAuthenticatedClient() (*jira.Client, error) {
	r := bufio.NewReader(os.Stdin)

	fmt.Print("Jira URL: ")
	jiraURL, _ := r.ReadString('\n')

	fmt.Print("Jira Username: ")
	username, _ := r.ReadString('\n')

	fmt.Print("Jira Password: ")
	bytePassword, _ := terminal.ReadPassword(int(syscall.Stdin))
	password := string(bytePassword)

	tp := jira.BasicAuthTransport{
		Username: strings.TrimSpace(username),
		Password: strings.TrimSpace(password),
	}

	return jira.NewClient(tp.Client(), strings.TrimSpace(jiraURL))
}
