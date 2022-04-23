package policy

import "github.com/securisec/cliam/shared"

var VoiceIdPolicies = []Service{
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
			aws_X_AMZ_TARGET:           "VoiceID.ListDomains",
		},
		Permission: "ListDomains",
	},
}
