package policy

// Microsoft_Authorization_authorization_ProviderOperationsCalls policy
var Microsoft_Authorization_authorization_ProviderOperationsCalls = map[string]Policy{
	"ProviderOperationsMetadata_Get": {
		Path:   "/providers/Microsoft.Authorization/providerOperations/{{.resourceProviderNamespace}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-04-01",
		},
		OperationID: "ProviderOperationsMetadata_Get",
		Resource:    "Microsoft.Authorization",
	},
	"ProviderOperationsMetadata_List": {
		Path:   "/providers/Microsoft.Authorization/providerOperations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-04-01",
		},
		OperationID: "ProviderOperationsMetadata_List",
		Resource:    "Microsoft.Authorization",
	},
}
