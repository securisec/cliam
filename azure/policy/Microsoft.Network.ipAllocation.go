package policy

// Microsoft_Network_ipAllocation policy
var Microsoft_Network_ipAllocation = map[string]Policy{
	"IpAllocations_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/IpAllocations/{{.ipAllocationName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "IpAllocations_Get",
		Resource:    "Microsoft.Network",
	},
	"IpAllocations_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Network/IpAllocations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "IpAllocations_List",
		Resource:    "Microsoft.Network",
	},
	"IpAllocations_ListByResourceGroup": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/IpAllocations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "IpAllocations_ListByResourceGroup",
		Resource:    "Microsoft.Network",
	},
}
