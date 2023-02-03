package policy

// Microsoft_Network_customIpPrefix policy
var Microsoft_Network_customIpPrefix = map[string]Policy{
	"CustomIPPrefixes_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/customIpPrefixes/{{.customIpPrefixName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "CustomIPPrefixes_Get",
		Resource:    "Microsoft.Network",
	},
	"CustomIPPrefixes_ListAll": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Network/customIpPrefixes",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "CustomIPPrefixes_ListAll",
		Resource:    "Microsoft.Network",
	},
	"CustomIPPrefixes_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/customIpPrefixes",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "CustomIPPrefixes_List",
		Resource:    "Microsoft.Network",
	},
}
