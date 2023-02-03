package policy

// Microsoft_Network_expressRouteProviderPort policy
var Microsoft_Network_expressRouteProviderPort = map[string]Policy{
	"ExpressRouteProviderPortsLocation_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Network/expressRouteProviderPorts",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "ExpressRouteProviderPortsLocation_List",
		Resource:    "Microsoft.Network",
	},
	"ExpressRouteProviderPort": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Network/expressRouteProviderPorts/{{.providerport}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "ExpressRouteProviderPort",
		Resource:    "Microsoft.Network",
	},
}
