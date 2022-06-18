package policy

import "github.com/securisec/cliam/shared"

var RobomakerPolicies = map[string]Service{
	"CreateWorldTemplate": {
		Method:        "POST",
		ServiceSuffix: "createWorldTemplate",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		},
		Permission: "CreateWorldTemplate",
	},
	"GetWorldTemplateBody": {
		Method:        "POST",
		ServiceSuffix: "getWorldTemplateBody",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		},
		Permission: "GetWorldTemplateBody",
	},
	"ListDeploymentJobs": {
		Method:        "POST",
		ServiceSuffix: "listDeploymentJobs",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		},
		Permission: "ListDeploymentJobs",
	},
	"ListFleets": {
		Method:        "POST",
		ServiceSuffix: "listFleets",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		},
		Permission: "ListFleets",
	},
	"ListRobotApplications": {
		Method:        "POST",
		ServiceSuffix: "listRobotApplications",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		},
		Permission: "ListRobotApplications",
	},
	"ListRobots": {
		Method:        "POST",
		ServiceSuffix: "listRobots",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		},
		Permission: "ListRobots",
	},
	"ListSimulationApplications": {
		Method:        "POST",
		ServiceSuffix: "listSimulationApplications",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		},
		Permission: "ListSimulationApplications",
	},
	"ListSimulationJobBatches": {
		Method:        "POST",
		ServiceSuffix: "listSimulationJobBatches",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		},
		Permission: "ListSimulationJobBatches",
	},
	"ListSimulationJobs": {
		Method:        "POST",
		ServiceSuffix: "listSimulationJobs",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		},
		Permission: "ListSimulationJobs",
	},
	"ListWorldExportJobs": {
		Method:        "POST",
		ServiceSuffix: "listWorldExportJobs",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		},
		Permission: "ListWorldExportJobs",
	},
	"ListWorldGenerationJobs": {
		Method:        "POST",
		ServiceSuffix: "listWorldGenerationJobs",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		},
		Permission: "ListWorldGenerationJobs",
	},
	"ListWorldTemplates": {
		Method:        "POST",
		ServiceSuffix: "listWorldTemplates",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		},
		Permission: "ListWorldTemplates",
	},
	"ListWorlds": {
		Method:        "POST",
		ServiceSuffix: "listWorlds",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		},
		Permission: "ListWorlds",
	},

	// extra
	"DescribeDeploymentJob": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeDeploymentJob",
			"Version": "2018-06-29",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "DescribeDeploymentJob",
		IsExtra:                true,
		ExtraComponentBodyKey:  "job",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "job",
	},
	"DescribeFleet": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeFleet",
			"Version": "2018-06-29",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "DescribeFleet",
		IsExtra:                true,
		ExtraComponentBodyKey:  "fleet",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "fleet",
	},
	"DescribeRobot": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeRobot",
			"Version": "2018-06-29",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "DescribeRobot",
		IsExtra:                true,
		ExtraComponentBodyKey:  "robot",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "robot",
	},
	"DescribeRobotApplication": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeRobotApplication",
			"Version": "2018-06-29",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "DescribeRobotApplication",
		IsExtra:                true,
		ExtraComponentBodyKey:  "application",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "application",
	},
	"DescribeSimulationApplication": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeSimulationApplication",
			"Version": "2018-06-29",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "DescribeSimulationApplication",
		IsExtra:                true,
		ExtraComponentBodyKey:  "application",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "application",
	},
	"DescribeSimulationJob": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeSimulationJob",
			"Version": "2018-06-29",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "DescribeSimulationJob",
		IsExtra:                true,
		ExtraComponentBodyKey:  "job",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "job",
	},
	"DescribeSimulationJobBatch": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeSimulationJobBatch",
			"Version": "2018-06-29",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "DescribeSimulationJobBatch",
		IsExtra:                true,
		ExtraComponentBodyKey:  "batch",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "batch",
	},
	"DescribeWorld": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeWorld",
			"Version": "2018-06-29",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "DescribeWorld",
		IsExtra:                true,
		ExtraComponentBodyKey:  "world",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "world",
	},
	"DescribeWorldExportJob": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeWorldExportJob",
			"Version": "2018-06-29",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "DescribeWorldExportJob",
		IsExtra:                true,
		ExtraComponentBodyKey:  "job",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "job",
	},
	"DescribeWorldGenerationJob": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeWorldGenerationJob",
			"Version": "2018-06-29",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "DescribeWorldGenerationJob",
		IsExtra:                true,
		ExtraComponentBodyKey:  "job",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "job",
	},
	"DescribeWorldTemplate": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeWorldTemplate",
			"Version": "2018-06-29",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "DescribeWorldTemplate",
		IsExtra:                true,
		ExtraComponentBodyKey:  "template",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "template",
	},
	"ListTagsForResource": {
		ServiceSuffix:          "/tags/{{.resource_arn}}",
		Permission:             "ListTagsForResource",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "resource_arn",
	},
}
