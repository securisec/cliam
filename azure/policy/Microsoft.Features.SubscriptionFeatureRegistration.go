package policy

var Microsoft_Features_SubscriptionFeatureRegistration = map[string]Policy{
	"SubscriptionFeatureRegistrations_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Features/featureProviders/{{.providerNamespace}}/subscriptionFeatureRegistrations/{{.featureName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "SubscriptionFeatureRegistrations_Get",
		Resource:    "Microsoft.Features",
	},
	"SubscriptionFeatureRegistrations_ListBySubscription": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Features/featureProviders/{{.providerNamespace}}/subscriptionFeatureRegistrations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "SubscriptionFeatureRegistrations_ListBySubscription",
		Resource:    "Microsoft.Features",
	},
	"SubscriptionFeatureRegistrations_ListAllBySubscription": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Features/subscriptionFeatureRegistrations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "SubscriptionFeatureRegistrations_ListAllBySubscription",
		Resource:    "Microsoft.Features",
	},
}
