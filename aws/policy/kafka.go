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
}
