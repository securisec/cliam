package policy

// Microsoft_Network_cloudServiceSwap policy
var Microsoft_Network_cloudServiceSwap = map[string]Policy{
	"VipSwap_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.groupName}}/providers/Microsoft.Compute/cloudServices/{{.resourceName}}/providers/Microsoft.Network/cloudServiceSlots/{{.singletonResource}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "VipSwap_Get",
		Resource:    "Microsoft.Network",
	},
	"VipSwap_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.groupName}}/providers/Microsoft.Compute/cloudServices/{{.resourceName}}/providers/Microsoft.Network/cloudServiceSlots",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "VipSwap_List",
		Resource:    "Microsoft.Network",
	},
}
