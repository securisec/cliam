package policy

var MQPolicies = []Service{
	{
		Method:        "GET",
		ServiceSuffix: "v1/broker-engine-types",
		Permission:    "DescribeBrokerEngineTypes",
	},
	{
		Method:        "GET",
		ServiceSuffix: "v1/broker-instance-options",
		Permission:    "DescribeBrokerInstanceOptions",
	},
	{
		Method:        "GET",
		ServiceSuffix: "v1/brokers",
		Permission:    "ListBrokers",
	},
	{
		Method:        "GET",
		ServiceSuffix: "v1/configurations",
		Permission:    "ListConfigurations",
	},
}
