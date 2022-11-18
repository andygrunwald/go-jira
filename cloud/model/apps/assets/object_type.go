package assets

type PutObjectType struct {
	Name               string `json:"name" validate:"required"`
	Description        string `json:"description"`
	IconID             string `json:"iconId" validate:"required"`
	ObjectSchemaID     string `json:"objectSchemaId" validate:"required"`
	ParentObjectTypeID string `json:"parentObjectTypeId"`
	Inherited          bool   `json:"inherited"`
	AbstractObjectType bool   `json:"abstractObjectType"`
}

type ObjectTypeAttributeOptions struct {
	OnlyValueEditable       bool   `url:"onlyValueEditable,omitempty"`
	OrderByName             bool   `url:"orderByName,omitempty"`
	Query                   string `url:"query,omitempty"`
	IncludeValuesExist      bool   `url:"includeValuesExist,omitempty"`
	ExcludeParentAttributes bool   `url:"excludeParentAttributes,omitempty"`
	IncludeChildren         bool   `url:"includeChildren,omitempty"`
	OrderByRequired         bool   `url:"orderByRequired,omitempty"`
}

type PostObjectTypePosition struct {
	ToObjectTypeID string `json:"toObjectTypeId,omitempty"`
	Position       int    `json:"position"`
}

type CreateObjectType struct {
	Name               string `json:"name" validate:"required"`
	Description        string `json:"description"`
	IconID             string `json:"iconId" validate:"required"`
	ObjectSchemaID     string `json:"objectSchemaId" validate:"required"`
	ParentObjectTypeID string `json:"parentObjectTypeId"`
	Inherited          bool   `json:"inherited"`
	AbstractObjectType bool   `json:"abstractObjectType"`
}
