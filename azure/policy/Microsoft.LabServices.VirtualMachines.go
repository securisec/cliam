package policy

// Microsoft_LabServices_VirtualMachines policy
var Microsoft_LabServices_VirtualMachines = map[string]Policy{
	"VirtualMachines_ListByLab": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.LabServices/labs/{{.labName}}/virtualMachines",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-08-01",
		},
		OperationID: "VirtualMachines_ListByLab",
		Resource:    "Microsoft.LabServices",
	},
	"VirtualMachines_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.LabServices/labs/{{.labName}}/virtualMachines/{{.virtualMachineName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-08-01",
		},
		OperationID: "VirtualMachines_Get",
		Resource:    "Microsoft.LabServices",
	},
	"VirtualMachines_Start": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.LabServices/labs/{{.labName}}/virtualMachines/{{.virtualMachineName}}/start",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-08-01",
		},
		OperationID: "VirtualMachines_Start",
		Resource:    "Microsoft.LabServices",
	},
	"VirtualMachines_Stop": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.LabServices/labs/{{.labName}}/virtualMachines/{{.virtualMachineName}}/stop",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-08-01",
		},
		OperationID: "VirtualMachines_Stop",
		Resource:    "Microsoft.LabServices",
	},
	"VirtualMachines_Reimage": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.LabServices/labs/{{.labName}}/virtualMachines/{{.virtualMachineName}}/reimage",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-08-01",
		},
		OperationID: "VirtualMachines_Reimage",
		Resource:    "Microsoft.LabServices",
	},
	"VirtualMachines_Redeploy": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.LabServices/labs/{{.labName}}/virtualMachines/{{.virtualMachineName}}/redeploy",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-08-01",
		},
		OperationID: "VirtualMachines_Redeploy",
		Resource:    "Microsoft.LabServices",
	},
	"VirtualMachines_ResetPassword": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.LabServices/labs/{{.labName}}/virtualMachines/{{.virtualMachineName}}/resetPassword",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-08-01",
		},
		OperationID: "VirtualMachines_ResetPassword",
		Resource:    "Microsoft.LabServices",
	},
}
