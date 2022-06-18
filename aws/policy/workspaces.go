package policy

import "github.com/securisec/cliam/shared"

var WorkspacesPolicies = map[string]Service{
	"DescribeAccount": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "WorkspacesService.DescribeAccount",
		},
		Permission: "DescribeAccount",
	},
	"DescribeAccountModifications": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "WorkspacesService.DescribeAccountModifications",
		},
		Permission: "DescribeAccountModifications",
	},
	"DescribeConnectionAliases": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "WorkspacesService.DescribeConnectionAliases",
		},
		Permission: "DescribeConnectionAliases",
	},
	"DescribeIpGroups": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "WorkspacesService.DescribeIpGroups",
		},
		Permission: "DescribeIpGroups",
	},
	"DescribeWorkspaceBundles": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "WorkspacesService.DescribeWorkspaceBundles",
		},
		Permission: "DescribeWorkspaceBundles",
	},
	"DescribeWorkspaceDirectories": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "WorkspacesService.DescribeWorkspaceDirectories",
		},
		Permission: "DescribeWorkspaceDirectories",
	},
	"DescribeWorkspaceImages": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "WorkspacesService.DescribeWorkspaceImages",
		},
		Permission: "DescribeWorkspaceImages",
	},
	"DescribeWorkspaces": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "WorkspacesService.DescribeWorkspaces",
		},
		Permission: "DescribeWorkspaces",
	},
	"DescribeWorkspacesConnectionStatus": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "WorkspacesService.DescribeWorkspacesConnectionStatus",
		},
		Permission: "DescribeWorkspacesConnectionStatus",
	},

	// extra
	"DescribeClientBranding": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "WorkspacesService.DescribeClientBranding",
		},
		Permission:             "DescribeClientBranding",
		IsExtra:                true,
		ExtraComponentBodyKey:  "ResourceId",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "resource_id",
	},
	"DescribeClientProperties": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "WorkspacesService.DescribeClientProperties",
		},
		Permission:             "DescribeClientProperties",
		IsExtra:                true,
		ExtraComponentBodyKey:  "ResourceIds",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "resource_ids",
	},
	"DescribeConnectClientAddIns": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "WorkspacesService.DescribeConnectClientAddIns",
		},
		Permission:             "DescribeConnectClientAddIns",
		IsExtra:                true,
		ExtraComponentBodyKey:  "ResourceId",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "resource_id",
	},
	"DescribeConnectionAliasPermissions": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "WorkspacesService.DescribeConnectionAliasPermissions",
		},
		Permission:             "DescribeConnectionAliasPermissions",
		IsExtra:                true,
		ExtraComponentBodyKey:  "AliasId",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "alias_id",
	},
	"DescribeTags": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "WorkspacesService.DescribeTags",
		},
		Permission:             "DescribeTags",
		IsExtra:                true,
		ExtraComponentBodyKey:  "ResourceId",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "resource_id",
	},
	"DescribeWorkspaceImagePermissions": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "WorkspacesService.DescribeWorkspaceImagePermissions",
		},
		Permission:             "DescribeWorkspaceImagePermissions",
		IsExtra:                true,
		ExtraComponentBodyKey:  "ImageId",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "image_id",
	},
	"DescribeWorkspaceSnapshots": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "WorkspacesService.DescribeWorkspaceSnapshots",
		},
		Permission:             "DescribeWorkspaceSnapshots",
		IsExtra:                true,
		ExtraComponentBodyKey:  "WorkspaceId",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "workspace_id",
	},
	"ListAvailableManagementCidrRanges": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "WorkspacesService.ListAvailableManagementCidrRanges",
		},
		Permission:             "ListAvailableManagementCidrRanges",
		IsExtra:                true,
		ExtraComponentBodyKey:  "ManagementCidrRangeConstraint",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "management_cidr_range_constraint",
	},
}
