package policy

// Microsoft_CostManagement_costmanagement_exports policy
var Microsoft_CostManagement_costmanagement_exports = map[string]Policy{
	"Exports_List": {
		Path:   "/{{.scope}}/providers/Microsoft.CostManagement/exports",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "Exports_List",
		Resource:    "Microsoft.CostManagement",
	},
	"Exports_Get": {
		Path:   "/{{.scope}}/providers/Microsoft.CostManagement/exports/{{.exportName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "Exports_Get",
		Resource:    "Microsoft.CostManagement",
	},
	"Exports_Execute": {
		Path:   "/{{.scope}}/providers/Microsoft.CostManagement/exports/{{.exportName}}/run",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "Exports_Execute",
		Resource:    "Microsoft.CostManagement",
	},
	"Exports_GetExecutionHistory": {
		Path:   "/{{.scope}}/providers/Microsoft.CostManagement/exports/{{.exportName}}/runHistory",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "Exports_GetExecutionHistory",
		Resource:    "Microsoft.CostManagement",
	},
}
