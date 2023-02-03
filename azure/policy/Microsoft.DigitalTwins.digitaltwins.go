package policy

// Microsoft_DigitalTwins_digitaltwins policy
var Microsoft_DigitalTwins_digitaltwins = map[string]Policy{
	"DigitalTwins_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DigitalTwins/digitalTwinsInstances/{{.resourceName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-10-31",
		},
		OperationID: "DigitalTwins_Get",
		Resource:    "Microsoft.DigitalTwins",
	},
	"DigitalTwinsEndpoint_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DigitalTwins/digitalTwinsInstances/{{.resourceName}}/endpoints",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-10-31",
		},
		OperationID: "DigitalTwinsEndpoint_List",
		Resource:    "Microsoft.DigitalTwins",
	},
	"DigitalTwinsEndpoint_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DigitalTwins/digitalTwinsInstances/{{.resourceName}}/endpoints/{{.endpointName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-10-31",
		},
		OperationID: "DigitalTwinsEndpoint_Get",
		Resource:    "Microsoft.DigitalTwins",
	},
	"DigitalTwins_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.DigitalTwins/digitalTwinsInstances",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-10-31",
		},
		OperationID: "DigitalTwins_List",
		Resource:    "Microsoft.DigitalTwins",
	},
	"DigitalTwins_ListByResourceGroup": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DigitalTwins/digitalTwinsInstances",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-10-31",
		},
		OperationID: "DigitalTwins_ListByResourceGroup",
		Resource:    "Microsoft.DigitalTwins",
	},
	"Operations_List": {
		Path:   "/providers/Microsoft.DigitalTwins/operations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-10-31",
		},
		OperationID: "Operations_List",
		Resource:    "Microsoft.DigitalTwins",
	},
	"DigitalTwins_CheckNameAvailability": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.DigitalTwins/locations/{{.location}}/checkNameAvailability",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-10-31",
		},
		OperationID: "DigitalTwins_CheckNameAvailability",
		Resource:    "Microsoft.DigitalTwins",
	},
	"PrivateLinkResources_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DigitalTwins/digitalTwinsInstances/{{.resourceName}}/privateLinkResources",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-10-31",
		},
		OperationID: "PrivateLinkResources_List",
		Resource:    "Microsoft.DigitalTwins",
	},
	"PrivateLinkResources_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DigitalTwins/digitalTwinsInstances/{{.resourceName}}/privateLinkResources/{{.resourceId}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-10-31",
		},
		OperationID: "PrivateLinkResources_Get",
		Resource:    "Microsoft.DigitalTwins",
	},
	"PrivateEndpointConnections_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DigitalTwins/digitalTwinsInstances/{{.resourceName}}/privateEndpointConnections",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-10-31",
		},
		OperationID: "PrivateEndpointConnections_List",
		Resource:    "Microsoft.DigitalTwins",
	},
	"PrivateEndpointConnections_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DigitalTwins/digitalTwinsInstances/{{.resourceName}}/privateEndpointConnections/{{.privateEndpointConnectionName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-10-31",
		},
		OperationID: "PrivateEndpointConnections_Get",
		Resource:    "Microsoft.DigitalTwins",
	},
	"TimeSeriesDatabaseConnections_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DigitalTwins/digitalTwinsInstances/{{.resourceName}}/timeSeriesDatabaseConnections",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-10-31",
		},
		OperationID: "TimeSeriesDatabaseConnections_List",
		Resource:    "Microsoft.DigitalTwins",
	},
	"TimeSeriesDatabaseConnections_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DigitalTwins/digitalTwinsInstances/{{.resourceName}}/timeSeriesDatabaseConnections/{{.timeSeriesDatabaseConnectionName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-10-31",
		},
		OperationID: "TimeSeriesDatabaseConnections_Get",
		Resource:    "Microsoft.DigitalTwins",
	},
}
