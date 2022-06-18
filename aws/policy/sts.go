package policy

import "github.com/securisec/cliam/shared"

var STSPolicies = map[string]Service{
	"GetCallerIdentity": {
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
	"GetSessionToken": {
		ServiceSuffix: "?Action=GetSessionToken&Version=2011-06-15",
		Permission:    "GetSessionToken",
	},

	// extra
	"GetAccessKeyInfo": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "GetAccessKeyInfo",
			"Version": "2011-06-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "GetAccessKeyInfo",
		IsExtra:                true,
		ExtraComponentBodyKey:  "AccessKeyId",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "access_key_id",
	},
	"GetFederationToken": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "GetFederationToken",
			"Version": "2011-06-15",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "GetFederationToken",
		IsExtra:                true,
		ExtraComponentBodyKey:  "Name",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "name",
	},
}
