package policy

// Microsoft_DelegatedNetwork_operations policy
var Microsoft_DelegatedNetwork_operations = map[string]Policy{
	"Operations_List": {
		Path:   "/providers/Microsoft.DelegatedNetwork/operations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-15",
		},
		OperationID: "Operations_List",
		Resource:    "Microsoft.DelegatedNetwork",
	},
}
