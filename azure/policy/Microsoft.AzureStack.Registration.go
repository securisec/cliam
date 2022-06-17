package policy

var Microsoft_AzureStack_Registration = map[string]Policy{
	"Registrations_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroup}}/providers/Microsoft.AzureStack/registrations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2017-06-01",
		},
		OperationID: "Registrations_List",
		Resource:    "Microsoft.AzureStack",
	},
	"Registrations_ListBySubscription": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.AzureStack/registrations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2017-06-01",
		},
		OperationID: "Registrations_ListBySubscription",
		Resource:    "Microsoft.AzureStack",
	},
	"Registrations_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroup}}/providers/Microsoft.AzureStack/registrations/{{.registrationName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2017-06-01",
		},
		OperationID: "Registrations_Get",
		Resource:    "Microsoft.AzureStack",
	},
	"Registrations_GetActivationKey": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroup}}/providers/Microsoft.AzureStack/registrations/{{.registrationName}}/getactivationkey",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2017-06-01",
		},
		OperationID: "Registrations_GetActivationKey",
		Resource:    "Microsoft.AzureStack",
	},
	"Registrations_EnableRemoteManagement": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroup}}/providers/Microsoft.AzureStack/registrations/{{.registrationName}}/enableRemoteManagement",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2017-06-01",
		},
		OperationID: "Registrations_EnableRemoteManagement",
		Resource:    "Microsoft.AzureStack",
	},
}
