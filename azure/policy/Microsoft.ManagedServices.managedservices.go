package policy

var Microsoft_ManagedServices_managedservices = []Policy{
	{
		Path:   "/{{.scope}}/providers/Microsoft.ManagedServices/registrationDefinitions/{{.registrationDefinitionId}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2019-09-01",
		},
		OperationID: "RegistrationDefinitions_Get",
		Resource:    "Microsoft.ManagedServices",
	}, {
		Path:   "/{{.scope}}/providers/Microsoft.ManagedServices/registrationAssignments/{{.registrationAssignmentId}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2019-09-01",
		},
		OperationID: "RegistrationAssignments_Get",
		Resource:    "Microsoft.ManagedServices",
	}, {
		Path:   "/{{.scope}}/providers/Microsoft.ManagedServices/registrationDefinitions",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2019-09-01",
		},
		OperationID: "RegistrationDefinitions_List",
		Resource:    "Microsoft.ManagedServices",
	}, {
		Path:   "/{{.scope}}/providers/Microsoft.ManagedServices/registrationAssignments",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2019-09-01",
		},
		OperationID: "RegistrationAssignments_List",
		Resource:    "Microsoft.ManagedServices",
	}, {
		Path:   "/{{.scope}}/providers/Microsoft.ManagedServices/marketplaceRegistrationDefinitions",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2019-09-01",
		},
		OperationID: "MarketplaceRegistrationDefinitions_List",
		Resource:    "Microsoft.ManagedServices",
	}, {
		Path:   "/providers/Microsoft.ManagedServices/marketplaceRegistrationDefinitions",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2019-09-01",
		},
		OperationID: "MarketplaceRegistrationDefinitionsWithoutScope_List",
		Resource:    "Microsoft.ManagedServices",
	}, {
		Path:   "/{{.scope}}/providers/Microsoft.ManagedServices/marketplaceRegistrationDefinitions/{{.marketplaceIdentifier}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2019-09-01",
		},
		OperationID: "MarketplaceRegistrationDefinitions_Get",
		Resource:    "Microsoft.ManagedServices",
	}, {
		Path:   "/providers/Microsoft.ManagedServices/marketplaceRegistrationDefinitions/{{.marketplaceIdentifier}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2019-09-01",
		},
		OperationID: "MarketplaceRegistrationDefinitionsWithoutScope_Get",
		Resource:    "Microsoft.ManagedServices",
	}, {
		Path:   "/providers/Microsoft.ManagedServices/operations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2019-09-01",
		},
		OperationID: "Operations_List",
		Resource:    "Microsoft.ManagedServices",
	},
}
