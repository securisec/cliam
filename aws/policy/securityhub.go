package policy

import "github.com/securisec/cliam/shared"

// SecurityHubPolicies policy
var SecurityHubPolicies = map[string]Service{
	"DescribeActionTargets": {
		Method:        "POST",
		ServiceSuffix: "actionTargets/get",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		},
		Permission: "DescribeActionTargets",
	},
	"DescribeHub": {
		Method:        "GET",
		ServiceSuffix: "accounts",
		Permission:    "DescribeHub",
	},
	"DescribeOrganizationConfiguration": {
		Method:        "GET",
		ServiceSuffix: "organization/configuration",
		Permission:    "DescribeOrganizationConfiguration",
	},
	"DescribeProducts": {
		Method:        "GET",
		ServiceSuffix: "products",
		Permission:    "DescribeProducts",
	},
	"DescribeStandards": {
		Method:        "GET",
		ServiceSuffix: "standards",
		Permission:    "DescribeStandards",
	},
	"DisassociateFromAdministratorAccount": {
		Method:        "POST",
		ServiceSuffix: "administrator/disassociate",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		},
		Permission: "DisassociateFromAdministratorAccount",
	},
	"DisassociateFromMasterAccount": {
		Method:        "POST",
		ServiceSuffix: "master/disassociate",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		},
		Permission: "DisassociateFromMasterAccount",
	},
	"EnableSecurityHub": {
		Method:        "POST",
		ServiceSuffix: "accounts",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		},
		Permission: "EnableSecurityHub",
	},
	"GetAdministratorAccount": {
		Method:        "GET",
		ServiceSuffix: "administrator",
		Permission:    "GetAdministratorAccount",
	},
	"GetEnabledStandards": {
		Method:        "POST",
		ServiceSuffix: "standards/get",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		},
		Permission: "GetEnabledStandards",
	},
	"GetFindings": {
		Method:        "POST",
		ServiceSuffix: "findings",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		},
		Permission: "GetFindings",
	},
	"GetInsights": {
		Method:        "POST",
		ServiceSuffix: "insights/get",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		},
		Permission: "GetInsights",
	},
	"GetInvitationsCount": {
		Method:        "GET",
		ServiceSuffix: "invitations/count",
		Permission:    "GetInvitationsCount",
	},
	"GetMasterAccount": {
		Method:        "GET",
		ServiceSuffix: "master",
		Permission:    "GetMasterAccount",
	},
	"ListEnabledProductsForImport": {
		Method:        "GET",
		ServiceSuffix: "productSubscriptions",
		Permission:    "ListEnabledProductsForImport",
	},
	"ListFindingAggregators": {
		Method:        "GET",
		ServiceSuffix: "findingAggregator/list",
		Permission:    "ListFindingAggregators",
	},
	"ListInvitations": {
		Method:        "GET",
		ServiceSuffix: "invitations",
		Permission:    "ListInvitations",
	},
	"ListMembers": {
		Method:        "GET",
		ServiceSuffix: "members",
		Permission:    "ListMembers",
	},
	"ListOrganizationAdminAccounts": {
		Method:        "GET",
		ServiceSuffix: "organization/admin",
		Permission:    "ListOrganizationAdminAccounts",
	},

	// extra
	"DescribeStandardsControls": {
		ServiceSuffix:          "/standards/controls/{{.standards_subscription_arn}}",
		Permission:             "DescribeStandardsControls",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "standards_subscription_arn",
	},
	"GetFindingAggregator": {
		ServiceSuffix:          "/findingAggregator/get/{{.finding_aggregator_arn}}",
		Permission:             "GetFindingAggregator",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "finding_aggregator_arn",
	},
	"GetInsightResults": {
		ServiceSuffix:          "/insights/results/{{.insight_arn}}",
		Permission:             "GetInsightResults",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "insight_arn",
	},
	"GetMembers": {
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
	"ListTagsForResource": {
		ServiceSuffix:          "/tags/{{.resource_arn}}",
		Permission:             "ListTagsForResource",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "resource_arn",
	},
}
