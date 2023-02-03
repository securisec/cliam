package policy

// Microsoft_Automation_module policy
var Microsoft_Automation_module = map[string]Policy{
	"Activity_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Automation/automationAccounts/{{.automationAccountName}}/modules/{{.moduleName}}/activities/{{.activityName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-08-08",
		},
		OperationID: "Activity_Get",
		Resource:    "Microsoft.Automation",
	},
	"Activity_ListByModule": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Automation/automationAccounts/{{.automationAccountName}}/modules/{{.moduleName}}/activities",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-08-08",
		},
		OperationID: "Activity_ListByModule",
		Resource:    "Microsoft.Automation",
	},
	"Module_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Automation/automationAccounts/{{.automationAccountName}}/modules/{{.moduleName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-08-08",
		},
		OperationID: "Module_Get",
		Resource:    "Microsoft.Automation",
	},
	"Module_ListByAutomationAccount": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Automation/automationAccounts/{{.automationAccountName}}/modules",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-08-08",
		},
		OperationID: "Module_ListByAutomationAccount",
		Resource:    "Microsoft.Automation",
	},
	"ObjectDataTypes_ListFieldsByModuleAndType": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Automation/automationAccounts/{{.automationAccountName}}/modules/{{.moduleName}}/objectDataTypes/{{.typeName}}/fields",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-08-08",
		},
		OperationID: "ObjectDataTypes_ListFieldsByModuleAndType",
		Resource:    "Microsoft.Automation",
	},
	"ObjectDataTypes_ListFieldsByType": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Automation/automationAccounts/{{.automationAccountName}}/objectDataTypes/{{.typeName}}/fields",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-08-08",
		},
		OperationID: "ObjectDataTypes_ListFieldsByType",
		Resource:    "Microsoft.Automation",
	},
	"Fields_ListByType": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Automation/automationAccounts/{{.automationAccountName}}/modules/{{.moduleName}}/types/{{.typeName}}/fields",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-08-08",
		},
		OperationID: "Fields_ListByType",
		Resource:    "Microsoft.Automation",
	},
}
