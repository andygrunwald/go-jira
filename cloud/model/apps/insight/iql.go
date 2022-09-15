package insight

type IQLGetObjectsOptions struct {
	IQL                   string `url:"iql,omitempty"`
	Page                  *int   `url:"page,omitempty"`
	ResultPerPage         *int   `url:"resultPerPage,omitempty"`
	IncludeAttributes     *bool  `url:"includeAttributes,omitempty"`
	IncludeAttributesDeep *int   `url:"includeAttributesDeep,omitempty"`
	IncludeTypeAttributes bool   `url:"includeTypeAttributes,omitempty"`
	IncludeExtendedInfo   bool   `url:"includeExtendedInfo,omitempty"`
}

type ObjectListResult struct {
	ObjectEntries        []Object              `json:"objectEntries" validate:"required"`
	ObjectTypeAttributes []ObjectTypeAttribute `json:"objectTypeAttributes"`
	// Deprecated field that shows which object type id the result is for. Not applicable when using IQL
	ObjectTypeID []string `json:"objectTypeId"`
	// Deprecated field should not be used.
	ObjectTypeIsInherited bool `json:"objectTypeIsInherited"`
	// Deprecated field should not be used.
	AbstractObjectType bool `json:"abstractObjectType"`
	TotalFilterCount   int  `json:"totalFilterCount"`
	StartIndex         int  `json:"startIndex"`
	ToIndex            int  `json:"toIndex"`
	PageObjectSize     int  `json:"pageObjectSize"`
	PageNumber         int  `json:"pageNumber"`
	// Deprecated field - The object type attribute id used for sorting
	OrderByTypeAttrID int `json:"orderByTypeAttrId"`
	// Deprecated field - The sort order, used in conjunction with the orderByTypeAttrId
	OrderWay string `json:"orderWay"`
	// Deprecated field - The field is used for basic search
	Filters            map[string]any `json:"filters"`
	IQL                string         `json:"iql" validate:"required"`
	IQLSearchResult    bool           `json:"iqlSearchResult"`
	ConversionPossible bool           `json:"conversionPossible"`
	// Deprecated field should not be used
	MatchedFilterValues map[string]any `json:"matchedFilterValues"`
	// Deprecated field should not be used
	InheritanceTree map[string]any `json:"inheritanceTree"`
}
