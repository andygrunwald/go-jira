package cloud

import "errors"

var ErrValidation = errors.New("the system cannot fulfill the request due to validation errors")
var ErrUnauthorized = errors.New("client must be authenticated to access this resource")
var ErrNotFound = errors.New("target resource do not exist")
var ErrUnknown = errors.New("internal server error")

const insightsURL = "https://api.atlassian.com"

// InsightsService handles Insights App for the Jira instance / API.
type InsightsService service
