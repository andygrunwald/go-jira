package jira

// StatusCategory represents the category a status belongs to.
// Those categories can be user defined in every JIRA instance.
type StatusCategory struct {
	Self      string `json:"self" structs:"self"`
	ID        int    `json:"id" structs:"id"`
	Name      string `json:"name" structs:"name"`
	Key       string `json:"key" structs:"key"`
	ColorName string `json:"colorName" structs:"colorName"`
}

// These constants are the keys of the default JIRA status categories
const (
	StatusCategoryComplete   = "done"
	StatusCategoryInProgress = "indeterminate"
	StatusCategoryToDo       = "new"
	StatusCategoryUndefined  = "undefined"
)
