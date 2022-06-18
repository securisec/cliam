package policy

import "github.com/securisec/cliam/shared"

var DSPolicies = map[string]Service{
	"DescribeDirectories": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "DirectoryService_20150416.DescribeDirectories",
		},
		Permission: "DescribeDirectories",
	},
	"DescribeEventTopics": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "DirectoryService_20150416.DescribeEventTopics",
		},
		Permission: "DescribeEventTopics",
	},
	"DescribeSnapshots": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "DirectoryService_20150416.DescribeSnapshots",
		},
		Permission: "DescribeSnapshots",
	},
	"DescribeTrusts": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "DirectoryService_20150416.DescribeTrusts",
		},
		Permission: "DescribeTrusts",
	},
	"GetDirectoryLimits": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "DirectoryService_20150416.GetDirectoryLimits",
		},
		Permission: "GetDirectoryLimits",
	},
	"ListLogSubscriptions": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "DirectoryService_20150416.ListLogSubscriptions",
		},
		Permission: "ListLogSubscriptions",
	},

	// extra
	"DescribeClientAuthenticationSettings": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "DirectoryService_20150416.DescribeClientAuthenticationSettings",
		},
		Permission:             "DescribeClientAuthenticationSettings",
		IsExtra:                true,
		ExtraComponentBodyKey:  "DirectoryId",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "directory_id",
	},
	"DescribeConditionalForwarders": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "DirectoryService_20150416.DescribeConditionalForwarders",
		},
		Permission:             "DescribeConditionalForwarders",
		IsExtra:                true,
		ExtraComponentBodyKey:  "DirectoryId",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "directory_id",
	},
	"DescribeDomainControllers": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "DirectoryService_20150416.DescribeDomainControllers",
		},
		Permission:             "DescribeDomainControllers",
		IsExtra:                true,
		ExtraComponentBodyKey:  "DirectoryId",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "directory_id",
	},
	"DescribeLDAPSSettings": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "DirectoryService_20150416.DescribeLDAPSSettings",
		},
		Permission:             "DescribeLDAPSSettings",
		IsExtra:                true,
		ExtraComponentBodyKey:  "DirectoryId",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "directory_id",
	},
	"DescribeRegions": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "DirectoryService_20150416.DescribeRegions",
		},
		Permission:             "DescribeRegions",
		IsExtra:                true,
		ExtraComponentBodyKey:  "DirectoryId",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "directory_id",
	},
	"DescribeSharedDirectories": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "DirectoryService_20150416.DescribeSharedDirectories",
		},
		Permission:             "DescribeSharedDirectories",
		IsExtra:                true,
		ExtraComponentBodyKey:  "OwnerDirectoryId",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "owner_directory_id",
	},
	"GetSnapshotLimits": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "DirectoryService_20150416.GetSnapshotLimits",
		},
		Permission:             "GetSnapshotLimits",
		IsExtra:                true,
		ExtraComponentBodyKey:  "DirectoryId",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "directory_id",
	},
	"ListCertificates": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "DirectoryService_20150416.ListCertificates",
		},
		Permission:             "ListCertificates",
		IsExtra:                true,
		ExtraComponentBodyKey:  "DirectoryId",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "directory_id",
	},
	"ListIpRoutes": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "DirectoryService_20150416.ListIpRoutes",
		},
		Permission:             "ListIpRoutes",
		IsExtra:                true,
		ExtraComponentBodyKey:  "DirectoryId",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "directory_id",
	},
	"ListSchemaExtensions": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "DirectoryService_20150416.ListSchemaExtensions",
		},
		Permission:             "ListSchemaExtensions",
		IsExtra:                true,
		ExtraComponentBodyKey:  "DirectoryId",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "directory_id",
	},
	"ListTagsForResource": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "DirectoryService_20150416.ListTagsForResource",
		},
		Permission:             "ListTagsForResource",
		IsExtra:                true,
		ExtraComponentBodyKey:  "ResourceId",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "resource_id",
	},
}
