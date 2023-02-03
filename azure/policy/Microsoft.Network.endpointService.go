package policy

// Microsoft_Network_endpointService policy
var Microsoft_Network_endpointService = map[string]Policy{
	"AvailableEndpointServices_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Network/locations/{{.location}}/virtualNetworkAvailableEndpointServices",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "AvailableEndpointServices_List",
		Resource:    "Microsoft.Network",
	},
}
