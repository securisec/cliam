package policy

import "github.com/securisec/cliam/shared"

// MigrationHubStratergyPolicies policy
var MigrationHubStratergyPolicies = map[string]Service{
	"GetLatestAssessmentId": {
		Method:        "GET",
		ServiceSuffix: "get-latest-assessment-id",
		Permission:    "GetLatestAssessmentId",
	},
	"GetPortfolioPreferences": {
		Method:        "GET",
		ServiceSuffix: "get-portfolio-preferences",
		Permission:    "GetPortfolioPreferences",
	},
	"GetPortfolioSummary": {
		Method:        "GET",
		ServiceSuffix: "get-portfolio-summary",
		Permission:    "GetPortfolioSummary",
	},
	"ListApplicationComponents": {
		Method:        "POST",
		ServiceSuffix: "list-applicationcomponents",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		},
		Permission: "ListApplicationComponents",
	},
	"ListCollectors": {
		Method:        "GET",
		ServiceSuffix: "list-collectors",
		Permission:    "ListCollectors",
	},
	"ListImportFileTask": {
		Method:        "GET",
		ServiceSuffix: "list-import-file-task",
		Permission:    "ListImportFileTask",
	},
	"ListServers": {
		Method:        "POST",
		ServiceSuffix: "list-servers",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		},
		Permission: "ListServers",
	},
	"PutPortfolioPreferences": {
		Method:        "POST",
		ServiceSuffix: "put-portfolio-preferences",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		},
		Permission: "PutPortfolioPreferences",
	},
	"StartAssessment": {
		Method:        "POST",
		ServiceSuffix: "start-assessment",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		},
		Permission: "StartAssessment",
	},
	"StartRecommendationReportGeneration": {
		Method:        "POST",
		ServiceSuffix: "start-recommendation-report-generation",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		},
		Permission: "StartRecommendationReportGeneration",
	},

	// extra
	"GetApplicationComponentDetails": {
		ServiceSuffix:          "/get-applicationcomponent-details/{{.application_component_id}}",
		Permission:             "GetApplicationComponentDetails",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "application_component_id",
	},
	"GetApplicationComponentStrategies": {
		ServiceSuffix:          "/get-applicationcomponent-strategies/{{.application_component_id}}",
		Permission:             "GetApplicationComponentStrategies",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "application_component_id",
	},
	"GetAssessment": {
		ServiceSuffix:          "/get-assessment/{{.id}}",
		Permission:             "GetAssessment",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "id",
	},
	"GetImportFileTask": {
		ServiceSuffix:          "/get-import-file-task/{{.id}}",
		Permission:             "GetImportFileTask",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "id",
	},
	"GetRecommendationReportDetails": {
		ServiceSuffix:          "/get-recommendation-report-details/{{.id}}",
		Permission:             "GetRecommendationReportDetails",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "id",
	},
	"GetServerDetails": {
		ServiceSuffix:          "/get-server-details/{{.server_id}}",
		Permission:             "GetServerDetails",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "server_id",
	},
	"GetServerStrategies": {
		ServiceSuffix:          "/get-server-strategies/{{.server_id}}",
		Permission:             "GetServerStrategies",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "server_id",
	},
}
