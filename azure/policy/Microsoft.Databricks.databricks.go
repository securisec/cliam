package policy

// Microsoft_Databricks_databricks policy
var Microsoft_Databricks_databricks = map[string]Policy{
	"Workspaces_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Databricks/workspaces/{{.workspaceName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2023-02-01",
		},
		OperationID: "Workspaces_Get",
		Resource:    "Microsoft.Databricks",
	},
	"Workspaces_ListByResourceGroup": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Databricks/workspaces",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2023-02-01",
		},
		OperationID: "Workspaces_ListByResourceGroup",
		Resource:    "Microsoft.Databricks",
	},
	"Workspaces_ListBySubscription": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Databricks/workspaces",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2023-02-01",
		},
		OperationID: "Workspaces_ListBySubscription",
		Resource:    "Microsoft.Databricks",
	},
	"Operations_List": {
		Path:   "/providers/Microsoft.Databricks/operations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2023-02-01",
		},
		OperationID: "Operations_List",
		Resource:    "Microsoft.Databricks",
	},
	"PrivateLinkResources_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Databricks/workspaces/{{.workspaceName}}/privateLinkResources",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2023-02-01",
		},
		OperationID: "PrivateLinkResources_List",
		Resource:    "Microsoft.Databricks",
	},
	"PrivateLinkResources_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Databricks/workspaces/{{.workspaceName}}/privateLinkResources/{{.groupId}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2023-02-01",
		},
		OperationID: "PrivateLinkResources_Get",
		Resource:    "Microsoft.Databricks",
	},
	"PrivateEndpointConnections_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Databricks/workspaces/{{.workspaceName}}/privateEndpointConnections",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2023-02-01",
		},
		OperationID: "PrivateEndpointConnections_List",
		Resource:    "Microsoft.Databricks",
	},
	"PrivateEndpointConnections_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Databricks/workspaces/{{.workspaceName}}/privateEndpointConnections/{{.privateEndpointConnectionName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2023-02-01",
		},
		OperationID: "PrivateEndpointConnections_Get",
		Resource:    "Microsoft.Databricks",
	},
	"OutboundNetworkDependenciesEndpoints_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Databricks/workspaces/{{.workspaceName}}/outboundNetworkDependenciesEndpoints",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2023-02-01",
		},
		OperationID: "OutboundNetworkDependenciesEndpoints_List",
		Resource:    "Microsoft.Databricks",
	},
}
