package policy

var Microsoft_ApiManagement_apimskus = []Policy{
    {
    Path: "/subscriptions/{{.subscriptionId}}/providers/Microsoft.ApiManagement/skus",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2021-08-01",
    },
	OperationID:    "ApiManagementSkus_List",
    Resource:       "Microsoft.ApiManagement",
},
}
