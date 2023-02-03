package policy

// Microsoft_App_ConnectedEnvironmentsStorages policy
var Microsoft_App_ConnectedEnvironmentsStorages = map[string]Policy{
	"ConnectedEnvironmentsStorages_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.App/connectedEnvironments/{{.connectedEnvironmentName}}/storages",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "ConnectedEnvironmentsStorages_List",
		Resource:    "Microsoft.App",
	},
	"ConnectedEnvironmentsStorages_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.App/connectedEnvironments/{{.connectedEnvironmentName}}/storages/{{.storageName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "ConnectedEnvironmentsStorages_Get",
		Resource:    "Microsoft.App",
	},
}
