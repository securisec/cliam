package policy

// Microsoft_CostManagement_costmanagement_benefits policy
var Microsoft_CostManagement_costmanagement_benefits = map[string]Policy{
	"BenefitRecommendations_List": {
		Path:   "/{{.billingScope}}/providers/Microsoft.CostManagement/benefitRecommendations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "BenefitRecommendations_List",
		Resource:    "Microsoft.CostManagement",
	},
	"BenefitUtilizationSummaries_ListByBillingAccountId": {
		Path:   "/providers/Microsoft.Billing/billingAccounts/{{.billingAccountId}}/providers/Microsoft.CostManagement/benefitUtilizationSummaries",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "BenefitUtilizationSummaries_ListByBillingAccountId",
		Resource:    "Microsoft.CostManagement",
	},
	"BenefitUtilizationSummaries_ListByBillingProfileId": {
		Path:   "/providers/Microsoft.Billing/billingAccounts/{{.billingAccountId}}/billingProfiles/{{.billingProfileId}}/providers/Microsoft.CostManagement/benefitUtilizationSummaries",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "BenefitUtilizationSummaries_ListByBillingProfileId",
		Resource:    "Microsoft.CostManagement",
	},
	"BenefitUtilizationSummaries_ListBySavingsPlanOrder": {
		Path:   "/providers/Microsoft.BillingBenefits/savingsPlanOrders/{{.savingsPlanOrderId}}/providers/Microsoft.CostManagement/benefitUtilizationSummaries",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "BenefitUtilizationSummaries_ListBySavingsPlanOrder",
		Resource:    "Microsoft.CostManagement",
	},
	"BenefitUtilizationSummaries_ListBySavingsPlanId": {
		Path:   "/providers/Microsoft.BillingBenefits/savingsPlanOrders/{{.savingsPlanOrderId}}/savingsPlans/{{.savingsPlanId}}/providers/Microsoft.CostManagement/benefitUtilizationSummaries",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "BenefitUtilizationSummaries_ListBySavingsPlanId",
		Resource:    "Microsoft.CostManagement",
	},
}
