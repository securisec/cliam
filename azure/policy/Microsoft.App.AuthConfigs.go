package policy

// Microsoft_App_AuthConfigs policy
var Microsoft_App_AuthConfigs = map[string]Policy{
	"ContainerAppsAuthConfigs_ListByContainerApp": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.App/containerApps/{{.containerAppName}}/authConfigs",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "ContainerAppsAuthConfigs_ListByContainerApp",
		Resource:    "Microsoft.App",
	},
	"ContainerAppsAuthConfigs_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.App/containerApps/{{.containerAppName}}/authConfigs/{{.authConfigName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "ContainerAppsAuthConfigs_Get",
		Resource:    "Microsoft.App",
	},
}
