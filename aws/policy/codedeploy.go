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
}
