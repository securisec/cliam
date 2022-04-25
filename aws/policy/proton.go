package policy

import "github.com/securisec/cliam/shared"

var ProtonPolicies = []Service{
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_0,
			aws_X_AMZ_TARGET:           "AwsProton20200720.GetAccountSettings",
		},
		Permission: "GetAccountSettings",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_0,
			aws_X_AMZ_TARGET:           "AwsProton20200720.ListEnvironmentTemplates",
		},
		Permission: "ListEnvironmentTemplates",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_0,
			aws_X_AMZ_TARGET:           "AwsProton20200720.ListEnvironments",
		},
		Permission: "ListEnvironments",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_0,
			aws_X_AMZ_TARGET:           "AwsProton20200720.ListRepositories",
		},
		Permission: "ListRepositories",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_0,
			aws_X_AMZ_TARGET:           "AwsProton20200720.ListServiceInstances",
		},
		Permission: "ListServiceInstances",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_0,
			aws_X_AMZ_TARGET:           "AwsProton20200720.ListServiceTemplates",
		},
		Permission: "ListServiceTemplates",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_0,
			aws_X_AMZ_TARGET:           "AwsProton20200720.ListServices",
		},
		Permission: "ListServices",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_0,
			aws_X_AMZ_TARGET:           "AwsProton20200720.UpdateAccountSettings",
		},
		Permission: "UpdateAccountSettings",
	},

	// extra
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_0,
			aws_X_AMZ_TARGET:           "AwsProton20200720.GetEnvironment",
		},
		Permission:             "GetEnvironment",
		IsExtra:                true,
		ExtraComponentBodyKey:  "name",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "name",
	},
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_0,
			aws_X_AMZ_TARGET:           "AwsProton20200720.GetEnvironmentAccountConnection",
		},
		Permission:             "GetEnvironmentAccountConnection",
		IsExtra:                true,
		ExtraComponentBodyKey:  "id",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "id",
	},
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_0,
			aws_X_AMZ_TARGET:           "AwsProton20200720.GetEnvironmentTemplate",
		},
		Permission:             "GetEnvironmentTemplate",
		IsExtra:                true,
		ExtraComponentBodyKey:  "name",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "name",
	},
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_0,
			aws_X_AMZ_TARGET:           "AwsProton20200720.GetService",
		},
		Permission:             "GetService",
		IsExtra:                true,
		ExtraComponentBodyKey:  "name",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "name",
	},
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_0,
			aws_X_AMZ_TARGET:           "AwsProton20200720.GetServiceTemplate",
		},
		Permission:             "GetServiceTemplate",
		IsExtra:                true,
		ExtraComponentBodyKey:  "name",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "name",
	},
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_0,
			aws_X_AMZ_TARGET:           "AwsProton20200720.ListEnvironmentAccountConnections",
		},
		Permission:             "ListEnvironmentAccountConnections",
		IsExtra:                true,
		ExtraComponentBodyKey:  "requestedBy",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "requested_by",
	},
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_0,
			aws_X_AMZ_TARGET:           "AwsProton20200720.ListEnvironmentOutputs",
		},
		Permission:             "ListEnvironmentOutputs",
		IsExtra:                true,
		ExtraComponentBodyKey:  "environmentName",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "environment_name",
	},
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_0,
			aws_X_AMZ_TARGET:           "AwsProton20200720.ListEnvironmentProvisionedResources",
		},
		Permission:             "ListEnvironmentProvisionedResources",
		IsExtra:                true,
		ExtraComponentBodyKey:  "environmentName",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "environment_name",
	},
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_0,
			aws_X_AMZ_TARGET:           "AwsProton20200720.ListEnvironmentTemplateVersions",
		},
		Permission:             "ListEnvironmentTemplateVersions",
		IsExtra:                true,
		ExtraComponentBodyKey:  "templateName",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "template_name",
	},
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_0,
			aws_X_AMZ_TARGET:           "AwsProton20200720.ListServicePipelineOutputs",
		},
		Permission:             "ListServicePipelineOutputs",
		IsExtra:                true,
		ExtraComponentBodyKey:  "serviceName",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "service_name",
	},
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_0,
			aws_X_AMZ_TARGET:           "AwsProton20200720.ListServicePipelineProvisionedResources",
		},
		Permission:             "ListServicePipelineProvisionedResources",
		IsExtra:                true,
		ExtraComponentBodyKey:  "serviceName",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "service_name",
	},
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_0,
			aws_X_AMZ_TARGET:           "AwsProton20200720.ListServiceTemplateVersions",
		},
		Permission:             "ListServiceTemplateVersions",
		IsExtra:                true,
		ExtraComponentBodyKey:  "templateName",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "template_name",
	},
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_0,
			aws_X_AMZ_TARGET:           "AwsProton20200720.ListTagsForResource",
		},
		Permission:             "ListTagsForResource",
		IsExtra:                true,
		ExtraComponentBodyKey:  "resourceArn",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "resource_arn",
	},
}
