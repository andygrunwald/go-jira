package insight

import "encoding/json"

type ObjectTypeAttributeType int

const (
	TypeDefault         ObjectTypeAttributeType = 0
	TypeObjectReference ObjectTypeAttributeType = 1
	TypeUser            ObjectTypeAttributeType = 2
	TypeGroup           ObjectTypeAttributeType = 4
	TypeStatus          ObjectTypeAttributeType = 7
)

type ObjectTypeAttributeDefaultType int

const (
	DefaultTypeNone      ObjectTypeAttributeDefaultType = -1
	DefaultTypeText      ObjectTypeAttributeDefaultType = 0
	DefaultTypeInteger   ObjectTypeAttributeDefaultType = 1
	DefaultTypeBoolean   ObjectTypeAttributeDefaultType = 2
	DefaultTypeDouble    ObjectTypeAttributeDefaultType = 3
	DefaultTypeDate      ObjectTypeAttributeDefaultType = 4
	DefaultTypeTime      ObjectTypeAttributeDefaultType = 5
	DefaultTypeDateTime  ObjectTypeAttributeDefaultType = 6
	DefaultTypeURL       ObjectTypeAttributeDefaultType = 7
	DefaultTypeEmail     ObjectTypeAttributeDefaultType = 8
	DefaultTypeTextarea  ObjectTypeAttributeDefaultType = 9
	DefaultTypeSelect    ObjectTypeAttributeDefaultType = 10
	DefaultTypeIPAddress ObjectTypeAttributeDefaultType = 11
)

type Object struct {
	WorkspaceID string            `json:"workspaceId" validate:"required"`
	GlobalID    string            `json:"globalId" validate:"required"`
	ID          string            `json:"id" validate:"required"`
	Label       string            `json:"label" validate:"required"`
	ObjectKey   string            `json:"objectKey" validate:"required"`
	Avatar      Avatar            `json:"avatar" validate:"required"`
	ObjectType  ObjectType        `json:"objectType" validate:"required"`
	Created     string            `json:"created" validate:"required"`
	Updated     string            `json:"updated" validate:"required"`
	HasAvatar   bool              `json:"hasAvatar"`
	Timestamp   int64             `json:"timestamp" validate:"required"`
	Attributes  []ObjectAttribute `json:"attributes" validate:"dive"`
	Links       json.RawMessage   `json:"_links " validate:"required"`
}

type ObjectType struct {
	WorkspaceID               string `json:"workspaceId" validate:"required"`
	GlobalID                  string `json:"globalId" validate:"required"`
	ID                        string `json:"id" validate:"required"`
	Name                      string `json:"name" validate:"required"`
	Description               string `json:"description" validate:"required"`
	Icon                      Icon   `json:"icon" validate:"required"`
	Position                  int    `json:"position"`
	Created                   string `json:"created" validate:"required"`
	Updated                   string `json:"updated" validate:"required"`
	ObjectCount               int    `json:"objectCount" validate:"required"`
	ParentObjectTypeID        int    `json:"parentObjectTypeId" validate:"required"`
	ObjectSchemaID            string `json:"objectSchemaId" validate:"required"`
	Inherited                 bool   `json:"inherited"`
	AbstractObjectType        bool   `json:"abstractObjectType"`
	ParentObjectTypeInherited bool   `json:"parentObjectTypeInherited"`
}

type ObjectAttribute struct {
	WorkspaceID           string                 `json:"workspaceId" validate:"required"`
	GlobalID              string                 `json:"globalId" validate:"required"`
	ID                    string                 `json:"id" validate:"required"`
	ObjectTypeAttribute   ObjectTypeAttribute    `json:"objectTypeAttribute" validate:"required"`
	ObjectTypeAttributeID string                 `json:"objectTypeAttributeId" validate:"required"`
	ObjectAttributeValues []ObjectAttributeValue `json:"objectAttributeValues" validate:"required,dive"`
}

type ObjectTypeAttribute struct {
	WorkspaceID             string                  `json:"workspaceId" validate:"required"`
	GlobalID                string                  `json:"globalId" validate:"required"`
	ID                      string                  `json:"id" validate:"required"`
	ObjectType              ObjectType              `json:"objectType" validate:"required"`
	Name                    string                  `json:"name,omitempty"`
	Label                   bool                    `json:"label,omitempty"`
	Type                    ObjectTypeAttributeType `json:"type"`
	Description             string                  `json:"description,omitempty"`
	DefaultType             DefaultType             `json:"defaultType"`
	TypeValue               string                  `json:"typeValue,omitempty"`
	TypeValueMulti          []string                `json:"typeValueMulti,omitempty"`
	AdditionalValue         string                  `json:"additionalValue,omitempty"`
	ReferenceType           *ReferenceType          `json:"referenceType,omitempty"`
	Editable                bool                    `json:"editable,omitempty"`
	System                  bool                    `json:"system,omitempty"`
	Indexed                 bool                    `json:"indexed,omitempty"`
	Sortable                bool                    `json:"sortable,omitempty"`
	Summable                bool                    `json:"summable,omitempty"`
	MinimumCardinality      int                     `json:"minimumCardinality,omitempty"`
	MaximumCardinality      int                     `json:"maximumCardinality,omitempty"`
	Suffix                  string                  `json:"suffix,omitempty"`
	Removable               bool                    `json:"removable,omitempty"`
	ObjectAttributeExists   bool                    `json:"objectAttributeExists,omitempty"`
	Hidden                  bool                    `json:"hidden,omitempty"`
	IncludeChildObjectTypes bool                    `json:"includeChildObjectTypes,omitempty"`
	UniqueAttribute         bool                    `json:"uniqueAttribute,omitempty"`
	RegexValidation         string                  `json:"regexValidation,omitempty"`
	IQL                     string                  `json:"iql,omitempty"`
	Options                 string                  `json:"options,omitempty"`
	Position                int                     `json:"position" validate:"required"`
}

type ObjectAttributeValue struct {
	Value            string  `json:"value"`
	DisplayValue     string  `json:"displayValue" validate:"required"`
	SearchValue      string  `json:"searchValue"`
	ReferencedObject *Object `json:"referencedObject"`
	User             *User   `json:"user"`
	Group            *Group  `json:"group"`
	Status           *Status `json:"status"`
	AdditionalValue  string  `json:"additionalValue"`
}

type ReferenceType struct {
	WorkspaceID    string `json:"workspaceId" validate:"required"`
	GlobalID       string `json:"globalId" validate:"required"`
	ID             string `json:"id,omitempty"`
	Name           string `json:"name" validate:"required"`
	Description    string `json:"description,omitempty"`
	Color          string `json:"color,omitempty"`
	URL16          string `json:"url16,omitempty"`
	Removable      bool   `json:"removable,omitempty"`
	ObjectSchemaID string `json:"objectSchemaId,omitempty"`
}

type DefaultType struct {
	ID   ObjectTypeAttributeDefaultType `json:"id"`
	Name string                         `json:"name"`
}

type PutObject struct {
	ObjectTypeID string              `json:"objectTypeId" validate:"required"`
	Attributes   []ObjectAttributeIn `json:"attributes" validate:"dive"`
	HasAvatar    bool                `json:"hasAvatar,omitempty"`
	AvatarUUID   string              `json:"avatarUUID,omitempty"`
}

type ObjectReferenceTypeInfo struct {
	ReferenceTypes            []ReferenceType `json:"referenceTypes" validate:"dive"`
	ObjectType                ObjectType      `json:"objectType"`
	NumberOfReferencedObjects int             `json:"numberOfReferencedObjects" validate:"required"`
	OpenIssuesExists          bool            `json:"openIssuesExists"`
}

type GenericList[T any] struct {
	StartAt    int `json:"startAt" validate:"required"`
	MaxResults int `json:"maxResults" validate:"required"`
	Total      int `json:"total" validate:"required"`
	Values     []T `json:"values" validate:"required"`
}

type ObjectHistoryOptions struct {
	AscendingOrder bool `url:"asc,omitempty"`
	Abbreviate     bool `url:"abbreviate,omitempty"`
}

type ObjectHistory struct {
	Actor             User   `json:"actor" validate:"required"`
	ID                string `json:"id" validate:"required"`
	AffectedAttribute string `json:"affectedAttribute"`
	OldValue          string `json:"oldValue"`
	NewValue          string `json:"newValue"`
	Type              string `json:"type" validate:"required"`
	Created           string `json:"created" validate:"required"`
	ObjectID          string `json:"objectId" validate:"required"`
}
