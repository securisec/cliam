package policy

var Microsoft_KeyVault_managedHsm = map[string]Policy{
	"ManagedHsms_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.KeyVault/managedHSMs/{{.name}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-10-01",
		},
		OperationID: "ManagedHsms_Get",
		Resource:    "Microsoft.KeyVault",
	},
	"ManagedHsms_ListByResourceGroup": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.KeyVault/managedHSMs",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-10-01",
		},
		OperationID: "ManagedHsms_ListByResourceGroup",
		Resource:    "Microsoft.KeyVault",
	},
	"ManagedHsms_ListBySubscription": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.KeyVault/managedHSMs",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-10-01",
		},
		OperationID: "ManagedHsms_ListBySubscription",
		Resource:    "Microsoft.KeyVault",
	},
	"MHSMPrivateEndpointConnections_ListByResource": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.KeyVault/managedHSMs/{{.name}}/privateEndpointConnections",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-10-01",
		},
		OperationID: "MHSMPrivateEndpointConnections_ListByResource",
		Resource:    "Microsoft.KeyVault",
	},
	"ManagedHsms_ListDeleted": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.KeyVault/deletedManagedHSMs",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-10-01",
		},
		OperationID: "ManagedHsms_ListDeleted",
		Resource:    "Microsoft.KeyVault",
	},
	"ManagedHsms_GetDeleted": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.KeyVault/locations/{{.location}}/deletedManagedHSMs/{{.name}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-10-01",
		},
		OperationID: "ManagedHsms_GetDeleted",
		Resource:    "Microsoft.KeyVault",
	},
	"ManagedHsms_PurgeDeleted": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.KeyVault/locations/{{.location}}/deletedManagedHSMs/{{.name}}/purge",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-10-01",
		},
		OperationID: "ManagedHsms_PurgeDeleted",
		Resource:    "Microsoft.KeyVault",
	},
	"MHSMPrivateEndpointConnections_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.KeyVault/managedHSMs/{{.name}}/privateEndpointConnections/{{.privateEndpointConnectionName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-10-01",
		},
		OperationID: "MHSMPrivateEndpointConnections_Get",
		Resource:    "Microsoft.KeyVault",
	},
	"MHSMPrivateLinkResources_ListByMHSMResource": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.KeyVault/managedHSMs/{{.name}}/privateLinkResources",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-10-01",
		},
		OperationID: "MHSMPrivateLinkResources_ListByMHSMResource",
		Resource:    "Microsoft.KeyVault",
	},
}
