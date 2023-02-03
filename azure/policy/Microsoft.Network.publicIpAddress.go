package policy

// Microsoft_Network_publicIpAddress policy
var Microsoft_Network_publicIpAddress = map[string]Policy{
	"PublicIPAddresses_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/publicIPAddresses/{{.publicIpAddressName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "PublicIPAddresses_Get",
		Resource:    "Microsoft.Network",
	},
	"PublicIPAddresses_ListAll": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Network/publicIPAddresses",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "PublicIPAddresses_ListAll",
		Resource:    "Microsoft.Network",
	},
	"PublicIPAddresses_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/publicIPAddresses",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "PublicIPAddresses_List",
		Resource:    "Microsoft.Network",
	},
	"PublicIPAddresses_DdosProtectionStatus": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/publicIPAddresses/{{.publicIpAddressName}}/ddosProtectionStatus",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "PublicIPAddresses_DdosProtectionStatus",
		Resource:    "Microsoft.Network",
	},
}
