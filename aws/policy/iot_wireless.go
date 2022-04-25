package policy

import "github.com/securisec/cliam/shared"

var IotWirelessPolicies = []Service{
	{
		Method:        "POST",
		ServiceSuffix: "device-profiles",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		},
		Permission: "CreateDeviceProfile",
	},
	{
		Method:        "POST",
		ServiceSuffix: "service-profiles",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
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
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		},
		Permission: "UpdateLogLevelsByResourceTypes",
	},

	// extra
	{
		ServiceSuffix:          "/destinations/{{.name}}",
		Permission:             "GetDestination",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "name",
	},
	{
		ServiceSuffix:          "/device-profiles/{{.id}}",
		Permission:             "GetDeviceProfile",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "id",
	},
	{
		ServiceSuffix:          "/fuota-tasks/{{.id}}",
		Permission:             "GetFuotaTask",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "id",
	},
	{
		ServiceSuffix:          "/multicast-groups/{{.id}}",
		Permission:             "GetMulticastGroup",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "id",
	},
	{
		ServiceSuffix:          "/multicast-groups/{{.id}}/session",
		Permission:             "GetMulticastGroupSession",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "id",
	},
	{
		ServiceSuffix:          "/network-analyzer-configurations/{{.configuration_name}}",
		Permission:             "GetNetworkAnalyzerConfiguration",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "configuration_name",
	},
	{
		ServiceSuffix:          "/service-profiles/{{.id}}",
		Permission:             "GetServiceProfile",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "id",
	},
	{
		ServiceSuffix:          "/wireless-devices/{{.wireless_device_id}}/statistics",
		Permission:             "GetWirelessDeviceStatistics",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "wireless_device_id",
	},
	{
		ServiceSuffix:          "/wireless-gateways/{{.id}}/certificate",
		Permission:             "GetWirelessGatewayCertificate",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "id",
	},
	{
		ServiceSuffix:          "/wireless-gateways/{{.id}}/firmware-information",
		Permission:             "GetWirelessGatewayFirmwareInformation",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "id",
	},
	{
		ServiceSuffix:          "/wireless-gateways/{{.wireless_gateway_id}}/statistics",
		Permission:             "GetWirelessGatewayStatistics",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "wireless_gateway_id",
	},
	{
		ServiceSuffix:          "/wireless-gateways/{{.id}}/tasks",
		Permission:             "GetWirelessGatewayTask",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "id",
	},
	{
		ServiceSuffix:          "/wireless-gateway-task-definitions/{{.id}}",
		Permission:             "GetWirelessGatewayTaskDefinition",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "id",
	},
	{
		ServiceSuffix:          "/fuota-tasks/{{.id}}/multicast-groups",
		Permission:             "ListMulticastGroupsByFuotaTask",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "id",
	},
	{
		ServiceSuffix:          "/wireless-devices/{{.id}}/data",
		Permission:             "ListQueuedMessages",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "id",
	},
	{
		ServiceSuffix:          "/tags",
		Permission:             "ListTagsForResource",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "resource_arn",
	},
}
