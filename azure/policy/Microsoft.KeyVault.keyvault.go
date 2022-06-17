package policy

var Microsoft_KeyVault_keyvault = map[string]Policy{
	"Vaults_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.KeyVault/vaults/{{.vaultName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-10-01",
		},
		OperationID: "Vaults_Get",
		Resource:    "Microsoft.KeyVault",
	},
	"Vaults_ListByResourceGroup": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.KeyVault/vaults",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-10-01",
		},
		OperationID: "Vaults_ListByResourceGroup",
		Resource:    "Microsoft.KeyVault",
	},
	"Vaults_ListBySubscription": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.KeyVault/vaults",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-10-01",
		},
		OperationID: "Vaults_ListBySubscription",
		Resource:    "Microsoft.KeyVault",
	},
	"Vaults_ListDeleted": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.KeyVault/deletedVaults",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-10-01",
		},
		OperationID: "Vaults_ListDeleted",
		Resource:    "Microsoft.KeyVault",
	},
	"Vaults_GetDeleted": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.KeyVault/locations/{{.location}}/deletedVaults/{{.vaultName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-10-01",
		},
		OperationID: "Vaults_GetDeleted",
		Resource:    "Microsoft.KeyVault",
	},
	"Vaults_PurgeDeleted": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.KeyVault/locations/{{.location}}/deletedVaults/{{.vaultName}}/purge",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-10-01",
		},
		OperationID: "Vaults_PurgeDeleted",
		Resource:    "Microsoft.KeyVault",
	},
	"Vaults_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resources",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-10-01",
		},
		OperationID: "Vaults_List",
		Resource:    "Microsoft.KeyVault",
	},
	"Vaults_CheckNameAvailability": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.KeyVault/checkNameAvailability",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-10-01",
		},
		OperationID: "Vaults_CheckNameAvailability",
		Resource:    "Microsoft.KeyVault",
	},
	"PrivateEndpointConnections_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.KeyVault/vaults/{{.vaultName}}/privateEndpointConnections/{{.privateEndpointConnectionName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-10-01",
		},
		OperationID: "PrivateEndpointConnections_Get",
		Resource:    "Microsoft.KeyVault",
	},
	"PrivateEndpointConnections_ListByResource": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.KeyVault/vaults/{{.vaultName}}/privateEndpointConnections",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-10-01",
		},
		OperationID: "PrivateEndpointConnections_ListByResource",
		Resource:    "Microsoft.KeyVault",
	},
	"PrivateLinkResources_ListByVault": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.KeyVault/vaults/{{.vaultName}}/privateLinkResources",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-10-01",
		},
		OperationID: "PrivateLinkResources_ListByVault",
		Resource:    "Microsoft.KeyVault",
	},
}
