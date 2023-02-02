package policy

import "github.com/securisec/cliam/shared"

// ProtonPolicies policy
var ProtonPolicies = map[string]Service{
	"GetAccountSettings": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AwsProton20200720.GetAccountSettings",
		},
		Permission: "GetAccountSettings",
	},
	"ListComponents": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AwsProton20200720.ListComponents",
		},
		Permission: "ListComponents",
	},
	"ListEnvironmentTemplates": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AwsProton20200720.ListEnvironmentTemplates",
		},
		Permission: "ListEnvironmentTemplates",
	},
	"ListEnvironments": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AwsProton20200720.ListEnvironments",
		},
		Permission: "ListEnvironments",
	},
	"ListRepositories": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AwsProton20200720.ListRepositories",
		},
		Permission: "ListRepositories",
	},
	"ListServiceInstances": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AwsProton20200720.ListServiceInstances",
		},
		Permission: "ListServiceInstances",
	},
	"ListServiceTemplates": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AwsProton20200720.ListServiceTemplates",
		},
		Permission: "ListServiceTemplates",
	},
	"ListServices": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AwsProton20200720.ListServices",
		},
		Permission: "ListServices",
	},
	"UpdateAccountSettings": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AwsProton20200720.UpdateAccountSettings",
		},
		Permission: "UpdateAccountSettings",
	},

	// extra
	"GetComponent": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_0,
			aws_X_AMZ_TARGET:           "AwsProton20200720.GetComponent",
		},
		Permission:             "GetComponent",
		IsExtra:                true,
		ExtraComponentBodyKey:  "name",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "name",
	},
	"GetEnvironment": {
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
	"GetEnvironmentAccountConnection": {
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
	"GetEnvironmentTemplate": {
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
	"GetService": {
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
	"GetServiceTemplate": {
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
	"ListComponentOutputs": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_0,
			aws_X_AMZ_TARGET:           "AwsProton20200720.ListComponentOutputs",
		},
		Permission:             "ListComponentOutputs",
		IsExtra:                true,
		ExtraComponentBodyKey:  "componentName",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "component_name",
	},
	"ListComponentProvisionedResources": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_0,
			aws_X_AMZ_TARGET:           "AwsProton20200720.ListComponentProvisionedResources",
		},
		Permission:             "ListComponentProvisionedResources",
		IsExtra:                true,
		ExtraComponentBodyKey:  "componentName",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "component_name",
	},
	"ListEnvironmentAccountConnections": {
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
	"ListEnvironmentOutputs": {
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
	"ListEnvironmentProvisionedResources": {
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
	"ListEnvironmentTemplateVersions": {
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
	"ListServicePipelineOutputs": {
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
	"ListServicePipelineProvisionedResources": {
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
	"ListServiceTemplateVersions": {
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
	"ListTagsForResource": {
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
