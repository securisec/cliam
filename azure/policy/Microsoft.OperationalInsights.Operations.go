package policy

// Microsoft_OperationalInsights_Operations policy
var Microsoft_OperationalInsights_Operations = map[string]Policy{
	"Operations_List": {
		Path:   "/providers/Microsoft.OperationalInsights/operations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "Operations_List",
		Resource:    "Microsoft.OperationalInsights",
	},
}
