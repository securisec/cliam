package policy

// Microsoft_MobileNetwork_slice policy
var Microsoft_MobileNetwork_slice = map[string]Policy{
	"Slices_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.MobileNetwork/mobileNetworks/{{.mobileNetworkName}}/slices/{{.sliceName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-11-01",
		},
		OperationID: "Slices_Get",
		Resource:    "Microsoft.MobileNetwork",
	},
	"Slices_ListByMobileNetwork": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.MobileNetwork/mobileNetworks/{{.mobileNetworkName}}/slices",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-11-01",
		},
		OperationID: "Slices_ListByMobileNetwork",
		Resource:    "Microsoft.MobileNetwork",
	},
}
