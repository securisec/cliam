package policy

var Microsoft_Cache_redis = map[string]Policy{
	"Operations_List": {
		Path:   "/providers/Microsoft.Cache/operations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "Operations_List",
		Resource:    "Microsoft.Cache",
	},
	"Redis_CheckNameAvailability": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Cache/CheckNameAvailability",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "Redis_CheckNameAvailability",
		Resource:    "Microsoft.Cache",
	},
	"Redis_ListUpgradeNotifications": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Cache/redis/{{.name}}/listUpgradeNotifications",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "Redis_ListUpgradeNotifications",
		Resource:    "Microsoft.Cache",
	},
	"Redis_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Cache/redis/{{.name}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "Redis_Get",
		Resource:    "Microsoft.Cache",
	},
	"Redis_ListByResourceGroup": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Cache/redis",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "Redis_ListByResourceGroup",
		Resource:    "Microsoft.Cache",
	},
	"Redis_ListBySubscription": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Cache/redis",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "Redis_ListBySubscription",
		Resource:    "Microsoft.Cache",
	},
	"Redis_ListKeys": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Cache/redis/{{.name}}/listKeys",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "Redis_ListKeys",
		Resource:    "Microsoft.Cache",
	},
	"Redis_RegenerateKey": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Cache/redis/{{.name}}/regenerateKey",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "Redis_RegenerateKey",
		Resource:    "Microsoft.Cache",
	},
	"Redis_ForceReboot": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Cache/redis/{{.name}}/forceReboot",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "Redis_ForceReboot",
		Resource:    "Microsoft.Cache",
	},
	"Redis_ImportData": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Cache/redis/{{.name}}/import",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "Redis_ImportData",
		Resource:    "Microsoft.Cache",
	},
	"Redis_ExportData": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Cache/redis/{{.name}}/export",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "Redis_ExportData",
		Resource:    "Microsoft.Cache",
	},
	"FirewallRules_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Cache/redis/{{.cacheName}}/firewallRules",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "FirewallRules_List",
		Resource:    "Microsoft.Cache",
	},
	"FirewallRules_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Cache/redis/{{.cacheName}}/firewallRules/{{.ruleName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "FirewallRules_Get",
		Resource:    "Microsoft.Cache",
	},
	"PatchSchedules_ListByRedisResource": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Cache/redis/{{.cacheName}}/patchSchedules",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "PatchSchedules_ListByRedisResource",
		Resource:    "Microsoft.Cache",
	},
	"PatchSchedules_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Cache/redis/{{.name}}/patchSchedules/{{.default}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "PatchSchedules_Get",
		Resource:    "Microsoft.Cache",
	},
	"LinkedServer_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Cache/redis/{{.name}}/linkedServers/{{.linkedServerName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "LinkedServer_Get",
		Resource:    "Microsoft.Cache",
	},
	"LinkedServer_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Cache/redis/{{.name}}/linkedServers",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "LinkedServer_List",
		Resource:    "Microsoft.Cache",
	},
	"PrivateEndpointConnections_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Cache/redis/{{.cacheName}}/privateEndpointConnections",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "PrivateEndpointConnections_List",
		Resource:    "Microsoft.Cache",
	},
	"PrivateEndpointConnections_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Cache/redis/{{.cacheName}}/privateEndpointConnections/{{.privateEndpointConnectionName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "PrivateEndpointConnections_Get",
		Resource:    "Microsoft.Cache",
	},
	"PrivateLinkResources_ListByRedisCache": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Cache/redis/{{.cacheName}}/privateLinkResources",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "PrivateLinkResources_ListByRedisCache",
		Resource:    "Microsoft.Cache",
	},
	"AsyncOperationStatus_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Cache/locations/{{.location}}/asyncOperations/{{.operationId}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "AsyncOperationStatus_Get",
		Resource:    "Microsoft.Cache",
	},
}
