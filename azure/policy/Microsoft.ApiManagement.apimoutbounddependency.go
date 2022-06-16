package policy

var Microsoft_ApiManagement_apimoutbounddependency = []Policy{
    {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ApiManagement/service/{{.serviceName}}/outboundNetworkDependenciesEndpoints",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2021-08-01",
    },
	OperationID:    "OutboundNetworkDependenciesEndpoints_ListByService",
    Resource:       "Microsoft.ApiManagement",
},
}
