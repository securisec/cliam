package policy

import "github.com/securisec/cliam/shared"

var WAFPolicies = []Service{
	{
		IgnoreRegion: true,
		Method:       "POST",
		JsonData:     map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSWAF_20150824.GetChangeToken",
		},
		Permission: "GetChangeToken",
	},
	{
		IgnoreRegion: true,
		Method:       "POST",
		JsonData:     map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSWAF_20150824.ListActivatedRulesInRuleGroup",
		},
		Permission: "ListActivatedRulesInRuleGroup",
	},
	{
		IgnoreRegion: true,
		Method:       "POST",
		JsonData:     map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSWAF_20150824.ListByteMatchSets",
		},
		Permission: "ListByteMatchSets",
	},
	{
		IgnoreRegion: true,
		Method:       "POST",
		JsonData:     map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSWAF_20150824.ListGeoMatchSets",
		},
		Permission: "ListGeoMatchSets",
	},
	{
		IgnoreRegion: true,
		Method:       "POST",
		JsonData:     map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSWAF_20150824.ListIpSets",
		},
		Permission: "ListIpSets",
	},
	{
		IgnoreRegion: true,
		Method:       "POST",
		JsonData:     map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSWAF_20150824.ListLoggingConfigurations",
		},
		Permission: "ListLoggingConfigurations",
	},
	{
		IgnoreRegion: true,
		Method:       "POST",
		JsonData:     map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSWAF_20150824.ListRateBasedRules",
		},
		Permission: "ListRateBasedRules",
	},
	{
		IgnoreRegion: true,
		Method:       "POST",
		JsonData:     map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSWAF_20150824.ListRegexMatchSets",
		},
		Permission: "ListRegexMatchSets",
	},
	{
		IgnoreRegion: true,
		Method:       "POST",
		JsonData:     map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSWAF_20150824.ListRegexPatternSets",
		},
		Permission: "ListRegexPatternSets",
	},
	{
		IgnoreRegion: true,
		Method:       "POST",
		JsonData:     map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSWAF_20150824.ListRuleGroups",
		},
		Permission: "ListRuleGroups",
	},
	{
		IgnoreRegion: true,
		Method:       "POST",
		JsonData:     map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSWAF_20150824.ListRules",
		},
		Permission: "ListRules",
	},
	{
		IgnoreRegion: true,
		Method:       "POST",
		JsonData:     map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSWAF_20150824.ListSizeConstraintSets",
		},
		Permission: "ListSizeConstraintSets",
	},
	{
		IgnoreRegion: true,
		Method:       "POST",
		JsonData:     map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSWAF_20150824.ListSqlInjectionMatchSets",
		},
		Permission: "ListSqlInjectionMatchSets",
	},
	{
		IgnoreRegion: true,
		Method:       "POST",
		JsonData:     map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSWAF_20150824.ListSubscribedRuleGroups",
		},
		Permission: "ListSubscribedRuleGroups",
	},
	{
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
	{
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
	{
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
	{
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
	{
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
	{
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
	{
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
	{
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
	{
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
	{
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
	{
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
	{
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
	{
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
	{
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
	{
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
	{
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
	{
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
	{
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
