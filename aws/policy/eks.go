package policy

var EKSPolicies = map[string]Service{
	"ListClusters": {
		ServiceSuffix: "clusters",
		Permission:    "ListClusters",
	},

	// extra policies
	"DescribeCluster": {
		ServiceSuffix:          "/clusters/{{.name}}",
		Permission:             "DescribeCluster",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "name",
	},
	"ListAddons": {
		ServiceSuffix:          "/clusters/{{.cluster_name}}/addons",
		Permission:             "ListAddons",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "cluster_name",
	},
	"ListFargateProfiles": {
		ServiceSuffix:          "/clusters/{{.cluster_name}}/fargate-profiles",
		Permission:             "ListFargateProfiles",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "cluster_name",
	},
	"ListIdentityProviderConfigs": {
		ServiceSuffix:          "/clusters/{{.cluster_name}}/identity-provider-configs",
		Permission:             "ListIdentityProviderConfigs",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "cluster_name",
	},
	"ListNodegroups": {
		ServiceSuffix:          "/clusters/{{.cluster_name}}/node-groups",
		Permission:             "ListNodegroups",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "cluster_name",
	},
	"ListTagsForResource": {
		ServiceSuffix:          "/tags/{{.resource_arn}}",
		Permission:             "ListTagsForResource",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "resource_arn",
	},
	"ListUpdates": {
		ServiceSuffix:          "/clusters/{{.name}}/updates",
		Permission:             "ListUpdates",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "name",
	},
}
