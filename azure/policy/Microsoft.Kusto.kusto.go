package policy

// Microsoft_Kusto_kusto policy
var Microsoft_Kusto_kusto = map[string]Policy{
	"Clusters_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Kusto/clusters/{{.clusterName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-12-29",
		},
		OperationID: "Clusters_Get",
		Resource:    "Microsoft.Kusto",
	},
	"Clusters_Stop": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Kusto/clusters/{{.clusterName}}/stop",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-12-29",
		},
		OperationID: "Clusters_Stop",
		Resource:    "Microsoft.Kusto",
	},
	"Clusters_Start": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Kusto/clusters/{{.clusterName}}/start",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-12-29",
		},
		OperationID: "Clusters_Start",
		Resource:    "Microsoft.Kusto",
	},
	"ClusterPrincipalAssignments_CheckNameAvailability": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Kusto/clusters/{{.clusterName}}/checkPrincipalAssignmentNameAvailability",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-12-29",
		},
		OperationID: "ClusterPrincipalAssignments_CheckNameAvailability",
		Resource:    "Microsoft.Kusto",
	},
	"ClusterPrincipalAssignments_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Kusto/clusters/{{.clusterName}}/principalAssignments/{{.principalAssignmentName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-12-29",
		},
		OperationID: "ClusterPrincipalAssignments_Get",
		Resource:    "Microsoft.Kusto",
	},
	"ClusterPrincipalAssignments_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Kusto/clusters/{{.clusterName}}/principalAssignments",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-12-29",
		},
		OperationID: "ClusterPrincipalAssignments_List",
		Resource:    "Microsoft.Kusto",
	},
	"Clusters_ListFollowerDatabases": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Kusto/clusters/{{.clusterName}}/listFollowerDatabases",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-12-29",
		},
		OperationID: "Clusters_ListFollowerDatabases",
		Resource:    "Microsoft.Kusto",
	},
	"Clusters_DetachFollowerDatabases": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Kusto/clusters/{{.clusterName}}/detachFollowerDatabases",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-12-29",
		},
		OperationID: "Clusters_DetachFollowerDatabases",
		Resource:    "Microsoft.Kusto",
	},
	"Clusters_DiagnoseVirtualNetwork": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Kusto/clusters/{{.clusterName}}/diagnoseVirtualNetwork",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-12-29",
		},
		OperationID: "Clusters_DiagnoseVirtualNetwork",
		Resource:    "Microsoft.Kusto",
	},
	"Clusters_ListByResourceGroup": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Kusto/clusters",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-12-29",
		},
		OperationID: "Clusters_ListByResourceGroup",
		Resource:    "Microsoft.Kusto",
	},
	"Clusters_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Kusto/clusters",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-12-29",
		},
		OperationID: "Clusters_List",
		Resource:    "Microsoft.Kusto",
	},
	"Clusters_ListSkus": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Kusto/skus",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-12-29",
		},
		OperationID: "Clusters_ListSkus",
		Resource:    "Microsoft.Kusto",
	},
	"Clusters_CheckNameAvailability": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Kusto/locations/{{.location}}/checkNameAvailability",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-12-29",
		},
		OperationID: "Clusters_CheckNameAvailability",
		Resource:    "Microsoft.Kusto",
	},
	"Skus_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Kusto/locations/{{.location}}/skus",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-12-29",
		},
		OperationID: "Skus_List",
		Resource:    "Microsoft.Kusto",
	},
	"Databases_CheckNameAvailability": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Kusto/clusters/{{.clusterName}}/checkNameAvailability",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-12-29",
		},
		OperationID: "Databases_CheckNameAvailability",
		Resource:    "Microsoft.Kusto",
	},
	"AttachedDatabaseConfigurations_CheckNameAvailability": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Kusto/clusters/{{.clusterName}}/attachedDatabaseConfigurationCheckNameAvailability",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-12-29",
		},
		OperationID: "AttachedDatabaseConfigurations_CheckNameAvailability",
		Resource:    "Microsoft.Kusto",
	},
	"ManagedPrivateEndpoints_CheckNameAvailability": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Kusto/clusters/{{.clusterName}}/managedPrivateEndpointsCheckNameAvailability",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-12-29",
		},
		OperationID: "ManagedPrivateEndpoints_CheckNameAvailability",
		Resource:    "Microsoft.Kusto",
	},
	"Clusters_ListSkusByResource": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Kusto/clusters/{{.clusterName}}/skus",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-12-29",
		},
		OperationID: "Clusters_ListSkusByResource",
		Resource:    "Microsoft.Kusto",
	},
	"Databases_ListByCluster": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Kusto/clusters/{{.clusterName}}/databases",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-12-29",
		},
		OperationID: "Databases_ListByCluster",
		Resource:    "Microsoft.Kusto",
	},
	"Databases_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Kusto/clusters/{{.clusterName}}/databases/{{.databaseName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-12-29",
		},
		OperationID: "Databases_Get",
		Resource:    "Microsoft.Kusto",
	},
	"DatabasePrincipalAssignments_CheckNameAvailability": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Kusto/clusters/{{.clusterName}}/databases/{{.databaseName}}/checkPrincipalAssignmentNameAvailability",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-12-29",
		},
		OperationID: "DatabasePrincipalAssignments_CheckNameAvailability",
		Resource:    "Microsoft.Kusto",
	},
	"DatabasePrincipalAssignments_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Kusto/clusters/{{.clusterName}}/databases/{{.databaseName}}/principalAssignments/{{.principalAssignmentName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-12-29",
		},
		OperationID: "DatabasePrincipalAssignments_Get",
		Resource:    "Microsoft.Kusto",
	},
	"DatabasePrincipalAssignments_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Kusto/clusters/{{.clusterName}}/databases/{{.databaseName}}/principalAssignments",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-12-29",
		},
		OperationID: "DatabasePrincipalAssignments_List",
		Resource:    "Microsoft.Kusto",
	},
	"Databases_ListPrincipals": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Kusto/clusters/{{.clusterName}}/databases/{{.databaseName}}/listPrincipals",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-12-29",
		},
		OperationID: "Databases_ListPrincipals",
		Resource:    "Microsoft.Kusto",
	},
	"Databases_AddPrincipals": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Kusto/clusters/{{.clusterName}}/databases/{{.databaseName}}/addPrincipals",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-12-29",
		},
		OperationID: "Databases_AddPrincipals",
		Resource:    "Microsoft.Kusto",
	},
	"Scripts_ListByDatabase": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Kusto/clusters/{{.clusterName}}/databases/{{.databaseName}}/scripts",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-12-29",
		},
		OperationID: "Scripts_ListByDatabase",
		Resource:    "Microsoft.Kusto",
	},
	"Scripts_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Kusto/clusters/{{.clusterName}}/databases/{{.databaseName}}/scripts/{{.scriptName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-12-29",
		},
		OperationID: "Scripts_Get",
		Resource:    "Microsoft.Kusto",
	},
	"ManagedPrivateEndpoints_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Kusto/clusters/{{.clusterName}}/managedPrivateEndpoints",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-12-29",
		},
		OperationID: "ManagedPrivateEndpoints_List",
		Resource:    "Microsoft.Kusto",
	},
	"ManagedPrivateEndpoints_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Kusto/clusters/{{.clusterName}}/managedPrivateEndpoints/{{.managedPrivateEndpointName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-12-29",
		},
		OperationID: "ManagedPrivateEndpoints_Get",
		Resource:    "Microsoft.Kusto",
	},
	"PrivateEndpointConnections_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Kusto/clusters/{{.clusterName}}/privateEndpointConnections",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-12-29",
		},
		OperationID: "PrivateEndpointConnections_List",
		Resource:    "Microsoft.Kusto",
	},
	"PrivateEndpointConnections_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Kusto/clusters/{{.clusterName}}/privateEndpointConnections/{{.privateEndpointConnectionName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-12-29",
		},
		OperationID: "PrivateEndpointConnections_Get",
		Resource:    "Microsoft.Kusto",
	},
	"PrivateLinkResources_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Kusto/clusters/{{.clusterName}}/privateLinkResources",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-12-29",
		},
		OperationID: "PrivateLinkResources_List",
		Resource:    "Microsoft.Kusto",
	},
	"PrivateLinkResources_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Kusto/clusters/{{.clusterName}}/privateLinkResources/{{.privateLinkResourceName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-12-29",
		},
		OperationID: "PrivateLinkResources_Get",
		Resource:    "Microsoft.Kusto",
	},
	"Clusters_ListOutboundNetworkDependenciesEndpoints": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Kusto/clusters/{{.clusterName}}/outboundNetworkDependenciesEndpoints",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-12-29",
		},
		OperationID: "Clusters_ListOutboundNetworkDependenciesEndpoints",
		Resource:    "Microsoft.Kusto",
	},
	"AttachedDatabaseConfigurations_ListByCluster": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Kusto/clusters/{{.clusterName}}/attachedDatabaseConfigurations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-12-29",
		},
		OperationID: "AttachedDatabaseConfigurations_ListByCluster",
		Resource:    "Microsoft.Kusto",
	},
	"AttachedDatabaseConfigurations_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Kusto/clusters/{{.clusterName}}/attachedDatabaseConfigurations/{{.attachedDatabaseConfigurationName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-12-29",
		},
		OperationID: "AttachedDatabaseConfigurations_Get",
		Resource:    "Microsoft.Kusto",
	},
	"Databases_RemovePrincipals": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Kusto/clusters/{{.clusterName}}/databases/{{.databaseName}}/removePrincipals",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-12-29",
		},
		OperationID: "Databases_RemovePrincipals",
		Resource:    "Microsoft.Kusto",
	},
	"DataConnections_ListByDatabase": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Kusto/clusters/{{.clusterName}}/databases/{{.databaseName}}/dataConnections",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-12-29",
		},
		OperationID: "DataConnections_ListByDatabase",
		Resource:    "Microsoft.Kusto",
	},
	"DataConnections_dataConnectionValidation": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Kusto/clusters/{{.clusterName}}/databases/{{.databaseName}}/dataConnectionValidation",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-12-29",
		},
		OperationID: "DataConnections_dataConnectionValidation",
		Resource:    "Microsoft.Kusto",
	},
	"DataConnections_CheckNameAvailability": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Kusto/clusters/{{.clusterName}}/databases/{{.databaseName}}/checkNameAvailability",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-12-29",
		},
		OperationID: "DataConnections_CheckNameAvailability",
		Resource:    "Microsoft.Kusto",
	},
	"Scripts_CheckNameAvailability": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Kusto/clusters/{{.clusterName}}/databases/{{.databaseName}}/scriptsCheckNameAvailability",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-12-29",
		},
		OperationID: "Scripts_CheckNameAvailability",
		Resource:    "Microsoft.Kusto",
	},
	"DataConnections_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Kusto/clusters/{{.clusterName}}/databases/{{.databaseName}}/dataConnections/{{.dataConnectionName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-12-29",
		},
		OperationID: "DataConnections_Get",
		Resource:    "Microsoft.Kusto",
	},
	"Operations_List": {
		Path:   "/providers/Microsoft.Kusto/operations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-12-29",
		},
		OperationID: "Operations_List",
		Resource:    "Microsoft.Kusto",
	},
	"Clusters_ListLanguageExtensions": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Kusto/clusters/{{.clusterName}}/listLanguageExtensions",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-12-29",
		},
		OperationID: "Clusters_ListLanguageExtensions",
		Resource:    "Microsoft.Kusto",
	},
	"Clusters_AddLanguageExtensions": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Kusto/clusters/{{.clusterName}}/addLanguageExtensions",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-12-29",
		},
		OperationID: "Clusters_AddLanguageExtensions",
		Resource:    "Microsoft.Kusto",
	},
	"Clusters_RemoveLanguageExtensions": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Kusto/clusters/{{.clusterName}}/removeLanguageExtensions",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-12-29",
		},
		OperationID: "Clusters_RemoveLanguageExtensions",
		Resource:    "Microsoft.Kusto",
	},
	"OperationsResults_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Kusto/locations/{{.location}}/operationResults/{{.operationId}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-12-29",
		},
		OperationID: "OperationsResults_Get",
		Resource:    "Microsoft.Kusto",
	},
}
