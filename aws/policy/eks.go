package policy

var EKSPolicies = []Service{
	{
		ServiceSuffix: "clusters",
		Permission:    "ListClusters",
	},

	// extra policies
	{
		ServiceSuffix:          "/clusters/{{.name}}",
		Permission:             "DescribeCluster",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "name",
	},
	{
		ServiceSuffix:          "/clusters/{{.cluster_name}}/addons",
		Permission:             "ListAddons",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "cluster_name",
	},
	{
		ServiceSuffix:          "/clusters/{{.cluster_name}}/fargate-profiles",
		Permission:             "ListFargateProfiles",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "cluster_name",
	},
	{
		ServiceSuffix:          "/clusters/{{.cluster_name}}/identity-provider-configs",
		Permission:             "ListIdentityProviderConfigs",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "cluster_name",
	},
	{
		ServiceSuffix:          "/clusters/{{.cluster_name}}/node-groups",
		Permission:             "ListNodegroups",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "cluster_name",
	},
	{
		ServiceSuffix:          "/tags/{{.resource_arn}}",
		Permission:             "ListTagsForResource",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "resource_arn",
	},
	{
		ServiceSuffix:          "/clusters/{{.name}}/updates",
		Permission:             "ListUpdates",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "name",
	},
}
