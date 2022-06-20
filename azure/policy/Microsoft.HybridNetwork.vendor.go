package policy

var Microsoft_HybridNetwork_vendor = map[string]Policy{
	"Vendors_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.HybridNetwork/vendors/{{.vendorName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-05-01",
		},
		OperationID: "Vendors_Get",
		Resource:    "Microsoft.HybridNetwork",
	},
	"Vendors_ListBySubscription": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.HybridNetwork/vendors",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-05-01",
		},
		OperationID: "Vendors_ListBySubscription",
		Resource:    "Microsoft.HybridNetwork",
	},
	"VendorSkus_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.HybridNetwork/vendors/{{.vendorName}}/vendorSkus/{{.skuName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-05-01",
		},
		OperationID: "VendorSkus_Get",
		Resource:    "Microsoft.HybridNetwork",
	},
	"VendorSkus_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.HybridNetwork/vendors/{{.vendorName}}/vendorSkus",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-05-01",
		},
		OperationID: "VendorSkus_List",
		Resource:    "Microsoft.HybridNetwork",
	},
	"VendorSkuPreview_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.HybridNetwork/vendors/{{.vendorName}}/vendorSkus/{{.skuName}}/previewSubscriptions",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-05-01",
		},
		OperationID: "VendorSkuPreview_List",
		Resource:    "Microsoft.HybridNetwork",
	},
	"VendorSkuPreview_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.HybridNetwork/vendors/{{.vendorName}}/vendorSkus/{{.skuName}}/previewSubscriptions/{{.previewSubscription}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-05-01",
		},
		OperationID: "VendorSkuPreview_Get",
		Resource:    "Microsoft.HybridNetwork",
	},
}
