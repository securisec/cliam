package policy

// Microsoft_ApiManagement_apimproducts policy
var Microsoft_ApiManagement_apimproducts = map[string]Policy{
	"Product_ListByService": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ApiManagement/service/{{.serviceName}}/products",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-08-01",
		},
		OperationID: "Product_ListByService",
		Resource:    "Microsoft.ApiManagement",
	},
	"Product_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ApiManagement/service/{{.serviceName}}/products/{{.productId}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-08-01",
		},
		OperationID: "Product_Get",
		Resource:    "Microsoft.ApiManagement",
	},
	"ProductApi_ListByProduct": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ApiManagement/service/{{.serviceName}}/products/{{.productId}}/apis",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-08-01",
		},
		OperationID: "ProductApi_ListByProduct",
		Resource:    "Microsoft.ApiManagement",
	},
	"ProductGroup_ListByProduct": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ApiManagement/service/{{.serviceName}}/products/{{.productId}}/groups",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-08-01",
		},
		OperationID: "ProductGroup_ListByProduct",
		Resource:    "Microsoft.ApiManagement",
	},
	"ProductSubscriptions_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ApiManagement/service/{{.serviceName}}/products/{{.productId}}/subscriptions",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-08-01",
		},
		OperationID: "ProductSubscriptions_List",
		Resource:    "Microsoft.ApiManagement",
	},
	"ProductPolicy_ListByProduct": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ApiManagement/service/{{.serviceName}}/products/{{.productId}}/policies",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-08-01",
		},
		OperationID: "ProductPolicy_ListByProduct",
		Resource:    "Microsoft.ApiManagement",
	},
	"ProductPolicy_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ApiManagement/service/{{.serviceName}}/products/{{.productId}}/policies/{{.policyId}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-08-01",
		},
		OperationID: "ProductPolicy_Get",
		Resource:    "Microsoft.ApiManagement",
	},
	"Tag_ListByProduct": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ApiManagement/service/{{.serviceName}}/products/{{.productId}}/tags",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-08-01",
		},
		OperationID: "Tag_ListByProduct",
		Resource:    "Microsoft.ApiManagement",
	},
	"Tag_GetByProduct": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ApiManagement/service/{{.serviceName}}/products/{{.productId}}/tags/{{.tagId}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-08-01",
		},
		OperationID: "Tag_GetByProduct",
		Resource:    "Microsoft.ApiManagement",
	},
}
