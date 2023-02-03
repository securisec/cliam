package policy

// Microsoft_MobileNetwork_service policy
var Microsoft_MobileNetwork_service = map[string]Policy{
	"Services_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.MobileNetwork/mobileNetworks/{{.mobileNetworkName}}/services/{{.serviceName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-11-01",
		},
		OperationID: "Services_Get",
		Resource:    "Microsoft.MobileNetwork",
	},
	"Services_ListByMobileNetwork": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.MobileNetwork/mobileNetworks/{{.mobileNetworkName}}/services",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-11-01",
		},
		OperationID: "Services_ListByMobileNetwork",
		Resource:    "Microsoft.MobileNetwork",
	},
}
