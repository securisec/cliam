package policy

// Microsoft_Network_vmssPublicIpAddress policy
var Microsoft_Network_vmssPublicIpAddress = map[string]Policy{
	"PublicIPAddresses_ListVirtualMachineScaleSetPublicIPAddresses": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Compute/virtualMachineScaleSets/{{.virtualMachineScaleSetName}}/publicipaddresses",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "PublicIPAddresses_ListVirtualMachineScaleSetPublicIPAddresses",
		Resource:    "Microsoft.Network",
	},
	"PublicIPAddresses_ListVirtualMachineScaleSetVMPublicIPAddresses": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Compute/virtualMachineScaleSets/{{.virtualMachineScaleSetName}}/virtualMachines/{{.virtualmachineIndex}}/networkInterfaces/{{.networkInterfaceName}}/ipconfigurations/{{.ipConfigurationName}}/publicipaddresses",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "PublicIPAddresses_ListVirtualMachineScaleSetVMPublicIPAddresses",
		Resource:    "Microsoft.Network",
	},
	"PublicIPAddresses_GetVirtualMachineScaleSetPublicIPAddress": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Compute/virtualMachineScaleSets/{{.virtualMachineScaleSetName}}/virtualMachines/{{.virtualmachineIndex}}/networkInterfaces/{{.networkInterfaceName}}/ipconfigurations/{{.ipConfigurationName}}/publicipaddresses/{{.publicIpAddressName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "PublicIPAddresses_GetVirtualMachineScaleSetPublicIPAddress",
		Resource:    "Microsoft.Network",
	},
}
