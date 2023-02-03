package policy

// Microsoft_App_SourceControls policy
var Microsoft_App_SourceControls = map[string]Policy{
	"ContainerAppsSourceControls_ListByContainerApp": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.App/containerApps/{{.containerAppName}}/sourcecontrols",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "ContainerAppsSourceControls_ListByContainerApp",
		Resource:    "Microsoft.App",
	},
	"ContainerAppsSourceControls_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.App/containerApps/{{.containerAppName}}/sourcecontrols/{{.sourceControlName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "ContainerAppsSourceControls_Get",
		Resource:    "Microsoft.App",
	},
}
