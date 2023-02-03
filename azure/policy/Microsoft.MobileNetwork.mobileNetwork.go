package policy

// Microsoft_MobileNetwork_mobileNetwork policy
var Microsoft_MobileNetwork_mobileNetwork = map[string]Policy{
	"MobileNetworks_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.MobileNetwork/mobileNetworks/{{.mobileNetworkName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-11-01",
		},
		OperationID: "MobileNetworks_Get",
		Resource:    "Microsoft.MobileNetwork",
	},
	"MobileNetworks_ListBySubscription": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.MobileNetwork/mobileNetworks",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-11-01",
		},
		OperationID: "MobileNetworks_ListBySubscription",
		Resource:    "Microsoft.MobileNetwork",
	},
	"MobileNetworks_ListByResourceGroup": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.MobileNetwork/mobileNetworks",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-11-01",
		},
		OperationID: "MobileNetworks_ListByResourceGroup",
		Resource:    "Microsoft.MobileNetwork",
	},
}
