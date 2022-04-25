package policy

import "github.com/securisec/cliam/shared"

var SecurityHubPolicies = []Service{
	{
		Method:        "POST",
		ServiceSuffix: "actionTargets/get",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		},
		Permission: "DescribeActionTargets",
	},
	{
		Method:        "GET",
		ServiceSuffix: "accounts",
		Permission:    "DescribeHub",
	},
	{
		Method:        "GET",
		ServiceSuffix: "organization/configuration",
		Permission:    "DescribeOrganizationConfiguration",
	},
	{
		Method:        "GET",
		ServiceSuffix: "products",
		Permission:    "DescribeProducts",
	},
	{
		Method:        "GET",
		ServiceSuffix: "standards",
		Permission:    "DescribeStandards",
	},
	{
		Method:        "POST",
		ServiceSuffix: "administrator/disassociate",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		},
		Permission: "DisassociateFromAdministratorAccount",
	},
	{
		Method:        "POST",
		ServiceSuffix: "master/disassociate",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		},
		Permission: "DisassociateFromMasterAccount",
	},
	{
		Method:        "POST",
		ServiceSuffix: "accounts",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		},
		Permission: "EnableSecurityHub",
	},
	{
		Method:        "GET",
		ServiceSuffix: "administrator",
		Permission:    "GetAdministratorAccount",
	},
	{
		Method:        "POST",
		ServiceSuffix: "standards/get",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		},
		Permission: "GetEnabledStandards",
	},
	{
		Method:        "POST",
		ServiceSuffix: "findings",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		},
		Permission: "GetFindings",
	},
	{
		Method:        "POST",
		ServiceSuffix: "insights/get",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		},
		Permission: "GetInsights",
	},
	{
		Method:        "GET",
		ServiceSuffix: "invitations/count",
		Permission:    "GetInvitationsCount",
	},
	{
		Method:        "GET",
		ServiceSuffix: "master",
		Permission:    "GetMasterAccount",
	},
	{
		Method:        "GET",
		ServiceSuffix: "productSubscriptions",
		Permission:    "ListEnabledProductsForImport",
	},
	{
		Method:        "GET",
		ServiceSuffix: "findingAggregator/list",
		Permission:    "ListFindingAggregators",
	},
	{
		Method:        "GET",
		ServiceSuffix: "invitations",
		Permission:    "ListInvitations",
	},
	{
		Method:        "GET",
		ServiceSuffix: "members",
		Permission:    "ListMembers",
	},
	{
		Method:        "GET",
		ServiceSuffix: "organization/admin",
		Permission:    "ListOrganizationAdminAccounts",
	},
}
