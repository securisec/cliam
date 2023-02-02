package policy

import "github.com/securisec/cliam/shared"

// WAFV2Policies policy
var WAFV2Policies = map[string]Service{
	"GetRuleGroup": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSWAF_20190729.GetRuleGroup",
		},
		Permission: "GetRuleGroup",
	},

	// extra
	"GetLoggingConfiguration": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSWAF_20190729.GetLoggingConfiguration",
		},
		Permission:             "GetLoggingConfiguration",
		IsExtra:                true,
		ExtraComponentBodyKey:  "ResourceArn",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "resource_arn",
	},
	"GetPermissionPolicy": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSWAF_20190729.GetPermissionPolicy",
		},
		Permission:             "GetPermissionPolicy",
		IsExtra:                true,
		ExtraComponentBodyKey:  "ResourceArn",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "resource_arn",
	},
	"GetWebACLForResource": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSWAF_20190729.GetWebACLForResource",
		},
		Permission:             "GetWebACLForResource",
		IsExtra:                true,
		ExtraComponentBodyKey:  "ResourceArn",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "resource_arn",
	},
	"ListAvailableManagedRuleGroups": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSWAF_20190729.ListAvailableManagedRuleGroups",
		},
		Permission:             "ListAvailableManagedRuleGroups",
		IsExtra:                true,
		ExtraComponentBodyKey:  "Scope",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "scope",
	},
	"ListIPSets": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSWAF_20190729.ListIPSets",
		},
		Permission:             "ListIPSets",
		IsExtra:                true,
		ExtraComponentBodyKey:  "Scope",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "scope",
	},
	"ListLoggingConfigurations": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSWAF_20190729.ListLoggingConfigurations",
		},
		Permission:             "ListLoggingConfigurations",
		IsExtra:                true,
		ExtraComponentBodyKey:  "Scope",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "scope",
	},
	"ListManagedRuleSets": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSWAF_20190729.ListManagedRuleSets",
		},
		Permission:             "ListManagedRuleSets",
		IsExtra:                true,
		ExtraComponentBodyKey:  "Scope",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "scope",
	},
	"ListMobileSdkReleases": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSWAF_20190729.ListMobileSdkReleases",
		},
		Permission:             "ListMobileSdkReleases",
		IsExtra:                true,
		ExtraComponentBodyKey:  "Platform",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "platform",
	},
	"ListRegexPatternSets": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSWAF_20190729.ListRegexPatternSets",
		},
		Permission:             "ListRegexPatternSets",
		IsExtra:                true,
		ExtraComponentBodyKey:  "Scope",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "scope",
	},
	"ListResourcesForWebACL": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSWAF_20190729.ListResourcesForWebACL",
		},
		Permission:             "ListResourcesForWebACL",
		IsExtra:                true,
		ExtraComponentBodyKey:  "WebACLArn",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "web_a_c_l_arn",
	},
	"ListRuleGroups": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSWAF_20190729.ListRuleGroups",
		},
		Permission:             "ListRuleGroups",
		IsExtra:                true,
		ExtraComponentBodyKey:  "Scope",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "scope",
	},
	"ListTagsForResource": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSWAF_20190729.ListTagsForResource",
		},
		Permission:             "ListTagsForResource",
		IsExtra:                true,
		ExtraComponentBodyKey:  "ResourceARN",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "resource_a_r_n",
	},
	"ListWebACLs": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSWAF_20190729.ListWebACLs",
		},
		Permission:             "ListWebACLs",
		IsExtra:                true,
		ExtraComponentBodyKey:  "Scope",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "scope",
	},
}
