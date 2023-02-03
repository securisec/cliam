package policy

// Microsoft_Network_expressRouteCircuit policy
var Microsoft_Network_expressRouteCircuit = map[string]Policy{
	"ExpressRouteCircuitAuthorizations_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/expressRouteCircuits/{{.circuitName}}/authorizations/{{.authorizationName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "ExpressRouteCircuitAuthorizations_Get",
		Resource:    "Microsoft.Network",
	},
	"ExpressRouteCircuitAuthorizations_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/expressRouteCircuits/{{.circuitName}}/authorizations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "ExpressRouteCircuitAuthorizations_List",
		Resource:    "Microsoft.Network",
	},
	"ExpressRouteCircuitPeerings_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/expressRouteCircuits/{{.circuitName}}/peerings/{{.peeringName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "ExpressRouteCircuitPeerings_Get",
		Resource:    "Microsoft.Network",
	},
	"ExpressRouteCircuitPeerings_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/expressRouteCircuits/{{.circuitName}}/peerings",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "ExpressRouteCircuitPeerings_List",
		Resource:    "Microsoft.Network",
	},
	"ExpressRouteCircuitConnections_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/expressRouteCircuits/{{.circuitName}}/peerings/{{.peeringName}}/connections/{{.connectionName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "ExpressRouteCircuitConnections_Get",
		Resource:    "Microsoft.Network",
	},
	"ExpressRouteCircuitConnections_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/expressRouteCircuits/{{.circuitName}}/peerings/{{.peeringName}}/connections",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "ExpressRouteCircuitConnections_List",
		Resource:    "Microsoft.Network",
	},
	"PeerExpressRouteCircuitConnections_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/expressRouteCircuits/{{.circuitName}}/peerings/{{.peeringName}}/peerConnections/{{.connectionName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "PeerExpressRouteCircuitConnections_Get",
		Resource:    "Microsoft.Network",
	},
	"PeerExpressRouteCircuitConnections_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/expressRouteCircuits/{{.circuitName}}/peerings/{{.peeringName}}/peerConnections",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "PeerExpressRouteCircuitConnections_List",
		Resource:    "Microsoft.Network",
	},
	"ExpressRouteCircuits_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/expressRouteCircuits/{{.circuitName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "ExpressRouteCircuits_Get",
		Resource:    "Microsoft.Network",
	},
	"ExpressRouteCircuits_ListArpTable": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/expressRouteCircuits/{{.circuitName}}/peerings/{{.peeringName}}/arpTables/{{.devicePath}}",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "ExpressRouteCircuits_ListArpTable",
		Resource:    "Microsoft.Network",
	},
	"ExpressRouteCircuits_ListRoutesTable": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/expressRouteCircuits/{{.circuitName}}/peerings/{{.peeringName}}/routeTables/{{.devicePath}}",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "ExpressRouteCircuits_ListRoutesTable",
		Resource:    "Microsoft.Network",
	},
	"ExpressRouteCircuits_ListRoutesTableSummary": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/expressRouteCircuits/{{.circuitName}}/peerings/{{.peeringName}}/routeTablesSummary/{{.devicePath}}",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "ExpressRouteCircuits_ListRoutesTableSummary",
		Resource:    "Microsoft.Network",
	},
	"ExpressRouteCircuits_GetStats": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/expressRouteCircuits/{{.circuitName}}/stats",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "ExpressRouteCircuits_GetStats",
		Resource:    "Microsoft.Network",
	},
	"ExpressRouteCircuits_GetPeeringStats": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/expressRouteCircuits/{{.circuitName}}/peerings/{{.peeringName}}/stats",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "ExpressRouteCircuits_GetPeeringStats",
		Resource:    "Microsoft.Network",
	},
	"ExpressRouteCircuits_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/expressRouteCircuits",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "ExpressRouteCircuits_List",
		Resource:    "Microsoft.Network",
	},
	"ExpressRouteCircuits_ListAll": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Network/expressRouteCircuits",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "ExpressRouteCircuits_ListAll",
		Resource:    "Microsoft.Network",
	},
	"ExpressRouteServiceProviders_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Network/expressRouteServiceProviders",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "ExpressRouteServiceProviders_List",
		Resource:    "Microsoft.Network",
	},
}
