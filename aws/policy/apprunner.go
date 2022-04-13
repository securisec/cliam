package policy

import "github.com/securisec/cliam/shared"

var AppRunnerPolicies = []Service{
	{
		Method:   "POST",
		JsonData: `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: "application/x-amz-json-1.0",
			aws_X_AMZ_TARGET:           "AppRunner.ListAutoScalingConfigurations",
		},
		Permission: "ListAutoScalingConfigurations",
	},
	{
		Method:   "POST",
		JsonData: `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: "application/x-amz-json-1.0",
			aws_X_AMZ_TARGET:           "AppRunner.ListConnections",
		},
		Permission: "ListConnections",
	},
	{
		Method:   "POST",
		JsonData: `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: "application/x-amz-json-1.0",
			aws_X_AMZ_TARGET:           "AppRunner.ListObservabilityConfigurations",
		},
		Permission: "ListObservabilityConfigurations",
	},
	{
		Method:   "POST",
		JsonData: `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: "application/x-amz-json-1.0",
			aws_X_AMZ_TARGET:           "AppRunner.ListServices",
		},
		Permission: "ListServices",
	},
	{
		Method:   "POST",
		JsonData: `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: "application/x-amz-json-1.0",
			aws_X_AMZ_TARGET:           "AppRunner.ListVpcConnectors",
		},
		Permission: "ListVpcConnectors",
	},
}
