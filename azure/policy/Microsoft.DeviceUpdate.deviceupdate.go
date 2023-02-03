package policy

// Microsoft_DeviceUpdate_deviceupdate policy
var Microsoft_DeviceUpdate_deviceupdate = map[string]Policy{
	"CheckNameAvailability": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.DeviceUpdate/checknameavailability",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "CheckNameAvailability",
		Resource:    "Microsoft.DeviceUpdate",
	},
	"Accounts_ListBySubscription": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.DeviceUpdate/accounts",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "Accounts_ListBySubscription",
		Resource:    "Microsoft.DeviceUpdate",
	},
	"Accounts_ListByResourceGroup": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DeviceUpdate/accounts",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "Accounts_ListByResourceGroup",
		Resource:    "Microsoft.DeviceUpdate",
	},
	"Accounts_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DeviceUpdate/accounts/{{.accountName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "Accounts_Get",
		Resource:    "Microsoft.DeviceUpdate",
	},
	"Instances_ListByAccount": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DeviceUpdate/accounts/{{.accountName}}/instances",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "Instances_ListByAccount",
		Resource:    "Microsoft.DeviceUpdate",
	},
	"Instances_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DeviceUpdate/accounts/{{.accountName}}/instances/{{.instanceName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "Instances_Get",
		Resource:    "Microsoft.DeviceUpdate",
	},
	"PrivateEndpointConnections_ListByAccount": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DeviceUpdate/accounts/{{.accountName}}/privateEndpointConnections",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "PrivateEndpointConnections_ListByAccount",
		Resource:    "Microsoft.DeviceUpdate",
	},
	"PrivateEndpointConnections_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DeviceUpdate/accounts/{{.accountName}}/privateEndpointConnections/{{.privateEndpointConnectionName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "PrivateEndpointConnections_Get",
		Resource:    "Microsoft.DeviceUpdate",
	},
	"PrivateLinkResources_ListByAccount": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DeviceUpdate/accounts/{{.accountName}}/privateLinkResources",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "PrivateLinkResources_ListByAccount",
		Resource:    "Microsoft.DeviceUpdate",
	},
	"PrivateLinkResources_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DeviceUpdate/accounts/{{.accountName}}/privateLinkResources/{{.groupId}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "PrivateLinkResources_Get",
		Resource:    "Microsoft.DeviceUpdate",
	},
	"PrivateEndpointConnectionProxies_ListByAccount": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DeviceUpdate/accounts/{{.accountName}}/privateEndpointConnectionProxies",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "PrivateEndpointConnectionProxies_ListByAccount",
		Resource:    "Microsoft.DeviceUpdate",
	},
	"PrivateEndpointConnectionProxies_Validate": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DeviceUpdate/accounts/{{.accountName}}/privateEndpointConnectionProxies/{{.privateEndpointConnectionProxyId}}/validate",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "PrivateEndpointConnectionProxies_Validate",
		Resource:    "Microsoft.DeviceUpdate",
	},
	"PrivateEndpointConnectionProxies_UpdatePrivateEndpointProperties": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DeviceUpdate/accounts/{{.accountName}}/privateEndpointConnectionProxies/{{.privateEndpointConnectionProxyId}}/updatePrivateEndpointProperties",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "PrivateEndpointConnectionProxies_UpdatePrivateEndpointProperties",
		Resource:    "Microsoft.DeviceUpdate",
	},
	"PrivateEndpointConnectionProxies_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DeviceUpdate/accounts/{{.accountName}}/privateEndpointConnectionProxies/{{.privateEndpointConnectionProxyId}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "PrivateEndpointConnectionProxies_Get",
		Resource:    "Microsoft.DeviceUpdate",
	},
	"Operations_List": {
		Path:   "/providers/Microsoft.DeviceUpdate/operations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "Operations_List",
		Resource:    "Microsoft.DeviceUpdate",
	},
}
