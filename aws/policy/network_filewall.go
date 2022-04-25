package policy

import "github.com/securisec/cliam/shared"

var NetworkFirewallPolicies = []Service{
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_0,
			aws_X_AMZ_TARGET:           "NetworkFirewall_20201112.DeleteFirewall",
		},
		Permission: "DeleteFirewall",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_0,
			aws_X_AMZ_TARGET:           "NetworkFirewall_20201112.DeleteFirewallPolicy",
		},
		Permission: "DeleteFirewallPolicy",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_0,
			aws_X_AMZ_TARGET:           "NetworkFirewall_20201112.DeleteRuleGroup",
		},
		Permission: "DeleteRuleGroup",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_0,
			aws_X_AMZ_TARGET:           "NetworkFirewall_20201112.DescribeFirewall",
		},
		Permission: "DescribeFirewall",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_0,
			aws_X_AMZ_TARGET:           "NetworkFirewall_20201112.DescribeFirewallPolicy",
		},
		Permission: "DescribeFirewallPolicy",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_0,
			aws_X_AMZ_TARGET:           "NetworkFirewall_20201112.DescribeLoggingConfiguration",
		},
		Permission: "DescribeLoggingConfiguration",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_0,
			aws_X_AMZ_TARGET:           "NetworkFirewall_20201112.DescribeRuleGroup",
		},
		Permission: "DescribeRuleGroup",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_0,
			aws_X_AMZ_TARGET:           "NetworkFirewall_20201112.DescribeRuleGroupMetadata",
		},
		Permission: "DescribeRuleGroupMetadata",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_0,
			aws_X_AMZ_TARGET:           "NetworkFirewall_20201112.ListFirewallPolicies",
		},
		Permission: "ListFirewallPolicies",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_0,
			aws_X_AMZ_TARGET:           "NetworkFirewall_20201112.ListFirewalls",
		},
		Permission: "ListFirewalls",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_0,
			aws_X_AMZ_TARGET:           "NetworkFirewall_20201112.ListRuleGroups",
		},
		Permission: "ListRuleGroups",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_0,
			aws_X_AMZ_TARGET:           "NetworkFirewall_20201112.UpdateFirewallDescription",
		},
		Permission: "UpdateFirewallDescription",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_0,
			aws_X_AMZ_TARGET:           "NetworkFirewall_20201112.UpdateLoggingConfiguration",
		},
		Permission: "UpdateLoggingConfiguration",
	},

	// extra
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_0,
			aws_X_AMZ_TARGET:           "NetworkFirewall_20201112.DescribeResourcePolicy",
		},
		Permission:             "DescribeResourcePolicy",
		IsExtra:                true,
		ExtraComponentBodyKey:  "ResourceArn",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "resource_arn",
	},
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_0,
			aws_X_AMZ_TARGET:           "NetworkFirewall_20201112.ListTagsForResource",
		},
		Permission:             "ListTagsForResource",
		IsExtra:                true,
		ExtraComponentBodyKey:  "ResourceArn",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "resource_arn",
	},
}
