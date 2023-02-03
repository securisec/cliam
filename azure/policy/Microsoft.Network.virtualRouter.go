package policy

// Microsoft_Network_virtualRouter policy
var Microsoft_Network_virtualRouter = map[string]Policy{
	"VirtualRouters_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/virtualRouters/{{.virtualRouterName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "VirtualRouters_Get",
		Resource:    "Microsoft.Network",
	},
	"VirtualRouters_ListByResourceGroup": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/virtualRouters",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "VirtualRouters_ListByResourceGroup",
		Resource:    "Microsoft.Network",
	},
	"VirtualRouters_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Network/virtualRouters",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "VirtualRouters_List",
		Resource:    "Microsoft.Network",
	},
	"VirtualRouterPeerings_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/virtualRouters/{{.virtualRouterName}}/peerings/{{.peeringName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "VirtualRouterPeerings_Get",
		Resource:    "Microsoft.Network",
	},
	"VirtualRouterPeerings_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/virtualRouters/{{.virtualRouterName}}/peerings",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "VirtualRouterPeerings_List",
		Resource:    "Microsoft.Network",
	},
}
