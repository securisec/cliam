package policy

// Microsoft_App_Global policy
var Microsoft_App_Global = map[string]Policy{
	"Operations_List": {
		Path:   "/providers/Microsoft.App/operations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "Operations_List",
		Resource:    "Microsoft.App",
	},
}
