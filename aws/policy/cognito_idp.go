package policy

import "github.com/securisec/cliam/shared"

var CognitoIDPPolicies = map[string]Service{
	"AssociateSoftwareToken": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSCognitoIdentityProviderService.AssociateSoftwareToken",
		},
		Permission: "AssociateSoftwareToken",
	},

	// extra
	"DescribeRiskConfiguration": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSCognitoIdentityProviderService.DescribeRiskConfiguration",
		},
		Permission:             "DescribeRiskConfiguration",
		IsExtra:                true,
		ExtraComponentBodyKey:  "UserPoolId",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "user_pool_id",
	},
	"DescribeUserPool": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSCognitoIdentityProviderService.DescribeUserPool",
		},
		Permission:             "DescribeUserPool",
		IsExtra:                true,
		ExtraComponentBodyKey:  "UserPoolId",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "user_pool_id",
	},
	"DescribeUserPoolDomain": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSCognitoIdentityProviderService.DescribeUserPoolDomain",
		},
		Permission:             "DescribeUserPoolDomain",
		IsExtra:                true,
		ExtraComponentBodyKey:  "Domain",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "domain",
	},
	"GetCSVHeader": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSCognitoIdentityProviderService.GetCSVHeader",
		},
		Permission:             "GetCSVHeader",
		IsExtra:                true,
		ExtraComponentBodyKey:  "UserPoolId",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "user_pool_id",
	},
	"GetDevice": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSCognitoIdentityProviderService.GetDevice",
		},
		Permission:             "GetDevice",
		IsExtra:                true,
		ExtraComponentBodyKey:  "DeviceKey",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "device_key",
	},
	"GetSigningCertificate": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSCognitoIdentityProviderService.GetSigningCertificate",
		},
		Permission:             "GetSigningCertificate",
		IsExtra:                true,
		ExtraComponentBodyKey:  "UserPoolId",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "user_pool_id",
	},
	"GetUICustomization": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSCognitoIdentityProviderService.GetUICustomization",
		},
		Permission:             "GetUICustomization",
		IsExtra:                true,
		ExtraComponentBodyKey:  "UserPoolId",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "user_pool_id",
	},
	"GetUser": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSCognitoIdentityProviderService.GetUser",
		},
		Permission:             "GetUser",
		IsExtra:                true,
		ExtraComponentBodyKey:  "AccessToken",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "access_token",
	},
	"GetUserPoolMfaConfig": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSCognitoIdentityProviderService.GetUserPoolMfaConfig",
		},
		Permission:             "GetUserPoolMfaConfig",
		IsExtra:                true,
		ExtraComponentBodyKey:  "UserPoolId",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "user_pool_id",
	},
	"ListDevices": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSCognitoIdentityProviderService.ListDevices",
		},
		Permission:             "ListDevices",
		IsExtra:                true,
		ExtraComponentBodyKey:  "AccessToken",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "access_token",
	},
	"ListGroups": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSCognitoIdentityProviderService.ListGroups",
		},
		Permission:             "ListGroups",
		IsExtra:                true,
		ExtraComponentBodyKey:  "UserPoolId",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "user_pool_id",
	},
	"ListIdentityProviders": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSCognitoIdentityProviderService.ListIdentityProviders",
		},
		Permission:             "ListIdentityProviders",
		IsExtra:                true,
		ExtraComponentBodyKey:  "UserPoolId",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "user_pool_id",
	},
	"ListResourceServers": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSCognitoIdentityProviderService.ListResourceServers",
		},
		Permission:             "ListResourceServers",
		IsExtra:                true,
		ExtraComponentBodyKey:  "UserPoolId",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "user_pool_id",
	},
	"ListTagsForResource": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSCognitoIdentityProviderService.ListTagsForResource",
		},
		Permission:             "ListTagsForResource",
		IsExtra:                true,
		ExtraComponentBodyKey:  "ResourceArn",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "resource_arn",
	},
	"ListUserPoolClients": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSCognitoIdentityProviderService.ListUserPoolClients",
		},
		Permission:             "ListUserPoolClients",
		IsExtra:                true,
		ExtraComponentBodyKey:  "UserPoolId",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "user_pool_id",
	},
	"ListUserPools": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSCognitoIdentityProviderService.ListUserPools",
		},
		Permission:             "ListUserPools",
		IsExtra:                true,
		ExtraComponentBodyKey:  "MaxResults",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "max_results",
	},
	"ListUsers": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSCognitoIdentityProviderService.ListUsers",
		},
		Permission:             "ListUsers",
		IsExtra:                true,
		ExtraComponentBodyKey:  "UserPoolId",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "user_pool_id",
	},
}
