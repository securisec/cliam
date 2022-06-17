package policy

var Microsoft_ContainerService_managedClusters = map[string]Policy{
	"Operations_List": {
		Path:   "/providers/Microsoft.ContainerService/operations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-04-01",
		},
		OperationID: "Operations_List",
		Resource:    "Microsoft.ContainerService",
	},
	"ManagedClusters_GetOSOptions": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.ContainerService/locations/{{.location}}/osOptions/default",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-04-01",
		},
		OperationID: "ManagedClusters_GetOSOptions",
		Resource:    "Microsoft.ContainerService",
	},
	"ManagedClusters_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.ContainerService/managedClusters",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-04-01",
		},
		OperationID: "ManagedClusters_List",
		Resource:    "Microsoft.ContainerService",
	},
	"ManagedClusters_ListByResourceGroup": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ContainerService/managedClusters",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-04-01",
		},
		OperationID: "ManagedClusters_ListByResourceGroup",
		Resource:    "Microsoft.ContainerService",
	},
	"ManagedClusters_GetUpgradeProfile": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ContainerService/managedClusters/{{.resourceName}}/upgradeProfiles/default",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-04-01",
		},
		OperationID: "ManagedClusters_GetUpgradeProfile",
		Resource:    "Microsoft.ContainerService",
	},
	"ManagedClusters_GetAccessProfile": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ContainerService/managedClusters/{{.resourceName}}/accessProfiles/{{.roleName}}/listCredential",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-04-01",
		},
		OperationID: "ManagedClusters_GetAccessProfile",
		Resource:    "Microsoft.ContainerService",
	},
	"ManagedClusters_ListClusterAdminCredentials": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ContainerService/managedClusters/{{.resourceName}}/listClusterAdminCredential",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-04-01",
		},
		OperationID: "ManagedClusters_ListClusterAdminCredentials",
		Resource:    "Microsoft.ContainerService",
	},
	"ManagedClusters_ListClusterUserCredentials": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ContainerService/managedClusters/{{.resourceName}}/listClusterUserCredential",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-04-01",
		},
		OperationID: "ManagedClusters_ListClusterUserCredentials",
		Resource:    "Microsoft.ContainerService",
	},
	"ManagedClusters_ListClusterMonitoringUserCredentials": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ContainerService/managedClusters/{{.resourceName}}/listClusterMonitoringUserCredential",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-04-01",
		},
		OperationID: "ManagedClusters_ListClusterMonitoringUserCredentials",
		Resource:    "Microsoft.ContainerService",
	},
	"ManagedClusters_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ContainerService/managedClusters/{{.resourceName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-04-01",
		},
		OperationID: "ManagedClusters_Get",
		Resource:    "Microsoft.ContainerService",
	},
	"MaintenanceConfigurations_ListByManagedCluster": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ContainerService/managedClusters/{{.resourceName}}/maintenanceConfigurations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-04-01",
		},
		OperationID: "MaintenanceConfigurations_ListByManagedCluster",
		Resource:    "Microsoft.ContainerService",
	},
	"MaintenanceConfigurations_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ContainerService/managedClusters/{{.resourceName}}/maintenanceConfigurations/{{.configName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-04-01",
		},
		OperationID: "MaintenanceConfigurations_Get",
		Resource:    "Microsoft.ContainerService",
	},
	"AgentPools_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ContainerService/managedClusters/{{.resourceName}}/agentPools",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-04-01",
		},
		OperationID: "AgentPools_List",
		Resource:    "Microsoft.ContainerService",
	},
	"AgentPools_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ContainerService/managedClusters/{{.resourceName}}/agentPools/{{.agentPoolName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-04-01",
		},
		OperationID: "AgentPools_Get",
		Resource:    "Microsoft.ContainerService",
	},
	"AgentPools_GetUpgradeProfile": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ContainerService/managedClusters/{{.resourceName}}/agentPools/{{.agentPoolName}}/upgradeProfiles/default",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-04-01",
		},
		OperationID: "AgentPools_GetUpgradeProfile",
		Resource:    "Microsoft.ContainerService",
	},
	"AgentPools_GetAvailableAgentPoolVersions": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ContainerService/managedClusters/{{.resourceName}}/availableAgentPoolVersions",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-04-01",
		},
		OperationID: "AgentPools_GetAvailableAgentPoolVersions",
		Resource:    "Microsoft.ContainerService",
	},
	"ManagedClusters_ResetServicePrincipalProfile": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ContainerService/managedClusters/{{.resourceName}}/resetServicePrincipalProfile",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-04-01",
		},
		OperationID: "ManagedClusters_ResetServicePrincipalProfile",
		Resource:    "Microsoft.ContainerService",
	},
	"ManagedClusters_ResetAADProfile": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ContainerService/managedClusters/{{.resourceName}}/resetAADProfile",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-04-01",
		},
		OperationID: "ManagedClusters_ResetAADProfile",
		Resource:    "Microsoft.ContainerService",
	},
	"ManagedClusters_RotateClusterCertificates": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ContainerService/managedClusters/{{.resourceName}}/rotateClusterCertificates",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-04-01",
		},
		OperationID: "ManagedClusters_RotateClusterCertificates",
		Resource:    "Microsoft.ContainerService",
	},
	"ManagedClusters_Stop": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ContainerService/managedClusters/{{.resourceName}}/stop",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-04-01",
		},
		OperationID: "ManagedClusters_Stop",
		Resource:    "Microsoft.ContainerService",
	},
	"ManagedClusters_Start": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ContainerService/managedClusters/{{.resourceName}}/start",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-04-01",
		},
		OperationID: "ManagedClusters_Start",
		Resource:    "Microsoft.ContainerService",
	},
	"PrivateEndpointConnections_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ContainerService/managedClusters/{{.resourceName}}/privateEndpointConnections",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-04-01",
		},
		OperationID: "PrivateEndpointConnections_List",
		Resource:    "Microsoft.ContainerService",
	},
	"PrivateEndpointConnections_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ContainerService/managedClusters/{{.resourceName}}/privateEndpointConnections/{{.privateEndpointConnectionName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-04-01",
		},
		OperationID: "PrivateEndpointConnections_Get",
		Resource:    "Microsoft.ContainerService",
	},
	"AgentPools_UpgradeNodeImageVersion": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ContainerService/managedClusters/{{.resourceName}}/agentPools/{{.agentPoolName}}/upgradeNodeImageVersion",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-04-01",
		},
		OperationID: "AgentPools_UpgradeNodeImageVersion",
		Resource:    "Microsoft.ContainerService",
	},
	"PrivateLinkResources_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ContainerService/managedClusters/{{.resourceName}}/privateLinkResources",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-04-01",
		},
		OperationID: "PrivateLinkResources_List",
		Resource:    "Microsoft.ContainerService",
	},
	"ResolvePrivateLinkServiceId_POST": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ContainerService/managedClusters/{{.resourceName}}/resolvePrivateLinkServiceId",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-04-01",
		},
		OperationID: "ResolvePrivateLinkServiceId_POST",
		Resource:    "Microsoft.ContainerService",
	},
	"ManagedClusters_RunCommand": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ContainerService/managedClusters/{{.resourceName}}/runCommand",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-04-01",
		},
		OperationID: "ManagedClusters_RunCommand",
		Resource:    "Microsoft.ContainerService",
	},
	"ManagedClusters_GetCommandResult": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ContainerService/managedClusters/{{.resourceName}}/commandResults/{{.commandId}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-04-01",
		},
		OperationID: "ManagedClusters_GetCommandResult",
		Resource:    "Microsoft.ContainerService",
	},
	"ManagedClusters_ListOutboundNetworkDependenciesEndpoints": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ContainerService/managedClusters/{{.resourceName}}/outboundNetworkDependenciesEndpoints",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-04-01",
		},
		OperationID: "ManagedClusters_ListOutboundNetworkDependenciesEndpoints",
		Resource:    "Microsoft.ContainerService",
	},
	"Snapshots_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.ContainerService/snapshots",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-04-01",
		},
		OperationID: "Snapshots_List",
		Resource:    "Microsoft.ContainerService",
	},
	"Snapshots_ListByResourceGroup": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ContainerService/snapshots",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-04-01",
		},
		OperationID: "Snapshots_ListByResourceGroup",
		Resource:    "Microsoft.ContainerService",
	},
	"Snapshots_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ContainerService/snapshots/{{.resourceName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-04-01",
		},
		OperationID: "Snapshots_Get",
		Resource:    "Microsoft.ContainerService",
	},
}
