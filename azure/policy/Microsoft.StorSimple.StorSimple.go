package policy

// Microsoft_StorSimple_storsimple policy
var Microsoft_StorSimple_storsimple = map[string]Policy{
	"Operations_List": {
		Path:   "/providers/Microsoft.StorSimple/operations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2017-06-01",
		},
		OperationID: "Operations_List",
		Resource:    "Microsoft.StorSimple",
	},
	"Managers_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.StorSimple/managers",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2017-06-01",
		},
		OperationID: "Managers_List",
		Resource:    "Microsoft.StorSimple",
	},
	"Managers_ListByResourceGroup": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.StorSimple/managers",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2017-06-01",
		},
		OperationID: "Managers_ListByResourceGroup",
		Resource:    "Microsoft.StorSimple",
	},
	"Managers_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.StorSimple/managers/{{.managerName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2017-06-01",
		},
		OperationID: "Managers_Get",
		Resource:    "Microsoft.StorSimple",
	},
	"AccessControlRecords_ListByManager": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.StorSimple/managers/{{.managerName}}/accessControlRecords",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2017-06-01",
		},
		OperationID: "AccessControlRecords_ListByManager",
		Resource:    "Microsoft.StorSimple",
	},
	"AccessControlRecords_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.StorSimple/managers/{{.managerName}}/accessControlRecords/{{.accessControlRecordName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2017-06-01",
		},
		OperationID: "AccessControlRecords_Get",
		Resource:    "Microsoft.StorSimple",
	},
	"Alerts_ListByManager": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.StorSimple/managers/{{.managerName}}/alerts",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2017-06-01",
		},
		OperationID: "Alerts_ListByManager",
		Resource:    "Microsoft.StorSimple",
	},
	"BandwidthSettings_ListByManager": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.StorSimple/managers/{{.managerName}}/bandwidthSettings",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2017-06-01",
		},
		OperationID: "BandwidthSettings_ListByManager",
		Resource:    "Microsoft.StorSimple",
	},
	"BandwidthSettings_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.StorSimple/managers/{{.managerName}}/bandwidthSettings/{{.bandwidthSettingName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2017-06-01",
		},
		OperationID: "BandwidthSettings_Get",
		Resource:    "Microsoft.StorSimple",
	},
	"Alerts_Clear": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.StorSimple/managers/{{.managerName}}/clearAlerts",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2017-06-01",
		},
		OperationID: "Alerts_Clear",
		Resource:    "Microsoft.StorSimple",
	},
	"CloudAppliances_ListSupportedConfigurations": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.StorSimple/managers/{{.managerName}}/cloudApplianceConfigurations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2017-06-01",
		},
		OperationID: "CloudAppliances_ListSupportedConfigurations",
		Resource:    "Microsoft.StorSimple",
	},
	"Devices_Configure": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.StorSimple/managers/{{.managerName}}/configureDevice",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2017-06-01",
		},
		OperationID: "Devices_Configure",
		Resource:    "Microsoft.StorSimple",
	},
	"Devices_ListByManager": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.StorSimple/managers/{{.managerName}}/devices",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2017-06-01",
		},
		OperationID: "Devices_ListByManager",
		Resource:    "Microsoft.StorSimple",
	},
	"Devices_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.StorSimple/managers/{{.managerName}}/devices/{{.deviceName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2017-06-01",
		},
		OperationID: "Devices_Get",
		Resource:    "Microsoft.StorSimple",
	},
	"DeviceSettings_GetAlertSettings": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.StorSimple/managers/{{.managerName}}/devices/{{.deviceName}}/alertSettings/default",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2017-06-01",
		},
		OperationID: "DeviceSettings_GetAlertSettings",
		Resource:    "Microsoft.StorSimple",
	},
	"Devices_AuthorizeForServiceEncryptionKeyRollover": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.StorSimple/managers/{{.managerName}}/devices/{{.deviceName}}/authorizeForServiceEncryptionKeyRollover",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2017-06-01",
		},
		OperationID: "Devices_AuthorizeForServiceEncryptionKeyRollover",
		Resource:    "Microsoft.StorSimple",
	},
	"BackupPolicies_ListByDevice": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.StorSimple/managers/{{.managerName}}/devices/{{.deviceName}}/backupPolicies",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2017-06-01",
		},
		OperationID: "BackupPolicies_ListByDevice",
		Resource:    "Microsoft.StorSimple",
	},
	"BackupPolicies_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.StorSimple/managers/{{.managerName}}/devices/{{.deviceName}}/backupPolicies/{{.backupPolicyName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2017-06-01",
		},
		OperationID: "BackupPolicies_Get",
		Resource:    "Microsoft.StorSimple",
	},
	"BackupPolicies_BackupNow": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.StorSimple/managers/{{.managerName}}/devices/{{.deviceName}}/backupPolicies/{{.backupPolicyName}}/backup",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2017-06-01",
		},
		OperationID: "BackupPolicies_BackupNow",
		Resource:    "Microsoft.StorSimple",
	},
	"BackupSchedules_ListByBackupPolicy": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.StorSimple/managers/{{.managerName}}/devices/{{.deviceName}}/backupPolicies/{{.backupPolicyName}}/schedules",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2017-06-01",
		},
		OperationID: "BackupSchedules_ListByBackupPolicy",
		Resource:    "Microsoft.StorSimple",
	},
	"BackupSchedules_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.StorSimple/managers/{{.managerName}}/devices/{{.deviceName}}/backupPolicies/{{.backupPolicyName}}/schedules/{{.backupScheduleName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2017-06-01",
		},
		OperationID: "BackupSchedules_Get",
		Resource:    "Microsoft.StorSimple",
	},
	"Backups_ListByDevice": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.StorSimple/managers/{{.managerName}}/devices/{{.deviceName}}/backups",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2017-06-01",
		},
		OperationID: "Backups_ListByDevice",
		Resource:    "Microsoft.StorSimple",
	},
	"Backups_Clone": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.StorSimple/managers/{{.managerName}}/devices/{{.deviceName}}/backups/{{.backupName}}/elements/{{.backupElementName}}/clone",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2017-06-01",
		},
		OperationID: "Backups_Clone",
		Resource:    "Microsoft.StorSimple",
	},
	"Backups_Restore": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.StorSimple/managers/{{.managerName}}/devices/{{.deviceName}}/backups/{{.backupName}}/restore",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2017-06-01",
		},
		OperationID: "Backups_Restore",
		Resource:    "Microsoft.StorSimple",
	},
	"Devices_Deactivate": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.StorSimple/managers/{{.managerName}}/devices/{{.deviceName}}/deactivate",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2017-06-01",
		},
		OperationID: "Devices_Deactivate",
		Resource:    "Microsoft.StorSimple",
	},
	"HardwareComponentGroups_ListByDevice": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.StorSimple/managers/{{.managerName}}/devices/{{.deviceName}}/hardwareComponentGroups",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2017-06-01",
		},
		OperationID: "HardwareComponentGroups_ListByDevice",
		Resource:    "Microsoft.StorSimple",
	},
	"HardwareComponentGroups_ChangeControllerPowerState": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.StorSimple/managers/{{.managerName}}/devices/{{.deviceName}}/hardwareComponentGroups/{{.hardwareComponentGroupName}}/changeControllerPowerState",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2017-06-01",
		},
		OperationID: "HardwareComponentGroups_ChangeControllerPowerState",
		Resource:    "Microsoft.StorSimple",
	},
	"Devices_InstallUpdates": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.StorSimple/managers/{{.managerName}}/devices/{{.deviceName}}/installUpdates",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2017-06-01",
		},
		OperationID: "Devices_InstallUpdates",
		Resource:    "Microsoft.StorSimple",
	},
	"Jobs_ListByDevice": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.StorSimple/managers/{{.managerName}}/devices/{{.deviceName}}/jobs",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2017-06-01",
		},
		OperationID: "Jobs_ListByDevice",
		Resource:    "Microsoft.StorSimple",
	},
	"Jobs_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.StorSimple/managers/{{.managerName}}/devices/{{.deviceName}}/jobs/{{.jobName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2017-06-01",
		},
		OperationID: "Jobs_Get",
		Resource:    "Microsoft.StorSimple",
	},
	"Jobs_Cancel": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.StorSimple/managers/{{.managerName}}/devices/{{.deviceName}}/jobs/{{.jobName}}/cancel",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2017-06-01",
		},
		OperationID: "Jobs_Cancel",
		Resource:    "Microsoft.StorSimple",
	},
	"Devices_ListFailoverSets": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.StorSimple/managers/{{.managerName}}/devices/{{.deviceName}}/listFailoverSets",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2017-06-01",
		},
		OperationID: "Devices_ListFailoverSets",
		Resource:    "Microsoft.StorSimple",
	},
	"Devices_ListMetrics": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.StorSimple/managers/{{.managerName}}/devices/{{.deviceName}}/metrics",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2017-06-01",
		},
		OperationID: "Devices_ListMetrics",
		Resource:    "Microsoft.StorSimple",
	},
	"Devices_ListMetricDefinition": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.StorSimple/managers/{{.managerName}}/devices/{{.deviceName}}/metricsDefinitions",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2017-06-01",
		},
		OperationID: "Devices_ListMetricDefinition",
		Resource:    "Microsoft.StorSimple",
	},
	"DeviceSettings_GetNetworkSettings": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.StorSimple/managers/{{.managerName}}/devices/{{.deviceName}}/networkSettings/default",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2017-06-01",
		},
		OperationID: "DeviceSettings_GetNetworkSettings",
		Resource:    "Microsoft.StorSimple",
	},
	"Managers_GetDevicePublicEncryptionKey": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.StorSimple/managers/{{.managerName}}/devices/{{.deviceName}}/publicEncryptionKey",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2017-06-01",
		},
		OperationID: "Managers_GetDevicePublicEncryptionKey",
		Resource:    "Microsoft.StorSimple",
	},
	"Devices_ScanForUpdates": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.StorSimple/managers/{{.managerName}}/devices/{{.deviceName}}/scanForUpdates",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2017-06-01",
		},
		OperationID: "Devices_ScanForUpdates",
		Resource:    "Microsoft.StorSimple",
	},
	"DeviceSettings_GetSecuritySettings": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.StorSimple/managers/{{.managerName}}/devices/{{.deviceName}}/securitySettings/default",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2017-06-01",
		},
		OperationID: "DeviceSettings_GetSecuritySettings",
		Resource:    "Microsoft.StorSimple",
	},
	"DeviceSettings_SyncRemotemanagementCertificate": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.StorSimple/managers/{{.managerName}}/devices/{{.deviceName}}/securitySettings/default/syncRemoteManagementCertificate",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2017-06-01",
		},
		OperationID: "DeviceSettings_SyncRemotemanagementCertificate",
		Resource:    "Microsoft.StorSimple",
	},
	"Alerts_SendTestEmail": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.StorSimple/managers/{{.managerName}}/devices/{{.deviceName}}/sendTestAlertEmail",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2017-06-01",
		},
		OperationID: "Alerts_SendTestEmail",
		Resource:    "Microsoft.StorSimple",
	},
	"DeviceSettings_GetTimeSettings": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.StorSimple/managers/{{.managerName}}/devices/{{.deviceName}}/timeSettings/default",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2017-06-01",
		},
		OperationID: "DeviceSettings_GetTimeSettings",
		Resource:    "Microsoft.StorSimple",
	},
	"Devices_GetUpdateSummary": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.StorSimple/managers/{{.managerName}}/devices/{{.deviceName}}/updateSummary/default",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2017-06-01",
		},
		OperationID: "Devices_GetUpdateSummary",
		Resource:    "Microsoft.StorSimple",
	},
	"VolumeContainers_ListByDevice": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.StorSimple/managers/{{.managerName}}/devices/{{.deviceName}}/volumeContainers",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2017-06-01",
		},
		OperationID: "VolumeContainers_ListByDevice",
		Resource:    "Microsoft.StorSimple",
	},
	"VolumeContainers_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.StorSimple/managers/{{.managerName}}/devices/{{.deviceName}}/volumeContainers/{{.volumeContainerName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2017-06-01",
		},
		OperationID: "VolumeContainers_Get",
		Resource:    "Microsoft.StorSimple",
	},
	"VolumeContainers_ListMetrics": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.StorSimple/managers/{{.managerName}}/devices/{{.deviceName}}/volumeContainers/{{.volumeContainerName}}/metrics",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2017-06-01",
		},
		OperationID: "VolumeContainers_ListMetrics",
		Resource:    "Microsoft.StorSimple",
	},
	"VolumeContainers_ListMetricDefinition": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.StorSimple/managers/{{.managerName}}/devices/{{.deviceName}}/volumeContainers/{{.volumeContainerName}}/metricsDefinitions",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2017-06-01",
		},
		OperationID: "VolumeContainers_ListMetricDefinition",
		Resource:    "Microsoft.StorSimple",
	},
	"Volumes_ListByVolumeContainer": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.StorSimple/managers/{{.managerName}}/devices/{{.deviceName}}/volumeContainers/{{.volumeContainerName}}/volumes",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2017-06-01",
		},
		OperationID: "Volumes_ListByVolumeContainer",
		Resource:    "Microsoft.StorSimple",
	},
	"Volumes_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.StorSimple/managers/{{.managerName}}/devices/{{.deviceName}}/volumeContainers/{{.volumeContainerName}}/volumes/{{.volumeName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2017-06-01",
		},
		OperationID: "Volumes_Get",
		Resource:    "Microsoft.StorSimple",
	},
	"Volumes_ListMetrics": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.StorSimple/managers/{{.managerName}}/devices/{{.deviceName}}/volumeContainers/{{.volumeContainerName}}/volumes/{{.volumeName}}/metrics",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2017-06-01",
		},
		OperationID: "Volumes_ListMetrics",
		Resource:    "Microsoft.StorSimple",
	},
	"Volumes_ListMetricDefinition": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.StorSimple/managers/{{.managerName}}/devices/{{.deviceName}}/volumeContainers/{{.volumeContainerName}}/volumes/{{.volumeName}}/metricsDefinitions",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2017-06-01",
		},
		OperationID: "Volumes_ListMetricDefinition",
		Resource:    "Microsoft.StorSimple",
	},
	"Volumes_ListByDevice": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.StorSimple/managers/{{.managerName}}/devices/{{.deviceName}}/volumes",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2017-06-01",
		},
		OperationID: "Volumes_ListByDevice",
		Resource:    "Microsoft.StorSimple",
	},
	"Devices_Failover": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.StorSimple/managers/{{.managerName}}/devices/{{.sourceDeviceName}}/failover",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2017-06-01",
		},
		OperationID: "Devices_Failover",
		Resource:    "Microsoft.StorSimple",
	},
	"Devices_ListFailoverTargets": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.StorSimple/managers/{{.managerName}}/devices/{{.sourceDeviceName}}/listFailoverTargets",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2017-06-01",
		},
		OperationID: "Devices_ListFailoverTargets",
		Resource:    "Microsoft.StorSimple",
	},
	"Managers_GetEncryptionSettings": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.StorSimple/managers/{{.managerName}}/encryptionSettings/default",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2017-06-01",
		},
		OperationID: "Managers_GetEncryptionSettings",
		Resource:    "Microsoft.StorSimple",
	},
	"Managers_GetExtendedInfo": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.StorSimple/managers/{{.managerName}}/extendedInformation/vaultExtendedInfo",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2017-06-01",
		},
		OperationID: "Managers_GetExtendedInfo",
		Resource:    "Microsoft.StorSimple",
	},
	"Managers_ListFeatureSupportStatus": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.StorSimple/managers/{{.managerName}}/features",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2017-06-01",
		},
		OperationID: "Managers_ListFeatureSupportStatus",
		Resource:    "Microsoft.StorSimple",
	},
	"Jobs_ListByManager": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.StorSimple/managers/{{.managerName}}/jobs",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2017-06-01",
		},
		OperationID: "Jobs_ListByManager",
		Resource:    "Microsoft.StorSimple",
	},
	"Managers_GetActivationKey": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.StorSimple/managers/{{.managerName}}/listActivationKey",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2017-06-01",
		},
		OperationID: "Managers_GetActivationKey",
		Resource:    "Microsoft.StorSimple",
	},
	"Managers_GetPublicEncryptionKey": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.StorSimple/managers/{{.managerName}}/listPublicEncryptionKey",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2017-06-01",
		},
		OperationID: "Managers_GetPublicEncryptionKey",
		Resource:    "Microsoft.StorSimple",
	},
	"Managers_ListMetrics": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.StorSimple/managers/{{.managerName}}/metrics",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2017-06-01",
		},
		OperationID: "Managers_ListMetrics",
		Resource:    "Microsoft.StorSimple",
	},
	"Managers_ListMetricDefinition": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.StorSimple/managers/{{.managerName}}/metricsDefinitions",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2017-06-01",
		},
		OperationID: "Managers_ListMetricDefinition",
		Resource:    "Microsoft.StorSimple",
	},
	"CloudAppliances_Provision": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.StorSimple/managers/{{.managerName}}/provisionCloudAppliance",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2017-06-01",
		},
		OperationID: "CloudAppliances_Provision",
		Resource:    "Microsoft.StorSimple",
	},
	"Managers_RegenerateActivationKey": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.StorSimple/managers/{{.managerName}}/regenerateActivationKey",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2017-06-01",
		},
		OperationID: "Managers_RegenerateActivationKey",
		Resource:    "Microsoft.StorSimple",
	},
	"StorageAccountCredentials_ListByManager": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.StorSimple/managers/{{.managerName}}/storageAccountCredentials",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2017-06-01",
		},
		OperationID: "StorageAccountCredentials_ListByManager",
		Resource:    "Microsoft.StorSimple",
	},
	"StorageAccountCredentials_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.StorSimple/managers/{{.managerName}}/storageAccountCredentials/{{.storageAccountCredentialName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2017-06-01",
		},
		OperationID: "StorageAccountCredentials_Get",
		Resource:    "Microsoft.StorSimple",
	},
}
