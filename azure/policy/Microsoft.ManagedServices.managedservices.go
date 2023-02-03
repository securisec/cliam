package policy

// Microsoft_ManagedServices_managedservices policy
var Microsoft_ManagedServices_managedservices = map[string]Policy{
	"RegistrationDefinitions_Get": {
		Path:   "/{{.scope}}/providers/Microsoft.ManagedServices/registrationDefinitions/{{.registrationDefinitionId}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "RegistrationDefinitions_Get",
		Resource:    "Microsoft.ManagedServices",
	},
	"RegistrationAssignments_Get": {
		Path:   "/{{.scope}}/providers/Microsoft.ManagedServices/registrationAssignments/{{.registrationAssignmentId}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "RegistrationAssignments_Get",
		Resource:    "Microsoft.ManagedServices",
	},
	"RegistrationDefinitions_List": {
		Path:   "/{{.scope}}/providers/Microsoft.ManagedServices/registrationDefinitions",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "RegistrationDefinitions_List",
		Resource:    "Microsoft.ManagedServices",
	},
	"RegistrationAssignments_List": {
		Path:   "/{{.scope}}/providers/Microsoft.ManagedServices/registrationAssignments",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "RegistrationAssignments_List",
		Resource:    "Microsoft.ManagedServices",
	},
	"MarketplaceRegistrationDefinitions_List": {
		Path:   "/{{.scope}}/providers/Microsoft.ManagedServices/marketplaceRegistrationDefinitions",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "MarketplaceRegistrationDefinitions_List",
		Resource:    "Microsoft.ManagedServices",
	},
	"MarketplaceRegistrationDefinitionsWithoutScope_List": {
		Path:   "/providers/Microsoft.ManagedServices/marketplaceRegistrationDefinitions",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "MarketplaceRegistrationDefinitionsWithoutScope_List",
		Resource:    "Microsoft.ManagedServices",
	},
	"MarketplaceRegistrationDefinitions_Get": {
		Path:   "/{{.scope}}/providers/Microsoft.ManagedServices/marketplaceRegistrationDefinitions/{{.marketplaceIdentifier}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "MarketplaceRegistrationDefinitions_Get",
		Resource:    "Microsoft.ManagedServices",
	},
	"MarketplaceRegistrationDefinitionsWithoutScope_Get": {
		Path:   "/providers/Microsoft.ManagedServices/marketplaceRegistrationDefinitions/{{.marketplaceIdentifier}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "MarketplaceRegistrationDefinitionsWithoutScope_Get",
		Resource:    "Microsoft.ManagedServices",
	},
	"Operations_List": {
		Path:   "/providers/Microsoft.ManagedServices/operations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "Operations_List",
		Resource:    "Microsoft.ManagedServices",
	},
	"OperationsWithScope_List": {
		Path:   "/{{.scope}}/providers/Microsoft.ManagedServices/operations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "OperationsWithScope_List",
		Resource:    "Microsoft.ManagedServices",
	},
}
