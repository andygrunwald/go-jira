package servicedesk

type UserDTO struct {
	AccountID    string       `json:"accountId,omitempty" structs:"accountId,omitempty"`
	Name         string       `json:"name,omitempty" structs:"name,omitempty"`
	Key          string       `json:"key,omitempty" structs:"key,omitempty"`
	EmailAddress string       `json:"emailAddress,omitempty" structs:"emailAddress,omitempty"`
	DisplayName  string       `json:"displayName,omitempty" structs:"displayName,omitempty"`
	Active       bool         `json:"active" structs:"active"`
	TimeZone     string       `json:"timeZone,omitempty" structs:"timeZone,omitempty"`
	Links        *UserLinkDTO `json:"_links,omitempty" structs:"_links,omitempty"`
}
type UserLinkDTO struct {
	Self       string            `json:"self,omitempty" structs:"self,omitempty"`
	JIRARest   string            `json:"jiraRest,omitempty" structs:"jiraRest,omitempty"`
	AvatarUrls map[string]string `json:"avatarUrls,omitempty" structs:"avatarUrls,omitempty"`
}
type RenderedValueDTO struct {
	HTML string `json:"html,omitempty" structs:"html,omitempty"`
}
