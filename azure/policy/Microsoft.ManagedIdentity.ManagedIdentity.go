package policy

// Microsoft_ManagedIdentity_ManagedIdentity policy
var Microsoft_ManagedIdentity_ManagedIdentity = map[string]Policy{
	"SystemAssignedIdentities_GetByScope": {
		Path:   "/{{.scope}}/providers/Microsoft.ManagedIdentity/identities/default",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2023-01-31",
		},
		OperationID: "SystemAssignedIdentities_GetByScope",
		Resource:    "Microsoft.ManagedIdentity",
	},
	"Operations_List": {
		Path:   "/providers/Microsoft.ManagedIdentity/operations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2023-01-31",
		},
		OperationID: "Operations_List",
		Resource:    "Microsoft.ManagedIdentity",
	},
	"UserAssignedIdentities_ListBySubscription": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.ManagedIdentity/userAssignedIdentities",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2023-01-31",
		},
		OperationID: "UserAssignedIdentities_ListBySubscription",
		Resource:    "Microsoft.ManagedIdentity",
	},
	"UserAssignedIdentities_ListByResourceGroup": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ManagedIdentity/userAssignedIdentities",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2023-01-31",
		},
		OperationID: "UserAssignedIdentities_ListByResourceGroup",
		Resource:    "Microsoft.ManagedIdentity",
	},
	"UserAssignedIdentities_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ManagedIdentity/userAssignedIdentities/{{.resourceName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2023-01-31",
		},
		OperationID: "UserAssignedIdentities_Get",
		Resource:    "Microsoft.ManagedIdentity",
	},
	"FederatedIdentityCredentials_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ManagedIdentity/userAssignedIdentities/{{.resourceName}}/federatedIdentityCredentials",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2023-01-31",
		},
		OperationID: "FederatedIdentityCredentials_List",
		Resource:    "Microsoft.ManagedIdentity",
	},
	"FederatedIdentityCredentials_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ManagedIdentity/userAssignedIdentities/{{.resourceName}}/federatedIdentityCredentials/{{.federatedIdentityCredentialResourceName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2023-01-31",
		},
		OperationID: "FederatedIdentityCredentials_Get",
		Resource:    "Microsoft.ManagedIdentity",
	},
}
