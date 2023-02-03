package policy

// Microsoft_Network_azureWebCategory policy
var Microsoft_Network_azureWebCategory = map[string]Policy{
	"WebCategories_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Network/azureWebCategories/{{.name}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "WebCategories_Get",
		Resource:    "Microsoft.Network",
	},
	"WebCategories_ListBySubscription": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Network/azureWebCategories",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "WebCategories_ListBySubscription",
		Resource:    "Microsoft.Network",
	},
}
