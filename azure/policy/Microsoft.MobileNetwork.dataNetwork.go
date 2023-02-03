package policy

// Microsoft_MobileNetwork_dataNetwork policy
var Microsoft_MobileNetwork_dataNetwork = map[string]Policy{
	"DataNetworks_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.MobileNetwork/mobileNetworks/{{.mobileNetworkName}}/dataNetworks/{{.dataNetworkName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-11-01",
		},
		OperationID: "DataNetworks_Get",
		Resource:    "Microsoft.MobileNetwork",
	},
	"DataNetworks_ListByMobileNetwork": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.MobileNetwork/mobileNetworks/{{.mobileNetworkName}}/dataNetworks",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-11-01",
		},
		OperationID: "DataNetworks_ListByMobileNetwork",
		Resource:    "Microsoft.MobileNetwork",
	},
}
