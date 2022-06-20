package policy

var Microsoft_SecurityAndCompliance_privateLinkServicesForM365ComplianceCenter = map[string]Policy{
	"privateLinkServicesForM365ComplianceCenter_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.SecurityAndCompliance/privateLinkServicesForM365ComplianceCenter/{{.resourceName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-08",
		},
		OperationID: "privateLinkServicesForM365ComplianceCenter_Get",
		Resource:    "Microsoft.SecurityAndCompliance",
	},
	"privateLinkServicesForM365ComplianceCenter_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.SecurityAndCompliance/privateLinkServicesForM365ComplianceCenter",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-08",
		},
		OperationID: "privateLinkServicesForM365ComplianceCenter_List",
		Resource:    "Microsoft.SecurityAndCompliance",
	},
	"privateLinkServicesForM365ComplianceCenter_ListByResourceGroup": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.SecurityAndCompliance/privateLinkServicesForM365ComplianceCenter",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-08",
		},
		OperationID: "privateLinkServicesForM365ComplianceCenter_ListByResourceGroup",
		Resource:    "Microsoft.SecurityAndCompliance",
	},
	"PrivateEndpointConnectionsComp_ListByService": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.SecurityAndCompliance/privateLinkServicesForM365ComplianceCenter/{{.resourceName}}/privateEndpointConnections",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-08",
		},
		OperationID: "PrivateEndpointConnectionsComp_ListByService",
		Resource:    "Microsoft.SecurityAndCompliance",
	},
	"PrivateEndpointConnectionsComp_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.SecurityAndCompliance/privateLinkServicesForM365ComplianceCenter/{{.resourceName}}/privateEndpointConnections/{{.privateEndpointConnectionName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-08",
		},
		OperationID: "PrivateEndpointConnectionsComp_Get",
		Resource:    "Microsoft.SecurityAndCompliance",
	},
	"PrivateLinkResourcesComp_ListByService": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.SecurityAndCompliance/privateLinkServicesForM365ComplianceCenter/{{.resourceName}}/privateLinkResources",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-08",
		},
		OperationID: "PrivateLinkResourcesComp_ListByService",
		Resource:    "Microsoft.SecurityAndCompliance",
	},
	"PrivateLinkResourcesComp_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.SecurityAndCompliance/privateLinkServicesForM365ComplianceCenter/{{.resourceName}}/privateLinkResources/{{.groupName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-08",
		},
		OperationID: "PrivateLinkResourcesComp_Get",
		Resource:    "Microsoft.SecurityAndCompliance",
	},
}
