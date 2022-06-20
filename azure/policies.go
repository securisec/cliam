package azure

import "github.com/securisec/cliam/azure/policy"

func GetPolicyKeys() []string {
	hold := make([]string, 0, len(Policies))
	for k := range Policies {
		hold = append(hold, k)
	}
	return hold
}

var Policies = map[string]map[string]policy.Policy{
	"Microsoft.Web.WebApps":                                                           policy.Microsoft_Web_WebApps,
	"Microsoft.Web.StaticSites":                                                       policy.Microsoft_Web_StaticSites,
	"Microsoft.Web.ResourceProvider":                                                  policy.Microsoft_Web_ResourceProvider,
	"Microsoft.Web.ResourceHealthMetadata":                                            policy.Microsoft_Web_ResourceHealthMetadata,
	"Microsoft.Web.ContainerApps":                                                     policy.Microsoft_Web_ContainerApps,
	"Microsoft.Web.DeletedWebApps":                                                    policy.Microsoft_Web_DeletedWebApps,
	"Microsoft.Web.Certificates":                                                      policy.Microsoft_Web_Certificates,
	"Microsoft.Web.AppServiceEnvironments":                                            policy.Microsoft_Web_AppServiceEnvironments,
	"Microsoft.Web.Diagnostics":                                                       policy.Microsoft_Web_Diagnostics,
	"Microsoft.Web.AppServicePlans":                                                   policy.Microsoft_Web_AppServicePlans,
	"Microsoft.Web.Global":                                                            policy.Microsoft_Web_Global,
	"Microsoft.Web.Provider":                                                          policy.Microsoft_Web_Provider,
	"Microsoft.Web.KubeEnvironments":                                                  policy.Microsoft_Web_KubeEnvironments,
	"Microsoft.Web.ContainerAppsRevisions":                                            policy.Microsoft_Web_ContainerAppsRevisions,
	"Microsoft.Web.Recommendations":                                                   policy.Microsoft_Web_Recommendations,
	"Microsoft.Web.CommonDefinitions":                                                 policy.Microsoft_Web_CommonDefinitions,
	"Microsoft.Storage.storage":                                                       policy.Microsoft_Storage_storage,
	"Microsoft.Storage.blob":                                                          policy.Microsoft_Storage_blob,
	"Microsoft.Storage.privatelinks":                                                  policy.Microsoft_Storage_privatelinks,
	"Microsoft.Storage.common":                                                        policy.Microsoft_Storage_common,
	"Microsoft.Storage.file":                                                          policy.Microsoft_Storage_file,
	"Microsoft.Storage.queue":                                                         policy.Microsoft_Storage_queue,
	"Microsoft.Storage.table":                                                         policy.Microsoft_Storage_table,
	"Microsoft.ContainerRegistry.containerregistry":                                   policy.Microsoft_ContainerRegistry_containerregistry,
	"Microsoft.DocumentDB.restorable":                                                 policy.Microsoft_DocumentDB_restorable,
	"Microsoft.DocumentDB.managedCassandra":                                           policy.Microsoft_DocumentDB_managedCassandra,
	"Microsoft.DocumentDB.rbac":                                                       policy.Microsoft_DocumentDB_rbac,
	"Microsoft.DocumentDB.cosmos-db":                                                  policy.Microsoft_DocumentDB_cosmos_db,
	"Microsoft.DocumentDB.notebook":                                                   policy.Microsoft_DocumentDB_notebook,
	"Microsoft.DocumentDB.privateEndpointConnection":                                  policy.Microsoft_DocumentDB_privateEndpointConnection,
	"Microsoft.DocumentDB.privateLinkResources":                                       policy.Microsoft_DocumentDB_privateLinkResources,
	"Microsoft.Aadiam.privateLinkForAzureAD":                                          policy.Microsoft_Aadiam_privateLinkForAzureAD,
	"Microsoft.Aadiam.privateLinkResources":                                           policy.Microsoft_Aadiam_privateLinkResources,
	"Microsoft.Aadiam.privateEndpointConnections":                                     policy.Microsoft_Aadiam_privateEndpointConnections,
	"Microsoft.Cache.redis":                                                           policy.Microsoft_Cache_redis,
	"Microsoft.Compute.runCommands":                                                   policy.Microsoft_Compute_runCommands,
	"Microsoft.Compute.communityGallery":                                              policy.Microsoft_Compute_communityGallery,
	"Microsoft.Compute.gallery":                                                       policy.Microsoft_Compute_gallery,
	"Microsoft.Compute.skus":                                                          policy.Microsoft_Compute_skus,
	"Microsoft.Compute.sharedGallery":                                                 policy.Microsoft_Compute_sharedGallery,
	"Microsoft.Compute.compute":                                                       policy.Microsoft_Compute_compute,
	"Microsoft.Portal.CloudShell":                                                     policy.Microsoft_Portal_CloudShell,
	"Microsoft.SignalRService.webpubsub":                                              policy.Microsoft_SignalRService_webpubsub,
	"Microsoft.AVS.vmware":                                                            policy.Microsoft_AVS_vmware,
	"Microsoft.DBforPostgreSQL.PrivateDnsZone":                                        policy.Microsoft_DBforPostgreSQL_PrivateDnsZone,
	"Microsoft.DBforPostgreSQL.Databases":                                             policy.Microsoft_DBforPostgreSQL_Databases,
	"Microsoft.DBforPostgreSQL.postgresql":                                            policy.Microsoft_DBforPostgreSQL_postgresql,
	"Microsoft.DBforMySQL.mysql":                                                      policy.Microsoft_DBforMySQL_mysql,
	"Microsoft.DBforMariaDB.Servers":                                                  policy.Microsoft_DBforMariaDB_Servers,
	"Microsoft.DevOps.devops":                                                         policy.Microsoft_DevOps_devops,
	"Microsoft.ApiManagement.apimskus":                                                policy.Microsoft_ApiManagement_apimskus,
	"Microsoft.ApiManagement.apimreports":                                             policy.Microsoft_ApiManagement_apimreports,
	"Microsoft.ApiManagement.apimapisByTags":                                          policy.Microsoft_ApiManagement_apimapisByTags,
	"Microsoft.ApiManagement.apimprivatelink":                                         policy.Microsoft_ApiManagement_apimprivatelink,
	"Microsoft.ApiManagement.apimnamedvalues":                                         policy.Microsoft_ApiManagement_apimnamedvalues,
	"Microsoft.ApiManagement.apimdeployment":                                          policy.Microsoft_ApiManagement_apimdeployment,
	"Microsoft.ApiManagement.apimcontenttypes":                                        policy.Microsoft_ApiManagement_apimcontenttypes,
	"Microsoft.ApiManagement.apimproducts":                                            policy.Microsoft_ApiManagement_apimproducts,
	"Microsoft.ApiManagement.apimtagresources":                                        policy.Microsoft_ApiManagement_apimtagresources,
	"Microsoft.ApiManagement.apimopenidconnectproviders":                              policy.Microsoft_ApiManagement_apimopenidconnectproviders,
	"Microsoft.ApiManagement.apimtenant":                                              policy.Microsoft_ApiManagement_apimtenant,
	"Microsoft.ApiManagement.apimanagement":                                           policy.Microsoft_ApiManagement_apimanagement,
	"Microsoft.ApiManagement.apimemailtemplates":                                      policy.Microsoft_ApiManagement_apimemailtemplates,
	"Microsoft.ApiManagement.apimapis":                                                policy.Microsoft_ApiManagement_apimapis,
	"Microsoft.ApiManagement.apimauthorizationservers":                                policy.Microsoft_ApiManagement_apimauthorizationservers,
	"Microsoft.ApiManagement.apimnotifications":                                       policy.Microsoft_ApiManagement_apimnotifications,
	"Microsoft.ApiManagement.apimtags":                                                policy.Microsoft_ApiManagement_apimtags,
	"Microsoft.ApiManagement.apimloggers":                                             policy.Microsoft_ApiManagement_apimloggers,
	"Microsoft.ApiManagement.apimquotas":                                              policy.Microsoft_ApiManagement_apimquotas,
	"Microsoft.ApiManagement.apimportalsettings":                                      policy.Microsoft_ApiManagement_apimportalsettings,
	"Microsoft.ApiManagement.apimsubscriptions":                                       policy.Microsoft_ApiManagement_apimsubscriptions,
	"Microsoft.ApiManagement.apimconnectivitycheck":                                   policy.Microsoft_ApiManagement_apimconnectivitycheck,
	"Microsoft.ApiManagement.apimportalrevisions":                                     policy.Microsoft_ApiManagement_apimportalrevisions,
	"Microsoft.ApiManagement.apimoutbounddependency":                                  policy.Microsoft_ApiManagement_apimoutbounddependency,
	"Microsoft.ApiManagement.apimidentityprovider":                                    policy.Microsoft_ApiManagement_apimidentityprovider,
	"Microsoft.ApiManagement.apimregions":                                             policy.Microsoft_ApiManagement_apimregions,
	"Microsoft.ApiManagement.apimissues":                                              policy.Microsoft_ApiManagement_apimissues,
	"Microsoft.ApiManagement.apimcaches":                                              policy.Microsoft_ApiManagement_apimcaches,
	"Microsoft.ApiManagement.apimpolicydescriptions":                                  policy.Microsoft_ApiManagement_apimpolicydescriptions,
	"Microsoft.ApiManagement.apimpolicies":                                            policy.Microsoft_ApiManagement_apimpolicies,
	"Microsoft.ApiManagement.apimdiagnostics":                                         policy.Microsoft_ApiManagement_apimdiagnostics,
	"Microsoft.ApiManagement.apimschema":                                              policy.Microsoft_ApiManagement_apimschema,
	"Microsoft.ApiManagement.apimsettings":                                            policy.Microsoft_ApiManagement_apimsettings,
	"Microsoft.ApiManagement.apimgateways":                                            policy.Microsoft_ApiManagement_apimgateways,
	"Microsoft.ApiManagement.apimgroups":                                              policy.Microsoft_ApiManagement_apimgroups,
	"Microsoft.ApiManagement.definitions":                                             policy.Microsoft_ApiManagement_definitions,
	"Microsoft.ApiManagement.apimdeletedservices":                                     policy.Microsoft_ApiManagement_apimdeletedservices,
	"Microsoft.ApiManagement.apimbackends":                                            policy.Microsoft_ApiManagement_apimbackends,
	"Microsoft.ApiManagement.apimusers":                                               policy.Microsoft_ApiManagement_apimusers,
	"Microsoft.ApiManagement.apimcertificates":                                        policy.Microsoft_ApiManagement_apimcertificates,
	"Microsoft.ApiManagement.apimnetworkstatus":                                       policy.Microsoft_ApiManagement_apimnetworkstatus,
	"Microsoft.ApiManagement.apimproductsByTags":                                      policy.Microsoft_ApiManagement_apimproductsByTags,
	"Microsoft.ApiManagement.apimapiversionsets":                                      policy.Microsoft_ApiManagement_apimapiversionsets,
	"Microsoft.ContainerService.containerService":                                     policy.Microsoft_ContainerService_containerService,
	"Microsoft.ContainerService.managedClusters":                                      policy.Microsoft_ContainerService_managedClusters,
	"Microsoft.ContainerInstance.containerInstance":                                   policy.Microsoft_ContainerInstance_containerInstance,
	"Microsoft.MachineLearningServices.machineLearningServices":                       policy.Microsoft_MachineLearningServices_machineLearningServices,
	"Microsoft.MachineLearningServices.mfe":                                           policy.Microsoft_MachineLearningServices_mfe,
	"Microsoft.MachineLearningServices.workspaceFeatures":                             policy.Microsoft_MachineLearningServices_workspaceFeatures,
	"Microsoft.Authorization.RoleManagementPolicy":                                    policy.Microsoft_Authorization_RoleManagementPolicy,
	"Microsoft.Authorization.RoleManagementPolicyAssignment":                          policy.Microsoft_Authorization_RoleManagementPolicyAssignment,
	"Microsoft.Authorization.RoleEligibilityScheduleInstance":                         policy.Microsoft_Authorization_RoleEligibilityScheduleInstance,
	"Microsoft.Authorization.common-types":                                            policy.Microsoft_Authorization_common_types,
	"Microsoft.Authorization.RoleEligibilitySchedule":                                 policy.Microsoft_Authorization_RoleEligibilitySchedule,
	"Microsoft.Authorization.RoleEligibilityScheduleRequest":                          policy.Microsoft_Authorization_RoleEligibilityScheduleRequest,
	"Microsoft.Authorization.RoleAssignmentScheduleInstance":                          policy.Microsoft_Authorization_RoleAssignmentScheduleInstance,
	"Microsoft.Authorization.RoleAssignmentScheduleRequest":                           policy.Microsoft_Authorization_RoleAssignmentScheduleRequest,
	"Microsoft.Authorization.RoleAssignmentSchedule":                                  policy.Microsoft_Authorization_RoleAssignmentSchedule,
	"Microsoft.Authorization.EligibleChildResources":                                  policy.Microsoft_Authorization_EligibleChildResources,
	"Microsoft.AzureStack.CustomerSubscription":                                       policy.Microsoft_AzureStack_CustomerSubscription,
	"Microsoft.AzureStack.Registration":                                               policy.Microsoft_AzureStack_Registration,
	"Microsoft.AzureStack.Product":                                                    policy.Microsoft_AzureStack_Product,
	"Microsoft.AzureStack.AzureStack":                                                 policy.Microsoft_AzureStack_AzureStack,
	"Microsoft.Network.dns":                                                           policy.Microsoft_Network_dns,
	"Microsoft.Elastic.elastic":                                                       policy.Microsoft_Elastic_elastic,
	"Microsoft.Features.features":                                                     policy.Microsoft_Features_features,
	"Microsoft.Features.SubscriptionFeatureRegistration":                              policy.Microsoft_Features_SubscriptionFeatureRegistration,
	"Microsoft.Authorization.policyAssignments":                                       policy.Microsoft_Authorization_policyAssignments,
	"Microsoft.Authorization.policySetDefinitions":                                    policy.Microsoft_Authorization_policySetDefinitions,
	"Microsoft.Authorization.policyDefinitions":                                       policy.Microsoft_Authorization_policyDefinitions,
	"Microsoft.Resources.changes":                                                     policy.Microsoft_Resources_changes,
	"Microsoft.Solutions.managedapplications":                                         policy.Microsoft_Solutions_managedapplications,
	"Microsoft.ManagedServices.managedservices":                                       policy.Microsoft_ManagedServices_managedservices,
	"Microsoft.Logz.logz":                                                             policy.Microsoft_Logz_logz,
	"Microsoft.KeyVault.keyvault":                                                     policy.Microsoft_KeyVault_keyvault,
	"Microsoft.KeyVault.common":                                                       policy.Microsoft_KeyVault_common,
	"Microsoft.KeyVault.providers":                                                    policy.Microsoft_KeyVault_providers,
	"Microsoft.KeyVault.secrets":                                                      policy.Microsoft_KeyVault_secrets,
	"Microsoft.KeyVault.keys":                                                         policy.Microsoft_KeyVault_keys,
	"Microsoft.KeyVault.managedHsm":                                                   policy.Microsoft_KeyVault_managedHsm,
	"Microsoft.VirtualMachineImages.imagebuilder":                                     policy.Microsoft_VirtualMachineImages_imagebuilder,
	"Microsoft.DevSpaces.devspaces":                                                   policy.Microsoft_DevSpaces_devspaces,
	"Microsoft.AppPlatform.appplatform":                                               policy.Microsoft_AppPlatform_appplatform,
	"Microsoft.App.ContainerApps":                                                     policy.Microsoft_App_ContainerApps,
	"Microsoft.App.AuthConfigs":                                                       policy.Microsoft_App_AuthConfigs,
	"Microsoft.App.SourceControls":                                                    policy.Microsoft_App_SourceControls,
	"Microsoft.App.ManagedEnvironments":                                               policy.Microsoft_App_ManagedEnvironments,
	"Microsoft.App.Global":                                                            policy.Microsoft_App_Global,
	"Microsoft.App.ContainerAppsRevisions":                                            policy.Microsoft_App_ContainerAppsRevisions,
	"Microsoft.App.CommonDefinitions":                                                 policy.Microsoft_App_CommonDefinitions,
	"Microsoft.App.DaprComponents":                                                    policy.Microsoft_App_DaprComponents,
	"Microsoft.App.ManagedEnvironmentsStorages":                                       policy.Microsoft_App_ManagedEnvironmentsStorages,
	"Microsoft.VMwareCloudSimple.vmwarecloudsimple":                                   policy.Microsoft_VMwareCloudSimple_vmwarecloudsimple,
	"Microsoft.Subscription.subscriptions":                                            policy.Microsoft_Subscription_subscriptions,
	"Microsoft.StorageCache.storagecache":                                             policy.Microsoft_StorageCache_storagecache,
	"Microsoft.SoftwarePlan.softwareplan":                                             policy.Microsoft_SoftwarePlan_softwareplan,
	"Microsoft.ServiceFabric.cluster":                                                 policy.Microsoft_ServiceFabric_cluster,
	"Microsoft.ServiceFabric.application":                                             policy.Microsoft_ServiceFabric_application,
	"Microsoft.ServiceBus.AuthorizationRules":                                         policy.Microsoft_ServiceBus_AuthorizationRules,
	"Microsoft.ServiceBus.migrationconfigs":                                           policy.Microsoft_ServiceBus_migrationconfigs,
	"Microsoft.ServiceBus.CheckNameAvailability":                                      policy.Microsoft_ServiceBus_CheckNameAvailability,
	"Microsoft.ServiceBus.DisasterRecoveryConfig":                                     policy.Microsoft_ServiceBus_DisasterRecoveryConfig,
	"Microsoft.ServiceBus.namespace-preview":                                          policy.Microsoft_ServiceBus_namespace_preview,
	"Microsoft.ServiceBus.subscriptions":                                              policy.Microsoft_ServiceBus_subscriptions,
	"Microsoft.ServiceBus.Queue":                                                      policy.Microsoft_ServiceBus_Queue,
	"Microsoft.ServiceBus.Rules":                                                      policy.Microsoft_ServiceBus_Rules,
	"Microsoft.ServiceBus.networksets":                                                policy.Microsoft_ServiceBus_networksets,
	"Microsoft.ServiceBus.operations":                                                 policy.Microsoft_ServiceBus_operations,
	"Microsoft.ServiceBus.topics":                                                     policy.Microsoft_ServiceBus_topics,
	"Microsoft.SerialConsole.serialconsole":                                           policy.Microsoft_SerialConsole_serialconsole,
	"Microsoft.SerialConsole.serialport":                                              policy.Microsoft_SerialConsole_serialport,
	"Microsoft.Search.search":                                                         policy.Microsoft_Search_search,
	"Microsoft.AppConfiguration.appconfiguration":                                     policy.Microsoft_AppConfiguration_appconfiguration,
	"Microsoft.Insights.workbooks_API":                                                policy.Microsoft_Insights_workbooks_API,
	"Microsoft.Automanage.automanage":                                                 policy.Microsoft_Automanage_automanage,
	"Microsoft.AzureArcData.dataControllers":                                          policy.Microsoft_AzureArcData_dataControllers,
	"Microsoft.AzureArcData.common":                                                   policy.Microsoft_AzureArcData_common,
	"Microsoft.AzureArcData.azurearcdata":                                             policy.Microsoft_AzureArcData_azurearcdata,
	"Microsoft.AzureArcData.sqlManagedInstances":                                      policy.Microsoft_AzureArcData_sqlManagedInstances,
	"Microsoft.AzureArcData.sqlServerInstances":                                       policy.Microsoft_AzureArcData_sqlServerInstances,
	"Microsoft.AzureArcData.operations":                                               policy.Microsoft_AzureArcData_operations,
	"Microsoft.Batch.BatchManagement":                                                 policy.Microsoft_Batch_BatchManagement,
	"Microsoft.BotService.botservice":                                                 policy.Microsoft_BotService_botservice,
	"Microsoft.Communication.CommunicationService":                                    policy.Microsoft_Communication_CommunicationService,
	"Microsoft.DesktopVirtualization.desktopvirtualization":                           policy.Microsoft_DesktopVirtualization_desktopvirtualization,
	"Microsoft.AAD.oucontainer":                                                       policy.Microsoft_AAD_oucontainer,
	"Microsoft.AAD.domainservices":                                                    policy.Microsoft_AAD_domainservices,
	"Microsoft.Network.webapplicationfirewall":                                        policy.Microsoft_Network_webapplicationfirewall,
	"Microsoft.Network.network":                                                       policy.Microsoft_Network_network,
	"Microsoft.GuestConfiguration.guestconfiguration":                                 policy.Microsoft_GuestConfiguration_guestconfiguration,
	"Microsoft.HybridCompute.privateLinkScopes":                                       policy.Microsoft_HybridCompute_privateLinkScopes,
	"Microsoft.HybridCompute.HybridCompute":                                           policy.Microsoft_HybridCompute_HybridCompute,
	"Microsoft.Kubernetes.connectedClusters":                                          policy.Microsoft_Kubernetes_connectedClusters,
	"Microsoft.HybridData.hybriddata":                                                 policy.Microsoft_HybridData_hybriddata,
	"Microsoft.HybridNetwork.vendorNetworkFunction":                                   policy.Microsoft_HybridNetwork_vendorNetworkFunction,
	"Microsoft.HybridNetwork.vendor":                                                  policy.Microsoft_HybridNetwork_vendor,
	"Microsoft.HybridNetwork.common":                                                  policy.Microsoft_HybridNetwork_common,
	"Microsoft.HybridNetwork.device":                                                  policy.Microsoft_HybridNetwork_device,
	"Microsoft.HybridNetwork.networkFunction":                                         policy.Microsoft_HybridNetwork_networkFunction,
	"Microsoft.HybridNetwork.operation":                                               policy.Microsoft_HybridNetwork_operation,
	"Microsoft.HybridNetwork.networkFunctionVendor":                                   policy.Microsoft_HybridNetwork_networkFunctionVendor,
	"Microsoft.MachineLearning.workspaces":                                            policy.Microsoft_MachineLearning_workspaces,
	"Microsoft.Management.management":                                                 policy.Microsoft_Management_management,
	"Microsoft.NetApp.netapp":                                                         policy.Microsoft_NetApp_netapp,
	"Microsoft.RecoveryServices.bms":                                                  policy.Microsoft_RecoveryServices_bms,
	"Microsoft.RecoveryServices.vaults":                                               policy.Microsoft_RecoveryServices_vaults,
	"Microsoft.RecoveryServices.registeredidentities":                                 policy.Microsoft_RecoveryServices_registeredidentities,
	"Microsoft.RecoveryServices.vaultusages":                                          policy.Microsoft_RecoveryServices_vaultusages,
	"Microsoft.RecoveryServices.replicationusages":                                    policy.Microsoft_RecoveryServices_replicationusages,
	"Microsoft.RedHatOpenShift.redhatopenshift":                                       policy.Microsoft_RedHatOpenShift_redhatopenshift,
	"Microsoft.Cache.redisenterprise":                                                 policy.Microsoft_Cache_redisenterprise,
	"Microsoft.ResourceGraph.resourcegraph":                                           policy.Microsoft_ResourceGraph_resourcegraph,
	"Microsoft.SecurityAndCompliance.common-types":                                    policy.Microsoft_SecurityAndCompliance_common_types,
	"Microsoft.SecurityAndCompliance.privateLinkServicesForSCCPowershell":             policy.Microsoft_SecurityAndCompliance_privateLinkServicesForSCCPowershell,
	"Microsoft.SecurityAndCompliance.privateLinkServicesForEDMUpload":                 policy.Microsoft_SecurityAndCompliance_privateLinkServicesForEDMUpload,
	"Microsoft.SecurityAndCompliance.privateLinkServicesForM365ComplianceCenter":      policy.Microsoft_SecurityAndCompliance_privateLinkServicesForM365ComplianceCenter,
	"Microsoft.SecurityAndCompliance.privateLinkServicesForMIPPolicySync":             policy.Microsoft_SecurityAndCompliance_privateLinkServicesForMIPPolicySync,
	"Microsoft.SecurityAndCompliance.privateLinkServicesForM365SecurityCenter":        policy.Microsoft_SecurityAndCompliance_privateLinkServicesForM365SecurityCenter,
	"Microsoft.SecurityAndCompliance.privateLinkServicesForO365ManagementActivityAPI": policy.Microsoft_SecurityAndCompliance_privateLinkServicesForO365ManagementActivityAPI,
	"Microsoft.StoragePool.storagepool":                                               policy.Microsoft_StoragePool_storagepool,
	"Microsoft.StorageSync.storagesync":                                               policy.Microsoft_StorageSync_storagesync,
}
