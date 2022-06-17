package policy

var Microsoft_App_ContainerApps = map[string]Policy{
	"ContainerApps_ListBySubscription": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.App/containerApps",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "ContainerApps_ListBySubscription",
		Resource:    "Microsoft.App",
	},
	"ContainerApps_ListByResourceGroup": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.App/containerApps",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "ContainerApps_ListByResourceGroup",
		Resource:    "Microsoft.App",
	},
	"ContainerApps_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.App/containerApps/{{.containerAppName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "ContainerApps_Get",
		Resource:    "Microsoft.App",
	},
	"ContainerApps_ListCustomHostNameAnalysis": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.App/containerApps/{{.containerAppName}}/listCustomHostNameAnalysis",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "ContainerApps_ListCustomHostNameAnalysis",
		Resource:    "Microsoft.App",
	},
	"ContainerApps_ListSecrets": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.App/containerApps/{{.containerAppName}}/listSecrets",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "ContainerApps_ListSecrets",
		Resource:    "Microsoft.App",
	},
}
