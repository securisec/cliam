package policy

import "github.com/securisec/cliam/shared"

var IotWirelessPolicies = []Service{
	{
		Method:        "POST",
		ServiceSuffix: "device-profiles",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
		},
		Permission: "CreateDeviceProfile",
	},
	{
		Method:        "POST",
		ServiceSuffix: "service-profiles",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
		},
		Permission: "CreateServiceProfile",
	},
	{
		Method:        "GET",
		ServiceSuffix: "log-levels",
		Permission:    "GetLogLevelsByResourceTypes",
	},
	{
		Method:        "GET",
		ServiceSuffix: "service-endpoint",
		Permission:    "GetServiceEndpoint",
	},
	{
		Method:        "GET",
		ServiceSuffix: "destinations",
		Permission:    "ListDestinations",
	},
	{
		Method:        "GET",
		ServiceSuffix: "device-profiles",
		Permission:    "ListDeviceProfiles",
	},
	{
		Method:        "GET",
		ServiceSuffix: "fuota-tasks",
		Permission:    "ListFuotaTasks",
	},
	{
		Method:        "GET",
		ServiceSuffix: "multicast-groups",
		Permission:    "ListMulticastGroups",
	},
	{
		Method:        "GET",
		ServiceSuffix: "partner-accounts",
		Permission:    "ListPartnerAccounts",
	},
	{
		Method:        "GET",
		ServiceSuffix: "service-profiles",
		Permission:    "ListServiceProfiles",
	},
	{
		Method:        "GET",
		ServiceSuffix: "wireless-devices",
		Permission:    "ListWirelessDevices",
	},
	{
		Method:        "GET",
		ServiceSuffix: "wireless-gateway-task-definitions",
		Permission:    "ListWirelessGatewayTaskDefinitions",
	},
	{
		Method:        "GET",
		ServiceSuffix: "wireless-gateways",
		Permission:    "ListWirelessGateways",
	},
	{
		Method:        "POST",
		ServiceSuffix: "log-levels",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
		},
		Permission: "UpdateLogLevelsByResourceTypes",
	},
}
