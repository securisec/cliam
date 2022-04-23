package policy

import "github.com/securisec/cliam/shared"

var Appstream2Policies = []Service{
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
			aws_X_AMZ_TARGET:           "PhotonAdminProxyService.DescribeAppBlocks",
		},
		Permission: "DescribeAppBlocks",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
			aws_X_AMZ_TARGET:           "PhotonAdminProxyService.DescribeApplicationFleetAssociations",
		},
		Permission: "DescribeApplicationFleetAssociations",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
			aws_X_AMZ_TARGET:           "PhotonAdminProxyService.DescribeDirectoryConfigs",
		},
		Permission: "DescribeDirectoryConfigs",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
			aws_X_AMZ_TARGET:           "PhotonAdminProxyService.DescribeFleets",
		},
		Permission: "DescribeFleets",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
			aws_X_AMZ_TARGET:           "PhotonAdminProxyService.DescribeImageBuilders",
		},
		Permission: "DescribeImageBuilders",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
			aws_X_AMZ_TARGET:           "PhotonAdminProxyService.DescribeImages",
		},
		Permission: "DescribeImages",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
			aws_X_AMZ_TARGET:           "PhotonAdminProxyService.DescribeUsageReportSubscriptions",
		},
		Permission: "DescribeUsageReportSubscriptions",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
			aws_X_AMZ_TARGET:           "PhotonAdminProxyService.DescribeUserStackAssociations",
		},
		Permission: "DescribeUserStackAssociations",
	},
}
