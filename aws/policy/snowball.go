package policy

import "github.com/securisec/cliam/shared"

var SnowballPolicies = []Service{
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSIESnowballJobManagementService.CreateJob",
		},
		Permission: "CreateJob",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSIESnowballJobManagementService.DescribeAddresses",
		},
		Permission: "DescribeAddresses",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSIESnowballJobManagementService.GetSnowballUsage",
		},
		Permission: "GetSnowballUsage",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSIESnowballJobManagementService.ListClusters",
		},
		Permission: "ListClusters",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSIESnowballJobManagementService.ListCompatibleImages",
		},
		Permission: "ListCompatibleImages",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSIESnowballJobManagementService.ListJobs",
		},
		Permission: "ListJobs",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSIESnowballJobManagementService.ListLongTermPricing",
		},
		Permission: "ListLongTermPricing",
	},

	// extra
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSIESnowballJobManagementService.DescribeAddress",
		},
		Permission:             "DescribeAddress",
		IsExtra:                true,
		ExtraComponentBodyKey:  "AddressId",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "address_id",
	},
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSIESnowballJobManagementService.DescribeCluster",
		},
		Permission:             "DescribeCluster",
		IsExtra:                true,
		ExtraComponentBodyKey:  "ClusterId",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "cluster_id",
	},
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSIESnowballJobManagementService.DescribeJob",
		},
		Permission:             "DescribeJob",
		IsExtra:                true,
		ExtraComponentBodyKey:  "JobId",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "job_id",
	},
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSIESnowballJobManagementService.DescribeReturnShippingLabel",
		},
		Permission:             "DescribeReturnShippingLabel",
		IsExtra:                true,
		ExtraComponentBodyKey:  "JobId",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "job_id",
	},
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSIESnowballJobManagementService.GetJobManifest",
		},
		Permission:             "GetJobManifest",
		IsExtra:                true,
		ExtraComponentBodyKey:  "JobId",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "job_id",
	},
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSIESnowballJobManagementService.GetJobUnlockCode",
		},
		Permission:             "GetJobUnlockCode",
		IsExtra:                true,
		ExtraComponentBodyKey:  "JobId",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "job_id",
	},
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSIESnowballJobManagementService.GetSoftwareUpdates",
		},
		Permission:             "GetSoftwareUpdates",
		IsExtra:                true,
		ExtraComponentBodyKey:  "JobId",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "job_id",
	},
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSIESnowballJobManagementService.ListClusterJobs",
		},
		Permission:             "ListClusterJobs",
		IsExtra:                true,
		ExtraComponentBodyKey:  "ClusterId",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "cluster_id",
	},
}
