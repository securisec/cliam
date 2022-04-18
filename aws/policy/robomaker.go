package policy

import "github.com/securisec/cliam/shared"

var RobomakerPolicies = []Service{
	{
		Method:        "POST",
		ServiceSuffix: "createWorldTemplate",
		JsonData:      `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
		},
		Permission: "CreateWorldTemplate",
	},
	{
		Method:        "POST",
		ServiceSuffix: "getWorldTemplateBody",
		JsonData:      `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
		},
		Permission: "GetWorldTemplateBody",
	},
	{
		Method:        "POST",
		ServiceSuffix: "listDeploymentJobs",
		JsonData:      `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
		},
		Permission: "ListDeploymentJobs",
	},
	{
		Method:        "POST",
		ServiceSuffix: "listFleets",
		JsonData:      `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
		},
		Permission: "ListFleets",
	},
	{
		Method:        "POST",
		ServiceSuffix: "listRobotApplications",
		JsonData:      `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
		},
		Permission: "ListRobotApplications",
	},
	{
		Method:        "POST",
		ServiceSuffix: "listRobots",
		JsonData:      `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
		},
		Permission: "ListRobots",
	},
	{
		Method:        "POST",
		ServiceSuffix: "listSimulationApplications",
		JsonData:      `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
		},
		Permission: "ListSimulationApplications",
	},
	{
		Method:        "POST",
		ServiceSuffix: "listSimulationJobBatches",
		JsonData:      `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
		},
		Permission: "ListSimulationJobBatches",
	},
	{
		Method:        "POST",
		ServiceSuffix: "listSimulationJobs",
		JsonData:      `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
		},
		Permission: "ListSimulationJobs",
	},
	{
		Method:        "POST",
		ServiceSuffix: "listWorldExportJobs",
		JsonData:      `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
		},
		Permission: "ListWorldExportJobs",
	},
	{
		Method:        "POST",
		ServiceSuffix: "listWorldGenerationJobs",
		JsonData:      `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
		},
		Permission: "ListWorldGenerationJobs",
	},
	{
		Method:        "POST",
		ServiceSuffix: "listWorldTemplates",
		JsonData:      `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
		},
		Permission: "ListWorldTemplates",
	},
	{
		Method:        "POST",
		ServiceSuffix: "listWorlds",
		JsonData:      `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
		},
		Permission: "ListWorlds",
	},
}
