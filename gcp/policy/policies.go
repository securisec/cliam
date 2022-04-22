package policy

type Service struct {
	Actions []string `json:"actions"`
	Method  string   `json:"method"`
}

var Resources = map[string][]Service{
	AccessApproval: {
		{Method: accessapprovalRequests, Actions: []string{"approve", "dismiss", "get", "list"}},
		{Method: accessapprovalSettings, Actions: []string{"get", "update"}},
	},

	AndroidManagement: {
		{Method: androidmanagementEnterprises, Actions: []string{"manage"}},
	},

	AppEngine: {
		{Method: appengineApplications, Actions: []string{"create", "get", "update"}},
		{Method: appengineInstances, Actions: []string{"delete", "get", "list"}},
		{Method: appengineMemcache, Actions: []string{"addKey", "flush", "get", "getKey", "list", "update"}},
		{Method: appengineOperations, Actions: []string{"get", "list"}},
		{Method: appengineRuntimes, Actions: []string{"actAsAdmin"}},
		{Method: appengineServices, Actions: []string{"delete", "get", "list", "update"}},
		{Method: appengineVersions, Actions: []string{"create", "delete", "get", "getFileContents", "list", "update"}},
	},

	AutoMl: {
		{Method: automlAnnotationSpecs, Actions: []string{"create", "delete", "get", "list", "update"}},
		{Method: automlAnnotations, Actions: []string{"approve", "create", "list", "manipulate", "reject"}},
		{Method: automlColumnSpecs, Actions: []string{"get", "list", "update"}},
		{Method: automlDatasets, Actions: []string{"create", "delete", "export", "get", "getIamPolicy", "import", "list", "setIamPolicy", "update"}},
		{Method: automlExamples, Actions: []string{"delete", "get", "list"}},
		{Method: automlHumanAnnotationTasks, Actions: []string{"create", "delete", "get", "list"}},
		{Method: automlLocations, Actions: []string{"get", "getIamPolicy", "list", "setIamPolicy"}},
		{Method: automlModelEvaluations, Actions: []string{"create", "get", "list"}},
		{Method: automlModels, Actions: []string{"create", "delete", "deploy", "export", "get", "getIamPolicy", "list", "predict", "setIamPolicy", "undeploy"}},
		{Method: automlOperations, Actions: []string{"cancel", "delete", "get", "list"}},
		{Method: automlTableSpecs, Actions: []string{"get", "list", "update"}},
		{Method: automlrecommendationsApiKeys, Actions: []string{"create", "delete", "list"}},
		{Method: automlrecommendationsCatalogItems, Actions: []string{"create", "delete", "get", "list", "update"}},
		{Method: automlrecommendationsCatalogs, Actions: []string{"getStats", "list"}},
		{Method: automlrecommendationsEventStores, Actions: []string{"getStats"}},
		{Method: automlrecommendationsEvents, Actions: []string{"create", "list", "purge"}},
		{Method: automlrecommendationsPlacements, Actions: []string{"getStats", "list"}},
		{Method: automlrecommendationsRecommendations, Actions: []string{"list"}},
	},

	BigQuery: {
		{Method: bigqueryConfig, Actions: []string{"get", "update"}},
		{Method: bigqueryConnections, Actions: []string{"create", "delete", "get", "getIamPolicy", "list", "setIamPolicy", "update", "use"}},
		{Method: bigqueryDatasets, Actions: []string{"create", "delete", "get", "getIamPolicy", "setIamPolicy", "update", "updateTag"}},
		{Method: bigqueryJobs, Actions: []string{"create", "get", "list", "listAll", "update"}},
		{Method: bigqueryModels, Actions: []string{"create", "delete", "getData", "getMetadata", "list", "updateData", "updateMetadata", "updateTag"}},
		{Method: bigqueryReadsessions, Actions: []string{"create"}},
		{Method: bigqueryRoutines, Actions: []string{"create", "delete", "get", "list", "update"}},
		{Method: bigquerySavedqueries, Actions: []string{"create", "delete", "get", "list", "update"}},
		{Method: bigqueryTransfers, Actions: []string{"get", "update"}},
	},

	BigTable: {
		{Method: bigtableAppProfiles, Actions: []string{"create", "delete", "get", "list", "update"}},
		{Method: bigtableClusters, Actions: []string{"create", "delete", "get", "list", "update"}},
		{Method: bigtableInstances, Actions: []string{"create", "delete", "get", "getIamPolicy", "list", "setIamPolicy", "update"}},
		{Method: bigtableLocations, Actions: []string{"list"}},
		{Method: bigtableTables, Actions: []string{"checkConsistency", "create", "delete", "generateConsistencyToken", "get", "list", "mutateRows", "readRows", "sampleRowKeys", "update"}},
	},

	Billing: {
		{Method: billingResourceCosts, Actions: []string{"get"}},
	},

	BinaryAuthorization: {
		{Method: binaryauthorizationAttestors, Actions: []string{"create", "delete", "get", "getIamPolicy", "list", "setIamPolicy", "update", "verifyImageAttested"}},
		{Method: binaryauthorizationPolicy, Actions: []string{"get", "getIamPolicy", "setIamPolicy", "update"}},
	},

	CloudAsset: {
		{Method: cloudassetAssets, Actions: []string{"exportIamPolicy", "exportResource"}},
	},

	CloudBuild: {
		{Method: cloudbuildBuilds, Actions: []string{"create", "get", "list", "update"}},
	},

	CloudConfig: {
		{Method: cloudconfigConfigs, Actions: []string{"get", "update"}},
	},

	Clouddebugger: {
		{Method: clouddebuggerBreakpoints, Actions: []string{"create", "delete", "get", "list", "listActive", "update"}},
		{Method: clouddebuggerDebuggees, Actions: []string{"create", "list"}},
	},

	CloudFunctions: {
		{Method: cloudfunctionsFunctions, Actions: []string{"call", "create", "delete", "get", "getIamPolicy", "invoke", "list", "setIamPolicy", "sourceCodeGet", "sourceCodeSet", "update"}},
		{Method: cloudfunctionsLocations, Actions: []string{"list"}},
		{Method: cloudfunctionsOperations, Actions: []string{"get", "list"}},
	},

	CloudIOT: {
		{Method: cloudiotDevices, Actions: []string{"bindGateway", "create", "delete", "get", "list", "sendCommand", "unbindGateway", "update", "updateConfig"}},
		{Method: cloudiotRegistries, Actions: []string{"create", "delete", "get", "getIamPolicy", "list", "setIamPolicy", "update"}},
		{Method: cloudiottokenTokensettings, Actions: []string{"get", "update"}},
	},

	CloudJobDiscovery: {
		{Method: cloudjobdiscoveryCompanies, Actions: []string{"create", "delete", "get", "list", "update"}},
		{Method: cloudjobdiscoveryEvents, Actions: []string{"create"}},
		{Method: cloudjobdiscoveryJobs, Actions: []string{"create", "delete", "get", "search", "update"}},
		{Method: cloudjobdiscoveryProfiles, Actions: []string{"create", "delete", "get", "search", "update"}},
		{Method: cloudjobdiscoveryTenants, Actions: []string{"create", "delete", "get", "update"}},
		{Method: cloudjobdiscoveryTools, Actions: []string{"access"}},
	},

	CloudKMS: {
		{Method: cloudkmsCryptoKeyVersions, Actions: []string{"create", "destroy", "get", "list", "restore", "update", "useToDecrypt", "useToEncrypt", "useToSign", "viewPublicKey"}},
		{Method: cloudkmsCryptoKeys, Actions: []string{"create", "get", "getIamPolicy", "list", "setIamPolicy", "update"}},
		{Method: cloudkmsKeyRings, Actions: []string{"create", "get", "getIamPolicy", "list", "setIamPolicy"}},
	},

	CloudMessaging: {
		{Method: cloudmessagingMessages, Actions: []string{"create"}},
	},

	CloudMigration: {
		{Method: cloudmigrationVelostrataendpoints, Actions: []string{"connect"}},
	},

	CloudNotifications: {
		{Method: cloudnotificationsActivities, Actions: []string{"list"}},
	},

	CloudPrivateCatalog: {
		{Method: cloudprivatecatalogTargets, Actions: []string{"get"}},
		{Method: cloudprivatecatalogproducerTargets, Actions: []string{"associate", "unassociate"}},
	},

	Cloudprofiler: {
		{Method: cloudprofilerProfiles, Actions: []string{"create", "list", "update"}},
	},

	CloudScheduler: {
		{Method: cloudschedulerJobs, Actions: []string{"create", "delete", "enable", "fullView", "get", "list", "pause", "run", "update"}},
		{Method: cloudschedulerLocations, Actions: []string{"get", "list"}},
	},

	CloudSecurityScanner: {
		{Method: cloudsecurityscannerCrawledurls, Actions: []string{"list"}},
		{Method: cloudsecurityscannerResults, Actions: []string{"get", "list"}},
		{Method: cloudsecurityscannerScanruns, Actions: []string{"get", "getSummary", "list", "stop"}},
		{Method: cloudsecurityscannerScans, Actions: []string{"create", "delete", "get", "list", "run", "update"}},
	},

	CloudSQL: {
		{Method: cloudsqlBackupRuns, Actions: []string{"create", "delete", "get", "list"}},
		{Method: cloudsqlDatabases, Actions: []string{"create", "delete", "get", "list", "update"}},
		{Method: cloudsqlInstances, Actions: []string{"addServerCa", "clone", "connect", "create", "delete", "demoteMaster", "export", "failover", "get", "import", "list", "listServerCas", "promoteReplica", "resetSslConfig", "restart", "restoreBackup", "rotateServerCa", "startReplica", "stopReplica", "truncateLog", "update"}},
		{Method: cloudsqlSslCerts, Actions: []string{"create", "createEphemeral", "delete", "get", "list"}},
		{Method: cloudsqlUsers, Actions: []string{"create", "delete", "list", "update"}},
	},

	CloudTasks: {
		{Method: cloudtasksLocations, Actions: []string{"get", "list"}},
		{Method: cloudtasksQueues, Actions: []string{"create", "delete", "get", "getIamPolicy", "list", "pause", "purge", "resume", "setIamPolicy", "update"}},
		{Method: cloudtasksTasks, Actions: []string{"create", "delete", "fullView", "get", "list", "run"}},
	},

	CloudTestService: {
		{Method: cloudtestserviceEnvironmentcatalog, Actions: []string{"get"}},
		{Method: cloudtestserviceMatrices, Actions: []string{"create", "get", "update"}},
	},

	CloudToolResults: {
		{Method: cloudtoolresultsExecutions, Actions: []string{"create", "get", "list", "update"}},
		{Method: cloudtoolresultsHistories, Actions: []string{"create", "get", "list"}},
		{Method: cloudtoolresultsSettings, Actions: []string{"create", "get", "update"}},
		{Method: cloudtoolresultsSteps, Actions: []string{"create", "get", "list", "update"}},
	},

	CloudTrace: {
		{Method: cloudtraceInsights, Actions: []string{"get", "list"}},
		{Method: cloudtraceStats, Actions: []string{"get"}},
		{Method: cloudtraceTasks, Actions: []string{"create", "delete", "get", "list"}},
		{Method: cloudtraceTraces, Actions: []string{"get", "list", "patch"}},
	},

	CloudTranslate: {
		{Method: cloudtranslateGeneralModels, Actions: []string{"batchPredict", "get", "predict"}},
		{Method: cloudtranslateGlossaries, Actions: []string{"batchPredict", "create", "delete", "get", "list", "predict"}},
		{Method: cloudtranslateLanguageDetectionModels, Actions: []string{"predict"}},
		{Method: cloudtranslateLocations, Actions: []string{"get", "list"}},
		{Method: cloudtranslateOperations, Actions: []string{"cancel", "delete", "get", "list", "wait"}},
	},

	Composer: {
		{Method: composerEnvironments, Actions: []string{"create", "delete", "get", "list", "update"}},
		{Method: composerImageversions, Actions: []string{"list"}},
		{Method: composerOperations, Actions: []string{"delete", "get", "list"}},
	},

	Compute: {
		{Method: computeAcceleratorTypes, Actions: []string{"get", "list"}},
		{Method: computeAddresses, Actions: []string{"create", "createInternal", "delete", "deleteInternal", "get", "list", "setLabels", "use", "useInternal"}},
		{Method: computeAutoscalers, Actions: []string{"create", "delete", "get", "list", "update"}},
		{Method: computeBackendBuckets, Actions: []string{"create", "delete", "get", "list", "update", "use"}},
		{Method: computeBackendServices, Actions: []string{"create", "delete", "get", "list", "setSecurityPolicy", "update", "use"}},
		{Method: computeCommitments, Actions: []string{"create", "get", "list"}},
		{Method: computeDiskTypes, Actions: []string{"get", "list"}},
		{Method: computeDisks, Actions: []string{"addResourcePolicies", "create", "createSnapshot", "delete", "get", "getIamPolicy", "list", "removeResourcePolicies", "resize", "setIamPolicy", "setLabels", "update", "use", "useReadOnly"}},
		{Method: computeFirewalls, Actions: []string{"create", "delete", "get", "list", "update"}},
		{Method: computeForwardingRules, Actions: []string{"create", "delete", "get", "list", "setLabels", "setTarget"}},
		{Method: computeGlobalAddresses, Actions: []string{"create", "createInternal", "delete", "deleteInternal", "get", "list", "setLabels", "use"}},
		{Method: computeGlobalForwardingRules, Actions: []string{"create", "delete", "get", "list", "setLabels", "setTarget"}},
		{Method: computeGlobalOperations, Actions: []string{"delete", "get", "getIamPolicy", "list", "setIamPolicy"}},
		{Method: computeHealthChecks, Actions: []string{"create", "delete", "get", "list", "update", "use", "useReadOnly"}},
		{Method: computeHttpHealthChecks, Actions: []string{"create", "delete", "get", "list", "update", "use", "useReadOnly"}},
		{Method: computeHttpsHealthChecks, Actions: []string{"create", "delete", "get", "list", "update", "use", "useReadOnly"}},
		{Method: computeImages, Actions: []string{"create", "delete", "deprecate", "get", "getFromFamily", "getIamPolicy", "list", "setIamPolicy", "setLabels", "update", "useReadOnly"}},
		{Method: computeInstanceGroupManagers, Actions: []string{"create", "delete", "get", "list", "update", "use"}},
		{Method: computeInstanceGroups, Actions: []string{"create", "delete", "get", "list", "update", "use"}},
		{Method: computeInstanceTemplates, Actions: []string{"create", "delete", "get", "getIamPolicy", "list", "setIamPolicy", "useReadOnly"}},
		{Method: computeInstances, Actions: []string{"addAccessConfig", "addMaintenancePolicies", "attachDisk", "create", "delete", "deleteAccessConfig", "detachDisk", "get", "getGuestAttributes", "getIamPolicy", "getSerialPortOutput", "getShieldedInstanceIdentity", "getShieldedVmIdentity", "list", "listReferrers", "osAdminLogin", "osLogin", "removeMaintenancePolicies", "reset", "resume", "setDeletionProtection", "setDiskAutoDelete", "setIamPolicy", "setLabels", "setMachineResources", "setMachineType", "setMetadata", "setMinCpuPlatform", "setScheduling", "setServiceAccount", "setShieldedInstanceIntegrityPolicy", "setShieldedVmIntegrityPolicy", "setTags", "start", "startWithEncryptionKey", "stop", "suspend", "update", "updateAccessConfig", "updateDisplayDevice", "updateNetworkInterface", "updateShieldedInstanceConfig", "updateShieldedVmConfig", "use"}},
		{Method: computeInterconnectAttachments, Actions: []string{"create", "delete", "get", "list", "setLabels", "update", "use"}},
		{Method: computeInterconnectLocations, Actions: []string{"get", "list"}},
		{Method: computeInterconnects, Actions: []string{"create", "delete", "get", "list", "setLabels", "update", "use"}},
		{Method: computeLicenseCodes, Actions: []string{"get", "getIamPolicy", "list", "setIamPolicy", "update", "use"}},
		{Method: computeLicenses, Actions: []string{"create", "delete", "get", "getIamPolicy", "list", "setIamPolicy"}},
		{Method: computeMachineTypes, Actions: []string{"get", "list"}},
		{Method: computeMaintenancePolicies, Actions: []string{"create", "delete", "get", "getIamPolicy", "list", "setIamPolicy", "use"}},
		{Method: computeNetworkEndpointGroups, Actions: []string{"attachNetworkEndpoints", "create", "delete", "detachNetworkEndpoints", "get", "getIamPolicy", "list", "setIamPolicy", "use"}},
		{Method: computeNetworks, Actions: []string{"addPeering", "create", "delete", "get", "list", "removePeering", "switchToCustomMode", "update", "updatePeering", "updatePolicy", "use", "useExternalIp"}},
		{Method: computeNodeGroups, Actions: []string{"addNodes", "create", "delete", "deleteNodes", "get", "getIamPolicy", "list", "setIamPolicy", "setNodeTemplate"}},
		{Method: computeNodeTemplates, Actions: []string{"create", "delete", "get", "getIamPolicy", "list", "setIamPolicy"}},
		{Method: computeNodeTypes, Actions: []string{"get", "list"}},
		{Method: computeProjects, Actions: []string{"get", "setDefaultNetworkTier", "setDefaultServiceAccount", "setUsageExportBucket"}},
		{Method: computeRegionBackendServices, Actions: []string{"create", "delete", "get", "list", "setSecurityPolicy", "update", "use"}},
		{Method: computeRegionOperations, Actions: []string{"delete", "get", "getIamPolicy", "list", "setIamPolicy"}},
		{Method: computeRegions, Actions: []string{"get", "list"}},
		{Method: computeReservations, Actions: []string{"create", "delete", "get", "list", "resize"}},
		{Method: computeResourcePolicies, Actions: []string{"create", "delete", "get", "list", "use"}},
		{Method: computeRouters, Actions: []string{"create", "delete", "get", "list", "update", "use"}},
		{Method: computeRoutes, Actions: []string{"create", "delete", "get", "list"}},
		{Method: computeSecurityPolicies, Actions: []string{"create", "delete", "get", "getIamPolicy", "list", "setIamPolicy", "update", "use"}},
		{Method: computeSnapshots, Actions: []string{"create", "delete", "get", "getIamPolicy", "list", "setIamPolicy", "setLabels", "useReadOnly"}},
		{Method: computeSslCertificates, Actions: []string{"create", "delete", "get", "list"}},
		{Method: computeSslPolicies, Actions: []string{"create", "delete", "get", "list", "listAvailableFeatures", "update", "use"}},
		{Method: computeSubnetworks, Actions: []string{"create", "delete", "expandIpCidrRange", "get", "getIamPolicy", "list", "setIamPolicy", "setPrivateIpGoogleAccess", "update", "use", "useExternalIp"}},
		{Method: computeTargetHttpProxies, Actions: []string{"create", "delete", "get", "list", "setUrlMap", "use"}},
		{Method: computeTargetHttpsProxies, Actions: []string{"create", "delete", "get", "list", "setSslCertificates", "setSslPolicy", "setUrlMap", "use"}},
		{Method: computeTargetInstances, Actions: []string{"create", "delete", "get", "list", "use"}},
		{Method: computeTargetPools, Actions: []string{"addHealthCheck", "addInstance", "create", "delete", "get", "list", "removeHealthCheck", "removeInstance", "update", "use"}},
		{Method: computeTargetSslProxies, Actions: []string{"create", "delete", "get", "list", "setBackendService", "setProxyHeader", "setSslCertificates", "use"}},
		{Method: computeTargetTcpProxies, Actions: []string{"create", "delete", "get", "list", "update", "use"}},
		{Method: computeTargetVpnGateways, Actions: []string{"create", "delete", "get", "list", "setLabels", "use"}},
		{Method: computeUrlMaps, Actions: []string{"create", "delete", "get", "invalidateCache", "list", "update", "use", "validate"}},
		{Method: computeVpnGateways, Actions: []string{"create", "delete", "get", "list", "setLabels", "use"}},
		{Method: computeVpnTunnels, Actions: []string{"create", "delete", "get", "list", "setLabels"}},
		{Method: computeZoneOperations, Actions: []string{"delete", "get", "getIamPolicy", "list", "setIamPolicy"}},
		{Method: computeZones, Actions: []string{"get", "list"}},
	},

	Container: {
		{Method: containerApiServices, Actions: []string{"create", "delete", "get", "list", "update", "updateStatus"}},
		{Method: containerBackendConfigs, Actions: []string{"create", "delete", "get", "list", "update"}},
		{Method: containerBindings, Actions: []string{"create", "delete", "get", "list", "update"}},
		{Method: containerCertificateSigningRequests, Actions: []string{"approve", "create", "delete", "get", "list", "update", "updateStatus"}},
		{Method: containerClusterRoleBindings, Actions: []string{"create", "delete", "get", "list", "update"}},
		{Method: containerClusterRoles, Actions: []string{"bind", "create", "delete", "get", "list", "update"}},
		{Method: containerClusters, Actions: []string{"create", "delete", "get", "getCredentials", "list", "update"}},
		{Method: containerComponentStatuses, Actions: []string{"get", "list"}},
		{Method: containerConfigMaps, Actions: []string{"create", "delete", "get", "list", "update"}},
		{Method: containerControllerRevisions, Actions: []string{"create", "delete", "get", "list", "update"}},
		{Method: containerCronJobs, Actions: []string{"create", "delete", "get", "getStatus", "list", "update", "updateStatus"}},
		{Method: containerCustomResourceDefinitions, Actions: []string{"create", "delete", "get", "list", "update", "updateStatus"}},
		{Method: containerDaemonSets, Actions: []string{"create", "delete", "get", "getStatus", "list", "update", "updateStatus"}},
		{Method: containerDeployments, Actions: []string{"create", "delete", "get", "getScale", "getStatus", "list", "rollback", "update", "updateScale", "updateStatus"}},
		{Method: containerEndpoints, Actions: []string{"create", "delete", "get", "list", "update"}},
		{Method: containerEvents, Actions: []string{"create", "delete", "get", "list", "update"}},
		{Method: containerHorizontalPodAutoscalers, Actions: []string{"create", "delete", "get", "getStatus", "list", "update", "updateStatus"}},
		{Method: containerIngresses, Actions: []string{"create", "delete", "get", "getStatus", "list", "update", "updateStatus"}},
		{Method: containerInitializerConfigurations, Actions: []string{"create", "delete", "get", "list", "update"}},
		{Method: containerJobs, Actions: []string{"create", "delete", "get", "getStatus", "list", "update", "updateStatus"}},
		{Method: containerLimitRanges, Actions: []string{"create", "delete", "get", "list", "update"}},
		{Method: containerLocalSubjectAccessReviews, Actions: []string{"create", "list"}},
		{Method: containerNamespaces, Actions: []string{"create", "delete", "get", "getStatus", "list", "update", "updateStatus"}},
		{Method: containerNetworkPolicies, Actions: []string{"create", "delete", "get", "list", "update"}},
		{Method: containerNodes, Actions: []string{"create", "delete", "get", "getStatus", "list", "proxy", "update", "updateStatus"}},
		{Method: containerOperations, Actions: []string{"get", "list"}},
		{Method: containerPersistentVolumeClaims, Actions: []string{"create", "delete", "get", "getStatus", "list", "update", "updateStatus"}},
		{Method: containerPersistentVolumes, Actions: []string{"create", "delete", "get", "getStatus", "list", "update", "updateStatus"}},
		{Method: containerPetSets, Actions: []string{"create", "delete", "get", "list", "update", "updateStatus"}},
		{Method: containerPodDisruptionBudgets, Actions: []string{"create", "delete", "get", "getStatus", "list", "update", "updateStatus"}},
		{Method: containerPodPresets, Actions: []string{"create", "delete", "get", "list", "update"}},
		{Method: containerPodSecurityPolicies, Actions: []string{"create", "delete", "get", "list", "update", "use"}},
		{Method: containerPodTemplates, Actions: []string{"create", "delete", "get", "list", "update"}},
		{Method: containerPods, Actions: []string{"attach", "create", "delete", "evict", "exec", "get", "getLogs", "getStatus", "initialize", "list", "portForward", "proxy", "update", "updateStatus"}},
		{Method: containerReplicaSets, Actions: []string{"create", "delete", "get", "getScale", "getStatus", "list", "update", "updateScale", "updateStatus"}},
		{Method: containerReplicationControllers, Actions: []string{"create", "delete", "get", "getScale", "getStatus", "list", "update", "updateScale", "updateStatus"}},
		{Method: containerResourceQuotas, Actions: []string{"create", "delete", "get", "getStatus", "list", "update", "updateStatus"}},
		{Method: containerRoleBindings, Actions: []string{"create", "delete", "get", "list", "update"}},
		{Method: containerRoles, Actions: []string{"bind", "create", "delete", "get", "list", "update"}},
		{Method: containerScheduledJobs, Actions: []string{"create", "delete", "get", "list", "update", "updateStatus"}},
		{Method: containerSecrets, Actions: []string{"create", "delete", "get", "list", "update"}},
		{Method: containerSelfSubjectAccessReviews, Actions: []string{"create", "list"}},
		{Method: containerServiceAccounts, Actions: []string{"create", "delete", "get", "list", "update"}},
		{Method: containerServices, Actions: []string{"create", "delete", "get", "getStatus", "list", "proxy", "update", "updateStatus"}},
		{Method: containerStatefulSets, Actions: []string{"create", "delete", "get", "getScale", "getStatus", "list", "update", "updateScale", "updateStatus"}},
		{Method: containerStorageClasses, Actions: []string{"create", "delete", "get", "list", "update"}},
		{Method: containerSubjectAccessReviews, Actions: []string{"create", "list"}},
		{Method: containerThirdPartyObjects, Actions: []string{"create", "delete", "get", "list", "update"}},
		{Method: containerThirdPartyResources, Actions: []string{"create", "delete", "get", "list", "update"}},
		{Method: containerTokenReviews, Actions: []string{"create"}},
	},

	DataCatalog: {
		{Method: datacatalogEntries, Actions: []string{"create", "delete", "get", "getIamPolicy", "setIamPolicy", "update"}},
		{Method: datacatalogEntryGroups, Actions: []string{"create", "delete", "get", "getIamPolicy", "setIamPolicy"}},
		{Method: datacatalogTagTemplates, Actions: []string{"create", "delete", "get", "getIamPolicy", "getTag", "setIamPolicy", "update", "use"}},
	},

	DataFlowJobs: {
		{Method: dataflowJobs, Actions: []string{"cancel", "create", "get", "list", "updateContents"}},
		{Method: dataflowMessages, Actions: []string{"list"}},
		{Method: dataflowMetrics, Actions: []string{"get"}},
	},

	DataFusion: {
		{Method: datafusionInstances, Actions: []string{"create", "delete", "get", "getIamPolicy", "list", "restart", "setIamPolicy", "update", "upgrade"}},
		{Method: datafusionLocations, Actions: []string{"get", "list"}},
		{Method: datafusionOperations, Actions: []string{"cancel", "get", "list"}},
	},

	DataLabeling: {
		{Method: datalabelingAnnotateddatasets, Actions: []string{"delete", "get", "label", "list"}},
		{Method: datalabelingAnnotationspecsets, Actions: []string{"create", "delete", "get", "list"}},
		{Method: datalabelingDataitems, Actions: []string{"get", "list"}},
		{Method: datalabelingDatasets, Actions: []string{"create", "delete", "export", "get", "import", "list"}},
		{Method: datalabelingExamples, Actions: []string{"get", "list"}},
		{Method: datalabelingInstructions, Actions: []string{"create", "delete", "get", "list"}},
		{Method: datalabelingOperations, Actions: []string{"cancel", "get", "list"}},
	},

	DataPrep: {
		{Method: dataprepProjects, Actions: []string{"use"}},
	},

	DataProc: {
		{Method: dataprocAgents, Actions: []string{"create", "delete", "get", "list", "update"}},
		{Method: dataprocClusters, Actions: []string{"create", "delete", "get", "getIamPolicy", "list", "setIamPolicy", "update", "use"}},
		{Method: dataprocJobs, Actions: []string{"cancel", "create", "delete", "get", "getIamPolicy", "list", "setIamPolicy", "update"}},
		{Method: dataprocOperations, Actions: []string{"cancel", "delete", "get", "getIamPolicy", "list", "setIamPolicy"}},
		{Method: dataprocTasks, Actions: []string{"lease", "listInvalidatedLeases", "reportStatus"}},
		{Method: dataprocWorkflowTemplates, Actions: []string{"create", "delete", "get", "getIamPolicy", "instantiate", "instantiateInline", "list", "setIamPolicy", "update"}},
	},

	DataStore: {
		{Method: datastoreDatabases, Actions: []string{"create", "delete", "export", "get", "getIamPolicy", "import", "list", "setIamPolicy", "update"}},
		{Method: datastoreEntities, Actions: []string{"allocateIds", "create", "delete", "get", "list", "update"}},
		{Method: datastoreIndexes, Actions: []string{"create", "delete", "get", "list", "update"}},
		{Method: datastoreLocations, Actions: []string{"get", "list"}},
		{Method: datastoreNamespaces, Actions: []string{"get", "getIamPolicy", "list", "setIamPolicy"}},
		{Method: datastoreOperations, Actions: []string{"cancel", "delete", "get", "list"}},
		{Method: datastoreStatistics, Actions: []string{"get", "list"}},
	},

	DeploymentManager: {
		{Method: deploymentmanagerCompositeTypes, Actions: []string{"create", "delete", "get", "list", "update"}},
		{Method: deploymentmanagerDeployments, Actions: []string{"cancelPreview", "create", "delete", "get", "getIamPolicy", "list", "setIamPolicy", "stop", "update"}},
		{Method: deploymentmanagerManifests, Actions: []string{"get", "list"}},
		{Method: deploymentmanagerOperations, Actions: []string{"get", "list"}},
		{Method: deploymentmanagerResources, Actions: []string{"get", "list"}},
		{Method: deploymentmanagerTypeProviders, Actions: []string{"create", "delete", "get", "getType", "list", "listTypes", "update"}},
		{Method: deploymentmanagerTypes, Actions: []string{"create", "delete", "get", "list", "update"}},
	},

	DialogFlow: {
		{Method: dialogflowAgents, Actions: []string{"create", "delete", "export", "get", "import", "restore", "search", "train", "update"}},
		{Method: dialogflowContexts, Actions: []string{"create", "delete", "get", "list", "update"}},
		{Method: dialogflowEntityTypes, Actions: []string{"create", "createEntity", "delete", "deleteEntity", "get", "list", "update", "updateEntity"}},
		{Method: dialogflowIntents, Actions: []string{"create", "delete", "get", "list", "update"}},
		{Method: dialogflowOperations, Actions: []string{"get"}},
		{Method: dialogflowSessionEntityTypes, Actions: []string{"create", "delete", "get", "list", "update"}},
		{Method: dialogflowSessions, Actions: []string{"detectIntent", "streamingDetectIntent"}},
	},

	DLP: {
		{Method: dlpAnalyzeRiskTemplates, Actions: []string{"create", "delete", "get", "list", "update"}},
		{Method: dlpDeidentifyTemplates, Actions: []string{"create", "delete", "get", "list", "update"}},
		{Method: dlpInspectTemplates, Actions: []string{"create", "delete", "get", "list", "update"}},
		{Method: dlpJobTriggers, Actions: []string{"create", "delete", "get", "list", "update"}},
		{Method: dlpJobs, Actions: []string{"cancel", "create", "delete", "get", "list"}},
		{Method: dlpKms, Actions: []string{"encrypt"}},
		{Method: dlpStoredInfoTypes, Actions: []string{"create", "delete", "get", "list", "update"}},
	},

	DNS: {
		{Method: dnsChanges, Actions: []string{"create", "get", "list"}},
		{Method: dnsDnsKeys, Actions: []string{"get", "list"}},
		{Method: dnsManagedZoneOperations, Actions: []string{"get", "list"}},
		{Method: dnsManagedZones, Actions: []string{"create", "delete", "get", "list", "update"}},
		{Method: dnsNetworks, Actions: []string{"bindPrivateDNSPolicy", "bindPrivateDNSZone", "targetWithPeeringZone"}},
		{Method: dnsPolicies, Actions: []string{"create", "delete", "get", "getIamPolicy", "list", "setIamPolicy", "update"}},
		{Method: dnsProjects, Actions: []string{"get"}},
		{Method: dnsResourceRecordSets, Actions: []string{"create", "delete", "list", "update"}},
	},

	Endpoints: {
		{Method: endpointsPortals, Actions: []string{"attachCustomDomain", "detachCustomDomain", "listCustomDomains", "update"}},
	},

	ErrorReporting: {
		{Method: errorreportingApplications, Actions: []string{"list"}},
		{Method: errorreportingErrorEvents, Actions: []string{"create", "delete", "list"}},
		{Method: errorreportingGroupMetadata, Actions: []string{"get", "update"}},
		{Method: errorreportingGroups, Actions: []string{"list"}},
	},

	File: {
		{Method: fileInstances, Actions: []string{"create", "delete", "get", "list", "restore", "update"}},
		{Method: fileLocations, Actions: []string{"get", "list"}},
		{Method: fileOperations, Actions: []string{"cancel", "delete", "get", "list"}},
		{Method: fileSnapshots, Actions: []string{"create", "delete", "get", "list", "update"}},
	},

	Firebase: {
		{Method: firebaseBillingPlans, Actions: []string{"get", "update"}},
		{Method: firebaseClients, Actions: []string{"create", "delete", "get"}},
		{Method: firebaseLinks, Actions: []string{"create", "delete", "list", "update"}},
		{Method: firebaseProjects, Actions: []string{"delete", "get", "update"}},
	},

	FirebaseAbt: {
		{Method: firebaseabtExperimentresults, Actions: []string{"get"}},
		{Method: firebaseabtExperiments, Actions: []string{"create", "delete", "get", "list", "update"}},
		{Method: firebaseabtProjectmetadata, Actions: []string{"get"}},
	},

	FirebaseAnalytics: {
		{Method: firebaseanalyticsResources, Actions: []string{"googleAnalyticsEdit", "googleAnalyticsReadAndAnalyze"}},
	},

	FirebaseAuth: {
		{Method: firebaseauthConfigs, Actions: []string{"create", "get", "getHashConfig", "update"}},
		{Method: firebaseauthUsers, Actions: []string{"create", "createSession", "delete", "get", "sendEmail", "update"}},
	},

	FirebaseCrash: {
		{Method: firebasecrashIssues, Actions: []string{"update"}},
		{Method: firebasecrashReports, Actions: []string{"get"}},
		{Method: firebasecrashlyticsConfig, Actions: []string{"get", "update"}},
		{Method: firebasecrashlyticsData, Actions: []string{"get"}},
		{Method: firebasecrashlyticsIssues, Actions: []string{"get", "list", "update"}},
		{Method: firebasecrashlyticsSessions, Actions: []string{"get"}},
	},

	FirebaseDatabase: {
		{Method: firebasedatabaseInstances, Actions: []string{"create", "get", "list", "update"}},
	},

	FirebaseDynamicLinks: {
		{Method: firebasedynamiclinksDestinations, Actions: []string{"list", "update"}},
		{Method: firebasedynamiclinksDomains, Actions: []string{"create", "delete", "get", "list", "update"}},
		{Method: firebasedynamiclinksLinks, Actions: []string{"create", "get", "list", "update"}},
		{Method: firebasedynamiclinksStats, Actions: []string{"get"}},
	},

	FirebaseExtensions: {
		{Method: firebaseextensionsConfigs, Actions: []string{"create", "delete", "list", "update"}},
	},

	FirebaseHosting: {
		{Method: firebasehostingSites, Actions: []string{"create", "delete", "get", "list", "update"}},
	},

	FirebaseInAppMessaging: {
		{Method: firebaseinappmessagingCampaigns, Actions: []string{"create", "delete", "get", "list", "update"}},
	},

	FirebaseML: {
		{Method: firebasemlCompressionjobs, Actions: []string{"create", "delete", "get", "list", "start", "update"}},
		{Method: firebasemlModels, Actions: []string{"create", "delete", "get", "list"}},
		{Method: firebasemlModelversions, Actions: []string{"create", "get", "list", "update"}},
	},

	FirebaseNotifications: {
		{Method: firebasenotificationsMessages, Actions: []string{"create", "delete", "get", "list", "update"}},
	},

	FirebasePerformance: {
		{Method: firebaseperformanceConfig, Actions: []string{"create", "delete", "update"}},
		{Method: firebaseperformanceData, Actions: []string{"get"}},
	},

	FirebasePredictions: {
		{Method: firebasepredictionsPredictions, Actions: []string{"create", "delete", "list", "update"}},
	},

	FirebaseRules: {
		{Method: firebaserulesReleases, Actions: []string{"create", "delete", "get", "getExecutable", "list", "update"}},
		{Method: firebaserulesRulesets, Actions: []string{"create", "delete", "get", "list", "test"}},
	},

	Genomics: {
		{Method: genomicsDatasets, Actions: []string{"create", "delete", "get", "getIamPolicy", "list", "setIamPolicy", "update"}},
		{Method: genomicsOperations, Actions: []string{"cancel", "create", "get", "list"}},
	},

	Healthcare: {
		{Method: healthcareDatasets, Actions: []string{"create", "deidentify", "delete", "get", "getIamPolicy", "list", "setIamPolicy", "update"}},
		{Method: healthcareDicomStores, Actions: []string{"create", "delete", "dicomWebDelete", "dicomWebRead", "dicomWebWrite", "export", "get", "getIamPolicy", "import", "list", "setIamPolicy", "update"}},
		{Method: healthcareFhirResources, Actions: []string{"create", "delete", "get", "patch", "purge", "update"}},
		{Method: healthcareFhirStores, Actions: []string{"create", "delete", "export", "get", "getIamPolicy", "import", "list", "searchResources", "setIamPolicy", "update"}},
		{Method: healthcareHl7V2Messages, Actions: []string{"create", "delete", "get", "ingest", "list", "update"}},
		{Method: healthcareHl7V2Stores, Actions: []string{"create", "delete", "get", "getIamPolicy", "list", "setIamPolicy", "update"}},
		{Method: healthcareOperations, Actions: []string{"cancel", "get", "list"}},
	},

	Iam: {
		{Method: iamRoles, Actions: []string{"create", "delete", "get", "list", "undelete", "update"}},
		{Method: iamServiceAccountKeys, Actions: []string{"create", "delete", "get", "list"}},
		{Method: iamServiceAccounts, Actions: []string{"actAs", "create", "delete", "get", "getIamPolicy", "list", "setIamPolicy", "update"}},
	},

	Iap: {
		{Method: iapTunnel, Actions: []string{"getIamPolicy", "setIamPolicy"}},
		{Method: iapTunnelInstances, Actions: []string{"accessViaIAP", "getIamPolicy", "setIamPolicy"}},
		{Method: iapTunnelZones, Actions: []string{"getIamPolicy", "setIamPolicy"}},
		{Method: iapWeb, Actions: []string{"getIamPolicy", "setIamPolicy"}},
		{Method: iapWebServiceVersions, Actions: []string{"getIamPolicy", "setIamPolicy"}},
		{Method: iapWebServices, Actions: []string{"getIamPolicy", "setIamPolicy"}},
		{Method: iapWebTypes, Actions: []string{"getIamPolicy", "setIamPolicy"}},
	},

	Logging: {
		{Method: loggingExclusions, Actions: []string{"create", "delete", "get", "list", "update"}},
		{Method: loggingLogEntries, Actions: []string{"create", "list"}},
		{Method: loggingLogMetrics, Actions: []string{"create", "delete", "get", "list", "update"}},
		{Method: loggingLogServiceIndexes, Actions: []string{"list"}},
		{Method: loggingLogServices, Actions: []string{"list"}},
		{Method: loggingLogs, Actions: []string{"delete", "list"}},
		{Method: loggingPrivateLogEntries, Actions: []string{"list"}},
		{Method: loggingSinks, Actions: []string{"create", "delete", "get", "list", "update"}},
		{Method: loggingUsage, Actions: []string{"get"}},
	},

	ManagedIdentities: {
		{Method: managedidentitiesDomains, Actions: []string{"attachTrust", "create", "delete", "detachTrust", "get", "getIamPolicy", "list", "reconfigureTrust", "resetpassword", "setIamPolicy", "update", "validateTrust"}},
		{Method: managedidentitiesLocations, Actions: []string{"get", "list"}},
		{Method: managedidentitiesOperations, Actions: []string{"cancel", "delete", "get", "list"}},
	},

	ML: {
		{Method: mlJobs, Actions: []string{"cancel", "create", "get", "getIamPolicy", "list", "setIamPolicy", "update"}},
		{Method: mlLocations, Actions: []string{"get", "list"}},
		{Method: mlModels, Actions: []string{"create", "delete", "get", "getIamPolicy", "list", "predict", "setIamPolicy", "update"}},
		{Method: mlOperations, Actions: []string{"cancel", "get", "list"}},
		{Method: mlProjects, Actions: []string{"getConfig"}},
		{Method: mlVersions, Actions: []string{"create", "delete", "get", "list", "predict", "update"}},
	},

	Monitoring: {
		{Method: monitoringAlertPolicies, Actions: []string{"create", "delete", "get", "list", "update"}},
		{Method: monitoringDashboards, Actions: []string{"create", "delete", "get", "list", "update"}},
		{Method: monitoringGroups, Actions: []string{"create", "delete", "get", "list", "update"}},
		{Method: monitoringMetricDescriptors, Actions: []string{"create", "delete", "get", "list"}},
		{Method: monitoringMonitoredResourceDescriptors, Actions: []string{"get", "list"}},
		{Method: monitoringNotificationChannelDescriptors, Actions: []string{"get", "list"}},
		{Method: monitoringNotificationChannels, Actions: []string{"create", "delete", "get", "getVerificationCode", "list", "sendVerificationCode", "update", "verify"}},
		{Method: monitoringPublicWidgets, Actions: []string{"create", "delete", "get", "list", "update"}},
		{Method: monitoringTimeSeries, Actions: []string{"create", "list"}},
		{Method: monitoringUptimeCheckConfigs, Actions: []string{"create", "delete", "get", "list", "update"}},
	},

	OrgPolicy: {
		{Method: orgpolicyPolicy, Actions: []string{"get"}},
	},

	ProximityBeacon: {
		{Method: proximitybeaconAttachments, Actions: []string{"create", "delete", "get", "list"}},
		{Method: proximitybeaconBeacons, Actions: []string{"attach", "create", "get", "getIamPolicy", "list", "setIamPolicy", "update"}},
		{Method: proximitybeaconNamespaces, Actions: []string{"create", "delete", "get", "getIamPolicy", "list", "setIamPolicy", "update"}},
	},

	Pubsub: {
		{Method: pubsubSnapshots, Actions: []string{"create", "delete", "get", "getIamPolicy", "list", "seek", "setIamPolicy", "update"}},
		{Method: pubsubSubscriptions, Actions: []string{"consume", "create", "delete", "get", "getIamPolicy", "list", "setIamPolicy", "update"}},
		{Method: pubsubTopics, Actions: []string{"attachSubscription", "create", "delete", "get", "getIamPolicy", "list", "publish", "setIamPolicy", "update", "updateTag"}},
	},

	Redis: {
		{Method: redisInstances, Actions: []string{"create", "delete", "export", "get", "import", "list", "update"}},
		{Method: redisLocations, Actions: []string{"get", "list"}},
		{Method: redisOperations, Actions: []string{"cancel", "delete", "get", "list"}},
	},

	RemoteBuildExecution: {
		{Method: remotebuildexecutionActions, Actions: []string{"create", "get", "update"}},
		{Method: remotebuildexecutionBlobs, Actions: []string{"create", "get"}},
		{Method: remotebuildexecutionBotsessions, Actions: []string{"create", "update"}},
		{Method: remotebuildexecutionInstances, Actions: []string{"create", "delete", "get", "list"}},
		{Method: remotebuildexecutionLogstreams, Actions: []string{"create", "get", "update"}},
		{Method: remotebuildexecutionWorkerpools, Actions: []string{"create", "delete", "get", "list", "update"}},
	},

	ResourceManager: {
		{Method: resourcemanagerProjects, Actions: []string{"createBillingAssignment", "delete", "deleteBillingAssignment", "get", "getIamPolicy", "setIamPolicy", "undelete", "update", "updateLiens"}},
	},

	Run: {
		{Method: runConfigurations, Actions: []string{"get", "list"}},
		{Method: runLocations, Actions: []string{"list"}},
		{Method: runRevisions, Actions: []string{"delete", "get", "list"}},
		{Method: runRoutes, Actions: []string{"get", "invoke", "list"}},
		{Method: runServices, Actions: []string{"create", "delete", "get", "getIamPolicy", "list", "setIamPolicy", "update"}},
		{Method: runtimeconfigConfigs, Actions: []string{"create", "delete", "get", "getIamPolicy", "list", "setIamPolicy", "update"}},
		{Method: runtimeconfigOperations, Actions: []string{"get", "list"}},
		{Method: runtimeconfigVariables, Actions: []string{"create", "delete", "get", "getIamPolicy", "list", "setIamPolicy", "update", "watch"}},
		{Method: runtimeconfigWaiters, Actions: []string{"create", "delete", "get", "getIamPolicy", "list", "setIamPolicy", "update"}},
	},

	ServiceBroker: {
		{Method: servicebrokerBindingoperations, Actions: []string{"get", "list"}},
		{Method: servicebrokerBindings, Actions: []string{"create", "delete", "get", "getIamPolicy", "list", "setIamPolicy"}},
		{Method: servicebrokerCatalogs, Actions: []string{"create", "delete", "get", "getIamPolicy", "list", "setIamPolicy", "validate"}},
		{Method: servicebrokerInstanceoperations, Actions: []string{"get", "list"}},
		{Method: servicebrokerInstances, Actions: []string{"create", "delete", "get", "getIamPolicy", "list", "setIamPolicy", "update"}},
	},

	SecretManager: {
		{Method: secretManager, Actions: []string{"create", "delete", "get", "getIamPolicy", "list", "setIamPolicy", "update"}},
	},

	ServiceConsumerManagement: {
		{Method: serviceconsumermanagementConsumers, Actions: []string{"get"}},
		{Method: serviceconsumermanagementQuota, Actions: []string{"get", "update"}},
		{Method: serviceconsumermanagementTenancyu, Actions: []string{"addResource", "create", "delete", "list", "removeResource"}},
		{Method: servicemanagementServices, Actions: []string{"bind", "check", "create", "delete", "get", "getIamPolicy", "list", "quota", "report", "setIamPolicy", "update"}},
	},

	ServiceNetworking: {
		{Method: servicenetworkingOperations, Actions: []string{"cancel", "delete", "get", "list"}},
		{Method: servicenetworkingServices, Actions: []string{"addPeering", "addSubnetwork", "get"}},
	},

	ServiceUsage: {
		{Method: serviceusageApiKeys, Actions: []string{"create", "delete", "get", "getProjectForKey", "list", "regenerate", "revert", "update"}},
		{Method: serviceusageOperations, Actions: []string{"cancel", "delete", "get", "list"}},
		{Method: serviceusageQuotas, Actions: []string{"get", "update"}},
		{Method: serviceusageServices, Actions: []string{"disable", "enable", "get", "list", "use"}},
	},

	Source: {
		{Method: sourceRepos, Actions: []string{"create", "delete", "get", "getIamPolicy", "getProjectConfig", "list", "setIamPolicy", "update", "updateProjectConfig", "updateRepoConfig"}},
	},

	Spanner: {
		{Method: spannerDatabaseOperations, Actions: []string{"cancel", "delete", "get", "list"}},
		{Method: spannerDatabases, Actions: []string{"beginOrRollbackReadWriteTransaction", "beginReadOnlyTransaction", "create", "drop", "get", "getDdl", "getIamPolicy", "list", "read", "select", "setIamPolicy", "update", "updateDdl", "write"}},
		{Method: spannerInstanceConfigs, Actions: []string{"get", "list"}},
		{Method: spannerInstanceOperations, Actions: []string{"cancel", "delete", "get", "list"}},
		{Method: spannerInstances, Actions: []string{"create", "delete", "get", "getIamPolicy", "list", "setIamPolicy", "update"}},
		{Method: spannerSessions, Actions: []string{"create", "delete", "get", "list"}},
	},

	StackDriver: {
		{Method: stackdriverProjects, Actions: []string{"edit", "get"}},
		{Method: stackdriverResourceMetadata, Actions: []string{"write"}},
	},

	Storage: {
		{Method: storageBuckets, Actions: []string{"create", "delete", "list"}},
		{Method: storageObjects, Actions: []string{"create", "delete", "get", "getIamPolicy", "list", "setIamPolicy", "update"}},
		{Method: storageHmacKeys, Actions: []string{"create", "delete", "get", "list", "update"}},
		{Method: storagetransferJobs, Actions: []string{"create", "delete", "get", "list", "update"}},
		{Method: storagetransferOperations, Actions: []string{"cancel", "get", "list", "pause", "resume"}},
		{Method: storagetransferProjects, Actions: []string{"getServiceAccount"}},
	},

	SubscribeTithGoogleDeveloper: {
		{Method: subscribewithgoogledeveloperTools, Actions: []string{"get"}},
	},

	ThreatDetection: {
		{Method: threatdetectionDetectorSettings, Actions: []string{"clear", "get", "update"}},
	},

	Tpu: {
		{Method: tpuAcceleratortypes, Actions: []string{"get", "list"}},
		{Method: tpuLocations, Actions: []string{"get", "list"}},
		{Method: tpuNodes, Actions: []string{"create", "delete", "get", "list", "reimage", "reset", "start", "stop"}},
		{Method: tpuOperations, Actions: []string{"get", "list"}},
		{Method: tpuTensorflowversions, Actions: []string{"get", "list"}},
	},

	VpcAccess: {
		{Method: vpcaccessConnectors, Actions: []string{"create", "delete", "get", "list", "use"}},
		{Method: vpcaccessLocations, Actions: []string{"list"}},
		{Method: vpcaccessOperations, Actions: []string{"get", "list"}},
	},
}
