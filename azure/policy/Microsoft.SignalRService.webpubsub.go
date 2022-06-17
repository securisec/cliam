package policy

var Microsoft_SignalRService_webpubsub = map[string]Policy{
	"Operations_List": {
		Path:   "/providers/Microsoft.SignalRService/operations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-10-01",
		},
		OperationID: "Operations_List",
		Resource:    "Microsoft.SignalRService",
	},
	"WebPubSub_CheckNameAvailability": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.SignalRService/locations/{{.location}}/checkNameAvailability",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-10-01",
		},
		OperationID: "WebPubSub_CheckNameAvailability",
		Resource:    "Microsoft.SignalRService",
	},
	"Usages_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.SignalRService/locations/{{.location}}/usages",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-10-01",
		},
		OperationID: "Usages_List",
		Resource:    "Microsoft.SignalRService",
	},
	"WebPubSub_ListBySubscription": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.SignalRService/webPubSub",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-10-01",
		},
		OperationID: "WebPubSub_ListBySubscription",
		Resource:    "Microsoft.SignalRService",
	},
	"WebPubSub_ListByResourceGroup": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.SignalRService/webPubSub",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-10-01",
		},
		OperationID: "WebPubSub_ListByResourceGroup",
		Resource:    "Microsoft.SignalRService",
	},
	"WebPubSub_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.SignalRService/webPubSub/{{.resourceName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-10-01",
		},
		OperationID: "WebPubSub_Get",
		Resource:    "Microsoft.SignalRService",
	},
	"WebPubSubHubs_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.SignalRService/webPubSub/{{.resourceName}}/hubs",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-10-01",
		},
		OperationID: "WebPubSubHubs_List",
		Resource:    "Microsoft.SignalRService",
	},
	"WebPubSubHubs_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.SignalRService/webPubSub/{{.resourceName}}/hubs/{{.hubName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-10-01",
		},
		OperationID: "WebPubSubHubs_Get",
		Resource:    "Microsoft.SignalRService",
	},
	"WebPubSub_ListKeys": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.SignalRService/webPubSub/{{.resourceName}}/listKeys",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-10-01",
		},
		OperationID: "WebPubSub_ListKeys",
		Resource:    "Microsoft.SignalRService",
	},
	"WebPubSubPrivateEndpointConnections_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.SignalRService/webPubSub/{{.resourceName}}/privateEndpointConnections",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-10-01",
		},
		OperationID: "WebPubSubPrivateEndpointConnections_List",
		Resource:    "Microsoft.SignalRService",
	},
	"WebPubSubPrivateEndpointConnections_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.SignalRService/webPubSub/{{.resourceName}}/privateEndpointConnections/{{.privateEndpointConnectionName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-10-01",
		},
		OperationID: "WebPubSubPrivateEndpointConnections_Get",
		Resource:    "Microsoft.SignalRService",
	},
	"WebPubSubPrivateLinkResources_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.SignalRService/webPubSub/{{.resourceName}}/privateLinkResources",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-10-01",
		},
		OperationID: "WebPubSubPrivateLinkResources_List",
		Resource:    "Microsoft.SignalRService",
	},
	"WebPubSub_RegenerateKey": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.SignalRService/webPubSub/{{.resourceName}}/regenerateKey",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-10-01",
		},
		OperationID: "WebPubSub_RegenerateKey",
		Resource:    "Microsoft.SignalRService",
	},
	"WebPubSub_Restart": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.SignalRService/webPubSub/{{.resourceName}}/restart",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-10-01",
		},
		OperationID: "WebPubSub_Restart",
		Resource:    "Microsoft.SignalRService",
	},
	"WebPubSubSharedPrivateLinkResources_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.SignalRService/webPubSub/{{.resourceName}}/sharedPrivateLinkResources",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-10-01",
		},
		OperationID: "WebPubSubSharedPrivateLinkResources_List",
		Resource:    "Microsoft.SignalRService",
	},
	"WebPubSubSharedPrivateLinkResources_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.SignalRService/webPubSub/{{.resourceName}}/sharedPrivateLinkResources/{{.sharedPrivateLinkResourceName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-10-01",
		},
		OperationID: "WebPubSubSharedPrivateLinkResources_Get",
		Resource:    "Microsoft.SignalRService",
	},
	"WebPubSub_ListSkus": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.SignalRService/webPubSub/{{.resourceName}}/skus",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-10-01",
		},
		OperationID: "WebPubSub_ListSkus",
		Resource:    "Microsoft.SignalRService",
	},
}
