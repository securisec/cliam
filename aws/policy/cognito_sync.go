package policy

import "github.com/securisec/cliam/shared"

var CognitoSyncPolicies = map[string]Service{
	"ListIdentityPoolUsage": {
		Method:        "GET",
		ServiceSuffix: "identitypools",
		Permission:    "ListIdentityPoolUsage",
	},

	// extra
	"DescribeIdentityPoolUsage": {
		ServiceSuffix:          "/identitypools/{{.identity_pool_id}}",
		Permission:             "DescribeIdentityPoolUsage",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "identity_pool_id",
	},
	"GetBulkPublishDetails": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "GetBulkPublishDetails",
			"Version": "2014-06-30",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "GetBulkPublishDetails",
		IsExtra:                true,
		ExtraComponentBodyKey:  "IdentityPoolId",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "identity_pool_id",
	},
	"GetCognitoEvents": {
		ServiceSuffix:          "/identitypools/{{.identity_pool_id}}/events",
		Permission:             "GetCognitoEvents",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "identity_pool_id",
	},
	"GetIdentityPoolConfiguration": {
		ServiceSuffix:          "/identitypools/{{.identity_pool_id}}/configuration",
		Permission:             "GetIdentityPoolConfiguration",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "identity_pool_id",
	},
}
