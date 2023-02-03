package policy

// Microsoft_CostManagement_common_types policy
var Microsoft_CostManagement_common_types = map[string]Policy{
	"Operations_List": {
		Path:   "/providers/Microsoft.CostManagement/operations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "Operations_List",
		Resource:    "Microsoft.CostManagement",
	},
}
