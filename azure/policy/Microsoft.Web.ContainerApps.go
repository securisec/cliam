package policy

// Microsoft_Web_ContainerApps policy
var Microsoft_Web_ContainerApps = map[string]Policy{
	"ContainerApps_ListBySubscription": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Web/containerApps",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "ContainerApps_ListBySubscription",
		Resource:    "Microsoft.Web",
	},
	"ContainerApps_ListByResourceGroup": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/containerApps",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "ContainerApps_ListByResourceGroup",
		Resource:    "Microsoft.Web",
	},
	"ContainerApps_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/containerApps/{{.name}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "ContainerApps_Get",
		Resource:    "Microsoft.Web",
	},
	"ContainerApps_ListSecrets": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Web/containerApps/{{.name}}/listSecrets",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "ContainerApps_ListSecrets",
		Resource:    "Microsoft.Web",
	},
}
