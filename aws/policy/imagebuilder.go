package policy

import "github.com/securisec/cliam/shared"

var ImageBuilderPolicies = []Service{
	{
		Method:        "POST",
		ServiceSuffix: "ListComponents",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
		},
		Permission: "ListComponents",
	},
	{
		Method:        "POST",
		ServiceSuffix: "ListContainerRecipes",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
		},
		Permission: "ListContainerRecipes",
	},
	{
		Method:        "POST",
		ServiceSuffix: "ListDistributionConfigurations",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
		},
		Permission: "ListDistributionConfigurations",
	},
	{
		Method:        "POST",
		ServiceSuffix: "ListImagePipelines",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
		},
		Permission: "ListImagePipelines",
	},
	{
		Method:        "POST",
		ServiceSuffix: "ListImageRecipes",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
		},
		Permission: "ListImageRecipes",
	},
	{
		Method:        "POST",
		ServiceSuffix: "ListImages",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
		},
		Permission: "ListImages",
	},
	{
		Method:        "POST",
		ServiceSuffix: "ListInfrastructureConfigurations",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
		},
		Permission: "ListInfrastructureConfigurations",
	},
}
