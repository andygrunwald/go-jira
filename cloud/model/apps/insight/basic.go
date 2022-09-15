package insight

type Avatar struct {
	WorkspaceID string `json:"workspaceId" validate:"required"`
	GlobalID    string `json:"globalId" validate:"required"`
	ID          string `json:"id" validate:"required"`
	URL16       string `json:"url16" validate:"required"`
	URL48       string `json:"url48" validate:"required"`
	URL72       string `json:"url72" validate:"required"`
	URL144      string `json:"url144" validate:"required"`
	URL288      string `json:"url288" validate:"required"`
	ObjectID    string `json:"objectId" validate:"required"`
}

type User struct {
	AvatarURL       string `json:"avatarUrl,omitempty"`
	DisplayName     string `json:"displayName,omitempty"`
	Name            string `json:"name,omitempty"`
	Key             string `json:"key,omitempty"`
	EmailAddress    string `json:"emailAddress,omitempty"`
	HTML            string `json:"html,omitempty"`
	RenderedLink    string `json:"renderedLink,omitempty"`
	IsDeleted       bool   `json:"isDeleted,omitempty"`
	LastSeenVersion string `json:"lastSeenVersion,omitempty"`
	Self            string `json:"self,omitempty"`
}

type Group struct {
	AvatarURL string `json:"avatarUrl,omitempty"`
	Name      string `json:"name,omitempty"`
}

type StatusCategory int

const (
	CategoryActive   StatusCategory = 1
	CategoryInactive StatusCategory = 0
	CategoryPending  StatusCategory = 2
)

type Status struct {
	ID             string         `json:"id" validate:"required"`
	Name           string         `json:"name" validate:"required"`
	Description    string         `json:"description,omitempty"`
	Category       StatusCategory `json:"category"`
	ObjectSchemaID string         `json:"objectSchemaId,omitempty"`
}
