package policy

// Microsoft_SignalRService_signalr policy
var Microsoft_SignalRService_signalr = map[string]Policy{
	"Operations_List": {
		Path:   "/providers/Microsoft.SignalRService/operations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-02-01",
		},
		OperationID: "Operations_List",
		Resource:    "Microsoft.SignalRService",
	},
	"SignalR_CheckNameAvailability": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.SignalRService/locations/{{.location}}/checkNameAvailability",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-02-01",
		},
		OperationID: "SignalR_CheckNameAvailability",
		Resource:    "Microsoft.SignalRService",
	},
	"Usages_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.SignalRService/locations/{{.location}}/usages",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-02-01",
		},
		OperationID: "Usages_List",
		Resource:    "Microsoft.SignalRService",
	},
	"SignalR_ListBySubscription": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.SignalRService/signalR",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-02-01",
		},
		OperationID: "SignalR_ListBySubscription",
		Resource:    "Microsoft.SignalRService",
	},
	"SignalR_ListByResourceGroup": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.SignalRService/signalR",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-02-01",
		},
		OperationID: "SignalR_ListByResourceGroup",
		Resource:    "Microsoft.SignalRService",
	},
	"SignalR_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.SignalRService/signalR/{{.resourceName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-02-01",
		},
		OperationID: "SignalR_Get",
		Resource:    "Microsoft.SignalRService",
	},
	"SignalRCustomCertificates_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.SignalRService/signalR/{{.resourceName}}/customCertificates",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-02-01",
		},
		OperationID: "SignalRCustomCertificates_List",
		Resource:    "Microsoft.SignalRService",
	},
	"SignalRCustomCertificates_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.SignalRService/signalR/{{.resourceName}}/customCertificates/{{.certificateName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-02-01",
		},
		OperationID: "SignalRCustomCertificates_Get",
		Resource:    "Microsoft.SignalRService",
	},
	"SignalRCustomDomains_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.SignalRService/signalR/{{.resourceName}}/customDomains",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-02-01",
		},
		OperationID: "SignalRCustomDomains_List",
		Resource:    "Microsoft.SignalRService",
	},
	"SignalRCustomDomains_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.SignalRService/signalR/{{.resourceName}}/customDomains/{{.name}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-02-01",
		},
		OperationID: "SignalRCustomDomains_Get",
		Resource:    "Microsoft.SignalRService",
	},
	"SignalR_ListKeys": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.SignalRService/signalR/{{.resourceName}}/listKeys",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-02-01",
		},
		OperationID: "SignalR_ListKeys",
		Resource:    "Microsoft.SignalRService",
	},
	"SignalRPrivateEndpointConnections_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.SignalRService/signalR/{{.resourceName}}/privateEndpointConnections",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-02-01",
		},
		OperationID: "SignalRPrivateEndpointConnections_List",
		Resource:    "Microsoft.SignalRService",
	},
	"SignalRPrivateEndpointConnections_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.SignalRService/signalR/{{.resourceName}}/privateEndpointConnections/{{.privateEndpointConnectionName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-02-01",
		},
		OperationID: "SignalRPrivateEndpointConnections_Get",
		Resource:    "Microsoft.SignalRService",
	},
	"SignalRPrivateLinkResources_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.SignalRService/signalR/{{.resourceName}}/privateLinkResources",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-02-01",
		},
		OperationID: "SignalRPrivateLinkResources_List",
		Resource:    "Microsoft.SignalRService",
	},
	"SignalR_RegenerateKey": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.SignalRService/signalR/{{.resourceName}}/regenerateKey",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-02-01",
		},
		OperationID: "SignalR_RegenerateKey",
		Resource:    "Microsoft.SignalRService",
	},
	"SignalR_Restart": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.SignalRService/signalR/{{.resourceName}}/restart",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-02-01",
		},
		OperationID: "SignalR_Restart",
		Resource:    "Microsoft.SignalRService",
	},
	"SignalRSharedPrivateLinkResources_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.SignalRService/signalR/{{.resourceName}}/sharedPrivateLinkResources",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-02-01",
		},
		OperationID: "SignalRSharedPrivateLinkResources_List",
		Resource:    "Microsoft.SignalRService",
	},
	"SignalRSharedPrivateLinkResources_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.SignalRService/signalR/{{.resourceName}}/sharedPrivateLinkResources/{{.sharedPrivateLinkResourceName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-02-01",
		},
		OperationID: "SignalRSharedPrivateLinkResources_Get",
		Resource:    "Microsoft.SignalRService",
	},
	"SignalR_ListSkus": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.SignalRService/signalR/{{.resourceName}}/skus",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-02-01",
		},
		OperationID: "SignalR_ListSkus",
		Resource:    "Microsoft.SignalRService",
	},
}
