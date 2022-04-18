package policy

import "github.com/securisec/cliam/shared"

var WafV2Policies = []Service{
	{
		Method:   "POST",
		JsonData: `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
			aws_X_AMZ_TARGET:           "AWSWAF_20190729.GetRuleGroup",
		},
		Permission: "GetRuleGroup",
	},
}
