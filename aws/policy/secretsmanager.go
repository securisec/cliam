package policy

import "github.com/securisec/cliam/shared"

var SecretsManagerPolicies = []Service{
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "secretsmanager.GetRandomPassword",
		},
		Permission: "GetRandomPassword",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "secretsmanager.ListSecrets",
		},
		Permission: "ListSecrets",
	},

	// extra
	{
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
	{
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
	{
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
	{
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
