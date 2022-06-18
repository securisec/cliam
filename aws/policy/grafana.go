package policy

var GrafanaPolicies = map[string]Service{
	"ListWorkspaces": {
		ServiceSuffix: "workspaces",
		Permission:    "ListWorkspaces",
	},

	// extra
	"DescribeWorkspace": {
		ServiceSuffix:          "/workspaces/{{.workspace_id}}",
		Permission:             "DescribeWorkspace",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "workspace_id",
	},
	"DescribeWorkspaceAuthentication": {
		ServiceSuffix:          "/workspaces/{{.workspace_id}}/authentication",
		Permission:             "DescribeWorkspaceAuthentication",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "workspace_id",
	},
	"ListPermissions": {
		ServiceSuffix:          "/workspaces/{{.workspace_id}}/permissions",
		Permission:             "ListPermissions",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "workspace_id",
	},
	"ListTagsForResource": {
		ServiceSuffix:          "/tags/{{.resource_arn}}",
		Permission:             "ListTagsForResource",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "resource_arn",
	},
}
