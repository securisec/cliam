package policy

    var Microsoft_App_SourceControls = []Policy{
        {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.App/containerApps/{{.containerAppName}}/sourcecontrols",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2022-03-01",
    },
	OperationID:    "ContainerAppsSourceControls_ListByContainerApp",
    Resource:       "Microsoft.App",
},{
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.App/containerApps/{{.containerAppName}}/sourcecontrols/{{.sourceControlName}}",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2022-03-01",
    },
	OperationID:    "ContainerAppsSourceControls_Get",
    Resource:       "Microsoft.App",
},
    }
    