package policy

// Microsoft_Network_virtualNetworkTap policy
var Microsoft_Network_virtualNetworkTap = map[string]Policy{
	"VirtualNetworkTaps_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/virtualNetworkTaps/{{.tapName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "VirtualNetworkTaps_Get",
		Resource:    "Microsoft.Network",
	},
	"VirtualNetworkTaps_ListAll": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Network/virtualNetworkTaps",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "VirtualNetworkTaps_ListAll",
		Resource:    "Microsoft.Network",
	},
	"VirtualNetworkTaps_ListByResourceGroup": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/virtualNetworkTaps",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "VirtualNetworkTaps_ListByResourceGroup",
		Resource:    "Microsoft.Network",
	},
}
