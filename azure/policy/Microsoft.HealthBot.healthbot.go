package policy

// Microsoft_HealthBot_healthbot policy
var Microsoft_HealthBot_healthbot = map[string]Policy{
	"Bots_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.HealthBot/healthBots/{{.botName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-08-08",
		},
		OperationID: "Bots_Get",
		Resource:    "Microsoft.HealthBot",
	},
	"Bots_ListSecrets": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.HealthBot/healthBots/{{.botName}}/listSecrets",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-08-08",
		},
		OperationID: "Bots_ListSecrets",
		Resource:    "Microsoft.HealthBot",
	},
	"Bots_RegenerateApiJwtSecret": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.HealthBot/healthBots/{{.botName}}/regenerateApiJwtSecret",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-08-08",
		},
		OperationID: "Bots_RegenerateApiJwtSecret",
		Resource:    "Microsoft.HealthBot",
	},
	"Bots_ListByResourceGroup": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.HealthBot/healthBots",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-08-08",
		},
		OperationID: "Bots_ListByResourceGroup",
		Resource:    "Microsoft.HealthBot",
	},
	"Bots_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.HealthBot/healthBots",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-08-08",
		},
		OperationID: "Bots_List",
		Resource:    "Microsoft.HealthBot",
	},
	"Operations_List": {
		Path:   "/providers/Microsoft.HealthBot/operations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-08-08",
		},
		OperationID: "Operations_List",
		Resource:    "Microsoft.HealthBot",
	},
}
