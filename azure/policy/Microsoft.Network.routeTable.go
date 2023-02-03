package policy

// Microsoft_Network_routeTable policy
var Microsoft_Network_routeTable = map[string]Policy{
	"RouteTables_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/routeTables/{{.routeTableName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "RouteTables_Get",
		Resource:    "Microsoft.Network",
	},
	"RouteTables_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/routeTables",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "RouteTables_List",
		Resource:    "Microsoft.Network",
	},
	"RouteTables_ListAll": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Network/routeTables",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "RouteTables_ListAll",
		Resource:    "Microsoft.Network",
	},
	"Routes_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/routeTables/{{.routeTableName}}/routes/{{.routeName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "Routes_Get",
		Resource:    "Microsoft.Network",
	},
	"Routes_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/routeTables/{{.routeTableName}}/routes",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "Routes_List",
		Resource:    "Microsoft.Network",
	},
}
