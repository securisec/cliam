package policy

// Microsoft_SqlVirtualMachine_sqlvm policy
var Microsoft_SqlVirtualMachine_sqlvm = map[string]Policy{
	"AvailabilityGroupListeners_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.SqlVirtualMachine/sqlVirtualMachineGroups/{{.sqlVirtualMachineGroupName}}/availabilityGroupListeners/{{.availabilityGroupListenerName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-02-01",
		},
		OperationID: "AvailabilityGroupListeners_Get",
		Resource:    "Microsoft.SqlVirtualMachine",
	},
	"AvailabilityGroupListeners_ListByGroup": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.SqlVirtualMachine/sqlVirtualMachineGroups/{{.sqlVirtualMachineGroupName}}/availabilityGroupListeners",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-02-01",
		},
		OperationID: "AvailabilityGroupListeners_ListByGroup",
		Resource:    "Microsoft.SqlVirtualMachine",
	},
	"Operations_List": {
		Path:   "/providers/Microsoft.SqlVirtualMachine/operations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-02-01",
		},
		OperationID: "Operations_List",
		Resource:    "Microsoft.SqlVirtualMachine",
	},
	"SqlVirtualMachineGroups_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.SqlVirtualMachine/sqlVirtualMachineGroups/{{.sqlVirtualMachineGroupName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-02-01",
		},
		OperationID: "SqlVirtualMachineGroups_Get",
		Resource:    "Microsoft.SqlVirtualMachine",
	},
	"SqlVirtualMachineGroups_ListByResourceGroup": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.SqlVirtualMachine/sqlVirtualMachineGroups",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-02-01",
		},
		OperationID: "SqlVirtualMachineGroups_ListByResourceGroup",
		Resource:    "Microsoft.SqlVirtualMachine",
	},
	"SqlVirtualMachineGroups_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.SqlVirtualMachine/sqlVirtualMachineGroups",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-02-01",
		},
		OperationID: "SqlVirtualMachineGroups_List",
		Resource:    "Microsoft.SqlVirtualMachine",
	},
	"SqlVirtualMachines_ListBySqlVmGroup": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.SqlVirtualMachine/sqlVirtualMachineGroups/{{.sqlVirtualMachineGroupName}}/sqlVirtualMachines",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-02-01",
		},
		OperationID: "SqlVirtualMachines_ListBySqlVmGroup",
		Resource:    "Microsoft.SqlVirtualMachine",
	},
	"SqlVirtualMachines_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.SqlVirtualMachine/sqlVirtualMachines",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-02-01",
		},
		OperationID: "SqlVirtualMachines_List",
		Resource:    "Microsoft.SqlVirtualMachine",
	},
	"SqlVirtualMachines_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.SqlVirtualMachine/sqlVirtualMachines/{{.sqlVirtualMachineName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-02-01",
		},
		OperationID: "SqlVirtualMachines_Get",
		Resource:    "Microsoft.SqlVirtualMachine",
	},
	"SqlVirtualMachines_ListByResourceGroup": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.SqlVirtualMachine/sqlVirtualMachines",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-02-01",
		},
		OperationID: "SqlVirtualMachines_ListByResourceGroup",
		Resource:    "Microsoft.SqlVirtualMachine",
	},
	"SqlVirtualMachines_Redeploy": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.SqlVirtualMachine/sqlVirtualMachines/{{.sqlVirtualMachineName}}/redeploy",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-02-01",
		},
		OperationID: "SqlVirtualMachines_Redeploy",
		Resource:    "Microsoft.SqlVirtualMachine",
	},
	"SqlVirtualMachines_StartAssessment": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.SqlVirtualMachine/sqlVirtualMachines/{{.sqlVirtualMachineName}}/startAssessment",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-02-01",
		},
		OperationID: "SqlVirtualMachines_StartAssessment",
		Resource:    "Microsoft.SqlVirtualMachine",
	},
}
