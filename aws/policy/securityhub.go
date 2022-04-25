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

	// extra
	{
		ServiceSuffix:          "/standards/controls/{{.standards_subscription_arn}}",
		Permission:             "DescribeStandardsControls",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "standards_subscription_arn",
	},
	{
		ServiceSuffix:          "/findingAggregator/get/{{.finding_aggregator_arn}}",
		Permission:             "GetFindingAggregator",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "finding_aggregator_arn",
	},
	{
		ServiceSuffix:          "/insights/results/{{.insight_arn}}",
		Permission:             "GetInsightResults",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "insight_arn",
	},
	{
		Method: "POST",
		FormData: map[string]string{
			"Action":  "GetMembers",
			"Version": "2018-10-26",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "GetMembers",
		IsExtra:                true,
		ExtraComponentBodyKey:  "AccountIds",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "account_ids",
	},
	{
		ServiceSuffix:          "/tags/{{.resource_arn}}",
		Permission:             "ListTagsForResource",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "resource_arn",
	},
}
