package policy

import "github.com/securisec/cliam/shared"

var ProtonPolicies = []Service{
	{
		Method:   "POST",
		JsonData: `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1,
			aws_X_AMZ_TARGET:           "AwsProton20200720.GetAccountSettings",
		},
		Permission: "GetAccountSettings",
	},
	{
		Method:   "POST",
		JsonData: `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1,
			aws_X_AMZ_TARGET:           "AwsProton20200720.ListEnvironmentTemplates",
		},
		Permission: "ListEnvironmentTemplates",
	},
	{
		Method:   "POST",
		JsonData: `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1,
			aws_X_AMZ_TARGET:           "AwsProton20200720.ListEnvironments",
		},
		Permission: "ListEnvironments",
	},
	{
		Method:   "POST",
		JsonData: `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1,
			aws_X_AMZ_TARGET:           "AwsProton20200720.ListRepositories",
		},
		Permission: "ListRepositories",
	},
	{
		Method:   "POST",
		JsonData: `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1,
			aws_X_AMZ_TARGET:           "AwsProton20200720.ListServiceInstances",
		},
		Permission: "ListServiceInstances",
	},
	{
		Method:   "POST",
		JsonData: `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1,
			aws_X_AMZ_TARGET:           "AwsProton20200720.ListServiceTemplates",
		},
		Permission: "ListServiceTemplates",
	},
	{
		Method:   "POST",
		JsonData: `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1,
			aws_X_AMZ_TARGET:           "AwsProton20200720.ListServices",
		},
		Permission: "ListServices",
	},
	{
		Method:   "POST",
		JsonData: `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1,
			aws_X_AMZ_TARGET:           "AwsProton20200720.UpdateAccountSettings",
		},
		Permission: "UpdateAccountSettings",
	},
}