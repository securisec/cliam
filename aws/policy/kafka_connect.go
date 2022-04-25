package policy

var KafkaConnectPolicies = []Service{
	{
		ServiceSuffix: "connect",
		Permission:    "ListConnectors",
	},
	{
		ServiceSuffix: "custom-plugins",
		Permission:    "ListCustomPlugins",
	},
	{
		ServiceSuffix: "worker-configurations",
		Permission:    "ListWorkerConfigurations",
	},

	// extra
	{
		ServiceSuffix:          "/v1/connectors/{{.connector_arn}}",
		Permission:             "DescribeConnector",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "connector_arn",
	},
	{
		ServiceSuffix:          "/v1/custom-plugins/{{.custom_plugin_arn}}",
		Permission:             "DescribeCustomPlugin",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "custom_plugin_arn",
	},
	{
		ServiceSuffix:          "/v1/worker-configurations/{{.worker_configuration_arn}}",
		Permission:             "DescribeWorkerConfiguration",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "worker_configuration_arn",
	},
}
