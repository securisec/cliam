package policy

// Microsoft_VoiceServices_voiceservices policy
var Microsoft_VoiceServices_voiceservices = map[string]Policy{
	"Operations_List": {
		Path:   "/providers/Microsoft.VoiceServices/operations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2023-01-31",
		},
		OperationID: "Operations_List",
		Resource:    "Microsoft.VoiceServices",
	},
	"CommunicationsGateways_ListBySubscription": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.VoiceServices/communicationsGateways",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2023-01-31",
		},
		OperationID: "CommunicationsGateways_ListBySubscription",
		Resource:    "Microsoft.VoiceServices",
	},
	"NameAvailability_CheckLocal": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.VoiceServices/locations/{{.location}}/checkNameAvailability",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2023-01-31",
		},
		OperationID: "NameAvailability_CheckLocal",
		Resource:    "Microsoft.VoiceServices",
	},
	"CommunicationsGateways_ListByResourceGroup": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.VoiceServices/communicationsGateways",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2023-01-31",
		},
		OperationID: "CommunicationsGateways_ListByResourceGroup",
		Resource:    "Microsoft.VoiceServices",
	},
	"CommunicationsGateways_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.VoiceServices/communicationsGateways/{{.communicationsGatewayName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2023-01-31",
		},
		OperationID: "CommunicationsGateways_Get",
		Resource:    "Microsoft.VoiceServices",
	},
	"TestLines_ListByCommunicationsGateway": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.VoiceServices/communicationsGateways/{{.communicationsGatewayName}}/testLines",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2023-01-31",
		},
		OperationID: "TestLines_ListByCommunicationsGateway",
		Resource:    "Microsoft.VoiceServices",
	},
	"TestLines_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.VoiceServices/communicationsGateways/{{.communicationsGatewayName}}/testLines/{{.testLineName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2023-01-31",
		},
		OperationID: "TestLines_Get",
		Resource:    "Microsoft.VoiceServices",
	},
}
