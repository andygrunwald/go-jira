package model

// FindGroupAndUsersQueryOptions specifies the optional parameters to the Edit issue
type FindGroupAndUsersQueryOptions struct {
	Query                string   `url:"query,omitempty"`
	MaxResults           int      `url:"maxResults,omitempty"`
	ShowAvatar           bool     `url:"showAvatar,omitempty"`
	FieldID              string   `url:"fieldId,omitempty"`
	ProjectID            []string `url:"projectId,omitempty"`
	IssueTypeID          []string `url:"issueTypeId,omitempty"`
	AvatarSize           string   `url:"avatarSize,omitempty"`
	CaseInsensitive      bool     `url:"caseInsensitive,omitempty"`
	ExcludeConnectAddons bool     `url:"excludeConnectAddons,omitempty"`
}

type FoundUsersAndGroups struct {
	Users  FoundUsers  `json:"users"`
	Groups FoundGroups `json:"groups"`
}

type FoundUsers struct {
	Users  []UserPickerUser `json:"users"`
	Total  int              `json:"total"`
	Header string           `json:"header"`
}

type UserPickerUser struct {
	AccountID   string `json:"accountId"`
	HTML        string `json:"html"`
	DisplayName string `json:"displayName"`
	AvatarURL   string `json:"avatarUrl"`
}

type FoundGroups struct {
	Groups []FoundGroup `json:"groups"`
	Total  int          `json:"total"`
	Header string       `json:"header"`
}

type FoundGroup struct {
	GroupID string       `json:"groupId"`
	Name    string       `json:"name"`
	HTML    string       `json:"html"`
	Labels  []GroupLabel `json:"labels"`
}

type GroupLabel struct {
	Text  string `json:"text"`
	Title string `json:"title"`
	Type  string `json:"type"`
}
