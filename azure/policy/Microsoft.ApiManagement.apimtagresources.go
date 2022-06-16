package policy

var Microsoft_ApiManagement_apimtagresources = []Policy{
    {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ApiManagement/service/{{.serviceName}}/tagResources",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2021-08-01",
    },
	OperationID:    "TagResource_ListByService",
    Resource:       "Microsoft.ApiManagement",
},
}
