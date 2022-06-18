package policy

import "github.com/securisec/cliam/shared"

var WAFPolicies = map[string]Service{
	"GetChangeToken": {
		IgnoreRegion: true,
		Method:       "POST",
		JsonData:     map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSWAF_20150824.GetChangeToken",
		},
		Permission: "GetChangeToken",
	},
	"ListActivatedRulesInRuleGroup": {
		IgnoreRegion: true,
		Method:       "POST",
		JsonData:     map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSWAF_20150824.ListActivatedRulesInRuleGroup",
		},
		Permission: "ListActivatedRulesInRuleGroup",
	},
	"ListByteMatchSets": {
		IgnoreRegion: true,
		Method:       "POST",
		JsonData:     map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSWAF_20150824.ListByteMatchSets",
		},
		Permission: "ListByteMatchSets",
	},
	"ListGeoMatchSets": {
		IgnoreRegion: true,
		Method:       "POST",
		JsonData:     map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSWAF_20150824.ListGeoMatchSets",
		},
		Permission: "ListGeoMatchSets",
	},
	"ListIpSets": {
		IgnoreRegion: true,
		Method:       "POST",
		JsonData:     map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSWAF_20150824.ListIpSets",
		},
		Permission: "ListIpSets",
	},
	"ListLoggingConfigurations": {
		IgnoreRegion: true,
		Method:       "POST",
		JsonData:     map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSWAF_20150824.ListLoggingConfigurations",
		},
		Permission: "ListLoggingConfigurations",
	},
	"ListRateBasedRules": {
		IgnoreRegion: true,
		Method:       "POST",
		JsonData:     map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSWAF_20150824.ListRateBasedRules",
		},
		Permission: "ListRateBasedRules",
	},
	"ListRegexMatchSets": {
		IgnoreRegion: true,
		Method:       "POST",
		JsonData:     map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSWAF_20150824.ListRegexMatchSets",
		},
		Permission: "ListRegexMatchSets",
	},
	"ListRegexPatternSets": {
		IgnoreRegion: true,
		Method:       "POST",
		JsonData:     map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSWAF_20150824.ListRegexPatternSets",
		},
		Permission: "ListRegexPatternSets",
	},
	"ListRuleGroups": {
		IgnoreRegion: true,
		Method:       "POST",
		JsonData:     map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSWAF_20150824.ListRuleGroups",
		},
		Permission: "ListRuleGroups",
	},
	"ListRules": {
		IgnoreRegion: true,
		Method:       "POST",
		JsonData:     map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSWAF_20150824.ListRules",
		},
		Permission: "ListRules",
	},
	"ListSizeConstraintSets": {
		IgnoreRegion: true,
		Method:       "POST",
		JsonData:     map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSWAF_20150824.ListSizeConstraintSets",
		},
		Permission: "ListSizeConstraintSets",
	},
	"ListSqlInjectionMatchSets": {
		IgnoreRegion: true,
		Method:       "POST",
		JsonData:     map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSWAF_20150824.ListSqlInjectionMatchSets",
		},
		Permission: "ListSqlInjectionMatchSets",
	},
	"ListSubscribedRuleGroups": {
		IgnoreRegion: true,
		Method:       "POST",
		JsonData:     map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSWAF_20150824.ListSubscribedRuleGroups",
		},
		Permission: "ListSubscribedRuleGroups",
	},
	"ListXssMatchSets": {
		IgnoreRegion: true,
		Method:       "POST",
		JsonData:     map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSWAF_20150824.ListXssMatchSets",
		},
		Permission: "ListXssMatchSets",
	},

	// extra
	"GetByteMatchSet": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSWAF_20150824.GetByteMatchSet",
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
			aws_X_AMZ_TARGET:           "AWSWAF_20150824.GetChangeTokenStatus",
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
			aws_X_AMZ_TARGET:           "AWSWAF_20150824.GetGeoMatchSet",
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
			aws_X_AMZ_TARGET:           "AWSWAF_20150824.GetIPSet",
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
			aws_X_AMZ_TARGET:           "AWSWAF_20150824.GetLoggingConfiguration",
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
			aws_X_AMZ_TARGET:           "AWSWAF_20150824.GetPermissionPolicy",
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
			aws_X_AMZ_TARGET:           "AWSWAF_20150824.GetRateBasedRule",
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
			aws_X_AMZ_TARGET:           "AWSWAF_20150824.GetRateBasedRuleManagedKeys",
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
			aws_X_AMZ_TARGET:           "AWSWAF_20150824.GetRegexMatchSet",
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
			aws_X_AMZ_TARGET:           "AWSWAF_20150824.GetRegexPatternSet",
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
			aws_X_AMZ_TARGET:           "AWSWAF_20150824.GetRule",
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
			aws_X_AMZ_TARGET:           "AWSWAF_20150824.GetRuleGroup",
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
			aws_X_AMZ_TARGET:           "AWSWAF_20150824.GetSizeConstraintSet",
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
			aws_X_AMZ_TARGET:           "AWSWAF_20150824.GetSqlInjectionMatchSet",
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
			aws_X_AMZ_TARGET:           "AWSWAF_20150824.GetWebACL",
		},
		Permission:             "GetWebACL",
		IsExtra:                true,
		ExtraComponentBodyKey:  "WebACLId",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "web_a_c_l_id",
	},
	"GetXssMatchSet": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSWAF_20150824.GetXssMatchSet",
		},
		Permission:             "GetXssMatchSet",
		IsExtra:                true,
		ExtraComponentBodyKey:  "XssMatchSetId",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "xss_match_set_id",
	},
	"ListTagsForResource": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSWAF_20150824.ListTagsForResource",
		},
		Permission:             "ListTagsForResource",
		IsExtra:                true,
		ExtraComponentBodyKey:  "ResourceARN",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "resource_arn",
	},
}
