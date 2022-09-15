package cloud

// needs to be var, so we can change it for testing
var insightURL = "https://api.atlassian.com"

// InsightService handles Insight App for the Jira instance / API.
type InsightService struct {
	common service

	Icon         *InsightIconService
	IQL          *InsightIQLService
	Object       *InsightObjectService
	ObjectSchema *InsightObjectSchemaService
	ObjectType   *InsightObjectTypeService
}
