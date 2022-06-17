package policy

var Microsoft_Subscription_subscriptions = map[string]Policy{
	"Subscription_Cancel": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Subscription/cancel",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-10-01",
		},
		OperationID: "Subscription_Cancel",
		Resource:    "Microsoft.Subscription",
	},
	"Subscription_Rename": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Subscription/rename",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-10-01",
		},
		OperationID: "Subscription_Rename",
		Resource:    "Microsoft.Subscription",
	},
	"Subscription_Enable": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Subscription/enable",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-10-01",
		},
		OperationID: "Subscription_Enable",
		Resource:    "Microsoft.Subscription",
	},
	"Operations_List": {
		Path:   "/providers/Microsoft.Subscription/operations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-10-01",
		},
		OperationID: "Operations_List",
		Resource:    "Microsoft.Subscription",
	},
	"Alias_Get": {
		Path:   "/providers/Microsoft.Subscription/aliases/{{.aliasName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-10-01",
		},
		OperationID: "Alias_Get",
		Resource:    "Microsoft.Subscription",
	},
	"Alias_List": {
		Path:   "/providers/Microsoft.Subscription/aliases",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-10-01",
		},
		OperationID: "Alias_List",
		Resource:    "Microsoft.Subscription",
	},
	"Subscription_AcceptOwnership": {
		Path:   "/providers/Microsoft.Subscription/subscriptions/{{.subscriptionId}}/acceptOwnership",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-10-01",
		},
		OperationID: "Subscription_AcceptOwnership",
		Resource:    "Microsoft.Subscription",
	},
	"Subscription_AcceptOwnershipStatus": {
		Path:   "/providers/Microsoft.Subscription/subscriptions/{{.subscriptionId}}/acceptOwnershipStatus",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-10-01",
		},
		OperationID: "Subscription_AcceptOwnershipStatus",
		Resource:    "Microsoft.Subscription",
	},
	"SubscriptionPolicy_GetPolicyForTenant": {
		Path:   "/providers/Microsoft.Subscription/policies/default",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-10-01",
		},
		OperationID: "SubscriptionPolicy_GetPolicyForTenant",
		Resource:    "Microsoft.Subscription",
	},
	"SubscriptionPolicy_ListPolicyForTenant": {
		Path:   "/providers/Microsoft.Subscription/policies",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-10-01",
		},
		OperationID: "SubscriptionPolicy_ListPolicyForTenant",
		Resource:    "Microsoft.Subscription",
	},
	"BillingAccount_GetPolicy": {
		Path:   "/providers/Microsoft.Billing/billingAccounts/{{.billingAccountId}}/providers/Microsoft.Subscription/policies/default",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-10-01",
		},
		OperationID: "BillingAccount_GetPolicy",
		Resource:    "Microsoft.Subscription",
	},
}
