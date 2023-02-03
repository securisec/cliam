package policy

// Microsoft_AzureStackHCI_operations policy
var Microsoft_AzureStackHCI_operations = map[string]Policy{
	"Operations_List": {
		Path:   "/providers/Microsoft.AzureStackHCI/operations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-12-01",
		},
		OperationID: "Operations_List",
		Resource:    "Microsoft.AzureStackHCI",
	},
}
