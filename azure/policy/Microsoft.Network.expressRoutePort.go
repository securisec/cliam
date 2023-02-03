package policy

// Microsoft_Network_expressRoutePort policy
var Microsoft_Network_expressRoutePort = map[string]Policy{
	"ExpressRoutePortsLocations_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Network/ExpressRoutePortsLocations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "ExpressRoutePortsLocations_List",
		Resource:    "Microsoft.Network",
	},
	"ExpressRoutePortsLocations_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Network/ExpressRoutePortsLocations/{{.locationName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "ExpressRoutePortsLocations_Get",
		Resource:    "Microsoft.Network",
	},
	"ExpressRoutePorts_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/ExpressRoutePorts/{{.expressRoutePortName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "ExpressRoutePorts_Get",
		Resource:    "Microsoft.Network",
	},
	"ExpressRoutePorts_ListByResourceGroup": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/ExpressRoutePorts",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "ExpressRoutePorts_ListByResourceGroup",
		Resource:    "Microsoft.Network",
	},
	"ExpressRoutePorts_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Network/ExpressRoutePorts",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "ExpressRoutePorts_List",
		Resource:    "Microsoft.Network",
	},
	"ExpressRouteLinks_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/ExpressRoutePorts/{{.expressRoutePortName}}/links/{{.linkName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "ExpressRouteLinks_Get",
		Resource:    "Microsoft.Network",
	},
	"ExpressRouteLinks_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/ExpressRoutePorts/{{.expressRoutePortName}}/links",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "ExpressRouteLinks_List",
		Resource:    "Microsoft.Network",
	},
	"ExpressRoutePorts_GenerateLOA": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/expressRoutePorts/{{.expressRoutePortName}}/generateLoa",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "ExpressRoutePorts_GenerateLOA",
		Resource:    "Microsoft.Network",
	},
	"ExpressRoutePortAuthorizations_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/expressRoutePorts/{{.expressRoutePortName}}/authorizations/{{.authorizationName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "ExpressRoutePortAuthorizations_Get",
		Resource:    "Microsoft.Network",
	},
	"ExpressRoutePortAuthorizations_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/expressRoutePorts/{{.expressRoutePortName}}/authorizations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "ExpressRoutePortAuthorizations_List",
		Resource:    "Microsoft.Network",
	},
}
