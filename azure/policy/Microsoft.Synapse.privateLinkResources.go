package policy

// Microsoft_Synapse_privateLinkResources policy
var Microsoft_Synapse_privateLinkResources = map[string]Policy{
	"PrivateLinkResources_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Synapse/workspaces/{{.workspaceName}}/privateLinkResources",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "PrivateLinkResources_List",
		Resource:    "Microsoft.Synapse",
	},
	"PrivateLinkResources_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Synapse/workspaces/{{.workspaceName}}/privateLinkResources/{{.privateLinkResourceName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "PrivateLinkResources_Get",
		Resource:    "Microsoft.Synapse",
	},
	"PrivateLinkHubPrivateLinkResources_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Synapse/privateLinkHubs/{{.privateLinkHubName}}/privateLinkResources",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "PrivateLinkHubPrivateLinkResources_List",
		Resource:    "Microsoft.Synapse",
	},
	"PrivateLinkHubPrivateLinkResources_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Synapse/privateLinkHubs/{{.privateLinkHubName}}/privateLinkResources/{{.privateLinkResourceName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "PrivateLinkHubPrivateLinkResources_Get",
		Resource:    "Microsoft.Synapse",
	},
}
