package policy

// Microsoft_Network_natGateway policy
var Microsoft_Network_natGateway = map[string]Policy{
	"NatGateways_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/natGateways/{{.natGatewayName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "NatGateways_Get",
		Resource:    "Microsoft.Network",
	},
	"NatGateways_ListAll": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Network/natGateways",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "NatGateways_ListAll",
		Resource:    "Microsoft.Network",
	},
	"NatGateways_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/natGateways",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "NatGateways_List",
		Resource:    "Microsoft.Network",
	},
}
