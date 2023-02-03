package policy

// Microsoft_EdgeOrder_edgeorder policy
var Microsoft_EdgeOrder_edgeorder = map[string]Policy{
	"ListOperations": {
		Path:   "/providers/Microsoft.EdgeOrder/operations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-12-01",
		},
		OperationID: "ListOperations",
		Resource:    "Microsoft.EdgeOrder",
	},
	"ListAddressesAtSubscriptionLevel": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.EdgeOrder/addresses",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-12-01",
		},
		OperationID: "ListAddressesAtSubscriptionLevel",
		Resource:    "Microsoft.EdgeOrder",
	},
	"ListProductFamilies": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.EdgeOrder/listProductFamilies",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-12-01",
		},
		OperationID: "ListProductFamilies",
		Resource:    "Microsoft.EdgeOrder",
	},
	"ListConfigurations": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.EdgeOrder/listConfigurations",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-12-01",
		},
		OperationID: "ListConfigurations",
		Resource:    "Microsoft.EdgeOrder",
	},
	"ListProductFamiliesMetadata": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.EdgeOrder/productFamiliesMetadata",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-12-01",
		},
		OperationID: "ListProductFamiliesMetadata",
		Resource:    "Microsoft.EdgeOrder",
	},
	"ListOrderAtSubscriptionLevel": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.EdgeOrder/orders",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-12-01",
		},
		OperationID: "ListOrderAtSubscriptionLevel",
		Resource:    "Microsoft.EdgeOrder",
	},
	"ListOrderItemsAtSubscriptionLevel": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.EdgeOrder/orderItems",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-12-01",
		},
		OperationID: "ListOrderItemsAtSubscriptionLevel",
		Resource:    "Microsoft.EdgeOrder",
	},
	"ListAddressesAtResourceGroupLevel": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.EdgeOrder/addresses",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-12-01",
		},
		OperationID: "ListAddressesAtResourceGroupLevel",
		Resource:    "Microsoft.EdgeOrder",
	},
	"GetAddressByName": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.EdgeOrder/addresses/{{.addressName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-12-01",
		},
		OperationID: "GetAddressByName",
		Resource:    "Microsoft.EdgeOrder",
	},
	"ListOrderAtResourceGroupLevel": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.EdgeOrder/orders",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-12-01",
		},
		OperationID: "ListOrderAtResourceGroupLevel",
		Resource:    "Microsoft.EdgeOrder",
	},
	"GetOrderByName": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.EdgeOrder/locations/{{.location}}/orders/{{.orderName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-12-01",
		},
		OperationID: "GetOrderByName",
		Resource:    "Microsoft.EdgeOrder",
	},
	"ListOrderItemsAtResourceGroupLevel": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.EdgeOrder/orderItems",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-12-01",
		},
		OperationID: "ListOrderItemsAtResourceGroupLevel",
		Resource:    "Microsoft.EdgeOrder",
	},
	"GetOrderItemByName": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.EdgeOrder/orderItems/{{.orderItemName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-12-01",
		},
		OperationID: "GetOrderItemByName",
		Resource:    "Microsoft.EdgeOrder",
	},
	"CancelOrderItem": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.EdgeOrder/orderItems/{{.orderItemName}}/cancel",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-12-01",
		},
		OperationID: "CancelOrderItem",
		Resource:    "Microsoft.EdgeOrder",
	},
	"ReturnOrderItem": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.EdgeOrder/orderItems/{{.orderItemName}}/return",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-12-01",
		},
		OperationID: "ReturnOrderItem",
		Resource:    "Microsoft.EdgeOrder",
	},
}
