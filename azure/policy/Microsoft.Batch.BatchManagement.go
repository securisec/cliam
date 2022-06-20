package policy

var Microsoft_Batch_BatchManagement = map[string]Policy{
	"BatchAccount_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Batch/batchAccounts/{{.accountName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-06-01",
		},
		OperationID: "BatchAccount_Get",
		Resource:    "Microsoft.Batch",
	},
	"BatchAccount_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Batch/batchAccounts",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-06-01",
		},
		OperationID: "BatchAccount_List",
		Resource:    "Microsoft.Batch",
	},
	"BatchAccount_ListByResourceGroup": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Batch/batchAccounts",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-06-01",
		},
		OperationID: "BatchAccount_ListByResourceGroup",
		Resource:    "Microsoft.Batch",
	},
	"BatchAccount_SynchronizeAutoStorageKeys": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Batch/batchAccounts/{{.accountName}}/syncAutoStorageKeys",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-06-01",
		},
		OperationID: "BatchAccount_SynchronizeAutoStorageKeys",
		Resource:    "Microsoft.Batch",
	},
	"BatchAccount_RegenerateKey": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Batch/batchAccounts/{{.accountName}}/regenerateKeys",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-06-01",
		},
		OperationID: "BatchAccount_RegenerateKey",
		Resource:    "Microsoft.Batch",
	},
	"BatchAccount_GetKeys": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Batch/batchAccounts/{{.accountName}}/listKeys",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-06-01",
		},
		OperationID: "BatchAccount_GetKeys",
		Resource:    "Microsoft.Batch",
	},
	"ApplicationPackage_Activate": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Batch/batchAccounts/{{.accountName}}/applications/{{.applicationName}}/versions/{{.versionName}}/activate",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-06-01",
		},
		OperationID: "ApplicationPackage_Activate",
		Resource:    "Microsoft.Batch",
	},
	"Application_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Batch/batchAccounts/{{.accountName}}/applications/{{.applicationName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-06-01",
		},
		OperationID: "Application_Get",
		Resource:    "Microsoft.Batch",
	},
	"ApplicationPackage_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Batch/batchAccounts/{{.accountName}}/applications/{{.applicationName}}/versions/{{.versionName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-06-01",
		},
		OperationID: "ApplicationPackage_Get",
		Resource:    "Microsoft.Batch",
	},
	"Application_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Batch/batchAccounts/{{.accountName}}/applications",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-06-01",
		},
		OperationID: "Application_List",
		Resource:    "Microsoft.Batch",
	},
	"ApplicationPackage_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Batch/batchAccounts/{{.accountName}}/applications/{{.applicationName}}/versions",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-06-01",
		},
		OperationID: "ApplicationPackage_List",
		Resource:    "Microsoft.Batch",
	},
	"Location_GetQuotas": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Batch/locations/{{.locationName}}/quotas",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-06-01",
		},
		OperationID: "Location_GetQuotas",
		Resource:    "Microsoft.Batch",
	},
	"Location_ListSupportedVirtualMachineSkus": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Batch/locations/{{.locationName}}/virtualMachineSkus",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-06-01",
		},
		OperationID: "Location_ListSupportedVirtualMachineSkus",
		Resource:    "Microsoft.Batch",
	},
	"Location_ListSupportedCloudServiceSkus": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Batch/locations/{{.locationName}}/cloudServiceSkus",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-06-01",
		},
		OperationID: "Location_ListSupportedCloudServiceSkus",
		Resource:    "Microsoft.Batch",
	},
	"Operations_List": {
		Path:   "/providers/Microsoft.Batch/operations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-06-01",
		},
		OperationID: "Operations_List",
		Resource:    "Microsoft.Batch",
	},
	"Location_CheckNameAvailability": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Batch/locations/{{.locationName}}/checkNameAvailability",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-06-01",
		},
		OperationID: "Location_CheckNameAvailability",
		Resource:    "Microsoft.Batch",
	},
	"Certificate_ListByBatchAccount": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Batch/batchAccounts/{{.accountName}}/certificates",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-06-01",
		},
		OperationID: "Certificate_ListByBatchAccount",
		Resource:    "Microsoft.Batch",
	},
	"Certificate_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Batch/batchAccounts/{{.accountName}}/certificates/{{.certificateName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-06-01",
		},
		OperationID: "Certificate_Get",
		Resource:    "Microsoft.Batch",
	},
	"Certificate_CancelDeletion": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Batch/batchAccounts/{{.accountName}}/certificates/{{.certificateName}}/cancelDelete",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-06-01",
		},
		OperationID: "Certificate_CancelDeletion",
		Resource:    "Microsoft.Batch",
	},
	"BatchAccount_ListDetectors": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Batch/batchAccounts/{{.accountName}}/detectors",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-06-01",
		},
		OperationID: "BatchAccount_ListDetectors",
		Resource:    "Microsoft.Batch",
	},
	"BatchAccount_GetDetector": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Batch/batchAccounts/{{.accountName}}/detectors/{{.detectorId}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-06-01",
		},
		OperationID: "BatchAccount_GetDetector",
		Resource:    "Microsoft.Batch",
	},
	"PrivateLinkResource_ListByBatchAccount": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Batch/batchAccounts/{{.accountName}}/privateLinkResources",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-06-01",
		},
		OperationID: "PrivateLinkResource_ListByBatchAccount",
		Resource:    "Microsoft.Batch",
	},
	"PrivateLinkResource_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Batch/batchAccounts/{{.accountName}}/privateLinkResources/{{.privateLinkResourceName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-06-01",
		},
		OperationID: "PrivateLinkResource_Get",
		Resource:    "Microsoft.Batch",
	},
	"PrivateEndpointConnection_ListByBatchAccount": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Batch/batchAccounts/{{.accountName}}/privateEndpointConnections",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-06-01",
		},
		OperationID: "PrivateEndpointConnection_ListByBatchAccount",
		Resource:    "Microsoft.Batch",
	},
	"PrivateEndpointConnection_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Batch/batchAccounts/{{.accountName}}/privateEndpointConnections/{{.privateEndpointConnectionName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-06-01",
		},
		OperationID: "PrivateEndpointConnection_Get",
		Resource:    "Microsoft.Batch",
	},
	"Pool_ListByBatchAccount": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Batch/batchAccounts/{{.accountName}}/pools",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-06-01",
		},
		OperationID: "Pool_ListByBatchAccount",
		Resource:    "Microsoft.Batch",
	},
	"Pool_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Batch/batchAccounts/{{.accountName}}/pools/{{.poolName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-06-01",
		},
		OperationID: "Pool_Get",
		Resource:    "Microsoft.Batch",
	},
	"Pool_DisableAutoScale": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Batch/batchAccounts/{{.accountName}}/pools/{{.poolName}}/disableAutoScale",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-06-01",
		},
		OperationID: "Pool_DisableAutoScale",
		Resource:    "Microsoft.Batch",
	},
	"Pool_StopResize": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Batch/batchAccounts/{{.accountName}}/pools/{{.poolName}}/stopResize",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-06-01",
		},
		OperationID: "Pool_StopResize",
		Resource:    "Microsoft.Batch",
	},
	"BatchAccount_ListOutboundNetworkDependenciesEndpoints": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Batch/batchAccounts/{{.accountName}}/outboundNetworkDependenciesEndpoints",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-06-01",
		},
		OperationID: "BatchAccount_ListOutboundNetworkDependenciesEndpoints",
		Resource:    "Microsoft.Batch",
	},
}
