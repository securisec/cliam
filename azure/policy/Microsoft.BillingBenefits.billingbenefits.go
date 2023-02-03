package policy

// Microsoft_BillingBenefits_billingbenefits policy
var Microsoft_BillingBenefits_billingbenefits = map[string]Policy{
	"Operations_List": {
		Path:   "/providers/Microsoft.BillingBenefits/operations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-11-01",
		},
		OperationID: "Operations_List",
		Resource:    "Microsoft.BillingBenefits",
	},
	"SavingsPlanOrderAlias_Get": {
		Path:   "/providers/Microsoft.BillingBenefits/savingsPlanOrderAliases/{{.savingsPlanOrderAliasName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-11-01",
		},
		OperationID: "SavingsPlanOrderAlias_Get",
		Resource:    "Microsoft.BillingBenefits",
	},
	"SavingsPlanOrder_Get": {
		Path:   "/providers/Microsoft.BillingBenefits/savingsPlanOrders/{{.savingsPlanOrderId}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-11-01",
		},
		OperationID: "SavingsPlanOrder_Get",
		Resource:    "Microsoft.BillingBenefits",
	},
	"SavingsPlanOrder_Elevate": {
		Path:   "/providers/Microsoft.BillingBenefits/savingsPlanOrders/{{.savingsPlanOrderId}}/elevate",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-11-01",
		},
		OperationID: "SavingsPlanOrder_Elevate",
		Resource:    "Microsoft.BillingBenefits",
	},
	"SavingsPlanOrder_List": {
		Path:   "/providers/Microsoft.BillingBenefits/savingsPlanOrders",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-11-01",
		},
		OperationID: "SavingsPlanOrder_List",
		Resource:    "Microsoft.BillingBenefits",
	},
	"SavingsPlan_List": {
		Path:   "/providers/Microsoft.BillingBenefits/savingsPlanOrders/{{.savingsPlanOrderId}}/savingsPlans",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-11-01",
		},
		OperationID: "SavingsPlan_List",
		Resource:    "Microsoft.BillingBenefits",
	},
	"SavingsPlan_ListAll": {
		Path:   "/providers/Microsoft.BillingBenefits/savingsPlans",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-11-01",
		},
		OperationID: "SavingsPlan_ListAll",
		Resource:    "Microsoft.BillingBenefits",
	},
	"SavingsPlan_Get": {
		Path:   "/providers/Microsoft.BillingBenefits/savingsPlanOrders/{{.savingsPlanOrderId}}/savingsPlans/{{.savingsPlanId}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-11-01",
		},
		OperationID: "SavingsPlan_Get",
		Resource:    "Microsoft.BillingBenefits",
	},
	"SavingsPlan_ValidateUpdate": {
		Path:   "/providers/Microsoft.BillingBenefits/savingsPlanOrders/{{.savingsPlanOrderId}}/savingsPlans/{{.savingsPlanId}}/validate",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-11-01",
		},
		OperationID: "SavingsPlan_ValidateUpdate",
		Resource:    "Microsoft.BillingBenefits",
	},
	"ValidatePurchase": {
		Path:   "/providers/Microsoft.BillingBenefits/validate",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-11-01",
		},
		OperationID: "ValidatePurchase",
		Resource:    "Microsoft.BillingBenefits",
	},
	"ReservationOrderAlias_Get": {
		Path:   "/providers/Microsoft.BillingBenefits/reservationOrderAliases/{{.reservationOrderAliasName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-11-01",
		},
		OperationID: "ReservationOrderAlias_Get",
		Resource:    "Microsoft.BillingBenefits",
	},
}
