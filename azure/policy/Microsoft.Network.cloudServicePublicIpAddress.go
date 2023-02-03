package policy

// Microsoft_Network_cloudServicePublicIpAddress policy
var Microsoft_Network_cloudServicePublicIpAddress = map[string]Policy{
	"PublicIPAddresses_ListCloudServicePublicIPAddresses": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Compute/cloudServices/{{.cloudServiceName}}/publicipaddresses",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "PublicIPAddresses_ListCloudServicePublicIPAddresses",
		Resource:    "Microsoft.Network",
	},
	"PublicIPAddresses_ListCloudServiceRoleInstancePublicIPAddresses": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Compute/cloudServices/{{.cloudServiceName}}/roleInstances/{{.roleInstanceName}}/networkInterfaces/{{.networkInterfaceName}}/ipconfigurations/{{.ipConfigurationName}}/publicipaddresses",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "PublicIPAddresses_ListCloudServiceRoleInstancePublicIPAddresses",
		Resource:    "Microsoft.Network",
	},
	"PublicIPAddresses_GetCloudServicePublicIPAddress": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Compute/cloudServices/{{.cloudServiceName}}/roleInstances/{{.roleInstanceName}}/networkInterfaces/{{.networkInterfaceName}}/ipconfigurations/{{.ipConfigurationName}}/publicipaddresses/{{.publicIpAddressName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "PublicIPAddresses_GetCloudServicePublicIPAddress",
		Resource:    "Microsoft.Network",
	},
}
