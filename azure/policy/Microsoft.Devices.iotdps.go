package policy

// Microsoft_Devices_iotdps policy
var Microsoft_Devices_iotdps = map[string]Policy{
	"Operations_List": {
		Path:   "/providers/Microsoft.Devices/operations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-12-12",
		},
		OperationID: "Operations_List",
		Resource:    "Microsoft.Devices",
	},
	"DpsCertificate_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Devices/provisioningServices/{{.provisioningServiceName}}/certificates/{{.certificateName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-12-12",
		},
		OperationID: "DpsCertificate_Get",
		Resource:    "Microsoft.Devices",
	},
	"IotDpsResource_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Devices/provisioningServices/{{.provisioningServiceName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-12-12",
		},
		OperationID: "IotDpsResource_Get",
		Resource:    "Microsoft.Devices",
	},
	"DpsCertificate_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Devices/provisioningServices/{{.provisioningServiceName}}/certificates",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-12-12",
		},
		OperationID: "DpsCertificate_List",
		Resource:    "Microsoft.Devices",
	},
	"IotDpsResource_ListBySubscription": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Devices/provisioningServices",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-12-12",
		},
		OperationID: "IotDpsResource_ListBySubscription",
		Resource:    "Microsoft.Devices",
	},
	"IotDpsResource_ListByResourceGroup": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Devices/provisioningServices",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-12-12",
		},
		OperationID: "IotDpsResource_ListByResourceGroup",
		Resource:    "Microsoft.Devices",
	},
	"IotDpsResource_GetOperationResult": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Devices/provisioningServices/{{.provisioningServiceName}}/operationresults/{{.operationId}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-12-12",
		},
		OperationID: "IotDpsResource_GetOperationResult",
		Resource:    "Microsoft.Devices",
	},
	"IotDpsResource_listValidSkus": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Devices/provisioningServices/{{.provisioningServiceName}}/skus",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-12-12",
		},
		OperationID: "IotDpsResource_listValidSkus",
		Resource:    "Microsoft.Devices",
	},
	"DpsCertificate_GenerateVerificationCode": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Devices/provisioningServices/{{.provisioningServiceName}}/certificates/{{.certificateName}}/generateVerificationCode",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-12-12",
		},
		OperationID: "DpsCertificate_GenerateVerificationCode",
		Resource:    "Microsoft.Devices",
	},
	"DpsCertificate_VerifyCertificate": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Devices/provisioningServices/{{.provisioningServiceName}}/certificates/{{.certificateName}}/verify",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-12-12",
		},
		OperationID: "DpsCertificate_VerifyCertificate",
		Resource:    "Microsoft.Devices",
	},
	"IotDpsResource_CheckProvisioningServiceNameAvailability": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Devices/checkProvisioningServiceNameAvailability",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-12-12",
		},
		OperationID: "IotDpsResource_CheckProvisioningServiceNameAvailability",
		Resource:    "Microsoft.Devices",
	},
	"IotDpsResource_ListKeys": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Devices/provisioningServices/{{.provisioningServiceName}}/listkeys",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-12-12",
		},
		OperationID: "IotDpsResource_ListKeys",
		Resource:    "Microsoft.Devices",
	},
	"IotDpsResource_ListKeysForKeyName": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Devices/provisioningServices/{{.provisioningServiceName}}/keys/{{.keyName}}/listkeys",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-12-12",
		},
		OperationID: "IotDpsResource_ListKeysForKeyName",
		Resource:    "Microsoft.Devices",
	},
	"IotDpsResource_ListPrivateLinkResources": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Devices/provisioningServices/{{.resourceName}}/privateLinkResources",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-12-12",
		},
		OperationID: "IotDpsResource_ListPrivateLinkResources",
		Resource:    "Microsoft.Devices",
	},
	"IotDpsResource_GetPrivateLinkResources": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Devices/provisioningServices/{{.resourceName}}/privateLinkResources/{{.groupId}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-12-12",
		},
		OperationID: "IotDpsResource_GetPrivateLinkResources",
		Resource:    "Microsoft.Devices",
	},
	"IotDpsResource_ListPrivateEndpointConnections": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Devices/provisioningServices/{{.resourceName}}/privateEndpointConnections",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-12-12",
		},
		OperationID: "IotDpsResource_ListPrivateEndpointConnections",
		Resource:    "Microsoft.Devices",
	},
	"IotDpsResource_GetPrivateEndpointConnection": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Devices/provisioningServices/{{.resourceName}}/privateEndpointConnections/{{.privateEndpointConnectionName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-12-12",
		},
		OperationID: "IotDpsResource_GetPrivateEndpointConnection",
		Resource:    "Microsoft.Devices",
	},
}
