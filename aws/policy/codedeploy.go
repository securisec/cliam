package policy

import "github.com/securisec/cliam/shared"

var CodeDeployPolicies = map[string]Service{
	"BatchGetDeploymentTargets": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "CodeDeploy_20141006.BatchGetDeploymentTargets",
		},
		Permission: "BatchGetDeploymentTargets",
	},
	"ContinueDeployment": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "CodeDeploy_20141006.ContinueDeployment",
		},
		Permission: "ContinueDeployment",
	},
	"DeleteGitHubAccountToken": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "CodeDeploy_20141006.DeleteGitHubAccountToken",
		},
		Permission: "DeleteGitHubAccountToken",
	},
	"DeleteResourcesByExternalId": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "CodeDeploy_20141006.DeleteResourcesByExternalId",
		},
		Permission: "DeleteResourcesByExternalId",
	},
	"GetDeploymentTarget": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "CodeDeploy_20141006.GetDeploymentTarget",
		},
		Permission: "GetDeploymentTarget",
	},
	"ListApplications": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "CodeDeploy_20141006.ListApplications",
		},
		Permission: "ListApplications",
	},
	"ListDeploymentConfigs": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "CodeDeploy_20141006.ListDeploymentConfigs",
		},
		Permission: "ListDeploymentConfigs",
	},
	"ListDeploymentTargets": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "CodeDeploy_20141006.ListDeploymentTargets",
		},
		Permission: "ListDeploymentTargets",
	},
	"ListDeployments": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "CodeDeploy_20141006.ListDeployments",
		},
		Permission: "ListDeployments",
	},
	"ListGitHubAccountTokenNames": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "CodeDeploy_20141006.ListGitHubAccountTokenNames",
		},
		Permission: "ListGitHubAccountTokenNames",
	},
	"ListOnPremisesInstances": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "CodeDeploy_20141006.ListOnPremisesInstances",
		},
		Permission: "ListOnPremisesInstances",
	},
	"PutLifecycleEventHookExecutionStatus": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "CodeDeploy_20141006.PutLifecycleEventHookExecutionStatus",
		},
		Permission: "PutLifecycleEventHookExecutionStatus",
	},
	"SkipWaitTimeForInstanceTermination": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "CodeDeploy_20141006.SkipWaitTimeForInstanceTermination",
		},
		Permission: "SkipWaitTimeForInstanceTermination",
	},
	"UpdateApplication": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "CodeDeploy_20141006.UpdateApplication",
		},
		Permission: "UpdateApplication",
	},

	// extra
	"GetApplication": {
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
	"GetDeployment": {
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
	"GetDeploymentConfig": {
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
	"GetOnPremisesInstance": {
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
	"ListApplicationRevisions": {
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
	"ListDeploymentGroups": {
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
	"ListDeploymentInstances": {
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
	"ListTagsForResource": {
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
