package policy

var Microsoft_Web_KubeEnvironments = []Policy{
	{
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Web/kubeEnvironments",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "KubeEnvironments_ListBySubscription",
		Resource:    "Microsoft.Web",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/kubeEnvironments",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "KubeEnvironments_ListByResourceGroup",
		Resource:    "Microsoft.Web",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/kubeEnvironments/{{.name}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "KubeEnvironments_Get",
		Resource:    "Microsoft.Web",
	},
}
