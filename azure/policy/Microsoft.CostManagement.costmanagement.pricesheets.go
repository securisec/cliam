package policy

// Microsoft_CostManagement_costmanagement_pricesheets policy
var Microsoft_CostManagement_costmanagement_pricesheets = map[string]Policy{
	"PriceSheet_Download": {
		Path:   "/providers/Microsoft.Billing/billingAccounts/{{.billingAccountName}}/billingProfiles/{{.billingProfileName}}/invoices/{{.invoiceName}}/providers/Microsoft.CostManagement/pricesheets/default/download",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "PriceSheet_Download",
		Resource:    "Microsoft.CostManagement",
	},
	"PriceSheet_DownloadByBillingProfile": {
		Path:   "/providers/Microsoft.Billing/billingAccounts/{{.billingAccountName}}/billingProfiles/{{.billingProfileName}}/providers/Microsoft.CostManagement/pricesheets/default/download",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "PriceSheet_DownloadByBillingProfile",
		Resource:    "Microsoft.CostManagement",
	},
}
