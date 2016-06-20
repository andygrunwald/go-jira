package jira

import (
	"fmt"
)

// TransitionService handles transitions for JIRA issue.
type TransitionService struct {
	client *Client
}

// Wrapper struct for search result
type transitionResult struct {
	Transitions []Transition `json:transitions`
}

// Transition represents an issue transition in JIRA
type Transition struct {
	ID     string                     `json:"id"`
	Name   string                     `json:"name"`
	Fields map[string]TransitionField `json:"fields"`
}

type TransitionField struct {
	Required bool `json:"required"`
}

// CreatePayload is used for creating new issue transitions
type CreateTransitionPayload struct {
	Transition TransitionPayload `json:"transition"`
}

type TransitionPayload struct {
	ID string `json:"id"`
}

// GetList gets transitions available for given issue
//
// JIRA API docs: https://docs.atlassian.com/jira/REST/latest/#api/2/issue-getTransitions
func (s *TransitionService) GetList(id string) ([]Transition, *Response, error) {
	apiEndpoint := fmt.Sprintf("rest/api/2/issue/%s/transitions?expand=transitions.fields", id)
	req, err := s.client.NewRequest("GET", apiEndpoint, nil)
	if err != nil {
		return nil, nil, err
	}

	result := new(transitionResult)
	resp, err := s.client.Do(req, result)
	return result.Transitions, resp, err
}

// Basic transition creation. This simply creates transition with given ID for issue
// with given ID. It doesn't yet support anything else.
//
// JIRA API docs: https://docs.atlassian.com/jira/REST/latest/#api/2/issue-doTransition
func (s *TransitionService) Create(ticketID, transitionID string) (*Response, error) {
	apiEndpoint := fmt.Sprintf("rest/api/2/issue/%s/transitions", ticketID)

	payload := CreateTransitionPayload{
		Transition: TransitionPayload{
			ID: transitionID,
		},
	}
	req, err := s.client.NewRequest("POST", apiEndpoint, payload)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req, nil)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
