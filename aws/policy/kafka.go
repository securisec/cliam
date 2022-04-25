package policy

var KafkaPolicies = []Service{
	{
		ServiceSuffix: "v1/clusters",
		Permission:    "ListClusters",
	},
	{
		ServiceSuffix: "api/v2/clusters",
		Permission:    "ListClustersV2",
	},
	{
		ServiceSuffix: "v1/compitable-kafka-versions",
		Permission:    "GetCompatibleKafkaVersions",
	},
	{
		ServiceSuffix: "v1/configurations",
		Permission:    "ListConfigurations",
	},
	{
		ServiceSuffix: "v1/kafka-versions",
		Permission:    "ListKafkaVersions",
	},

	// extra
	{
		ServiceSuffix:          "/v1/clusters/{{.cluster_arn}}",
		Permission:             "DescribeCluster",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "cluster_arn",
	},
	{
		ServiceSuffix:          "/api/v2/clusters/{{.cluster_arn}}",
		Permission:             "DescribeClusterV2",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "cluster_arn",
	},
	{
		ServiceSuffix:          "/v1/operations/{{.cluster_operation_arn}}",
		Permission:             "DescribeClusterOperation",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "cluster_operation_arn",
	},
	{
		ServiceSuffix:          "/v1/configurations/{{.arn}}",
		Permission:             "DescribeConfiguration",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "arn",
	},
	{
		ServiceSuffix:          "/v1/clusters/{{.cluster_arn}}/bootstrap-brokers",
		Permission:             "GetBootstrapBrokers",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "cluster_arn",
	},
	{
		ServiceSuffix:          "/v1/clusters/{{.cluster_arn}}/operations",
		Permission:             "ListClusterOperations",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "cluster_arn",
	},
	{
		ServiceSuffix:          "/v1/configurations/{{.arn}}/revisions",
		Permission:             "ListConfigurationRevisions",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "arn",
	},
	{
		ServiceSuffix:          "/v1/clusters/{{.cluster_arn}}/nodes",
		Permission:             "ListNodes",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "cluster_arn",
	},
	{
		ServiceSuffix:          "/v1/clusters/{{.cluster_arn}}/scram-secrets",
		Permission:             "ListScramSecrets",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "cluster_arn",
	},
	{
		ServiceSuffix:          "/v1/tags/{{.resource_arn}}",
		Permission:             "ListTagsForResource",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "resource_arn",
	},
}
