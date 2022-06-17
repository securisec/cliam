package policy

var Microsoft_Compute_runCommands = map[string]Policy{
	"VirtualMachineRunCommands_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Compute/locations/{{.location}}/runCommands",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "VirtualMachineRunCommands_List",
		Resource:    "Microsoft.Compute",
	},
	"VirtualMachineRunCommands_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Compute/locations/{{.location}}/runCommands/{{.commandId}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "VirtualMachineRunCommands_Get",
		Resource:    "Microsoft.Compute",
	},
	"VirtualMachines_RunCommand": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Compute/virtualMachines/{{.vmName}}/runCommand",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "VirtualMachines_RunCommand",
		Resource:    "Microsoft.Compute",
	},
	"VirtualMachineScaleSetVMs_RunCommand": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Compute/virtualMachineScaleSets/{{.vmScaleSetName}}/virtualmachines/{{.instanceId}}/runCommand",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "VirtualMachineScaleSetVMs_RunCommand",
		Resource:    "Microsoft.Compute",
	},
	"VirtualMachineRunCommands_GetByVirtualMachine": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Compute/virtualMachines/{{.vmName}}/runCommands/{{.runCommandName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "VirtualMachineRunCommands_GetByVirtualMachine",
		Resource:    "Microsoft.Compute",
	},
	"VirtualMachineRunCommands_ListByVirtualMachine": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Compute/virtualMachines/{{.vmName}}/runCommands",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "VirtualMachineRunCommands_ListByVirtualMachine",
		Resource:    "Microsoft.Compute",
	},
	"VirtualMachineScaleSetVMRunCommands_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Compute/virtualMachineScaleSets/{{.vmScaleSetName}}/virtualMachines/{{.instanceId}}/runCommands/{{.runCommandName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "VirtualMachineScaleSetVMRunCommands_Get",
		Resource:    "Microsoft.Compute",
	},
	"VirtualMachineScaleSetVMRunCommands_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Compute/virtualMachineScaleSets/{{.vmScaleSetName}}/virtualMachines/{{.instanceId}}/runCommands",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "VirtualMachineScaleSetVMRunCommands_List",
		Resource:    "Microsoft.Compute",
	},
}
