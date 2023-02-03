package policy

// Microsoft_Relay_Namespaces policy
var Microsoft_Relay_Namespaces = map[string]Policy{
	"Namespaces_CheckNameAvailability": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Relay/checkNameAvailability",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-11-01",
		},
		OperationID: "Namespaces_CheckNameAvailability",
		Resource:    "Microsoft.Relay",
	},
	"Namespaces_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Relay/namespaces",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-11-01",
		},
		OperationID: "Namespaces_List",
		Resource:    "Microsoft.Relay",
	},
	"Namespaces_ListByResourceGroup": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Relay/namespaces",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-11-01",
		},
		OperationID: "Namespaces_ListByResourceGroup",
		Resource:    "Microsoft.Relay",
	},
	"Namespaces_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Relay/namespaces/{{.namespaceName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-11-01",
		},
		OperationID: "Namespaces_Get",
		Resource:    "Microsoft.Relay",
	},
	"PrivateEndpointConnections_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Relay/namespaces/{{.namespaceName}}/privateEndpointConnections",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-11-01",
		},
		OperationID: "PrivateEndpointConnections_List",
		Resource:    "Microsoft.Relay",
	},
	"PrivateEndpointConnections_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Relay/namespaces/{{.namespaceName}}/privateEndpointConnections/{{.privateEndpointConnectionName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-11-01",
		},
		OperationID: "PrivateEndpointConnections_Get",
		Resource:    "Microsoft.Relay",
	},
	"PrivateLinkResources_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Relay/namespaces/{{.namespaceName}}/privateLinkResources/{{.privateLinkResourceName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-11-01",
		},
		OperationID: "PrivateLinkResources_Get",
		Resource:    "Microsoft.Relay",
	},
	"PrivateLinkResources_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Relay/namespaces/{{.namespaceName}}/privateLinkResources",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-11-01",
		},
		OperationID: "PrivateLinkResources_List",
		Resource:    "Microsoft.Relay",
	},
}
