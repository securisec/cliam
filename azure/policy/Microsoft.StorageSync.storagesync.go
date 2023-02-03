package policy

// Microsoft_StorageSync_storagesync policy
var Microsoft_StorageSync_storagesync = map[string]Policy{
	"Operations_List": {
		Path:   "/providers/Microsoft.StorageSync/operations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-06-01",
		},
		OperationID: "Operations_List",
		Resource:    "Microsoft.StorageSync",
	},
	"StorageSyncServices_CheckNameAvailability": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.StorageSync/locations/{{.locationName}}/checkNameAvailability",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-06-01",
		},
		OperationID: "StorageSyncServices_CheckNameAvailability",
		Resource:    "Microsoft.StorageSync",
	},
	"StorageSyncServices_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.StorageSync/storageSyncServices/{{.storageSyncServiceName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-06-01",
		},
		OperationID: "StorageSyncServices_Get",
		Resource:    "Microsoft.StorageSync",
	},
	"StorageSyncServices_ListByResourceGroup": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.StorageSync/storageSyncServices",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-06-01",
		},
		OperationID: "StorageSyncServices_ListByResourceGroup",
		Resource:    "Microsoft.StorageSync",
	},
	"StorageSyncServices_ListBySubscription": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.StorageSync/storageSyncServices",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-06-01",
		},
		OperationID: "StorageSyncServices_ListBySubscription",
		Resource:    "Microsoft.StorageSync",
	},
	"PrivateLinkResources_ListByStorageSyncService": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.StorageSync/storageSyncServices/{{.storageSyncServiceName}}/privateLinkResources",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-06-01",
		},
		OperationID: "PrivateLinkResources_ListByStorageSyncService",
		Resource:    "Microsoft.StorageSync",
	},
	"PrivateEndpointConnections_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.StorageSync/storageSyncServices/{{.storageSyncServiceName}}/privateEndpointConnections/{{.privateEndpointConnectionName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-06-01",
		},
		OperationID: "PrivateEndpointConnections_Get",
		Resource:    "Microsoft.StorageSync",
	},
	"PrivateEndpointConnections_ListByStorageSyncService": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.StorageSync/storageSyncServices/{{.storageSyncServiceName}}/privateEndpointConnections",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-06-01",
		},
		OperationID: "PrivateEndpointConnections_ListByStorageSyncService",
		Resource:    "Microsoft.StorageSync",
	},
	"SyncGroups_ListByStorageSyncService": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.StorageSync/storageSyncServices/{{.storageSyncServiceName}}/syncGroups",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-06-01",
		},
		OperationID: "SyncGroups_ListByStorageSyncService",
		Resource:    "Microsoft.StorageSync",
	},
	"SyncGroups_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.StorageSync/storageSyncServices/{{.storageSyncServiceName}}/syncGroups/{{.syncGroupName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-06-01",
		},
		OperationID: "SyncGroups_Get",
		Resource:    "Microsoft.StorageSync",
	},
	"CloudEndpoints_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.StorageSync/storageSyncServices/{{.storageSyncServiceName}}/syncGroups/{{.syncGroupName}}/cloudEndpoints/{{.cloudEndpointName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-06-01",
		},
		OperationID: "CloudEndpoints_Get",
		Resource:    "Microsoft.StorageSync",
	},
	"CloudEndpoints_ListBySyncGroup": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.StorageSync/storageSyncServices/{{.storageSyncServiceName}}/syncGroups/{{.syncGroupName}}/cloudEndpoints",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-06-01",
		},
		OperationID: "CloudEndpoints_ListBySyncGroup",
		Resource:    "Microsoft.StorageSync",
	},
	"CloudEndpoints_PreBackup": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.StorageSync/storageSyncServices/{{.storageSyncServiceName}}/syncGroups/{{.syncGroupName}}/cloudEndpoints/{{.cloudEndpointName}}/prebackup",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-06-01",
		},
		OperationID: "CloudEndpoints_PreBackup",
		Resource:    "Microsoft.StorageSync",
	},
	"CloudEndpoints_PostBackup": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.StorageSync/storageSyncServices/{{.storageSyncServiceName}}/syncGroups/{{.syncGroupName}}/cloudEndpoints/{{.cloudEndpointName}}/postbackup",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-06-01",
		},
		OperationID: "CloudEndpoints_PostBackup",
		Resource:    "Microsoft.StorageSync",
	},
	"CloudEndpoints_PreRestore": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.StorageSync/storageSyncServices/{{.storageSyncServiceName}}/syncGroups/{{.syncGroupName}}/cloudEndpoints/{{.cloudEndpointName}}/prerestore",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-06-01",
		},
		OperationID: "CloudEndpoints_PreRestore",
		Resource:    "Microsoft.StorageSync",
	},
	"CloudEndpoints_restoreheartbeat": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.StorageSync/storageSyncServices/{{.storageSyncServiceName}}/syncGroups/{{.syncGroupName}}/cloudEndpoints/{{.cloudEndpointName}}/restoreheartbeat",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-06-01",
		},
		OperationID: "CloudEndpoints_restoreheartbeat",
		Resource:    "Microsoft.StorageSync",
	},
	"CloudEndpoints_PostRestore": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.StorageSync/storageSyncServices/{{.storageSyncServiceName}}/syncGroups/{{.syncGroupName}}/cloudEndpoints/{{.cloudEndpointName}}/postrestore",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-06-01",
		},
		OperationID: "CloudEndpoints_PostRestore",
		Resource:    "Microsoft.StorageSync",
	},
	"CloudEndpoints_TriggerChangeDetection": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.StorageSync/storageSyncServices/{{.storageSyncServiceName}}/syncGroups/{{.syncGroupName}}/cloudEndpoints/{{.cloudEndpointName}}/triggerChangeDetection",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-06-01",
		},
		OperationID: "CloudEndpoints_TriggerChangeDetection",
		Resource:    "Microsoft.StorageSync",
	},
	"CloudEndpoints_AfsShareMetadataCertificatePublicKeys": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.StorageSync/storageSyncServices/{{.storageSyncServiceName}}/syncGroups/{{.syncGroupName}}/cloudEndpoints/{{.cloudEndpointName}}/afsShareMetadataCertificatePublicKeys",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-06-01",
		},
		OperationID: "CloudEndpoints_AfsShareMetadataCertificatePublicKeys",
		Resource:    "Microsoft.StorageSync",
	},
	"ServerEndpoints_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.StorageSync/storageSyncServices/{{.storageSyncServiceName}}/syncGroups/{{.syncGroupName}}/serverEndpoints/{{.serverEndpointName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-06-01",
		},
		OperationID: "ServerEndpoints_Get",
		Resource:    "Microsoft.StorageSync",
	},
	"ServerEndpoints_ListBySyncGroup": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.StorageSync/storageSyncServices/{{.storageSyncServiceName}}/syncGroups/{{.syncGroupName}}/serverEndpoints",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-06-01",
		},
		OperationID: "ServerEndpoints_ListBySyncGroup",
		Resource:    "Microsoft.StorageSync",
	},
	"ServerEndpoints_recallAction": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.StorageSync/storageSyncServices/{{.storageSyncServiceName}}/syncGroups/{{.syncGroupName}}/serverEndpoints/{{.serverEndpointName}}/recallAction",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-06-01",
		},
		OperationID: "ServerEndpoints_recallAction",
		Resource:    "Microsoft.StorageSync",
	},
	"RegisteredServers_ListByStorageSyncService": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.StorageSync/storageSyncServices/{{.storageSyncServiceName}}/registeredServers",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-06-01",
		},
		OperationID: "RegisteredServers_ListByStorageSyncService",
		Resource:    "Microsoft.StorageSync",
	},
	"RegisteredServers_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.StorageSync/storageSyncServices/{{.storageSyncServiceName}}/registeredServers/{{.serverId}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-06-01",
		},
		OperationID: "RegisteredServers_Get",
		Resource:    "Microsoft.StorageSync",
	},
	"RegisteredServers_triggerRollover": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.StorageSync/storageSyncServices/{{.storageSyncServiceName}}/registeredServers/{{.serverId}}/triggerRollover",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-06-01",
		},
		OperationID: "RegisteredServers_triggerRollover",
		Resource:    "Microsoft.StorageSync",
	},
	"Workflows_ListByStorageSyncService": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.StorageSync/storageSyncServices/{{.storageSyncServiceName}}/workflows",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-06-01",
		},
		OperationID: "Workflows_ListByStorageSyncService",
		Resource:    "Microsoft.StorageSync",
	},
	"Workflows_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.StorageSync/storageSyncServices/{{.storageSyncServiceName}}/workflows/{{.workflowId}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-06-01",
		},
		OperationID: "Workflows_Get",
		Resource:    "Microsoft.StorageSync",
	},
	"Workflows_Abort": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.StorageSync/storageSyncServices/{{.storageSyncServiceName}}/workflows/{{.workflowId}}/abort",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-06-01",
		},
		OperationID: "Workflows_Abort",
		Resource:    "Microsoft.StorageSync",
	},
	"OperationStatus_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.StorageSync/locations/{{.locationName}}/workflows/{{.workflowId}}/operations/{{.operationId}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-06-01",
		},
		OperationID: "OperationStatus_Get",
		Resource:    "Microsoft.StorageSync",
	},
	"LocationOperationStatus": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.StorageSync/locations/{{.locationName}}/operations/{{.operationId}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-06-01",
		},
		OperationID: "LocationOperationStatus",
		Resource:    "Microsoft.StorageSync",
	},
}
