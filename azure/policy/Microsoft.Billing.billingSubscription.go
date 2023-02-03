package policy

// Microsoft_Billing_billingSubscription policy
var Microsoft_Billing_billingSubscription = map[string]Policy{
	"BillingSubscriptions_ListByBillingAccount": {
		Path:   "/providers/Microsoft.Billing/billingAccounts/{{.billingAccountName}}/billingSubscriptions",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-10-01",
		},
		OperationID: "BillingSubscriptions_ListByBillingAccount",
		Resource:    "Microsoft.Billing",
	},
	"BillingSubscriptions_Get": {
		Path:   "/providers/Microsoft.Billing/billingAccounts/{{.billingAccountName}}/billingSubscriptions/{{.billingSubscriptionName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-10-01",
		},
		OperationID: "BillingSubscriptions_Get",
		Resource:    "Microsoft.Billing",
	},
	"BillingSubscriptions_Move": {
		Path:   "/providers/Microsoft.Billing/billingAccounts/{{.billingAccountName}}/billingSubscriptions/{{.billingSubscriptionName}}/move",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-10-01",
		},
		OperationID: "BillingSubscriptions_Move",
		Resource:    "Microsoft.Billing",
	},
	"BillingSubscriptions_ValidateMoveEligibility": {
		Path:   "/providers/Microsoft.Billing/billingAccounts/{{.billingAccountName}}/billingSubscriptions/{{.billingSubscriptionName}}/validateMoveEligibility",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-10-01",
		},
		OperationID: "BillingSubscriptions_ValidateMoveEligibility",
		Resource:    "Microsoft.Billing",
	},
	"BillingSubscriptions_Merge": {
		Path:   "/providers/Microsoft.Billing/billingAccounts/{{.billingAccountName}}/billingSubscriptions/{{.billingSubscriptionName}}/merge",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-10-01",
		},
		OperationID: "BillingSubscriptions_Merge",
		Resource:    "Microsoft.Billing",
	},
	"BillingSubscriptions_Split": {
		Path:   "/providers/Microsoft.Billing/billingAccounts/{{.billingAccountName}}/billingSubscriptions/{{.billingSubscriptionName}}/split",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-10-01",
		},
		OperationID: "BillingSubscriptions_Split",
		Resource:    "Microsoft.Billing",
	},
	"BillingSubscriptionsAliases_ListByBillingAccount": {
		Path:   "/providers/Microsoft.Billing/billingAccounts/{{.billingAccountName}}/billingSubscriptionAliases",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-10-01",
		},
		OperationID: "BillingSubscriptionsAliases_ListByBillingAccount",
		Resource:    "Microsoft.Billing",
	},
	"BillingSubscriptionsAliases_Get": {
		Path:   "/providers/Microsoft.Billing/billingAccounts/{{.billingAccountName}}/billingSubscriptionAliases/{{.aliasName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-10-01",
		},
		OperationID: "BillingSubscriptionsAliases_Get",
		Resource:    "Microsoft.Billing",
	},
}
