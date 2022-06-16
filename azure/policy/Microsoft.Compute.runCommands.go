package policy

var Microsoft_Compute_runCommands = []Policy{
	{
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Compute/locations/{{.location}}/runCommands",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "VirtualMachineRunCommands_List",
		Resource:    "Microsoft.Compute",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Compute/locations/{{.location}}/runCommands/{{.commandId}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "VirtualMachineRunCommands_Get",
		Resource:    "Microsoft.Compute",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Compute/virtualMachines/{{.vmName}}/runCommand",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "VirtualMachines_RunCommand",
		Resource:    "Microsoft.Compute",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Compute/virtualMachineScaleSets/{{.vmScaleSetName}}/virtualmachines/{{.instanceId}}/runCommand",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "VirtualMachineScaleSetVMs_RunCommand",
		Resource:    "Microsoft.Compute",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Compute/virtualMachines/{{.vmName}}/runCommands/{{.runCommandName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "VirtualMachineRunCommands_GetByVirtualMachine",
		Resource:    "Microsoft.Compute",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Compute/virtualMachines/{{.vmName}}/runCommands",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "VirtualMachineRunCommands_ListByVirtualMachine",
		Resource:    "Microsoft.Compute",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Compute/virtualMachineScaleSets/{{.vmScaleSetName}}/virtualMachines/{{.instanceId}}/runCommands/{{.runCommandName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "VirtualMachineScaleSetVMRunCommands_Get",
		Resource:    "Microsoft.Compute",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Compute/virtualMachineScaleSets/{{.vmScaleSetName}}/virtualMachines/{{.instanceId}}/runCommands",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "VirtualMachineScaleSetVMRunCommands_List",
		Resource:    "Microsoft.Compute",
	},
}
