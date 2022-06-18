package policy

import "github.com/securisec/cliam/shared"

var NetworkFirewallPolicies = map[string]Service{
	"DeleteFirewall": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_0,
			aws_X_AMZ_TARGET:           "NetworkFirewall_20201112.DeleteFirewall",
		},
		Permission: "DeleteFirewall",
	},
	"DeleteFirewallPolicy": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_0,
			aws_X_AMZ_TARGET:           "NetworkFirewall_20201112.DeleteFirewallPolicy",
		},
		Permission: "DeleteFirewallPolicy",
	},
	"DeleteRuleGroup": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_0,
			aws_X_AMZ_TARGET:           "NetworkFirewall_20201112.DeleteRuleGroup",
		},
		Permission: "DeleteRuleGroup",
	},
	"DescribeFirewall": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_0,
			aws_X_AMZ_TARGET:           "NetworkFirewall_20201112.DescribeFirewall",
		},
		Permission: "DescribeFirewall",
	},
	"DescribeFirewallPolicy": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_0,
			aws_X_AMZ_TARGET:           "NetworkFirewall_20201112.DescribeFirewallPolicy",
		},
		Permission: "DescribeFirewallPolicy",
	},
	"DescribeLoggingConfiguration": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_0,
			aws_X_AMZ_TARGET:           "NetworkFirewall_20201112.DescribeLoggingConfiguration",
		},
		Permission: "DescribeLoggingConfiguration",
	},
	"DescribeRuleGroup": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_0,
			aws_X_AMZ_TARGET:           "NetworkFirewall_20201112.DescribeRuleGroup",
		},
		Permission: "DescribeRuleGroup",
	},
	"DescribeRuleGroupMetadata": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_0,
			aws_X_AMZ_TARGET:           "NetworkFirewall_20201112.DescribeRuleGroupMetadata",
		},
		Permission: "DescribeRuleGroupMetadata",
	},
	"ListFirewallPolicies": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_0,
			aws_X_AMZ_TARGET:           "NetworkFirewall_20201112.ListFirewallPolicies",
		},
		Permission: "ListFirewallPolicies",
	},
	"ListFirewalls": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_0,
			aws_X_AMZ_TARGET:           "NetworkFirewall_20201112.ListFirewalls",
		},
		Permission: "ListFirewalls",
	},
	"ListRuleGroups": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_0,
			aws_X_AMZ_TARGET:           "NetworkFirewall_20201112.ListRuleGroups",
		},
		Permission: "ListRuleGroups",
	},
	"UpdateFirewallDescription": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_0,
			aws_X_AMZ_TARGET:           "NetworkFirewall_20201112.UpdateFirewallDescription",
		},
		Permission: "UpdateFirewallDescription",
	},
	"UpdateLoggingConfiguration": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_0,
			aws_X_AMZ_TARGET:           "NetworkFirewall_20201112.UpdateLoggingConfiguration",
		},
		Permission: "UpdateLoggingConfiguration",
	},

	// extra
	"DescribeResourcePolicy": {
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
	"ListTagsForResource": {
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
