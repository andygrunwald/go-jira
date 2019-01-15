package jira

type JiraError struct {
	Message []string `json:"errorMessages"`
	Errors map[string]string `json:"errors"`
	Status int `json:"status"`
}