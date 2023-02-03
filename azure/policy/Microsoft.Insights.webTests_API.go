package policy

// Microsoft_Insights_webTests_API policy
var Microsoft_Insights_webTests_API = map[string]Policy{
	"WebTests_ListByResourceGroup": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Insights/webtests",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-06-15",
		},
		OperationID: "WebTests_ListByResourceGroup",
		Resource:    "Microsoft.Insights",
	},
	"WebTests_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Insights/webtests/{{.webTestName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-06-15",
		},
		OperationID: "WebTests_Get",
		Resource:    "Microsoft.Insights",
	},
	"WebTests_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Insights/webtests",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-06-15",
		},
		OperationID: "WebTests_List",
		Resource:    "Microsoft.Insights",
	},
	"WebTests_ListByComponent": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Insights/components/{{.componentName}}/webtests",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-06-15",
		},
		OperationID: "WebTests_ListByComponent",
		Resource:    "Microsoft.Insights",
	},
}
