package policy

import "github.com/securisec/cliam/shared"

// FSXPolicies policy
var FSXPolicies = map[string]Service{
	"DescribeBackups": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSSimbaAPIService_v20180301.DescribeBackups",
		},
		Permission: "DescribeBackups",
	},
	"DescribeDataRepositoryAssociations": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSSimbaAPIService_v20180301.DescribeDataRepositoryAssociations",
		},
		Permission: "DescribeDataRepositoryAssociations",
	},
	"DescribeDataRepositoryTasks": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSSimbaAPIService_v20180301.DescribeDataRepositoryTasks",
		},
		Permission: "DescribeDataRepositoryTasks",
	},
	"DescribeFileCaches": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSSimbaAPIService_v20180301.DescribeFileCaches",
		},
		Permission: "DescribeFileCaches",
	},
	"DescribeFileSystems": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSSimbaAPIService_v20180301.DescribeFileSystems",
		},
		Permission: "DescribeFileSystems",
	},
	"DescribeSnapshots": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSSimbaAPIService_v20180301.DescribeSnapshots",
		},
		Permission: "DescribeSnapshots",
	},
	"DescribeStorageVirtualMachines": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSSimbaAPIService_v20180301.DescribeStorageVirtualMachines",
		},
		Permission: "DescribeStorageVirtualMachines",
	},
	"DescribeVolumes": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSSimbaAPIService_v20180301.DescribeVolumes",
		},
		Permission: "DescribeVolumes",
	},

	// extra
	"DescribeFileSystemAliases": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSSimbaAPIService_v20180301.DescribeFileSystemAliases",
		},
		Permission:             "DescribeFileSystemAliases",
		IsExtra:                true,
		ExtraComponentBodyKey:  "FileSystemId",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "file_system_id",
	},
	"ListTagsForResource": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSSimbaAPIService_v20180301.ListTagsForResource",
		},
		Permission:             "ListTagsForResource",
		IsExtra:                true,
		ExtraComponentBodyKey:  "ResourceARN",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "resource_arn",
	},
}
