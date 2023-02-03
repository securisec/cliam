package policy

// Microsoft_DataProtection_dataprotection policy
var Microsoft_DataProtection_dataprotection = map[string]Policy{
	"BackupVaults_GetInSubscription": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.DataProtection/backupVaults",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2023-01-01",
		},
		OperationID: "BackupVaults_GetInSubscription",
		Resource:    "Microsoft.DataProtection",
	},
	"OperationResult_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.DataProtection/locations/{{.location}}/operationResults/{{.operationId}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2023-01-01",
		},
		OperationID: "OperationResult_Get",
		Resource:    "Microsoft.DataProtection",
	},
	"OperationStatus_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.DataProtection/locations/{{.location}}/operationStatus/{{.operationId}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2023-01-01",
		},
		OperationID: "OperationStatus_Get",
		Resource:    "Microsoft.DataProtection",
	},
	"OperationStatusBackupVaultContext_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DataProtection/backupVaults/{{.vaultName}}/operationStatus/{{.operationId}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2023-01-01",
		},
		OperationID: "OperationStatusBackupVaultContext_Get",
		Resource:    "Microsoft.DataProtection",
	},
	"OperationStatusResourceGroupContext_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DataProtection/operationStatus/{{.operationId}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2023-01-01",
		},
		OperationID: "OperationStatusResourceGroupContext_Get",
		Resource:    "Microsoft.DataProtection",
	},
	"BackupVaults_GetInResourceGroup": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DataProtection/backupVaults",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2023-01-01",
		},
		OperationID: "BackupVaults_GetInResourceGroup",
		Resource:    "Microsoft.DataProtection",
	},
	"BackupVaults_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DataProtection/backupVaults/{{.vaultName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2023-01-01",
		},
		OperationID: "BackupVaults_Get",
		Resource:    "Microsoft.DataProtection",
	},
	"BackupVaultOperationResults_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DataProtection/backupVaults/{{.vaultName}}/operationResults/{{.operationId}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2023-01-01",
		},
		OperationID: "BackupVaultOperationResults_Get",
		Resource:    "Microsoft.DataProtection",
	},
	"BackupVaults_CheckNameAvailability": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DataProtection/locations/{{.location}}/checkNameAvailability",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2023-01-01",
		},
		OperationID: "BackupVaults_CheckNameAvailability",
		Resource:    "Microsoft.DataProtection",
	},
	"DataProtection_CheckFeatureSupport": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.DataProtection/locations/{{.location}}/checkFeatureSupport",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2023-01-01",
		},
		OperationID: "DataProtection_CheckFeatureSupport",
		Resource:    "Microsoft.DataProtection",
	},
	"DataProtectionOperations_List": {
		Path:   "/providers/Microsoft.DataProtection/operations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2023-01-01",
		},
		OperationID: "DataProtectionOperations_List",
		Resource:    "Microsoft.DataProtection",
	},
	"BackupPolicies_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DataProtection/backupVaults/{{.vaultName}}/backupPolicies",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2023-01-01",
		},
		OperationID: "BackupPolicies_List",
		Resource:    "Microsoft.DataProtection",
	},
	"BackupPolicies_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DataProtection/backupVaults/{{.vaultName}}/backupPolicies/{{.backupPolicyName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2023-01-01",
		},
		OperationID: "BackupPolicies_Get",
		Resource:    "Microsoft.DataProtection",
	},
	"BackupInstances_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DataProtection/backupVaults/{{.vaultName}}/backupInstances",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2023-01-01",
		},
		OperationID: "BackupInstances_List",
		Resource:    "Microsoft.DataProtection",
	},
	"BackupInstances_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DataProtection/backupVaults/{{.vaultName}}/backupInstances/{{.backupInstanceName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2023-01-01",
		},
		OperationID: "BackupInstances_Get",
		Resource:    "Microsoft.DataProtection",
	},
	"BackupInstances_AdhocBackup": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DataProtection/backupVaults/{{.vaultName}}/backupInstances/{{.backupInstanceName}}/backup",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2023-01-01",
		},
		OperationID: "BackupInstances_AdhocBackup",
		Resource:    "Microsoft.DataProtection",
	},
	"BackupInstances_ValidateForBackup": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DataProtection/backupVaults/{{.vaultName}}/validateForBackup",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2023-01-01",
		},
		OperationID: "BackupInstances_ValidateForBackup",
		Resource:    "Microsoft.DataProtection",
	},
	"BackupInstances_GetBackupInstanceOperationResult": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DataProtection/backupVaults/{{.vaultName}}/backupInstances/{{.backupInstanceName}}/operationResults/{{.operationId}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2023-01-01",
		},
		OperationID: "BackupInstances_GetBackupInstanceOperationResult",
		Resource:    "Microsoft.DataProtection",
	},
	"RecoveryPoints_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DataProtection/backupVaults/{{.vaultName}}/backupInstances/{{.backupInstanceName}}/recoveryPoints",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2023-01-01",
		},
		OperationID: "RecoveryPoints_List",
		Resource:    "Microsoft.DataProtection",
	},
	"RecoveryPoints_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DataProtection/backupVaults/{{.vaultName}}/backupInstances/{{.backupInstanceName}}/recoveryPoints/{{.recoveryPointId}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2023-01-01",
		},
		OperationID: "RecoveryPoints_Get",
		Resource:    "Microsoft.DataProtection",
	},
	"BackupInstances_TriggerRehydrate": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DataProtection/backupVaults/{{.vaultName}}/backupInstances/{{.backupInstanceName}}/rehydrate",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2023-01-01",
		},
		OperationID: "BackupInstances_TriggerRehydrate",
		Resource:    "Microsoft.DataProtection",
	},
	"BackupInstances_TriggerRestore": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DataProtection/backupVaults/{{.vaultName}}/backupInstances/{{.backupInstanceName}}/restore",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2023-01-01",
		},
		OperationID: "BackupInstances_TriggerRestore",
		Resource:    "Microsoft.DataProtection",
	},
	"BackupInstances_ResumeBackups": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DataProtection/backupVaults/{{.vaultName}}/backupInstances/{{.backupInstanceName}}/resumeBackups",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2023-01-01",
		},
		OperationID: "BackupInstances_ResumeBackups",
		Resource:    "Microsoft.DataProtection",
	},
	"BackupInstances_ResumeProtection": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DataProtection/backupVaults/{{.vaultName}}/backupInstances/{{.backupInstanceName}}/resumeProtection",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2023-01-01",
		},
		OperationID: "BackupInstances_ResumeProtection",
		Resource:    "Microsoft.DataProtection",
	},
	"BackupInstances_StopProtection": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DataProtection/backupVaults/{{.vaultName}}/backupInstances/{{.backupInstanceName}}/stopProtection",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2023-01-01",
		},
		OperationID: "BackupInstances_StopProtection",
		Resource:    "Microsoft.DataProtection",
	},
	"BackupInstances_SuspendBackups": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DataProtection/backupVaults/{{.vaultName}}/backupInstances/{{.backupInstanceName}}/suspendBackups",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2023-01-01",
		},
		OperationID: "BackupInstances_SuspendBackups",
		Resource:    "Microsoft.DataProtection",
	},
	"BackupInstances_SyncBackupInstance": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DataProtection/backupVaults/{{.vaultName}}/backupInstances/{{.backupInstanceName}}/sync",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2023-01-01",
		},
		OperationID: "BackupInstances_SyncBackupInstance",
		Resource:    "Microsoft.DataProtection",
	},
	"BackupInstances_ValidateForRestore": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DataProtection/backupVaults/{{.vaultName}}/backupInstances/{{.backupInstanceName}}/validateRestore",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2023-01-01",
		},
		OperationID: "BackupInstances_ValidateForRestore",
		Resource:    "Microsoft.DataProtection",
	},
	"Jobs_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DataProtection/backupVaults/{{.vaultName}}/backupJobs",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2023-01-01",
		},
		OperationID: "Jobs_List",
		Resource:    "Microsoft.DataProtection",
	},
	"RestorableTimeRanges_Find": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DataProtection/backupVaults/{{.vaultName}}/backupInstances/{{.backupInstanceName}}/findRestorableTimeRanges",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2023-01-01",
		},
		OperationID: "RestorableTimeRanges_Find",
		Resource:    "Microsoft.DataProtection",
	},
	"Jobs_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DataProtection/backupVaults/{{.vaultName}}/backupJobs/{{.jobId}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2023-01-01",
		},
		OperationID: "Jobs_Get",
		Resource:    "Microsoft.DataProtection",
	},
	"ExportJobs_Trigger": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DataProtection/backupVaults/{{.vaultName}}/exportBackupJobs",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2023-01-01",
		},
		OperationID: "ExportJobs_Trigger",
		Resource:    "Microsoft.DataProtection",
	},
	"ExportJobsOperationResult_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DataProtection/backupVaults/{{.vaultName}}/backupJobs/operations/{{.operationId}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2023-01-01",
		},
		OperationID: "ExportJobsOperationResult_Get",
		Resource:    "Microsoft.DataProtection",
	},
	"DeletedBackupInstances_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DataProtection/backupVaults/{{.vaultName}}/deletedBackupInstances",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2023-01-01",
		},
		OperationID: "DeletedBackupInstances_List",
		Resource:    "Microsoft.DataProtection",
	},
	"DeletedBackupInstances_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DataProtection/backupVaults/{{.vaultName}}/deletedBackupInstances/{{.backupInstanceName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2023-01-01",
		},
		OperationID: "DeletedBackupInstances_Get",
		Resource:    "Microsoft.DataProtection",
	},
	"DeletedBackupInstances_Undelete": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DataProtection/backupVaults/{{.vaultName}}/deletedBackupInstances/{{.backupInstanceName}}/undelete",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2023-01-01",
		},
		OperationID: "DeletedBackupInstances_Undelete",
		Resource:    "Microsoft.DataProtection",
	},
	"ResourceGuards_GetResourcesInSubscription": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.DataProtection/resourceGuards",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2023-01-01",
		},
		OperationID: "ResourceGuards_GetResourcesInSubscription",
		Resource:    "Microsoft.DataProtection",
	},
	"ResourceGuards_GetResourcesInResourceGroup": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DataProtection/resourceGuards",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2023-01-01",
		},
		OperationID: "ResourceGuards_GetResourcesInResourceGroup",
		Resource:    "Microsoft.DataProtection",
	},
	"ResourceGuards_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DataProtection/resourceGuards/{{.resourceGuardsName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2023-01-01",
		},
		OperationID: "ResourceGuards_Get",
		Resource:    "Microsoft.DataProtection",
	},
	"ResourceGuards_GetDisableSoftDeleteRequestsObjects": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DataProtection/resourceGuards/{{.resourceGuardsName}}/disableSoftDeleteRequests",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2023-01-01",
		},
		OperationID: "ResourceGuards_GetDisableSoftDeleteRequestsObjects",
		Resource:    "Microsoft.DataProtection",
	},
	"ResourceGuards_GetDeleteResourceGuardProxyRequestsObjects": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DataProtection/resourceGuards/{{.resourceGuardsName}}/deleteResourceGuardProxyRequests",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2023-01-01",
		},
		OperationID: "ResourceGuards_GetDeleteResourceGuardProxyRequestsObjects",
		Resource:    "Microsoft.DataProtection",
	},
	"ResourceGuards_GetBackupSecurityPINRequestsObjects": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DataProtection/resourceGuards/{{.resourceGuardsName}}/getBackupSecurityPINRequests",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2023-01-01",
		},
		OperationID: "ResourceGuards_GetBackupSecurityPINRequestsObjects",
		Resource:    "Microsoft.DataProtection",
	},
	"ResourceGuards_GetDeleteProtectedItemRequestsObjects": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DataProtection/resourceGuards/{{.resourceGuardsName}}/deleteProtectedItemRequests",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2023-01-01",
		},
		OperationID: "ResourceGuards_GetDeleteProtectedItemRequestsObjects",
		Resource:    "Microsoft.DataProtection",
	},
	"ResourceGuards_GetUpdateProtectionPolicyRequestsObjects": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DataProtection/resourceGuards/{{.resourceGuardsName}}/updateProtectionPolicyRequests",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2023-01-01",
		},
		OperationID: "ResourceGuards_GetUpdateProtectionPolicyRequestsObjects",
		Resource:    "Microsoft.DataProtection",
	},
	"ResourceGuards_GetUpdateProtectedItemRequestsObjects": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DataProtection/resourceGuards/{{.resourceGuardsName}}/updateProtectedItemRequests",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2023-01-01",
		},
		OperationID: "ResourceGuards_GetUpdateProtectedItemRequestsObjects",
		Resource:    "Microsoft.DataProtection",
	},
	"ResourceGuards_GetDefaultDisableSoftDeleteRequestsObject": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DataProtection/resourceGuards/{{.resourceGuardsName}}/disableSoftDeleteRequests/{{.requestName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2023-01-01",
		},
		OperationID: "ResourceGuards_GetDefaultDisableSoftDeleteRequestsObject",
		Resource:    "Microsoft.DataProtection",
	},
	"ResourceGuards_GetDefaultDeleteResourceGuardProxyRequestsObject": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DataProtection/resourceGuards/{{.resourceGuardsName}}/deleteResourceGuardProxyRequests/{{.requestName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2023-01-01",
		},
		OperationID: "ResourceGuards_GetDefaultDeleteResourceGuardProxyRequestsObject",
		Resource:    "Microsoft.DataProtection",
	},
	"ResourceGuards_GetDefaultBackupSecurityPINRequestsObject": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DataProtection/resourceGuards/{{.resourceGuardsName}}/getBackupSecurityPINRequests/{{.requestName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2023-01-01",
		},
		OperationID: "ResourceGuards_GetDefaultBackupSecurityPINRequestsObject",
		Resource:    "Microsoft.DataProtection",
	},
	"ResourceGuards_GetDefaultDeleteProtectedItemRequestsObject": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DataProtection/resourceGuards/{{.resourceGuardsName}}/deleteProtectedItemRequests/{{.requestName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2023-01-01",
		},
		OperationID: "ResourceGuards_GetDefaultDeleteProtectedItemRequestsObject",
		Resource:    "Microsoft.DataProtection",
	},
	"ResourceGuards_GetDefaultUpdateProtectionPolicyRequestsObject": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DataProtection/resourceGuards/{{.resourceGuardsName}}/updateProtectionPolicyRequests/{{.requestName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2023-01-01",
		},
		OperationID: "ResourceGuards_GetDefaultUpdateProtectionPolicyRequestsObject",
		Resource:    "Microsoft.DataProtection",
	},
	"ResourceGuards_GetDefaultUpdateProtectedItemRequestsObject": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DataProtection/resourceGuards/{{.resourceGuardsName}}/updateProtectedItemRequests/{{.requestName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2023-01-01",
		},
		OperationID: "ResourceGuards_GetDefaultUpdateProtectedItemRequestsObject",
		Resource:    "Microsoft.DataProtection",
	},
}
