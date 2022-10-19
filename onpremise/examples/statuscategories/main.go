package main

import (
	"context"
	"log"

	jira "github.com/andygrunwald/go-jira/v2/onpremise"
)

func main() {
	jiraClient, err := jira.NewClient("https://issues.apache.org/jira/", nil)
	if err != nil {
		panic(err)
	}

	// Showcase of StatusCategory.GetList:
	// Getting all status categories
	categories, resp, err := jiraClient.StatusCategory.GetList(context.TODO())
	if err != nil {
		log.Println(resp.StatusCode)
		panic(err)
	}

	for _, statusCategory := range categories {
		log.Println(statusCategory)
	}

	// Showcase of StatusCategory.Get
	// Getting a single status category
	category, resp, err := jiraClient.StatusCategory.Get(context.TODO(), "1")
	if err != nil {
		log.Println(resp.StatusCode)
		panic(err)
	}

	log.Println(category)
}
