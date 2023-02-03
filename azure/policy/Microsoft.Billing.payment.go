package policy

// Microsoft_Billing_payment policy
var Microsoft_Billing_payment = map[string]Policy{
	"PaymentMethods_ListByUser": {
		Path:   "/providers/Microsoft.Billing/paymentMethods",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-10-01",
		},
		OperationID: "PaymentMethods_ListByUser",
		Resource:    "Microsoft.Billing",
	},
	"PaymentMethods_GetByUser": {
		Path:   "/providers/Microsoft.Billing/paymentMethods/{{.paymentMethodName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-10-01",
		},
		OperationID: "PaymentMethods_GetByUser",
		Resource:    "Microsoft.Billing",
	},
	"PaymentMethods_ListByBillingAccount": {
		Path:   "/providers/Microsoft.Billing/billingAccounts/{{.billingAccountName}}/paymentMethods",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-10-01",
		},
		OperationID: "PaymentMethods_ListByBillingAccount",
		Resource:    "Microsoft.Billing",
	},
	"PaymentMethods_GetByBillingAccount": {
		Path:   "/providers/Microsoft.Billing/billingAccounts/{{.billingAccountName}}/paymentMethods/{{.paymentMethodName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-10-01",
		},
		OperationID: "PaymentMethods_GetByBillingAccount",
		Resource:    "Microsoft.Billing",
	},
	"PaymentMethods_ListByBillingProfile": {
		Path:   "/providers/Microsoft.Billing/billingAccounts/{{.billingAccountName}}/billingProfiles/{{.billingProfileName}}/paymentMethodLinks",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-10-01",
		},
		OperationID: "PaymentMethods_ListByBillingProfile",
		Resource:    "Microsoft.Billing",
	},
	"PaymentMethods_GetByBillingProfile": {
		Path:   "/providers/Microsoft.Billing/billingAccounts/{{.billingAccountName}}/billingProfiles/{{.billingProfileName}}/paymentMethodLinks/{{.paymentMethodName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-10-01",
		},
		OperationID: "PaymentMethods_GetByBillingProfile",
		Resource:    "Microsoft.Billing",
	},
	"Operations_List": {
		Path:   "/providers/Microsoft.Billing/operations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-10-01",
		},
		OperationID: "Operations_List",
		Resource:    "Microsoft.Billing",
	},
}
