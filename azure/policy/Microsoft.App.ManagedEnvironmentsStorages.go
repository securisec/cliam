package policy

    var Microsoft_App_ManagedEnvironmentsStorages = []Policy{
        {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.App/managedEnvironments/{{.environmentName}}/storages",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2022-03-01",
    },
	OperationID:    "ManagedEnvironmentsStorages_List",
    Resource:       "Microsoft.App",
},{
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.App/managedEnvironments/{{.environmentName}}/storages/{{.storageName}}",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2022-03-01",
    },
	OperationID:    "ManagedEnvironmentsStorages_Get",
    Resource:       "Microsoft.App",
},
    }
    