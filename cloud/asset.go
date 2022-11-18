package cloud

// needs to be var, so we can change it for testing
var assetURL = "https://api.atlassian.com"

// AssetsService handles Insight App for the Jira instance / API.
type AssetsService struct {
	common service

	Icon         *AssetsIconService
	IQL          *AssetsIQLService
	Object       *AssetsObjectService
	ObjectSchema *AssetsObjectSchemaService
	ObjectType   *AssetsObjectTypeService
}
