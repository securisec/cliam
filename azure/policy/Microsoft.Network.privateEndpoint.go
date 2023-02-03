package policy

// Microsoft_Network_privateEndpoint policy
var Microsoft_Network_privateEndpoint = map[string]Policy{
	"PrivateEndpoints_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/privateEndpoints/{{.privateEndpointName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "PrivateEndpoints_Get",
		Resource:    "Microsoft.Network",
	},
	"PrivateEndpoints_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/privateEndpoints",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "PrivateEndpoints_List",
		Resource:    "Microsoft.Network",
	},
	"PrivateEndpoints_ListBySubscription": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Network/privateEndpoints",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "PrivateEndpoints_ListBySubscription",
		Resource:    "Microsoft.Network",
	},
	"AvailablePrivateEndpointTypes_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Network/locations/{{.location}}/availablePrivateEndpointTypes",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "AvailablePrivateEndpointTypes_List",
		Resource:    "Microsoft.Network",
	},
	"AvailablePrivateEndpointTypes_ListByResourceGroup": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/locations/{{.location}}/availablePrivateEndpointTypes",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "AvailablePrivateEndpointTypes_ListByResourceGroup",
		Resource:    "Microsoft.Network",
	},
	"PrivateDnsZoneGroups_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/privateEndpoints/{{.privateEndpointName}}/privateDnsZoneGroups/{{.privateDnsZoneGroupName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "PrivateDnsZoneGroups_Get",
		Resource:    "Microsoft.Network",
	},
	"PrivateDnsZoneGroups_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/privateEndpoints/{{.privateEndpointName}}/privateDnsZoneGroups",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "PrivateDnsZoneGroups_List",
		Resource:    "Microsoft.Network",
	},
}
