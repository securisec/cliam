package policy

var Microsoft_AzureStack_Product = map[string]Policy{
	"Products_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroup}}/providers/Microsoft.AzureStack/registrations/{{.registrationName}}/products",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2017-06-01",
		},
		OperationID: "Products_List",
		Resource:    "Microsoft.AzureStack",
	},
	"Products_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroup}}/providers/Microsoft.AzureStack/registrations/{{.registrationName}}/products/{{.productName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2017-06-01",
		},
		OperationID: "Products_Get",
		Resource:    "Microsoft.AzureStack",
	},
	"Products_ListDetails": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroup}}/providers/Microsoft.AzureStack/registrations/{{.registrationName}}/products/{{.productName}}/listDetails",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2017-06-01",
		},
		OperationID: "Products_ListDetails",
		Resource:    "Microsoft.AzureStack",
	},
	"Products_GetProducts": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroup}}/providers/Microsoft.AzureStack/registrations/{{.registrationName}}/products/{{.productName}}/getProducts",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2017-06-01",
		},
		OperationID: "Products_GetProducts",
		Resource:    "Microsoft.AzureStack",
	},
	"Products_GetProduct": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroup}}/providers/Microsoft.AzureStack/registrations/{{.registrationName}}/products/{{.productName}}/getProduct",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2017-06-01",
		},
		OperationID: "Products_GetProduct",
		Resource:    "Microsoft.AzureStack",
	},
	"Products_UploadLog": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroup}}/providers/Microsoft.AzureStack/registrations/{{.registrationName}}/products/{{.productName}}/uploadProductLog",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2017-06-01",
		},
		OperationID: "Products_UploadLog",
		Resource:    "Microsoft.AzureStack",
	},
}
