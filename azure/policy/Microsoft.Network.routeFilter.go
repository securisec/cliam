package policy

// Microsoft_Network_routeFilter policy
var Microsoft_Network_routeFilter = map[string]Policy{
	"RouteFilters_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/routeFilters/{{.routeFilterName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "RouteFilters_Get",
		Resource:    "Microsoft.Network",
	},
	"RouteFilters_ListByResourceGroup": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/routeFilters",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "RouteFilters_ListByResourceGroup",
		Resource:    "Microsoft.Network",
	},
	"RouteFilters_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Network/routeFilters",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "RouteFilters_List",
		Resource:    "Microsoft.Network",
	},
	"RouteFilterRules_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/routeFilters/{{.routeFilterName}}/routeFilterRules/{{.ruleName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "RouteFilterRules_Get",
		Resource:    "Microsoft.Network",
	},
	"RouteFilterRules_ListByRouteFilter": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/routeFilters/{{.routeFilterName}}/routeFilterRules",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "RouteFilterRules_ListByRouteFilter",
		Resource:    "Microsoft.Network",
	},
}
