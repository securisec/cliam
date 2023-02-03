package policy

// Microsoft_Network_privateLinkService policy
var Microsoft_Network_privateLinkService = map[string]Policy{
	"PrivateLinkServices_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/privateLinkServices/{{.serviceName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "PrivateLinkServices_Get",
		Resource:    "Microsoft.Network",
	},
	"PrivateLinkServices_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/privateLinkServices",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "PrivateLinkServices_List",
		Resource:    "Microsoft.Network",
	},
	"PrivateLinkServices_ListBySubscription": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Network/privateLinkServices",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "PrivateLinkServices_ListBySubscription",
		Resource:    "Microsoft.Network",
	},
	"PrivateLinkServices_GetPrivateEndpointConnection": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/privateLinkServices/{{.serviceName}}/privateEndpointConnections/{{.peConnectionName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "PrivateLinkServices_GetPrivateEndpointConnection",
		Resource:    "Microsoft.Network",
	},
	"PrivateLinkServices_ListPrivateEndpointConnections": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/privateLinkServices/{{.serviceName}}/privateEndpointConnections",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "PrivateLinkServices_ListPrivateEndpointConnections",
		Resource:    "Microsoft.Network",
	},
	"PrivateLinkServices_CheckPrivateLinkServiceVisibility": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Network/locations/{{.location}}/checkPrivateLinkServiceVisibility",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "PrivateLinkServices_CheckPrivateLinkServiceVisibility",
		Resource:    "Microsoft.Network",
	},
	"PrivateLinkServices_CheckPrivateLinkServiceVisibilityByResourceGroup": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/locations/{{.location}}/checkPrivateLinkServiceVisibility",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "PrivateLinkServices_CheckPrivateLinkServiceVisibilityByResourceGroup",
		Resource:    "Microsoft.Network",
	},
	"PrivateLinkServices_ListAutoApprovedPrivateLinkServices": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Network/locations/{{.location}}/autoApprovedPrivateLinkServices",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "PrivateLinkServices_ListAutoApprovedPrivateLinkServices",
		Resource:    "Microsoft.Network",
	},
	"PrivateLinkServices_ListAutoApprovedPrivateLinkServicesByResourceGroup": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/locations/{{.location}}/autoApprovedPrivateLinkServices",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "PrivateLinkServices_ListAutoApprovedPrivateLinkServicesByResourceGroup",
		Resource:    "Microsoft.Network",
	},
}
