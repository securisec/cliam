package policy

import "github.com/securisec/cliam/shared"

var AppRunnerPolicies = map[string]Service{
	"ListAutoScalingConfigurations": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: "application/x-amz-json-1.0",
			aws_X_AMZ_TARGET:           "AppRunner.ListAutoScalingConfigurations",
		},
		Permission: "ListAutoScalingConfigurations",
	},
	"ListConnections": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: "application/x-amz-json-1.0",
			aws_X_AMZ_TARGET:           "AppRunner.ListConnections",
		},
		Permission: "ListConnections",
	},
	"ListObservabilityConfigurations": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: "application/x-amz-json-1.0",
			aws_X_AMZ_TARGET:           "AppRunner.ListObservabilityConfigurations",
		},
		Permission: "ListObservabilityConfigurations",
	},
	"ListServices": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: "application/x-amz-json-1.0",
			aws_X_AMZ_TARGET:           "AppRunner.ListServices",
		},
		Permission: "ListServices",
	},
	"ListVpcConnectors": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: "application/x-amz-json-1.0",
			aws_X_AMZ_TARGET:           "AppRunner.ListVpcConnectors",
		},
		Permission: "ListVpcConnectors",
	},

	// extra
	"DescribeAutoScalingConfiguration": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_0,
			aws_X_AMZ_TARGET:           "AppRunner.DescribeAutoScalingConfiguration",
		},
		Permission:             "DescribeAutoScalingConfiguration",
		IsExtra:                true,
		ExtraComponentBodyKey:  "AutoScalingConfigurationArn",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "auto_scaling_configuration_arn",
	},
	"DescribeCustomDomains": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_0,
			aws_X_AMZ_TARGET:           "AppRunner.DescribeCustomDomains",
		},
		Permission:             "DescribeCustomDomains",
		IsExtra:                true,
		ExtraComponentBodyKey:  "ServiceArn",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "service_arn",
	},
	"DescribeObservabilityConfiguration": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_0,
			aws_X_AMZ_TARGET:           "AppRunner.DescribeObservabilityConfiguration",
		},
		Permission:             "DescribeObservabilityConfiguration",
		IsExtra:                true,
		ExtraComponentBodyKey:  "ObservabilityConfigurationArn",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "observability_configuration_arn",
	},
	"DescribeService": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_0,
			aws_X_AMZ_TARGET:           "AppRunner.DescribeService",
		},
		Permission:             "DescribeService",
		IsExtra:                true,
		ExtraComponentBodyKey:  "ServiceArn",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "service_arn",
	},
	"DescribeVpcConnector": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_0,
			aws_X_AMZ_TARGET:           "AppRunner.DescribeVpcConnector",
		},
		Permission:             "DescribeVpcConnector",
		IsExtra:                true,
		ExtraComponentBodyKey:  "VpcConnectorArn",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "vpc_connector_arn",
	},
	"ListOperations": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_0,
			aws_X_AMZ_TARGET:           "AppRunner.ListOperations",
		},
		Permission:             "ListOperations",
		IsExtra:                true,
		ExtraComponentBodyKey:  "ServiceArn",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "service_arn",
	},
	"ListTagsForResource": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_0,
			aws_X_AMZ_TARGET:           "AppRunner.ListTagsForResource",
		},
		Permission:             "ListTagsForResource",
		IsExtra:                true,
		ExtraComponentBodyKey:  "ResourceArn",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "resource_arn",
	},
}
