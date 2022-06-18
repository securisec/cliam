package policy

import "github.com/securisec/cliam/shared"

var Inspector2Policies = map[string]Service{
	"BatchGetAccountStatus": {
		Method:        "POST",
		ServiceSuffix: "status/batch/get",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		},
		Permission: "BatchGetAccountStatus",
	},
	"DescribeOrganizationConfiguration": {
		Method:        "POST",
		ServiceSuffix: "organizationconfiguration/describe",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		},
		Permission: "DescribeOrganizationConfiguration",
	},
	"Disable": {
		Method:        "POST",
		ServiceSuffix: "disable",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		},
		Permission: "Disable",
	},
	"GetDelegatedAdminAccount": {
		Method:        "POST",
		ServiceSuffix: "delegatedadminaccounts/get",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		},
		Permission: "GetDelegatedAdminAccount",
	},
	"GetFindingsReportStatus": {
		Method:        "POST",
		ServiceSuffix: "reporting/status/get",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		},
		Permission: "GetFindingsReportStatus",
	},
	"ListAccountPermissions": {
		Method:        "POST",
		ServiceSuffix: "accountpermissions/list",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		},
		Permission: "ListAccountPermissions",
	},
	"ListCoverage": {
		Method:        "POST",
		ServiceSuffix: "coverage/list",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		},
		Permission: "ListCoverage",
	},
	"ListCoverageStatistics": {
		Method:        "POST",
		ServiceSuffix: "coverage/statistics/list",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		},
		Permission: "ListCoverageStatistics",
	},
	"ListDelegatedAdminAccounts": {
		Method:        "POST",
		ServiceSuffix: "delegatedadminaccounts/list",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		},
		Permission: "ListDelegatedAdminAccounts",
	},
	"ListFilters": {
		Method:        "POST",
		ServiceSuffix: "filters/list",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		},
		Permission: "ListFilters",
	},
	"ListFindings": {
		Method:        "POST",
		ServiceSuffix: "findings/list",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		},
		Permission: "ListFindings",
	},
	"ListMembers": {
		Method:        "POST",
		ServiceSuffix: "members/list",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		},
		Permission: "ListMembers",
	},
	"ListUsageTotals": {
		Method:        "POST",
		ServiceSuffix: "usage/list",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		},
		Permission: "ListUsageTotals",
	},

	// extra
	"GetMember": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "GetMember",
			"Version": "2020-06-08",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "GetMember",
		IsExtra:                true,
		ExtraComponentBodyKey:  "accountId",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "account_id",
	},
	"ListFindingAggregations": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "ListFindingAggregations",
			"Version": "2020-06-08",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "ListFindingAggregations",
		IsExtra:                true,
		ExtraComponentBodyKey:  "aggregationType",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "aggregation_type",
	},
	"ListTagsForResource": {
		ServiceSuffix:          "/tags/{{.resource_arn}}",
		Permission:             "ListTagsForResource",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "resource_arn",
	},
}
