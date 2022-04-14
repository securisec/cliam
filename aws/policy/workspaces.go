package policy

import "github.com/securisec/cliam/shared"

var WorkspacesPolicies = []Service{
	{
		Method:   "POST",
		JsonData: `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
			aws_X_AMZ_TARGET:           "WorkspacesService.DescribeAccount",
		},
		Permission: "DescribeAccount",
	},
	{
		Method:   "POST",
		JsonData: `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
			aws_X_AMZ_TARGET:           "WorkspacesService.DescribeAccountModifications",
		},
		Permission: "DescribeAccountModifications",
	},
	{
		Method:   "POST",
		JsonData: `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
			aws_X_AMZ_TARGET:           "WorkspacesService.DescribeConnectionAliases",
		},
		Permission: "DescribeConnectionAliases",
	},
	{
		Method:   "POST",
		JsonData: `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
			aws_X_AMZ_TARGET:           "WorkspacesService.DescribeIpGroups",
		},
		Permission: "DescribeIpGroups",
	},
	{
		Method:   "POST",
		JsonData: `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
			aws_X_AMZ_TARGET:           "WorkspacesService.DescribeWorkspaceBundles",
		},
		Permission: "DescribeWorkspaceBundles",
	},
	{
		Method:   "POST",
		JsonData: `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
			aws_X_AMZ_TARGET:           "WorkspacesService.DescribeWorkspaceDirectories",
		},
		Permission: "DescribeWorkspaceDirectories",
	},
	{
		Method:   "POST",
		JsonData: `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
			aws_X_AMZ_TARGET:           "WorkspacesService.DescribeWorkspaceImages",
		},
		Permission: "DescribeWorkspaceImages",
	},
	{
		Method:   "POST",
		JsonData: `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
			aws_X_AMZ_TARGET:           "WorkspacesService.DescribeWorkspaces",
		},
		Permission: "DescribeWorkspaces",
	},
	{
		Method:   "POST",
		JsonData: `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
			aws_X_AMZ_TARGET:           "WorkspacesService.DescribeWorkspacesConnectionStatus",
		},
		Permission: "DescribeWorkspacesConnectionStatus",
	},
}