package policy

import "github.com/securisec/cliam/shared"

// StorageGatewayPolicies policy
var StorageGatewayPolicies = map[string]Service{
	"DescribeTapeArchives": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "StorageGateway_20130630.DescribeTapeArchives",
		},
		Permission: "DescribeTapeArchives",
	},
	"ListAutomaticTapeCreationPolicies": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "StorageGateway_20130630.ListAutomaticTapeCreationPolicies",
		},
		Permission: "ListAutomaticTapeCreationPolicies",
	},
	"ListFileShares": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "StorageGateway_20130630.ListFileShares",
		},
		Permission: "ListFileShares",
	},
	"ListFileSystemAssociations": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "StorageGateway_20130630.ListFileSystemAssociations",
		},
		Permission: "ListFileSystemAssociations",
	},
	"ListGateways": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "StorageGateway_20130630.ListGateways",
		},
		Permission: "ListGateways",
	},
	"ListTapePools": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "StorageGateway_20130630.ListTapePools",
		},
		Permission: "ListTapePools",
	},
	"ListTapes": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "StorageGateway_20130630.ListTapes",
		},
		Permission: "ListTapes",
	},
	"ListVolumes": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "StorageGateway_20130630.ListVolumes",
		},
		Permission: "ListVolumes",
	},

	// extra
	"DescribeAvailabilityMonitorTest": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "StorageGateway_20130630.DescribeAvailabilityMonitorTest",
		},
		Permission:             "DescribeAvailabilityMonitorTest",
		IsExtra:                true,
		ExtraComponentBodyKey:  "GatewayARN",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "gateway_arn",
	},
	"DescribeBandwidthRateLimit": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "StorageGateway_20130630.DescribeBandwidthRateLimit",
		},
		Permission:             "DescribeBandwidthRateLimit",
		IsExtra:                true,
		ExtraComponentBodyKey:  "GatewayARN",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "gateway_arn",
	},
	"DescribeBandwidthRateLimitSchedule": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "StorageGateway_20130630.DescribeBandwidthRateLimitSchedule",
		},
		Permission:             "DescribeBandwidthRateLimitSchedule",
		IsExtra:                true,
		ExtraComponentBodyKey:  "GatewayARN",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "gateway_arn",
	},
	"DescribeCache": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "StorageGateway_20130630.DescribeCache",
		},
		Permission:             "DescribeCache",
		IsExtra:                true,
		ExtraComponentBodyKey:  "GatewayARN",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "gateway_arn",
	},
	"DescribeCachediSCSIVolumes": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "StorageGateway_20130630.DescribeCachediSCSIVolumes",
		},
		Permission:             "DescribeCachediSCSIVolumes",
		IsExtra:                true,
		ExtraComponentBodyKey:  "VolumeARNs",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "volume_arns",
	},
	"DescribeChapCredentials": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "StorageGateway_20130630.DescribeChapCredentials",
		},
		Permission:             "DescribeChapCredentials",
		IsExtra:                true,
		ExtraComponentBodyKey:  "TargetARN",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "target_arn",
	},
	"DescribeFileSystemAssociations": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "StorageGateway_20130630.DescribeFileSystemAssociations",
		},
		Permission:             "DescribeFileSystemAssociations",
		IsExtra:                true,
		ExtraComponentBodyKey:  "FileSystemAssociationARNList",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "file_system_association_arn_list",
	},
	"DescribeGatewayInformation": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "StorageGateway_20130630.DescribeGatewayInformation",
		},
		Permission:             "DescribeGatewayInformation",
		IsExtra:                true,
		ExtraComponentBodyKey:  "GatewayARN",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "gateway_arn",
	},
	"DescribeMaintenanceStartTime": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "StorageGateway_20130630.DescribeMaintenanceStartTime",
		},
		Permission:             "DescribeMaintenanceStartTime",
		IsExtra:                true,
		ExtraComponentBodyKey:  "GatewayARN",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "gateway_arn",
	},
	"DescribeNFSFileShares": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "StorageGateway_20130630.DescribeNFSFileShares",
		},
		Permission:             "DescribeNFSFileShares",
		IsExtra:                true,
		ExtraComponentBodyKey:  "FileShareARNList",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "file_share_arn_list",
	},
	"DescribeSMBFileShares": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "StorageGateway_20130630.DescribeSMBFileShares",
		},
		Permission:             "DescribeSMBFileShares",
		IsExtra:                true,
		ExtraComponentBodyKey:  "FileShareARNList",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "file_share_arn_list",
	},
	"DescribeSMBSettings": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "StorageGateway_20130630.DescribeSMBSettings",
		},
		Permission:             "DescribeSMBSettings",
		IsExtra:                true,
		ExtraComponentBodyKey:  "GatewayARN",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "gateway_arn",
	},
	"DescribeSnapshotSchedule": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "StorageGateway_20130630.DescribeSnapshotSchedule",
		},
		Permission:             "DescribeSnapshotSchedule",
		IsExtra:                true,
		ExtraComponentBodyKey:  "VolumeARN",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "volume_arn",
	},
	"DescribeStorediSCSIVolumes": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "StorageGateway_20130630.DescribeStorediSCSIVolumes",
		},
		Permission:             "DescribeStorediSCSIVolumes",
		IsExtra:                true,
		ExtraComponentBodyKey:  "VolumeARNs",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "volume_arns",
	},
	"DescribeTapeRecoveryPoints": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "StorageGateway_20130630.DescribeTapeRecoveryPoints",
		},
		Permission:             "DescribeTapeRecoveryPoints",
		IsExtra:                true,
		ExtraComponentBodyKey:  "GatewayARN",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "gateway_arn",
	},
	"DescribeTapes": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "StorageGateway_20130630.DescribeTapes",
		},
		Permission:             "DescribeTapes",
		IsExtra:                true,
		ExtraComponentBodyKey:  "GatewayARN",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "gateway_arn",
	},
	"DescribeUploadBuffer": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "StorageGateway_20130630.DescribeUploadBuffer",
		},
		Permission:             "DescribeUploadBuffer",
		IsExtra:                true,
		ExtraComponentBodyKey:  "GatewayARN",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "gateway_arn",
	},
	"DescribeVTLDevices": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "StorageGateway_20130630.DescribeVTLDevices",
		},
		Permission:             "DescribeVTLDevices",
		IsExtra:                true,
		ExtraComponentBodyKey:  "GatewayARN",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "gateway_arn",
	},
	"DescribeWorkingStorage": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "StorageGateway_20130630.DescribeWorkingStorage",
		},
		Permission:             "DescribeWorkingStorage",
		IsExtra:                true,
		ExtraComponentBodyKey:  "GatewayARN",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "gateway_arn",
	},
	"ListLocalDisks": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "StorageGateway_20130630.ListLocalDisks",
		},
		Permission:             "ListLocalDisks",
		IsExtra:                true,
		ExtraComponentBodyKey:  "GatewayARN",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "gateway_arn",
	},
	"ListTagsForResource": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "StorageGateway_20130630.ListTagsForResource",
		},
		Permission:             "ListTagsForResource",
		IsExtra:                true,
		ExtraComponentBodyKey:  "ResourceARN",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "resource_arn",
	},
	"ListVolumeInitiators": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "StorageGateway_20130630.ListVolumeInitiators",
		},
		Permission:             "ListVolumeInitiators",
		IsExtra:                true,
		ExtraComponentBodyKey:  "VolumeARN",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "volume_arn",
	},
	"ListVolumeRecoveryPoints": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "StorageGateway_20130630.ListVolumeRecoveryPoints",
		},
		Permission:             "ListVolumeRecoveryPoints",
		IsExtra:                true,
		ExtraComponentBodyKey:  "GatewayARN",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "gateway_arn",
	},
}
