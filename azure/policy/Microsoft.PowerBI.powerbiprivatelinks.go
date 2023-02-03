package policy

// Microsoft_PowerBI_powerbiprivatelinks policy
var Microsoft_PowerBI_powerbiprivatelinks = map[string]Policy{
	"Operations_List": {
		Path:   "/providers/Microsoft.PowerBI/operations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2020-06-01",
		},
		OperationID: "Operations_List",
		Resource:    "Microsoft.PowerBI",
	},
	"privateLinkServicesForPowerBI_ListBySubscriptionId": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.PowerBI/privateLinkServicesForPowerBI",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2020-06-01",
		},
		OperationID: "privateLinkServicesForPowerBI_ListBySubscriptionId",
		Resource:    "Microsoft.PowerBI",
	},
	"PrivateLinkServiceResourceOperationResults_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.PowerBI/operationResults/{{.operationId}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2020-06-01",
		},
		OperationID: "PrivateLinkServiceResourceOperationResults_Get",
		Resource:    "Microsoft.PowerBI",
	},
	"PrivateLinkServices_ListByResourceGroup": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.PowerBI/privateLinkServicesForPowerBI",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2020-06-01",
		},
		OperationID: "PrivateLinkServices_ListByResourceGroup",
		Resource:    "Microsoft.PowerBI",
	},
	"PowerBIResources_ListByResourceName": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.PowerBI/privateLinkServicesForPowerBI/{{.azureResourceName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2020-06-01",
		},
		OperationID: "PowerBIResources_ListByResourceName",
		Resource:    "Microsoft.PowerBI",
	},
	"PrivateLinkResources_ListByResource": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.PowerBI/privateLinkServicesForPowerBI/{{.azureResourceName}}/privateLinkResources",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2020-06-01",
		},
		OperationID: "PrivateLinkResources_ListByResource",
		Resource:    "Microsoft.PowerBI",
	},
	"PrivateLinkResources_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.PowerBI/privateLinkServicesForPowerBI/{{.azureResourceName}}/privateLinkResources/{{.privateLinkResourceName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2020-06-01",
		},
		OperationID: "PrivateLinkResources_Get",
		Resource:    "Microsoft.PowerBI",
	},
	"PrivateEndpointConnections_ListByResource": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.PowerBI/privateLinkServicesForPowerBI/{{.azureResourceName}}/privateEndpointConnections",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2020-06-01",
		},
		OperationID: "PrivateEndpointConnections_ListByResource",
		Resource:    "Microsoft.PowerBI",
	},
	"PrivateEndpointConnections_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.PowerBI/privateLinkServicesForPowerBI/{{.azureResourceName}}/privateEndpointConnections/{{.privateEndpointName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2020-06-01",
		},
		OperationID: "PrivateEndpointConnections_Get",
		Resource:    "Microsoft.PowerBI",
	},
}
