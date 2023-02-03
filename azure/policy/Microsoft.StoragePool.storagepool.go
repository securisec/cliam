package policy

// Microsoft_StoragePool_storagepool policy
var Microsoft_StoragePool_storagepool = map[string]Policy{
	"Operations_List": {
		Path:   "/providers/Microsoft.StoragePool/operations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-08-01",
		},
		OperationID: "Operations_List",
		Resource:    "Microsoft.StoragePool",
	},
	"DiskPools_ListBySubscription": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.StoragePool/diskPools",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-08-01",
		},
		OperationID: "DiskPools_ListBySubscription",
		Resource:    "Microsoft.StoragePool",
	},
	"DiskPools_ListByResourceGroup": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.StoragePool/diskPools",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-08-01",
		},
		OperationID: "DiskPools_ListByResourceGroup",
		Resource:    "Microsoft.StoragePool",
	},
	"DiskPools_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.StoragePool/diskPools/{{.diskPoolName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-08-01",
		},
		OperationID: "DiskPools_Get",
		Resource:    "Microsoft.StoragePool",
	},
	"DiskPoolZones_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.StoragePool/locations/{{.location}}/diskPoolZones",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-08-01",
		},
		OperationID: "DiskPoolZones_List",
		Resource:    "Microsoft.StoragePool",
	},
	"ResourceSkus_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.StoragePool/locations/{{.location}}/skus",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-08-01",
		},
		OperationID: "ResourceSkus_List",
		Resource:    "Microsoft.StoragePool",
	},
	"DiskPools_ListOutboundNetworkDependenciesEndpoints": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.StoragePool/diskPools/{{.diskPoolName}}/outboundNetworkDependenciesEndpoints",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-08-01",
		},
		OperationID: "DiskPools_ListOutboundNetworkDependenciesEndpoints",
		Resource:    "Microsoft.StoragePool",
	},
	"DiskPools_Start": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.StoragePool/diskPools/{{.diskPoolName}}/start",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-08-01",
		},
		OperationID: "DiskPools_Start",
		Resource:    "Microsoft.StoragePool",
	},
	"DiskPools_Deallocate": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.StoragePool/diskPools/{{.diskPoolName}}/deallocate",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-08-01",
		},
		OperationID: "DiskPools_Deallocate",
		Resource:    "Microsoft.StoragePool",
	},
	"DiskPools_Upgrade": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.StoragePool/diskPools/{{.diskPoolName}}/upgrade",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-08-01",
		},
		OperationID: "DiskPools_Upgrade",
		Resource:    "Microsoft.StoragePool",
	},
	"IscsiTargets_ListByDiskPool": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.StoragePool/diskPools/{{.diskPoolName}}/iscsiTargets",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-08-01",
		},
		OperationID: "IscsiTargets_ListByDiskPool",
		Resource:    "Microsoft.StoragePool",
	},
	"IscsiTargets_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.StoragePool/diskPools/{{.diskPoolName}}/iscsiTargets/{{.iscsiTargetName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-08-01",
		},
		OperationID: "IscsiTargets_Get",
		Resource:    "Microsoft.StoragePool",
	},
}
