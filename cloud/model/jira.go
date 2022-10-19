package model

// PagedDTOT is response of a paged list (PagedDTO) with generic support for values
type PagedDTOT[T any] struct {
	Size       int       `json:"size" structs:"size"`
	Start      int       `json:"start" structs:"start"`
	Limit      int       `json:"limit,omitempty" structs:"limit,omitempty"`
	IsLastPage bool      `json:"isLastPage,omitempty" structs:"isLastPage,omitempty"`
	Values     []T       `json:"values,omitempty" structs:"values,omitempty"`
	Expands    []string  `json:"_expands,omitempty" structs:"_expands,omitempty"`
	Links      *SelfLink `json:"_links,omitempty" structs:"_links,omitempty"`
}

type SelfLink struct {
	Self string `json:"self,omitempty" structs:"self,omitempty"`
}

type DateDTO struct {
	ISO8601     string `json:"iso8601,omitempty" structs:"iso8601,omitempty"`
	JIRA        string `json:"jira,omitempty" structs:"jira,omitempty"`
	Friendly    string `json:"friendly,omitempty" structs:"friendly,omitempty"`
	EpochMillis int64  `json:"epochMillis,omitempty" structs:"epochMillis,omitempty"`
}
