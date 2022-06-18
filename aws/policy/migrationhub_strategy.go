package policy

import "github.com/securisec/cliam/shared"

var MigrationHubStratergyPolicies = map[string]Service{
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
}
