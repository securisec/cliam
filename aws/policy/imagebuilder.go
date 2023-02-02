package policy

import "github.com/securisec/cliam/shared"

// ImageBuilderPolicies policy
var ImageBuilderPolicies = map[string]Service{
	"ListComponents": {
		Method:        "POST",
		ServiceSuffix: "ListComponents",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		},
		Permission: "ListComponents",
	},
	"ListContainerRecipes": {
		Method:        "POST",
		ServiceSuffix: "ListContainerRecipes",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		},
		Permission: "ListContainerRecipes",
	},
	"ListDistributionConfigurations": {
		Method:        "POST",
		ServiceSuffix: "ListDistributionConfigurations",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		},
		Permission: "ListDistributionConfigurations",
	},
	"ListImagePipelines": {
		Method:        "POST",
		ServiceSuffix: "ListImagePipelines",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		},
		Permission: "ListImagePipelines",
	},
	"ListImageRecipes": {
		Method:        "POST",
		ServiceSuffix: "ListImageRecipes",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		},
		Permission: "ListImageRecipes",
	},
	"ListImages": {
		Method:        "POST",
		ServiceSuffix: "ListImages",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		},
		Permission: "ListImages",
	},
	"ListInfrastructureConfigurations": {
		Method:        "POST",
		ServiceSuffix: "ListInfrastructureConfigurations",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		},
		Permission: "ListInfrastructureConfigurations",
	},

	// extra
	"GetComponent": {
		ServiceSuffix:          "/GetComponent",
		Permission:             "GetComponent",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "component_build_version_arn",
	},
	"GetComponentPolicy": {
		ServiceSuffix:          "/GetComponentPolicy",
		Permission:             "GetComponentPolicy",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "component_arn",
	},
	"GetContainerRecipe": {
		ServiceSuffix:          "/GetContainerRecipe",
		Permission:             "GetContainerRecipe",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "container_recipe_arn",
	},
	"GetContainerRecipePolicy": {
		ServiceSuffix:          "/GetContainerRecipePolicy",
		Permission:             "GetContainerRecipePolicy",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "container_recipe_arn",
	},
	"GetDistributionConfiguration": {
		ServiceSuffix:          "/GetDistributionConfiguration",
		Permission:             "GetDistributionConfiguration",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "distribution_configuration_arn",
	},
	"GetImage": {
		ServiceSuffix:          "/GetImage",
		Permission:             "GetImage",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "image_build_version_arn",
	},
	"GetImagePipeline": {
		ServiceSuffix:          "/GetImagePipeline",
		Permission:             "GetImagePipeline",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "image_pipeline_arn",
	},
	"GetImagePolicy": {
		ServiceSuffix:          "/GetImagePolicy",
		Permission:             "GetImagePolicy",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "image_arn",
	},
	"GetImageRecipe": {
		ServiceSuffix:          "/GetImageRecipe",
		Permission:             "GetImageRecipe",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "image_recipe_arn",
	},
	"GetImageRecipePolicy": {
		ServiceSuffix:          "/GetImageRecipePolicy",
		Permission:             "GetImageRecipePolicy",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "image_recipe_arn",
	},
	"GetInfrastructureConfiguration": {
		ServiceSuffix:          "/GetInfrastructureConfiguration",
		Permission:             "GetInfrastructureConfiguration",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "infrastructure_configuration_arn",
	},
	"ListComponentBuildVersions": {
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
	"ListImageBuildVersions": {
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
	"ListImagePackages": {
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
	"ListImagePipelineImages": {
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
	"ListTagsForResource": {
		ServiceSuffix:          "/tags/{{.resource_arn}}",
		Permission:             "ListTagsForResource",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "resource_arn",
	},
}
