package policy

// Microsoft_ApiManagement_apimskus policy
var Microsoft_ApiManagement_apimskus = map[string]Policy{
	"ApiManagementSkus_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.ApiManagement/skus",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-08-01",
		},
		OperationID: "ApiManagementSkus_List",
		Resource:    "Microsoft.ApiManagement",
	},
}
