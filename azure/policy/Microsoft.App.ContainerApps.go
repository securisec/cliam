package policy

    var Microsoft_App_ContainerApps = []Policy{
        {
    Path: "/subscriptions/{{.subscriptionId}}/providers/Microsoft.App/containerApps",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2022-03-01",
    },
	OperationID:    "ContainerApps_ListBySubscription",
    Resource:       "Microsoft.App",
},{
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.App/containerApps",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2022-03-01",
    },
	OperationID:    "ContainerApps_ListByResourceGroup",
    Resource:       "Microsoft.App",
},{
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.App/containerApps/{{.containerAppName}}",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2022-03-01",
    },
	OperationID:    "ContainerApps_Get",
    Resource:       "Microsoft.App",
},{
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.App/containerApps/{{.containerAppName}}/listCustomHostNameAnalysis",
	Method: "POST",
	QueryValues:   map[string]string{
        "api-version": "2022-03-01",
    },
	OperationID:    "ContainerApps_ListCustomHostNameAnalysis",
    Resource:       "Microsoft.App",
},{
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.App/containerApps/{{.containerAppName}}/listSecrets",
	Method: "POST",
	QueryValues:   map[string]string{
        "api-version": "2022-03-01",
    },
	OperationID:    "ContainerApps_ListSecrets",
    Resource:       "Microsoft.App",
},
    }
    