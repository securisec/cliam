package policy

var Microsoft_ContainerService_managedClusters = []Policy{
    {
    Path: "/providers/Microsoft.ContainerService/operations",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2022-04-01",
    },
	OperationID:    "Operations_List",
    Resource:       "Microsoft.ContainerService",
},{
    Path: "/subscriptions/{{.subscriptionId}}/providers/Microsoft.ContainerService/locations/{{.location}}/osOptions/default",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2022-04-01",
    },
	OperationID:    "ManagedClusters_GetOSOptions",
    Resource:       "Microsoft.ContainerService",
},{
    Path: "/subscriptions/{{.subscriptionId}}/providers/Microsoft.ContainerService/managedClusters",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2022-04-01",
    },
	OperationID:    "ManagedClusters_List",
    Resource:       "Microsoft.ContainerService",
},{
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ContainerService/managedClusters",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2022-04-01",
    },
	OperationID:    "ManagedClusters_ListByResourceGroup",
    Resource:       "Microsoft.ContainerService",
},{
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ContainerService/managedClusters/{{.resourceName}}/upgradeProfiles/default",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2022-04-01",
    },
	OperationID:    "ManagedClusters_GetUpgradeProfile",
    Resource:       "Microsoft.ContainerService",
},{
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ContainerService/managedClusters/{{.resourceName}}/accessProfiles/{{.roleName}}/listCredential",
	Method: "POST",
	QueryValues:   map[string]string{
        "api-version": "2022-04-01",
    },
	OperationID:    "ManagedClusters_GetAccessProfile",
    Resource:       "Microsoft.ContainerService",
},{
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ContainerService/managedClusters/{{.resourceName}}/listClusterAdminCredential",
	Method: "POST",
	QueryValues:   map[string]string{
        "api-version": "2022-04-01",
    },
	OperationID:    "ManagedClusters_ListClusterAdminCredentials",
    Resource:       "Microsoft.ContainerService",
},{
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ContainerService/managedClusters/{{.resourceName}}/listClusterUserCredential",
	Method: "POST",
	QueryValues:   map[string]string{
        "api-version": "2022-04-01",
    },
	OperationID:    "ManagedClusters_ListClusterUserCredentials",
    Resource:       "Microsoft.ContainerService",
},{
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ContainerService/managedClusters/{{.resourceName}}/listClusterMonitoringUserCredential",
	Method: "POST",
	QueryValues:   map[string]string{
        "api-version": "2022-04-01",
    },
	OperationID:    "ManagedClusters_ListClusterMonitoringUserCredentials",
    Resource:       "Microsoft.ContainerService",
},{
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ContainerService/managedClusters/{{.resourceName}}",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2022-04-01",
    },
	OperationID:    "ManagedClusters_Get",
    Resource:       "Microsoft.ContainerService",
},{
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ContainerService/managedClusters/{{.resourceName}}/maintenanceConfigurations",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2022-04-01",
    },
	OperationID:    "MaintenanceConfigurations_ListByManagedCluster",
    Resource:       "Microsoft.ContainerService",
},{
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ContainerService/managedClusters/{{.resourceName}}/maintenanceConfigurations/{{.configName}}",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2022-04-01",
    },
	OperationID:    "MaintenanceConfigurations_Get",
    Resource:       "Microsoft.ContainerService",
},{
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ContainerService/managedClusters/{{.resourceName}}/agentPools",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2022-04-01",
    },
	OperationID:    "AgentPools_List",
    Resource:       "Microsoft.ContainerService",
},{
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ContainerService/managedClusters/{{.resourceName}}/agentPools/{{.agentPoolName}}",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2022-04-01",
    },
	OperationID:    "AgentPools_Get",
    Resource:       "Microsoft.ContainerService",
},{
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ContainerService/managedClusters/{{.resourceName}}/agentPools/{{.agentPoolName}}/upgradeProfiles/default",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2022-04-01",
    },
	OperationID:    "AgentPools_GetUpgradeProfile",
    Resource:       "Microsoft.ContainerService",
},{
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ContainerService/managedClusters/{{.resourceName}}/availableAgentPoolVersions",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2022-04-01",
    },
	OperationID:    "AgentPools_GetAvailableAgentPoolVersions",
    Resource:       "Microsoft.ContainerService",
},{
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ContainerService/managedClusters/{{.resourceName}}/resetServicePrincipalProfile",
	Method: "POST",
	QueryValues:   map[string]string{
        "api-version": "2022-04-01",
    },
	OperationID:    "ManagedClusters_ResetServicePrincipalProfile",
    Resource:       "Microsoft.ContainerService",
},{
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ContainerService/managedClusters/{{.resourceName}}/resetAADProfile",
	Method: "POST",
	QueryValues:   map[string]string{
        "api-version": "2022-04-01",
    },
	OperationID:    "ManagedClusters_ResetAADProfile",
    Resource:       "Microsoft.ContainerService",
},{
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ContainerService/managedClusters/{{.resourceName}}/rotateClusterCertificates",
	Method: "POST",
	QueryValues:   map[string]string{
        "api-version": "2022-04-01",
    },
	OperationID:    "ManagedClusters_RotateClusterCertificates",
    Resource:       "Microsoft.ContainerService",
},{
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ContainerService/managedClusters/{{.resourceName}}/stop",
	Method: "POST",
	QueryValues:   map[string]string{
        "api-version": "2022-04-01",
    },
	OperationID:    "ManagedClusters_Stop",
    Resource:       "Microsoft.ContainerService",
},{
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ContainerService/managedClusters/{{.resourceName}}/start",
	Method: "POST",
	QueryValues:   map[string]string{
        "api-version": "2022-04-01",
    },
	OperationID:    "ManagedClusters_Start",
    Resource:       "Microsoft.ContainerService",
},{
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ContainerService/managedClusters/{{.resourceName}}/privateEndpointConnections",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2022-04-01",
    },
	OperationID:    "PrivateEndpointConnections_List",
    Resource:       "Microsoft.ContainerService",
},{
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ContainerService/managedClusters/{{.resourceName}}/privateEndpointConnections/{{.privateEndpointConnectionName}}",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2022-04-01",
    },
	OperationID:    "PrivateEndpointConnections_Get",
    Resource:       "Microsoft.ContainerService",
},{
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ContainerService/managedClusters/{{.resourceName}}/agentPools/{{.agentPoolName}}/upgradeNodeImageVersion",
	Method: "POST",
	QueryValues:   map[string]string{
        "api-version": "2022-04-01",
    },
	OperationID:    "AgentPools_UpgradeNodeImageVersion",
    Resource:       "Microsoft.ContainerService",
},{
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ContainerService/managedClusters/{{.resourceName}}/privateLinkResources",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2022-04-01",
    },
	OperationID:    "PrivateLinkResources_List",
    Resource:       "Microsoft.ContainerService",
},{
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ContainerService/managedClusters/{{.resourceName}}/resolvePrivateLinkServiceId",
	Method: "POST",
	QueryValues:   map[string]string{
        "api-version": "2022-04-01",
    },
	OperationID:    "ResolvePrivateLinkServiceId_POST",
    Resource:       "Microsoft.ContainerService",
},{
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ContainerService/managedClusters/{{.resourceName}}/runCommand",
	Method: "POST",
	QueryValues:   map[string]string{
        "api-version": "2022-04-01",
    },
	OperationID:    "ManagedClusters_RunCommand",
    Resource:       "Microsoft.ContainerService",
},{
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ContainerService/managedClusters/{{.resourceName}}/commandResults/{{.commandId}}",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2022-04-01",
    },
	OperationID:    "ManagedClusters_GetCommandResult",
    Resource:       "Microsoft.ContainerService",
},{
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ContainerService/managedClusters/{{.resourceName}}/outboundNetworkDependenciesEndpoints",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2022-04-01",
    },
	OperationID:    "ManagedClusters_ListOutboundNetworkDependenciesEndpoints",
    Resource:       "Microsoft.ContainerService",
},{
    Path: "/subscriptions/{{.subscriptionId}}/providers/Microsoft.ContainerService/snapshots",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2022-04-01",
    },
	OperationID:    "Snapshots_List",
    Resource:       "Microsoft.ContainerService",
},{
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ContainerService/snapshots",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2022-04-01",
    },
	OperationID:    "Snapshots_ListByResourceGroup",
    Resource:       "Microsoft.ContainerService",
},{
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ContainerService/snapshots/{{.resourceName}}",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2022-04-01",
    },
	OperationID:    "Snapshots_Get",
    Resource:       "Microsoft.ContainerService",
},
}
