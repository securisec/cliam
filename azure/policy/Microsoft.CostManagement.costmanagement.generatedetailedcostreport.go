package policy

// Microsoft_CostManagement_costmanagement_generatedetailedcostreport policy
var Microsoft_CostManagement_costmanagement_generatedetailedcostreport = map[string]Policy{
	"GenerateDetailedCostReport_CreateOperation": {
		Path:   "/{{.scope}}/providers/Microsoft.CostManagement/generateDetailedCostReport",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "GenerateDetailedCostReport_CreateOperation",
		Resource:    "Microsoft.CostManagement",
	},
	"GenerateDetailedCostReportOperationResults_Get": {
		Path:   "/{{.scope}}/providers/Microsoft.CostManagement/operationResults/{{.operationId}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "GenerateDetailedCostReportOperationResults_Get",
		Resource:    "Microsoft.CostManagement",
	},
	"GenerateDetailedCostReportOperationStatus_Get": {
		Path:   "/{{.scope}}/providers/Microsoft.CostManagement/operationStatus/{{.operationId}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "GenerateDetailedCostReportOperationStatus_Get",
		Resource:    "Microsoft.CostManagement",
	},
}
