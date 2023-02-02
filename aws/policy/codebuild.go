package policy

import "github.com/securisec/cliam/shared"

// CodeBuildPolicies policy
var CodeBuildPolicies = map[string]Service{
	"ListBuildBatches": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "CodeBuild_20161006.ListBuildBatches",
		},
		Permission: "ListBuildBatches",
	},
	"ListBuildBatchesForProject": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "CodeBuild_20161006.ListBuildBatchesForProject",
		},
		Permission: "ListBuildBatchesForProject",
	},
	"ListBuilds": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "CodeBuild_20161006.ListBuilds",
		},
		Permission: "ListBuilds",
	},
	"ListCuratedEnvironmentImages": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "CodeBuild_20161006.ListCuratedEnvironmentImages",
		},
		Permission: "ListCuratedEnvironmentImages",
	},
	"ListProjects": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "CodeBuild_20161006.ListProjects",
		},
		Permission: "ListProjects",
	},
	"ListReportGroups": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "CodeBuild_20161006.ListReportGroups",
		},
		Permission: "ListReportGroups",
	},
	"ListReports": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "CodeBuild_20161006.ListReports",
		},
		Permission: "ListReports",
	},
	"ListSharedProjects": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "CodeBuild_20161006.ListSharedProjects",
		},
		Permission: "ListSharedProjects",
	},
	"ListSharedReportGroups": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "CodeBuild_20161006.ListSharedReportGroups",
		},
		Permission: "ListSharedReportGroups",
	},
	"ListSourceCredentials": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "CodeBuild_20161006.ListSourceCredentials",
		},
		Permission: "ListSourceCredentials",
	},
	"RetryBuild": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "CodeBuild_20161006.RetryBuild",
		},
		Permission: "RetryBuild",
	},
	"RetryBuildBatch": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "CodeBuild_20161006.RetryBuildBatch",
		},
		Permission: "RetryBuildBatch",
	},

	// extra
	"DescribeCodeCoverages": {
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
	"DescribeTestCases": {
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
	"GetResourcePolicy": {
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
	"ListBuildsForProject": {
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
	"ListReportsForReportGroup": {
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
