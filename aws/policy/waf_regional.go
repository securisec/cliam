package policy

import "github.com/securisec/cliam/shared"

var WAFRegionalPolicies = []Service{
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSWAF_Regional_20161128.GetChangeToken",
		},
		Permission: "GetChangeToken",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSWAF_Regional_20161128.ListActivatedRulesInRuleGroup",
		},
		Permission: "ListActivatedRulesInRuleGroup",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSWAF_Regional_20161128.ListByteMatchSets",
		},
		Permission: "ListByteMatchSets",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSWAF_Regional_20161128.ListGeoMatchSets",
		},
		Permission: "ListGeoMatchSets",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSWAF_Regional_20161128.ListIpSets",
		},
		Permission: "ListIpSets",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSWAF_Regional_20161128.ListLoggingConfigurations",
		},
		Permission: "ListLoggingConfigurations",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSWAF_Regional_20161128.ListRateBasedRules",
		},
		Permission: "ListRateBasedRules",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSWAF_Regional_20161128.ListRegexMatchSets",
		},
		Permission: "ListRegexMatchSets",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSWAF_Regional_20161128.ListRegexPatternSets",
		},
		Permission: "ListRegexPatternSets",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSWAF_Regional_20161128.ListRuleGroups",
		},
		Permission: "ListRuleGroups",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSWAF_Regional_20161128.ListRules",
		},
		Permission: "ListRules",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSWAF_Regional_20161128.ListSizeConstraintSets",
		},
		Permission: "ListSizeConstraintSets",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSWAF_Regional_20161128.ListSqlInjectionMatchSets",
		},
		Permission: "ListSqlInjectionMatchSets",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSWAF_Regional_20161128.ListSubscribedRuleGroups",
		},
		Permission: "ListSubscribedRuleGroups",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSWAF_Regional_20161128.ListXssMatchSets",
		},
		Permission: "ListXssMatchSets",
	},
}
