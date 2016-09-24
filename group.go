package jira

import (
	"fmt"
)


type GroupService struct {
	client *Client
}

type GroupMembers struct {
	StartAt    int             `json:"startAt"`
	MaxResults int             `json:"maxResults"`
	Total      int             `json:"total"`
	Members   []Member `json:"values"`
}


type Member struct {
	Self      string `json:"self,omitempty"`
	Name        string `json:"name,omitempty"`
	Key  string `json:"key,omitempty"`
	EmailAddress    string  `json:"emailAddress,omitempty"`
	DisplayName   string `json:"displayName,omitempty"`
	Active      bool    `json:"active,omitempty"`
	TimeZone  string `json:"timeZone,omitempty"`
}



func (s *GroupService) Get(groupName string) ([]Member, *Response, error) {
	apiEndpoint := fmt.Sprintf("rest/api/2/group/member?groupname=%s", groupName)
	req, err := s.client.NewRequest("GET", apiEndpoint, nil)
	if err != nil {
		return nil, nil, err
	}

	group := new(GroupMembers)
	resp, err := s.client.Do(req, group)
	if err != nil {
		return nil, resp, err
	}

	return group.Members, resp, nil
}

