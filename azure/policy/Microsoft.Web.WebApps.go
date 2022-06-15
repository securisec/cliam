package policy

var Microsoft_Web_WebApps = []Policy{
	{
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Web/sites",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_List",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_ListByResourceGroup",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_Get",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/analyzeCustomHostname",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_AnalyzeCustomHostname",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/applySlotConfig",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_ApplySlotConfigToProduction",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/backup",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_Backup",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/backups",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_ListBackups",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/backups/{{.backupId}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_GetBackupStatus",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/backups/{{.backupId}}/list",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_ListBackupStatusSecrets",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/backups/{{.backupId}}/restore",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_Restore",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/basicPublishingCredentialsPolicies",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_ListBasicPublishingCredentialsPolicies",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/basicPublishingCredentialsPolicies/ftp",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_GetFtpAllowed",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/basicPublishingCredentialsPolicies/scm",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_GetScmAllowed",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/config",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_ListConfigurations",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/config/appsettings/list",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_ListApplicationSettings",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/config/authsettings/list",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_GetAuthSettings",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/config/authsettingsV2",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_GetAuthSettingsV2WithoutSecrets",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/config/authsettingsV2/list",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_GetAuthSettingsV2",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/config/azurestorageaccounts/list",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_ListAzureStorageAccounts",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/config/backup/list",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_GetBackupConfiguration",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/config/configreferences/appsettings",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_GetAppSettingsKeyVaultReferences",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/config/configreferences/appsettings/{{.appSettingKey}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_GetAppSettingKeyVaultReference",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/config/configreferences/connectionstrings",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_GetSiteConnectionStringKeyVaultReferences",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/config/configreferences/connectionstrings/{{.connectionStringKey}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_GetSiteConnectionStringKeyVaultReference",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/config/connectionstrings/list",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_ListConnectionStrings",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/config/logs",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_GetDiagnosticLogsConfiguration",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/config/metadata/list",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_ListMetadata",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/config/publishingcredentials/list",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_ListPublishingCredentials",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/config/pushsettings/list",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_ListSitePushSettings",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/config/slotConfigNames",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_ListSlotConfigurationNames",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/config/web",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_GetConfiguration",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/config/web/snapshots",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_ListConfigurationSnapshotInfo",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/config/web/snapshots/{{.snapshotId}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_GetConfigurationSnapshot",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/config/web/snapshots/{{.snapshotId}}/recover",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_RecoverSiteConfigurationSnapshot",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/containerlogs",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_GetWebSiteContainerLogs",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/containerlogs/zip/download",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_GetContainerLogsZip",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/continuouswebjobs",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_ListContinuousWebJobs",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/continuouswebjobs/{{.webJobName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_GetContinuousWebJob",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/continuouswebjobs/{{.webJobName}}/start",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_StartContinuousWebJob",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/continuouswebjobs/{{.webJobName}}/stop",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_StopContinuousWebJob",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/deployments",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_ListDeployments",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/deployments/{{.id}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_GetDeployment",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/deployments/{{.id}}/log",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_ListDeploymentLog",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/discoverbackup",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_DiscoverBackup",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/domainOwnershipIdentifiers",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_ListDomainOwnershipIdentifiers",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/domainOwnershipIdentifiers/{{.domainOwnershipIdentifierName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_GetDomainOwnershipIdentifier",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/extensions/MSDeploy",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_GetMSDeployStatus",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/extensions/MSDeploy/log",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_GetMSDeployLog",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/extensions/onedeploy",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_GetOneDeployStatus",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/functions",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_ListFunctions",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/functions/admin/token",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_GetFunctionsAdminToken",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/functions/{{.functionName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_GetFunction",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/functions/{{.functionName}}/listkeys",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_ListFunctionKeys",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/functions/{{.functionName}}/listsecrets",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_ListFunctionSecrets",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/host/default/listkeys",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_ListHostKeys",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/host/default/listsyncstatus",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_ListSyncStatus",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/host/default/sync",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_SyncFunctions",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/hostNameBindings",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_ListHostNameBindings",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/hostNameBindings/{{.hostName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_GetHostNameBinding",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/hybridConnectionNamespaces/{{.namespaceName}}/relays/{{.relayName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_GetHybridConnection",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/hybridConnectionRelays",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_ListHybridConnections",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/hybridconnection",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_ListRelayServiceConnections",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/hybridconnection/{{.entityName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_GetRelayServiceConnection",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/instances",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_ListInstanceIdentifiers",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/instances/{{.instanceId}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_GetInstanceInfo",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/instances/{{.instanceId}}/extensions/MSDeploy",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_GetInstanceMsDeployStatus",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/instances/{{.instanceId}}/extensions/MSDeploy/log",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_GetInstanceMSDeployLog",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/instances/{{.instanceId}}/processes",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_ListInstanceProcesses",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/instances/{{.instanceId}}/processes/{{.processId}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_GetInstanceProcess",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/instances/{{.instanceId}}/processes/{{.processId}}/dump",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_GetInstanceProcessDump",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/instances/{{.instanceId}}/processes/{{.processId}}/modules",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_ListInstanceProcessModules",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/instances/{{.instanceId}}/processes/{{.processId}}/modules/{{.baseAddress}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_GetInstanceProcessModule",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/instances/{{.instanceId}}/processes/{{.processId}}/threads",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_ListInstanceProcessThreads",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/iscloneable",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_IsCloneable",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/listbackups",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_ListSiteBackups",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/listsyncfunctiontriggerstatus",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_ListSyncFunctionTriggers",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/migratemysql",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_MigrateMySql",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/migratemysql/status",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_GetMigrateMySqlStatus",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/networkConfig/virtualNetwork",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_GetSwiftVirtualNetworkConnection",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/networkFeatures/{{.view}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_ListNetworkFeatures",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/networkTrace/operationresults/{{.operationId}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_GetNetworkTraceOperation",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/networkTrace/start",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_StartWebSiteNetworkTrace",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/networkTrace/startOperation",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_StartWebSiteNetworkTraceOperation",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/networkTrace/stop",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_StopWebSiteNetworkTrace",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/networkTrace/{{.operationId}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_GetNetworkTraces",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/networkTraces/current/operationresults/{{.operationId}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_GetNetworkTraceOperationV2",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/networkTraces/{{.operationId}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_GetNetworkTracesV2",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/newpassword",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_GenerateNewSitePublishingPassword",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/perfcounters",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_ListPerfMonCounters",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/phplogging",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_GetSitePhpErrorLogFlag",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/premieraddons",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_ListPremierAddOns",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/premieraddons/{{.premierAddOnName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_GetPremierAddOn",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/privateAccess/virtualNetworks",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_GetPrivateAccess",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/privateEndpointConnections",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_GetPrivateEndpointConnectionList",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/privateEndpointConnections/{{.privateEndpointConnectionName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_GetPrivateEndpointConnection",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/privateLinkResources",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_GetPrivateLinkResources",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/processes",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_ListProcesses",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/processes/{{.processId}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_GetProcess",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/processes/{{.processId}}/dump",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_GetProcessDump",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/processes/{{.processId}}/modules",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_ListProcessModules",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/processes/{{.processId}}/modules/{{.baseAddress}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_GetProcessModule",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/processes/{{.processId}}/threads",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_ListProcessThreads",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/publicCertificates",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_ListPublicCertificates",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/publicCertificates/{{.publicCertificateName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_GetPublicCertificate",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/publishxml",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_ListPublishingProfileXmlWithSecrets",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/resetSlotConfig",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_ResetProductionSlotConfig",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/restart",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_Restart",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/restoreFromBackupBlob",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_RestoreFromBackupBlob",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/restoreFromDeletedApp",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_RestoreFromDeletedApp",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/restoreSnapshot",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_RestoreSnapshot",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/siteextensions",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_ListSiteExtensions",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/siteextensions/{{.siteExtensionId}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_GetSiteExtension",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_ListSlots",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_GetSlot",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/analyzeCustomHostname",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_AnalyzeCustomHostnameSlot",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/applySlotConfig",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_ApplySlotConfigurationSlot",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/backup",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_BackupSlot",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/backups",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_ListBackupsSlot",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/backups/{{.backupId}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_GetBackupStatusSlot",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/backups/{{.backupId}}/list",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_ListBackupStatusSecretsSlot",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/backups/{{.backupId}}/restore",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_RestoreSlot",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/basicPublishingCredentialsPolicies",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_ListBasicPublishingCredentialsPoliciesSlot",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/basicPublishingCredentialsPolicies/ftp",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_GetFtpAllowedSlot",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/basicPublishingCredentialsPolicies/scm",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_GetScmAllowedSlot",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/config",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_ListConfigurationsSlot",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/config/appsettings/list",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_ListApplicationSettingsSlot",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/config/authsettings/list",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_GetAuthSettingsSlot",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/config/authsettingsV2",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_GetAuthSettingsV2WithoutSecretsSlot",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/config/authsettingsV2/list",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_GetAuthSettingsV2Slot",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/config/azurestorageaccounts/list",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_ListAzureStorageAccountsSlot",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/config/backup/list",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_GetBackupConfigurationSlot",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/config/configreferences/appsettings",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_GetAppSettingsKeyVaultReferencesSlot",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/config/configreferences/appsettings/{{.appSettingKey}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_GetAppSettingKeyVaultReferenceSlot",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/config/configreferences/connectionstrings",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_GetSiteConnectionStringKeyVaultReferencesSlot",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/config/configreferences/connectionstrings/{{.connectionStringKey}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_GetSiteConnectionStringKeyVaultReferenceSlot",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/config/connectionstrings/list",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_ListConnectionStringsSlot",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/config/logs",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_GetDiagnosticLogsConfigurationSlot",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/config/metadata/list",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_ListMetadataSlot",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/config/publishingcredentials/list",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_ListPublishingCredentialsSlot",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/config/pushsettings/list",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_ListSitePushSettingsSlot",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/config/web",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_GetConfigurationSlot",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/config/web/snapshots",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_ListConfigurationSnapshotInfoSlot",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/config/web/snapshots/{{.snapshotId}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_GetConfigurationSnapshotSlot",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/config/web/snapshots/{{.snapshotId}}/recover",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_RecoverSiteConfigurationSnapshotSlot",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/containerlogs",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_GetWebSiteContainerLogsSlot",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/containerlogs/zip/download",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_GetContainerLogsZipSlot",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/continuouswebjobs",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_ListContinuousWebJobsSlot",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/continuouswebjobs/{{.webJobName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_GetContinuousWebJobSlot",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/continuouswebjobs/{{.webJobName}}/start",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_StartContinuousWebJobSlot",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/continuouswebjobs/{{.webJobName}}/stop",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_StopContinuousWebJobSlot",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/deployments",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_ListDeploymentsSlot",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/deployments/{{.id}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_GetDeploymentSlot",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/deployments/{{.id}}/log",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_ListDeploymentLogSlot",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/discoverbackup",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_DiscoverBackupSlot",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/domainOwnershipIdentifiers",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_ListDomainOwnershipIdentifiersSlot",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/domainOwnershipIdentifiers/{{.domainOwnershipIdentifierName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_GetDomainOwnershipIdentifierSlot",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/extensions/MSDeploy",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_GetMSDeployStatusSlot",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/extensions/MSDeploy/log",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_GetMSDeployLogSlot",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/functions",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_ListInstanceFunctionsSlot",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/functions/admin/token",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_GetFunctionsAdminTokenSlot",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/functions/{{.functionName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_GetInstanceFunctionSlot",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/functions/{{.functionName}}/listkeys",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_ListFunctionKeysSlot",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/functions/{{.functionName}}/listsecrets",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_ListFunctionSecretsSlot",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/host/default/listkeys",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_ListHostKeysSlot",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/host/default/listsyncstatus",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_ListSyncStatusSlot",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/host/default/sync",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_SyncFunctionsSlot",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/hostNameBindings",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_ListHostNameBindingsSlot",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/hostNameBindings/{{.hostName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_GetHostNameBindingSlot",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/hybridConnectionNamespaces/{{.namespaceName}}/relays/{{.relayName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_GetHybridConnectionSlot",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/hybridConnectionRelays",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_ListHybridConnectionsSlot",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/hybridconnection",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_ListRelayServiceConnectionsSlot",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/hybridconnection/{{.entityName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_GetRelayServiceConnectionSlot",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/instances",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_ListInstanceIdentifiersSlot",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/instances/{{.instanceId}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_GetInstanceInfoSlot",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/instances/{{.instanceId}}/extensions/MSDeploy",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_GetInstanceMsDeployStatusSlot",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/instances/{{.instanceId}}/extensions/MSDeploy/log",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_GetInstanceMSDeployLogSlot",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/instances/{{.instanceId}}/processes",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_ListInstanceProcessesSlot",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/instances/{{.instanceId}}/processes/{{.processId}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_GetInstanceProcessSlot",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/instances/{{.instanceId}}/processes/{{.processId}}/dump",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_GetInstanceProcessDumpSlot",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/instances/{{.instanceId}}/processes/{{.processId}}/modules",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_ListInstanceProcessModulesSlot",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/instances/{{.instanceId}}/processes/{{.processId}}/modules/{{.baseAddress}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_GetInstanceProcessModuleSlot",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/instances/{{.instanceId}}/processes/{{.processId}}/threads",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_ListInstanceProcessThreadsSlot",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/iscloneable",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_IsCloneableSlot",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/listbackups",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_ListSiteBackupsSlot",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/listsyncfunctiontriggerstatus",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_ListSyncFunctionTriggersSlot",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/migratemysql/status",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_GetMigrateMySqlStatusSlot",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/networkConfig/virtualNetwork",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_GetSwiftVirtualNetworkConnectionSlot",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/networkFeatures/{{.view}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_ListNetworkFeaturesSlot",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/networkTrace/operationresults/{{.operationId}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_GetNetworkTraceOperationSlot",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/networkTrace/start",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_StartWebSiteNetworkTraceSlot",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/networkTrace/startOperation",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_StartWebSiteNetworkTraceOperationSlot",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/networkTrace/stop",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_StopWebSiteNetworkTraceSlot",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/networkTrace/{{.operationId}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_GetNetworkTracesSlot",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/networkTraces/current/operationresults/{{.operationId}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_GetNetworkTraceOperationSlotV2",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/networkTraces/{{.operationId}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_GetNetworkTracesSlotV2",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/newpassword",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_GenerateNewSitePublishingPasswordSlot",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/perfcounters",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_ListPerfMonCountersSlot",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/phplogging",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_GetSitePhpErrorLogFlagSlot",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/premieraddons",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_ListPremierAddOnsSlot",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/premieraddons/{{.premierAddOnName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_GetPremierAddOnSlot",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/privateAccess/virtualNetworks",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_GetPrivateAccessSlot",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/privateEndpointConnections",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_GetPrivateEndpointConnectionListSlot",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/privateEndpointConnections/{{.privateEndpointConnectionName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_GetPrivateEndpointConnectionSlot",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/privateLinkResources",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_GetPrivateLinkResourcesSlot",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/processes",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_ListProcessesSlot",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/processes/{{.processId}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_GetProcessSlot",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/processes/{{.processId}}/dump",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_GetProcessDumpSlot",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/processes/{{.processId}}/modules",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_ListProcessModulesSlot",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/processes/{{.processId}}/modules/{{.baseAddress}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_GetProcessModuleSlot",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/processes/{{.processId}}/threads",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_ListProcessThreadsSlot",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/publicCertificates",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_ListPublicCertificatesSlot",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/publicCertificates/{{.publicCertificateName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_GetPublicCertificateSlot",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/publishxml",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_ListPublishingProfileXmlWithSecretsSlot",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/resetSlotConfig",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_ResetSlotConfigurationSlot",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/restart",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_RestartSlot",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/restoreFromBackupBlob",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_RestoreFromBackupBlobSlot",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/restoreFromDeletedApp",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_RestoreFromDeletedAppSlot",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/restoreSnapshot",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_RestoreSnapshotSlot",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/siteextensions",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_ListSiteExtensionsSlot",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/siteextensions/{{.siteExtensionId}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_GetSiteExtensionSlot",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/slotsdiffs",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_ListSlotDifferencesSlot",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/slotsswap",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_SwapSlotSlot",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/snapshots",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_ListSnapshotsSlot",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/snapshotsdr",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_ListSnapshotsFromDRSecondarySlot",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/sourcecontrols/web",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_GetSourceControlSlot",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/start",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_StartSlot",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/startNetworkTrace",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_StartNetworkTraceSlot",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/stop",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_StopSlot",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/stopNetworkTrace",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_StopNetworkTraceSlot",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/sync",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_SyncRepositorySlot",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/syncfunctiontriggers",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_SyncFunctionTriggersSlot",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/triggeredwebjobs",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_ListTriggeredWebJobsSlot",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/triggeredwebjobs/{{.webJobName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_GetTriggeredWebJobSlot",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/triggeredwebjobs/{{.webJobName}}/history",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_ListTriggeredWebJobHistorySlot",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/triggeredwebjobs/{{.webJobName}}/history/{{.id}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_GetTriggeredWebJobHistorySlot",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/triggeredwebjobs/{{.webJobName}}/run",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_RunTriggeredWebJobSlot",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/usages",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_ListUsagesSlot",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/virtualNetworkConnections",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_ListVnetConnectionsSlot",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/virtualNetworkConnections/{{.vnetName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_GetVnetConnectionSlot",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/virtualNetworkConnections/{{.vnetName}}/gateways/{{.gatewayName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_GetVnetConnectionGatewaySlot",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/webjobs",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_ListWebJobsSlot",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slots/{{.slot}}/webjobs/{{.webJobName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_GetWebJobSlot",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slotsdiffs",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_ListSlotDifferencesFromProduction",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/slotsswap",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_SwapSlotWithProduction",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/snapshots",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_ListSnapshots",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/snapshotsdr",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_ListSnapshotsFromDRSecondary",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/sourcecontrols/web",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_GetSourceControl",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/start",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_Start",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/startNetworkTrace",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_StartNetworkTrace",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/stop",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_Stop",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/stopNetworkTrace",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_StopNetworkTrace",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/sync",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_SyncRepository",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/syncfunctiontriggers",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_SyncFunctionTriggers",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/triggeredwebjobs",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_ListTriggeredWebJobs",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/triggeredwebjobs/{{.webJobName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_GetTriggeredWebJob",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/triggeredwebjobs/{{.webJobName}}/history",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_ListTriggeredWebJobHistory",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/triggeredwebjobs/{{.webJobName}}/history/{{.id}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_GetTriggeredWebJobHistory",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/triggeredwebjobs/{{.webJobName}}/run",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_RunTriggeredWebJob",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/usages",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_ListUsages",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/virtualNetworkConnections",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_ListVnetConnections",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/virtualNetworkConnections/{{.vnetName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_GetVnetConnection",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/virtualNetworkConnections/{{.vnetName}}/gateways/{{.gatewayName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_GetVnetConnectionGateway",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/webjobs",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_ListWebJobs",
		Resource:    "Microsoft.Web",
	},
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Web/sites/{{.name}}/webjobs/{{.webJobName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "WebApps_GetWebJob",
		Resource:    "Microsoft.Web",
	},
}
