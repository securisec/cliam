package policy

// Microsoft_Web_WebApps policy
var Microsoft_Web_WebApps = map[string]Policy{
	"WebApps_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Web/sites",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_List",
		Resource:    "Microsoft.Web",
	},
	"WebApps_ListByResourceGroup": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_ListByResourceGroup",
		Resource:    "Microsoft.Web",
	},
	"WebApps_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_Get",
		Resource:    "Microsoft.Web",
	},
	"WebApps_AnalyzeCustomHostname": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/analyzeCustomHostname",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_AnalyzeCustomHostname",
		Resource:    "Microsoft.Web",
	},
	"WebApps_ApplySlotConfigToProduction": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/applySlotConfig",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_ApplySlotConfigToProduction",
		Resource:    "Microsoft.Web",
	},
	"WebApps_Backup": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/backup",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_Backup",
		Resource:    "Microsoft.Web",
	},
	"WebApps_ListBackups": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/backups",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_ListBackups",
		Resource:    "Microsoft.Web",
	},
	"WebApps_GetBackupStatus": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/backups/{{.backupId}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_GetBackupStatus",
		Resource:    "Microsoft.Web",
	},
	"WebApps_ListBackupStatusSecrets": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/backups/{{.backupId}}/list",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_ListBackupStatusSecrets",
		Resource:    "Microsoft.Web",
	},
	"WebApps_Restore": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/backups/{{.backupId}}/restore",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_Restore",
		Resource:    "Microsoft.Web",
	},
	"WebApps_ListBasicPublishingCredentialsPolicies": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/basicPublishingCredentialsPolicies",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_ListBasicPublishingCredentialsPolicies",
		Resource:    "Microsoft.Web",
	},
	"WebApps_GetFtpAllowed": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/basicPublishingCredentialsPolicies/ftp",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_GetFtpAllowed",
		Resource:    "Microsoft.Web",
	},
	"WebApps_GetScmAllowed": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/basicPublishingCredentialsPolicies/scm",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_GetScmAllowed",
		Resource:    "Microsoft.Web",
	},
	"WebApps_ListConfigurations": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/config",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_ListConfigurations",
		Resource:    "Microsoft.Web",
	},
	"WebApps_ListApplicationSettings": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/config/appsettings/list",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_ListApplicationSettings",
		Resource:    "Microsoft.Web",
	},
	"WebApps_GetAuthSettings": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/config/authsettings/list",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_GetAuthSettings",
		Resource:    "Microsoft.Web",
	},
	"WebApps_GetAuthSettingsV2WithoutSecrets": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/config/authsettingsV2",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_GetAuthSettingsV2WithoutSecrets",
		Resource:    "Microsoft.Web",
	},
	"WebApps_GetAuthSettingsV2": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/config/authsettingsV2/list",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_GetAuthSettingsV2",
		Resource:    "Microsoft.Web",
	},
	"WebApps_ListAzureStorageAccounts": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/config/azurestorageaccounts/list",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_ListAzureStorageAccounts",
		Resource:    "Microsoft.Web",
	},
	"WebApps_GetBackupConfiguration": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/config/backup/list",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_GetBackupConfiguration",
		Resource:    "Microsoft.Web",
	},
	"WebApps_GetAppSettingsKeyVaultReferences": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/config/configreferences/appsettings",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_GetAppSettingsKeyVaultReferences",
		Resource:    "Microsoft.Web",
	},
	"WebApps_GetAppSettingKeyVaultReference": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/config/configreferences/appsettings/{{.appSettingKey}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_GetAppSettingKeyVaultReference",
		Resource:    "Microsoft.Web",
	},
	"WebApps_GetSiteConnectionStringKeyVaultReferences": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/config/configreferences/connectionstrings",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_GetSiteConnectionStringKeyVaultReferences",
		Resource:    "Microsoft.Web",
	},
	"WebApps_GetSiteConnectionStringKeyVaultReference": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/config/configreferences/connectionstrings/{{.connectionStringKey}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_GetSiteConnectionStringKeyVaultReference",
		Resource:    "Microsoft.Web",
	},
	"WebApps_ListConnectionStrings": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/config/connectionstrings/list",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_ListConnectionStrings",
		Resource:    "Microsoft.Web",
	},
	"WebApps_GetDiagnosticLogsConfiguration": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/config/logs",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_GetDiagnosticLogsConfiguration",
		Resource:    "Microsoft.Web",
	},
	"WebApps_ListMetadata": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/config/metadata/list",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_ListMetadata",
		Resource:    "Microsoft.Web",
	},
	"WebApps_ListPublishingCredentials": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/config/publishingcredentials/list",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_ListPublishingCredentials",
		Resource:    "Microsoft.Web",
	},
	"WebApps_ListSitePushSettings": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/config/pushsettings/list",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_ListSitePushSettings",
		Resource:    "Microsoft.Web",
	},
	"WebApps_ListSlotConfigurationNames": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/config/slotConfigNames",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_ListSlotConfigurationNames",
		Resource:    "Microsoft.Web",
	},
	"WebApps_GetConfiguration": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/config/web",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_GetConfiguration",
		Resource:    "Microsoft.Web",
	},
	"WebApps_ListConfigurationSnapshotInfo": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/config/web/snapshots",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_ListConfigurationSnapshotInfo",
		Resource:    "Microsoft.Web",
	},
	"WebApps_GetConfigurationSnapshot": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/config/web/snapshots/{{.snapshotId}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_GetConfigurationSnapshot",
		Resource:    "Microsoft.Web",
	},
	"WebApps_RecoverSiteConfigurationSnapshot": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/config/web/snapshots/{{.snapshotId}}/recover",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_RecoverSiteConfigurationSnapshot",
		Resource:    "Microsoft.Web",
	},
	"WebApps_GetWebSiteContainerLogs": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/containerlogs",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_GetWebSiteContainerLogs",
		Resource:    "Microsoft.Web",
	},
	"WebApps_GetContainerLogsZip": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/containerlogs/zip/download",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_GetContainerLogsZip",
		Resource:    "Microsoft.Web",
	},
	"WebApps_ListContinuousWebJobs": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/continuouswebjobs",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_ListContinuousWebJobs",
		Resource:    "Microsoft.Web",
	},
	"WebApps_GetContinuousWebJob": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/continuouswebjobs/{{.webJobName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_GetContinuousWebJob",
		Resource:    "Microsoft.Web",
	},
	"WebApps_StartContinuousWebJob": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/continuouswebjobs/{{.webJobName}}/start",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_StartContinuousWebJob",
		Resource:    "Microsoft.Web",
	},
	"WebApps_StopContinuousWebJob": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/continuouswebjobs/{{.webJobName}}/stop",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_StopContinuousWebJob",
		Resource:    "Microsoft.Web",
	},
	"WebApps_ListProductionSiteDeploymentStatuses": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/deploymentStatus",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_ListProductionSiteDeploymentStatuses",
		Resource:    "Microsoft.Web",
	},
	"WebApps_GetProductionSiteDeploymentStatus": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/deploymentStatus/{{.deploymentStatusId}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_GetProductionSiteDeploymentStatus",
		Resource:    "Microsoft.Web",
	},
	"WebApps_ListDeployments": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/deployments",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_ListDeployments",
		Resource:    "Microsoft.Web",
	},
	"WebApps_GetDeployment": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/deployments/{{.id}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_GetDeployment",
		Resource:    "Microsoft.Web",
	},
	"WebApps_ListDeploymentLog": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/deployments/{{.id}}/log",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_ListDeploymentLog",
		Resource:    "Microsoft.Web",
	},
	"WebApps_DiscoverBackup": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/discoverbackup",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_DiscoverBackup",
		Resource:    "Microsoft.Web",
	},
	"WebApps_ListDomainOwnershipIdentifiers": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/domainOwnershipIdentifiers",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_ListDomainOwnershipIdentifiers",
		Resource:    "Microsoft.Web",
	},
	"WebApps_GetDomainOwnershipIdentifier": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/domainOwnershipIdentifiers/{{.domainOwnershipIdentifierName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_GetDomainOwnershipIdentifier",
		Resource:    "Microsoft.Web",
	},
	"WebApps_GetMSDeployStatus": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/extensions/MSDeploy",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_GetMSDeployStatus",
		Resource:    "Microsoft.Web",
	},
	"WebApps_GetMSDeployLog": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/extensions/MSDeploy/log",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_GetMSDeployLog",
		Resource:    "Microsoft.Web",
	},
	"WebApps_GetOneDeployStatus": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/extensions/onedeploy",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_GetOneDeployStatus",
		Resource:    "Microsoft.Web",
	},
	"WebApps_ListFunctions": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/functions",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_ListFunctions",
		Resource:    "Microsoft.Web",
	},
	"WebApps_GetFunctionsAdminToken": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/functions/admin/token",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_GetFunctionsAdminToken",
		Resource:    "Microsoft.Web",
	},
	"WebApps_GetFunction": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/functions/{{.functionName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_GetFunction",
		Resource:    "Microsoft.Web",
	},
	"WebApps_ListFunctionKeys": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/functions/{{.functionName}}/listkeys",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_ListFunctionKeys",
		Resource:    "Microsoft.Web",
	},
	"WebApps_ListFunctionSecrets": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/functions/{{.functionName}}/listsecrets",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_ListFunctionSecrets",
		Resource:    "Microsoft.Web",
	},
	"WebApps_ListHostKeys": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/host/default/listkeys",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_ListHostKeys",
		Resource:    "Microsoft.Web",
	},
	"WebApps_ListSyncStatus": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/host/default/listsyncstatus",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_ListSyncStatus",
		Resource:    "Microsoft.Web",
	},
	"WebApps_SyncFunctions": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/host/default/sync",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_SyncFunctions",
		Resource:    "Microsoft.Web",
	},
	"WebApps_ListHostNameBindings": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/hostNameBindings",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_ListHostNameBindings",
		Resource:    "Microsoft.Web",
	},
	"WebApps_GetHostNameBinding": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/hostNameBindings/{{.hostName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_GetHostNameBinding",
		Resource:    "Microsoft.Web",
	},
	"WebApps_GetHybridConnection": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/hybridConnectionNamespaces/{{.namespaceName}}/relays/{{.relayName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_GetHybridConnection",
		Resource:    "Microsoft.Web",
	},
	"WebApps_ListHybridConnections": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/hybridConnectionRelays",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_ListHybridConnections",
		Resource:    "Microsoft.Web",
	},
	"WebApps_ListRelayServiceConnections": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/hybridconnection",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_ListRelayServiceConnections",
		Resource:    "Microsoft.Web",
	},
	"WebApps_GetRelayServiceConnection": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/hybridconnection/{{.entityName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_GetRelayServiceConnection",
		Resource:    "Microsoft.Web",
	},
	"WebApps_ListInstanceIdentifiers": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/instances",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_ListInstanceIdentifiers",
		Resource:    "Microsoft.Web",
	},
	"WebApps_GetInstanceInfo": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/instances/{{.instanceId}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_GetInstanceInfo",
		Resource:    "Microsoft.Web",
	},
	"WebApps_GetInstanceMsDeployStatus": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/instances/{{.instanceId}}/extensions/MSDeploy",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_GetInstanceMsDeployStatus",
		Resource:    "Microsoft.Web",
	},
	"WebApps_GetInstanceMSDeployLog": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/instances/{{.instanceId}}/extensions/MSDeploy/log",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_GetInstanceMSDeployLog",
		Resource:    "Microsoft.Web",
	},
	"WebApps_ListInstanceProcesses": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/instances/{{.instanceId}}/processes",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_ListInstanceProcesses",
		Resource:    "Microsoft.Web",
	},
	"WebApps_GetInstanceProcess": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/instances/{{.instanceId}}/processes/{{.processId}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_GetInstanceProcess",
		Resource:    "Microsoft.Web",
	},
	"WebApps_GetInstanceProcessDump": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/instances/{{.instanceId}}/processes/{{.processId}}/dump",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_GetInstanceProcessDump",
		Resource:    "Microsoft.Web",
	},
	"WebApps_ListInstanceProcessModules": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/instances/{{.instanceId}}/processes/{{.processId}}/modules",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_ListInstanceProcessModules",
		Resource:    "Microsoft.Web",
	},
	"WebApps_GetInstanceProcessModule": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/instances/{{.instanceId}}/processes/{{.processId}}/modules/{{.baseAddress}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_GetInstanceProcessModule",
		Resource:    "Microsoft.Web",
	},
	"WebApps_ListInstanceProcessThreads": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/instances/{{.instanceId}}/processes/{{.processId}}/threads",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_ListInstanceProcessThreads",
		Resource:    "Microsoft.Web",
	},
	"WebApps_IsCloneable": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/iscloneable",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_IsCloneable",
		Resource:    "Microsoft.Web",
	},
	"WebApps_ListSiteBackups": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/listbackups",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_ListSiteBackups",
		Resource:    "Microsoft.Web",
	},
	"WebApps_ListSyncFunctionTriggers": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/listsyncfunctiontriggerstatus",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_ListSyncFunctionTriggers",
		Resource:    "Microsoft.Web",
	},
	"WebApps_MigrateMySql": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/migratemysql",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_MigrateMySql",
		Resource:    "Microsoft.Web",
	},
	"WebApps_GetMigrateMySqlStatus": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/migratemysql/status",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_GetMigrateMySqlStatus",
		Resource:    "Microsoft.Web",
	},
	"WebApps_GetSwiftVirtualNetworkConnection": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/networkConfig/virtualNetwork",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_GetSwiftVirtualNetworkConnection",
		Resource:    "Microsoft.Web",
	},
	"WebApps_ListNetworkFeatures": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/networkFeatures/{{.view}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_ListNetworkFeatures",
		Resource:    "Microsoft.Web",
	},
	"WebApps_GetNetworkTraceOperation": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/networkTrace/operationresults/{{.operationId}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_GetNetworkTraceOperation",
		Resource:    "Microsoft.Web",
	},
	"WebApps_StartWebSiteNetworkTrace": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/networkTrace/start",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_StartWebSiteNetworkTrace",
		Resource:    "Microsoft.Web",
	},
	"WebApps_StartWebSiteNetworkTraceOperation": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/networkTrace/startOperation",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_StartWebSiteNetworkTraceOperation",
		Resource:    "Microsoft.Web",
	},
	"WebApps_StopWebSiteNetworkTrace": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/networkTrace/stop",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_StopWebSiteNetworkTrace",
		Resource:    "Microsoft.Web",
	},
	"WebApps_GetNetworkTraces": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/networkTrace/{{.operationId}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_GetNetworkTraces",
		Resource:    "Microsoft.Web",
	},
	"WebApps_GetNetworkTraceOperationV2": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/networkTraces/current/operationresults/{{.operationId}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_GetNetworkTraceOperationV2",
		Resource:    "Microsoft.Web",
	},
	"WebApps_GetNetworkTracesV2": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/networkTraces/{{.operationId}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_GetNetworkTracesV2",
		Resource:    "Microsoft.Web",
	},
	"WebApps_GenerateNewSitePublishingPassword": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/newpassword",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_GenerateNewSitePublishingPassword",
		Resource:    "Microsoft.Web",
	},
	"WebApps_ListPerfMonCounters": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/perfcounters",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_ListPerfMonCounters",
		Resource:    "Microsoft.Web",
	},
	"WebApps_GetSitePhpErrorLogFlag": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/phplogging",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_GetSitePhpErrorLogFlag",
		Resource:    "Microsoft.Web",
	},
	"WebApps_ListPremierAddOns": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/premieraddons",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_ListPremierAddOns",
		Resource:    "Microsoft.Web",
	},
	"WebApps_GetPremierAddOn": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/premieraddons/{{.premierAddOnName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_GetPremierAddOn",
		Resource:    "Microsoft.Web",
	},
	"WebApps_GetPrivateAccess": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/privateAccess/virtualNetworks",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_GetPrivateAccess",
		Resource:    "Microsoft.Web",
	},
	"WebApps_GetPrivateEndpointConnectionList": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/privateEndpointConnections",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_GetPrivateEndpointConnectionList",
		Resource:    "Microsoft.Web",
	},
	"WebApps_GetPrivateEndpointConnection": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/privateEndpointConnections/{{.privateEndpointConnectionName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_GetPrivateEndpointConnection",
		Resource:    "Microsoft.Web",
	},
	"WebApps_GetPrivateLinkResources": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/privateLinkResources",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_GetPrivateLinkResources",
		Resource:    "Microsoft.Web",
	},
	"WebApps_ListProcesses": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/processes",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_ListProcesses",
		Resource:    "Microsoft.Web",
	},
	"WebApps_GetProcess": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/processes/{{.processId}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_GetProcess",
		Resource:    "Microsoft.Web",
	},
	"WebApps_GetProcessDump": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/processes/{{.processId}}/dump",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_GetProcessDump",
		Resource:    "Microsoft.Web",
	},
	"WebApps_ListProcessModules": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/processes/{{.processId}}/modules",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_ListProcessModules",
		Resource:    "Microsoft.Web",
	},
	"WebApps_GetProcessModule": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/processes/{{.processId}}/modules/{{.baseAddress}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_GetProcessModule",
		Resource:    "Microsoft.Web",
	},
	"WebApps_ListProcessThreads": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/processes/{{.processId}}/threads",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_ListProcessThreads",
		Resource:    "Microsoft.Web",
	},
	"WebApps_ListPublicCertificates": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/publicCertificates",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_ListPublicCertificates",
		Resource:    "Microsoft.Web",
	},
	"WebApps_GetPublicCertificate": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/publicCertificates/{{.publicCertificateName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_GetPublicCertificate",
		Resource:    "Microsoft.Web",
	},
	"WebApps_ListPublishingProfileXmlWithSecrets": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/publishxml",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_ListPublishingProfileXmlWithSecrets",
		Resource:    "Microsoft.Web",
	},
	"WebApps_ResetProductionSlotConfig": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/resetSlotConfig",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_ResetProductionSlotConfig",
		Resource:    "Microsoft.Web",
	},
	"WebApps_Restart": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/restart",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_Restart",
		Resource:    "Microsoft.Web",
	},
	"WebApps_RestoreFromBackupBlob": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/restoreFromBackupBlob",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_RestoreFromBackupBlob",
		Resource:    "Microsoft.Web",
	},
	"WebApps_RestoreFromDeletedApp": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/restoreFromDeletedApp",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_RestoreFromDeletedApp",
		Resource:    "Microsoft.Web",
	},
	"WebApps_RestoreSnapshot": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/restoreSnapshot",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_RestoreSnapshot",
		Resource:    "Microsoft.Web",
	},
	"WebApps_ListSiteExtensions": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/siteextensions",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_ListSiteExtensions",
		Resource:    "Microsoft.Web",
	},
	"WebApps_GetSiteExtension": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/siteextensions/{{.siteExtensionId}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_GetSiteExtension",
		Resource:    "Microsoft.Web",
	},
	"WebApps_ListSlots": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_ListSlots",
		Resource:    "Microsoft.Web",
	},
	"WebApps_GetSlot": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_GetSlot",
		Resource:    "Microsoft.Web",
	},
	"WebApps_AnalyzeCustomHostnameSlot": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/analyzeCustomHostname",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_AnalyzeCustomHostnameSlot",
		Resource:    "Microsoft.Web",
	},
	"WebApps_ApplySlotConfigurationSlot": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/applySlotConfig",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_ApplySlotConfigurationSlot",
		Resource:    "Microsoft.Web",
	},
	"WebApps_BackupSlot": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/backup",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_BackupSlot",
		Resource:    "Microsoft.Web",
	},
	"WebApps_ListBackupsSlot": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/backups",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_ListBackupsSlot",
		Resource:    "Microsoft.Web",
	},
	"WebApps_GetBackupStatusSlot": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/backups/{{.backupId}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_GetBackupStatusSlot",
		Resource:    "Microsoft.Web",
	},
	"WebApps_ListBackupStatusSecretsSlot": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/backups/{{.backupId}}/list",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_ListBackupStatusSecretsSlot",
		Resource:    "Microsoft.Web",
	},
	"WebApps_RestoreSlot": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/backups/{{.backupId}}/restore",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_RestoreSlot",
		Resource:    "Microsoft.Web",
	},
	"WebApps_ListBasicPublishingCredentialsPoliciesSlot": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/basicPublishingCredentialsPolicies",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_ListBasicPublishingCredentialsPoliciesSlot",
		Resource:    "Microsoft.Web",
	},
	"WebApps_GetFtpAllowedSlot": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/basicPublishingCredentialsPolicies/ftp",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_GetFtpAllowedSlot",
		Resource:    "Microsoft.Web",
	},
	"WebApps_GetScmAllowedSlot": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/basicPublishingCredentialsPolicies/scm",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_GetScmAllowedSlot",
		Resource:    "Microsoft.Web",
	},
	"WebApps_ListConfigurationsSlot": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/config",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_ListConfigurationsSlot",
		Resource:    "Microsoft.Web",
	},
	"WebApps_ListApplicationSettingsSlot": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/config/appsettings/list",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_ListApplicationSettingsSlot",
		Resource:    "Microsoft.Web",
	},
	"WebApps_GetAuthSettingsSlot": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/config/authsettings/list",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_GetAuthSettingsSlot",
		Resource:    "Microsoft.Web",
	},
	"WebApps_GetAuthSettingsV2WithoutSecretsSlot": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/config/authsettingsV2",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_GetAuthSettingsV2WithoutSecretsSlot",
		Resource:    "Microsoft.Web",
	},
	"WebApps_GetAuthSettingsV2Slot": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/config/authsettingsV2/list",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_GetAuthSettingsV2Slot",
		Resource:    "Microsoft.Web",
	},
	"WebApps_ListAzureStorageAccountsSlot": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/config/azurestorageaccounts/list",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_ListAzureStorageAccountsSlot",
		Resource:    "Microsoft.Web",
	},
	"WebApps_GetBackupConfigurationSlot": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/config/backup/list",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_GetBackupConfigurationSlot",
		Resource:    "Microsoft.Web",
	},
	"WebApps_GetAppSettingsKeyVaultReferencesSlot": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/config/configreferences/appsettings",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_GetAppSettingsKeyVaultReferencesSlot",
		Resource:    "Microsoft.Web",
	},
	"WebApps_GetAppSettingKeyVaultReferenceSlot": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/config/configreferences/appsettings/{{.appSettingKey}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_GetAppSettingKeyVaultReferenceSlot",
		Resource:    "Microsoft.Web",
	},
	"WebApps_GetSiteConnectionStringKeyVaultReferencesSlot": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/config/configreferences/connectionstrings",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_GetSiteConnectionStringKeyVaultReferencesSlot",
		Resource:    "Microsoft.Web",
	},
	"WebApps_GetSiteConnectionStringKeyVaultReferenceSlot": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/config/configreferences/connectionstrings/{{.connectionStringKey}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_GetSiteConnectionStringKeyVaultReferenceSlot",
		Resource:    "Microsoft.Web",
	},
	"WebApps_ListConnectionStringsSlot": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/config/connectionstrings/list",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_ListConnectionStringsSlot",
		Resource:    "Microsoft.Web",
	},
	"WebApps_GetDiagnosticLogsConfigurationSlot": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/config/logs",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_GetDiagnosticLogsConfigurationSlot",
		Resource:    "Microsoft.Web",
	},
	"WebApps_ListMetadataSlot": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/config/metadata/list",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_ListMetadataSlot",
		Resource:    "Microsoft.Web",
	},
	"WebApps_ListPublishingCredentialsSlot": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/config/publishingcredentials/list",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_ListPublishingCredentialsSlot",
		Resource:    "Microsoft.Web",
	},
	"WebApps_ListSitePushSettingsSlot": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/config/pushsettings/list",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_ListSitePushSettingsSlot",
		Resource:    "Microsoft.Web",
	},
	"WebApps_GetConfigurationSlot": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/config/web",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_GetConfigurationSlot",
		Resource:    "Microsoft.Web",
	},
	"WebApps_ListConfigurationSnapshotInfoSlot": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/config/web/snapshots",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_ListConfigurationSnapshotInfoSlot",
		Resource:    "Microsoft.Web",
	},
	"WebApps_GetConfigurationSnapshotSlot": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/config/web/snapshots/{{.snapshotId}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_GetConfigurationSnapshotSlot",
		Resource:    "Microsoft.Web",
	},
	"WebApps_RecoverSiteConfigurationSnapshotSlot": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/config/web/snapshots/{{.snapshotId}}/recover",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_RecoverSiteConfigurationSnapshotSlot",
		Resource:    "Microsoft.Web",
	},
	"WebApps_GetWebSiteContainerLogsSlot": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/containerlogs",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_GetWebSiteContainerLogsSlot",
		Resource:    "Microsoft.Web",
	},
	"WebApps_GetContainerLogsZipSlot": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/containerlogs/zip/download",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_GetContainerLogsZipSlot",
		Resource:    "Microsoft.Web",
	},
	"WebApps_ListContinuousWebJobsSlot": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/continuouswebjobs",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_ListContinuousWebJobsSlot",
		Resource:    "Microsoft.Web",
	},
	"WebApps_GetContinuousWebJobSlot": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/continuouswebjobs/{{.webJobName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_GetContinuousWebJobSlot",
		Resource:    "Microsoft.Web",
	},
	"WebApps_StartContinuousWebJobSlot": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/continuouswebjobs/{{.webJobName}}/start",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_StartContinuousWebJobSlot",
		Resource:    "Microsoft.Web",
	},
	"WebApps_StopContinuousWebJobSlot": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/continuouswebjobs/{{.webJobName}}/stop",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_StopContinuousWebJobSlot",
		Resource:    "Microsoft.Web",
	},
	"WebApps_ListSlotSiteDeploymentStatusesSlot": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/deploymentStatus",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_ListSlotSiteDeploymentStatusesSlot",
		Resource:    "Microsoft.Web",
	},
	"WebApps_GetSlotSiteDeploymentStatusSlot": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/deploymentStatus/{{.deploymentStatusId}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_GetSlotSiteDeploymentStatusSlot",
		Resource:    "Microsoft.Web",
	},
	"WebApps_ListDeploymentsSlot": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/deployments",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_ListDeploymentsSlot",
		Resource:    "Microsoft.Web",
	},
	"WebApps_GetDeploymentSlot": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/deployments/{{.id}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_GetDeploymentSlot",
		Resource:    "Microsoft.Web",
	},
	"WebApps_ListDeploymentLogSlot": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/deployments/{{.id}}/log",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_ListDeploymentLogSlot",
		Resource:    "Microsoft.Web",
	},
	"WebApps_DiscoverBackupSlot": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/discoverbackup",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_DiscoverBackupSlot",
		Resource:    "Microsoft.Web",
	},
	"WebApps_ListDomainOwnershipIdentifiersSlot": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/domainOwnershipIdentifiers",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_ListDomainOwnershipIdentifiersSlot",
		Resource:    "Microsoft.Web",
	},
	"WebApps_GetDomainOwnershipIdentifierSlot": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/domainOwnershipIdentifiers/{{.domainOwnershipIdentifierName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_GetDomainOwnershipIdentifierSlot",
		Resource:    "Microsoft.Web",
	},
	"WebApps_GetMSDeployStatusSlot": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/extensions/MSDeploy",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_GetMSDeployStatusSlot",
		Resource:    "Microsoft.Web",
	},
	"WebApps_GetMSDeployLogSlot": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/extensions/MSDeploy/log",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_GetMSDeployLogSlot",
		Resource:    "Microsoft.Web",
	},
	"WebApps_ListInstanceFunctionsSlot": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/functions",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_ListInstanceFunctionsSlot",
		Resource:    "Microsoft.Web",
	},
	"WebApps_GetFunctionsAdminTokenSlot": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/functions/admin/token",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_GetFunctionsAdminTokenSlot",
		Resource:    "Microsoft.Web",
	},
	"WebApps_GetInstanceFunctionSlot": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/functions/{{.functionName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_GetInstanceFunctionSlot",
		Resource:    "Microsoft.Web",
	},
	"WebApps_ListFunctionKeysSlot": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/functions/{{.functionName}}/listkeys",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_ListFunctionKeysSlot",
		Resource:    "Microsoft.Web",
	},
	"WebApps_ListFunctionSecretsSlot": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/functions/{{.functionName}}/listsecrets",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_ListFunctionSecretsSlot",
		Resource:    "Microsoft.Web",
	},
	"WebApps_ListHostKeysSlot": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/host/default/listkeys",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_ListHostKeysSlot",
		Resource:    "Microsoft.Web",
	},
	"WebApps_ListSyncStatusSlot": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/host/default/listsyncstatus",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_ListSyncStatusSlot",
		Resource:    "Microsoft.Web",
	},
	"WebApps_SyncFunctionsSlot": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/host/default/sync",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_SyncFunctionsSlot",
		Resource:    "Microsoft.Web",
	},
	"WebApps_ListHostNameBindingsSlot": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/hostNameBindings",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_ListHostNameBindingsSlot",
		Resource:    "Microsoft.Web",
	},
	"WebApps_GetHostNameBindingSlot": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/hostNameBindings/{{.hostName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_GetHostNameBindingSlot",
		Resource:    "Microsoft.Web",
	},
	"WebApps_GetHybridConnectionSlot": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/hybridConnectionNamespaces/{{.namespaceName}}/relays/{{.relayName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_GetHybridConnectionSlot",
		Resource:    "Microsoft.Web",
	},
	"WebApps_ListHybridConnectionsSlot": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/hybridConnectionRelays",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_ListHybridConnectionsSlot",
		Resource:    "Microsoft.Web",
	},
	"WebApps_ListRelayServiceConnectionsSlot": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/hybridconnection",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_ListRelayServiceConnectionsSlot",
		Resource:    "Microsoft.Web",
	},
	"WebApps_GetRelayServiceConnectionSlot": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/hybridconnection/{{.entityName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_GetRelayServiceConnectionSlot",
		Resource:    "Microsoft.Web",
	},
	"WebApps_ListInstanceIdentifiersSlot": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/instances",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_ListInstanceIdentifiersSlot",
		Resource:    "Microsoft.Web",
	},
	"WebApps_GetInstanceInfoSlot": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/instances/{{.instanceId}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_GetInstanceInfoSlot",
		Resource:    "Microsoft.Web",
	},
	"WebApps_GetInstanceMsDeployStatusSlot": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/instances/{{.instanceId}}/extensions/MSDeploy",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_GetInstanceMsDeployStatusSlot",
		Resource:    "Microsoft.Web",
	},
	"WebApps_GetInstanceMSDeployLogSlot": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/instances/{{.instanceId}}/extensions/MSDeploy/log",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_GetInstanceMSDeployLogSlot",
		Resource:    "Microsoft.Web",
	},
	"WebApps_ListInstanceProcessesSlot": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/instances/{{.instanceId}}/processes",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_ListInstanceProcessesSlot",
		Resource:    "Microsoft.Web",
	},
	"WebApps_GetInstanceProcessSlot": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/instances/{{.instanceId}}/processes/{{.processId}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_GetInstanceProcessSlot",
		Resource:    "Microsoft.Web",
	},
	"WebApps_GetInstanceProcessDumpSlot": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/instances/{{.instanceId}}/processes/{{.processId}}/dump",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_GetInstanceProcessDumpSlot",
		Resource:    "Microsoft.Web",
	},
	"WebApps_ListInstanceProcessModulesSlot": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/instances/{{.instanceId}}/processes/{{.processId}}/modules",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_ListInstanceProcessModulesSlot",
		Resource:    "Microsoft.Web",
	},
	"WebApps_GetInstanceProcessModuleSlot": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/instances/{{.instanceId}}/processes/{{.processId}}/modules/{{.baseAddress}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_GetInstanceProcessModuleSlot",
		Resource:    "Microsoft.Web",
	},
	"WebApps_ListInstanceProcessThreadsSlot": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/instances/{{.instanceId}}/processes/{{.processId}}/threads",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_ListInstanceProcessThreadsSlot",
		Resource:    "Microsoft.Web",
	},
	"WebApps_IsCloneableSlot": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/iscloneable",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_IsCloneableSlot",
		Resource:    "Microsoft.Web",
	},
	"WebApps_ListSiteBackupsSlot": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/listbackups",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_ListSiteBackupsSlot",
		Resource:    "Microsoft.Web",
	},
	"WebApps_ListSyncFunctionTriggersSlot": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/listsyncfunctiontriggerstatus",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_ListSyncFunctionTriggersSlot",
		Resource:    "Microsoft.Web",
	},
	"WebApps_GetMigrateMySqlStatusSlot": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/migratemysql/status",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_GetMigrateMySqlStatusSlot",
		Resource:    "Microsoft.Web",
	},
	"WebApps_GetSwiftVirtualNetworkConnectionSlot": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/networkConfig/virtualNetwork",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_GetSwiftVirtualNetworkConnectionSlot",
		Resource:    "Microsoft.Web",
	},
	"WebApps_ListNetworkFeaturesSlot": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/networkFeatures/{{.view}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_ListNetworkFeaturesSlot",
		Resource:    "Microsoft.Web",
	},
	"WebApps_GetNetworkTraceOperationSlot": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/networkTrace/operationresults/{{.operationId}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_GetNetworkTraceOperationSlot",
		Resource:    "Microsoft.Web",
	},
	"WebApps_StartWebSiteNetworkTraceSlot": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/networkTrace/start",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_StartWebSiteNetworkTraceSlot",
		Resource:    "Microsoft.Web",
	},
	"WebApps_StartWebSiteNetworkTraceOperationSlot": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/networkTrace/startOperation",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_StartWebSiteNetworkTraceOperationSlot",
		Resource:    "Microsoft.Web",
	},
	"WebApps_StopWebSiteNetworkTraceSlot": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/networkTrace/stop",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_StopWebSiteNetworkTraceSlot",
		Resource:    "Microsoft.Web",
	},
	"WebApps_GetNetworkTracesSlot": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/networkTrace/{{.operationId}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_GetNetworkTracesSlot",
		Resource:    "Microsoft.Web",
	},
	"WebApps_GetNetworkTraceOperationSlotV2": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/networkTraces/current/operationresults/{{.operationId}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_GetNetworkTraceOperationSlotV2",
		Resource:    "Microsoft.Web",
	},
	"WebApps_GetNetworkTracesSlotV2": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/networkTraces/{{.operationId}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_GetNetworkTracesSlotV2",
		Resource:    "Microsoft.Web",
	},
	"WebApps_GenerateNewSitePublishingPasswordSlot": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/newpassword",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_GenerateNewSitePublishingPasswordSlot",
		Resource:    "Microsoft.Web",
	},
	"WebApps_ListPerfMonCountersSlot": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/perfcounters",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_ListPerfMonCountersSlot",
		Resource:    "Microsoft.Web",
	},
	"WebApps_GetSitePhpErrorLogFlagSlot": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/phplogging",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_GetSitePhpErrorLogFlagSlot",
		Resource:    "Microsoft.Web",
	},
	"WebApps_ListPremierAddOnsSlot": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/premieraddons",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_ListPremierAddOnsSlot",
		Resource:    "Microsoft.Web",
	},
	"WebApps_GetPremierAddOnSlot": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/premieraddons/{{.premierAddOnName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_GetPremierAddOnSlot",
		Resource:    "Microsoft.Web",
	},
	"WebApps_GetPrivateAccessSlot": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/privateAccess/virtualNetworks",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_GetPrivateAccessSlot",
		Resource:    "Microsoft.Web",
	},
	"WebApps_GetPrivateEndpointConnectionListSlot": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/privateEndpointConnections",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_GetPrivateEndpointConnectionListSlot",
		Resource:    "Microsoft.Web",
	},
	"WebApps_GetPrivateEndpointConnectionSlot": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/privateEndpointConnections/{{.privateEndpointConnectionName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_GetPrivateEndpointConnectionSlot",
		Resource:    "Microsoft.Web",
	},
	"WebApps_GetPrivateLinkResourcesSlot": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/privateLinkResources",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_GetPrivateLinkResourcesSlot",
		Resource:    "Microsoft.Web",
	},
	"WebApps_ListProcessesSlot": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/processes",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_ListProcessesSlot",
		Resource:    "Microsoft.Web",
	},
	"WebApps_GetProcessSlot": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/processes/{{.processId}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_GetProcessSlot",
		Resource:    "Microsoft.Web",
	},
	"WebApps_GetProcessDumpSlot": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/processes/{{.processId}}/dump",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_GetProcessDumpSlot",
		Resource:    "Microsoft.Web",
	},
	"WebApps_ListProcessModulesSlot": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/processes/{{.processId}}/modules",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_ListProcessModulesSlot",
		Resource:    "Microsoft.Web",
	},
	"WebApps_GetProcessModuleSlot": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/processes/{{.processId}}/modules/{{.baseAddress}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_GetProcessModuleSlot",
		Resource:    "Microsoft.Web",
	},
	"WebApps_ListProcessThreadsSlot": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/processes/{{.processId}}/threads",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_ListProcessThreadsSlot",
		Resource:    "Microsoft.Web",
	},
	"WebApps_ListPublicCertificatesSlot": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/publicCertificates",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_ListPublicCertificatesSlot",
		Resource:    "Microsoft.Web",
	},
	"WebApps_GetPublicCertificateSlot": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/publicCertificates/{{.publicCertificateName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_GetPublicCertificateSlot",
		Resource:    "Microsoft.Web",
	},
	"WebApps_ListPublishingProfileXmlWithSecretsSlot": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/publishxml",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_ListPublishingProfileXmlWithSecretsSlot",
		Resource:    "Microsoft.Web",
	},
	"WebApps_ResetSlotConfigurationSlot": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/resetSlotConfig",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_ResetSlotConfigurationSlot",
		Resource:    "Microsoft.Web",
	},
	"WebApps_RestartSlot": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/restart",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_RestartSlot",
		Resource:    "Microsoft.Web",
	},
	"WebApps_RestoreFromBackupBlobSlot": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/restoreFromBackupBlob",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_RestoreFromBackupBlobSlot",
		Resource:    "Microsoft.Web",
	},
	"WebApps_RestoreFromDeletedAppSlot": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/restoreFromDeletedApp",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_RestoreFromDeletedAppSlot",
		Resource:    "Microsoft.Web",
	},
	"WebApps_RestoreSnapshotSlot": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/restoreSnapshot",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_RestoreSnapshotSlot",
		Resource:    "Microsoft.Web",
	},
	"WebApps_ListSiteExtensionsSlot": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/siteextensions",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_ListSiteExtensionsSlot",
		Resource:    "Microsoft.Web",
	},
	"WebApps_GetSiteExtensionSlot": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/siteextensions/{{.siteExtensionId}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_GetSiteExtensionSlot",
		Resource:    "Microsoft.Web",
	},
	"WebApps_ListSlotDifferencesSlot": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/slotsdiffs",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_ListSlotDifferencesSlot",
		Resource:    "Microsoft.Web",
	},
	"WebApps_SwapSlotSlot": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/slotsswap",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_SwapSlotSlot",
		Resource:    "Microsoft.Web",
	},
	"WebApps_ListSnapshotsSlot": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/snapshots",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_ListSnapshotsSlot",
		Resource:    "Microsoft.Web",
	},
	"WebApps_ListSnapshotsFromDRSecondarySlot": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/snapshotsdr",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_ListSnapshotsFromDRSecondarySlot",
		Resource:    "Microsoft.Web",
	},
	"WebApps_GetSourceControlSlot": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/sourcecontrols/web",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_GetSourceControlSlot",
		Resource:    "Microsoft.Web",
	},
	"WebApps_StartSlot": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/start",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_StartSlot",
		Resource:    "Microsoft.Web",
	},
	"WebApps_StartNetworkTraceSlot": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/startNetworkTrace",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_StartNetworkTraceSlot",
		Resource:    "Microsoft.Web",
	},
	"WebApps_StopSlot": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/stop",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_StopSlot",
		Resource:    "Microsoft.Web",
	},
	"WebApps_StopNetworkTraceSlot": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/stopNetworkTrace",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_StopNetworkTraceSlot",
		Resource:    "Microsoft.Web",
	},
	"WebApps_SyncRepositorySlot": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/sync",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_SyncRepositorySlot",
		Resource:    "Microsoft.Web",
	},
	"WebApps_SyncFunctionTriggersSlot": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/syncfunctiontriggers",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_SyncFunctionTriggersSlot",
		Resource:    "Microsoft.Web",
	},
	"WebApps_ListTriggeredWebJobsSlot": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/triggeredwebjobs",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_ListTriggeredWebJobsSlot",
		Resource:    "Microsoft.Web",
	},
	"WebApps_GetTriggeredWebJobSlot": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/triggeredwebjobs/{{.webJobName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_GetTriggeredWebJobSlot",
		Resource:    "Microsoft.Web",
	},
	"WebApps_ListTriggeredWebJobHistorySlot": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/triggeredwebjobs/{{.webJobName}}/history",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_ListTriggeredWebJobHistorySlot",
		Resource:    "Microsoft.Web",
	},
	"WebApps_GetTriggeredWebJobHistorySlot": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/triggeredwebjobs/{{.webJobName}}/history/{{.id}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_GetTriggeredWebJobHistorySlot",
		Resource:    "Microsoft.Web",
	},
	"WebApps_RunTriggeredWebJobSlot": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/triggeredwebjobs/{{.webJobName}}/run",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_RunTriggeredWebJobSlot",
		Resource:    "Microsoft.Web",
	},
	"WebApps_ListUsagesSlot": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/usages",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_ListUsagesSlot",
		Resource:    "Microsoft.Web",
	},
	"WebApps_ListVnetConnectionsSlot": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/virtualNetworkConnections",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_ListVnetConnectionsSlot",
		Resource:    "Microsoft.Web",
	},
	"WebApps_GetVnetConnectionSlot": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/virtualNetworkConnections/{{.vnetName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_GetVnetConnectionSlot",
		Resource:    "Microsoft.Web",
	},
	"WebApps_GetVnetConnectionGatewaySlot": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/virtualNetworkConnections/{{.vnetName}}/gateways/{{.gatewayName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_GetVnetConnectionGatewaySlot",
		Resource:    "Microsoft.Web",
	},
	"WebApps_ListWebJobsSlot": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/webjobs",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_ListWebJobsSlot",
		Resource:    "Microsoft.Web",
	},
	"WebApps_GetWebJobSlot": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/webjobs/{{.webJobName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_GetWebJobSlot",
		Resource:    "Microsoft.Web",
	},
	"WebApps_ListSlotDifferencesFromProduction": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slotsdiffs",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_ListSlotDifferencesFromProduction",
		Resource:    "Microsoft.Web",
	},
	"WebApps_SwapSlotWithProduction": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slotsswap",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_SwapSlotWithProduction",
		Resource:    "Microsoft.Web",
	},
	"WebApps_ListSnapshots": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/snapshots",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_ListSnapshots",
		Resource:    "Microsoft.Web",
	},
	"WebApps_ListSnapshotsFromDRSecondary": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/snapshotsdr",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_ListSnapshotsFromDRSecondary",
		Resource:    "Microsoft.Web",
	},
	"WebApps_GetSourceControl": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/sourcecontrols/web",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_GetSourceControl",
		Resource:    "Microsoft.Web",
	},
	"WebApps_Start": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/start",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_Start",
		Resource:    "Microsoft.Web",
	},
	"WebApps_StartNetworkTrace": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/startNetworkTrace",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_StartNetworkTrace",
		Resource:    "Microsoft.Web",
	},
	"WebApps_Stop": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/stop",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_Stop",
		Resource:    "Microsoft.Web",
	},
	"WebApps_StopNetworkTrace": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/stopNetworkTrace",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_StopNetworkTrace",
		Resource:    "Microsoft.Web",
	},
	"WebApps_SyncRepository": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/sync",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_SyncRepository",
		Resource:    "Microsoft.Web",
	},
	"WebApps_SyncFunctionTriggers": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/syncfunctiontriggers",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_SyncFunctionTriggers",
		Resource:    "Microsoft.Web",
	},
	"WebApps_ListTriggeredWebJobs": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/triggeredwebjobs",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_ListTriggeredWebJobs",
		Resource:    "Microsoft.Web",
	},
	"WebApps_GetTriggeredWebJob": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/triggeredwebjobs/{{.webJobName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_GetTriggeredWebJob",
		Resource:    "Microsoft.Web",
	},
	"WebApps_ListTriggeredWebJobHistory": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/triggeredwebjobs/{{.webJobName}}/history",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_ListTriggeredWebJobHistory",
		Resource:    "Microsoft.Web",
	},
	"WebApps_GetTriggeredWebJobHistory": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/triggeredwebjobs/{{.webJobName}}/history/{{.id}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_GetTriggeredWebJobHistory",
		Resource:    "Microsoft.Web",
	},
	"WebApps_RunTriggeredWebJob": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/triggeredwebjobs/{{.webJobName}}/run",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_RunTriggeredWebJob",
		Resource:    "Microsoft.Web",
	},
	"WebApps_ListUsages": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/usages",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_ListUsages",
		Resource:    "Microsoft.Web",
	},
	"WebApps_ListVnetConnections": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/virtualNetworkConnections",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_ListVnetConnections",
		Resource:    "Microsoft.Web",
	},
	"WebApps_GetVnetConnection": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/virtualNetworkConnections/{{.vnetName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_GetVnetConnection",
		Resource:    "Microsoft.Web",
	},
	"WebApps_GetVnetConnectionGateway": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/virtualNetworkConnections/{{.vnetName}}/gateways/{{.gatewayName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_GetVnetConnectionGateway",
		Resource:    "Microsoft.Web",
	},
	"WebApps_ListWebJobs": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/webjobs",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_ListWebJobs",
		Resource:    "Microsoft.Web",
	},
	"WebApps_GetWebJob": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/webjobs/{{.webJobName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WebApps_GetWebJob",
		Resource:    "Microsoft.Web",
	},
	"Workflows_RegenerateAccessKey": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/hostruntime/runtime/webhooks/workflow/api/management/workflows/{{.workflowName}}/regenerateAccessKey",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "Workflows_RegenerateAccessKey",
		Resource:    "Microsoft.Web",
	},
	"WorkflowRuns_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/hostruntime/runtime/webhooks/workflow/api/management/workflows/{{.workflowName}}/runs",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WorkflowRuns_List",
		Resource:    "Microsoft.Web",
	},
	"WorkflowRuns_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/hostruntime/runtime/webhooks/workflow/api/management/workflows/{{.workflowName}}/runs/{{.runName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WorkflowRuns_Get",
		Resource:    "Microsoft.Web",
	},
	"WorkflowRunActions_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/hostruntime/runtime/webhooks/workflow/api/management/workflows/{{.workflowName}}/runs/{{.runName}}/actions",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WorkflowRunActions_List",
		Resource:    "Microsoft.Web",
	},
	"WorkflowRunActions_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/hostruntime/runtime/webhooks/workflow/api/management/workflows/{{.workflowName}}/runs/{{.runName}}/actions/{{.actionName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WorkflowRunActions_Get",
		Resource:    "Microsoft.Web",
	},
	"WorkflowRunActions_ListExpressionTraces": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/hostruntime/runtime/webhooks/workflow/api/management/workflows/{{.workflowName}}/runs/{{.runName}}/actions/{{.actionName}}/listExpressionTraces",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WorkflowRunActions_ListExpressionTraces",
		Resource:    "Microsoft.Web",
	},
	"WorkflowRunActionRepetitions_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/hostruntime/runtime/webhooks/workflow/api/management/workflows/{{.workflowName}}/runs/{{.runName}}/actions/{{.actionName}}/repetitions",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WorkflowRunActionRepetitions_List",
		Resource:    "Microsoft.Web",
	},
	"WorkflowRunActionRepetitions_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/hostruntime/runtime/webhooks/workflow/api/management/workflows/{{.workflowName}}/runs/{{.runName}}/actions/{{.actionName}}/repetitions/{{.repetitionName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WorkflowRunActionRepetitions_Get",
		Resource:    "Microsoft.Web",
	},
	"WorkflowRunActionRepetitions_ListExpressionTraces": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/hostruntime/runtime/webhooks/workflow/api/management/workflows/{{.workflowName}}/runs/{{.runName}}/actions/{{.actionName}}/repetitions/{{.repetitionName}}/listExpressionTraces",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WorkflowRunActionRepetitions_ListExpressionTraces",
		Resource:    "Microsoft.Web",
	},
	"WorkflowRunActionRepetitionsRequestHistories_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/hostruntime/runtime/webhooks/workflow/api/management/workflows/{{.workflowName}}/runs/{{.runName}}/actions/{{.actionName}}/repetitions/{{.repetitionName}}/requestHistories",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WorkflowRunActionRepetitionsRequestHistories_List",
		Resource:    "Microsoft.Web",
	},
	"WorkflowRunActionRepetitionsRequestHistories_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/hostruntime/runtime/webhooks/workflow/api/management/workflows/{{.workflowName}}/runs/{{.runName}}/actions/{{.actionName}}/repetitions/{{.repetitionName}}/requestHistories/{{.requestHistoryName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WorkflowRunActionRepetitionsRequestHistories_Get",
		Resource:    "Microsoft.Web",
	},
	"WorkflowRunActionScopeRepetitions_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/hostruntime/runtime/webhooks/workflow/api/management/workflows/{{.workflowName}}/runs/{{.runName}}/actions/{{.actionName}}/scopeRepetitions",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WorkflowRunActionScopeRepetitions_List",
		Resource:    "Microsoft.Web",
	},
	"WorkflowRunActionScopeRepetitions_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/hostruntime/runtime/webhooks/workflow/api/management/workflows/{{.workflowName}}/runs/{{.runName}}/actions/{{.actionName}}/scopeRepetitions/{{.repetitionName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WorkflowRunActionScopeRepetitions_Get",
		Resource:    "Microsoft.Web",
	},
	"WorkflowRuns_Cancel": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/hostruntime/runtime/webhooks/workflow/api/management/workflows/{{.workflowName}}/runs/{{.runName}}/cancel",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WorkflowRuns_Cancel",
		Resource:    "Microsoft.Web",
	},
	"WorkflowTriggers_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/hostruntime/runtime/webhooks/workflow/api/management/workflows/{{.workflowName}}/triggers",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WorkflowTriggers_List",
		Resource:    "Microsoft.Web",
	},
	"WorkflowTriggers_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/hostruntime/runtime/webhooks/workflow/api/management/workflows/{{.workflowName}}/triggers/{{.triggerName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WorkflowTriggers_Get",
		Resource:    "Microsoft.Web",
	},
	"WorkflowTriggerHistories_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/hostruntime/runtime/webhooks/workflow/api/management/workflows/{{.workflowName}}/triggers/{{.triggerName}}/histories",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WorkflowTriggerHistories_List",
		Resource:    "Microsoft.Web",
	},
	"WorkflowTriggerHistories_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/hostruntime/runtime/webhooks/workflow/api/management/workflows/{{.workflowName}}/triggers/{{.triggerName}}/histories/{{.historyName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WorkflowTriggerHistories_Get",
		Resource:    "Microsoft.Web",
	},
	"WorkflowTriggerHistories_Resubmit": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/hostruntime/runtime/webhooks/workflow/api/management/workflows/{{.workflowName}}/triggers/{{.triggerName}}/histories/{{.historyName}}/resubmit",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WorkflowTriggerHistories_Resubmit",
		Resource:    "Microsoft.Web",
	},
	"WorkflowTriggers_ListCallbackUrl": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/hostruntime/runtime/webhooks/workflow/api/management/workflows/{{.workflowName}}/triggers/{{.triggerName}}/listCallbackUrl",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WorkflowTriggers_ListCallbackUrl",
		Resource:    "Microsoft.Web",
	},
	"WorkflowTriggers_Run": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/hostruntime/runtime/webhooks/workflow/api/management/workflows/{{.workflowName}}/triggers/{{.triggerName}}/run",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WorkflowTriggers_Run",
		Resource:    "Microsoft.Web",
	},
	"WorkflowTriggers_GetSchemaJson": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/hostruntime/runtime/webhooks/workflow/api/management/workflows/{{.workflowName}}/triggers/{{.triggerName}}/schemas/json",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WorkflowTriggers_GetSchemaJson",
		Resource:    "Microsoft.Web",
	},
	"Workflows_Validate": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/hostruntime/runtime/webhooks/workflow/api/management/workflows/{{.workflowName}}/validate",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "Workflows_Validate",
		Resource:    "Microsoft.Web",
	},
	"WorkflowVersions_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/hostruntime/runtime/webhooks/workflow/api/management/workflows/{{.workflowName}}/versions",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WorkflowVersions_List",
		Resource:    "Microsoft.Web",
	},
	"WorkflowVersions_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/hostruntime/runtime/webhooks/workflow/api/management/workflows/{{.workflowName}}/versions/{{.versionId}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "WorkflowVersions_Get",
		Resource:    "Microsoft.Web",
	},
}
