package policy

// Microsoft_Network_dscpConfiguration policy
var Microsoft_Network_dscpConfiguration = map[string]Policy{
	"DscpConfiguration_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/dscpConfigurations/{{.dscpConfigurationName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "DscpConfiguration_Get",
		Resource:    "Microsoft.Network",
	},
	"DscpConfiguration_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/dscpConfigurations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "DscpConfiguration_List",
		Resource:    "Microsoft.Network",
	},
	"DscpConfiguration_ListAll": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Network/dscpConfigurations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "DscpConfiguration_ListAll",
		Resource:    "Microsoft.Network",
	},
}
