package policy

import "github.com/securisec/cliam/shared"

var MigrationHubStratergyPolicies = []Service{
	{
		Method:        "GET",
		ServiceSuffix: "get-portfolio-preferences",
		Permission:    "GetPortfolioPreferences",
	},
	{
		Method:        "GET",
		ServiceSuffix: "get-portfolio-summary",
		Permission:    "GetPortfolioSummary",
	},
	{
		Method:        "POST",
		ServiceSuffix: "list-applicationcomponents",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
		},
		Permission: "ListApplicationComponents",
	},
	{
		Method:        "GET",
		ServiceSuffix: "list-collectors",
		Permission:    "ListCollectors",
	},
	{
		Method:        "GET",
		ServiceSuffix: "list-import-file-task",
		Permission:    "ListImportFileTask",
	},
	{
		Method:        "POST",
		ServiceSuffix: "list-servers",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
		},
		Permission: "ListServers",
	},
	{
		Method:        "POST",
		ServiceSuffix: "put-portfolio-preferences",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
		},
		Permission: "PutPortfolioPreferences",
	},
	{
		Method:        "POST",
		ServiceSuffix: "start-assessment",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
		},
		Permission: "StartAssessment",
	},
	{
		Method:        "POST",
		ServiceSuffix: "start-recommendation-report-generation",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
		},
		Permission: "StartRecommendationReportGeneration",
	},
}
