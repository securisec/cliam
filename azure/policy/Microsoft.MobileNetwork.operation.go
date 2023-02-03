package policy

// Microsoft_MobileNetwork_operation policy
var Microsoft_MobileNetwork_operation = map[string]Policy{
	"Operations_List": {
		Path:   "/providers/Microsoft.MobileNetwork/operations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-11-01",
		},
		OperationID: "Operations_List",
		Resource:    "Microsoft.MobileNetwork",
	},
}
