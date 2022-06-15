package policy

var Microsoft_Web_ContainerApps = []Policy{
	{
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Web/containerApps",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "ContainerApps_ListBySubscription",
		Resource:    "Microsoft.Web",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/containerApps",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "ContainerApps_ListByResourceGroup",
		Resource:    "Microsoft.Web",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/containerApps/{{.name}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "ContainerApps_Get",
		Resource:    "Microsoft.Web",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Web/containerApps/{{.name}}/listSecrets",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "ContainerApps_ListSecrets",
		Resource:    "Microsoft.Web",
	},
}
