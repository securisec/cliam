package policy

var KafkaPolicies = map[string]Service{
	"ListClusters": {
		ServiceSuffix: "v1/clusters",
		Permission:    "ListClusters",
	},
	"ListClustersV2": {
		ServiceSuffix: "api/v2/clusters",
		Permission:    "ListClustersV2",
	},
	"GetCompatibleKafkaVersions": {
		ServiceSuffix: "v1/compitable-kafka-versions",
		Permission:    "GetCompatibleKafkaVersions",
	},
	"ListConfigurations": {
		ServiceSuffix: "v1/configurations",
		Permission:    "ListConfigurations",
	},
	"ListKafkaVersions": {
		ServiceSuffix: "v1/kafka-versions",
		Permission:    "ListKafkaVersions",
	},

	// extra
	"DescribeCluster": {
		ServiceSuffix:          "/v1/clusters/{{.cluster_arn}}",
		Permission:             "DescribeCluster",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "cluster_arn",
	},
	"DescribeClusterV2": {
		ServiceSuffix:          "/api/v2/clusters/{{.cluster_arn}}",
		Permission:             "DescribeClusterV2",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "cluster_arn",
	},
	"DescribeClusterOperation": {
		ServiceSuffix:          "/v1/operations/{{.cluster_operation_arn}}",
		Permission:             "DescribeClusterOperation",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "cluster_operation_arn",
	},
	"DescribeConfiguration": {
		ServiceSuffix:          "/v1/configurations/{{.arn}}",
		Permission:             "DescribeConfiguration",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "arn",
	},
	"GetBootstrapBrokers": {
		ServiceSuffix:          "/v1/clusters/{{.cluster_arn}}/bootstrap-brokers",
		Permission:             "GetBootstrapBrokers",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "cluster_arn",
	},
	"ListClusterOperations": {
		ServiceSuffix:          "/v1/clusters/{{.cluster_arn}}/operations",
		Permission:             "ListClusterOperations",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "cluster_arn",
	},
	"ListConfigurationRevisions": {
		ServiceSuffix:          "/v1/configurations/{{.arn}}/revisions",
		Permission:             "ListConfigurationRevisions",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "arn",
	},
	"ListNodes": {
		ServiceSuffix:          "/v1/clusters/{{.cluster_arn}}/nodes",
		Permission:             "ListNodes",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "cluster_arn",
	},
	"ListScramSecrets": {
		ServiceSuffix:          "/v1/clusters/{{.cluster_arn}}/scram-secrets",
		Permission:             "ListScramSecrets",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "cluster_arn",
	},
	"ListTagsForResource": {
		ServiceSuffix:          "/v1/tags/{{.resource_arn}}",
		Permission:             "ListTagsForResource",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "resource_arn",
	},
}
