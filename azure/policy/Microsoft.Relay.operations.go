package policy

// Microsoft_Relay_operations policy
var Microsoft_Relay_operations = map[string]Policy{
	"Operations_List": {
		Path:   "/providers/Microsoft.Relay/operations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-11-01",
		},
		OperationID: "Operations_List",
		Resource:    "Microsoft.Relay",
	},
}
