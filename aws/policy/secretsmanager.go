package policy

import "github.com/securisec/cliam/shared"

var SecretsManagerPolicies = []Service{
	{
		Method:   "POST",
		JsonData: `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
			aws_X_AMZ_TARGET:           "secretsmanager.GetRandomPassword",
		},
		Permission: "GetRandomPassword",
	},
	{
		Method:   "POST",
		JsonData: `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
			aws_X_AMZ_TARGET:           "secretsmanager.ListSecrets",
		},
		Permission: "ListSecrets",
	},
}
