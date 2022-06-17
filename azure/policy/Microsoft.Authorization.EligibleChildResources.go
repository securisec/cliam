package policy

var Microsoft_Authorization_EligibleChildResources = map[string]Policy{
	"EligibleChildResources_Get": {
		Path:   "/{{.scope}}/providers/Microsoft.Authorization/eligibleChildResources",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2020-10-01",
		},
		OperationID: "EligibleChildResources_Get",
		Resource:    "Microsoft.Authorization",
	},
}
