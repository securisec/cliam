package policy

var Microsoft_RecoveryServices_bms = map[string]Policy{
	"BackupResourceStorageConfigsNonCRR_Get": {
		Path:   "/Subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.RecoveryServices/vaults/{{.vaultName}}/backupstorageconfig/vaultstorageconfig",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "BackupResourceStorageConfigsNonCRR_Get",
		Resource:    "Microsoft.RecoveryServices",
	},
	"ProtectionIntent_Validate": {
		Path:   "/Subscriptions/{{.subscriptionId}}/providers/Microsoft.RecoveryServices/locations/{{.azureRegion}}/backupPreValidateProtection",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "ProtectionIntent_Validate",
		Resource:    "Microsoft.RecoveryServices",
	},
	"BackupStatus_Get": {
		Path:   "/Subscriptions/{{.subscriptionId}}/providers/Microsoft.RecoveryServices/locations/{{.azureRegion}}/backupStatus",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "BackupStatus_Get",
		Resource:    "Microsoft.RecoveryServices",
	},
	"FeatureSupport_Validate": {
		Path:   "/Subscriptions/{{.subscriptionId}}/providers/Microsoft.RecoveryServices/locations/{{.azureRegion}}/backupValidateFeatures",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "FeatureSupport_Validate",
		Resource:    "Microsoft.RecoveryServices",
	},
	"ProtectionIntent_Get": {
		Path:   "/Subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.RecoveryServices/vaults/{{.vaultName}}/backupFabrics/{{.fabricName}}/backupProtectionIntent/{{.intentObjectName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "ProtectionIntent_Get",
		Resource:    "Microsoft.RecoveryServices",
	},
	"BackupProtectionIntent_List": {
		Path:   "/Subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.RecoveryServices/vaults/{{.vaultName}}/backupProtectionIntents",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "BackupProtectionIntent_List",
		Resource:    "Microsoft.RecoveryServices",
	},
	"BackupUsageSummaries_List": {
		Path:   "/Subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.RecoveryServices/vaults/{{.vaultName}}/backupUsageSummaries",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "BackupUsageSummaries_List",
		Resource:    "Microsoft.RecoveryServices",
	},
	"Operations_List": {
		Path:   "/providers/Microsoft.RecoveryServices/operations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "Operations_List",
		Resource:    "Microsoft.RecoveryServices",
	},
	"BackupResourceVaultConfigs_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.RecoveryServices/vaults/{{.vaultName}}/backupconfig/vaultconfig",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "BackupResourceVaultConfigs_Get",
		Resource:    "Microsoft.RecoveryServices",
	},
	"BackupResourceEncryptionConfigs_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.RecoveryServices/vaults/{{.vaultName}}/backupEncryptionConfigs/backupResourceEncryptionConfig",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "BackupResourceEncryptionConfigs_Get",
		Resource:    "Microsoft.RecoveryServices",
	},
	"PrivateEndpointConnection_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.RecoveryServices/vaults/{{.vaultName}}/privateEndpointConnections/{{.privateEndpointConnectionName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "PrivateEndpointConnection_Get",
		Resource:    "Microsoft.RecoveryServices",
	},
	"PrivateEndpoint_GetOperationStatus": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.RecoveryServices/vaults/{{.vaultName}}/privateEndpointConnections/{{.privateEndpointConnectionName}}/operationsStatus/{{.operationId}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "PrivateEndpoint_GetOperationStatus",
		Resource:    "Microsoft.RecoveryServices",
	},
	"GetOperationStatus": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.RecoveryServices/vaults/{{.vaultName}}/backupstorageconfig/vaultstorageconfig/operationStatus/{{.operationId}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "GetOperationStatus",
		Resource:    "Microsoft.RecoveryServices",
	},
	"BMSPrepareDataMove": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.RecoveryServices/vaults/{{.vaultName}}/backupstorageconfig/vaultstorageconfig/prepareDataMove",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "BMSPrepareDataMove",
		Resource:    "Microsoft.RecoveryServices",
	},
	"BMSPrepareDataMoveOperationResult_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.RecoveryServices/vaults/{{.vaultName}}/backupstorageconfig/vaultstorageconfig/operationResults/{{.operationId}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "BMSPrepareDataMoveOperationResult_Get",
		Resource:    "Microsoft.RecoveryServices",
	},
	"BMSTriggerDataMove": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.RecoveryServices/vaults/{{.vaultName}}/backupstorageconfig/vaultstorageconfig/triggerDataMove",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "BMSTriggerDataMove",
		Resource:    "Microsoft.RecoveryServices",
	},
	"ProtectedItems_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.RecoveryServices/vaults/{{.vaultName}}/backupFabrics/{{.fabricName}}/protectionContainers/{{.containerName}}/protectedItems/{{.protectedItemName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "ProtectedItems_Get",
		Resource:    "Microsoft.RecoveryServices",
	},
	"ProtectedItemOperationResults_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.RecoveryServices/vaults/{{.vaultName}}/backupFabrics/{{.fabricName}}/protectionContainers/{{.containerName}}/protectedItems/{{.protectedItemName}}/operationResults/{{.operationId}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "ProtectedItemOperationResults_Get",
		Resource:    "Microsoft.RecoveryServices",
	},
	"RecoveryPoints_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.RecoveryServices/vaults/{{.vaultName}}/backupFabrics/{{.fabricName}}/protectionContainers/{{.containerName}}/protectedItems/{{.protectedItemName}}/recoveryPoints",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "RecoveryPoints_List",
		Resource:    "Microsoft.RecoveryServices",
	},
	"RecoveryPoints_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.RecoveryServices/vaults/{{.vaultName}}/backupFabrics/{{.fabricName}}/protectionContainers/{{.containerName}}/protectedItems/{{.protectedItemName}}/recoveryPoints/{{.recoveryPointId}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "RecoveryPoints_Get",
		Resource:    "Microsoft.RecoveryServices",
	},
	"Restores_Trigger": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.RecoveryServices/vaults/{{.vaultName}}/backupFabrics/{{.fabricName}}/protectionContainers/{{.containerName}}/protectedItems/{{.protectedItemName}}/recoveryPoints/{{.recoveryPointId}}/restore",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "Restores_Trigger",
		Resource:    "Microsoft.RecoveryServices",
	},
	"BackupPolicies_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.RecoveryServices/vaults/{{.vaultName}}/backupPolicies",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "BackupPolicies_List",
		Resource:    "Microsoft.RecoveryServices",
	},
	"ProtectionPolicies_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.RecoveryServices/vaults/{{.vaultName}}/backupPolicies/{{.policyName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "ProtectionPolicies_Get",
		Resource:    "Microsoft.RecoveryServices",
	},
	"ProtectionPolicyOperationResults_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.RecoveryServices/vaults/{{.vaultName}}/backupPolicies/{{.policyName}}/operationResults/{{.operationId}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "ProtectionPolicyOperationResults_Get",
		Resource:    "Microsoft.RecoveryServices",
	},
	"BackupJobs_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.RecoveryServices/vaults/{{.vaultName}}/backupJobs",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "BackupJobs_List",
		Resource:    "Microsoft.RecoveryServices",
	},
	"JobDetails_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.RecoveryServices/vaults/{{.vaultName}}/backupJobs/{{.jobName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "JobDetails_Get",
		Resource:    "Microsoft.RecoveryServices",
	},
	"JobCancellations_Trigger": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.RecoveryServices/vaults/{{.vaultName}}/backupJobs/{{.jobName}}/cancel",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "JobCancellations_Trigger",
		Resource:    "Microsoft.RecoveryServices",
	},
	"JobOperationResults_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.RecoveryServices/vaults/{{.vaultName}}/backupJobs/{{.jobName}}/operationResults/{{.operationId}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "JobOperationResults_Get",
		Resource:    "Microsoft.RecoveryServices",
	},
	"ExportJobsOperationResults_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.RecoveryServices/vaults/{{.vaultName}}/backupJobs/operationResults/{{.operationId}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "ExportJobsOperationResults_Get",
		Resource:    "Microsoft.RecoveryServices",
	},
	"Jobs_Export": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.RecoveryServices/vaults/{{.vaultName}}/backupJobsExport",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "Jobs_Export",
		Resource:    "Microsoft.RecoveryServices",
	},
	"BackupProtectedItems_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.RecoveryServices/vaults/{{.vaultName}}/backupProtectedItems",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "BackupProtectedItems_List",
		Resource:    "Microsoft.RecoveryServices",
	},
	"Operation_Validate": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.RecoveryServices/vaults/{{.vaultName}}/backupValidateOperation",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "Operation_Validate",
		Resource:    "Microsoft.RecoveryServices",
	},
	"ValidateOperation_Trigger": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.RecoveryServices/vaults/{{.vaultName}}/backupTriggerValidateOperation",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "ValidateOperation_Trigger",
		Resource:    "Microsoft.RecoveryServices",
	},
	"ValidateOperationResults_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.RecoveryServices/vaults/{{.vaultName}}/backupValidateOperationResults/{{.operationId}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "ValidateOperationResults_Get",
		Resource:    "Microsoft.RecoveryServices",
	},
	"ValidateOperationStatuses_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.RecoveryServices/vaults/{{.vaultName}}/backupValidateOperationsStatuses/{{.operationId}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "ValidateOperationStatuses_Get",
		Resource:    "Microsoft.RecoveryServices",
	},
	"BackupEngines_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.RecoveryServices/vaults/{{.vaultName}}/backupEngines",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "BackupEngines_List",
		Resource:    "Microsoft.RecoveryServices",
	},
	"BackupEngines_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.RecoveryServices/vaults/{{.vaultName}}/backupEngines/{{.backupEngineName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "BackupEngines_Get",
		Resource:    "Microsoft.RecoveryServices",
	},
	"ProtectionContainerRefreshOperationResults_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.RecoveryServices/vaults/{{.vaultName}}/backupFabrics/{{.fabricName}}/operationResults/{{.operationId}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "ProtectionContainerRefreshOperationResults_Get",
		Resource:    "Microsoft.RecoveryServices",
	},
	"ProtectableContainers_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.RecoveryServices/vaults/{{.vaultName}}/backupFabrics/{{.fabricName}}/protectableContainers",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "ProtectableContainers_List",
		Resource:    "Microsoft.RecoveryServices",
	},
	"ProtectionContainers_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.RecoveryServices/vaults/{{.vaultName}}/backupFabrics/{{.fabricName}}/protectionContainers/{{.containerName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "ProtectionContainers_Get",
		Resource:    "Microsoft.RecoveryServices",
	},
	"ProtectionContainers_Inquire": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.RecoveryServices/vaults/{{.vaultName}}/backupFabrics/{{.fabricName}}/protectionContainers/{{.containerName}}/inquire",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "ProtectionContainers_Inquire",
		Resource:    "Microsoft.RecoveryServices",
	},
	"BackupWorkloadItems_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.RecoveryServices/vaults/{{.vaultName}}/backupFabrics/{{.fabricName}}/protectionContainers/{{.containerName}}/items",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "BackupWorkloadItems_List",
		Resource:    "Microsoft.RecoveryServices",
	},
	"ProtectionContainerOperationResults_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.RecoveryServices/vaults/{{.vaultName}}/backupFabrics/{{.fabricName}}/protectionContainers/{{.containerName}}/operationResults/{{.operationId}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "ProtectionContainerOperationResults_Get",
		Resource:    "Microsoft.RecoveryServices",
	},
	"Backups_Trigger": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.RecoveryServices/vaults/{{.vaultName}}/backupFabrics/{{.fabricName}}/protectionContainers/{{.containerName}}/protectedItems/{{.protectedItemName}}/backup",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "Backups_Trigger",
		Resource:    "Microsoft.RecoveryServices",
	},
	"ProtectedItemOperationStatuses_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.RecoveryServices/vaults/{{.vaultName}}/backupFabrics/{{.fabricName}}/protectionContainers/{{.containerName}}/protectedItems/{{.protectedItemName}}/operationsStatus/{{.operationId}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "ProtectedItemOperationStatuses_Get",
		Resource:    "Microsoft.RecoveryServices",
	},
	"ItemLevelRecoveryConnections_Provision": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.RecoveryServices/vaults/{{.vaultName}}/backupFabrics/{{.fabricName}}/protectionContainers/{{.containerName}}/protectedItems/{{.protectedItemName}}/recoveryPoints/{{.recoveryPointId}}/provisionInstantItemRecovery",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "ItemLevelRecoveryConnections_Provision",
		Resource:    "Microsoft.RecoveryServices",
	},
	"ItemLevelRecoveryConnections_Revoke": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.RecoveryServices/vaults/{{.vaultName}}/backupFabrics/{{.fabricName}}/protectionContainers/{{.containerName}}/protectedItems/{{.protectedItemName}}/recoveryPoints/{{.recoveryPointId}}/revokeInstantItemRecovery",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "ItemLevelRecoveryConnections_Revoke",
		Resource:    "Microsoft.RecoveryServices",
	},
	"ProtectionContainers_Refresh": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.RecoveryServices/vaults/{{.vaultName}}/backupFabrics/{{.fabricName}}/refreshContainers",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "ProtectionContainers_Refresh",
		Resource:    "Microsoft.RecoveryServices",
	},
	"BackupOperationResults_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.RecoveryServices/vaults/{{.vaultName}}/backupOperationResults/{{.operationId}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "BackupOperationResults_Get",
		Resource:    "Microsoft.RecoveryServices",
	},
	"BackupOperationStatuses_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.RecoveryServices/vaults/{{.vaultName}}/backupOperations/{{.operationId}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "BackupOperationStatuses_Get",
		Resource:    "Microsoft.RecoveryServices",
	},
	"ProtectionPolicyOperationStatuses_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.RecoveryServices/vaults/{{.vaultName}}/backupPolicies/{{.policyName}}/operations/{{.operationId}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "ProtectionPolicyOperationStatuses_Get",
		Resource:    "Microsoft.RecoveryServices",
	},
	"BackupProtectableItems_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.RecoveryServices/vaults/{{.vaultName}}/backupProtectableItems",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "BackupProtectableItems_List",
		Resource:    "Microsoft.RecoveryServices",
	},
	"BackupProtectionContainers_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.RecoveryServices/vaults/{{.vaultName}}/backupProtectionContainers",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "BackupProtectionContainers_List",
		Resource:    "Microsoft.RecoveryServices",
	},
	"SecurityPINs_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.RecoveryServices/vaults/{{.vaultName}}/backupSecurityPIN",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "SecurityPINs_Get",
		Resource:    "Microsoft.RecoveryServices",
	},
	"MoveRecoveryPoint": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.RecoveryServices/vaults/{{.vaultName}}/backupFabrics/{{.fabricName}}/protectionContainers/{{.containerName}}/protectedItems/{{.protectedItemName}}/recoveryPoints/{{.recoveryPointId}}/move",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "MoveRecoveryPoint",
		Resource:    "Microsoft.RecoveryServices",
	},
	"RecoveryPointsRecommendedForMove_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.RecoveryServices/vaults/{{.vaultName}}/backupFabrics/{{.fabricName}}/protectionContainers/{{.containerName}}/protectedItems/{{.protectedItemName}}/recoveryPointsRecommendedForMove",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "RecoveryPointsRecommendedForMove_List",
		Resource:    "Microsoft.RecoveryServices",
	},
	"ResourceGuardProxies_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.RecoveryServices/vaults/{{.vaultName}}/backupResourceGuardProxies",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "ResourceGuardProxies_Get",
		Resource:    "Microsoft.RecoveryServices",
	},
	"ResourceGuardProxy_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.RecoveryServices/vaults/{{.vaultName}}/backupResourceGuardProxies/{{.resourceGuardProxyName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "ResourceGuardProxy_Get",
		Resource:    "Microsoft.RecoveryServices",
	},
	"ResourceGuardProxy_UnlockDelete": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.RecoveryServices/vaults/{{.vaultName}}/backupResourceGuardProxies/{{.resourceGuardProxyName}}/unlockDelete",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "ResourceGuardProxy_UnlockDelete",
		Resource:    "Microsoft.RecoveryServices",
	},
}
