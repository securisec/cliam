package policy

// Microsoft_Network_publicIpPrefix policy
var Microsoft_Network_publicIpPrefix = map[string]Policy{
	"PublicIPPrefixes_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/publicIPPrefixes/{{.publicIpPrefixName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "PublicIPPrefixes_Get",
		Resource:    "Microsoft.Network",
	},
	"PublicIPPrefixes_ListAll": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Network/publicIPPrefixes",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "PublicIPPrefixes_ListAll",
		Resource:    "Microsoft.Network",
	},
	"PublicIPPrefixes_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/publicIPPrefixes",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "PublicIPPrefixes_List",
		Resource:    "Microsoft.Network",
	},
}
