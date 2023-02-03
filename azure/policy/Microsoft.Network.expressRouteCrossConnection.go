package policy

// Microsoft_Network_expressRouteCrossConnection policy
var Microsoft_Network_expressRouteCrossConnection = map[string]Policy{
	"ExpressRouteCrossConnections_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Network/expressRouteCrossConnections",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "ExpressRouteCrossConnections_List",
		Resource:    "Microsoft.Network",
	},
	"ExpressRouteCrossConnections_ListByResourceGroup": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/expressRouteCrossConnections",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "ExpressRouteCrossConnections_ListByResourceGroup",
		Resource:    "Microsoft.Network",
	},
	"ExpressRouteCrossConnections_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/expressRouteCrossConnections/{{.crossConnectionName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "ExpressRouteCrossConnections_Get",
		Resource:    "Microsoft.Network",
	},
	"ExpressRouteCrossConnectionPeerings_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/expressRouteCrossConnections/{{.crossConnectionName}}/peerings",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "ExpressRouteCrossConnectionPeerings_List",
		Resource:    "Microsoft.Network",
	},
	"ExpressRouteCrossConnectionPeerings_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/expressRouteCrossConnections/{{.crossConnectionName}}/peerings/{{.peeringName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "ExpressRouteCrossConnectionPeerings_Get",
		Resource:    "Microsoft.Network",
	},
	"ExpressRouteCrossConnections_ListArpTable": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/expressRouteCrossConnections/{{.crossConnectionName}}/peerings/{{.peeringName}}/arpTables/{{.devicePath}}",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "ExpressRouteCrossConnections_ListArpTable",
		Resource:    "Microsoft.Network",
	},
	"ExpressRouteCrossConnections_ListRoutesTableSummary": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/expressRouteCrossConnections/{{.crossConnectionName}}/peerings/{{.peeringName}}/routeTablesSummary/{{.devicePath}}",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "ExpressRouteCrossConnections_ListRoutesTableSummary",
		Resource:    "Microsoft.Network",
	},
	"ExpressRouteCrossConnections_ListRoutesTable": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/expressRouteCrossConnections/{{.crossConnectionName}}/peerings/{{.peeringName}}/routeTables/{{.devicePath}}",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "ExpressRouteCrossConnections_ListRoutesTable",
		Resource:    "Microsoft.Network",
	},
}
