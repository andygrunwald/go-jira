package model

type SearchUsersQueryOptions struct {
	Query      string `url:"query,omitempty"`
	Username   string `url:"username,omitempty"`
	AccountID  string `url:"accountId,omitempty"`
	StartAt    string `url:"startAt,omitempty"`
	MaxResults int    `url:"maxResults,omitempty"`
	Property   string `url:"property,omitempty"`
}

type SearchUsersQueryResult struct {
	Self        string `json:"self"`
	AccountID   string `json:"accountId"`
	AccountType string `json:"accountType"`
	Name        string `json:"name"`
	DisplayName string `json:"displayName"`
	Active      bool   `json:"active"`
	TimeZone    string `json:"timeZone"`
	Locale      string `json:"locale"`
}
