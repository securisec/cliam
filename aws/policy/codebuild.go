package policy

import "github.com/securisec/cliam/shared"

var CodebuildPolicies = []Service{
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "CodeBuild_20161006.ListBuildBatches",
		},
		Permission: "ListBuildBatches",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "CodeBuild_20161006.ListBuildBatchesForProject",
		},
		Permission: "ListBuildBatchesForProject",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "CodeBuild_20161006.ListBuilds",
		},
		Permission: "ListBuilds",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "CodeBuild_20161006.ListCuratedEnvironmentImages",
		},
		Permission: "ListCuratedEnvironmentImages",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "CodeBuild_20161006.ListProjects",
		},
		Permission: "ListProjects",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "CodeBuild_20161006.ListReportGroups",
		},
		Permission: "ListReportGroups",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "CodeBuild_20161006.ListReports",
		},
		Permission: "ListReports",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "CodeBuild_20161006.ListSharedProjects",
		},
		Permission: "ListSharedProjects",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "CodeBuild_20161006.ListSourceCredentials",
		},
		Permission: "ListSourceCredentials",
	},

	// extra
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "CodeBuild_20161006.DescribeCodeCoverages",
		},
		Permission:             "DescribeCodeCoverages",
		IsExtra:                true,
		ExtraComponentBodyKey:  "reportArn",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "report_arn",
	},
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "CodeBuild_20161006.DescribeTestCases",
		},
		Permission:             "DescribeTestCases",
		IsExtra:                true,
		ExtraComponentBodyKey:  "reportArn",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "report_arn",
	},
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "CodeBuild_20161006.GetResourcePolicy",
		},
		Permission:             "GetResourcePolicy",
		IsExtra:                true,
		ExtraComponentBodyKey:  "resourceArn",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "resource_arn",
	},
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "CodeBuild_20161006.ListBuildsForProject",
		},
		Permission:             "ListBuildsForProject",
		IsExtra:                true,
		ExtraComponentBodyKey:  "projectName",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "project_name",
	},
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "CodeBuild_20161006.ListReportsForReportGroup",
		},
		Permission:             "ListReportsForReportGroup",
		IsExtra:                true,
		ExtraComponentBodyKey:  "reportGroupArn",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "report_group_arn",
	},
}
