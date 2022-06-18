package policy

import "github.com/securisec/cliam/shared"

var IotWirelessPolicies = map[string]Service{
	"CreateDeviceProfile": {
		Method:        "POST",
		ServiceSuffix: "device-profiles",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		},
		Permission: "CreateDeviceProfile",
	},
	"CreateServiceProfile": {
		Method:        "POST",
		ServiceSuffix: "service-profiles",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		},
		Permission: "CreateServiceProfile",
	},
	"GetLogLevelsByResourceTypes": {
		Method:        "GET",
		ServiceSuffix: "log-levels",
		Permission:    "GetLogLevelsByResourceTypes",
	},
	"GetServiceEndpoint": {
		Method:        "GET",
		ServiceSuffix: "service-endpoint",
		Permission:    "GetServiceEndpoint",
	},
	"ListDestinations": {
		Method:        "GET",
		ServiceSuffix: "destinations",
		Permission:    "ListDestinations",
	},
	"ListDeviceProfiles": {
		Method:        "GET",
		ServiceSuffix: "device-profiles",
		Permission:    "ListDeviceProfiles",
	},
	"ListFuotaTasks": {
		Method:        "GET",
		ServiceSuffix: "fuota-tasks",
		Permission:    "ListFuotaTasks",
	},
	"ListMulticastGroups": {
		Method:        "GET",
		ServiceSuffix: "multicast-groups",
		Permission:    "ListMulticastGroups",
	},
	"ListPartnerAccounts": {
		Method:        "GET",
		ServiceSuffix: "partner-accounts",
		Permission:    "ListPartnerAccounts",
	},
	"ListServiceProfiles": {
		Method:        "GET",
		ServiceSuffix: "service-profiles",
		Permission:    "ListServiceProfiles",
	},
	"ListWirelessDevices": {
		Method:        "GET",
		ServiceSuffix: "wireless-devices",
		Permission:    "ListWirelessDevices",
	},
	"ListWirelessGatewayTaskDefinitions": {
		Method:        "GET",
		ServiceSuffix: "wireless-gateway-task-definitions",
		Permission:    "ListWirelessGatewayTaskDefinitions",
	},
	"ListWirelessGateways": {
		Method:        "GET",
		ServiceSuffix: "wireless-gateways",
		Permission:    "ListWirelessGateways",
	},
	"UpdateLogLevelsByResourceTypes": {
		Method:        "POST",
		ServiceSuffix: "log-levels",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		},
		Permission: "UpdateLogLevelsByResourceTypes",
	},

	// extra
	"GetDestination": {
		ServiceSuffix:          "/destinations/{{.name}}",
		Permission:             "GetDestination",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "name",
	},
	"GetDeviceProfile": {
		ServiceSuffix:          "/device-profiles/{{.id}}",
		Permission:             "GetDeviceProfile",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "id",
	},
	"GetFuotaTask": {
		ServiceSuffix:          "/fuota-tasks/{{.id}}",
		Permission:             "GetFuotaTask",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "id",
	},
	"GetMulticastGroup": {
		ServiceSuffix:          "/multicast-groups/{{.id}}",
		Permission:             "GetMulticastGroup",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "id",
	},
	"GetMulticastGroupSession": {
		ServiceSuffix:          "/multicast-groups/{{.id}}/session",
		Permission:             "GetMulticastGroupSession",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "id",
	},
	"GetNetworkAnalyzerConfiguration": {
		ServiceSuffix:          "/network-analyzer-configurations/{{.configuration_name}}",
		Permission:             "GetNetworkAnalyzerConfiguration",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "configuration_name",
	},
	"GetServiceProfile": {
		ServiceSuffix:          "/service-profiles/{{.id}}",
		Permission:             "GetServiceProfile",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "id",
	},
	"GetWirelessDeviceStatistics": {
		ServiceSuffix:          "/wireless-devices/{{.wireless_device_id}}/statistics",
		Permission:             "GetWirelessDeviceStatistics",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "wireless_device_id",
	},
	"GetWirelessGatewayCertificate": {
		ServiceSuffix:          "/wireless-gateways/{{.id}}/certificate",
		Permission:             "GetWirelessGatewayCertificate",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "id",
	},
	"GetWirelessGatewayFirmwareInformation": {
		ServiceSuffix:          "/wireless-gateways/{{.id}}/firmware-information",
		Permission:             "GetWirelessGatewayFirmwareInformation",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "id",
	},
	"GetWirelessGatewayStatistics": {
		ServiceSuffix:          "/wireless-gateways/{{.wireless_gateway_id}}/statistics",
		Permission:             "GetWirelessGatewayStatistics",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "wireless_gateway_id",
	},
	"GetWirelessGatewayTask": {
		ServiceSuffix:          "/wireless-gateways/{{.id}}/tasks",
		Permission:             "GetWirelessGatewayTask",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "id",
	},
	"GetWirelessGatewayTaskDefinition": {
		ServiceSuffix:          "/wireless-gateway-task-definitions/{{.id}}",
		Permission:             "GetWirelessGatewayTaskDefinition",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "id",
	},
	"ListMulticastGroupsByFuotaTask": {
		ServiceSuffix:          "/fuota-tasks/{{.id}}/multicast-groups",
		Permission:             "ListMulticastGroupsByFuotaTask",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "id",
	},
	"ListQueuedMessages": {
		ServiceSuffix:          "/wireless-devices/{{.id}}/data",
		Permission:             "ListQueuedMessages",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "id",
	},
	"ListTagsForResource": {
		ServiceSuffix:          "/tags",
		Permission:             "ListTagsForResource",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "resource_arn",
	},
}
