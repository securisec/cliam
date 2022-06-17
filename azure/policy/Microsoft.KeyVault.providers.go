package policy

    var Microsoft_KeyVault_providers = []Policy{
        {
    Path: "/providers/Microsoft.KeyVault/operations",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2021-10-01",
    },
	OperationID:    "Operations_List",
    Resource:       "Microsoft.KeyVault",
},
    }
    