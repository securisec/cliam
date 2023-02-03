package policy

// Microsoft_Network_cloudServiceNetworkInterface policy
var Microsoft_Network_cloudServiceNetworkInterface = map[string]Policy{
	"NetworkInterfaces_ListCloudServiceRoleInstanceNetworkInterfaces": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Compute/cloudServices/{{.cloudServiceName}}/roleInstances/{{.roleInstanceName}}/networkInterfaces",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "NetworkInterfaces_ListCloudServiceRoleInstanceNetworkInterfaces",
		Resource:    "Microsoft.Network",
	},
	"NetworkInterfaces_ListCloudServiceNetworkInterfaces": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Compute/cloudServices/{{.cloudServiceName}}/networkInterfaces",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "NetworkInterfaces_ListCloudServiceNetworkInterfaces",
		Resource:    "Microsoft.Network",
	},
	"NetworkInterfaces_GetCloudServiceNetworkInterface": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Compute/cloudServices/{{.cloudServiceName}}/roleInstances/{{.roleInstanceName}}/networkInterfaces/{{.networkInterfaceName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "NetworkInterfaces_GetCloudServiceNetworkInterface",
		Resource:    "Microsoft.Network",
	},
}
