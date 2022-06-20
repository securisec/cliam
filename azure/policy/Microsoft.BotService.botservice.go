package policy

var Microsoft_BotService_botservice = map[string]Policy{
	"Bots_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.BotService/botServices/{{.resourceName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "Bots_Get",
		Resource:    "Microsoft.BotService",
	},
	"Bots_ListByResourceGroup": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.BotService/botServices",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "Bots_ListByResourceGroup",
		Resource:    "Microsoft.BotService",
	},
	"Bots_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.BotService/botServices",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "Bots_List",
		Resource:    "Microsoft.BotService",
	},
	"Channels_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.BotService/botServices/{{.resourceName}}/channels/{{.channelName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "Channels_Get",
		Resource:    "Microsoft.BotService",
	},
	"Channels_ListWithKeys": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.BotService/botServices/{{.resourceName}}/channels/{{.channelName}}/listChannelWithKeys",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "Channels_ListWithKeys",
		Resource:    "Microsoft.BotService",
	},
	"Channels_ListByResourceGroup": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.BotService/botServices/{{.resourceName}}/channels",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "Channels_ListByResourceGroup",
		Resource:    "Microsoft.BotService",
	},
	"DirectLine_RegenerateKeys": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.BotService/botServices/{{.resourceName}}/channels/{{.channelName}}/regeneratekeys",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "DirectLine_RegenerateKeys",
		Resource:    "Microsoft.BotService",
	},
	"Bots_GetCheckNameAvailability": {
		Path:   "/providers/Microsoft.BotService/checkNameAvailability",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "Bots_GetCheckNameAvailability",
		Resource:    "Microsoft.BotService",
	},
	"Operations_List": {
		Path:   "/providers/Microsoft.BotService/operations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "Operations_List",
		Resource:    "Microsoft.BotService",
	},
	"BotConnection_ListServiceProviders": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.BotService/listAuthServiceProviders",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "BotConnection_ListServiceProviders",
		Resource:    "Microsoft.BotService",
	},
	"BotConnection_ListWithSecrets": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.BotService/botServices/{{.resourceName}}/connections/{{.connectionName}}/listWithSecrets",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "BotConnection_ListWithSecrets",
		Resource:    "Microsoft.BotService",
	},
	"BotConnection_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.BotService/botServices/{{.resourceName}}/connections/{{.connectionName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "BotConnection_Get",
		Resource:    "Microsoft.BotService",
	},
	"BotConnection_ListByBotService": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.BotService/botServices/{{.resourceName}}/connections",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "BotConnection_ListByBotService",
		Resource:    "Microsoft.BotService",
	},
	"HostSettings_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.BotService/hostSettings",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "HostSettings_Get",
		Resource:    "Microsoft.BotService",
	},
}
