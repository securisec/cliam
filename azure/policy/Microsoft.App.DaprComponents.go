package policy

    var Microsoft_App_DaprComponents = []Policy{
        {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.App/managedEnvironments/{{.environmentName}}/daprComponents",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2022-03-01",
    },
	OperationID:    "DaprComponents_List",
    Resource:       "Microsoft.App",
},{
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.App/managedEnvironments/{{.environmentName}}/daprComponents/{{.componentName}}",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2022-03-01",
    },
	OperationID:    "DaprComponents_Get",
    Resource:       "Microsoft.App",
},{
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.App/managedEnvironments/{{.environmentName}}/daprComponents/{{.componentName}}/listSecrets",
	Method: "POST",
	QueryValues:   map[string]string{
        "api-version": "2022-03-01",
    },
	OperationID:    "DaprComponents_ListSecrets",
    Resource:       "Microsoft.App",
},
    }
    