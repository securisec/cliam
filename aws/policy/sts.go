package policy

import "github.com/securisec/enumerate/shared"

var STSPolicies = []Service{
	{
		ServiceSuffix: "",
		Permission:    "GetCallerIdentity",
		FormData: map[string]string{
			"Action":  "GetCallerIdentity",
			"Version": "2011-06-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Method: "POST",
	},
	{
		ServiceSuffix: "?Action=GetSessionToken&Version=2011-06-15",
		Permission:    "GetSessionToken",
	},
}
