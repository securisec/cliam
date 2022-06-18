package policy

var MQPolicies = map[string]Service{
	"DescribeBrokerEngineTypes": {
		Method:        "GET",
		ServiceSuffix: "v1/broker-engine-types",
		Permission:    "DescribeBrokerEngineTypes",
	},
	"DescribeBrokerInstanceOptions": {
		Method:        "GET",
		ServiceSuffix: "v1/broker-instance-options",
		Permission:    "DescribeBrokerInstanceOptions",
	},
	"ListBrokers": {
		Method:        "GET",
		ServiceSuffix: "v1/brokers",
		Permission:    "ListBrokers",
	},
	"ListConfigurations": {
		Method:        "GET",
		ServiceSuffix: "v1/configurations",
		Permission:    "ListConfigurations",
	},
}
