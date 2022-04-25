package policy

import "github.com/securisec/cliam/shared"

var CodeDeployPolicies = []Service{
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "CodeDeploy_20141006.BatchGetDeploymentTargets",
		},
		Permission: "BatchGetDeploymentTargets",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "CodeDeploy_20141006.ContinueDeployment",
		},
		Permission: "ContinueDeployment",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "CodeDeploy_20141006.DeleteGitHubAccountToken",
		},
		Permission: "DeleteGitHubAccountToken",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "CodeDeploy_20141006.DeleteResourcesByExternalId",
		},
		Permission: "DeleteResourcesByExternalId",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "CodeDeploy_20141006.GetDeploymentTarget",
		},
		Permission: "GetDeploymentTarget",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "CodeDeploy_20141006.ListApplications",
		},
		Permission: "ListApplications",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "CodeDeploy_20141006.ListDeploymentConfigs",
		},
		Permission: "ListDeploymentConfigs",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "CodeDeploy_20141006.ListDeploymentTargets",
		},
		Permission: "ListDeploymentTargets",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "CodeDeploy_20141006.ListDeployments",
		},
		Permission: "ListDeployments",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "CodeDeploy_20141006.ListGitHubAccountTokenNames",
		},
		Permission: "ListGitHubAccountTokenNames",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "CodeDeploy_20141006.ListOnPremisesInstances",
		},
		Permission: "ListOnPremisesInstances",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "CodeDeploy_20141006.PutLifecycleEventHookExecutionStatus",
		},
		Permission: "PutLifecycleEventHookExecutionStatus",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "CodeDeploy_20141006.SkipWaitTimeForInstanceTermination",
		},
		Permission: "SkipWaitTimeForInstanceTermination",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "CodeDeploy_20141006.UpdateApplication",
		},
		Permission: "UpdateApplication",
	},

	// extra
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "CodeDeploy_20141006.GetApplication",
		},
		Permission:             "GetApplication",
		IsExtra:                true,
		ExtraComponentBodyKey:  "applicationName",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "application_name",
	},
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "CodeDeploy_20141006.GetDeployment",
		},
		Permission:             "GetDeployment",
		IsExtra:                true,
		ExtraComponentBodyKey:  "deploymentId",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "deployment_id",
	},
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "CodeDeploy_20141006.GetDeploymentConfig",
		},
		Permission:             "GetDeploymentConfig",
		IsExtra:                true,
		ExtraComponentBodyKey:  "deploymentConfigName",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "deployment_config_name",
	},
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "CodeDeploy_20141006.GetOnPremisesInstance",
		},
		Permission:             "GetOnPremisesInstance",
		IsExtra:                true,
		ExtraComponentBodyKey:  "instanceName",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "instance_name",
	},
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "CodeDeploy_20141006.ListApplicationRevisions",
		},
		Permission:             "ListApplicationRevisions",
		IsExtra:                true,
		ExtraComponentBodyKey:  "applicationName",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "application_name",
	},
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "CodeDeploy_20141006.ListDeploymentGroups",
		},
		Permission:             "ListDeploymentGroups",
		IsExtra:                true,
		ExtraComponentBodyKey:  "applicationName",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "application_name",
	},
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "CodeDeploy_20141006.ListDeploymentInstances",
		},
		Permission:             "ListDeploymentInstances",
		IsExtra:                true,
		ExtraComponentBodyKey:  "deploymentId",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "deployment_id",
	},
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "CodeDeploy_20141006.ListTagsForResource",
		},
		Permission:             "ListTagsForResource",
		IsExtra:                true,
		ExtraComponentBodyKey:  "ResourceArn",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "resource_arn",
	},
}
