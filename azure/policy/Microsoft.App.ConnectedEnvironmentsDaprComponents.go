package policy

// Microsoft_App_ConnectedEnvironmentsDaprComponents policy
var Microsoft_App_ConnectedEnvironmentsDaprComponents = map[string]Policy{
	"ConnectedEnvironmentsDaprComponents_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.App/connectedEnvironments/{{.connectedEnvironmentName}}/daprComponents",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "ConnectedEnvironmentsDaprComponents_List",
		Resource:    "Microsoft.App",
	},
	"ConnectedEnvironmentsDaprComponents_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.App/connectedEnvironments/{{.connectedEnvironmentName}}/daprComponents/{{.componentName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "ConnectedEnvironmentsDaprComponents_Get",
		Resource:    "Microsoft.App",
	},
	"ConnectedEnvironmentsDaprComponents_ListSecrets": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.App/connectedEnvironments/{{.connectedEnvironmentName}}/daprComponents/{{.componentName}}/listSecrets",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "ConnectedEnvironmentsDaprComponents_ListSecrets",
		Resource:    "Microsoft.App",
	},
}
