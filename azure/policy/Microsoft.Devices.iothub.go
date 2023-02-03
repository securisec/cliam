package policy

// Microsoft_Devices_iothub policy
var Microsoft_Devices_iothub = map[string]Policy{
	"Operations_List": {
		Path:   "/providers/Microsoft.Devices/operations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-07-02",
		},
		OperationID: "Operations_List",
		Resource:    "Microsoft.Devices",
	},
	"IotHubResource_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Devices/IotHubs/{{.resourceName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-07-02",
		},
		OperationID: "IotHubResource_Get",
		Resource:    "Microsoft.Devices",
	},
	"IotHubResource_ListBySubscription": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Devices/IotHubs",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-07-02",
		},
		OperationID: "IotHubResource_ListBySubscription",
		Resource:    "Microsoft.Devices",
	},
	"IotHubResource_ListByResourceGroup": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Devices/IotHubs",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-07-02",
		},
		OperationID: "IotHubResource_ListByResourceGroup",
		Resource:    "Microsoft.Devices",
	},
	"IotHubResource_GetStats": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Devices/IotHubs/{{.resourceName}}/IotHubStats",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-07-02",
		},
		OperationID: "IotHubResource_GetStats",
		Resource:    "Microsoft.Devices",
	},
	"IotHubResource_GetValidSkus": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Devices/IotHubs/{{.resourceName}}/skus",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-07-02",
		},
		OperationID: "IotHubResource_GetValidSkus",
		Resource:    "Microsoft.Devices",
	},
	"IotHubResource_ListEventHubConsumerGroups": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Devices/IotHubs/{{.resourceName}}/eventHubEndpoints/{{.eventHubEndpointName}}/ConsumerGroups",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-07-02",
		},
		OperationID: "IotHubResource_ListEventHubConsumerGroups",
		Resource:    "Microsoft.Devices",
	},
	"IotHubResource_GetEventHubConsumerGroup": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Devices/IotHubs/{{.resourceName}}/eventHubEndpoints/{{.eventHubEndpointName}}/ConsumerGroups/{{.name}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-07-02",
		},
		OperationID: "IotHubResource_GetEventHubConsumerGroup",
		Resource:    "Microsoft.Devices",
	},
	"IotHubResource_ListJobs": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Devices/IotHubs/{{.resourceName}}/jobs",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-07-02",
		},
		OperationID: "IotHubResource_ListJobs",
		Resource:    "Microsoft.Devices",
	},
	"IotHubResource_GetJob": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Devices/IotHubs/{{.resourceName}}/jobs/{{.jobId}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-07-02",
		},
		OperationID: "IotHubResource_GetJob",
		Resource:    "Microsoft.Devices",
	},
	"IotHubResource_GetQuotaMetrics": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Devices/IotHubs/{{.resourceName}}/quotaMetrics",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-07-02",
		},
		OperationID: "IotHubResource_GetQuotaMetrics",
		Resource:    "Microsoft.Devices",
	},
	"IotHubResource_GetEndpointHealth": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Devices/IotHubs/{{.iotHubName}}/routingEndpointsHealth",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-07-02",
		},
		OperationID: "IotHubResource_GetEndpointHealth",
		Resource:    "Microsoft.Devices",
	},
	"IotHubResource_CheckNameAvailability": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Devices/checkNameAvailability",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-07-02",
		},
		OperationID: "IotHubResource_CheckNameAvailability",
		Resource:    "Microsoft.Devices",
	},
	"ResourceProviderCommon_GetSubscriptionQuota": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Devices/usages",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-07-02",
		},
		OperationID: "ResourceProviderCommon_GetSubscriptionQuota",
		Resource:    "Microsoft.Devices",
	},
	"IotHubResource_TestAllRoutes": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Devices/IotHubs/{{.iotHubName}}/routing/routes/$testall",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-07-02",
		},
		OperationID: "IotHubResource_TestAllRoutes",
		Resource:    "Microsoft.Devices",
	},
	"IotHubResource_TestRoute": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Devices/IotHubs/{{.iotHubName}}/routing/routes/$testnew",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-07-02",
		},
		OperationID: "IotHubResource_TestRoute",
		Resource:    "Microsoft.Devices",
	},
	"IotHubResource_ListKeys": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Devices/IotHubs/{{.resourceName}}/listkeys",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-07-02",
		},
		OperationID: "IotHubResource_ListKeys",
		Resource:    "Microsoft.Devices",
	},
	"IotHubResource_GetKeysForKeyName": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Devices/IotHubs/{{.resourceName}}/IotHubKeys/{{.keyName}}/listkeys",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-07-02",
		},
		OperationID: "IotHubResource_GetKeysForKeyName",
		Resource:    "Microsoft.Devices",
	},
	"IotHubResource_ExportDevices": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Devices/IotHubs/{{.resourceName}}/exportDevices",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-07-02",
		},
		OperationID: "IotHubResource_ExportDevices",
		Resource:    "Microsoft.Devices",
	},
	"IotHubResource_ImportDevices": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Devices/IotHubs/{{.resourceName}}/importDevices",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-07-02",
		},
		OperationID: "IotHubResource_ImportDevices",
		Resource:    "Microsoft.Devices",
	},
	"Certificates_ListByIotHub": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Devices/IotHubs/{{.resourceName}}/certificates",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-07-02",
		},
		OperationID: "Certificates_ListByIotHub",
		Resource:    "Microsoft.Devices",
	},
	"Certificates_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Devices/IotHubs/{{.resourceName}}/certificates/{{.certificateName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-07-02",
		},
		OperationID: "Certificates_Get",
		Resource:    "Microsoft.Devices",
	},
	"Certificates_GenerateVerificationCode": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Devices/IotHubs/{{.resourceName}}/certificates/{{.certificateName}}/generateVerificationCode",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-07-02",
		},
		OperationID: "Certificates_GenerateVerificationCode",
		Resource:    "Microsoft.Devices",
	},
	"Certificates_Verify": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Devices/IotHubs/{{.resourceName}}/certificates/{{.certificateName}}/verify",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-07-02",
		},
		OperationID: "Certificates_Verify",
		Resource:    "Microsoft.Devices",
	},
	"IotHub_ManualFailover": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Devices/IotHubs/{{.iotHubName}}/failover",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-07-02",
		},
		OperationID: "IotHub_ManualFailover",
		Resource:    "Microsoft.Devices",
	},
	"PrivateLinkResources_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Devices/iotHubs/{{.resourceName}}/privateLinkResources",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-07-02",
		},
		OperationID: "PrivateLinkResources_List",
		Resource:    "Microsoft.Devices",
	},
	"PrivateLinkResources_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Devices/iotHubs/{{.resourceName}}/privateLinkResources/{{.groupId}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-07-02",
		},
		OperationID: "PrivateLinkResources_Get",
		Resource:    "Microsoft.Devices",
	},
	"PrivateEndpointConnections_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Devices/iotHubs/{{.resourceName}}/privateEndpointConnections",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-07-02",
		},
		OperationID: "PrivateEndpointConnections_List",
		Resource:    "Microsoft.Devices",
	},
	"PrivateEndpointConnections_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Devices/iotHubs/{{.resourceName}}/privateEndpointConnections/{{.privateEndpointConnectionName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-07-02",
		},
		OperationID: "PrivateEndpointConnections_Get",
		Resource:    "Microsoft.Devices",
	},
}
