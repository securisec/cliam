package policy

// Microsoft_MobileNetwork_site policy
var Microsoft_MobileNetwork_site = map[string]Policy{
	"Sites_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.MobileNetwork/mobileNetworks/{{.mobileNetworkName}}/sites/{{.siteName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-11-01",
		},
		OperationID: "Sites_Get",
		Resource:    "Microsoft.MobileNetwork",
	},
	"Sites_ListByMobileNetwork": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.MobileNetwork/mobileNetworks/{{.mobileNetworkName}}/sites",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-11-01",
		},
		OperationID: "Sites_ListByMobileNetwork",
		Resource:    "Microsoft.MobileNetwork",
	},
}
