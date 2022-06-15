package policy

var Microsoft_Web_Provider = []Policy{
	{
		Path:   "/providers/Microsoft.Web/availableStacks",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "Provider_GetAvailableStacks",
		Resource:    "Microsoft.Web",
	}, {
		Path:   "/providers/Microsoft.Web/functionAppStacks",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "Provider_GetFunctionAppStacks",
		Resource:    "Microsoft.Web",
	}, {
		Path:   "/providers/Microsoft.Web/locations/{{.location}}/functionAppStacks",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "Provider_GetFunctionAppStacksForLocation",
		Resource:    "Microsoft.Web",
	}, {
		Path:   "/providers/Microsoft.Web/locations/{{.location}}/webAppStacks",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "Provider_GetWebAppStacksForLocation",
		Resource:    "Microsoft.Web",
	}, {
		Path:   "/providers/Microsoft.Web/operations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "Provider_ListOperations",
		Resource:    "Microsoft.Web",
	}, {
		Path:   "/providers/Microsoft.Web/webAppStacks",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "Provider_GetWebAppStacks",
		Resource:    "Microsoft.Web",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Web/availableStacks",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "Provider_GetAvailableStacksOnPrem",
		Resource:    "Microsoft.Web",
	},
}
