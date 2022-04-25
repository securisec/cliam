package policy

var GrafanaPolicies = []Service{
	{
		ServiceSuffix: "workspaces",
		Permission:    "ListWorkspaces",
	},

	// extra
	{
		ServiceSuffix:          "/workspaces/{{.workspace_id}}",
		Permission:             "DescribeWorkspace",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "workspace_id",
	},
	{
		ServiceSuffix:          "/workspaces/{{.workspace_id}}/authentication",
		Permission:             "DescribeWorkspaceAuthentication",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "workspace_id",
	},
	{
		ServiceSuffix:          "/workspaces/{{.workspace_id}}/permissions",
		Permission:             "ListPermissions",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "workspace_id",
	},
	{
		ServiceSuffix:          "/tags/{{.resource_arn}}",
		Permission:             "ListTagsForResource",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "resource_arn",
	},
}
