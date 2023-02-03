package policy

// Microsoft_App_ManagedEnvironmentsStorages policy
var Microsoft_App_ManagedEnvironmentsStorages = map[string]Policy{
	"ManagedEnvironmentsStorages_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.App/managedEnvironments/{{.environmentName}}/storages",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "ManagedEnvironmentsStorages_List",
		Resource:    "Microsoft.App",
	},
	"ManagedEnvironmentsStorages_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.App/managedEnvironments/{{.environmentName}}/storages/{{.storageName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "ManagedEnvironmentsStorages_Get",
		Resource:    "Microsoft.App",
	},
}
