package policy

// Microsoft_Automation_credential policy
var Microsoft_Automation_credential = map[string]Policy{
	"Credential_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Automation/automationAccounts/{{.automationAccountName}}/credentials/{{.credentialName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-08-08",
		},
		OperationID: "Credential_Get",
		Resource:    "Microsoft.Automation",
	},
	"Credential_ListByAutomationAccount": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Automation/automationAccounts/{{.automationAccountName}}/credentials",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-08-08",
		},
		OperationID: "Credential_ListByAutomationAccount",
		Resource:    "Microsoft.Automation",
	},
}
