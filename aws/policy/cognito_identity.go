package policy

import "github.com/securisec/cliam/shared"

var CognitoIdentityPolicies = map[string]Service{

	// extra
	"DescribeIdentity": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSCognitoIdentityService.DescribeIdentity",
		},
		Permission:             "DescribeIdentity",
		IsExtra:                true,
		ExtraComponentBodyKey:  "IdentityId",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "identity_id",
	},
	"DescribeIdentityPool": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSCognitoIdentityService.DescribeIdentityPool",
		},
		Permission:             "DescribeIdentityPool",
		IsExtra:                true,
		ExtraComponentBodyKey:  "IdentityPoolId",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "identity_pool_id",
	},
	"GetCredentialsForIdentity": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSCognitoIdentityService.GetCredentialsForIdentity",
		},
		Permission:             "GetCredentialsForIdentity",
		IsExtra:                true,
		ExtraComponentBodyKey:  "IdentityId",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "identity_id",
	},
	"GetId": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSCognitoIdentityService.GetId",
		},
		Permission:             "GetId",
		IsExtra:                true,
		ExtraComponentBodyKey:  "IdentityPoolId",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "identity_pool_id",
	},
	"GetIdentityPoolRoles": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSCognitoIdentityService.GetIdentityPoolRoles",
		},
		Permission:             "GetIdentityPoolRoles",
		IsExtra:                true,
		ExtraComponentBodyKey:  "IdentityPoolId",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "identity_pool_id",
	},
	"GetOpenIdToken": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSCognitoIdentityService.GetOpenIdToken",
		},
		Permission:             "GetOpenIdToken",
		IsExtra:                true,
		ExtraComponentBodyKey:  "IdentityId",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "identity_id",
	},
	"ListIdentityPools": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSCognitoIdentityService.ListIdentityPools",
		},
		Permission:             "ListIdentityPools",
		IsExtra:                true,
		ExtraComponentBodyKey:  "MaxResults",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "max_results",
	},
	"ListTagsForResource": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSCognitoIdentityService.ListTagsForResource",
		},
		Permission:             "ListTagsForResource",
		IsExtra:                true,
		ExtraComponentBodyKey:  "ResourceArn",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "resource_arn",
	},
}
