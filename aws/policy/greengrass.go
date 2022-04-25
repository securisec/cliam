package policy

import "github.com/securisec/cliam/shared"

var GreengrassPolicies = []Service{
	{
		Method:        "POST",
		ServiceSuffix: "greengrass/definition/connectors",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		},
		Permission: "CreateConnectorDefinition",
	},
	{
		Method:        "POST",
		ServiceSuffix: "greengrass/definition/cores",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		},
		Permission: "CreateCoreDefinition",
	},
	{
		Method:        "POST",
		ServiceSuffix: "greengrass/definition/devices",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		},
		Permission: "CreateDeviceDefinition",
	},
	{
		Method:        "POST",
		ServiceSuffix: "greengrass/definition/functions",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		},
		Permission: "CreateFunctionDefinition",
	},
	{
		Method:        "POST",
		ServiceSuffix: "greengrass/definition/loggers",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		},
		Permission: "CreateLoggerDefinition",
	},
	{
		Method:        "POST",
		ServiceSuffix: "greengrass/definition/resources",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		},
		Permission: "CreateResourceDefinition",
	},
	{
		Method:        "POST",
		ServiceSuffix: "greengrass/definition/subscriptions",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		},
		Permission: "CreateSubscriptionDefinition",
	},
	{
		Method:        "GET",
		ServiceSuffix: "greengrass/servicerole",
		Permission:    "GetServiceRoleForAccount",
	},
	{
		Method:        "GET",
		ServiceSuffix: "greengrass/bulk/deployments",
		Permission:    "ListBulkDeployments",
	},
	{
		Method:        "GET",
		ServiceSuffix: "greengrass/definition/connectors",
		Permission:    "ListConnectorDefinitions",
	},
	{
		Method:        "GET",
		ServiceSuffix: "greengrass/definition/cores",
		Permission:    "ListCoreDefinitions",
	},
	{
		Method:        "GET",
		ServiceSuffix: "greengrass/definition/devices",
		Permission:    "ListDeviceDefinitions",
	},
	{
		Method:        "GET",
		ServiceSuffix: "greengrass/definition/functions",
		Permission:    "ListFunctionDefinitions",
	},
	{
		Method:        "GET",
		ServiceSuffix: "greengrass/groups",
		Permission:    "ListGroups",
	},
	{
		Method:        "GET",
		ServiceSuffix: "greengrass/definition/loggers",
		Permission:    "ListLoggerDefinitions",
	},
	{
		Method:        "GET",
		ServiceSuffix: "greengrass/definition/resources",
		Permission:    "ListResourceDefinitions",
	},
	{
		Method:        "GET",
		ServiceSuffix: "greengrass/definition/subscriptions",
		Permission:    "ListSubscriptionDefinitions",
	},
}
