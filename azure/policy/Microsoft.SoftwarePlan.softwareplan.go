package policy

// Microsoft_SoftwarePlan_softwareplan policy
var Microsoft_SoftwarePlan_softwareplan = map[string]Policy{
	"SoftwarePlan_Register": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.SoftwarePlan/register",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2019-12-01",
		},
		OperationID: "SoftwarePlan_Register",
		Resource:    "Microsoft.SoftwarePlan",
	},
	"HybridUseBenefit_List": {
		Path:   "/{{.scope}}/providers/Microsoft.SoftwarePlan/hybridUseBenefits",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2019-12-01",
		},
		OperationID: "HybridUseBenefit_List",
		Resource:    "Microsoft.SoftwarePlan",
	},
	"HybridUseBenefit_Get": {
		Path:   "/{{.scope}}/providers/Microsoft.SoftwarePlan/hybridUseBenefits/{{.planId}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2019-12-01",
		},
		OperationID: "HybridUseBenefit_Get",
		Resource:    "Microsoft.SoftwarePlan",
	},
	"HybridUseBenefitRevision_List": {
		Path:   "/{{.scope}}/providers/Microsoft.SoftwarePlan/hybridUseBenefits/{{.planId}}/revisions",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2019-12-01",
		},
		OperationID: "HybridUseBenefitRevision_List",
		Resource:    "Microsoft.SoftwarePlan",
	},
	"Operations_List": {
		Path:   "/{{.scope}}/providers/Microsoft.SoftwarePlan/operations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2019-12-01",
		},
		OperationID: "Operations_List",
		Resource:    "Microsoft.SoftwarePlan",
	},
}
