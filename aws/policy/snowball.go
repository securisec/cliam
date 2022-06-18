package policy

import "github.com/securisec/cliam/shared"

var SnowballPolicies = map[string]Service{
	"CreateJob": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSIESnowballJobManagementService.CreateJob",
		},
		Permission: "CreateJob",
	},
	"DescribeAddresses": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSIESnowballJobManagementService.DescribeAddresses",
		},
		Permission: "DescribeAddresses",
	},
	"GetSnowballUsage": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSIESnowballJobManagementService.GetSnowballUsage",
		},
		Permission: "GetSnowballUsage",
	},
	"ListClusters": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSIESnowballJobManagementService.ListClusters",
		},
		Permission: "ListClusters",
	},
	"ListCompatibleImages": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSIESnowballJobManagementService.ListCompatibleImages",
		},
		Permission: "ListCompatibleImages",
	},
	"ListJobs": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSIESnowballJobManagementService.ListJobs",
		},
		Permission: "ListJobs",
	},
	"ListLongTermPricing": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSIESnowballJobManagementService.ListLongTermPricing",
		},
		Permission: "ListLongTermPricing",
	},

	// extra
	"DescribeAddress": {
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
	"DescribeCluster": {
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
	"DescribeJob": {
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
	"DescribeReturnShippingLabel": {
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
	"GetJobManifest": {
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
	"GetJobUnlockCode": {
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
	"GetSoftwareUpdates": {
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
	"ListClusterJobs": {
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
