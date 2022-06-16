package policy

var Microsoft_ApiManagement_apimnetworkstatus = []Policy{
    {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ApiManagement/service/{{.serviceName}}/networkstatus",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2021-08-01",
    },
	OperationID:    "NetworkStatus_ListByService",
    Resource:       "Microsoft.ApiManagement",
},{
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ApiManagement/service/{{.serviceName}}/locations/{{.locationName}}/networkstatus",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2021-08-01",
    },
	OperationID:    "NetworkStatus_ListByLocation",
    Resource:       "Microsoft.ApiManagement",
},
}
