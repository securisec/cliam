package policy

var Microsoft_Communication_CommunicationService = map[string]Policy{
	"Operations_List": {
		Path:   "/providers/Microsoft.Communication/operations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2020-08-20",
		},
		OperationID: "Operations_List",
		Resource:    "Microsoft.Communication",
	},
	"CommunicationService_CheckNameAvailability": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Communication/checkNameAvailability",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2020-08-20",
		},
		OperationID: "CommunicationService_CheckNameAvailability",
		Resource:    "Microsoft.Communication",
	},
	"CommunicationService_LinkNotificationHub": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Communication/communicationServices/{{.communicationServiceName}}/linkNotificationHub",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2020-08-20",
		},
		OperationID: "CommunicationService_LinkNotificationHub",
		Resource:    "Microsoft.Communication",
	},
	"CommunicationService_ListBySubscription": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Communication/communicationServices",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2020-08-20",
		},
		OperationID: "CommunicationService_ListBySubscription",
		Resource:    "Microsoft.Communication",
	},
	"CommunicationService_ListByResourceGroup": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Communication/communicationServices",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2020-08-20",
		},
		OperationID: "CommunicationService_ListByResourceGroup",
		Resource:    "Microsoft.Communication",
	},
	"CommunicationService_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Communication/communicationServices/{{.communicationServiceName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2020-08-20",
		},
		OperationID: "CommunicationService_Get",
		Resource:    "Microsoft.Communication",
	},
	"CommunicationService_ListKeys": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Communication/communicationServices/{{.communicationServiceName}}/listKeys",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2020-08-20",
		},
		OperationID: "CommunicationService_ListKeys",
		Resource:    "Microsoft.Communication",
	},
	"CommunicationService_RegenerateKey": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Communication/communicationServices/{{.communicationServiceName}}/regenerateKey",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2020-08-20",
		},
		OperationID: "CommunicationService_RegenerateKey",
		Resource:    "Microsoft.Communication",
	},
}
