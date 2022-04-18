package policy

import "github.com/securisec/cliam/shared"

var SecurityHubPolicies = []Service{
	{
		Method:        "POST",
		ServiceSuffix: "actionTargets/get",
		JsonData:      `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
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
		JsonData:      `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
		},
		Permission: "DisassociateFromAdministratorAccount",
	},
	{
		Method:        "POST",
		ServiceSuffix: "master/disassociate",
		JsonData:      `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
		},
		Permission: "DisassociateFromMasterAccount",
	},
	{
		Method:        "POST",
		ServiceSuffix: "accounts",
		JsonData:      `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
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
		JsonData:      `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
		},
		Permission: "GetEnabledStandards",
	},
	{
		Method:        "POST",
		ServiceSuffix: "findings",
		JsonData:      `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
		},
		Permission: "GetFindings",
	},
	{
		Method:        "POST",
		ServiceSuffix: "insights/get",
		JsonData:      `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
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
