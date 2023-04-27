package azure

import "github.com/securisec/cliam/azure/policy"

// GetGroupPolicyKeys get keys for grouped policies
func GetGroupPolicyKeys() []string {
	hold := make([]string, 0, len(GroupedPolicies))
	for k := range GroupedPolicies {
		hold = append(hold, k)
	}
	return hold
}

// GroupedPolicies map of grouped policies as specified by the rest api spec
var GroupedPolicies = map[string][]map[string]policy.Policy{
	"addons": {},

	"adhybridhealthservice": {
		policy.Microsoft_ADHybridHealthService_ADHybridHealthService,
	},

	"adp": {},

	"advisor": {
		policy.Microsoft_Advisor_advisor,
	},

	"agrifood": {},

	"alertsmanagement": {
		policy.Microsoft_AlertsManagement_AlertProcessingRules,
	},

	"analysisservices": {
		policy.Microsoft_AnalysisServices_analysisservices,
	},

	"apimanagement": {
		policy.Microsoft_ApiManagement_apimskus,
		policy.Microsoft_ApiManagement_apimreports,
		policy.Microsoft_ApiManagement_apimapisByTags,
		policy.Microsoft_ApiManagement_apimprivatelink,
		policy.Microsoft_ApiManagement_apimnamedvalues,
		policy.Microsoft_ApiManagement_apimdeployment,
		policy.Microsoft_ApiManagement_apimcontenttypes,
		policy.Microsoft_ApiManagement_apimproducts,
		policy.Microsoft_ApiManagement_apimtagresources,
		policy.Microsoft_ApiManagement_apimopenidconnectproviders,
		policy.Microsoft_ApiManagement_apimtenant,
		policy.Microsoft_ApiManagement_apimanagement,
		policy.Microsoft_ApiManagement_apimemailtemplates,
		policy.Microsoft_ApiManagement_apimapis,
		policy.Microsoft_ApiManagement_apimauthorizationservers,
		policy.Microsoft_ApiManagement_apimnotifications,
		policy.Microsoft_ApiManagement_apimtags,
		policy.Microsoft_ApiManagement_apimloggers,
		policy.Microsoft_ApiManagement_apimquotas,
		policy.Microsoft_ApiManagement_apimportalsettings,
		policy.Microsoft_ApiManagement_apimsubscriptions,
		policy.Microsoft_ApiManagement_apimconnectivitycheck,
		policy.Microsoft_ApiManagement_apimportalrevisions,
		policy.Microsoft_ApiManagement_apimoutbounddependency,
		policy.Microsoft_ApiManagement_apimidentityprovider,
		policy.Microsoft_ApiManagement_apimregions,
		policy.Microsoft_ApiManagement_apimissues,
		policy.Microsoft_ApiManagement_apimcaches,
		policy.Microsoft_ApiManagement_apimpolicydescriptions,
		policy.Microsoft_ApiManagement_apimpolicies,
		policy.Microsoft_ApiManagement_apimdiagnostics,
		policy.Microsoft_ApiManagement_apimschema,
		policy.Microsoft_ApiManagement_apimsettings,
		policy.Microsoft_ApiManagement_apimgateways,
		policy.Microsoft_ApiManagement_apimgroups,
		policy.Microsoft_ApiManagement_definitions,
		policy.Microsoft_ApiManagement_apimdeletedservices,
		policy.Microsoft_ApiManagement_apimbackends,
		policy.Microsoft_ApiManagement_apimusers,
		policy.Microsoft_ApiManagement_apimcertificates,
		policy.Microsoft_ApiManagement_apimnetworkstatus,
		policy.Microsoft_ApiManagement_apimproductsByTags,
		policy.Microsoft_ApiManagement_apimapiversionsets,
	},

	"app": {
		policy.Microsoft_App_ConnectedEnvironmentsStorages,
		policy.Microsoft_App_BillingMeters,
		policy.Microsoft_App_ManagedEnvironmentsDaprComponents,
		policy.Microsoft_App_ConnectedEnvironments,
		policy.Microsoft_App_ContainerApps,
		policy.Microsoft_App_ConnectedEnvironmentsDaprComponents,
		policy.Microsoft_App_AuthConfigs,
		policy.Microsoft_App_Diagnostics,
		policy.Microsoft_App_AvailableWorkloadProfiles,
		policy.Microsoft_App_ConnectedEnvironmentsCertificates,
		policy.Microsoft_App_SourceControls,
		policy.Microsoft_App_ManagedEnvironments,
		policy.Microsoft_App_Global,
		policy.Microsoft_App_ContainerAppsRevisions,
		policy.Microsoft_App_CommonDefinitions,
		policy.Microsoft_App_ManagedEnvironmentsStorages,
	},

	"appcomplianceautomation": {},

	"appconfiguration": {
		policy.Microsoft_AppConfiguration_appconfiguration,
	},

	"applicationinsights": {
		policy.Microsoft_Insights_webTests_API,
	},

	"appplatform": {
		policy.Microsoft_AppPlatform_appplatform,
	},

	"attestation": {
		policy.Microsoft_Attestation_attestation,
	},

	"authorization": {
		policy.Microsoft_Authorization_authorization_RoleDefinitionsCalls,
		policy.Microsoft_Authorization_authorization_DenyAssignmentCalls,
		policy.Microsoft_Authorization_authorization_RoleAssignmentsCalls,
		policy.Microsoft_Authorization_authorization_ProviderOperationsCalls,
	},

	"automanage": {
		policy.Microsoft_Automanage_automanage,
	},

	"automation": {
		policy.Microsoft_Automation_python2package,
		policy.Microsoft_Automation_sourceControl,
		policy.Microsoft_Automation_connection,
		policy.Microsoft_Automation_credential,
		policy.Microsoft_Automation_softwareUpdateConfigurationMachineRun,
		policy.Microsoft_Automation_variable,
		policy.Microsoft_Automation_account,
		policy.Microsoft_Automation_certificate,
		policy.Microsoft_Automation_linkedWorkspace,
		policy.Microsoft_Automation_python3package,
		policy.Microsoft_Automation_connectionType,
		policy.Microsoft_Automation_job,
		policy.Microsoft_Automation_jobSchedule,
		policy.Microsoft_Automation_schedule,
		policy.Microsoft_Automation_softwareUpdateConfigurationRun,
		policy.Microsoft_Automation_module,
		policy.Microsoft_Automation_dscNodeConfiguration,
		policy.Microsoft_Automation_runbook,
		policy.Microsoft_Automation_hybridRunbookWorkerGroup,
		policy.Microsoft_Automation_sourceControlSyncJobStreams,
		policy.Microsoft_Automation_sourceControlSyncJob,
		policy.Microsoft_Automation_hybridRunbookWorker,
		policy.Microsoft_Automation_operations,
		policy.Microsoft_Automation_dscConfiguration,
	},

	"azsadmin": {},

	"azure-kusto": {
		policy.Microsoft_Kusto_kusto,
	},

	"azureactivedirectory": {
		policy.Microsoft_Aadiam_privateLinkForAzureAD,
		policy.Microsoft_Aadiam_privateLinkResources,
		policy.Microsoft_Aadiam_privateEndpointConnections,
	},

	"azurearcdata": {
		policy.Microsoft_AzureArcData_dataControllers,
		policy.Microsoft_AzureArcData_common,
		policy.Microsoft_AzureArcData_azurearcdata,
		policy.Microsoft_AzureArcData_sqlManagedInstances,
		policy.Microsoft_AzureArcData_sqlServerInstances,
		policy.Microsoft_AzureArcData_operations,
	},

	"azuredata": {},

	"azurestack": {
		policy.Microsoft_AzureStack_CustomerSubscription,
		policy.Microsoft_AzureStack_Registration,
		policy.Microsoft_AzureStack_Product,
		policy.Microsoft_AzureStack_AzureStack,
	},

	"azurestackhci": {
		policy.Microsoft_AzureStackHCI_offers,
		policy.Microsoft_AzureStackHCI_publishers,
		policy.Microsoft_AzureStackHCI_updates,
		policy.Microsoft_AzureStackHCI_extensions,
		policy.Microsoft_AzureStackHCI_skus,
		policy.Microsoft_AzureStackHCI_updateRuns,
		policy.Microsoft_AzureStackHCI_updateSummaries,
		policy.Microsoft_AzureStackHCI_clusters,
		policy.Microsoft_AzureStackHCI_operations,
		policy.Microsoft_AzureStackHCI_arcSettings,
	},

	"baremetalinfrastructure": {
		policy.Microsoft_BareMetalInfrastructure_baremetalinfrastructure,
	},

	"batch": {
		policy.Microsoft_Batch_BatchManagement,
	},

	"billing": {
		policy.Microsoft_Billing_types,
		policy.Microsoft_Billing_billingSubscription,
		policy.Microsoft_Billing_payment,
	},

	"billingbenefits": {
		policy.Microsoft_BillingBenefits_billingbenefits,
	},

	"blockchain": {},

	"blueprint": {},

	"botservice": {
		policy.Microsoft_BotService_privatelinks,
		policy.Microsoft_BotService_botservice,
	},

	"cdn": {
		policy.Microsoft_Cdn_afdx,
		policy.Microsoft_Cdn_cdnwebapplicationfirewall,
		policy.Microsoft_Cdn_cdn,
	},

	"changeanalysis": {
		policy.Microsoft_ChangeAnalysis_changeanalysis,
	},

	"chaos": {},

	"cloudshell": {
		policy.Microsoft_Portal_CloudShell,
	},

	"cognitiveservices": {
		policy.Microsoft_CognitiveServices_cognitiveservices,
	},

	"commerce": {},

	"common-types": {},

	"communication": {
		policy.Microsoft_Communication_CommunicationService,
	},

	"compute": {},

	"confidentialledger": {
		policy.Microsoft_ConfidentialLedger_confidentialledger,
	},

	"confluent": {
		policy.Microsoft_Confluent_confluent,
	},

	"connectedvmware": {},

	"consumption": {
		policy.Microsoft_Consumption_consumption,
	},

	"containerinstance": {
		policy.Microsoft_ContainerInstance_containerInstance,
	},

	"containerregistry": {
		policy.Microsoft_ContainerRegistry_containerregistry,
	},

	"containerservice": {
		policy.Microsoft_ContainerService_managedClusters,
	},

	"cosmos-db": {
		policy.Microsoft_DocumentDB_restorable,
		policy.Microsoft_DocumentDB_services,
		policy.Microsoft_DocumentDB_managedCassandra,
		policy.Microsoft_DocumentDB_rbac,
		policy.Microsoft_DocumentDB_cosmos_db,
		policy.Microsoft_DocumentDB_notebook,
		policy.Microsoft_DocumentDB_privateEndpointConnection,
		policy.Microsoft_DocumentDB_privateLinkResources,
		policy.Microsoft_DocumentDB_mongorbac,
	},

	"cost-management": {
		policy.Microsoft_CostManagement_common_types,
		policy.Microsoft_CostManagement_costmanagement_generatecostdetailsreport,
		policy.Microsoft_CostManagement_scheduledActions,
		policy.Microsoft_CostManagement_costmanagement_exports,
		policy.Microsoft_CostManagement_costmanagement_generatedetailedcostreport,
		policy.Microsoft_CostManagement_costmanagement_benefits,
		policy.Microsoft_CostManagement_costmanagement,
		policy.Microsoft_CostManagement_costmanagement_pricesheets,
	},

	"cpim": {
		policy.Microsoft_AzureActiveDirectory_externalIdentities,
	},

	"customer-insights": {
		policy.Microsoft_CustomerInsights_customer_insights,
	},

	"customerlockbox": {},

	"customproviders": {},

	"dashboard": {
		policy.Microsoft_Dashboard_grafana,
	},

	"databox": {
		policy.Microsoft_DataBox_databox,
	},

	"databoxedge": {
		policy.Microsoft_DataBoxEdge_databoxedge,
	},

	"databricks": {
		policy.Microsoft_Databricks_vnetpeering,
		policy.Microsoft_Databricks_databricks,
	},

	"datacatalog": {
		policy.Microsoft_DataCatalog_datacatalog,
	},

	"datadog": {
		policy.Microsoft_Datadog_datadog,
	},

	"datafactory": {
		policy.Microsoft_DataFactory_datafactory,
	},

	"datalake-analytics": {
		policy.Microsoft_DataLakeAnalytics_account,
	},

	"datalake-store": {
		policy.Microsoft_DataLakeStore_account,
	},

	"datamigration": {
		policy.Microsoft_DataMigration_datamigration,
	},

	"dataprotection": {
		policy.Microsoft_DataProtection_dataprotection,
	},

	"datashare": {
		policy.Microsoft_DataShare_DataShare,
	},

	"deploymentmanager": {},

	"desktopvirtualization": {
		policy.Microsoft_DesktopVirtualization_desktopvirtualization,
	},

	"devcenter": {},

	"developerhub": {},

	"deviceprovisioningservices": {
		policy.Microsoft_Devices_iotdps,
	},

	"deviceupdate": {
		policy.Microsoft_DeviceUpdate_privatelinks,
		policy.Microsoft_DeviceUpdate_deviceupdate,
	},

	"devops": {},

	"devspaces": {
		policy.Microsoft_DevSpaces_devspaces,
	},

	"devtestlabs": {
		policy.Microsoft_DevTestLab_DTL,
	},

	"dfp": {},

	"digitaltwins": {
		policy.Microsoft_DigitalTwins_digitaltwins,
	},

	"dnc": {
		policy.Microsoft_DelegatedNetwork_common_types,
		policy.Microsoft_DelegatedNetwork_delegatedSubnets,
		policy.Microsoft_DelegatedNetwork_controller,
		policy.Microsoft_DelegatedNetwork_orchestrators,
		policy.Microsoft_DelegatedNetwork_operations,
	},

	"dns": {
		policy.Microsoft_Network_dns,
	},

	"dnsresolver": {
		policy.Microsoft_Network_dnsresolver,
	},

	"domainservices": {
		policy.Microsoft_AAD_oucontainer,
		policy.Microsoft_AAD_domainservices,
	},

	"dynamicstelemetry": {},

	"dynatrace": {
		policy.Dynatrace_Observability_dynatrace,
	},

	"edgeorder": {
		policy.Microsoft_EdgeOrder_edgeorder,
	},

	"edgeorderpartner": {},

	"education": {},

	"elastic": {
		policy.Microsoft_Elastic_elastic,
	},

	"elasticsan": {},

	"engagementfabric": {},

	"eventgrid": {
		policy.Microsoft_EventGrid_EventGrid,
	},

	"eventhub": {},

	"extendedlocation": {
		policy.Microsoft_ExtendedLocation_customlocations,
	},

	"fluidrelay": {
		policy.Microsoft_FluidRelay_fluidrelay,
	},

	"frontdoor": {
		policy.Microsoft_Network_webapplicationfirewall,
		policy.Microsoft_Network_network,
	},

	"graphrbac": {},

	"guestconfiguration": {
		policy.Microsoft_GuestConfiguration_guestconfiguration,
	},

	"hanaonazure": {},

	"hardwaresecuritymodules": {
		policy.Microsoft_HardwareSecurityModules_dedicatedhsm,
	},

	"hdinsight": {
		policy.Microsoft_HDInsight_locations,
		policy.Microsoft_HDInsight_scriptActions,
		policy.Microsoft_HDInsight_virtualMachines,
		policy.Microsoft_HDInsight_extensions,
		policy.Microsoft_HDInsight_cluster,
		policy.Microsoft_HDInsight_privateLinkResources,
		policy.Microsoft_HDInsight_privateEndpointConnections,
		policy.Microsoft_HDInsight_applications,
		policy.Microsoft_HDInsight_operations,
		policy.Microsoft_HDInsight_configurations,
	},

	"healthbot": {
		policy.Microsoft_HealthBot_healthbot,
	},

	"healthcareapis": {
		policy.Microsoft_HealthcareApis_healthcare_apis,
	},

	"hybridaks": {},

	"hybridcompute": {
		policy.Microsoft_HybridCompute_privateLinkScopes,
		policy.Microsoft_HybridCompute_HybridCompute,
	},

	"hybridconnectivity": {},

	"hybriddatamanager": {
		policy.Microsoft_HybridData_hybriddata,
	},

	"hybridkubernetes": {
		policy.Microsoft_Kubernetes_connectedClusters,
	},

	"hybridnetwork": {
		policy.Microsoft_HybridNetwork_vendorNetworkFunction,
		policy.Microsoft_HybridNetwork_vendor,
		policy.Microsoft_HybridNetwork_common,
		policy.Microsoft_HybridNetwork_device,
		policy.Microsoft_HybridNetwork_networkFunction,
		policy.Microsoft_HybridNetwork_operation,
		policy.Microsoft_HybridNetwork_networkFunctionVendor,
	},

	"imagebuilder": {
		policy.Microsoft_VirtualMachineImages_imagebuilder,
	},

	"imds": {},

	"intune": {},

	"iotcentral": {
		policy.Microsoft_IoTCentral_iotcentral,
	},

	"iothub": {
		policy.Microsoft_Devices_iothub,
	},

	"iotsecurity": {},

	"iotspaces": {},

	"keyvault": {
		policy.Microsoft_KeyVault_keyvault,
		policy.Microsoft_KeyVault_common,
		policy.Microsoft_KeyVault_providers,
		policy.Microsoft_KeyVault_secrets,
		policy.Microsoft_KeyVault_keys,
		policy.Microsoft_KeyVault_managedHsm,
	},

	"kubernetesconfiguration": {
		policy.Microsoft_KubernetesConfiguration_kubernetesconfiguration,
		policy.Microsoft_KubernetesConfiguration_fluxconfiguration,
		policy.Microsoft_KubernetesConfiguration_extensions,
		policy.Microsoft_KubernetesConfiguration_operations,
	},

	"labservices": {
		policy.Microsoft_LabServices_LabPlans,
		policy.Microsoft_LabServices_LabServices,
		policy.Microsoft_LabServices_Images,
		policy.Microsoft_LabServices_Users,
		policy.Microsoft_LabServices_Labs,
		policy.Microsoft_LabServices_VirtualMachines,
		policy.Microsoft_LabServices_OperationResults,
		policy.Microsoft_LabServices_Skus,
		policy.Microsoft_LabServices_Schedules,
		policy.Microsoft_LabServices_Usages,
	},

	"liftrqumulo": {},

	"loadtestservice": {
		policy.Microsoft_LoadTestService_loadtestservice,
	},

	"logic": {
		policy.Microsoft_Logic_logic,
	},

	"logz": {
		policy.Microsoft_Logz_logz,
	},

	"m365securityandcompliance": {},

	"machinelearning": {
		policy.Microsoft_MachineLearning_workspaces,
	},

	"machinelearningcompute": {},

	"machinelearningexperimentation": {},

	"machinelearningservices": {
		policy.Microsoft_MachineLearningServices_machineLearningServices,
		policy.Microsoft_MachineLearningServices_mfe,
		policy.Microsoft_MachineLearningServices_workspaceFeatures,
	},

	"maintenance": {
		policy.Microsoft_Maintenance_Maintenance,
	},

	"managednetwork": {},

	"managedservices": {
		policy.Microsoft_ManagedServices_managedservices,
	},

	"managementgroups": {
		policy.Microsoft_Management_management,
	},

	"managementpartner": {},

	"maps": {
		policy.Microsoft_Maps_maps_management,
	},

	"mariadb": {
		policy.Microsoft_DBforMariaDB_Servers,
	},

	"marketplace": {
		policy.Microsoft_Marketplace_Marketplace,
	},

	"marketplacecatalog": {
		policy.Microsoft_Marketplace_marketplacecatalog,
		policy.Search_search,
	},

	"marketplacenotifications": {
		policy.Microsoft_MarketplaceNotifications_MarketplaceNotifications,
	},

	"marketplaceordering": {
		policy.Microsoft_MarketplaceOrdering_Agreements,
	},

	"mediaservices": {},

	"migrate": {
		policy.Microsoft_Migrate_hubmigrate,
		policy.Microsoft_OffAzure_migrate,
	},

	"migrateprojects": {},

	"mixedreality": {
		policy.Microsoft_MixedReality_proxy,
		policy.Microsoft_MixedReality_common,
		policy.Microsoft_MixedReality_remote_rendering,
	},

	"mobilenetwork": {
		policy.Microsoft_MobileNetwork_dataNetwork,
		policy.Microsoft_MobileNetwork_attachedDataNetwork,
		policy.Microsoft_MobileNetwork_site,
		policy.Microsoft_MobileNetwork_mobileNetwork,
		policy.Microsoft_MobileNetwork_common,
		policy.Microsoft_MobileNetwork_service,
		policy.Microsoft_MobileNetwork_slice,
		policy.Microsoft_MobileNetwork_packetCoreControlPlane,
		policy.Microsoft_MobileNetwork_simPolicy,
		policy.Microsoft_MobileNetwork_operation,
		policy.Microsoft_MobileNetwork_ts29571,
		policy.Microsoft_MobileNetwork_packetCoreDataPlane,
		policy.Microsoft_MobileNetwork_simGroup,
		policy.Microsoft_MobileNetwork_sim,
	},

	"monitor": {
		policy.Microsoft_Insights_autoscale_API,
	},

	"msi": {
		policy.Microsoft_ManagedIdentity_ManagedIdentity,
	},

	"mysql": {
		policy.Microsoft_DBforMySQL_mysql,
	},

	"netapp": {
		policy.Microsoft_NetApp_netapp,
	},

	"network": {
		policy.Microsoft_Network_networkManagerConnection,
		policy.Microsoft_Network_networkInterface,
		policy.Microsoft_Network_availableServiceAliases,
		policy.Microsoft_Network_networkManagerGroup,
		policy.Microsoft_Network_usage,
		policy.Microsoft_Network_checkDnsAvailability,
		policy.Microsoft_Network_ddosProtectionPlan,
		policy.Microsoft_Network_customIpPrefix,
		policy.Microsoft_Network_cloudServicePublicIpAddress,
		policy.Microsoft_Network_expressRouteCircuit,
		policy.Microsoft_Network_routeTable,
		policy.Microsoft_Network_serviceTags,
		policy.Microsoft_Network_applicationGatewayWafDynamicManifests,
		policy.Microsoft_Network_availableDelegations,
		policy.Microsoft_Network_ddosCustomPolicy,
		policy.Microsoft_Network_expressRouteCrossConnection,
		policy.Microsoft_Network_networkManagerScopeConnection,
		policy.Microsoft_Network_publicIpAddress,
		policy.Microsoft_Network_networkManager,
		policy.Microsoft_Network_virtualNetworkGateway,
		policy.Microsoft_Network_networkManagerConnectivityConfiguration,
		policy.Microsoft_Network_ipGroups,
		policy.Microsoft_Network_webapplicationfirewall,
		policy.Microsoft_Network_networkSecurityGroup,
		policy.Microsoft_Network_networkVirtualAppliance,
		policy.Microsoft_Network_networkProfile,
		policy.Microsoft_Network_ipAllocation,
		policy.Microsoft_Network_endpointService,
		policy.Microsoft_Network_azureFirewallFqdnTag,
		policy.Microsoft_Network_natGateway,
		policy.Microsoft_Network_virtualNetwork,
		policy.Microsoft_Network_vmssNetworkInterface,
		policy.Microsoft_Network_dscpConfiguration,
		policy.Microsoft_Network_networkManagerEffectiveConfiguration,
		policy.Microsoft_Network_virtualWan,
		policy.Microsoft_Network_firewallPolicy,
		policy.Microsoft_Network_loadBalancer,
		policy.Microsoft_Network_network,
		policy.Microsoft_Network_cloudServiceNetworkInterface,
		policy.Microsoft_Network_publicIpPrefix,
		policy.Microsoft_Network_operation,
		policy.Microsoft_Network_vmssPublicIpAddress,
		policy.Microsoft_Network_networkManagerSecurityAdminConfiguration,
		policy.Microsoft_Network_azureWebCategory,
		policy.Microsoft_Network_securityPartnerProvider,
		policy.Microsoft_Network_serviceCommunity,
		policy.Microsoft_Network_cloudServiceSwap,
		policy.Microsoft_Network_expressRouteProviderPort,
		policy.Microsoft_Network_applicationGateway,
		policy.Microsoft_Network_expressRoutePort,
		policy.Microsoft_Network_privateEndpoint,
		policy.Microsoft_Network_applicationSecurityGroup,
		policy.Microsoft_Network_serviceEndpointPolicy,
		policy.Microsoft_Network_virtualRouter,
		policy.Microsoft_Network_networkManagerActiveConfiguration,
		policy.Microsoft_Network_virtualNetworkTap,
		policy.Microsoft_Network_azureFirewall,
		policy.Microsoft_Network_networkWatcher,
		policy.Microsoft_Network_privateLinkService,
		policy.Microsoft_Network_bastionHost,
		policy.Microsoft_Network_routeFilter,
	},

	"networkfunction": {
		policy.Microsoft_NetworkFunction_AzureTrafficCollector,
	},

	"nginx": {
		policy.NGINX_NGINXPLUS_swagger,
	},

	"notificationhubs": {
		policy.Microsoft_NotificationHubs_notificationhubs,
	},

	"oep": {},

	"operationalinsights": {
		policy.Microsoft_OperationalInsights_Tables,
		policy.Microsoft_OperationalInsights_Workspaces,
		policy.Microsoft_OperationalInsights_Operations,
	},

	"operationsmanagement": {},

	"orbital": {
		policy.Microsoft_Orbital_orbital,
	},

	"peering": {
		policy.Microsoft_Peering_peering,
	},

	"policyinsights": {
		policy.Microsoft_PolicyInsights_attestations,
	},

	"portal": {},

	"portalservices": {},

	"postgresql": {},

	"postgresqlhsc": {},

	"powerbidedicated": {
		policy.Microsoft_PowerBIdedicated_autoScaleVCores,
		policy.Microsoft_PowerBIdedicated_powerbidedicated,
	},

	"powerbiembedded": {
		policy.Microsoft_PowerBI_powerbiembedded,
	},

	"powerbiprivatelinks": {
		policy.Microsoft_PowerBI_powerbiprivatelinks,
	},

	"powerplatform": {},

	"privatedns": {
		policy.Microsoft_Network_privatedns,
	},

	"providerhub": {
		policy.Microsoft_ProviderHub_providerhub,
	},

	"purview": {
		policy.Microsoft_Purview_purview,
	},

	"quantum": {},

	"quota": {
		policy.Microsoft_Quota_quota,
	},

	"recommendationsservice": {
		policy.Microsoft_RecommendationsService_recommendationsservice,
	},

	"recoveryservices": {
		policy.Microsoft_RecoveryServices_vaults,
		policy.Microsoft_RecoveryServices_registeredidentities,
		policy.Microsoft_RecoveryServices_vaultusages,
		policy.Microsoft_RecoveryServices_replicationusages,
	},

	"recoveryservicesbackup": {
		policy.Microsoft_RecoveryServices_bms,
	},

	"recoveryservicessiterecovery": {
		policy.Microsoft_RecoveryServices_service,
	},

	"redhatopenshift": {
		policy.Microsoft_RedHatOpenShift_redhatopenshift,
	},

	"redis": {
		policy.Microsoft_Cache_redis,
	},

	"redisenterprise": {
		policy.Microsoft_Cache_redisenterprise,
	},

	"relay": {
		policy.Microsoft_Relay_authorizationRules,
		policy.Microsoft_Relay_wcfRelays,
		policy.Microsoft_Relay_NetworkRuleSets,
		policy.Microsoft_Relay_Namespaces,
		policy.Microsoft_Relay_hybridConnections,
		policy.Microsoft_Relay_operations,
	},

	"reservations": {
		policy.Microsoft_Capacity_reservations,
	},

	"resourceconnector": {
		policy.Microsoft_ResourceConnector_appliances,
	},

	"resourcegraph": {
		policy.Microsoft_ResourceGraph_resourcegraph,
	},

	"resourcehealth": {
		policy.Microsoft_ResourceHealth_ResourceHealth,
	},

	"resourcemover": {
		policy.Microsoft_Migrate_resourcemovercollection,
	},

	"resources": {
		policy.Microsoft_Features_features,
		policy.Microsoft_Features_SubscriptionFeatureRegistration,
		policy.Microsoft_Authorization_policyAssignments,
		policy.Microsoft_Resources_resources,
		policy.Microsoft_Solutions_managedapplications,
	},

	"riskiq": {},

	"saas": {},

	"scheduler": {
		policy.Microsoft_Scheduler_scheduler,
	},

	"schemaregistry": {},

	"scvmm": {},

	"search": {
		policy.Microsoft_Search_search,
	},

	"security": {
		policy.Microsoft_Security_settings,
	},

	"securityandcompliance": {
		policy.Microsoft_SecurityAndCompliance_common_types,
		policy.Microsoft_SecurityAndCompliance_privateLinkServicesForSCCPowershell,
		policy.Microsoft_SecurityAndCompliance_privateLinkServicesForEDMUpload,
		policy.Microsoft_SecurityAndCompliance_privateLinkServicesForM365ComplianceCenter,
		policy.Microsoft_SecurityAndCompliance_privateLinkServicesForMIPPolicySync,
		policy.Microsoft_SecurityAndCompliance_privateLinkServicesForM365SecurityCenter,
		policy.Microsoft_SecurityAndCompliance_privateLinkServicesForO365ManagementActivityAPI,
	},

	"securitydevops": {},

	"securityinsights": {},

	"serialconsole": {
		policy.Microsoft_SerialConsole_serialconsole,
		policy.Microsoft_SerialConsole_serialport,
	},

	"service-map": {},

	"servicebus": {
		policy.Microsoft_ServiceBus_AuthorizationRules,
		policy.Microsoft_ServiceBus_migrationconfigs,
		policy.Microsoft_ServiceBus_CheckNameAvailability,
		policy.Microsoft_ServiceBus_DisasterRecoveryConfig,
		policy.Microsoft_ServiceBus_namespace_preview,
		policy.Microsoft_ServiceBus_subscriptions,
		policy.Microsoft_ServiceBus_Queue,
		policy.Microsoft_ServiceBus_Rules,
		policy.Microsoft_ServiceBus_networksets,
		policy.Microsoft_ServiceBus_operations,
		policy.Microsoft_ServiceBus_topics,
	},

	"servicefabric": {
		policy.Microsoft_ServiceFabric_cluster,
		policy.Microsoft_ServiceFabric_application,
	},

	"servicefabricmanagedclusters": {
		policy.Microsoft_ServiceFabric_nodetype,
		policy.Microsoft_ServiceFabric_managedapplication,
		policy.Microsoft_ServiceFabric_managedcluster,
	},

	"servicefabricmesh": {},

	"servicelinker": {
		policy.Microsoft_ServiceLinker_servicelinker,
	},

	"servicenetworking": {},

	"signalr": {
		policy.Microsoft_SignalRService_signalr,
	},

	"softwareplan": {
		policy.Microsoft_SoftwarePlan_softwareplan,
	},

	"solutions": {
		policy.Microsoft_Solutions_managedapplications,
	},

	"sql": {},

	"sqlvirtualmachine": {
		policy.Microsoft_SqlVirtualMachine_sqlvm,
	},

	"storage": {
		policy.Microsoft_Storage_blob,
		policy.Microsoft_Storage_privatelinks,
		policy.Microsoft_Storage_common,
		policy.Microsoft_Storage_file,
		policy.Microsoft_Storage_queue,
		policy.Microsoft_Storage_table,
		policy.Microsoft_Storage_storage,
	},

	"storagecache": {
		policy.Microsoft_StorageCache_storagecache,
	},

	"storageimportexport": {
		policy.Microsoft_ImportExport_storageimportexport,
	},

	"storagemover": {
		policy.Microsoft_StorageMover_storagemover,
	},

	"storagepool": {
		policy.Microsoft_StoragePool_storagepool,
	},

	"storagesync": {
		policy.Microsoft_StorageSync_storagesync,
	},

	"storsimple8000series": {
		policy.Microsoft_StorSimple_storsimple,
	},

	"streamanalytics": {
		policy.Microsoft_StreamAnalytics_transformations,
		policy.Microsoft_StreamAnalytics_privateEndpoints,
		policy.Microsoft_StreamAnalytics_streamingjobs,
		policy.Microsoft_StreamAnalytics_outputs,
		policy.Microsoft_StreamAnalytics_subscriptions,
		policy.Microsoft_StreamAnalytics_inputs,
		policy.Microsoft_StreamAnalytics_clusters,
		policy.Microsoft_StreamAnalytics_functions,
	},

	"subscription": {
		policy.Microsoft_Subscription_subscriptions,
	},

	"support": {
		policy.Microsoft_Support_support,
	},

	"synapse": {
		policy.Microsoft_Synapse_integrationRuntime,
		policy.Microsoft_Synapse_workspace,
		policy.Microsoft_Synapse_checkNameAvailability,
		policy.Microsoft_Synapse_azureADOnlyAuthentication,
		policy.Microsoft_Synapse_sqlPool,
		policy.Microsoft_Synapse_library,
		policy.Microsoft_Synapse_keys,
		policy.Microsoft_Synapse_privateLinkResources,
		policy.Microsoft_Synapse_privatelinkhub,
		policy.Microsoft_Synapse_privateEndpointConnections,
		policy.Microsoft_Synapse_sqlServer,
		policy.Microsoft_Synapse_firewallRule,
		policy.Microsoft_Synapse_bigDataPool,
		policy.Microsoft_Synapse_operations,
	},

	"syntex": {},

	"testbase": {},

	"timeseriesinsights": {
		policy.Microsoft_TimeSeriesInsights_timeseriesinsights,
	},

	"trafficmanager": {
		policy.Microsoft_Network_trafficmanager,
	},

	"vi": {
		policy.Microsoft_VideoIndexer_vi,
	},

	"videoanalyzer": {},

	"visualstudio": {},

	"vmware": {
		policy.Microsoft_AVS_vmware,
	},

	"vmwarecloudsimple": {
		policy.Microsoft_VMwareCloudSimple_vmwarecloudsimple,
	},

	"voiceservices": {
		policy.Microsoft_VoiceServices_voiceservices,
	},

	"web": {
		policy.Microsoft_Web_ContainerApps,
		policy.Microsoft_Web_ResourceHealthMetadata,
		policy.Microsoft_Web_ResourceProvider,
		policy.Microsoft_Web_DeletedWebApps,
		policy.Microsoft_Web_Certificates,
		policy.Microsoft_Web_AppServiceEnvironments,
		policy.Microsoft_Web_Diagnostics,
		policy.Microsoft_Web_AppServicePlans,
		policy.Microsoft_Web_StaticSites,
		policy.Microsoft_Web_Global,
		policy.Microsoft_Web_Provider,
		policy.Microsoft_Web_KubeEnvironments,
		policy.Microsoft_Web_ContainerAppsRevisions,
		policy.Microsoft_Web_Recommendations,
		policy.Microsoft_Web_CommonDefinitions,
		policy.Microsoft_Web_WebApps,
		policy.Microsoft_DomainRegistration_TopLevelDomains,
		policy.Microsoft_DomainRegistration_DomainRegistrationProvider,
		policy.Microsoft_DomainRegistration_Domains,
		policy.Microsoft_CertificateRegistration_CertificateRegistrationProvider,
		policy.Microsoft_CertificateRegistration_AppServiceCertificateOrders,
		policy.Microsoft_CertificateRegistration_CertificateOrdersDiagnostics,
	},

	"webpubsub": {
		policy.Microsoft_SignalRService_webpubsub,
	},

	"windowsesu": {},

	"windowsiot": {
		policy.Microsoft_WindowsIoT_WindowsIotServices,
	},

	"workloadmonitor": {},

	"workloads": {},
}
