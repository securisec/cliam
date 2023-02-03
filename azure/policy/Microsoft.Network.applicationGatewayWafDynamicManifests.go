package policy

// Microsoft_Network_applicationGatewayWafDynamicManifests policy
var Microsoft_Network_applicationGatewayWafDynamicManifests = map[string]Policy{
	"ApplicationGatewayWafDynamicManifestsDefault_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Network/locations/{{.location}}/applicationGatewayWafDynamicManifests/dafault",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "ApplicationGatewayWafDynamicManifestsDefault_Get",
		Resource:    "Microsoft.Network",
	},
	"ApplicationGatewayWafDynamicManifests_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Network/locations/{{.location}}/applicationGatewayWafDynamicManifests",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "ApplicationGatewayWafDynamicManifests_Get",
		Resource:    "Microsoft.Network",
	},
}
