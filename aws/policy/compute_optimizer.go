package policy

import "github.com/securisec/cliam/shared"

var ComputeOptimizerPolicies = []Service{
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "ComputeOptimizerService.DescribeRecommendationExportJobs",
		},
		Permission: "DescribeRecommendationExportJobs",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "ComputeOptimizerService.GetAutoScalingGroupRecommendations",
		},
		Permission: "GetAutoScalingGroupRecommendations",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "ComputeOptimizerService.GetEBSVolumeRecommendations",
		},
		Permission: "GetEBSVolumeRecommendations",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "ComputeOptimizerService.GetEC2InstanceRecommendations",
		},
		Permission: "GetEC2InstanceRecommendations",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "ComputeOptimizerService.GetEnrollmentStatus",
		},
		Permission: "GetEnrollmentStatus",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "ComputeOptimizerService.GetEnrollmentStatusesForOrganization",
		},
		Permission: "GetEnrollmentStatusesForOrganization",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "ComputeOptimizerService.GetLambdaFunctionRecommendations",
		},
		Permission: "GetLambdaFunctionRecommendations",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "ComputeOptimizerService.GetRecommendationSummaries",
		},
		Permission: "GetRecommendationSummaries",
	},

	// extra
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_0,
			aws_X_AMZ_TARGET:           "ComputeOptimizerService.GetEffectiveRecommendationPreferences",
		},
		Permission:             "GetEffectiveRecommendationPreferences",
		IsExtra:                true,
		ExtraComponentBodyKey:  "resourceArn",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "resource_arn",
	},
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_0,
			aws_X_AMZ_TARGET:           "ComputeOptimizerService.GetRecommendationPreferences",
		},
		Permission:             "GetRecommendationPreferences",
		IsExtra:                true,
		ExtraComponentBodyKey:  "resourceType",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "resource_type",
	},
}
