package policy

var Microsoft_Web_Provider = map[string]Policy{
	"Provider_GetAvailableStacks": {
		Path:   "/providers/Microsoft.Web/availableStacks",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "Provider_GetAvailableStacks",
		Resource:    "Microsoft.Web",
	},
	"Provider_GetFunctionAppStacks": {
		Path:   "/providers/Microsoft.Web/functionAppStacks",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "Provider_GetFunctionAppStacks",
		Resource:    "Microsoft.Web",
	},
	"Provider_GetFunctionAppStacksForLocation": {
		Path:   "/providers/Microsoft.Web/locations/{{.location}}/functionAppStacks",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "Provider_GetFunctionAppStacksForLocation",
		Resource:    "Microsoft.Web",
	},
	"Provider_GetWebAppStacksForLocation": {
		Path:   "/providers/Microsoft.Web/locations/{{.location}}/webAppStacks",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "Provider_GetWebAppStacksForLocation",
		Resource:    "Microsoft.Web",
	},
	"Provider_ListOperations": {
		Path:   "/providers/Microsoft.Web/operations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "Provider_ListOperations",
		Resource:    "Microsoft.Web",
	},
	"Provider_GetWebAppStacks": {
		Path:   "/providers/Microsoft.Web/webAppStacks",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "Provider_GetWebAppStacks",
		Resource:    "Microsoft.Web",
	},
	"Provider_GetAvailableStacksOnPrem": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Web/availableStacks",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "Provider_GetAvailableStacksOnPrem",
		Resource:    "Microsoft.Web",
	},
}
