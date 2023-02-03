package policy

// Microsoft_Confluent_confluent policy
var Microsoft_Confluent_confluent = map[string]Policy{
	"MarketplaceAgreements_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Confluent/agreements",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-12-01",
		},
		OperationID: "MarketplaceAgreements_List",
		Resource:    "Microsoft.Confluent",
	},
	"OrganizationOperations_List": {
		Path:   "/providers/Microsoft.Confluent/operations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-12-01",
		},
		OperationID: "OrganizationOperations_List",
		Resource:    "Microsoft.Confluent",
	},
	"Organization_ListBySubscription": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Confluent/organizations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-12-01",
		},
		OperationID: "Organization_ListBySubscription",
		Resource:    "Microsoft.Confluent",
	},
	"Organization_ListByResourceGroup": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Confluent/organizations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-12-01",
		},
		OperationID: "Organization_ListByResourceGroup",
		Resource:    "Microsoft.Confluent",
	},
	"Organization_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Confluent/organizations/{{.organizationName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-12-01",
		},
		OperationID: "Organization_Get",
		Resource:    "Microsoft.Confluent",
	},
	"Validations_ValidateOrganization": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Confluent/validations/{{.organizationName}}/orgvalidate",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-12-01",
		},
		OperationID: "Validations_ValidateOrganization",
		Resource:    "Microsoft.Confluent",
	},
}
