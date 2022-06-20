package policy

var Microsoft_Cache_redisenterprise = map[string]Policy{
	"Operations_List": {
		Path:   "/providers/Microsoft.Cache/operations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-01-01",
		},
		OperationID: "Operations_List",
		Resource:    "Microsoft.Cache",
	},
	"OperationsStatus_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Cache/locations/{{.location}}/operationsStatus/{{.operationId}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-01-01",
		},
		OperationID: "OperationsStatus_Get",
		Resource:    "Microsoft.Cache",
	},
	"RedisEnterprise_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Cache/redisEnterprise/{{.clusterName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-01-01",
		},
		OperationID: "RedisEnterprise_Get",
		Resource:    "Microsoft.Cache",
	},
	"RedisEnterprise_ListByResourceGroup": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Cache/redisEnterprise",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-01-01",
		},
		OperationID: "RedisEnterprise_ListByResourceGroup",
		Resource:    "Microsoft.Cache",
	},
	"RedisEnterprise_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Cache/redisEnterprise",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-01-01",
		},
		OperationID: "RedisEnterprise_List",
		Resource:    "Microsoft.Cache",
	},
	"Databases_ListByCluster": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Cache/redisEnterprise/{{.clusterName}}/databases",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-01-01",
		},
		OperationID: "Databases_ListByCluster",
		Resource:    "Microsoft.Cache",
	},
	"Databases_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Cache/redisEnterprise/{{.clusterName}}/databases/{{.databaseName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-01-01",
		},
		OperationID: "Databases_Get",
		Resource:    "Microsoft.Cache",
	},
	"Databases_ListKeys": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Cache/redisEnterprise/{{.clusterName}}/databases/{{.databaseName}}/listKeys",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-01-01",
		},
		OperationID: "Databases_ListKeys",
		Resource:    "Microsoft.Cache",
	},
	"Databases_RegenerateKey": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Cache/redisEnterprise/{{.clusterName}}/databases/{{.databaseName}}/regenerateKey",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-01-01",
		},
		OperationID: "Databases_RegenerateKey",
		Resource:    "Microsoft.Cache",
	},
	"Databases_Import": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Cache/redisEnterprise/{{.clusterName}}/databases/{{.databaseName}}/import",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-01-01",
		},
		OperationID: "Databases_Import",
		Resource:    "Microsoft.Cache",
	},
	"Databases_Export": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Cache/redisEnterprise/{{.clusterName}}/databases/{{.databaseName}}/export",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-01-01",
		},
		OperationID: "Databases_Export",
		Resource:    "Microsoft.Cache",
	},
	"PrivateEndpointConnections_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Cache/redisEnterprise/{{.clusterName}}/privateEndpointConnections",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-01-01",
		},
		OperationID: "PrivateEndpointConnections_List",
		Resource:    "Microsoft.Cache",
	},
	"PrivateEndpointConnections_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Cache/redisEnterprise/{{.clusterName}}/privateEndpointConnections/{{.privateEndpointConnectionName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-01-01",
		},
		OperationID: "PrivateEndpointConnections_Get",
		Resource:    "Microsoft.Cache",
	},
	"PrivateLinkResources_ListByCluster": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Cache/redisEnterprise/{{.clusterName}}/privateLinkResources",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-01-01",
		},
		OperationID: "PrivateLinkResources_ListByCluster",
		Resource:    "Microsoft.Cache",
	},
	"Databases_ForceUnlink": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Cache/redisEnterprise/{{.clusterName}}/databases/{{.databaseName}}/forceUnlink",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-01-01",
		},
		OperationID: "Databases_ForceUnlink",
		Resource:    "Microsoft.Cache",
	},
}
