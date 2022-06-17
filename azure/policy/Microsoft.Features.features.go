package policy

var Microsoft_Features_features = map[string]Policy{
	"ListOperations": {
		Path:   "/providers/Microsoft.Features/operations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "ListOperations",
		Resource:    "Microsoft.Features",
	},
	"Features_ListAll": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Features/features",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "Features_ListAll",
		Resource:    "Microsoft.Features",
	},
	"Features_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Features/providers/{{.resourceProviderNamespace}}/features",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "Features_List",
		Resource:    "Microsoft.Features",
	},
	"Features_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Features/providers/{{.resourceProviderNamespace}}/features/{{.featureName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "Features_Get",
		Resource:    "Microsoft.Features",
	},
	"Features_Register": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Features/providers/{{.resourceProviderNamespace}}/features/{{.featureName}}/register",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "Features_Register",
		Resource:    "Microsoft.Features",
	},
	"Features_Unregister": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Features/providers/{{.resourceProviderNamespace}}/features/{{.featureName}}/unregister",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "Features_Unregister",
		Resource:    "Microsoft.Features",
	},
}
