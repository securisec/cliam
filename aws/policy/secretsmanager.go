package policy

import "github.com/securisec/cliam/shared"

var SecretsManagerPolicies = map[string]Service{
	"GetRandomPassword": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "secretsmanager.GetRandomPassword",
		},
		Permission: "GetRandomPassword",
	},
	"ListSecrets": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "secretsmanager.ListSecrets",
		},
		Permission: "ListSecrets",
	},

	// extra
	"DescribeSecret": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "secretsmanager.DescribeSecret",
		},
		Permission:             "DescribeSecret",
		IsExtra:                true,
		ExtraComponentBodyKey:  "SecretId",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "secret_id",
	},
	"GetResourcePolicy": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "secretsmanager.GetResourcePolicy",
		},
		Permission:             "GetResourcePolicy",
		IsExtra:                true,
		ExtraComponentBodyKey:  "SecretId",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "secret_id",
	},
	"GetSecretValue": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "secretsmanager.GetSecretValue",
		},
		Permission:             "GetSecretValue",
		IsExtra:                true,
		ExtraComponentBodyKey:  "SecretId",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "secret_id",
	},
	"ListSecretVersionIds": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "secretsmanager.ListSecretVersionIds",
		},
		Permission:             "ListSecretVersionIds",
		IsExtra:                true,
		ExtraComponentBodyKey:  "SecretId",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "secret_id",
	},
}
