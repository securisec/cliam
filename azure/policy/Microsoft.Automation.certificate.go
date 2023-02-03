package policy

// Microsoft_Automation_certificate policy
var Microsoft_Automation_certificate = map[string]Policy{
	"Certificate_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Automation/automationAccounts/{{.automationAccountName}}/certificates/{{.certificateName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-08-08",
		},
		OperationID: "Certificate_Get",
		Resource:    "Microsoft.Automation",
	},
	"Certificate_ListByAutomationAccount": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Automation/automationAccounts/{{.automationAccountName}}/certificates",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-08-08",
		},
		OperationID: "Certificate_ListByAutomationAccount",
		Resource:    "Microsoft.Automation",
	},
}
