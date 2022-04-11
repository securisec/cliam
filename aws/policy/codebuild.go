package policy

import "github.com/securisec/enumerate/shared"

var CodebuildPolicies = []Service{
	{
		Method:   "POST",
		JsonData: `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
			aws_X_AMZ_TARGET:           "CodeBuild_20161006.ListBuildBatches",
		},
		Permission: "ListBuildBatches",
	},
	{
		Method:   "POST",
		JsonData: `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
			aws_X_AMZ_TARGET:           "CodeBuild_20161006.ListBuildBatchesForProject",
		},
		Permission: "ListBuildBatchesForProject",
	},
	{
		Method:   "POST",
		JsonData: `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
			aws_X_AMZ_TARGET:           "CodeBuild_20161006.ListBuilds",
		},
		Permission: "ListBuilds",
	},
	{
		Method:   "POST",
		JsonData: `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
			aws_X_AMZ_TARGET:           "CodeBuild_20161006.ListCuratedEnvironmentImages",
		},
		Permission: "ListCuratedEnvironmentImages",
	},
	{
		Method:   "POST",
		JsonData: `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
			aws_X_AMZ_TARGET:           "CodeBuild_20161006.ListProjects",
		},
		Permission: "ListProjects",
	},
	{
		Method:   "POST",
		JsonData: `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
			aws_X_AMZ_TARGET:           "CodeBuild_20161006.ListReportGroups",
		},
		Permission: "ListReportGroups",
	},
	{
		Method:   "POST",
		JsonData: `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
			aws_X_AMZ_TARGET:           "CodeBuild_20161006.ListReports",
		},
		Permission: "ListReports",
	},
	{
		Method:   "POST",
		JsonData: `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
			aws_X_AMZ_TARGET:           "CodeBuild_20161006.ListSharedProjects",
		},
		Permission: "ListSharedProjects",
	},
	{
		Method:   "POST",
		JsonData: `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
			aws_X_AMZ_TARGET:           "CodeBuild_20161006.ListSourceCredentials",
		},
		Permission: "ListSourceCredentials",
	},
}
