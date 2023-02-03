package policy

// Microsoft_Network_vmssNetworkInterface policy
var Microsoft_Network_vmssNetworkInterface = map[string]Policy{
	"NetworkInterfaces_ListVirtualMachineScaleSetVMNetworkInterfaces": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/microsoft.Compute/virtualMachineScaleSets/{{.virtualMachineScaleSetName}}/virtualMachines/{{.virtualmachineIndex}}/networkInterfaces",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "NetworkInterfaces_ListVirtualMachineScaleSetVMNetworkInterfaces",
		Resource:    "Microsoft.Network",
	},
	"NetworkInterfaces_ListVirtualMachineScaleSetNetworkInterfaces": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/microsoft.Compute/virtualMachineScaleSets/{{.virtualMachineScaleSetName}}/networkInterfaces",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "NetworkInterfaces_ListVirtualMachineScaleSetNetworkInterfaces",
		Resource:    "Microsoft.Network",
	},
	"NetworkInterfaces_GetVirtualMachineScaleSetNetworkInterface": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/microsoft.Compute/virtualMachineScaleSets/{{.virtualMachineScaleSetName}}/virtualMachines/{{.virtualmachineIndex}}/networkInterfaces/{{.networkInterfaceName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "NetworkInterfaces_GetVirtualMachineScaleSetNetworkInterface",
		Resource:    "Microsoft.Network",
	},
	"NetworkInterfaces_ListVirtualMachineScaleSetIpConfigurations": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/microsoft.Compute/virtualMachineScaleSets/{{.virtualMachineScaleSetName}}/virtualMachines/{{.virtualmachineIndex}}/networkInterfaces/{{.networkInterfaceName}}/ipConfigurations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "NetworkInterfaces_ListVirtualMachineScaleSetIpConfigurations",
		Resource:    "Microsoft.Network",
	},
	"NetworkInterfaces_GetVirtualMachineScaleSetIpConfiguration": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/microsoft.Compute/virtualMachineScaleSets/{{.virtualMachineScaleSetName}}/virtualMachines/{{.virtualmachineIndex}}/networkInterfaces/{{.networkInterfaceName}}/ipConfigurations/{{.ipConfigurationName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "NetworkInterfaces_GetVirtualMachineScaleSetIpConfiguration",
		Resource:    "Microsoft.Network",
	},
}
