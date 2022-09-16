package insight

type ObjectSchemaList struct {
	StartAt    int            `json:"startAt" validate:"required"`
	MaxResults int            `json:"maxResults" validate:"required"`
	Total      int            `json:"total" validate:"required"`
	Values     []ObjectSchema `json:"values" validate:"required"`
}

type ObjectSchema struct {
	WorkspaceID     string `json:"workspaceId" validate:"required"`
	GlobalID        string `json:"globalId" validate:"required"`
	ID              string `json:"id" validate:"required"`
	Name            string `json:"name" validate:"required"`
	ObjectSchemaKey string `json:"objectSchemaKey" validate:"required"`
	Description     string `json:"description,omitempty"`
	Status          string `json:"status,omitempty"`
	Created         string `json:"created" validate:"required"`
	Updated         string `json:"updated" validate:"required"`
	ObjectCount     int    `json:"objectCount"`
	ObjectTypeCount int    `json:"objectTypeCount"`
	CanManage       bool   `json:"canManage,omitempty"`
}

type ObjectSchemaAttributeOptions struct {
	OnlyValueEditable bool   `url:"onlyValueEditable,omitempty"`
	Extended          bool   `url:"extended,omitempty"`
	Query             string `url:"query,omitempty"`
}

type CreateOrUpdateObjectSchema struct {
	Name            string `json:"name" validate:"required"`
	ObjectSchemaKey string `json:"objectSchemaKey" validate:"required"`
	Description     string `json:"description,omitempty"`
}
