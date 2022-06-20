package policy

    var Microsoft_HybridNetwork_networkFunctionVendor = map[string]Policy{
        "networkFunctionVendors_List": {
    Path: "/subscriptions/{{.subscriptionId}}/providers/Microsoft.HybridNetwork/networkFunctionVendors",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2021-05-01",
    },
	OperationID:    "networkFunctionVendors_List",
    Resource:       "Microsoft.HybridNetwork",
},
"NetworkFunctionVendorSkus_ListByVendor": {
    Path: "/subscriptions/{{.subscriptionId}}/providers/Microsoft.HybridNetwork/networkFunctionVendors/{{.vendorName}}/vendorSkus",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2021-05-01",
    },
	OperationID:    "NetworkFunctionVendorSkus_ListByVendor",
    Resource:       "Microsoft.HybridNetwork",
},
"networkFunctionVendorSkus_ListBySku": {
    Path: "/subscriptions/{{.subscriptionId}}/providers/Microsoft.HybridNetwork/networkFunctionVendors/{{.vendorName}}/vendorSkus/{{.vendorSkuName}}",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2021-05-01",
    },
	OperationID:    "networkFunctionVendorSkus_ListBySku",
    Resource:       "Microsoft.HybridNetwork",
},
    }
    