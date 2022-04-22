package policy

import "github.com/securisec/cliam/shared"

var ComputeOptimizerPolicies = []Service{
	{
		Method:   "POST",
		JsonData: `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
			aws_X_AMZ_TARGET:           "ComputeOptimizerService.DescribeRecommendationExportJobs",
		},
		Permission: "DescribeRecommendationExportJobs",
	},
	{
		Method:   "POST",
		JsonData: `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
			aws_X_AMZ_TARGET:           "ComputeOptimizerService.GetAutoScalingGroupRecommendations",
		},
		Permission: "GetAutoScalingGroupRecommendations",
	},
	{
		Method:   "POST",
		JsonData: `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
			aws_X_AMZ_TARGET:           "ComputeOptimizerService.GetEBSVolumeRecommendations",
		},
		Permission: "GetEBSVolumeRecommendations",
	},
	{
		Method:   "POST",
		JsonData: `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
			aws_X_AMZ_TARGET:           "ComputeOptimizerService.GetEC2InstanceRecommendations",
		},
		Permission: "GetEC2InstanceRecommendations",
	},
	{
		Method:   "POST",
		JsonData: `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
			aws_X_AMZ_TARGET:           "ComputeOptimizerService.GetEnrollmentStatus",
		},
		Permission: "GetEnrollmentStatus",
	},
	{
		Method:   "POST",
		JsonData: `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
			aws_X_AMZ_TARGET:           "ComputeOptimizerService.GetEnrollmentStatusesForOrganization",
		},
		Permission: "GetEnrollmentStatusesForOrganization",
	},
	{
		Method:   "POST",
		JsonData: `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
			aws_X_AMZ_TARGET:           "ComputeOptimizerService.GetLambdaFunctionRecommendations",
		},
		Permission: "GetLambdaFunctionRecommendations",
	},
	{
		Method:   "POST",
		JsonData: `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
			aws_X_AMZ_TARGET:           "ComputeOptimizerService.GetRecommendationSummaries",
		},
		Permission: "GetRecommendationSummaries",
	},
}