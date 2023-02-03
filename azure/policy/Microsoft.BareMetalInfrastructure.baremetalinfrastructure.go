package policy

// Microsoft_BareMetalInfrastructure_baremetalinfrastructure policy
var Microsoft_BareMetalInfrastructure_baremetalinfrastructure = map[string]Policy{
	"AzureBareMetalInstances_ListBySubscription": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.BareMetalInfrastructure/bareMetalInstances",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-08-09",
		},
		OperationID: "AzureBareMetalInstances_ListBySubscription",
		Resource:    "Microsoft.BareMetalInfrastructure",
	},
	"AzureBareMetalInstances_ListByResourceGroup": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.BareMetalInfrastructure/bareMetalInstances",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-08-09",
		},
		OperationID: "AzureBareMetalInstances_ListByResourceGroup",
		Resource:    "Microsoft.BareMetalInfrastructure",
	},
	"Operations_List": {
		Path:   "/providers/Microsoft.BareMetalInfrastructure/operations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-08-09",
		},
		OperationID: "Operations_List",
		Resource:    "Microsoft.BareMetalInfrastructure",
	},
	"AzureBareMetalInstances_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.BareMetalInfrastructure/bareMetalInstances/{{.azureBareMetalInstanceName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-08-09",
		},
		OperationID: "AzureBareMetalInstances_Get",
		Resource:    "Microsoft.BareMetalInfrastructure",
	},
}
