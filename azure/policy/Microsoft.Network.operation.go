package policy

// Microsoft_Network_operation policy
var Microsoft_Network_operation = map[string]Policy{
	"Operations_List": {
		Path:   "/providers/Microsoft.Network/operations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "Operations_List",
		Resource:    "Microsoft.Network",
	},
}
