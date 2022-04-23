package policy

import "github.com/securisec/cliam/shared"

var Inspector2Policies = []Service{
	{
		Method:        "POST",
		ServiceSuffix: "status/batch/get",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
		},
		Permission: "BatchGetAccountStatus",
	},
	{
		Method:        "POST",
		ServiceSuffix: "organizationconfiguration/describe",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
		},
		Permission: "DescribeOrganizationConfiguration",
	},
	{
		Method:        "POST",
		ServiceSuffix: "disable",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
		},
		Permission: "Disable",
	},
	{
		Method:        "POST",
		ServiceSuffix: "delegatedadminaccounts/get",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
		},
		Permission: "GetDelegatedAdminAccount",
	},
	{
		Method:        "POST",
		ServiceSuffix: "reporting/status/get",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
		},
		Permission: "GetFindingsReportStatus",
	},
	{
		Method:        "POST",
		ServiceSuffix: "accountpermissions/list",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
		},
		Permission: "ListAccountPermissions",
	},
	{
		Method:        "POST",
		ServiceSuffix: "coverage/list",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
		},
		Permission: "ListCoverage",
	},
	{
		Method:        "POST",
		ServiceSuffix: "coverage/statistics/list",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
		},
		Permission: "ListCoverageStatistics",
	},
	{
		Method:        "POST",
		ServiceSuffix: "delegatedadminaccounts/list",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
		},
		Permission: "ListDelegatedAdminAccounts",
	},
	{
		Method:        "POST",
		ServiceSuffix: "filters/list",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
		},
		Permission: "ListFilters",
	},
	{
		Method:        "POST",
		ServiceSuffix: "findings/list",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
		},
		Permission: "ListFindings",
	},
	{
		Method:        "POST",
		ServiceSuffix: "members/list",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
		},
		Permission: "ListMembers",
	},
	{
		Method:        "POST",
		ServiceSuffix: "usage/list",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
		},
		Permission: "ListUsageTotals",
	},
}
