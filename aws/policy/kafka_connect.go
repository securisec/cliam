package policy

// KafkaConnectPolicies policy
var KafkaConnectPolicies = map[string]Service{
	"ListConnectors": {
		Method:        "GET",
		ServiceSuffix: "v1/connectors",
		Permission:    "ListConnectors",
	},
	"ListCustomPlugins": {
		Method:        "GET",
		ServiceSuffix: "v1/custom-plugins",
		Permission:    "ListCustomPlugins",
	},
	"ListWorkerConfigurations": {
		Method:        "GET",
		ServiceSuffix: "v1/worker-configurations",
		Permission:    "ListWorkerConfigurations",
	},

	// extra
	"DescribeConnector": {
		ServiceSuffix:          "/v1/connectors/{{.connector_arn}}",
		Permission:             "DescribeConnector",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "connector_arn",
	},
	"DescribeCustomPlugin": {
		ServiceSuffix:          "/v1/custom-plugins/{{.custom_plugin_arn}}",
		Permission:             "DescribeCustomPlugin",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "custom_plugin_arn",
	},
	"DescribeWorkerConfiguration": {
		ServiceSuffix:          "/v1/worker-configurations/{{.worker_configuration_arn}}",
		Permission:             "DescribeWorkerConfiguration",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "worker_configuration_arn",
	},
}
