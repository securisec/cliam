package policy

var Microsoft_SecurityAndCompliance_privateLinkServicesForSCCPowershell = map[string]Policy{
	"privateLinkServicesForSCCPowershell_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.SecurityAndCompliance/privateLinkServicesForSCCPowershell/{{.resourceName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-08",
		},
		OperationID: "privateLinkServicesForSCCPowershell_Get",
		Resource:    "Microsoft.SecurityAndCompliance",
	},
	"privateLinkServicesForSCCPowershell_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.SecurityAndCompliance/privateLinkServicesForSCCPowershell",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-08",
		},
		OperationID: "privateLinkServicesForSCCPowershell_List",
		Resource:    "Microsoft.SecurityAndCompliance",
	},
	"privateLinkServicesForSCCPowershell_ListByResourceGroup": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.SecurityAndCompliance/privateLinkServicesForSCCPowershell",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-08",
		},
		OperationID: "privateLinkServicesForSCCPowershell_ListByResourceGroup",
		Resource:    "Microsoft.SecurityAndCompliance",
	},
	"PrivateEndpointConnectionsForSCCPowershell_ListByService": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.SecurityAndCompliance/privateLinkServicesForSCCPowershell/{{.resourceName}}/privateEndpointConnections",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-08",
		},
		OperationID: "PrivateEndpointConnectionsForSCCPowershell_ListByService",
		Resource:    "Microsoft.SecurityAndCompliance",
	},
	"PrivateEndpointConnectionsForSCCPowershell_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.SecurityAndCompliance/privateLinkServicesForSCCPowershell/{{.resourceName}}/privateEndpointConnections/{{.privateEndpointConnectionName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-08",
		},
		OperationID: "PrivateEndpointConnectionsForSCCPowershell_Get",
		Resource:    "Microsoft.SecurityAndCompliance",
	},
	"PrivateLinkResourcesForSCCPowershell_ListByService": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.SecurityAndCompliance/privateLinkServicesForSCCPowershell/{{.resourceName}}/privateLinkResources",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-08",
		},
		OperationID: "PrivateLinkResourcesForSCCPowershell_ListByService",
		Resource:    "Microsoft.SecurityAndCompliance",
	},
	"PrivateLinkResourcesForSCCPowershell_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.SecurityAndCompliance/privateLinkServicesForSCCPowershell/{{.resourceName}}/privateLinkResources/{{.groupName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-08",
		},
		OperationID: "PrivateLinkResourcesForSCCPowershell_Get",
		Resource:    "Microsoft.SecurityAndCompliance",
	},
}
