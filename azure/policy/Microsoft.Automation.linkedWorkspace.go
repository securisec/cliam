package policy

// Microsoft_Automation_linkedWorkspace policy
var Microsoft_Automation_linkedWorkspace = map[string]Policy{
	"LinkedWorkspace_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Automation/automationAccounts/{{.automationAccountName}}/linkedWorkspace",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-08-08",
		},
		OperationID: "LinkedWorkspace_Get",
		Resource:    "Microsoft.Automation",
	},
}
