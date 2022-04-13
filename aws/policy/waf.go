package policy

import "github.com/securisec/cliam/shared"

var WAFPolicies = []Service{
	{
		IgnoreRegion: true,
		Method:       "POST",
		JsonData:     `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
			aws_X_AMZ_TARGET:           "AWSWAF_20150824.GetChangeToken",
		},
		Permission: "GetChangeToken",
	},
	{
		IgnoreRegion: true,
		Method:       "POST",
		JsonData:     `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
			aws_X_AMZ_TARGET:           "AWSWAF_20150824.ListActivatedRulesInRuleGroup",
		},
		Permission: "ListActivatedRulesInRuleGroup",
	},
	{
		IgnoreRegion: true,
		Method:       "POST",
		JsonData:     `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
			aws_X_AMZ_TARGET:           "AWSWAF_20150824.ListByteMatchSets",
		},
		Permission: "ListByteMatchSets",
	},
	{
		IgnoreRegion: true,
		Method:       "POST",
		JsonData:     `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
			aws_X_AMZ_TARGET:           "AWSWAF_20150824.ListGeoMatchSets",
		},
		Permission: "ListGeoMatchSets",
	},
	{
		IgnoreRegion: true,
		Method:       "POST",
		JsonData:     `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
			aws_X_AMZ_TARGET:           "AWSWAF_20150824.ListIpSets",
		},
		Permission: "ListIpSets",
	},
	{
		IgnoreRegion: true,
		Method:       "POST",
		JsonData:     `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
			aws_X_AMZ_TARGET:           "AWSWAF_20150824.ListLoggingConfigurations",
		},
		Permission: "ListLoggingConfigurations",
	},
	{
		IgnoreRegion: true,
		Method:       "POST",
		JsonData:     `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
			aws_X_AMZ_TARGET:           "AWSWAF_20150824.ListRateBasedRules",
		},
		Permission: "ListRateBasedRules",
	},
	{
		IgnoreRegion: true,
		Method:       "POST",
		JsonData:     `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
			aws_X_AMZ_TARGET:           "AWSWAF_20150824.ListRegexMatchSets",
		},
		Permission: "ListRegexMatchSets",
	},
	{
		IgnoreRegion: true,
		Method:       "POST",
		JsonData:     `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
			aws_X_AMZ_TARGET:           "AWSWAF_20150824.ListRegexPatternSets",
		},
		Permission: "ListRegexPatternSets",
	},
	{
		IgnoreRegion: true,
		Method:       "POST",
		JsonData:     `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
			aws_X_AMZ_TARGET:           "AWSWAF_20150824.ListRuleGroups",
		},
		Permission: "ListRuleGroups",
	},
	{
		IgnoreRegion: true,
		Method:       "POST",
		JsonData:     `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
			aws_X_AMZ_TARGET:           "AWSWAF_20150824.ListRules",
		},
		Permission: "ListRules",
	},
	{
		IgnoreRegion: true,
		Method:       "POST",
		JsonData:     `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
			aws_X_AMZ_TARGET:           "AWSWAF_20150824.ListSizeConstraintSets",
		},
		Permission: "ListSizeConstraintSets",
	},
	{
		IgnoreRegion: true,
		Method:       "POST",
		JsonData:     `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
			aws_X_AMZ_TARGET:           "AWSWAF_20150824.ListSqlInjectionMatchSets",
		},
		Permission: "ListSqlInjectionMatchSets",
	},
	{
		IgnoreRegion: true,
		Method:       "POST",
		JsonData:     `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
			aws_X_AMZ_TARGET:           "AWSWAF_20150824.ListSubscribedRuleGroups",
		},
		Permission: "ListSubscribedRuleGroups",
	},
	{
		IgnoreRegion: true,
		Method:       "POST",
		JsonData:     `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
			aws_X_AMZ_TARGET:           "AWSWAF_20150824.ListXssMatchSets",
		},
		Permission: "ListXssMatchSets",
	},
}
