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
}
