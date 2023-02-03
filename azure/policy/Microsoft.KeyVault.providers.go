package policy

// Microsoft_KeyVault_providers policy
var Microsoft_KeyVault_providers = map[string]Policy{
	"Operations_List": {
		Path:   "/providers/Microsoft.KeyVault/operations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "Operations_List",
		Resource:    "Microsoft.KeyVault",
	},
}
