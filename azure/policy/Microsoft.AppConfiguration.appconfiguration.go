package policy

var Microsoft_AppConfiguration_appconfiguration = map[string]Policy{
	"ConfigurationStores_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.AppConfiguration/configurationStores",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-05-01",
		},
		OperationID: "ConfigurationStores_List",
		Resource:    "Microsoft.AppConfiguration",
	},
	"ConfigurationStores_ListByResourceGroup": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.AppConfiguration/configurationStores",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-05-01",
		},
		OperationID: "ConfigurationStores_ListByResourceGroup",
		Resource:    "Microsoft.AppConfiguration",
	},
	"ConfigurationStores_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.AppConfiguration/configurationStores/{{.configStoreName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-05-01",
		},
		OperationID: "ConfigurationStores_Get",
		Resource:    "Microsoft.AppConfiguration",
	},
	"Operations_CheckNameAvailability": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.AppConfiguration/checkNameAvailability",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-05-01",
		},
		OperationID: "Operations_CheckNameAvailability",
		Resource:    "Microsoft.AppConfiguration",
	},
	"ConfigurationStores_ListKeys": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.AppConfiguration/configurationStores/{{.configStoreName}}/listKeys",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-05-01",
		},
		OperationID: "ConfigurationStores_ListKeys",
		Resource:    "Microsoft.AppConfiguration",
	},
	"ConfigurationStores_RegenerateKey": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.AppConfiguration/configurationStores/{{.configStoreName}}/regenerateKey",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-05-01",
		},
		OperationID: "ConfigurationStores_RegenerateKey",
		Resource:    "Microsoft.AppConfiguration",
	},
	"Operations_List": {
		Path:   "/providers/Microsoft.AppConfiguration/operations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-05-01",
		},
		OperationID: "Operations_List",
		Resource:    "Microsoft.AppConfiguration",
	},
	"PrivateEndpointConnections_ListByConfigurationStore": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.AppConfiguration/configurationStores/{{.configStoreName}}/privateEndpointConnections",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-05-01",
		},
		OperationID: "PrivateEndpointConnections_ListByConfigurationStore",
		Resource:    "Microsoft.AppConfiguration",
	},
	"PrivateEndpointConnections_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.AppConfiguration/configurationStores/{{.configStoreName}}/privateEndpointConnections/{{.privateEndpointConnectionName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-05-01",
		},
		OperationID: "PrivateEndpointConnections_Get",
		Resource:    "Microsoft.AppConfiguration",
	},
	"PrivateLinkResources_ListByConfigurationStore": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.AppConfiguration/configurationStores/{{.configStoreName}}/privateLinkResources",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-05-01",
		},
		OperationID: "PrivateLinkResources_ListByConfigurationStore",
		Resource:    "Microsoft.AppConfiguration",
	},
	"PrivateLinkResources_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.AppConfiguration/configurationStores/{{.configStoreName}}/privateLinkResources/{{.groupName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-05-01",
		},
		OperationID: "PrivateLinkResources_Get",
		Resource:    "Microsoft.AppConfiguration",
	},
	"KeyValues_ListByConfigurationStore": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.AppConfiguration/configurationStores/{{.configStoreName}}/keyValues",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-05-01",
		},
		OperationID: "KeyValues_ListByConfigurationStore",
		Resource:    "Microsoft.AppConfiguration",
	},
	"KeyValues_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.AppConfiguration/configurationStores/{{.configStoreName}}/keyValues/{{.keyValueName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-05-01",
		},
		OperationID: "KeyValues_Get",
		Resource:    "Microsoft.AppConfiguration",
	},
	"ConfigurationStores_ListDeleted": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.AppConfiguration/deletedConfigurationStores",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-05-01",
		},
		OperationID: "ConfigurationStores_ListDeleted",
		Resource:    "Microsoft.AppConfiguration",
	},
	"ConfigurationStores_GetDeleted": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.AppConfiguration/locations/{{.location}}/deletedConfigurationStores/{{.configStoreName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-05-01",
		},
		OperationID: "ConfigurationStores_GetDeleted",
		Resource:    "Microsoft.AppConfiguration",
	},
	"ConfigurationStores_PurgeDeleted": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.AppConfiguration/locations/{{.location}}/deletedConfigurationStores/{{.configStoreName}}/purge",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-05-01",
		},
		OperationID: "ConfigurationStores_PurgeDeleted",
		Resource:    "Microsoft.AppConfiguration",
	},
	"Operations_RegionalCheckNameAvailability": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.AppConfiguration/locations/{{.location}}/checkNameAvailability",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-05-01",
		},
		OperationID: "Operations_RegionalCheckNameAvailability",
		Resource:    "Microsoft.AppConfiguration",
	},
}
