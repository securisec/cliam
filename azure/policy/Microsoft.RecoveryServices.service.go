package policy

// Microsoft_RecoveryServices_service policy
var Microsoft_RecoveryServices_service = map[string]Policy{
	"Operations_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.RecoveryServices/operations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "Operations_List",
		Resource:    "Microsoft.RecoveryServices",
	},
	"ReplicationAlertSettings_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.RecoveryServices/vaults/{{.resourceName}}/replicationAlertSettings",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "ReplicationAlertSettings_List",
		Resource:    "Microsoft.RecoveryServices",
	},
	"ReplicationAlertSettings_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.RecoveryServices/vaults/{{.resourceName}}/replicationAlertSettings/{{.alertSettingName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "ReplicationAlertSettings_Get",
		Resource:    "Microsoft.RecoveryServices",
	},
	"ReplicationAppliances_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.RecoveryServices/vaults/{{.resourceName}}/replicationAppliances",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "ReplicationAppliances_List",
		Resource:    "Microsoft.RecoveryServices",
	},
	"ReplicationEligibilityResults_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Compute/virtualMachines/{{.virtualMachineName}}/providers/Microsoft.RecoveryServices/replicationEligibilityResults",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "ReplicationEligibilityResults_List",
		Resource:    "Microsoft.RecoveryServices",
	},
	"ReplicationEligibilityResults_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Compute/virtualMachines/{{.virtualMachineName}}/providers/Microsoft.RecoveryServices/replicationEligibilityResults/default",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "ReplicationEligibilityResults_Get",
		Resource:    "Microsoft.RecoveryServices",
	},
	"ReplicationEvents_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.RecoveryServices/vaults/{{.resourceName}}/replicationEvents",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "ReplicationEvents_List",
		Resource:    "Microsoft.RecoveryServices",
	},
	"ReplicationEvents_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.RecoveryServices/vaults/{{.resourceName}}/replicationEvents/{{.eventName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "ReplicationEvents_Get",
		Resource:    "Microsoft.RecoveryServices",
	},
	"ReplicationFabrics_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.RecoveryServices/vaults/{{.resourceName}}/replicationFabrics",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "ReplicationFabrics_List",
		Resource:    "Microsoft.RecoveryServices",
	},
	"ReplicationFabrics_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.RecoveryServices/vaults/{{.resourceName}}/replicationFabrics/{{.fabricName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "ReplicationFabrics_Get",
		Resource:    "Microsoft.RecoveryServices",
	},
	"ReplicationFabrics_CheckConsistency": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.RecoveryServices/vaults/{{.resourceName}}/replicationFabrics/{{.fabricName}}/checkConsistency",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "ReplicationFabrics_CheckConsistency",
		Resource:    "Microsoft.RecoveryServices",
	},
	"ReplicationFabrics_MigrateToAad": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.RecoveryServices/vaults/{{.resourceName}}/replicationFabrics/{{.fabricName}}/migratetoaad",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "ReplicationFabrics_MigrateToAad",
		Resource:    "Microsoft.RecoveryServices",
	},
	"ReplicationFabrics_ReassociateGateway": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.RecoveryServices/vaults/{{.resourceName}}/replicationFabrics/{{.fabricName}}/reassociateGateway",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "ReplicationFabrics_ReassociateGateway",
		Resource:    "Microsoft.RecoveryServices",
	},
	"ReplicationFabrics_Delete": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.RecoveryServices/vaults/{{.resourceName}}/replicationFabrics/{{.fabricName}}/remove",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "ReplicationFabrics_Delete",
		Resource:    "Microsoft.RecoveryServices",
	},
	"ReplicationFabrics_RenewCertificate": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.RecoveryServices/vaults/{{.resourceName}}/replicationFabrics/{{.fabricName}}/renewCertificate",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "ReplicationFabrics_RenewCertificate",
		Resource:    "Microsoft.RecoveryServices",
	},
	"ReplicationLogicalNetworks_ListByReplicationFabrics": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.RecoveryServices/vaults/{{.resourceName}}/replicationFabrics/{{.fabricName}}/replicationLogicalNetworks",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "ReplicationLogicalNetworks_ListByReplicationFabrics",
		Resource:    "Microsoft.RecoveryServices",
	},
	"ReplicationLogicalNetworks_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.RecoveryServices/vaults/{{.resourceName}}/replicationFabrics/{{.fabricName}}/replicationLogicalNetworks/{{.logicalNetworkName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "ReplicationLogicalNetworks_Get",
		Resource:    "Microsoft.RecoveryServices",
	},
	"ReplicationNetworks_ListByReplicationFabrics": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.RecoveryServices/vaults/{{.resourceName}}/replicationFabrics/{{.fabricName}}/replicationNetworks",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "ReplicationNetworks_ListByReplicationFabrics",
		Resource:    "Microsoft.RecoveryServices",
	},
	"ReplicationNetworks_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.RecoveryServices/vaults/{{.resourceName}}/replicationFabrics/{{.fabricName}}/replicationNetworks/{{.networkName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "ReplicationNetworks_Get",
		Resource:    "Microsoft.RecoveryServices",
	},
	"ReplicationNetworkMappings_ListByReplicationNetworks": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.RecoveryServices/vaults/{{.resourceName}}/replicationFabrics/{{.fabricName}}/replicationNetworks/{{.networkName}}/replicationNetworkMappings",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "ReplicationNetworkMappings_ListByReplicationNetworks",
		Resource:    "Microsoft.RecoveryServices",
	},
	"ReplicationNetworkMappings_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.RecoveryServices/vaults/{{.resourceName}}/replicationFabrics/{{.fabricName}}/replicationNetworks/{{.networkName}}/replicationNetworkMappings/{{.networkMappingName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "ReplicationNetworkMappings_Get",
		Resource:    "Microsoft.RecoveryServices",
	},
	"ReplicationProtectionContainers_ListByReplicationFabrics": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.RecoveryServices/vaults/{{.resourceName}}/replicationFabrics/{{.fabricName}}/replicationProtectionContainers",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "ReplicationProtectionContainers_ListByReplicationFabrics",
		Resource:    "Microsoft.RecoveryServices",
	},
	"ReplicationProtectionContainers_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.RecoveryServices/vaults/{{.resourceName}}/replicationFabrics/{{.fabricName}}/replicationProtectionContainers/{{.protectionContainerName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "ReplicationProtectionContainers_Get",
		Resource:    "Microsoft.RecoveryServices",
	},
	"ReplicationProtectionContainers_DiscoverProtectableItem": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.RecoveryServices/vaults/{{.resourceName}}/replicationFabrics/{{.fabricName}}/replicationProtectionContainers/{{.protectionContainerName}}/discoverProtectableItem",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "ReplicationProtectionContainers_DiscoverProtectableItem",
		Resource:    "Microsoft.RecoveryServices",
	},
	"ReplicationProtectionContainers_Delete": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.RecoveryServices/vaults/{{.resourceName}}/replicationFabrics/{{.fabricName}}/replicationProtectionContainers/{{.protectionContainerName}}/remove",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "ReplicationProtectionContainers_Delete",
		Resource:    "Microsoft.RecoveryServices",
	},
	"ReplicationMigrationItems_ListByReplicationProtectionContainers": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.RecoveryServices/vaults/{{.resourceName}}/replicationFabrics/{{.fabricName}}/replicationProtectionContainers/{{.protectionContainerName}}/replicationMigrationItems",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "ReplicationMigrationItems_ListByReplicationProtectionContainers",
		Resource:    "Microsoft.RecoveryServices",
	},
	"ReplicationMigrationItems_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.RecoveryServices/vaults/{{.resourceName}}/replicationFabrics/{{.fabricName}}/replicationProtectionContainers/{{.protectionContainerName}}/replicationMigrationItems/{{.migrationItemName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "ReplicationMigrationItems_Get",
		Resource:    "Microsoft.RecoveryServices",
	},
	"ReplicationMigrationItems_Migrate": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.RecoveryServices/vaults/{{.resourceName}}/replicationFabrics/{{.fabricName}}/replicationProtectionContainers/{{.protectionContainerName}}/replicationMigrationItems/{{.migrationItemName}}/migrate",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "ReplicationMigrationItems_Migrate",
		Resource:    "Microsoft.RecoveryServices",
	},
	"MigrationRecoveryPoints_ListByReplicationMigrationItems": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.RecoveryServices/vaults/{{.resourceName}}/replicationFabrics/{{.fabricName}}/replicationProtectionContainers/{{.protectionContainerName}}/replicationMigrationItems/{{.migrationItemName}}/migrationRecoveryPoints",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "MigrationRecoveryPoints_ListByReplicationMigrationItems",
		Resource:    "Microsoft.RecoveryServices",
	},
	"MigrationRecoveryPoints_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.RecoveryServices/vaults/{{.resourceName}}/replicationFabrics/{{.fabricName}}/replicationProtectionContainers/{{.protectionContainerName}}/replicationMigrationItems/{{.migrationItemName}}/migrationRecoveryPoints/{{.migrationRecoveryPointName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "MigrationRecoveryPoints_Get",
		Resource:    "Microsoft.RecoveryServices",
	},
	"ReplicationMigrationItems_PauseReplication": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.RecoveryServices/vaults/{{.resourceName}}/replicationFabrics/{{.fabricName}}/replicationProtectionContainers/{{.protectionContainerName}}/replicationMigrationItems/{{.migrationItemName}}/pauseReplication",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "ReplicationMigrationItems_PauseReplication",
		Resource:    "Microsoft.RecoveryServices",
	},
	"ReplicationMigrationItems_ResumeReplication": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.RecoveryServices/vaults/{{.resourceName}}/replicationFabrics/{{.fabricName}}/replicationProtectionContainers/{{.protectionContainerName}}/replicationMigrationItems/{{.migrationItemName}}/resumeReplication",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "ReplicationMigrationItems_ResumeReplication",
		Resource:    "Microsoft.RecoveryServices",
	},
	"ReplicationMigrationItems_Resync": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.RecoveryServices/vaults/{{.resourceName}}/replicationFabrics/{{.fabricName}}/replicationProtectionContainers/{{.protectionContainerName}}/replicationMigrationItems/{{.migrationItemName}}/resync",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "ReplicationMigrationItems_Resync",
		Resource:    "Microsoft.RecoveryServices",
	},
	"ReplicationMigrationItems_TestMigrate": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.RecoveryServices/vaults/{{.resourceName}}/replicationFabrics/{{.fabricName}}/replicationProtectionContainers/{{.protectionContainerName}}/replicationMigrationItems/{{.migrationItemName}}/testMigrate",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "ReplicationMigrationItems_TestMigrate",
		Resource:    "Microsoft.RecoveryServices",
	},
	"ReplicationMigrationItems_TestMigrateCleanup": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.RecoveryServices/vaults/{{.resourceName}}/replicationFabrics/{{.fabricName}}/replicationProtectionContainers/{{.protectionContainerName}}/replicationMigrationItems/{{.migrationItemName}}/testMigrateCleanup",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "ReplicationMigrationItems_TestMigrateCleanup",
		Resource:    "Microsoft.RecoveryServices",
	},
	"ReplicationProtectableItems_ListByReplicationProtectionContainers": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.RecoveryServices/vaults/{{.resourceName}}/replicationFabrics/{{.fabricName}}/replicationProtectionContainers/{{.protectionContainerName}}/replicationProtectableItems",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "ReplicationProtectableItems_ListByReplicationProtectionContainers",
		Resource:    "Microsoft.RecoveryServices",
	},
	"ReplicationProtectableItems_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.RecoveryServices/vaults/{{.resourceName}}/replicationFabrics/{{.fabricName}}/replicationProtectionContainers/{{.protectionContainerName}}/replicationProtectableItems/{{.protectableItemName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "ReplicationProtectableItems_Get",
		Resource:    "Microsoft.RecoveryServices",
	},
	"ReplicationProtectedItems_ListByReplicationProtectionContainers": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.RecoveryServices/vaults/{{.resourceName}}/replicationFabrics/{{.fabricName}}/replicationProtectionContainers/{{.protectionContainerName}}/replicationProtectedItems",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "ReplicationProtectedItems_ListByReplicationProtectionContainers",
		Resource:    "Microsoft.RecoveryServices",
	},
	"ReplicationProtectedItems_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.RecoveryServices/vaults/{{.resourceName}}/replicationFabrics/{{.fabricName}}/replicationProtectionContainers/{{.protectionContainerName}}/replicationProtectedItems/{{.replicatedProtectedItemName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "ReplicationProtectedItems_Get",
		Resource:    "Microsoft.RecoveryServices",
	},
	"ReplicationProtectedItems_AddDisks": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.RecoveryServices/vaults/{{.resourceName}}/replicationFabrics/{{.fabricName}}/replicationProtectionContainers/{{.protectionContainerName}}/replicationProtectedItems/{{.replicatedProtectedItemName}}/addDisks",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "ReplicationProtectedItems_AddDisks",
		Resource:    "Microsoft.RecoveryServices",
	},
	"ReplicationProtectedItems_ApplyRecoveryPoint": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.RecoveryServices/vaults/{{.resourceName}}/replicationFabrics/{{.fabricName}}/replicationProtectionContainers/{{.protectionContainerName}}/replicationProtectedItems/{{.replicatedProtectedItemName}}/applyRecoveryPoint",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "ReplicationProtectedItems_ApplyRecoveryPoint",
		Resource:    "Microsoft.RecoveryServices",
	},
	"ReplicationProtectedItems_FailoverCancel": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.RecoveryServices/vaults/{{.resourceName}}/replicationFabrics/{{.fabricName}}/replicationProtectionContainers/{{.protectionContainerName}}/replicationProtectedItems/{{.replicatedProtectedItemName}}/failoverCancel",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "ReplicationProtectedItems_FailoverCancel",
		Resource:    "Microsoft.RecoveryServices",
	},
	"ReplicationProtectedItems_FailoverCommit": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.RecoveryServices/vaults/{{.resourceName}}/replicationFabrics/{{.fabricName}}/replicationProtectionContainers/{{.protectionContainerName}}/replicationProtectedItems/{{.replicatedProtectedItemName}}/failoverCommit",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "ReplicationProtectedItems_FailoverCommit",
		Resource:    "Microsoft.RecoveryServices",
	},
	"ReplicationProtectedItems_PlannedFailover": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.RecoveryServices/vaults/{{.resourceName}}/replicationFabrics/{{.fabricName}}/replicationProtectionContainers/{{.protectionContainerName}}/replicationProtectedItems/{{.replicatedProtectedItemName}}/plannedFailover",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "ReplicationProtectedItems_PlannedFailover",
		Resource:    "Microsoft.RecoveryServices",
	},
	"RecoveryPoints_ListByReplicationProtectedItems": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.RecoveryServices/vaults/{{.resourceName}}/replicationFabrics/{{.fabricName}}/replicationProtectionContainers/{{.protectionContainerName}}/replicationProtectedItems/{{.replicatedProtectedItemName}}/recoveryPoints",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "RecoveryPoints_ListByReplicationProtectedItems",
		Resource:    "Microsoft.RecoveryServices",
	},
	"RecoveryPoints_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.RecoveryServices/vaults/{{.resourceName}}/replicationFabrics/{{.fabricName}}/replicationProtectionContainers/{{.protectionContainerName}}/replicationProtectedItems/{{.replicatedProtectedItemName}}/recoveryPoints/{{.recoveryPointName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "RecoveryPoints_Get",
		Resource:    "Microsoft.RecoveryServices",
	},
	"ReplicationProtectedItems_Delete": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.RecoveryServices/vaults/{{.resourceName}}/replicationFabrics/{{.fabricName}}/replicationProtectionContainers/{{.protectionContainerName}}/replicationProtectedItems/{{.replicatedProtectedItemName}}/remove",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "ReplicationProtectedItems_Delete",
		Resource:    "Microsoft.RecoveryServices",
	},
	"ReplicationProtectedItems_RemoveDisks": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.RecoveryServices/vaults/{{.resourceName}}/replicationFabrics/{{.fabricName}}/replicationProtectionContainers/{{.protectionContainerName}}/replicationProtectedItems/{{.replicatedProtectedItemName}}/removeDisks",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "ReplicationProtectedItems_RemoveDisks",
		Resource:    "Microsoft.RecoveryServices",
	},
	"ReplicationProtectedItems_RepairReplication": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.RecoveryServices/vaults/{{.resourceName}}/replicationFabrics/{{.fabricName}}/replicationProtectionContainers/{{.protectionContainerName}}/replicationProtectedItems/{{.replicatedProtectedItemName}}/repairReplication",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "ReplicationProtectedItems_RepairReplication",
		Resource:    "Microsoft.RecoveryServices",
	},
	"ReplicationProtectedItems_Reprotect": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.RecoveryServices/vaults/{{.resourceName}}/replicationFabrics/{{.fabricName}}/replicationProtectionContainers/{{.protectionContainerName}}/replicationProtectedItems/{{.replicatedProtectedItemName}}/reProtect",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "ReplicationProtectedItems_Reprotect",
		Resource:    "Microsoft.RecoveryServices",
	},
	"ReplicationProtectedItems_ResolveHealthErrors": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.RecoveryServices/vaults/{{.resourceName}}/replicationFabrics/{{.fabricName}}/replicationProtectionContainers/{{.protectionContainerName}}/replicationProtectedItems/{{.replicatedProtectedItemName}}/resolveHealthErrors",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "ReplicationProtectedItems_ResolveHealthErrors",
		Resource:    "Microsoft.RecoveryServices",
	},
	"ReplicationProtectedItems_SwitchProvider": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.RecoveryServices/vaults/{{.resourceName}}/replicationFabrics/{{.fabricName}}/replicationProtectionContainers/{{.protectionContainerName}}/replicationProtectedItems/{{.replicatedProtectedItemName}}/switchProvider",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "ReplicationProtectedItems_SwitchProvider",
		Resource:    "Microsoft.RecoveryServices",
	},
	"TargetComputeSizes_ListByReplicationProtectedItems": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.RecoveryServices/vaults/{{.resourceName}}/replicationFabrics/{{.fabricName}}/replicationProtectionContainers/{{.protectionContainerName}}/replicationProtectedItems/{{.replicatedProtectedItemName}}/targetComputeSizes",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "TargetComputeSizes_ListByReplicationProtectedItems",
		Resource:    "Microsoft.RecoveryServices",
	},
	"ReplicationProtectedItems_TestFailover": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.RecoveryServices/vaults/{{.resourceName}}/replicationFabrics/{{.fabricName}}/replicationProtectionContainers/{{.protectionContainerName}}/replicationProtectedItems/{{.replicatedProtectedItemName}}/testFailover",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "ReplicationProtectedItems_TestFailover",
		Resource:    "Microsoft.RecoveryServices",
	},
	"ReplicationProtectedItems_TestFailoverCleanup": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.RecoveryServices/vaults/{{.resourceName}}/replicationFabrics/{{.fabricName}}/replicationProtectionContainers/{{.protectionContainerName}}/replicationProtectedItems/{{.replicatedProtectedItemName}}/testFailoverCleanup",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "ReplicationProtectedItems_TestFailoverCleanup",
		Resource:    "Microsoft.RecoveryServices",
	},
	"ReplicationProtectedItems_UnplannedFailover": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.RecoveryServices/vaults/{{.resourceName}}/replicationFabrics/{{.fabricName}}/replicationProtectionContainers/{{.protectionContainerName}}/replicationProtectedItems/{{.replicatedProtectedItemName}}/unplannedFailover",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "ReplicationProtectedItems_UnplannedFailover",
		Resource:    "Microsoft.RecoveryServices",
	},
	"ReplicationProtectedItems_UpdateAppliance": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.RecoveryServices/vaults/{{.resourceName}}/replicationFabrics/{{.fabricName}}/replicationProtectionContainers/{{.protectionContainerName}}/replicationProtectedItems/{{.replicatedProtectedItemName}}/updateAppliance",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "ReplicationProtectedItems_UpdateAppliance",
		Resource:    "Microsoft.RecoveryServices",
	},
	"ReplicationProtectedItems_UpdateMobilityService": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.RecoveryServices/vaults/{{.resourceName}}/replicationFabrics/{{.fabricName}}/replicationProtectionContainers/{{.protectionContainerName}}/replicationProtectedItems/{{.replicatedProtectedItemName}}/updateMobilityService",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "ReplicationProtectedItems_UpdateMobilityService",
		Resource:    "Microsoft.RecoveryServices",
	},
	"ReplicationProtectionContainerMappings_ListByReplicationProtectionContainers": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.RecoveryServices/vaults/{{.resourceName}}/replicationFabrics/{{.fabricName}}/replicationProtectionContainers/{{.protectionContainerName}}/replicationProtectionContainerMappings",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "ReplicationProtectionContainerMappings_ListByReplicationProtectionContainers",
		Resource:    "Microsoft.RecoveryServices",
	},
	"ReplicationProtectionContainerMappings_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.RecoveryServices/vaults/{{.resourceName}}/replicationFabrics/{{.fabricName}}/replicationProtectionContainers/{{.protectionContainerName}}/replicationProtectionContainerMappings/{{.mappingName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "ReplicationProtectionContainerMappings_Get",
		Resource:    "Microsoft.RecoveryServices",
	},
	"ReplicationProtectionContainerMappings_Delete": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.RecoveryServices/vaults/{{.resourceName}}/replicationFabrics/{{.fabricName}}/replicationProtectionContainers/{{.protectionContainerName}}/replicationProtectionContainerMappings/{{.mappingName}}/remove",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "ReplicationProtectionContainerMappings_Delete",
		Resource:    "Microsoft.RecoveryServices",
	},
	"ReplicationProtectionContainers_SwitchProtection": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.RecoveryServices/vaults/{{.resourceName}}/replicationFabrics/{{.fabricName}}/replicationProtectionContainers/{{.protectionContainerName}}/switchprotection",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "ReplicationProtectionContainers_SwitchProtection",
		Resource:    "Microsoft.RecoveryServices",
	},
	"ReplicationRecoveryServicesProviders_ListByReplicationFabrics": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.RecoveryServices/vaults/{{.resourceName}}/replicationFabrics/{{.fabricName}}/replicationRecoveryServicesProviders",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "ReplicationRecoveryServicesProviders_ListByReplicationFabrics",
		Resource:    "Microsoft.RecoveryServices",
	},
	"ReplicationRecoveryServicesProviders_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.RecoveryServices/vaults/{{.resourceName}}/replicationFabrics/{{.fabricName}}/replicationRecoveryServicesProviders/{{.providerName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "ReplicationRecoveryServicesProviders_Get",
		Resource:    "Microsoft.RecoveryServices",
	},
	"ReplicationRecoveryServicesProviders_RefreshProvider": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.RecoveryServices/vaults/{{.resourceName}}/replicationFabrics/{{.fabricName}}/replicationRecoveryServicesProviders/{{.providerName}}/refreshProvider",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "ReplicationRecoveryServicesProviders_RefreshProvider",
		Resource:    "Microsoft.RecoveryServices",
	},
	"ReplicationRecoveryServicesProviders_Delete": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.RecoveryServices/vaults/{{.resourceName}}/replicationFabrics/{{.fabricName}}/replicationRecoveryServicesProviders/{{.providerName}}/remove",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "ReplicationRecoveryServicesProviders_Delete",
		Resource:    "Microsoft.RecoveryServices",
	},
	"ReplicationStorageClassifications_ListByReplicationFabrics": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.RecoveryServices/vaults/{{.resourceName}}/replicationFabrics/{{.fabricName}}/replicationStorageClassifications",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "ReplicationStorageClassifications_ListByReplicationFabrics",
		Resource:    "Microsoft.RecoveryServices",
	},
	"ReplicationStorageClassifications_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.RecoveryServices/vaults/{{.resourceName}}/replicationFabrics/{{.fabricName}}/replicationStorageClassifications/{{.storageClassificationName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "ReplicationStorageClassifications_Get",
		Resource:    "Microsoft.RecoveryServices",
	},
	"ReplicationStorageClassificationMappings_ListByReplicationStorageClassifications": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.RecoveryServices/vaults/{{.resourceName}}/replicationFabrics/{{.fabricName}}/replicationStorageClassifications/{{.storageClassificationName}}/replicationStorageClassificationMappings",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "ReplicationStorageClassificationMappings_ListByReplicationStorageClassifications",
		Resource:    "Microsoft.RecoveryServices",
	},
	"ReplicationStorageClassificationMappings_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.RecoveryServices/vaults/{{.resourceName}}/replicationFabrics/{{.fabricName}}/replicationStorageClassifications/{{.storageClassificationName}}/replicationStorageClassificationMappings/{{.storageClassificationMappingName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "ReplicationStorageClassificationMappings_Get",
		Resource:    "Microsoft.RecoveryServices",
	},
	"ReplicationvCenters_ListByReplicationFabrics": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.RecoveryServices/vaults/{{.resourceName}}/replicationFabrics/{{.fabricName}}/replicationvCenters",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "ReplicationvCenters_ListByReplicationFabrics",
		Resource:    "Microsoft.RecoveryServices",
	},
	"ReplicationvCenters_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.RecoveryServices/vaults/{{.resourceName}}/replicationFabrics/{{.fabricName}}/replicationvCenters/{{.vcenterName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "ReplicationvCenters_Get",
		Resource:    "Microsoft.RecoveryServices",
	},
	"ReplicationJobs_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.RecoveryServices/vaults/{{.resourceName}}/replicationJobs",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "ReplicationJobs_List",
		Resource:    "Microsoft.RecoveryServices",
	},
	"ReplicationJobs_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.RecoveryServices/vaults/{{.resourceName}}/replicationJobs/{{.jobName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "ReplicationJobs_Get",
		Resource:    "Microsoft.RecoveryServices",
	},
	"ReplicationJobs_Cancel": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.RecoveryServices/vaults/{{.resourceName}}/replicationJobs/{{.jobName}}/cancel",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "ReplicationJobs_Cancel",
		Resource:    "Microsoft.RecoveryServices",
	},
	"ReplicationJobs_Restart": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.RecoveryServices/vaults/{{.resourceName}}/replicationJobs/{{.jobName}}/restart",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "ReplicationJobs_Restart",
		Resource:    "Microsoft.RecoveryServices",
	},
	"ReplicationJobs_Resume": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.RecoveryServices/vaults/{{.resourceName}}/replicationJobs/{{.jobName}}/resume",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "ReplicationJobs_Resume",
		Resource:    "Microsoft.RecoveryServices",
	},
	"ReplicationJobs_Export": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.RecoveryServices/vaults/{{.resourceName}}/replicationJobs/export",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "ReplicationJobs_Export",
		Resource:    "Microsoft.RecoveryServices",
	},
	"ReplicationMigrationItems_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.RecoveryServices/vaults/{{.resourceName}}/replicationMigrationItems",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "ReplicationMigrationItems_List",
		Resource:    "Microsoft.RecoveryServices",
	},
	"ReplicationNetworkMappings_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.RecoveryServices/vaults/{{.resourceName}}/replicationNetworkMappings",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "ReplicationNetworkMappings_List",
		Resource:    "Microsoft.RecoveryServices",
	},
	"ReplicationNetworks_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.RecoveryServices/vaults/{{.resourceName}}/replicationNetworks",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "ReplicationNetworks_List",
		Resource:    "Microsoft.RecoveryServices",
	},
	"ReplicationPolicies_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.RecoveryServices/vaults/{{.resourceName}}/replicationPolicies",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "ReplicationPolicies_List",
		Resource:    "Microsoft.RecoveryServices",
	},
	"ReplicationPolicies_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.RecoveryServices/vaults/{{.resourceName}}/replicationPolicies/{{.policyName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "ReplicationPolicies_Get",
		Resource:    "Microsoft.RecoveryServices",
	},
	"ReplicationProtectedItems_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.RecoveryServices/vaults/{{.resourceName}}/replicationProtectedItems",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "ReplicationProtectedItems_List",
		Resource:    "Microsoft.RecoveryServices",
	},
	"ReplicationProtectionContainerMappings_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.RecoveryServices/vaults/{{.resourceName}}/replicationProtectionContainerMappings",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "ReplicationProtectionContainerMappings_List",
		Resource:    "Microsoft.RecoveryServices",
	},
	"ReplicationProtectionContainers_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.RecoveryServices/vaults/{{.resourceName}}/replicationProtectionContainers",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "ReplicationProtectionContainers_List",
		Resource:    "Microsoft.RecoveryServices",
	},
	"ReplicationProtectionIntents_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.RecoveryServices/vaults/{{.resourceName}}/replicationProtectionIntents",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "ReplicationProtectionIntents_List",
		Resource:    "Microsoft.RecoveryServices",
	},
	"ReplicationProtectionIntents_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.RecoveryServices/vaults/{{.resourceName}}/replicationProtectionIntents/{{.intentObjectName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "ReplicationProtectionIntents_Get",
		Resource:    "Microsoft.RecoveryServices",
	},
	"ReplicationRecoveryPlans_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.RecoveryServices/vaults/{{.resourceName}}/replicationRecoveryPlans",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "ReplicationRecoveryPlans_List",
		Resource:    "Microsoft.RecoveryServices",
	},
	"ReplicationRecoveryPlans_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.RecoveryServices/vaults/{{.resourceName}}/replicationRecoveryPlans/{{.recoveryPlanName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "ReplicationRecoveryPlans_Get",
		Resource:    "Microsoft.RecoveryServices",
	},
	"ReplicationRecoveryPlans_FailoverCancel": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.RecoveryServices/vaults/{{.resourceName}}/replicationRecoveryPlans/{{.recoveryPlanName}}/failoverCancel",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "ReplicationRecoveryPlans_FailoverCancel",
		Resource:    "Microsoft.RecoveryServices",
	},
	"ReplicationRecoveryPlans_FailoverCommit": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.RecoveryServices/vaults/{{.resourceName}}/replicationRecoveryPlans/{{.recoveryPlanName}}/failoverCommit",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "ReplicationRecoveryPlans_FailoverCommit",
		Resource:    "Microsoft.RecoveryServices",
	},
	"ReplicationRecoveryPlans_PlannedFailover": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.RecoveryServices/vaults/{{.resourceName}}/replicationRecoveryPlans/{{.recoveryPlanName}}/plannedFailover",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "ReplicationRecoveryPlans_PlannedFailover",
		Resource:    "Microsoft.RecoveryServices",
	},
	"ReplicationRecoveryPlans_Reprotect": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.RecoveryServices/vaults/{{.resourceName}}/replicationRecoveryPlans/{{.recoveryPlanName}}/reProtect",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "ReplicationRecoveryPlans_Reprotect",
		Resource:    "Microsoft.RecoveryServices",
	},
	"ReplicationRecoveryPlans_TestFailover": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.RecoveryServices/vaults/{{.resourceName}}/replicationRecoveryPlans/{{.recoveryPlanName}}/testFailover",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "ReplicationRecoveryPlans_TestFailover",
		Resource:    "Microsoft.RecoveryServices",
	},
	"ReplicationRecoveryPlans_TestFailoverCleanup": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.RecoveryServices/vaults/{{.resourceName}}/replicationRecoveryPlans/{{.recoveryPlanName}}/testFailoverCleanup",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "ReplicationRecoveryPlans_TestFailoverCleanup",
		Resource:    "Microsoft.RecoveryServices",
	},
	"ReplicationRecoveryPlans_UnplannedFailover": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.RecoveryServices/vaults/{{.resourceName}}/replicationRecoveryPlans/{{.recoveryPlanName}}/unplannedFailover",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "ReplicationRecoveryPlans_UnplannedFailover",
		Resource:    "Microsoft.RecoveryServices",
	},
	"ReplicationRecoveryServicesProviders_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.RecoveryServices/vaults/{{.resourceName}}/replicationRecoveryServicesProviders",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "ReplicationRecoveryServicesProviders_List",
		Resource:    "Microsoft.RecoveryServices",
	},
	"ReplicationStorageClassificationMappings_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.RecoveryServices/vaults/{{.resourceName}}/replicationStorageClassificationMappings",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "ReplicationStorageClassificationMappings_List",
		Resource:    "Microsoft.RecoveryServices",
	},
	"ReplicationStorageClassifications_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.RecoveryServices/vaults/{{.resourceName}}/replicationStorageClassifications",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "ReplicationStorageClassifications_List",
		Resource:    "Microsoft.RecoveryServices",
	},
	"SupportedOperatingSystems_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.RecoveryServices/vaults/{{.resourceName}}/replicationSupportedOperatingSystems",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "SupportedOperatingSystems_Get",
		Resource:    "Microsoft.RecoveryServices",
	},
	"ReplicationVaultHealth_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.RecoveryServices/vaults/{{.resourceName}}/replicationVaultHealth",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "ReplicationVaultHealth_Get",
		Resource:    "Microsoft.RecoveryServices",
	},
	"ReplicationVaultHealth_Refresh": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.RecoveryServices/vaults/{{.resourceName}}/replicationVaultHealth/default/refresh",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "ReplicationVaultHealth_Refresh",
		Resource:    "Microsoft.RecoveryServices",
	},
	"ReplicationVaultSetting_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.RecoveryServices/vaults/{{.resourceName}}/replicationVaultSettings",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "ReplicationVaultSetting_List",
		Resource:    "Microsoft.RecoveryServices",
	},
	"ReplicationVaultSetting_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.RecoveryServices/vaults/{{.resourceName}}/replicationVaultSettings/{{.vaultSettingName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "ReplicationVaultSetting_Get",
		Resource:    "Microsoft.RecoveryServices",
	},
	"ReplicationvCenters_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.RecoveryServices/vaults/{{.resourceName}}/replicationvCenters",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "ReplicationvCenters_List",
		Resource:    "Microsoft.RecoveryServices",
	},
}
