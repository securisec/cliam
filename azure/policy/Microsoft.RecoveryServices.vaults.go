package policy

// Microsoft_RecoveryServices_vaults policy
var Microsoft_RecoveryServices_vaults = map[string]Policy{
	"PrivateLinkResources_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.RecoveryServices/vaults/{{.vaultName}}/privateLinkResources",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2023-01-01",
		},
		OperationID: "PrivateLinkResources_List",
		Resource:    "Microsoft.RecoveryServices",
	},
	"PrivateLinkResources_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.RecoveryServices/vaults/{{.vaultName}}/privateLinkResources/{{.privateLinkResourceName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2023-01-01",
		},
		OperationID: "PrivateLinkResources_Get",
		Resource:    "Microsoft.RecoveryServices",
	},
	"RecoveryServices_CheckNameAvailability": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.RecoveryServices/locations/{{.location}}/checkNameAvailability",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2023-01-01",
		},
		OperationID: "RecoveryServices_CheckNameAvailability",
		Resource:    "Microsoft.RecoveryServices",
	},
	"RecoveryServices_Capabilities": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.RecoveryServices/locations/{{.location}}/capabilities",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2023-01-01",
		},
		OperationID: "RecoveryServices_Capabilities",
		Resource:    "Microsoft.RecoveryServices",
	},
	"Vaults_ListBySubscriptionId": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.RecoveryServices/vaults",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2023-01-01",
		},
		OperationID: "Vaults_ListBySubscriptionId",
		Resource:    "Microsoft.RecoveryServices",
	},
	"Operations_List": {
		Path:   "/providers/Microsoft.RecoveryServices/operations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2023-01-01",
		},
		OperationID: "Operations_List",
		Resource:    "Microsoft.RecoveryServices",
	},
	"Vaults_ListByResourceGroup": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.RecoveryServices/vaults",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2023-01-01",
		},
		OperationID: "Vaults_ListByResourceGroup",
		Resource:    "Microsoft.RecoveryServices",
	},
	"Vaults_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.RecoveryServices/vaults/{{.vaultName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2023-01-01",
		},
		OperationID: "Vaults_Get",
		Resource:    "Microsoft.RecoveryServices",
	},
	"VaultExtendedInfo_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.RecoveryServices/vaults/{{.vaultName}}/extendedInformation/vaultExtendedInfo",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2023-01-01",
		},
		OperationID: "VaultExtendedInfo_Get",
		Resource:    "Microsoft.RecoveryServices",
	},
	"GetOperationStatus": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.RecoveryServices/vaults/{{.vaultName}}/operationStatus/{{.operationId}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2023-01-01",
		},
		OperationID: "GetOperationStatus",
		Resource:    "Microsoft.RecoveryServices",
	},
	"GetOperationResult": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.RecoveryServices/vaults/{{.vaultName}}/operationResults/{{.operationId}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2023-01-01",
		},
		OperationID: "GetOperationResult",
		Resource:    "Microsoft.RecoveryServices",
	},
}
