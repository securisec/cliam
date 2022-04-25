package policy

import "github.com/securisec/cliam/shared"

var RobomakerPolicies = []Service{
	{
		Method:        "POST",
		ServiceSuffix: "createWorldTemplate",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		},
		Permission: "CreateWorldTemplate",
	},
	{
		Method:        "POST",
		ServiceSuffix: "getWorldTemplateBody",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		},
		Permission: "GetWorldTemplateBody",
	},
	{
		Method:        "POST",
		ServiceSuffix: "listDeploymentJobs",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		},
		Permission: "ListDeploymentJobs",
	},
	{
		Method:        "POST",
		ServiceSuffix: "listFleets",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		},
		Permission: "ListFleets",
	},
	{
		Method:        "POST",
		ServiceSuffix: "listRobotApplications",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		},
		Permission: "ListRobotApplications",
	},
	{
		Method:        "POST",
		ServiceSuffix: "listRobots",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		},
		Permission: "ListRobots",
	},
	{
		Method:        "POST",
		ServiceSuffix: "listSimulationApplications",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		},
		Permission: "ListSimulationApplications",
	},
	{
		Method:        "POST",
		ServiceSuffix: "listSimulationJobBatches",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		},
		Permission: "ListSimulationJobBatches",
	},
	{
		Method:        "POST",
		ServiceSuffix: "listSimulationJobs",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		},
		Permission: "ListSimulationJobs",
	},
	{
		Method:        "POST",
		ServiceSuffix: "listWorldExportJobs",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		},
		Permission: "ListWorldExportJobs",
	},
	{
		Method:        "POST",
		ServiceSuffix: "listWorldGenerationJobs",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		},
		Permission: "ListWorldGenerationJobs",
	},
	{
		Method:        "POST",
		ServiceSuffix: "listWorldTemplates",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		},
		Permission: "ListWorldTemplates",
	},
	{
		Method:        "POST",
		ServiceSuffix: "listWorlds",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		},
		Permission: "ListWorlds",
	},
}
