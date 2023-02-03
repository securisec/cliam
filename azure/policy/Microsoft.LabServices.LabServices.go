package policy

// Microsoft_LabServices_LabServices policy
var Microsoft_LabServices_LabServices = map[string]Policy{
	"Operations_List": {
		Path:   "/providers/Microsoft.LabServices/operations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-08-01",
		},
		OperationID: "Operations_List",
		Resource:    "Microsoft.LabServices",
	},
}
