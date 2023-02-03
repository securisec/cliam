package policy

// Microsoft_ServiceBus_namespace_preview policy
var Microsoft_ServiceBus_namespace_preview = map[string]Policy{
	"Namespaces_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.ServiceBus/namespaces",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-11-01",
		},
		OperationID: "Namespaces_List",
		Resource:    "Microsoft.ServiceBus",
	},
	"Namespaces_ListByResourceGroup": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ServiceBus/namespaces",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-11-01",
		},
		OperationID: "Namespaces_ListByResourceGroup",
		Resource:    "Microsoft.ServiceBus",
	},
	"Namespaces_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ServiceBus/namespaces/{{.namespaceName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-11-01",
		},
		OperationID: "Namespaces_Get",
		Resource:    "Microsoft.ServiceBus",
	},
	"PrivateEndpointConnections_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ServiceBus/namespaces/{{.namespaceName}}/privateEndpointConnections",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-11-01",
		},
		OperationID: "PrivateEndpointConnections_List",
		Resource:    "Microsoft.ServiceBus",
	},
	"PrivateEndpointConnections_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ServiceBus/namespaces/{{.namespaceName}}/privateEndpointConnections/{{.privateEndpointConnectionName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-11-01",
		},
		OperationID: "PrivateEndpointConnections_Get",
		Resource:    "Microsoft.ServiceBus",
	},
	"PrivateLinkResources_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ServiceBus/namespaces/{{.namespaceName}}/privateLinkResources",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-11-01",
		},
		OperationID: "PrivateLinkResources_Get",
		Resource:    "Microsoft.ServiceBus",
	},
}
