package policy

import "github.com/securisec/cliam/shared"

var SnowballPolicies = []Service{
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
			aws_X_AMZ_TARGET:           "AWSIESnowballJobManagementService.CreateJob",
		},
		Permission: "CreateJob",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
			aws_X_AMZ_TARGET:           "AWSIESnowballJobManagementService.DescribeAddresses",
		},
		Permission: "DescribeAddresses",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
			aws_X_AMZ_TARGET:           "AWSIESnowballJobManagementService.GetSnowballUsage",
		},
		Permission: "GetSnowballUsage",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
			aws_X_AMZ_TARGET:           "AWSIESnowballJobManagementService.ListClusters",
		},
		Permission: "ListClusters",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
			aws_X_AMZ_TARGET:           "AWSIESnowballJobManagementService.ListCompatibleImages",
		},
		Permission: "ListCompatibleImages",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
			aws_X_AMZ_TARGET:           "AWSIESnowballJobManagementService.ListJobs",
		},
		Permission: "ListJobs",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
			aws_X_AMZ_TARGET:           "AWSIESnowballJobManagementService.ListLongTermPricing",
		},
		Permission: "ListLongTermPricing",
	},
}
