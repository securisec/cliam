package policy

// Microsoft_Synapse_privatelinkhub policy
var Microsoft_Synapse_privatelinkhub = map[string]Policy{
	"PrivateLinkHubs_ListByResourceGroup": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Synapse/privateLinkHubs",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "PrivateLinkHubs_ListByResourceGroup",
		Resource:    "Microsoft.Synapse",
	},
	"PrivateLinkHubs_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Synapse/privateLinkHubs/{{.privateLinkHubName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "PrivateLinkHubs_Get",
		Resource:    "Microsoft.Synapse",
	},
	"PrivateLinkHubs_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Synapse/privateLinkHubs",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "PrivateLinkHubs_List",
		Resource:    "Microsoft.Synapse",
	},
	"PrivateEndpointConnectionsPrivateLinkHub_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Synapse/privateLinkHubs/{{.privateLinkHubName}}/privateEndpointConnections",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "PrivateEndpointConnectionsPrivateLinkHub_List",
		Resource:    "Microsoft.Synapse",
	},
	"PrivateEndpointConnectionsPrivateLinkHub_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Synapse/privateLinkHubs/{{.privateLinkHubName}}/privateEndpointConnections/{{.privateEndpointConnectionName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "PrivateEndpointConnectionsPrivateLinkHub_Get",
		Resource:    "Microsoft.Synapse",
	},
}
