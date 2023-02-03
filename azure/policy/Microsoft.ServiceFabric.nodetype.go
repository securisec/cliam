package policy

// Microsoft_ServiceFabric_nodetype policy
var Microsoft_ServiceFabric_nodetype = map[string]Policy{
	"NodeTypes_ListByManagedClusters": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourcegroups/{{.resourceGroupName}}/providers/Microsoft.ServiceFabric/managedClusters/{{.clusterName}}/nodeTypes",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-01-01",
		},
		OperationID: "NodeTypes_ListByManagedClusters",
		Resource:    "Microsoft.ServiceFabric",
	},
	"NodeTypes_Restart": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ServiceFabric/managedClusters/{{.clusterName}}/nodeTypes/{{.nodeTypeName}}/restart",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-01-01",
		},
		OperationID: "NodeTypes_Restart",
		Resource:    "Microsoft.ServiceFabric",
	},
	"NodeTypes_Reimage": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ServiceFabric/managedClusters/{{.clusterName}}/nodeTypes/{{.nodeTypeName}}/reimage",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-01-01",
		},
		OperationID: "NodeTypes_Reimage",
		Resource:    "Microsoft.ServiceFabric",
	},
	"NodeTypes_DeleteNode": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ServiceFabric/managedClusters/{{.clusterName}}/nodeTypes/{{.nodeTypeName}}/deleteNode",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-01-01",
		},
		OperationID: "NodeTypes_DeleteNode",
		Resource:    "Microsoft.ServiceFabric",
	},
	"NodeTypeSkus_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ServiceFabric/managedClusters/{{.clusterName}}/nodeTypes/{{.nodeTypeName}}/skus",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-01-01",
		},
		OperationID: "NodeTypeSkus_List",
		Resource:    "Microsoft.ServiceFabric",
	},
	"NodeTypes_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ServiceFabric/managedClusters/{{.clusterName}}/nodeTypes/{{.nodeTypeName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-01-01",
		},
		OperationID: "NodeTypes_Get",
		Resource:    "Microsoft.ServiceFabric",
	},
}
