package policy

import "github.com/securisec/cliam/shared"

var ImageBuilderPolicies = []Service{
	{
		Method:        "POST",
		ServiceSuffix: "ListComponents",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		},
		Permission: "ListComponents",
	},
	{
		Method:        "POST",
		ServiceSuffix: "ListContainerRecipes",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		},
		Permission: "ListContainerRecipes",
	},
	{
		Method:        "POST",
		ServiceSuffix: "ListDistributionConfigurations",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		},
		Permission: "ListDistributionConfigurations",
	},
	{
		Method:        "POST",
		ServiceSuffix: "ListImagePipelines",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		},
		Permission: "ListImagePipelines",
	},
	{
		Method:        "POST",
		ServiceSuffix: "ListImageRecipes",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		},
		Permission: "ListImageRecipes",
	},
	{
		Method:        "POST",
		ServiceSuffix: "ListImages",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		},
		Permission: "ListImages",
	},
	{
		Method:        "POST",
		ServiceSuffix: "ListInfrastructureConfigurations",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		},
		Permission: "ListInfrastructureConfigurations",
	},
}
