package policy

var Microsoft_Authorization_EligibleChildResources = []Policy{
	{
		Path:   "/{{.scope}}/providers/Microsoft.Authorization/eligibleChildResources",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2020-10-01",
		},
		OperationID: "EligibleChildResources_Get",
		Resource:    "Microsoft.Authorization",
	},
}
