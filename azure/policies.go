package azure

import "github.com/securisec/cliam/azure/policy"

func GetPolicyKeys() []string {
	hold := make([]string, 0, len(Policies))
	for k := range Policies {
		hold = append(hold, k)
	}
	return hold
}

var Policies = map[string][]policy.Policy{
	"Microsoft.Web.WebApps":                          policy.Microsoft_Web_WebApps,
	"Microsoft.Web.StaticSites":                      policy.Microsoft_Web_StaticSites,
	"Microsoft.Web.ResourceProvider":                 policy.Microsoft_Web_ResourceProvider,
	"Microsoft.Web.ResourceHealthMetadata":           policy.Microsoft_Web_ResourceHealthMetadata,
	"Microsoft.Web.ContainerApps":                    policy.Microsoft_Web_ContainerApps,
	"Microsoft.Web.DeletedWebApps":                   policy.Microsoft_Web_DeletedWebApps,
	"Microsoft.Web.Certificates":                     policy.Microsoft_Web_Certificates,
	"Microsoft.Web.AppServiceEnvironments":           policy.Microsoft_Web_AppServiceEnvironments,
	"Microsoft.Web.Diagnostics":                      policy.Microsoft_Web_Diagnostics,
	"Microsoft.Web.AppServicePlans":                  policy.Microsoft_Web_AppServicePlans,
	"Microsoft.Web.Global":                           policy.Microsoft_Web_Global,
	"Microsoft.Web.Provider":                         policy.Microsoft_Web_Provider,
	"Microsoft.Web.KubeEnvironments":                 policy.Microsoft_Web_KubeEnvironments,
	"Microsoft.Web.ContainerAppsRevisions":           policy.Microsoft_Web_ContainerAppsRevisions,
	"Microsoft.Web.Recommendations":                  policy.Microsoft_Web_Recommendations,
	"Microsoft.Web.CommonDefinitions":                policy.Microsoft_Web_CommonDefinitions,
	"Microsoft.Storage.storage":                      policy.Microsoft_Storage_storage,
	"Microsoft.Storage.blob":                         policy.Microsoft_Storage_blob,
	"Microsoft.Storage.privatelinks":                 policy.Microsoft_Storage_privatelinks,
	"Microsoft.Storage.common":                       policy.Microsoft_Storage_common,
	"Microsoft.Storage.file":                         policy.Microsoft_Storage_file,
	"Microsoft.Storage.queue":                        policy.Microsoft_Storage_queue,
	"Microsoft.Storage.table":                        policy.Microsoft_Storage_table,
	"Microsoft.ContainerRegistry.containerregistry":  policy.Microsoft_ContainerRegistry_containerregistry,
	"Microsoft.DocumentDB.restorable":                policy.Microsoft_DocumentDB_restorable,
	"Microsoft.DocumentDB.managedCassandra":          policy.Microsoft_DocumentDB_managedCassandra,
	"Microsoft.DocumentDB.rbac":                      policy.Microsoft_DocumentDB_rbac,
	"Microsoft.DocumentDB.cosmos-db":                 policy.Microsoft_DocumentDB_cosmos_db,
	"Microsoft.DocumentDB.notebook":                  policy.Microsoft_DocumentDB_notebook,
	"Microsoft.DocumentDB.privateEndpointConnection": policy.Microsoft_DocumentDB_privateEndpointConnection,
	"Microsoft.DocumentDB.privateLinkResources":      policy.Microsoft_DocumentDB_privateLinkResources,
	"Microsoft.Aadiam.privateLinkForAzureAD":         policy.Microsoft_Aadiam_privateLinkForAzureAD,
	"Microsoft.Aadiam.privateLinkResources":          policy.Microsoft_Aadiam_privateLinkResources,
	"Microsoft.Aadiam.privateEndpointConnections":    policy.Microsoft_Aadiam_privateEndpointConnections,
	"Microsoft.Cache.redis":                          policy.Microsoft_Cache_redis,
	"Microsoft.ContainerService.containerService":    policy.Microsoft_ContainerService_containerService,
	"Microsoft.Compute.runCommands":                  policy.Microsoft_Compute_runCommands,
	"Microsoft.Compute.communityGallery":             policy.Microsoft_Compute_communityGallery,
	"Microsoft.Compute.gallery":                      policy.Microsoft_Compute_gallery,
	"Microsoft.Compute.skus":                         policy.Microsoft_Compute_skus,
	"Microsoft.Compute.sharedGallery":                policy.Microsoft_Compute_sharedGallery,
	"Microsoft.Compute.compute":                      policy.Microsoft_Compute_compute,
	"Microsoft.Portal.CloudShell":                    policy.Microsoft_Portal_CloudShell,
	"Microsoft.SignalRService.webpubsub":             policy.Microsoft_SignalRService_webpubsub,
	"Microsoft.AVS.vmware":                           policy.Microsoft_AVS_vmware,
	"Microsoft.DBforPostgreSQL.PrivateDnsZone":       policy.Microsoft_DBforPostgreSQL_PrivateDnsZone,
	"Microsoft.DBforPostgreSQL.Databases":            policy.Microsoft_DBforPostgreSQL_Databases,
	"Microsoft.DBforPostgreSQL.postgresql":           policy.Microsoft_DBforPostgreSQL_postgresql,
	"Microsoft.DevOps.devops":                        policy.Microsoft_DevOps_devops,
	"Microsoft.DBforMySQL.mysql":                     policy.Microsoft_DBforMySQL_mysql,
}
