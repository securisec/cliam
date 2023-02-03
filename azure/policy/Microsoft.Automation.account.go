package policy

// Microsoft_Automation_account policy
var Microsoft_Automation_account = map[string]Policy{
	"AutomationAccount_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Automation/automationAccounts/{{.automationAccountName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-08-08",
		},
		OperationID: "AutomationAccount_Get",
		Resource:    "Microsoft.Automation",
	},
	"AutomationAccount_ListByResourceGroup": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Automation/automationAccounts",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-08-08",
		},
		OperationID: "AutomationAccount_ListByResourceGroup",
		Resource:    "Microsoft.Automation",
	},
	"AutomationAccount_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Automation/automationAccounts",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-08-08",
		},
		OperationID: "AutomationAccount_List",
		Resource:    "Microsoft.Automation",
	},
	"Statistics_ListByAutomationAccount": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Automation/automationAccounts/{{.automationAccountName}}/statistics",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-08-08",
		},
		OperationID: "Statistics_ListByAutomationAccount",
		Resource:    "Microsoft.Automation",
	},
	"Usages_ListByAutomationAccount": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Automation/automationAccounts/{{.automationAccountName}}/usages",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-08-08",
		},
		OperationID: "Usages_ListByAutomationAccount",
		Resource:    "Microsoft.Automation",
	},
	"Keys_ListByAutomationAccount": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Automation/automationAccounts/{{.automationAccountName}}/listKeys",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-08-08",
		},
		OperationID: "Keys_ListByAutomationAccount",
		Resource:    "Microsoft.Automation",
	},
}
