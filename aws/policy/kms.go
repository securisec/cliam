package policy

import "github.com/securisec/cliam/shared"

var KMSPolicies = map[string]Service{
	"DescribeCustomKeyStores": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "TrentService.DescribeCustomKeyStores",
		},
		Permission: "DescribeCustomKeyStores",
	},
	"ListAliases": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "TrentService.ListAliases",
		},
		Permission: "ListAliases",
	},
	"ListKeys": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "TrentService.ListKeys",
		},
		Permission: "ListKeys",
	},

	// extra
	"DescribeKey": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "TrentService.DescribeKey",
		},
		Permission:             "DescribeKey",
		IsExtra:                true,
		ExtraComponentBodyKey:  "KeyId",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "key_id",
	},
	"GetKeyRotationStatus": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "TrentService.GetKeyRotationStatus",
		},
		Permission:             "GetKeyRotationStatus",
		IsExtra:                true,
		ExtraComponentBodyKey:  "KeyId",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "key_id",
	},
	"GetPublicKey": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "TrentService.GetPublicKey",
		},
		Permission:             "GetPublicKey",
		IsExtra:                true,
		ExtraComponentBodyKey:  "KeyId",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "key_id",
	},
	"ListGrants": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "TrentService.ListGrants",
		},
		Permission:             "ListGrants",
		IsExtra:                true,
		ExtraComponentBodyKey:  "KeyId",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "key_id",
	},
	"ListKeyPolicies": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "TrentService.ListKeyPolicies",
		},
		Permission:             "ListKeyPolicies",
		IsExtra:                true,
		ExtraComponentBodyKey:  "KeyId",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "key_id",
	},
	"ListResourceTags": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "TrentService.ListResourceTags",
		},
		Permission:             "ListResourceTags",
		IsExtra:                true,
		ExtraComponentBodyKey:  "KeyId",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "key_id",
	},
	"ListRetirableGrants": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "TrentService.ListRetirableGrants",
		},
		Permission:             "ListRetirableGrants",
		IsExtra:                true,
		ExtraComponentBodyKey:  "RetiringPrincipal",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "retiring_principal",
	},
}
