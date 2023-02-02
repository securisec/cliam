package policy

// MQPolicies policy
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

	// extra
	"DescribeBroker": {
		ServiceSuffix:          "/v1/brokers/{{.broker_id}}",
		Permission:             "DescribeBroker",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "broker_id",
	},
	"DescribeConfiguration": {
		ServiceSuffix:          "/v1/configurations/{{.configuration_id}}",
		Permission:             "DescribeConfiguration",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "configuration_id",
	},
	"ListConfigurationRevisions": {
		ServiceSuffix:          "/v1/configurations/{{.configuration_id}}/revisions",
		Permission:             "ListConfigurationRevisions",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "configuration_id",
	},
	"ListTags": {
		ServiceSuffix:          "/v1/tags/{{.resource_arn}}",
		Permission:             "ListTags",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "resource_arn",
	},
	"ListUsers": {
		ServiceSuffix:          "/v1/brokers/{{.broker_id}}/users",
		Permission:             "ListUsers",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "broker_id",
	},
}
