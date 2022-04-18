package policy

import "github.com/securisec/cliam/shared"

var ImageBuilderPolicies = []Service{
	{
		Method:        "POST",
		ServiceSuffix: "ListComponents",
		JsonData:      `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
		},
		Permission: "ListComponents",
	},
	{
		Method:        "POST",
		ServiceSuffix: "ListContainerRecipes",
		JsonData:      `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
		},
		Permission: "ListContainerRecipes",
	},
	{
		Method:        "POST",
		ServiceSuffix: "ListDistributionConfigurations",
		JsonData:      `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
		},
		Permission: "ListDistributionConfigurations",
	},
	{
		Method:        "POST",
		ServiceSuffix: "ListImagePipelines",
		JsonData:      `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
		},
		Permission: "ListImagePipelines",
	},
	{
		Method:        "POST",
		ServiceSuffix: "ListImageRecipes",
		JsonData:      `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
		},
		Permission: "ListImageRecipes",
	},
	{
		Method:        "POST",
		ServiceSuffix: "ListImages",
		JsonData:      `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
		},
		Permission: "ListImages",
	},
	{
		Method:        "POST",
		ServiceSuffix: "ListInfrastructureConfigurations",
		JsonData:      `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
		},
		Permission: "ListInfrastructureConfigurations",
	},
}
