package policy

// Microsoft_Network_networkVirtualAppliance policy
var Microsoft_Network_networkVirtualAppliance = map[string]Policy{
	"NetworkVirtualAppliances_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/networkVirtualAppliances/{{.networkVirtualApplianceName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "NetworkVirtualAppliances_Get",
		Resource:    "Microsoft.Network",
	},
	"NetworkVirtualAppliances_ListByResourceGroup": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/networkVirtualAppliances",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "NetworkVirtualAppliances_ListByResourceGroup",
		Resource:    "Microsoft.Network",
	},
	"NetworkVirtualAppliances_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Network/networkVirtualAppliances",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "NetworkVirtualAppliances_List",
		Resource:    "Microsoft.Network",
	},
	"VirtualApplianceSites_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/networkVirtualAppliances/{{.networkVirtualApplianceName}}/virtualApplianceSites/{{.siteName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "VirtualApplianceSites_Get",
		Resource:    "Microsoft.Network",
	},
	"VirtualApplianceSites_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/networkVirtualAppliances/{{.networkVirtualApplianceName}}/virtualApplianceSites",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "VirtualApplianceSites_List",
		Resource:    "Microsoft.Network",
	},
	"VirtualApplianceSkus_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Network/networkVirtualApplianceSkus",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "VirtualApplianceSkus_List",
		Resource:    "Microsoft.Network",
	},
	"VirtualApplianceSkus_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Network/networkVirtualApplianceSkus/{{.skuName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "VirtualApplianceSkus_Get",
		Resource:    "Microsoft.Network",
	},
}
