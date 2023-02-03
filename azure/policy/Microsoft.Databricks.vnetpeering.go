package policy

// Microsoft_Databricks_vnetpeering policy
var Microsoft_Databricks_vnetpeering = map[string]Policy{
	"vNetPeering_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Databricks/workspaces/{{.workspaceName}}/virtualNetworkPeerings/{{.peeringName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2023-02-01",
		},
		OperationID: "vNetPeering_Get",
		Resource:    "Microsoft.Databricks",
	},
	"vNetPeering_ListByWorkspace": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Databricks/workspaces/{{.workspaceName}}/virtualNetworkPeerings",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2023-02-01",
		},
		OperationID: "vNetPeering_ListByWorkspace",
		Resource:    "Microsoft.Databricks",
	},
}
