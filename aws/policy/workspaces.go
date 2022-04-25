package policy

import "github.com/securisec/cliam/shared"

var WorkspacesPolicies = []Service{
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "WorkspacesService.DescribeAccount",
		},
		Permission: "DescribeAccount",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "WorkspacesService.DescribeAccountModifications",
		},
		Permission: "DescribeAccountModifications",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "WorkspacesService.DescribeConnectionAliases",
		},
		Permission: "DescribeConnectionAliases",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "WorkspacesService.DescribeIpGroups",
		},
		Permission: "DescribeIpGroups",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "WorkspacesService.DescribeWorkspaceBundles",
		},
		Permission: "DescribeWorkspaceBundles",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "WorkspacesService.DescribeWorkspaceDirectories",
		},
		Permission: "DescribeWorkspaceDirectories",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "WorkspacesService.DescribeWorkspaceImages",
		},
		Permission: "DescribeWorkspaceImages",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "WorkspacesService.DescribeWorkspaces",
		},
		Permission: "DescribeWorkspaces",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "WorkspacesService.DescribeWorkspacesConnectionStatus",
		},
		Permission: "DescribeWorkspacesConnectionStatus",
	},
}
