package policy

import "github.com/securisec/cliam/shared"

var FSXPolicies = []Service{
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
			aws_X_AMZ_TARGET:           "AWSSimbaAPIService_v20180301.CreateBackup",
		},
		Permission: "CreateBackup",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
			aws_X_AMZ_TARGET:           "AWSSimbaAPIService_v20180301.DescribeBackups",
		},
		Permission: "DescribeBackups",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
			aws_X_AMZ_TARGET:           "AWSSimbaAPIService_v20180301.DescribeDataRepositoryAssociations",
		},
		Permission: "DescribeDataRepositoryAssociations",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
			aws_X_AMZ_TARGET:           "AWSSimbaAPIService_v20180301.DescribeDataRepositoryTasks",
		},
		Permission: "DescribeDataRepositoryTasks",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
			aws_X_AMZ_TARGET:           "AWSSimbaAPIService_v20180301.DescribeFileSystems",
		},
		Permission: "DescribeFileSystems",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
			aws_X_AMZ_TARGET:           "AWSSimbaAPIService_v20180301.DescribeSnapshots",
		},
		Permission: "DescribeSnapshots",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
			aws_X_AMZ_TARGET:           "AWSSimbaAPIService_v20180301.DescribeStorageVirtualMachines",
		},
		Permission: "DescribeStorageVirtualMachines",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
			aws_X_AMZ_TARGET:           "AWSSimbaAPIService_v20180301.DescribeVolumes",
		},
		Permission: "DescribeVolumes",
	},
}
