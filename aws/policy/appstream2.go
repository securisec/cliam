package policy

import "github.com/securisec/cliam/shared"

var Appstream2Policies = []Service{
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "PhotonAdminProxyService.DescribeAppBlocks",
		},
		Permission: "DescribeAppBlocks",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "PhotonAdminProxyService.DescribeApplicationFleetAssociations",
		},
		Permission: "DescribeApplicationFleetAssociations",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "PhotonAdminProxyService.DescribeDirectoryConfigs",
		},
		Permission: "DescribeDirectoryConfigs",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "PhotonAdminProxyService.DescribeFleets",
		},
		Permission: "DescribeFleets",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "PhotonAdminProxyService.DescribeImageBuilders",
		},
		Permission: "DescribeImageBuilders",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "PhotonAdminProxyService.DescribeImages",
		},
		Permission: "DescribeImages",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "PhotonAdminProxyService.DescribeUsageReportSubscriptions",
		},
		Permission: "DescribeUsageReportSubscriptions",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "PhotonAdminProxyService.DescribeUserStackAssociations",
		},
		Permission: "DescribeUserStackAssociations",
	},

	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "PhotonAdminProxyService.DescribeEntitlements",
		},
		Permission:             "DescribeEntitlements",
		IsExtra:                true,
		ExtraComponentBodyKey:  "StackName",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "stack_name",
	},
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "PhotonAdminProxyService.DescribeImagePermissions",
		},
		Permission:             "DescribeImagePermissions",
		IsExtra:                true,
		ExtraComponentBodyKey:  "Name",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "name",
	},
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "PhotonAdminProxyService.DescribeUsers",
		},
		Permission:             "DescribeUsers",
		IsExtra:                true,
		ExtraComponentBodyKey:  "AuthenticationType",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "authentication_type",
	},
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "PhotonAdminProxyService.ListAssociatedFleets",
		},
		Permission:             "ListAssociatedFleets",
		IsExtra:                true,
		ExtraComponentBodyKey:  "StackName",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "stack_name",
	},
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "PhotonAdminProxyService.ListAssociatedStacks",
		},
		Permission:             "ListAssociatedStacks",
		IsExtra:                true,
		ExtraComponentBodyKey:  "FleetName",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "fleet_name",
	},
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "PhotonAdminProxyService.ListTagsForResource",
		},
		Permission:             "ListTagsForResource",
		IsExtra:                true,
		ExtraComponentBodyKey:  "ResourceArn",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "resource_arn",
	},
}
