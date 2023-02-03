package policy

// Microsoft_App_ConnectedEnvironments policy
var Microsoft_App_ConnectedEnvironments = map[string]Policy{
	"ConnectedEnvironments_ListBySubscription": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.App/connectedEnvironments",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "ConnectedEnvironments_ListBySubscription",
		Resource:    "Microsoft.App",
	},
	"ConnectedEnvironments_ListByResourceGroup": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.App/connectedEnvironments",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "ConnectedEnvironments_ListByResourceGroup",
		Resource:    "Microsoft.App",
	},
	"ConnectedEnvironments_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourcegroups/{{.resourceGroupName}}/providers/Microsoft.App/connectedEnvironments/{{.connectedEnvironmentName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "ConnectedEnvironments_Get",
		Resource:    "Microsoft.App",
	},
	"ConnectedEnvironments_CheckNameAvailability": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.App/connectedEnvironments/{{.connectedEnvironmentName}}/checkNameAvailability",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "ConnectedEnvironments_CheckNameAvailability",
		Resource:    "Microsoft.App",
	},
}
