package policy

// Microsoft_App_BillingMeters policy
var Microsoft_App_BillingMeters = map[string]Policy{
	"BillingMeters_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.App/locations/{{.location}}/billingMeters",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "BillingMeters_Get",
		Resource:    "Microsoft.App",
	},
}
