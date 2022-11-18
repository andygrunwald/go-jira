package assets

type ObjectCreate struct {
	ObjectTypeID string              `json:"objectTypeId" validate:"required"`
	Attributes   []ObjectAttributeIn `json:"attributes" validate:"required,dive"`
	HasAvatar    bool                `json:"hasAvatar,omitempty"`
	AvatarUUID   string              `json:"avatarUUID,omitempty"`
}

type ObjectAttributeIn struct {
	ObjectTypeAttributeID string                   `json:"objectTypeAttributeId" validate:"required"`
	ObjectAttributeValues []ObjectAttributeValueIn `json:"objectAttributeValues" validate:"required,dive"`
}

type ObjectAttributeValueIn struct {
	Value string `json:"value" validate:"required"`
}
