package policy

var Microsoft_AzureStack_Product = []Policy{
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroup}}/providers/Microsoft.AzureStack/registrations/{{.registrationName}}/products",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2017-06-01",
		},
		OperationID: "Products_List",
		Resource:    "Microsoft.AzureStack",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroup}}/providers/Microsoft.AzureStack/registrations/{{.registrationName}}/products/{{.productName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2017-06-01",
		},
		OperationID: "Products_Get",
		Resource:    "Microsoft.AzureStack",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroup}}/providers/Microsoft.AzureStack/registrations/{{.registrationName}}/products/{{.productName}}/listDetails",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2017-06-01",
		},
		OperationID: "Products_ListDetails",
		Resource:    "Microsoft.AzureStack",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroup}}/providers/Microsoft.AzureStack/registrations/{{.registrationName}}/products/{{.productName}}/getProducts",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2017-06-01",
		},
		OperationID: "Products_GetProducts",
		Resource:    "Microsoft.AzureStack",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroup}}/providers/Microsoft.AzureStack/registrations/{{.registrationName}}/products/{{.productName}}/getProduct",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2017-06-01",
		},
		OperationID: "Products_GetProduct",
		Resource:    "Microsoft.AzureStack",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroup}}/providers/Microsoft.AzureStack/registrations/{{.registrationName}}/products/{{.productName}}/uploadProductLog",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2017-06-01",
		},
		OperationID: "Products_UploadLog",
		Resource:    "Microsoft.AzureStack",
	},
}
