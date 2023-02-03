package policy

// Microsoft_ApiManagement_apimdeployment policy
var Microsoft_ApiManagement_apimdeployment = map[string]Policy{
	"ApiManagementOperations_List": {
		Path:   "/providers/Microsoft.ApiManagement/operations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-08-01",
		},
		OperationID: "ApiManagementOperations_List",
		Resource:    "Microsoft.ApiManagement",
	},
	"ApiManagementServiceSkus_ListAvailableServiceSkus": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ApiManagement/service/{{.serviceName}}/skus",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-08-01",
		},
		OperationID: "ApiManagementServiceSkus_ListAvailableServiceSkus",
		Resource:    "Microsoft.ApiManagement",
	},
	"ApiManagementService_Restore": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ApiManagement/service/{{.serviceName}}/restore",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-08-01",
		},
		OperationID: "ApiManagementService_Restore",
		Resource:    "Microsoft.ApiManagement",
	},
	"ApiManagementService_Backup": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ApiManagement/service/{{.serviceName}}/backup",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-08-01",
		},
		OperationID: "ApiManagementService_Backup",
		Resource:    "Microsoft.ApiManagement",
	},
	"ApiManagementService_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ApiManagement/service/{{.serviceName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-08-01",
		},
		OperationID: "ApiManagementService_Get",
		Resource:    "Microsoft.ApiManagement",
	},
	"ApiManagementService_ListByResourceGroup": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ApiManagement/service",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-08-01",
		},
		OperationID: "ApiManagementService_ListByResourceGroup",
		Resource:    "Microsoft.ApiManagement",
	},
	"ApiManagementService_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.ApiManagement/service",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-08-01",
		},
		OperationID: "ApiManagementService_List",
		Resource:    "Microsoft.ApiManagement",
	},
	"ApiManagementService_GetSsoToken": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ApiManagement/service/{{.serviceName}}/getssotoken",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-08-01",
		},
		OperationID: "ApiManagementService_GetSsoToken",
		Resource:    "Microsoft.ApiManagement",
	},
	"ApiManagementService_CheckNameAvailability": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.ApiManagement/checkNameAvailability",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-08-01",
		},
		OperationID: "ApiManagementService_CheckNameAvailability",
		Resource:    "Microsoft.ApiManagement",
	},
	"ApiManagementService_GetDomainOwnershipIdentifier": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.ApiManagement/getDomainOwnershipIdentifier",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-08-01",
		},
		OperationID: "ApiManagementService_GetDomainOwnershipIdentifier",
		Resource:    "Microsoft.ApiManagement",
	},
	"ApiManagementService_ApplyNetworkConfigurationUpdates": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ApiManagement/service/{{.serviceName}}/applynetworkconfigurationupdates",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-08-01",
		},
		OperationID: "ApiManagementService_ApplyNetworkConfigurationUpdates",
		Resource:    "Microsoft.ApiManagement",
	},
}
