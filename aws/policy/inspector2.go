package policy

import "github.com/securisec/cliam/shared"

var Inspector2Policies = []Service{
	{
		Method:        "POST",
		ServiceSuffix: "status/batch/get",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		},
		Permission: "BatchGetAccountStatus",
	},
	{
		Method:        "POST",
		ServiceSuffix: "organizationconfiguration/describe",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		},
		Permission: "DescribeOrganizationConfiguration",
	},
	{
		Method:        "POST",
		ServiceSuffix: "disable",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		},
		Permission: "Disable",
	},
	{
		Method:        "POST",
		ServiceSuffix: "delegatedadminaccounts/get",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		},
		Permission: "GetDelegatedAdminAccount",
	},
	{
		Method:        "POST",
		ServiceSuffix: "reporting/status/get",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		},
		Permission: "GetFindingsReportStatus",
	},
	{
		Method:        "POST",
		ServiceSuffix: "accountpermissions/list",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		},
		Permission: "ListAccountPermissions",
	},
	{
		Method:        "POST",
		ServiceSuffix: "coverage/list",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		},
		Permission: "ListCoverage",
	},
	{
		Method:        "POST",
		ServiceSuffix: "coverage/statistics/list",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		},
		Permission: "ListCoverageStatistics",
	},
	{
		Method:        "POST",
		ServiceSuffix: "delegatedadminaccounts/list",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		},
		Permission: "ListDelegatedAdminAccounts",
	},
	{
		Method:        "POST",
		ServiceSuffix: "filters/list",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		},
		Permission: "ListFilters",
	},
	{
		Method:        "POST",
		ServiceSuffix: "findings/list",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		},
		Permission: "ListFindings",
	},
	{
		Method:        "POST",
		ServiceSuffix: "members/list",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		},
		Permission: "ListMembers",
	},
	{
		Method:        "POST",
		ServiceSuffix: "usage/list",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		},
		Permission: "ListUsageTotals",
	},
}
