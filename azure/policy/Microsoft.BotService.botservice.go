package policy

// Microsoft_BotService_botservice policy
var Microsoft_BotService_botservice = map[string]Policy{
	"Bots_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.BotService/botServices/{{.resourceName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-09-15",
		},
		OperationID: "Bots_Get",
		Resource:    "Microsoft.BotService",
	},
	"Bots_ListByResourceGroup": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.BotService/botServices",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-09-15",
		},
		OperationID: "Bots_ListByResourceGroup",
		Resource:    "Microsoft.BotService",
	},
	"Bots_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.BotService/botServices",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-09-15",
		},
		OperationID: "Bots_List",
		Resource:    "Microsoft.BotService",
	},
	"Channels_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.BotService/botServices/{{.resourceName}}/channels/{{.channelName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-09-15",
		},
		OperationID: "Channels_Get",
		Resource:    "Microsoft.BotService",
	},
	"Channels_ListWithKeys": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.BotService/botServices/{{.resourceName}}/channels/{{.channelName}}/listChannelWithKeys",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-09-15",
		},
		OperationID: "Channels_ListWithKeys",
		Resource:    "Microsoft.BotService",
	},
	"Channels_ListByResourceGroup": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.BotService/botServices/{{.resourceName}}/channels",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-09-15",
		},
		OperationID: "Channels_ListByResourceGroup",
		Resource:    "Microsoft.BotService",
	},
	"DirectLine_RegenerateKeys": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.BotService/botServices/{{.resourceName}}/channels/{{.channelName}}/regeneratekeys",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-09-15",
		},
		OperationID: "DirectLine_RegenerateKeys",
		Resource:    "Microsoft.BotService",
	},
	"Email_CreateSignInUrl": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.BotService/botServices/{{.resourceName}}/createEmailSignInUrl",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-09-15",
		},
		OperationID: "Email_CreateSignInUrl",
		Resource:    "Microsoft.BotService",
	},
	"Bots_GetCheckNameAvailability": {
		Path:   "/providers/Microsoft.BotService/checkNameAvailability",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-09-15",
		},
		OperationID: "Bots_GetCheckNameAvailability",
		Resource:    "Microsoft.BotService",
	},
	"Operations_List": {
		Path:   "/providers/Microsoft.BotService/operations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-09-15",
		},
		OperationID: "Operations_List",
		Resource:    "Microsoft.BotService",
	},
	"BotConnection_ListServiceProviders": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.BotService/listAuthServiceProviders",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-09-15",
		},
		OperationID: "BotConnection_ListServiceProviders",
		Resource:    "Microsoft.BotService",
	},
	"QnAMakerEndpointKeys_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.BotService/listQnAMakerEndpointKeys",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-09-15",
		},
		OperationID: "QnAMakerEndpointKeys_Get",
		Resource:    "Microsoft.BotService",
	},
	"BotConnection_ListWithSecrets": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.BotService/botServices/{{.resourceName}}/connections/{{.connectionName}}/listWithSecrets",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-09-15",
		},
		OperationID: "BotConnection_ListWithSecrets",
		Resource:    "Microsoft.BotService",
	},
	"BotConnection_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.BotService/botServices/{{.resourceName}}/connections/{{.connectionName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-09-15",
		},
		OperationID: "BotConnection_Get",
		Resource:    "Microsoft.BotService",
	},
	"BotConnection_ListByBotService": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.BotService/botServices/{{.resourceName}}/connections",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-09-15",
		},
		OperationID: "BotConnection_ListByBotService",
		Resource:    "Microsoft.BotService",
	},
	"HostSettings_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.BotService/hostSettings",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-09-15",
		},
		OperationID: "HostSettings_Get",
		Resource:    "Microsoft.BotService",
	},
	"OperationResults_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.BotService/operationresults/{{.operationResultId}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-09-15",
		},
		OperationID: "OperationResults_Get",
		Resource:    "Microsoft.BotService",
	},
	"PrivateEndpointConnections_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.BotService/botServices/{{.resourceName}}/privateEndpointConnections",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-09-15",
		},
		OperationID: "PrivateEndpointConnections_List",
		Resource:    "Microsoft.BotService",
	},
	"PrivateEndpointConnections_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.BotService/botServices/{{.resourceName}}/privateEndpointConnections/{{.privateEndpointConnectionName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-09-15",
		},
		OperationID: "PrivateEndpointConnections_Get",
		Resource:    "Microsoft.BotService",
	},
	"PrivateLinkResources_ListByBotResource": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.BotService/botServices/{{.resourceName}}/privateLinkResources",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-09-15",
		},
		OperationID: "PrivateLinkResources_ListByBotResource",
		Resource:    "Microsoft.BotService",
	},
}
