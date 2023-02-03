package policy

// Microsoft_CostManagement_costmanagement_generatecostdetailsreport policy
var Microsoft_CostManagement_costmanagement_generatecostdetailsreport = map[string]Policy{
	"GenerateCostDetailsReport_CreateOperation": {
		Path:   "/{{.scope}}/providers/Microsoft.CostManagement/generateCostDetailsReport",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "GenerateCostDetailsReport_CreateOperation",
		Resource:    "Microsoft.CostManagement",
	},
	"GenerateCostDetailsReport_GetOperationResults": {
		Path:   "/{{.scope}}/providers/Microsoft.CostManagement/costDetailsOperationResults/{{.operationId}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "GenerateCostDetailsReport_GetOperationResults",
		Resource:    "Microsoft.CostManagement",
	},
}
