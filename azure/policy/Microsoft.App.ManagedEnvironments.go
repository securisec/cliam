package policy

var Microsoft_App_ManagedEnvironments = map[string]Policy{
	"ManagedEnvironments_ListBySubscription": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.App/managedEnvironments",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "ManagedEnvironments_ListBySubscription",
		Resource:    "Microsoft.App",
	},
	"ManagedEnvironments_ListByResourceGroup": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.App/managedEnvironments",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "ManagedEnvironments_ListByResourceGroup",
		Resource:    "Microsoft.App",
	},
	"ManagedEnvironments_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.App/managedEnvironments/{{.environmentName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "ManagedEnvironments_Get",
		Resource:    "Microsoft.App",
	},
	"Certificates_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.App/managedEnvironments/{{.environmentName}}/certificates",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "Certificates_List",
		Resource:    "Microsoft.App",
	},
	"Certificates_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.App/managedEnvironments/{{.environmentName}}/certificates/{{.certificateName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "Certificates_Get",
		Resource:    "Microsoft.App",
	},
	"Namespaces_CheckNameAvailability": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.App/managedEnvironments/{{.environmentName}}/checkNameAvailability",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "Namespaces_CheckNameAvailability",
		Resource:    "Microsoft.App",
	},
}
