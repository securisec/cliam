package policy

// Microsoft_Network_securityPartnerProvider policy
var Microsoft_Network_securityPartnerProvider = map[string]Policy{
	"SecurityPartnerProviders_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/securityPartnerProviders/{{.securityPartnerProviderName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "SecurityPartnerProviders_Get",
		Resource:    "Microsoft.Network",
	},
	"SecurityPartnerProviders_ListByResourceGroup": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/securityPartnerProviders",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "SecurityPartnerProviders_ListByResourceGroup",
		Resource:    "Microsoft.Network",
	},
	"SecurityPartnerProviders_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Network/securityPartnerProviders",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "SecurityPartnerProviders_List",
		Resource:    "Microsoft.Network",
	},
}
