package policy

import "github.com/securisec/cliam/shared"

var NetworkFirewallPolicies = []Service{
	{
		Method:   "POST",
		JsonData: `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1,
			aws_X_AMZ_TARGET:           "NetworkFirewall_20201112.DeleteFirewall",
		},
		Permission: "DeleteFirewall",
	},
	{
		Method:   "POST",
		JsonData: `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1,
			aws_X_AMZ_TARGET:           "NetworkFirewall_20201112.DeleteFirewallPolicy",
		},
		Permission: "DeleteFirewallPolicy",
	},
	{
		Method:   "POST",
		JsonData: `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1,
			aws_X_AMZ_TARGET:           "NetworkFirewall_20201112.DeleteRuleGroup",
		},
		Permission: "DeleteRuleGroup",
	},
	{
		Method:   "POST",
		JsonData: `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1,
			aws_X_AMZ_TARGET:           "NetworkFirewall_20201112.DescribeFirewall",
		},
		Permission: "DescribeFirewall",
	},
	{
		Method:   "POST",
		JsonData: `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1,
			aws_X_AMZ_TARGET:           "NetworkFirewall_20201112.DescribeFirewallPolicy",
		},
		Permission: "DescribeFirewallPolicy",
	},
	{
		Method:   "POST",
		JsonData: `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1,
			aws_X_AMZ_TARGET:           "NetworkFirewall_20201112.DescribeLoggingConfiguration",
		},
		Permission: "DescribeLoggingConfiguration",
	},
	{
		Method:   "POST",
		JsonData: `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1,
			aws_X_AMZ_TARGET:           "NetworkFirewall_20201112.DescribeRuleGroup",
		},
		Permission: "DescribeRuleGroup",
	},
	{
		Method:   "POST",
		JsonData: `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1,
			aws_X_AMZ_TARGET:           "NetworkFirewall_20201112.DescribeRuleGroupMetadata",
		},
		Permission: "DescribeRuleGroupMetadata",
	},
	{
		Method:   "POST",
		JsonData: `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1,
			aws_X_AMZ_TARGET:           "NetworkFirewall_20201112.ListFirewallPolicies",
		},
		Permission: "ListFirewallPolicies",
	},
	{
		Method:   "POST",
		JsonData: `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1,
			aws_X_AMZ_TARGET:           "NetworkFirewall_20201112.ListFirewalls",
		},
		Permission: "ListFirewalls",
	},
	{
		Method:   "POST",
		JsonData: `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1,
			aws_X_AMZ_TARGET:           "NetworkFirewall_20201112.ListRuleGroups",
		},
		Permission: "ListRuleGroups",
	},
	{
		Method:   "POST",
		JsonData: `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1,
			aws_X_AMZ_TARGET:           "NetworkFirewall_20201112.UpdateFirewallDescription",
		},
		Permission: "UpdateFirewallDescription",
	},
	{
		Method:   "POST",
		JsonData: `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1,
			aws_X_AMZ_TARGET:           "NetworkFirewall_20201112.UpdateLoggingConfiguration",
		},
		Permission: "UpdateLoggingConfiguration",
	},
}