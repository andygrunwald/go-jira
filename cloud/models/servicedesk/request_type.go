package servicedesk

import "github.com/andygrunwald/go-jira/v2/cloud/models"

// RequestType contains RequestType data
type RequestType struct {
	ID            string           `json:"id,omitempty" structs:"id,omitempty"`
	Name          string           `json:"name,omitempty" structs:"name,omitempty"`
	Description   string           `json:"description,omitempty" structs:"description,omitempty"`
	HelpText      string           `json:"helpText,omitempty" structs:"helpText,omitempty"`
	IssueTypeID   string           `json:"issueTypeId,omitempty" structs:"issueTypeId,omitempty"`
	ServiceDeskID string           `json:"serviceDeskId,omitempty" structs:"serviceDeskId,omitempty"`
	GroupIDs      []string         `json:"groupIds,omitempty" structs:"groupIds,omitempty"`
	Icon          *RequestTypeIcon `json:"icon,omitempty" structs:"icon,omitempty"`
	Practice      string           `json:"practice,omitempty" structs:"practice,omitempty"`
	Expands       []string         `json:"_expands,omitempty" structs:"_expands,omitempty"`
	Links         *models.SelfLink `json:"_links,omitempty" structs:"_links,omitempty"`
}

// RequestTypeIcon contains RequestType icon data
type RequestTypeIcon struct {
	ID    string               `json:"id,omitempty" structs:"id,omitempty"`
	Links *RequestTypeIconLink `json:"_links,omitempty" structs:"_links,omitempty"`
}

// RequestTypeIconLink contains RequestTypeIcon link data
type RequestTypeIconLink struct {
	IconURLs map[string]string `json:"iconUrls,omitempty" structs:"iconUrls,omitempty"`
}

// RequestTypeOptions is the query options for listing request types.
type RequestTypeOptions struct {
	SearchQuery string `url:"searchQuery,omitempty" query:"searchQuery"`
	GroupID     int    `url:"groupId,omitempty" query:"groupId"`
}

type CustomerRequestCreateMeta struct {
	RequestTypeFields         []RequestTypeField `json:"requestTypeFields,omitempty" structs:"requestTypeFields,omitempty"`
	CanRaiseOnBehalfOf        bool               `json:"canRaiseOnBehalfOf,omitempty" structs:"canRaiseOnBehalfOf,omitempty"`
	CanAddRequestParticipants bool               `json:"canAddRequestParticipants,omitempty" structs:"canAddRequestParticipants,omitempty"`
}

type RequestTypeField struct {
	FieldID       string                  `json:"fieldId,omitempty" structs:"fieldId,omitempty"`
	Name          string                  `json:"name,omitempty" structs:"name,omitempty"`
	Description   string                  `json:"description,omitempty" structs:"description,omitempty"`
	Required      bool                    `json:"required" structs:"required"`
	DefaultValues []RequestTypeFieldValue `json:"defaultValues,omitempty" structs:"defaultValues,omitempty"`
	ValidValues   []RequestTypeFieldValue `json:"validValues,omitempty" structs:"validValues,omitempty"`
	PresetValues  []string                `json:"presetValues,omitempty" structs:"presetValues,omitempty"`
	JiraSchema    *JiraSchema             `json:"jiraSchema,omitempty" structs:"jiraSchema,omitempty"`
	Visible       bool                    `json:"visible" structs:"visible"`
}

type RequestTypeFieldValue struct {
	Value    string                  `json:"value,omitempty" structs:"value,omitempty"`
	Label    string                  `json:"label,omitempty" structs:"label,omitempty"`
	Children []RequestTypeFieldValue `json:"children,omitempty" structs:"children,omitempty"`
}

type JiraSchema struct {
	Type          string            `json:"type,omitempty" structs:"type,omitempty"`
	Items         string            `json:"items,omitempty" structs:"items,omitempty"`
	System        string            `json:"system,omitempty" structs:"system,omitempty"`
	Custom        string            `json:"custom,omitempty" structs:"custom,omitempty"`
	CustomID      int               `json:"customId,omitempty" structs:"customId,omitempty"`
	Configuration map[string]string `json:"configuration,omitempty" structs:"configuration,omitempty"`
}

type RequestTypeGroup struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
