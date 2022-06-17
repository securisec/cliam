package policy

var Microsoft_Web_KubeEnvironments = map[string]Policy{
	"KubeEnvironments_ListBySubscription": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Web/kubeEnvironments",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "KubeEnvironments_ListBySubscription",
		Resource:    "Microsoft.Web",
	},
	"KubeEnvironments_ListByResourceGroup": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/kubeEnvironments",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "KubeEnvironments_ListByResourceGroup",
		Resource:    "Microsoft.Web",
	},
	"KubeEnvironments_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/kubeEnvironments/{{.name}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "KubeEnvironments_Get",
		Resource:    "Microsoft.Web",
	},
}
