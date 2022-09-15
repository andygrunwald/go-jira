package cloud

import "errors"

var ErrValidation = errors.New("the system cannot fulfill the request due to validation errors")
var ErrUnauthorized = errors.New("client must be authenticated to access this resource")
var ErrNotFound = errors.New("target resource do not exist")
var ErrUnknown = errors.New("internal server error")

// needs to be var, so we can change it for testing
var insightURL = "https://api.atlassian.com"

// InsightService handles Insight App for the Jira instance / API.
type InsightService struct {
	common service

	IQL          *InsightIQLService
	Object       *InsightObjectService
	ObjectSchema *InsightObjectSchemaService
	ObjectType   *InsightObjectTypeService
}
