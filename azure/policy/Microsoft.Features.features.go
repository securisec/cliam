package policy

var Microsoft_Features_features = []Policy{
	{
		Path:   "/providers/Microsoft.Features/operations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "ListOperations",
		Resource:    "Microsoft.Features",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Features/features",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "Features_ListAll",
		Resource:    "Microsoft.Features",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Features/providers/{{.resourceProviderNamespace}}/features",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "Features_List",
		Resource:    "Microsoft.Features",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Features/providers/{{.resourceProviderNamespace}}/features/{{.featureName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "Features_Get",
		Resource:    "Microsoft.Features",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Features/providers/{{.resourceProviderNamespace}}/features/{{.featureName}}/register",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "Features_Register",
		Resource:    "Microsoft.Features",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Features/providers/{{.resourceProviderNamespace}}/features/{{.featureName}}/unregister",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "Features_Unregister",
		Resource:    "Microsoft.Features",
	},
}
