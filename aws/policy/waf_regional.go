package policy

import "github.com/securisec/cliam/shared"

// WAFRegionalPolicies policy
var WAFRegionalPolicies = map[string]Service{
	"GetChangeToken": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSWAF_Regional_20161128.GetChangeToken",
		},
		Permission: "GetChangeToken",
	},
	"ListActivatedRulesInRuleGroup": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSWAF_Regional_20161128.ListActivatedRulesInRuleGroup",
		},
		Permission: "ListActivatedRulesInRuleGroup",
	},
	"ListByteMatchSets": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSWAF_Regional_20161128.ListByteMatchSets",
		},
		Permission: "ListByteMatchSets",
	},
	"ListGeoMatchSets": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSWAF_Regional_20161128.ListGeoMatchSets",
		},
		Permission: "ListGeoMatchSets",
	},
	"ListIPSets": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSWAF_Regional_20161128.ListIPSets",
		},
		Permission: "ListIPSets",
	},
	"ListLoggingConfigurations": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSWAF_Regional_20161128.ListLoggingConfigurations",
		},
		Permission: "ListLoggingConfigurations",
	},
	"ListRateBasedRules": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSWAF_Regional_20161128.ListRateBasedRules",
		},
		Permission: "ListRateBasedRules",
	},
	"ListRegexMatchSets": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSWAF_Regional_20161128.ListRegexMatchSets",
		},
		Permission: "ListRegexMatchSets",
	},
	"ListRegexPatternSets": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSWAF_Regional_20161128.ListRegexPatternSets",
		},
		Permission: "ListRegexPatternSets",
	},
	"ListRuleGroups": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSWAF_Regional_20161128.ListRuleGroups",
		},
		Permission: "ListRuleGroups",
	},
	"ListRules": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSWAF_Regional_20161128.ListRules",
		},
		Permission: "ListRules",
	},
	"ListSizeConstraintSets": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSWAF_Regional_20161128.ListSizeConstraintSets",
		},
		Permission: "ListSizeConstraintSets",
	},
	"ListSqlInjectionMatchSets": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSWAF_Regional_20161128.ListSqlInjectionMatchSets",
		},
		Permission: "ListSqlInjectionMatchSets",
	},
	"ListSubscribedRuleGroups": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSWAF_Regional_20161128.ListSubscribedRuleGroups",
		},
		Permission: "ListSubscribedRuleGroups",
	},
	"ListWebACLs": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSWAF_Regional_20161128.ListWebACLs",
		},
		Permission: "ListWebACLs",
	},
	"ListXssMatchSets": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSWAF_Regional_20161128.ListXssMatchSets",
		},
		Permission: "ListXssMatchSets",
	},

	// extra
	"GetByteMatchSet": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSWAF_Regional_20161128.GetByteMatchSet",
		},
		Permission:             "GetByteMatchSet",
		IsExtra:                true,
		ExtraComponentBodyKey:  "ByteMatchSetId",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "byte_match_set_id",
	},
	"GetChangeTokenStatus": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSWAF_Regional_20161128.GetChangeTokenStatus",
		},
		Permission:             "GetChangeTokenStatus",
		IsExtra:                true,
		ExtraComponentBodyKey:  "ChangeToken",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "change_token",
	},
	"GetGeoMatchSet": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSWAF_Regional_20161128.GetGeoMatchSet",
		},
		Permission:             "GetGeoMatchSet",
		IsExtra:                true,
		ExtraComponentBodyKey:  "GeoMatchSetId",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "geo_match_set_id",
	},
	"GetIPSet": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSWAF_Regional_20161128.GetIPSet",
		},
		Permission:             "GetIPSet",
		IsExtra:                true,
		ExtraComponentBodyKey:  "IPSetId",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "i_p_set_id",
	},
	"GetLoggingConfiguration": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSWAF_Regional_20161128.GetLoggingConfiguration",
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
			aws_X_AMZ_TARGET:           "AWSWAF_Regional_20161128.GetPermissionPolicy",
		},
		Permission:             "GetPermissionPolicy",
		IsExtra:                true,
		ExtraComponentBodyKey:  "ResourceArn",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "resource_arn",
	},
	"GetRateBasedRule": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSWAF_Regional_20161128.GetRateBasedRule",
		},
		Permission:             "GetRateBasedRule",
		IsExtra:                true,
		ExtraComponentBodyKey:  "RuleId",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "rule_id",
	},
	"GetRateBasedRuleManagedKeys": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSWAF_Regional_20161128.GetRateBasedRuleManagedKeys",
		},
		Permission:             "GetRateBasedRuleManagedKeys",
		IsExtra:                true,
		ExtraComponentBodyKey:  "RuleId",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "rule_id",
	},
	"GetRegexMatchSet": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSWAF_Regional_20161128.GetRegexMatchSet",
		},
		Permission:             "GetRegexMatchSet",
		IsExtra:                true,
		ExtraComponentBodyKey:  "RegexMatchSetId",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "regex_match_set_id",
	},
	"GetRegexPatternSet": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSWAF_Regional_20161128.GetRegexPatternSet",
		},
		Permission:             "GetRegexPatternSet",
		IsExtra:                true,
		ExtraComponentBodyKey:  "RegexPatternSetId",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "regex_pattern_set_id",
	},
	"GetRule": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSWAF_Regional_20161128.GetRule",
		},
		Permission:             "GetRule",
		IsExtra:                true,
		ExtraComponentBodyKey:  "RuleId",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "rule_id",
	},
	"GetRuleGroup": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSWAF_Regional_20161128.GetRuleGroup",
		},
		Permission:             "GetRuleGroup",
		IsExtra:                true,
		ExtraComponentBodyKey:  "RuleGroupId",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "rule_group_id",
	},
	"GetSizeConstraintSet": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSWAF_Regional_20161128.GetSizeConstraintSet",
		},
		Permission:             "GetSizeConstraintSet",
		IsExtra:                true,
		ExtraComponentBodyKey:  "SizeConstraintSetId",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "size_constraint_set_id",
	},
	"GetSqlInjectionMatchSet": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSWAF_Regional_20161128.GetSqlInjectionMatchSet",
		},
		Permission:             "GetSqlInjectionMatchSet",
		IsExtra:                true,
		ExtraComponentBodyKey:  "SqlInjectionMatchSetId",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "sql_injection_match_set_id",
	},
	"GetWebACL": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSWAF_Regional_20161128.GetWebACL",
		},
		Permission:             "GetWebACL",
		IsExtra:                true,
		ExtraComponentBodyKey:  "WebACLId",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "web_a_c_l_id",
	},
	"GetWebACLForResource": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSWAF_Regional_20161128.GetWebACLForResource",
		},
		Permission:             "GetWebACLForResource",
		IsExtra:                true,
		ExtraComponentBodyKey:  "ResourceArn",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "resource_arn",
	},
	"GetXssMatchSet": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSWAF_Regional_20161128.GetXssMatchSet",
		},
		Permission:             "GetXssMatchSet",
		IsExtra:                true,
		ExtraComponentBodyKey:  "XssMatchSetId",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "xss_match_set_id",
	},
	"ListResourcesForWebACL": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSWAF_Regional_20161128.ListResourcesForWebACL",
		},
		Permission:             "ListResourcesForWebACL",
		IsExtra:                true,
		ExtraComponentBodyKey:  "WebACLId",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "web_a_c_l_id",
	},
	"ListTagsForResource": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSWAF_Regional_20161128.ListTagsForResource",
		},
		Permission:             "ListTagsForResource",
		IsExtra:                true,
		ExtraComponentBodyKey:  "ResourceARN",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "resource_a_r_n",
	},
}
