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

	// extra
	{
		ServiceSuffix:          "/GetComponent",
		Permission:             "GetComponent",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "component_build_version_arn",
	},
	{
		ServiceSuffix:          "/GetComponentPolicy",
		Permission:             "GetComponentPolicy",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "component_arn",
	},
	{
		ServiceSuffix:          "/GetContainerRecipe",
		Permission:             "GetContainerRecipe",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "container_recipe_arn",
	},
	{
		ServiceSuffix:          "/GetContainerRecipePolicy",
		Permission:             "GetContainerRecipePolicy",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "container_recipe_arn",
	},
	{
		ServiceSuffix:          "/GetDistributionConfiguration",
		Permission:             "GetDistributionConfiguration",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "distribution_configuration_arn",
	},
	{
		ServiceSuffix:          "/GetImage",
		Permission:             "GetImage",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "image_build_version_arn",
	},
	{
		ServiceSuffix:          "/GetImagePipeline",
		Permission:             "GetImagePipeline",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "image_pipeline_arn",
	},
	{
		ServiceSuffix:          "/GetImagePolicy",
		Permission:             "GetImagePolicy",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "image_arn",
	},
	{
		ServiceSuffix:          "/GetImageRecipe",
		Permission:             "GetImageRecipe",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "image_recipe_arn",
	},
	{
		ServiceSuffix:          "/GetImageRecipePolicy",
		Permission:             "GetImageRecipePolicy",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "image_recipe_arn",
	},
	{
		ServiceSuffix:          "/GetInfrastructureConfiguration",
		Permission:             "GetInfrastructureConfiguration",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "infrastructure_configuration_arn",
	},
	{
		Method: "POST",
		FormData: map[string]string{
			"Action":  "ListComponentBuildVersions",
			"Version": "2019-12-02",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "ListComponentBuildVersions",
		IsExtra:                true,
		ExtraComponentBodyKey:  "componentVersionArn",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "component_version_arn",
	},
	{
		Method: "POST",
		FormData: map[string]string{
			"Action":  "ListImageBuildVersions",
			"Version": "2019-12-02",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "ListImageBuildVersions",
		IsExtra:                true,
		ExtraComponentBodyKey:  "imageVersionArn",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "image_version_arn",
	},
	{
		Method: "POST",
		FormData: map[string]string{
			"Action":  "ListImagePackages",
			"Version": "2019-12-02",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "ListImagePackages",
		IsExtra:                true,
		ExtraComponentBodyKey:  "imageBuildVersionArn",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "image_build_version_arn",
	},
	{
		Method: "POST",
		FormData: map[string]string{
			"Action":  "ListImagePipelineImages",
			"Version": "2019-12-02",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "ListImagePipelineImages",
		IsExtra:                true,
		ExtraComponentBodyKey:  "imagePipelineArn",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "image_pipeline_arn",
	},
	{
		ServiceSuffix:          "/tags/{{.resource_arn}}",
		Permission:             "ListTagsForResource",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "resource_arn",
	},
}
