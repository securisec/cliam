package policy

import "github.com/securisec/cliam/shared"

// AppStreamPolicies policy
var AppStreamPolicies = map[string]Service{
	"DeleteUsageReportSubscription": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "PhotonAdminProxyService.DeleteUsageReportSubscription",
		},
		Permission: "DeleteUsageReportSubscription",
	},
	"DescribeAppBlocks": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "PhotonAdminProxyService.DescribeAppBlocks",
		},
		Permission: "DescribeAppBlocks",
	},
	"DescribeApplicationFleetAssociations": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "PhotonAdminProxyService.DescribeApplicationFleetAssociations",
		},
		Permission: "DescribeApplicationFleetAssociations",
	},
	"DescribeApplications": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "PhotonAdminProxyService.DescribeApplications",
		},
		Permission: "DescribeApplications",
	},
	"DescribeDirectoryConfigs": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "PhotonAdminProxyService.DescribeDirectoryConfigs",
		},
		Permission: "DescribeDirectoryConfigs",
	},
	"DescribeFleets": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "PhotonAdminProxyService.DescribeFleets",
		},
		Permission: "DescribeFleets",
	},
	"DescribeImageBuilders": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "PhotonAdminProxyService.DescribeImageBuilders",
		},
		Permission: "DescribeImageBuilders",
	},
	"DescribeImages": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "PhotonAdminProxyService.DescribeImages",
		},
		Permission: "DescribeImages",
	},
	"DescribeStacks": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "PhotonAdminProxyService.DescribeStacks",
		},
		Permission: "DescribeStacks",
	},
	"DescribeUsageReportSubscriptions": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "PhotonAdminProxyService.DescribeUsageReportSubscriptions",
		},
		Permission: "DescribeUsageReportSubscriptions",
	},
	"DescribeUserStackAssociations": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "PhotonAdminProxyService.DescribeUserStackAssociations",
		},
		Permission: "DescribeUserStackAssociations",
	},
	"UpdateFleet": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "PhotonAdminProxyService.UpdateFleet",
		},
		Permission: "UpdateFleet",
	},

	// extra
	"DescribeEntitlements": {
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
	"DescribeImagePermissions": {
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
	"DescribeUsers": {
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
	"ListAssociatedFleets": {
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
	"ListAssociatedStacks": {
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
	"ListTagsForResource": {
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
