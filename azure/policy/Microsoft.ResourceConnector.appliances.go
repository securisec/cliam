package policy

// Microsoft_ResourceConnector_appliances policy
var Microsoft_ResourceConnector_appliances = map[string]Policy{
	"Appliances_ListOperations": {
		Path:   "/providers/Microsoft.ResourceConnector/operations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-10-27",
		},
		OperationID: "Appliances_ListOperations",
		Resource:    "Microsoft.ResourceConnector",
	},
	"Appliances_ListBySubscription": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.ResourceConnector/appliances",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-10-27",
		},
		OperationID: "Appliances_ListBySubscription",
		Resource:    "Microsoft.ResourceConnector",
	},
	"Appliances_GetTelemetryConfig": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.ResourceConnector/telemetryconfig",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-10-27",
		},
		OperationID: "Appliances_GetTelemetryConfig",
		Resource:    "Microsoft.ResourceConnector",
	},
	"Appliances_ListByResourceGroup": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ResourceConnector/appliances",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-10-27",
		},
		OperationID: "Appliances_ListByResourceGroup",
		Resource:    "Microsoft.ResourceConnector",
	},
	"Appliances_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ResourceConnector/appliances/{{.resourceName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-10-27",
		},
		OperationID: "Appliances_Get",
		Resource:    "Microsoft.ResourceConnector",
	},
	"Appliances_ListClusterUserCredential": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ResourceConnector/appliances/{{.resourceName}}/listClusterUserCredential",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-10-27",
		},
		OperationID: "Appliances_ListClusterUserCredential",
		Resource:    "Microsoft.ResourceConnector",
	},
	"Appliances_ListKeys": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ResourceConnector/appliances/{{.resourceName}}/listkeys",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-10-27",
		},
		OperationID: "Appliances_ListKeys",
		Resource:    "Microsoft.ResourceConnector",
	},
	"Appliances_GetUpgradeGraph": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ResourceConnector/appliances/{{.resourceName}}/upgradeGraphs/{{.upgradeGraph}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-10-27",
		},
		OperationID: "Appliances_GetUpgradeGraph",
		Resource:    "Microsoft.ResourceConnector",
	},
}
