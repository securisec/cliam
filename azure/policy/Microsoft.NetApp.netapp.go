package policy

var Microsoft_NetApp_netapp = map[string]Policy{
	"Operations_List": {
		Path:   "/providers/Microsoft.NetApp/operations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-01-01",
		},
		OperationID: "Operations_List",
		Resource:    "Microsoft.NetApp",
	},
	"NetAppResource_CheckNameAvailability": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.NetApp/locations/{{.location}}/checkNameAvailability",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-01-01",
		},
		OperationID: "NetAppResource_CheckNameAvailability",
		Resource:    "Microsoft.NetApp",
	},
	"NetAppResource_CheckFilePathAvailability": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.NetApp/locations/{{.location}}/checkFilePathAvailability",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-01-01",
		},
		OperationID: "NetAppResource_CheckFilePathAvailability",
		Resource:    "Microsoft.NetApp",
	},
	"NetAppResource_CheckQuotaAvailability": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.NetApp/locations/{{.location}}/checkQuotaAvailability",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-01-01",
		},
		OperationID: "NetAppResource_CheckQuotaAvailability",
		Resource:    "Microsoft.NetApp",
	},
	"NetAppResourceQuotaLimits_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.NetApp/locations/{{.location}}/quotaLimits",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-01-01",
		},
		OperationID: "NetAppResourceQuotaLimits_List",
		Resource:    "Microsoft.NetApp",
	},
	"NetAppResourceQuotaLimits_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.NetApp/locations/{{.location}}/quotaLimits/{{.quotaLimitName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-01-01",
		},
		OperationID: "NetAppResourceQuotaLimits_Get",
		Resource:    "Microsoft.NetApp",
	},
	"Accounts_ListBySubscription": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.NetApp/netAppAccounts",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-01-01",
		},
		OperationID: "Accounts_ListBySubscription",
		Resource:    "Microsoft.NetApp",
	},
	"Accounts_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.NetApp/netAppAccounts",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-01-01",
		},
		OperationID: "Accounts_List",
		Resource:    "Microsoft.NetApp",
	},
	"Accounts_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.NetApp/netAppAccounts/{{.accountName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-01-01",
		},
		OperationID: "Accounts_Get",
		Resource:    "Microsoft.NetApp",
	},
	"Pools_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.NetApp/netAppAccounts/{{.accountName}}/capacityPools",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-01-01",
		},
		OperationID: "Pools_List",
		Resource:    "Microsoft.NetApp",
	},
	"Pools_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.NetApp/netAppAccounts/{{.accountName}}/capacityPools/{{.poolName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-01-01",
		},
		OperationID: "Pools_Get",
		Resource:    "Microsoft.NetApp",
	},
	"Volumes_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.NetApp/netAppAccounts/{{.accountName}}/capacityPools/{{.poolName}}/volumes",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-01-01",
		},
		OperationID: "Volumes_List",
		Resource:    "Microsoft.NetApp",
	},
	"Volumes_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.NetApp/netAppAccounts/{{.accountName}}/capacityPools/{{.poolName}}/volumes/{{.volumeName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-01-01",
		},
		OperationID: "Volumes_Get",
		Resource:    "Microsoft.NetApp",
	},
	"Volumes_Revert": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.NetApp/netAppAccounts/{{.accountName}}/capacityPools/{{.poolName}}/volumes/{{.volumeName}}/revert",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-01-01",
		},
		OperationID: "Volumes_Revert",
		Resource:    "Microsoft.NetApp",
	},
	"Volumes_ResetCifsPassword": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.NetApp/netAppAccounts/{{.accountName}}/capacityPools/{{.poolName}}/volumes/{{.volumeName}}/resetCifsPassword",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-01-01",
		},
		OperationID: "Volumes_ResetCifsPassword",
		Resource:    "Microsoft.NetApp",
	},
	"Volumes_BreakReplication": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.NetApp/netAppAccounts/{{.accountName}}/capacityPools/{{.poolName}}/volumes/{{.volumeName}}/breakReplication",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-01-01",
		},
		OperationID: "Volumes_BreakReplication",
		Resource:    "Microsoft.NetApp",
	},
	"Volumes_ReplicationStatus": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.NetApp/netAppAccounts/{{.accountName}}/capacityPools/{{.poolName}}/volumes/{{.volumeName}}/replicationStatus",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-01-01",
		},
		OperationID: "Volumes_ReplicationStatus",
		Resource:    "Microsoft.NetApp",
	},
	"Volumes_ListReplications": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.NetApp/netAppAccounts/{{.accountName}}/capacityPools/{{.poolName}}/volumes/{{.volumeName}}/listReplications",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-01-01",
		},
		OperationID: "Volumes_ListReplications",
		Resource:    "Microsoft.NetApp",
	},
	"Volumes_ResyncReplication": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.NetApp/netAppAccounts/{{.accountName}}/capacityPools/{{.poolName}}/volumes/{{.volumeName}}/resyncReplication",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-01-01",
		},
		OperationID: "Volumes_ResyncReplication",
		Resource:    "Microsoft.NetApp",
	},
	"Volumes_DeleteReplication": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.NetApp/netAppAccounts/{{.accountName}}/capacityPools/{{.poolName}}/volumes/{{.volumeName}}/deleteReplication",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-01-01",
		},
		OperationID: "Volumes_DeleteReplication",
		Resource:    "Microsoft.NetApp",
	},
	"Volumes_AuthorizeReplication": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.NetApp/netAppAccounts/{{.accountName}}/capacityPools/{{.poolName}}/volumes/{{.volumeName}}/authorizeReplication",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-01-01",
		},
		OperationID: "Volumes_AuthorizeReplication",
		Resource:    "Microsoft.NetApp",
	},
	"Volumes_ReInitializeReplication": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.NetApp/netAppAccounts/{{.accountName}}/capacityPools/{{.poolName}}/volumes/{{.volumeName}}/reinitializeReplication",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-01-01",
		},
		OperationID: "Volumes_ReInitializeReplication",
		Resource:    "Microsoft.NetApp",
	},
	"Volumes_PoolChange": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.NetApp/netAppAccounts/{{.accountName}}/capacityPools/{{.poolName}}/volumes/{{.volumeName}}/poolChange",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-01-01",
		},
		OperationID: "Volumes_PoolChange",
		Resource:    "Microsoft.NetApp",
	},
	"Volumes_Relocate": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.NetApp/netAppAccounts/{{.accountName}}/capacityPools/{{.poolName}}/volumes/{{.volumeName}}/relocate",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-01-01",
		},
		OperationID: "Volumes_Relocate",
		Resource:    "Microsoft.NetApp",
	},
	"Volumes_FinalizeRelocation": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.NetApp/netAppAccounts/{{.accountName}}/capacityPools/{{.poolName}}/volumes/{{.volumeName}}/finalizeRelocation",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-01-01",
		},
		OperationID: "Volumes_FinalizeRelocation",
		Resource:    "Microsoft.NetApp",
	},
	"Volumes_RevertRelocation": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.NetApp/netAppAccounts/{{.accountName}}/capacityPools/{{.poolName}}/volumes/{{.volumeName}}/revertRelocation",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-01-01",
		},
		OperationID: "Volumes_RevertRelocation",
		Resource:    "Microsoft.NetApp",
	},
	"Snapshots_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.NetApp/netAppAccounts/{{.accountName}}/capacityPools/{{.poolName}}/volumes/{{.volumeName}}/snapshots",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-01-01",
		},
		OperationID: "Snapshots_List",
		Resource:    "Microsoft.NetApp",
	},
	"Snapshots_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.NetApp/netAppAccounts/{{.accountName}}/capacityPools/{{.poolName}}/volumes/{{.volumeName}}/snapshots/{{.snapshotName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-01-01",
		},
		OperationID: "Snapshots_Get",
		Resource:    "Microsoft.NetApp",
	},
	"Snapshots_RestoreFiles": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.NetApp/netAppAccounts/{{.accountName}}/capacityPools/{{.poolName}}/volumes/{{.volumeName}}/snapshots/{{.snapshotName}}/restoreFiles",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-01-01",
		},
		OperationID: "Snapshots_RestoreFiles",
		Resource:    "Microsoft.NetApp",
	},
	"SnapshotPolicies_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.NetApp/netAppAccounts/{{.accountName}}/snapshotPolicies",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-01-01",
		},
		OperationID: "SnapshotPolicies_List",
		Resource:    "Microsoft.NetApp",
	},
	"SnapshotPolicies_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.NetApp/netAppAccounts/{{.accountName}}/snapshotPolicies/{{.snapshotPolicyName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-01-01",
		},
		OperationID: "SnapshotPolicies_Get",
		Resource:    "Microsoft.NetApp",
	},
	"SnapshotPolicies_ListVolumes": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.NetApp/netAppAccounts/{{.accountName}}/snapshotPolicies/{{.snapshotPolicyName}}/volumes",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-01-01",
		},
		OperationID: "SnapshotPolicies_ListVolumes",
		Resource:    "Microsoft.NetApp",
	},
	"Backups_GetStatus": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.NetApp/netAppAccounts/{{.accountName}}/capacityPools/{{.poolName}}/volumes/{{.volumeName}}/backupStatus",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-01-01",
		},
		OperationID: "Backups_GetStatus",
		Resource:    "Microsoft.NetApp",
	},
	"Backups_GetVolumeRestoreStatus": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.NetApp/netAppAccounts/{{.accountName}}/capacityPools/{{.poolName}}/volumes/{{.volumeName}}/restoreStatus",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-01-01",
		},
		OperationID: "Backups_GetVolumeRestoreStatus",
		Resource:    "Microsoft.NetApp",
	},
	"AccountBackups_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.NetApp/netAppAccounts/{{.accountName}}/accountBackups",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-01-01",
		},
		OperationID: "AccountBackups_List",
		Resource:    "Microsoft.NetApp",
	},
	"AccountBackups_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.NetApp/netAppAccounts/{{.accountName}}/accountBackups/{{.backupName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-01-01",
		},
		OperationID: "AccountBackups_Get",
		Resource:    "Microsoft.NetApp",
	},
	"Backups_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.NetApp/netAppAccounts/{{.accountName}}/capacityPools/{{.poolName}}/volumes/{{.volumeName}}/backups",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-01-01",
		},
		OperationID: "Backups_List",
		Resource:    "Microsoft.NetApp",
	},
	"Backups_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.NetApp/netAppAccounts/{{.accountName}}/capacityPools/{{.poolName}}/volumes/{{.volumeName}}/backups/{{.backupName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-01-01",
		},
		OperationID: "Backups_Get",
		Resource:    "Microsoft.NetApp",
	},
	"BackupPolicies_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.NetApp/netAppAccounts/{{.accountName}}/backupPolicies",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-01-01",
		},
		OperationID: "BackupPolicies_List",
		Resource:    "Microsoft.NetApp",
	},
	"BackupPolicies_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.NetApp/netAppAccounts/{{.accountName}}/backupPolicies/{{.backupPolicyName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-01-01",
		},
		OperationID: "BackupPolicies_Get",
		Resource:    "Microsoft.NetApp",
	},
	"VolumeQuotaRules_ListByVolume": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.NetApp/netAppAccounts/{{.accountName}}/capacityPools/{{.poolName}}/volumes/{{.volumeName}}/volumeQuotaRules",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-01-01",
		},
		OperationID: "VolumeQuotaRules_ListByVolume",
		Resource:    "Microsoft.NetApp",
	},
	"VolumeQuotaRules_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.NetApp/netAppAccounts/{{.accountName}}/capacityPools/{{.poolName}}/volumes/{{.volumeName}}/volumeQuotaRules/{{.volumeQuotaRuleName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-01-01",
		},
		OperationID: "VolumeQuotaRules_Get",
		Resource:    "Microsoft.NetApp",
	},
	"Vaults_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.NetApp/netAppAccounts/{{.accountName}}/vaults",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-01-01",
		},
		OperationID: "Vaults_List",
		Resource:    "Microsoft.NetApp",
	},
	"VolumeGroups_ListByNetAppAccount": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.NetApp/netAppAccounts/{{.accountName}}/volumeGroups",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-01-01",
		},
		OperationID: "VolumeGroups_ListByNetAppAccount",
		Resource:    "Microsoft.NetApp",
	},
	"VolumeGroups_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.NetApp/netAppAccounts/{{.accountName}}/volumeGroups/{{.volumeGroupName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-01-01",
		},
		OperationID: "VolumeGroups_Get",
		Resource:    "Microsoft.NetApp",
	},
	"Subvolumes_ListByVolume": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.NetApp/netAppAccounts/{{.accountName}}/capacityPools/{{.poolName}}/volumes/{{.volumeName}}/subvolumes",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-01-01",
		},
		OperationID: "Subvolumes_ListByVolume",
		Resource:    "Microsoft.NetApp",
	},
	"Subvolumes_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.NetApp/netAppAccounts/{{.accountName}}/capacityPools/{{.poolName}}/volumes/{{.volumeName}}/subvolumes/{{.subvolumeName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-01-01",
		},
		OperationID: "Subvolumes_Get",
		Resource:    "Microsoft.NetApp",
	},
	"Subvolumes_GetMetadata": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.NetApp/netAppAccounts/{{.accountName}}/capacityPools/{{.poolName}}/volumes/{{.volumeName}}/subvolumes/{{.subvolumeName}}/getMetadata",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-01-01",
		},
		OperationID: "Subvolumes_GetMetadata",
		Resource:    "Microsoft.NetApp",
	},
}
